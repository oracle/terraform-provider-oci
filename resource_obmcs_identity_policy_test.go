// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityPolicyTestSuite struct {
	suite.Suite
	Client      mockableClient
	Provider    terraform.ResourceProvider
	Providers   map[string]terraform.ResourceProvider
	TimeCreated time.Time
	Config      string
	PolicyName  string
	Policy      *baremetal.Policy
}

func (s *ResourceIdentityPolicyTestSuite) SetupTest() {
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
		name = "-tf-group"
		description = "automated test group"
	}
	data "baremetal_identity_compartments" "t" {
		compartment_id = "${var.compartment_id}"
	}
	resource "baremetal_identity_policy" "p" {
		name = "-tf-policy"
		description = "automated test policy"
		compartment_id = "${data.baremetal_identity_compartments.t.compartments.0.id}"
		statements = ["Allow group ${baremetal_identity_group.t.name} to read instances in compartment ${data.baremetal_identity_compartments.t.compartments.0.name}"]
	}
	`
	s.Config += testProviderConfig()
	s.PolicyName = "baremetal_identity_policy.p"

}

func (s *ResourceIdentityPolicyTestSuite) TestCreateResourceIdentityPolicy() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.PolicyName, "id"),
					resource.TestCheckResourceAttrSet(s.PolicyName, "time_created"),
					resource.TestCheckResourceAttr(s.PolicyName, "statements.#", "1"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityPolicyTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityPolicyTestSuite))
}
