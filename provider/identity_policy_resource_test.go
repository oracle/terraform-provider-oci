// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityPolicyTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
	Token        string
	TokenFn      func(string, map[string]string) string
}

func (s *ResourceIdentityPolicyTestSuite) SetupTest() {
	s.Token, s.TokenFn = tokenize()
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + s.TokenFn(`
	resource "oci_identity_compartment" "t" {
		name = "-tf-compartment"
		description = "tf test compartment"
	}
	
	resource "oci_identity_group" "t" {
		name = "{{.token}}"
		description = "automated test group"
	}`, nil)
	s.ResourceName = "oci_identity_policy.p"
}

func (s *ResourceIdentityPolicyTestSuite) TestAccResourceIdentityPolicy_basic() {
	var policyHash string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					name = "p1-{{.token}}"
					description = "automated test policy"
					version_date = "2018-04-17"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}"]
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ETag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "lastUpdateETag"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "policyHash"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", "p1-"+s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test policy"),
					resource.TestCheckResourceAttr(s.ResourceName, "statements.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.PolicyLifecycleStateActive)),
					resource.TestCheckResourceAttr(s.ResourceName, "version_date", "2018-04-17"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(s *terraform.State) (err error) {
						policyHash, err = fromInstanceState(s, "oci_identity_policy.p", "policyHash")
						return err
					},
				),
			},
			// verify update
			{
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					name = "p2-{{.token}}"
					description = "automated test policy (updated)"
					version_date = "2018-04-18"
					statements = [
						"Allow group ${oci_identity_group.t.name} to inspect instances in compartment ${oci_identity_compartment.t.name}",
						"Allow group ${oci_identity_group.t.name} to read instances in compartment ${oci_identity_compartment.t.name}"
					]
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "p2-"+s.Token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test policy (updated)"),
					resource.TestCheckResourceAttr(s.ResourceName, "version_date", "2018-04-18"),
					resource.TestCheckResourceAttr(s.ResourceName, "statements.#", "2"),
					func(s *terraform.State) (err error) {
						newHash, err := fromInstanceState(s, "oci_identity_policy.p", "policyHash")
						if policyHash == newHash {
							return fmt.Errorf("Expected new hash, got same")
						}
						return err
					},
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
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					name = "{{.token}}"
					description = "automated test policy"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in >> compartment ${oci_identity_compartment.t.name}"]
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "statements.0", "Allow group "+s.Token+" to read instances in compartment -tf-compartment"),
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
				Config: s.Config + s.TokenFn(`
				resource "oci_identity_policy" "p" {
					compartment_id = "${oci_identity_compartment.t.id}"
					name = "{{.token}}"
					description = "automated test policy"
					statements = ["Allow group ${oci_identity_group.t.name} to read instances in >> compartment ${oci_identity_compartment.t.name}"]
				}`, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
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
