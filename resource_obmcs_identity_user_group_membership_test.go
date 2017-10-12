// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserGroupMembershipTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceIdentityUserGroupMembershipTestSuite) SetupTest() {
	var _, tokenFn = tokenize()
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "tf test user"
	}
	
	resource "oci_identity_group" "t" {
		name = "{{.token}}"
		description = "tf test group"
	}`, nil)
	s.ResourceName = "oci_identity_user_group_membership.t"
}

func (s *ResourceIdentityUserGroupMembershipTestSuite) TestAccResourceUserGroupMemberships_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
				resource "oci_identity_user_group_membership" "t" {
					compartment_id = "${var.tenancy_ocid}"
					user_id = "${oci_identity_user.t.id}"
					group_id = "${oci_identity_group.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "group_id"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityUserGroupMembershipTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserGroupMembershipTestSuite))
}
