// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserGroupMembershipTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceIdentityUserGroupMembershipTestSuite) SetupTest() {
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
    	compartment_id = "${var.tenancy_ocid}"
	user_id = "${baremetal_identity_user.u.id}"
	group_id = "${baremetal_identity_group.g.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "baremetal_identity_user_group_membership.ug_membership"
}

func (s *ResourceIdentityUserGroupMembershipTestSuite) TestGetUserGroupMembershipsByGroup() {
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
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("baremetal_identity_user_group_membership.ug_membership", "user_id"),
				),
			},
		},
	},
	)
}

func TestResourceIdentityUserGroupMembershipTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserGroupMembershipTestSuite))
}
