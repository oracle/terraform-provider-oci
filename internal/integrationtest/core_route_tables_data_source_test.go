// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreRouteTableTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreRouteTableTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
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
	compartmentID := acctest.GetCompartmentIDForLegacyTests()
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
					data "oci_core_route_tables" "t" {
						compartment_id = "${oci_core_route_table.t.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "1"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.display_name", "Default Route Table for -tf-vcn"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.state", string(core.RouteTableLifecycleStateAvailable)),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.vcn_id", "oci_core_virtual_network.t", "id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.id", "oci_core_virtual_network.t", "default_route_table_id"),
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "1"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.display_name", "oci_core_route_table.t", "display_name"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.0.state", string(core.RouteTableLifecycleStateAvailable)),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.vcn_id", "oci_core_virtual_network.t", "id"),
					acctest.TestCheckResourceAttributesEqual(s.ResourceName, "route_tables.0.id", "oci_core_route_table.t", "id"),
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vcn_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "route_tables.#", "0"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: core/virtualNetwork
func TestDatasourceCoreRouteTableTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreRouteTableTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreRouteTableTestSuite))
}
