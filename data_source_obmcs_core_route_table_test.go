// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreRouteTablesTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreRouteTablesTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = `
resource "oci_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
}
resource "oci_core_internet_gateway" "CompleteIG" {
    compartment_id = "${var.compartment_id}"
    display_name = "CompleteIG"
    vcn_id = "${oci_core_virtual_network.t.id}"
}
resource "oci_core_route_table" "t" {
	compartment_id = "${var.compartment_id}"
	display_name = "display_name"
	route_rules {
		cidr_block = "0.0.0.0/0"
		network_entity_id = "${oci_core_internet_gateway.CompleteIG.id}"
	}
	vcn_id = "${oci_core_virtual_network.t.id}"
}
    data "oci_core_route_tables" "t" {
      compartment_id = "${var.compartment_id}"
      vcn_id = "${oci_core_virtual_network.t.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.oci_core_route_tables.t"

}

func (s *ResourceCoreRouteTablesTestSuite) TestResourceListRouteTables() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "route_tables.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "1"),
				),
			},
		},
	},
	)
}

func TestResourceCoreRouteTablesTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreRouteTablesTestSuite))
}
