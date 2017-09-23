// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreRouteTableTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreRouteTableTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		cidr_block = "10.0.0.0/16"
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
	}
	resource "oci_core_internet_gateway" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-internet_gateway"
		vcn_id = "${oci_core_virtual_network.t.id}"
	}
	resource "oci_core_route_table" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "display_name"
		route_rules {
			cidr_block = "0.0.0.0/0"
			network_entity_id = "${oci_core_internet_gateway.t.id}"
		}
		vcn_id = "${oci_core_virtual_network.t.id}"
	}`

	s.ResourceName = "data.oci_core_route_tables.t"
}


func (s *DatasourceCoreRouteTableTestSuite) TestAccDatasourceRouteTable_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config + `
				data "oci_core_route_tables" "t" {
					compartment_id = "${var.compartment_id}"
					vcn_id = "${oci_core_virtual_network.t.id}"
				}`,
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

func TestDatasourceCoreRouteTableTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreRouteTableTestSuite))
}
