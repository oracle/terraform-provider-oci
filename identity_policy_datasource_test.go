// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

var testPoliciesConfig = `
  data "baremetal_identity_policies" "p" {
    compartment_id = "%s"
  }
`

type ResourceIdentityPoliciesTestSuite struct {
	suite.Suite
	Client      *mocks.BareMetalClient
	Provider    terraform.ResourceProvider
	Providers   map[string]terraform.ResourceProvider
	TimeCreated time.Time
	Config      string
	PoliciesName  string
	Policies      baremetal.ListPolicies
}

func (s *ResourceIdentityPoliciesTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	},
	)
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = fmt.Sprintf(testProviderConfig+testPoliciesConfig, "7")
	s.PoliciesName = "data.baremetal_identity_policies.p"
	s.Policies = baremetal.ListPolicies{
		Policies: []baremetal.Policy{
			baremetal.Policy{
				ID:            "123",
				Name:          "pol",
				Description:   "desc",
				CompartmentID: "7",
				State:         baremetal.ResourceActive,
				TimeCreated:   s.TimeCreated,
				Statements:    []string{"statementX", "statementY"},
			},
			baremetal.Policy{
				ID:            "234",
				Name:          "pol2",
				Description:   "desc2",
				CompartmentID: "7",
				State:         baremetal.ResourceActive,
				TimeCreated:   s.TimeCreated,
				Statements:    []string{"statementY", "statementZ"},
			},
		},
	}

	s.Client.On(
		"ListPolicies",
		"7",
		(*baremetal.ListOptions)(nil),
	).Return(&s.Policies, nil)

}

func (s *ResourceIdentityPoliciesTestSuite) TestListResourceIdentityPolicies() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.PoliciesName, "policies.0.id", s.Policies.Policies[0].ID),
					resource.TestCheckResourceAttr(s.PoliciesName, "policies.0.statements.1", s.Policies.Policies[0].Statements[1]),
				),
			},
		},
	},
	)
}

func TestResourceIdentityPoliciesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityPoliciesTestSuite))
}
