// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityPolicyTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityPolicyTestSuite) SetupTest() {
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
	}`
	s.ResourceName = "oci_identity_policy.p"
}

func (s *ResourceIdentityPolicyTestSuite) TestAccResourceIdentityPolicy_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_policy" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					name = "-tf-policy"
					description = "automated test policy"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}"]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ETag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "lastUpdateETag"),
					resource.TestCheckResourceAttr(s.ResourceName, "policyHash", "5b4814dcd284428ec85a969f94273fd0"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-policy"),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test policy"),
					resource.TestCheckResourceAttr(s.ResourceName, "statements.#", "1"),
				),
			},
			// verify update
			{
				Config: s.Config + `
				resource "oci_identity_policy" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					name = "-tf-policy-update"
					description = "automated test policy (updated)"
					statements = [
						"Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}",
						"Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}"
					]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "-tf-policy-update"),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test policy (updated)"),
					resource.TestCheckResourceAttr(s.ResourceName, "statements.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "policyHash", "84af348fe943dfb06bf36fbd0903b136"),
				),
			},
		},
	},
	)
}

func (s *ResourceIdentityPolicyTestSuite) TestAccResourceIdentityPolicy_formattingDiff() {
	var lastUpdateETag, policyHash string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// create policy with bad formatting
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_policy" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					name = "-tf-policy"
					description = "automated test policy"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in >> compartment ${oci_identity_compartment.t.name}"]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "statements.0", "Allow group -tf-group to read instances in compartment -tf-compartment"),
					func(s *terraform.State) (err error) {
						if policyHash, err = fromInstanceState(s, "oci_identity_policy.p", "policyHash"); err == nil {
							lastUpdateETag, err = fromInstanceState(s, "oci_identity_policy.p", "lastUpdateETag")
						}
						return err
					},
				),
			},
			// verify update does not change the hash and ETag value
			{
				Config: s.Config + `
				resource "oci_identity_policy" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					name = "-tf-policy"
					description = "automated test policy"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in >> compartment ${oci_identity_compartment.t.name}"]
				}`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) (err error) {
						resource.TestCheckResourceAttr("oci_identity_policy.p", "policyHash", policyHash)
						resource.TestCheckResourceAttr("oci_identity_policy.p", "lastUpdateETag", lastUpdateETag)
						return err
					},
				),
			},
		},
	},
	)
}

func TestResourceIdentityPolicyTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityPolicyTestSuite))
}
