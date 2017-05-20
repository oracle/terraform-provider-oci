// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/MustWin/baremetal-sdk-go"
)

var testPoliciesConfig = `
  data "baremetal_identity_policies" "p" {
    compartment_id = "%s"
  }
`

type ResourceIdentityPoliciesTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	PoliciesName string
	Policies     baremetal.ListPolicies
}

func (s *ResourceIdentityPoliciesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	},
	)
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `
	resource "baremetal_identity_group" "t" {
		name = "HelpDesk"
		description = "group desc!"
	}
	data "baremetal_identity_compartments" "t" {
     		compartment_id = "${var.compartment_id}"
        }
	  resource "baremetal_identity_policy" "p" {
	    name = "HelpdeskUsers"
	    description = "description"
	    compartment_id = "${data.baremetal_identity_compartments.t.compartments.0.id}"
	    statements = ["Allow group HelpDesk to read instances in compartment ${data.baremetal_identity_compartments.t.compartments.0.name}"]

	    depends_on = ["baremetal_identity_group.t"]
	  }
	data "baremetal_identity_policies" "p" {
		compartment_id = "${data.baremetal_identity_compartments.t.compartments.0.id}"
	}
	  `
	s.Config += testProviderConfig()
	s.PoliciesName = "data.baremetal_identity_policies.p"
}

func (s *ResourceIdentityPoliciesTestSuite) TestListResourceIdentityPolicies() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.PoliciesName, "policies.0.id"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityPoliciesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityPoliciesTestSuite))
}
