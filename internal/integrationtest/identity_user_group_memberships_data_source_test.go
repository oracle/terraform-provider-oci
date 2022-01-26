// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/identity"
	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityUserGroupMembershipsTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceIdentityUserGroupMembershipsTestSuite) SetupTest() {
	_, tokenFn := acctest.TokenizeWithHttpReplay("identity_user_group_data_source")
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + tokenFn(`
	resource "oci_identity_user" "t" {
		name = "{{.token}}"
		description = "tf test user"
		compartment_id = "${var.tenancy_ocid}"
	}
	
	resource "oci_identity_group" "t" {
		name = "{{.token}}"
		description = "tf test Group"
		compartment_id = "${var.tenancy_ocid}"
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
				Config: s.Config,
			},
			//verify membership by specifying both user and Group id
			{
				Config: s.Config + `			
				data "oci_identity_user_group_memberships" "t" {
					compartment_id = "${var.tenancy_ocid}"
					user_id = "${oci_identity_user.t.id}"
					group_id = "${oci_identity_group.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "memberships.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "memberships.0.state", string(identity.UserGroupMembershipLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "memberships.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "memberships.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "memberships.0.user_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "memberships.0.group_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "memberships.0.time_created"),
					// TODO: This field is not being returned by the service call but is still showing up in the datasource
					// resource.TestCheckNoResourceAttr(s.ResourceName, "memberships.0.inactive_state"),
				),
			},
			// verify membership by Group
			{
				Config: s.Config + `
				data "oci_identity_user_group_memberships" "t" {
					compartment_id = "${var.tenancy_ocid}"
					group_id = "${oci_identity_group.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "memberships.#", "1"),
				),
			},
			// verify filtering
			{
				Config: s.Config + `
				data "oci_identity_user_group_memberships" "t" {
					compartment_id = "${var.tenancy_ocid}"
					group_id = "${oci_identity_group.t.id}"
					filter {
						name = "user_id"
						values = ["${oci_identity_user.t.id}"]
					}
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "memberships.#", "1"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: identity/default
func TestDatasourceIdentityUserGroupMembershipsTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceIdentityUserGroupMembershipsTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceIdentityUserGroupMembershipsTestSuite))
}
