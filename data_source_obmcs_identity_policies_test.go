// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityPolicyTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      TokenFn
}

func (s *DatasourceIdentityPolicyTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + s.TokenFn(`
	resource "oci_identity_compartment" "t" {
		name = "-tf-compartment"
		description = "tf test compartment"
	}

	resource "oci_identity_group" "t" {
		name = "{{.token}}"
		description = "automated test group"
	}

	resource "oci_identity_policy" "p" {
		name = "{{.token}}"
		description = "automated test policy"
		compartment_id = "${oci_identity_compartment.t.id}"
		statements = ["Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}"]
	}`, nil)
	s.ResourceName = "data.oci_identity_policies.p"
}

func (s *DatasourceIdentityPolicyTestSuite) TestAccDatasourceIdentityPolicies_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_identity_policies" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "policies.#"),
				),
			},
			{
				Config: s.Config + s.TokenFn(`
				data "oci_identity_policies" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					filter {
						name   = "name"
						values = ["{{.token}}"]
					}
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "policies.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "policies.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "policies.0.name", s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "policies.0.description", "automated test policy"),
					resource.TestCheckResourceAttr(s.ResourceName, "policies.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(s.ResourceName, "policies.0.inactive_state", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "policies.0.statements.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "policies.0.time_created"),
				),
			},
			// Test filter against array of strings
			{
				Config: s.Config + s.TokenFn(`
				data "oci_identity_policies" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					filter {
						name   = "statements"
						values = ["Allow group {{.token}} to read instances in compartment -tf-compartment"]
					}
				}`, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "policies.#", "1"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityPoliciesTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityPolicyTestSuite))
}
