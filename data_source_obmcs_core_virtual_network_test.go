// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVirtualNetworkTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVirtualNetworkTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
	data "oci_core_virtual_networks" "t" {
		compartment_id = "${oci_core_virtual_network.t.compartment_id}"
		limit = 1
	}`
	s.ResourceName = "data.oci_core_virtual_networks.t"
}

func (s *DatasourceCoreVirtualNetworkTestSuite) TestAccDatasourceCoreVirtualNetwork_basic() {

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "virtual_networks.0.cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "virtual_networks.#"),
				),
			},
		},
	},
	)

}

func TestDatasourceCoreVirtualNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVirtualNetworkTestSuite))
}
