// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/client/mocks"
)

type ResourceIdentityUserGroupMembershipsTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListUserGroupMemberships
}

func (s *ResourceIdentityUserGroupMembershipsTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
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
	compartment_id = "cid"
	user_id = "${baremetal_identity_user.u.id}"
	group_id = "${baremetal_identity_group.g.id}"
    }
  `
	s.Config += testProviderConfig
	s.ResourceName = "baremetal_identity_user_group_membership.ug_membership"
	u1 := &baremetal.User{
		ID:            "user_id",
		Name:          "user_name",
		Description:   "user desc",
		CompartmentID: "cid",
		State:         baremetal.ResourceActive,
		TimeCreated:   time.Now(),
	}
	g1 := &baremetal.Group{
		ID:            "group_id",
		Name:          "group_name",
		Description:   "group desc",
		CompartmentID: "cid",
		State:         baremetal.ResourceActive,
		TimeCreated:   time.Now(),
	}
	m1 := &baremetal.UserGroupMembership{
		CompartmentID: "cid",
		GroupID:       g1.ID,
		ID:            "user_group_id",
		State:         baremetal.ResourceActive,
		TimeCreated:   time.Now(),
		UserID:        u1.ID,
	}

	s.List = &baremetal.ListUserGroupMemberships{
		Memberships: []baremetal.UserGroupMembership{*m1},
	}
	s.Client.On("CreateUser", "user_name", "user desc", (*baremetal.RetryTokenOptions)(nil)).
		Return(u1, nil)
	s.Client.On("GetUser", u1.ID).Return(u1, nil)
	s.Client.On("CreateGroup", "group_name", "group desc", (*baremetal.RetryTokenOptions)(nil)).Return(g1, nil)
	s.Client.On("GetGroup", g1.ID).Return(g1, nil)
	s.Client.On("AddUserToGroup", "user_id", "group_id", (*baremetal.RetryTokenOptions)(nil)).Return(m1, nil)
	s.Client.On("ListUserGroupMemberships", (*baremetal.ListMembershipsOptions)(nil)).Return(s.List, nil)
	s.Client.On("GetUserGroupMembership", m1.ID).Return(m1, nil)
	s.Client.On("DeleteUser", "user_id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
	s.Client.On("DeleteGroup", "group_id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
	s.Client.On("DeleteUserGroupMembership", "user_group_id", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceIdentityUserGroupMembershipsTestSuite) TestCreateUserGroupMemberships() {
	fmt.Println(s.Config)
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "id", "user_group_id"),
				),
			},
		},
	},
	)
}

func (s *ResourceIdentityUserGroupMembershipsTestSuite) TestGetUserGroupMembershipsByGroup() {
	config := s.Config
	config += `
	data "baremetal_identity_user_group_memberships" "g_memberships" {
	    compartment_id = "cid"
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
					resource.TestCheckResourceAttr(s.ResourceName, "id", "user_group_id"),
				),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.baremetal_identity_user_group_memberships.g_memberships", "memberships.0.id", "user_group_id"),
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
		compartment_id = "cid"
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
					resource.TestCheckResourceAttr(s.ResourceName, "id", "user_group_id"),
				),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.baremetal_identity_user_group_memberships.u_memberships", "memberships.0.id", "user_group_id"),
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
	    compartment_id = "cid"
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
					resource.TestCheckResourceAttr(s.ResourceName, "id", "user_group_id"),
				),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.baremetal_identity_user_group_memberships.ug_memberships", "memberships.0.id", "user_group_id"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityUserGroupMembershipsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserGroupMembershipsTestSuite))
}
