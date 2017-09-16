// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityPolicyTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
}

func (s *DatasourceIdentityPolicyTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_identity_compartment" "t" {
		name = "-tf-compartment"
		description = "tf test compartment"
	}

	resource "oci_identity_group" "t" {
		name = "-tf-group"
		description = "automated test group"
	}

	resource "oci_identity_policy" "p" {
		name = "-tf-policy"
		description = "automated test policy"
		compartment_id = "${oci_identity_compartment.t.id}"
		statements = ["Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}"]
	}`
	s.ResourceName = "data.oci_identity_policies.p"
}

func (s *DatasourceIdentityPolicyTestSuite) TestAccDatasourceIdentityPolicies_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `	
				data "oci_identity_policies" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "policies.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "policies.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "policies.0.name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "policies.0.description"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "policies.0.statements.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityPoliciesTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityPolicyTestSuite))
}
