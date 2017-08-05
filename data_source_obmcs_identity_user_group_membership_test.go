// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityUserGroupMembershipsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListUserGroupMemberships
}

func (s *DatasourceIdentityUserGroupMembershipsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}

	s.Config = `
    resource "oci_identity_user" "u" {
		name = "user_name"
		description = "user desc"
    }
    resource "oci_identity_group" "g" {
		name = "group_name"
		description = "group desc"
    }
    resource "oci_identity_user_group_membership" "ug_membership" {
    	compartment_id = "${var.tenancy_ocid}"
		user_id = "${oci_identity_user.u.id}"
		group_id = "${oci_identity_group.g.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "oci_identity_user_group_membership.ug_membership"
}

func (s *DatasourceIdentityUserGroupMembershipsTestSuite) TestGetUserGroupMembershipsByGroup() {
	config := `
	data "oci_identity_user_group_memberships" "g_memberships" {
	    compartment_id = "${var.tenancy_ocid}"
	    group_id = "${oci_identity_group.g.id}"
	}

	data "oci_identity_user_group_memberships" "u_memberships" {
		compartment_id = "${var.tenancy_ocid}"
		user_id = "${oci_identity_user.u.id}"
	}

	data "oci_identity_user_group_memberships" "ug_memberships" {
	    compartment_id = "${var.tenancy_ocid}"
	    user_id = "${oci_identity_user.u.id}"
	    group_id = "${oci_identity_group.g.id}"
	}
	`
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
					resource.TestCheckResourceAttrSet("data.oci_identity_user_group_memberships.g_memberships", "memberships.0.id"),
					resource.TestCheckResourceAttrSet("data.oci_identity_user_group_memberships.u_memberships", "memberships.0.id"),
					resource.TestCheckResourceAttrSet("data.oci_identity_user_group_memberships.ug_memberships", "memberships.0.id"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityUserGroupMembershipsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityUserGroupMembershipsTestSuite))
}
