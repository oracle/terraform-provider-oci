// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreRouteTableTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreRouteTableTestSuite) SetupTest() {
	s.Providers = testAccProviders
	s.Config = legacyTestProviderConfig() + `
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		display_name = "-tf-vcn"
		cidr_block = "10.0.0.0/16"
	}
	
	resource "oci_core_route_table" "t" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.t.id}"
		display_name = "-tf-route-table"
	}`

	s.ResourceName = "data.oci_core_route_tables.t"
}

func (s *DatasourceCoreRouteTableTestSuite) TestAccDatasourceRouteTable_basic() {
	compartmentID := getCompartmentIDForLegacyTests()
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config: s.Config + `
					data "oci_core_route_tables" "t" {
						compartment_id = "${oci_core_route_table.t.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "2"),
				),
			},
			// Server-side filtering tests.
			{
				Config: s.Config + `
					data "oci_core_route_tables" "t" {
						compartment_id = "${oci_core_route_table.t.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "Default Route Table for -tf-vcn"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.display_name", "Default Route Table for -tf-vcn"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.state", string(core.RouteTableLifecycleStateAvailable)),
					TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.vcn_id", "oci_core_virtual_network.t", "id"),
					TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.id", "oci_core_virtual_network.t", "default_route_table_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.compartment_id", compartmentID),
					resource.TestCheckResourceAttrSet(s.ResourceName, "route_tables.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.route_rules.#", "0"),
				),
			},
			{
				Config: s.Config + `
					data "oci_core_route_tables" "t" {
						compartment_id = "${oci_core_route_table.t.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						display_name = "-tf-route-table"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "1"),
					TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.display_name", "oci_core_route_table.t", "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.state", string(core.RouteTableLifecycleStateAvailable)),
					TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.vcn_id", "oci_core_virtual_network.t", "id"),
					TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.id", "oci_core_route_table.t", "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.compartment_id", compartmentID),
					resource.TestCheckResourceAttrSet(s.ResourceName, "route_tables.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.route_rules.#", "0"),
				),
			},
			{
				Config: s.Config + `
					data "oci_core_route_tables" "t" {
						compartment_id = "${oci_core_route_table.t.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						state = "` + string(core.RouteTableLifecycleStateProvisioning) + `"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "0"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreRouteTableTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreRouteTableTestSuite))
}
