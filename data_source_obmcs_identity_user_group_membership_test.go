// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/MustWin/baremetal-sdk-go"
)

type ResourceIdentityUserGroupMembershipsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListUserGroupMemberships
}

func (s *ResourceIdentityUserGroupMembershipsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.Config = `
    resource "baremetal_identity_user" "u" {
	name = "user_name"
	description = "user desc"
    }
    resource "baremetal_identity_group" "g" {
	name = "group_name"
	description = "group desc"
    }
    resource "baremetal_identity_user_group_membership" "ug_membership" {
    	compartment_id = "${var.compartment_id}"
	user_id = "${baremetal_identity_user.u.id}"
	group_id = "${baremetal_identity_group.g.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "baremetal_identity_user_group_membership.ug_membership"
}

func (s *ResourceIdentityUserGroupMembershipsTestSuite) TestGetUserGroupMembershipsByGroup() {
	config := s.Config
	config += `
	data "baremetal_identity_user_group_memberships" "g_memberships" {
	    compartment_id = "${var.compartment_id}"
	    group_id = "${baremetal_identity_group.g.id}"
        }`
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				),
			},
			{
				Config: s.Config + config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.baremetal_identity_user_group_memberships.g_memberships", "memberships.0.id"),
				),
			},
		},
	},
	)
}

func (s *ResourceIdentityUserGroupMembershipsTestSuite) TestGetUserGroupMembershipsByUser() {
	config := s.Config
	config += `
	data "baremetal_identity_user_group_memberships" "u_memberships" {
		compartment_id = "${var.compartment_id}"
		user_id = "${baremetal_identity_user.u.id}"
	   }`
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				),
			},
			{
				Config: s.Config + config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.baremetal_identity_user_group_memberships.u_memberships", "memberships.0.id"),
				),
			},
		},
	},
	)
}

func (s *ResourceIdentityUserGroupMembershipsTestSuite) TestGetUserGroupMembershipsByUserAndGroup() {
	config := s.Config
	config += `
	data "baremetal_identity_user_group_memberships" "ug_memberships" {
	    compartment_id = "${var.compartment_id}"
	    user_id = "${baremetal_identity_user.u.id}"
	    group_id = "${baremetal_identity_group.g.id}"
        }`
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				),
			},
			{
				Config: s.Config + config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.baremetal_identity_user_group_memberships.ug_memberships", "memberships.0.id"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityUserGroupMembershipsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserGroupMembershipsTestSuite))
}
