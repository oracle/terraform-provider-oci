// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserGroupMembershipTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceIdentityUserGroupMembershipTestSuite) SetupTest() {
	token, tokenFn := tokenize()
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t1" {
		name = "{{.token}}"
		description = "tf test user 1"
	}

	resource "oci_identity_user" "t2" {
		name = "{{.token2}}"
		description = "tf test user 2"
	}
	
	resource "oci_identity_group" "t" {
		name = "{{.token}}"
		description = "tf test group"
	}`, map[string]string{"token2": token + "2"})
	s.ResourceName = "oci_identity_user_group_membership.t"
}

func (s *ResourceIdentityUserGroupMembershipTestSuite) TestAccResourceUserGroupMemberships_basic() {
	var resId, resId2 string

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			// Verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_user_group_membership" "t" {
					user_id = "${oci_identity_user.t1.id}"
					group_id = "${oci_identity_group.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "group_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.UserGroupMembershipLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					func(st *terraform.State) (err error) {
						resId, err = fromInstanceState(st, s.ResourceName, "id")
						return err
					},
				),
			},
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_user_group_membership" "t" {
					user_id = "${oci_identity_user.t2.id}"
					group_id = "${oci_identity_group.t.id}"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "group_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(identity.UserGroupMembershipLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckNoResourceAttr(s.ResourceName, "inactive_state"),
					// Verify that changing the user_id causes ForceNew
					func(st *terraform.State) (err error) {
						resId2, err = fromInstanceState(st, s.ResourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
		},
	},
	)
}

func TestResourceIdentityUserGroupMembershipTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserGroupMembershipTestSuite))
}
