package main

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserTestSuite struct {
	suite.Suite
	Client       *MockClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremtlsdk.IdentityResource
}

func (s *ResourceIdentityUserTestSuite) SetupTest() {
	s.Client = &MockClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `

		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "desc!"
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_identity_user.t"
	s.Res = &baremtlsdk.IdentityResource{
		ID:            "id!",
		Name:          "name!",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         baremtlsdk.ResourceCreated,
		TimeCreated:   s.TimeCreated,
		TimeModified:  s.TimeCreated,
	}
	s.Client.On("CreateUser", "name!", "desc!").Return(s.Res, nil)
	s.Client.On("DeleteUser", "id!").Return(nil)
}

func (s *ResourceIdentityUserTestSuite) TestCreateResourceIdentityUser() {
	s.Client.On("GetUser", "id!").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", s.Res.Name),
					resource.TestCheckResourceAttr(s.ResourceName, "description", s.Res.Description),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.Res.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
					resource.TestCheckResourceAttr(s.ResourceName, "time_modified", s.Res.TimeModified.String()),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestCreateResourceIdentityUserPolling() {
	s.Res.State = baremtlsdk.ResourceCreating
	s.Client.On("GetUser", "id!").Return(s.Res, nil).Once()

	u := *s.Res
	u.State = baremtlsdk.ResourceCreated
	s.Client.On("GetUser", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremtlsdk.ResourceCreated),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestUpdateResourceIdentityUserDescription() {
	s.Client.On("GetUser", "id!").Return(s.Res, nil).Twice()

	c := `

		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "newdesc!"
		}
	`
	c += testProviderConfig

	t := s.TimeCreated.Add(5 * time.Minute)
	u := *s.Res
	u.Description = "newdesc!"
	u.TimeModified = t
	s.Client.On("UpdateUser", "id!", "newdesc!").Return(&u, nil)
	s.Client.On("GetUser", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
					resource.TestCheckResourceAttr(s.ResourceName, "time_modified", t.String()),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestFailedUpdateResourceIdentityUserDescription() {
	s.Client.On("GetUser", "id!").Return(s.Res, nil).Times(3)

	c := `

		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "newdesc!"
		}

	`

	c += testProviderConfig

	s.Client.On("UpdateUser", "id!", "newdesc!").Return(nil, errors.New("FAILED!")).Once()

	t := s.TimeCreated.Add(5 * time.Minute)
	u := *s.Res
	u.Description = "newdesc!"
	u.TimeModified = t
	s.Client.On("UpdateUser", "id!", "newdesc!").Return(&u, nil)
	s.Client.On("GetUser", "id!").Return(&u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config:      c,
				ExpectError: regexp.MustCompile(`FAILED`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "desc!"),
				),
			},
			resource.TestStep{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestUpdateResourceIdentityUserNameShouldCreateNew() {
	s.Client.On("GetUser", "id!").Return(s.Res, nil)

	c := `
		resource "baremetal_identity_user" "t" {
			name = "newname!"
			description = "desc!"
		}
	`

	c += testProviderConfig

	u := *s.Res
	u.ID = "newid!"
	u.Name = "newname!"
	s.Client.On("CreateUser", "newname!", "desc!").Return(&u, nil)
	s.Client.On("GetUser", "newid!").Return(&u, nil)
	s.Client.On("DeleteUser", "newid!").Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "newname!"),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestDeleteResourceIdentityUser() {
	s.Client.On("GetUser", "id!").Return(s.Res, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteUser", "id!")
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
