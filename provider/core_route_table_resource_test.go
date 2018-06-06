// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type ResourceCoreRouteTableTestSuite struct {
	suite.Suite
	Providers           map[string]terraform.ResourceProvider
	Config              string
	ResourceName        string
	DefaultResourceName string
}

var defaultRouteTable = `
resource "oci_core_default_route_table" "default" {
	manage_default_resource_id = "${oci_core_virtual_network.t.default_route_table_id}"
	route_rules {
		cidr_block = "0.0.0.0/0"
		network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
	}
}
`

func (s *ResourceCoreRouteTableTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
		resource "oci_core_virtual_network" "t" {
			compartment_id = "${var.compartment_id}"
			cidr_block = "10.0.0.0/16"
			display_name = "-tf-vcn"
		}

		resource "oci_core_internet_gateway" "internet-gateway1" {
			compartment_id = "${var.compartment_id}"
			vcn_id = "${oci_core_virtual_network.t.id}"
			display_name = "-tf-internet-gateway"
		}`

	s.ResourceName = "oci_core_route_table.t"
	s.DefaultResourceName = "oci_core_default_route_table.default"
}

func (s *ResourceCoreRouteTableTestSuite) TestAccResourceCoreRouteTable_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// verify create without rules
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					resource "oci_core_route_table" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}

					resource "oci_core_default_route_table" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_route_table_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.#", "0"),
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "display_name"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "route_rules.#", "0"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
				),
			},
			// verify add rule
			{
				Config: s.Config + `
					resource "oci_core_route_table" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						route_rules {
							cidr_block = "0.0.0.0/0"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
					}` + defaultRouteTable,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
				),
			},
			// verify update
			{
				Config: s.Config + `
					resource "oci_core_route_table" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "-tf-route-table"
						route_rules {
							cidr_block = "0.0.0.0/0"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
						route_rules {
							cidr_block = "10.0.0.0/8"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
					}
					resource "oci_core_default_route_table" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_route_table_id}"
						display_name = "default-tf-route-table"
						route_rules {
							cidr_block = "0.0.0.0/0"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
						route_rules {
							cidr_block = "10.0.0.0/8"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-route-table"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.#", "2"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_rules.1.cidr_block", "10.0.0.0/8"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "display_name", "default-tf-route-table"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "route_rules.#", "2"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "route_rules.1.cidr_block", "10.0.0.0/8"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
				),
			},
			// verify default resource delete
			{
				Config: s.Config,
				Check:  nil,
			},
			// verify adding the default resource back to the config
			{
				Config: s.Config + defaultRouteTable,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(s.DefaultResourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttr(s.DefaultResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
				),
			},
		},
	})
}

func TestResourceCoreRouteTableTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreRouteTableTestSuite))
}
