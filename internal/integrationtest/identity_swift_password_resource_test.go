// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/identity"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentitySwiftPasswordTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

func (s *ResourceIdentitySwiftPasswordTestSuite) SetupTest() {
	_, tokenFn := acctest.TokenizeWithHttpReplay("swiff_pass_resource")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "tf test user"
		compartment_id = "${var.tenancy_ocid}"
	}
  resource "oci_identity_user" "t2" {
		name = "{{.token}}2"
		description = "tf test user 2"
		compartment_id = "${var.tenancy_ocid}"
	}`, nil)

	s.ResourceName = "oci_identity_swift_password.t"
}

func (s *ResourceIdentitySwiftPasswordTestSuite) TestAccResourceIdentitySwiftPassword_basic() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: s.Config + `
				resource "oci_identity_swift_password" "t" {
					user_id = "${oci_identity_user.t.id}"
					description = "tf test swift password"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test swift password"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "expires_on"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.SwiftPasswordLifecycleStateActive)),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_identity_swift_password.t", "id")
						return err
					},
				),
			},
			// verify Update
			{
				Config: s.Config + `
				resource "oci_identity_swift_password" "t" {
					user_id = "${oci_identity_user.t.id}"
					description = "tf test swift password (updated)"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test swift password (updated)"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "expires_on"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.SwiftPasswordLifecycleStateActive)),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, "oci_identity_swift_password.t", "id")
						if resId != resId2 {
							return fmt.Errorf("resource was recreated when it should not have been")
						}
						return err
					},
				),
			},
			// Verify updating user_id causes ForceNew
			{
				Config: s.Config + `
				resource "oci_identity_swift_password" "t" {
					user_id = "${oci_identity_user.t2.id}"
					description = "tf test swift password (user_id updated)"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "tf test swift password (user_id updated)"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "password"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "expires_on"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.SwiftPasswordLifecycleStateActive)),
					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, "oci_identity_swift_password.t", "id")
						if resId == resId2 {
							return fmt.Errorf("resource was updated when it should have been ForceNew")
						}
						return err
					},
				),
			},
		},
	})
}

// issue-routing-tag: identity/default
func TestResourceIdentitySwiftPasswordTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceIdentitySwiftPasswordTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceIdentitySwiftPasswordTestSuite))
}
