// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/bmcs-go-sdk"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityUserGroupMembershipsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceIdentityUserGroupMembershipsTestSuite) SetupTest() {
	_, tokenFn := tokenize()
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
	}
	
	resource "oci_identity_user_group_membership" "t" {
		compartment_id = "${var.tenancy_ocid}"
		user_id = "${oci_identity_user.t.id}"
		group_id = "${oci_identity_group.t.id}"
	}`, nil)
	s.ResourceName = "data.oci_identity_user_group_memberships.t"
}

func (s *DatasourceIdentityUserGroupMembershipsTestSuite) TestAccIdentityUserGroupMemberships_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			// verify import state
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			// verify membership by group
			{
				Config: s.Config + `
				data "oci_identity_user_group_memberships" "t" {
					compartment_id = "${var.tenancy_ocid}"
					group_id = "${oci_identity_group.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "memberships.#", "1"),
				),
			},
			// verify membership by user
			{
				Config: s.Config + `
				data "oci_identity_user_group_memberships" "t" {
					compartment_id = "${var.tenancy_ocid}"
					user_id = "${oci_identity_user.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "memberships.#", "1"),
				),
			},
			//verify membership by specifying both user and group id
			{
				Config: s.Config + `			
				data "oci_identity_user_group_memberships" "t" {
					compartment_id = "${var.tenancy_ocid}"
					user_id = "${oci_identity_user.t.id}"
					group_id = "${oci_identity_group.t.id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "memberships.#", "1"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityUserGroupMembershipsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityUserGroupMembershipsTestSuite))
}
