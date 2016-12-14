package main

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

var testPolicyConfig = `
  resource "baremetal_identity_policy" "p" {
    name = "%s"
    description = "%s"
    statements = %s
  }
`

// Custom assertion for TypeList attributes.
func testCheckAttributeTypeList(resourceName, attributeName string, expecteds []string) resource.TestCheckFunc {
	return func(s *terraform.State) (e error) {
		ms := s.RootModule()

		rs, ok := ms.Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		is := rs.Primary
		if is == nil {
			return fmt.Errorf("No primary instance: %s", resourceName)
		}

		for i, expected := range expecteds {
			// Keys for individual list elements are represented in Terraform as
			// key.0 key.1 ... key.(N-1)
			key := fmt.Sprintf("%s.%d", attributeName, i)
			actual := is.Attributes[key]
			if actual != expected {
				return fmt.Errorf("Expected '%s' but got '%s'", actual, expected)
			}
		}

		return
	}
}

type ResourceIdentityPolicyTestSuite struct {
	suite.Suite
	Client      *mocks.BareMetalClient
	Provider    terraform.ResourceProvider
	Providers   map[string]terraform.ResourceProvider
	TimeCreated time.Time
	Config      string
	PolicyName  string
	Policy      *baremetal.Policy
}

func (s *ResourceIdentityPolicyTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	},
	)
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = fmt.Sprintf(testProviderConfig+testPolicyConfig,
		"pol",
		"desc",
		`["statementX","statementY"]`,
	)
	s.PolicyName = "baremetal_identity_policy.p"
	s.Policy = &baremetal.Policy{
		ID:            "123",
		Name:          "pol",
		Description:   "desc",
		CompartmentID: "7",
		State:         baremetal.ResourceActive,
		TimeCreated:   s.TimeCreated,
		Statements:    []string{"statementX", "statementY"},
	}

	s.Client.On(
		"CreatePolicy",
		"pol",
		"desc",
		[]string{"statementX", "statementY"},
		(*baremetal.CreatePolicyOptions)(nil),
	).Return(s.Policy, nil)

}

func (s *ResourceIdentityPolicyTestSuite) TestCreateResourceIdentityPolicy() {
	s.Client.On("GetPolicy", s.Policy.ID).Return(s.Policy, nil)
	s.Client.On("DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.PolicyName, "name", s.Policy.Name),
					resource.TestCheckResourceAttr(s.PolicyName, "description", s.Policy.Description),
					resource.TestCheckResourceAttr(s.PolicyName, "compartment_id", s.Policy.CompartmentID),
					resource.TestCheckResourceAttr(s.PolicyName, "state", s.Policy.State),
					resource.TestCheckResourceAttr(s.PolicyName, "time_created", s.Policy.TimeCreated.String()),
					testCheckAttributeTypeList(s.PolicyName, "statements", s.Policy.Statements),
				),
			},
		},
	},
	)
}

func (s *ResourceIdentityPolicyTestSuite) TestUpdateResourceIdentityPolicy() {
	s.Client.On("GetPolicy", s.Policy.ID).Return(s.Policy, nil).Twice()

	s.Client.On("DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	config := fmt.Sprintf(testProviderConfig+testPolicyConfig,
		"pol",
		"newdesc",
		`["statementA","statementY", "statementZ"]`,
	)

	updated := *s.Policy
	updated.Description = "newdesc"
	updated.Statements = []string{"statementA", "statementY", "statementZ"}
	opts := &baremetal.UpdatePolicyOptions{}
	opts.Description = updated.Description
	opts.Statements = updated.Statements
	s.Client.On(
		"UpdatePolicy",
		updated.ID,
		opts,
	).Return(&updated, nil)

	s.Client.On("GetPolicy", updated.ID).Return(&updated, nil)

	resource.UnitTest(s.T(),
		resource.TestCase{
			Providers: s.Providers,
			Steps: []resource.TestStep{
				{
					Config: s.Config,
				},
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(s.PolicyName, "description", updated.Description),
						testCheckAttributeTypeList(s.PolicyName, "statements", updated.Statements),
					),
				},
			},
		},
	)

}

func (s *ResourceIdentityPolicyTestSuite) TestFailedUpdateResourceIdentityPolicy() {
	s.Client.On("GetPolicy", s.Policy.ID).Return(s.Policy, nil).Times(3)
	s.Client.On("DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	config := fmt.Sprintf(testProviderConfig+testPolicyConfig,
		"pol",
		"newdesc",
		`["statementA", "statementB"]`,
	)

	newStatements := []string{"statementA", "statementB"}
	opts := &baremetal.UpdatePolicyOptions{}
	opts.Description = "newdesc"
	opts.Statements = newStatements
	s.Client.On("UpdatePolicy",
		s.Policy.ID,
		opts,
	).Return(nil, errors.New("FAILED")).Once()

	u := *s.Policy
	u.Description = "newdesc"
	u.Statements = newStatements
	s.Client.On("UpdatePolicy",
		s.Policy.ID,
		opts,
	).Return(&u, nil)
	s.Client.On("GetPolicy", s.Policy.ID).Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config:      config,
				ExpectError: regexp.MustCompile(`FAILED`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.PolicyName, "description", "desc"),
					testCheckAttributeTypeList(s.PolicyName, "statements", s.Policy.Statements),
				),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.PolicyName, "description", "newdesc"),
					testCheckAttributeTypeList(s.PolicyName, "statements", u.Statements),
				),
			},
		},
	})
}

func (s *ResourceIdentityPolicyTestSuite) TestUpdateResourceIdentityPolicyNameShouldCreateNew() {
	s.Client.On("GetPolicy", s.Policy.ID).Return(s.Policy, nil)
	s.Client.On("DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	config := fmt.Sprintf(testProviderConfig+testPolicyConfig,
		"newname",
		"desc",
		`["statementX", "statementY"]`,
	)

	statements := []string{"statementX", "statementY"}

	u := *s.Policy
	u.ID = "999"
	u.Name = "newname"
	s.Client.On(
		"CreatePolicy",
		"newname",
		"desc",
		statements,
		(*baremetal.CreatePolicyOptions)(nil),
	).Return(&u, nil)
	s.Client.On("GetPolicy", "999").Return(&u, nil)
	s.Client.On("DeletePolicy", "999", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config: config,
				Check:  resource.TestCheckResourceAttr(s.PolicyName, "name", "newname"),
			},
		},
	})
}

func (s *ResourceIdentityPolicyTestSuite) TestDeleteResourceIdentityPolicy() {
	s.Client.On("GetPolicy", "123").Return(s.Policy, nil)
	s.Client.On("DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil))
}

func (s *ResourceIdentityPolicyTestSuite) TestDeleteFailureResourceIdentityPolicy() {
	s.Client.On("GetPolicy", "123").Return(s.Policy, nil)
	s.Client.On("DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil)).
		Return(errors.New("XXX")).Once()
	s.Client.On("DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil)).Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			{
				Config:      s.Config,
				ExpectError: regexp.MustCompile(`XXX`),
				Destroy:     true,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeletePolicy", "123", (*baremetal.IfMatchOptions)(nil))
}

func TestResourceIdentityPolicyTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityPolicyTestSuite))
}
