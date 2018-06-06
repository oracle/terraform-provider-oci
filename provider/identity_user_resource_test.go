// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"

	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentityUserTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig()

	s.ResourceName = "oci_identity_user.t"
}

func (s *ResourceIdentityUserTestSuite) TestAccResourceIdentityUser_basic() {
	var resId, resId2 string
	token, tokenFn := tokenize()
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create w/ compartment
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config +
					tokenFn(
						`resource "oci_identity_user" "t" {
							name = "{{.token}}"
							description = "{{.description}}"
							compartment_id = "${var.compartment_id}"
						}`,
						map[string]string{"description": "automated test user"}),
				ExpectError: regexp.MustCompile("Tenant id is not equal to compartment id"),
			},
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config +
					tokenFn(
						`resource "oci_identity_user" "t" {
							name = "{{.token}}"
							description = "{{.description}}"
							compartment_id = "${var.tenancy_ocid}"
						}`,
						map[string]string{"description": "automated test user"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", getRequiredEnvSetting("tenancy_ocid")),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test user"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.UserLifecycleStateActive)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_identity_user.t", "id")
						return err
					},
				),
			},
			// verify create w/o compartment, check that it defaults to tenancy
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config +
					tokenFn(
						`resource "oci_identity_user" "t" {
							name = "{{.token}}"
							description = "{{.description}}"
						}`,
						map[string]string{"description": "automated test user"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", getRequiredEnvSetting("tenancy_ocid")),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test user"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.UserLifecycleStateActive)),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, "oci_identity_user.t", "id")
						return err
					},
				),
			},
			// verify update
			{
				Config: s.Config + tokenFn(
					`resource "oci_identity_user" "t" {
						name = "{{.token}}"
						description = "{{.description}}"
					}`,
					map[string]string{"description": "automated test user (updated)"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test user (updated)"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_identity_user.t", "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						resId = resId2
						return err
					},
				),
			},
			// verify force new update
			{
				Config: s.Config + tokenFn(
					`resource "oci_identity_user" "t" {
						name  = "{{.new_name}}"
						description = "{{.description}}"
					}`,
					map[string]string{"new_name": token + "_new", "description": "automated test user (updated)"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "automated test user (updated)"),
					resource.TestCheckResourceAttr(s.ResourceName, "name", token+"_new"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, "oci_identity_user.t", "id")
						if resId2 == resId {
							return fmt.Errorf("resource expected to be recreated but was not")
						}
						return err
					},
				),
			},
		},
	})
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
