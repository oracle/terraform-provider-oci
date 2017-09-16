// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceIdentityGroupsTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListGroups
}

func (s *DatasourceIdentityGroupsTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_identity_group" "t" {
		name = "-tf-group"
		description = "automated test group"
	}`
	s.ResourceName = "data.oci_identity_groups.t"
}

func (s *DatasourceIdentityGroupsTestSuite) TestAccDatasourceIdentityGroups_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				data "oci_identity_groups" "t" {
					compartment_id = "${var.compartment_id}"
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "groups.0.id"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityGroupsTestSuite))
}
