// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	RouteTableRequiredOnlyResource = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		cidr_block = "${var.route_table_route_rules_cidr_block}"
		network_entity_id = "${oci_core_internet_gateway.test_network_entity.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	RouteTableResourceConfig = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		cidr_block = "${var.route_table_route_rules_cidr_block}"
		network_entity_id = "${oci_core_internet_gateway.test_network_entity.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.route_table_display_name}"
}
`
	RouteTablePropertyVariables = `
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_route_rules_cidr_block" { default = "0.0.0.0/0" }
variable "route_table_state" { default = "state" }

`
	RouteTableResourceDependencies = VcnPropertyVariables + VcnResourceConfig + `
	resource "oci_core_internet_gateway" "test_network_entity" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_vcn.id}"
		display_name = "-tf-internet-gateway"
	}
	`
)

func TestCoreRouteTableResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_route_table.test_route_table"
	datasourceName := "data.oci_core_route_tables.test_route_tables"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + RouteTablePropertyVariables + compartmentIdVariableStr + RouteTableRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RouteTableResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + RouteTablePropertyVariables + compartmentIdVariableStr + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyRouteTable"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "route_table_display_name" { default = "displayName2" }
variable "route_table_route_rules_cidr_block" { default = "0.0.0.0/0" }
variable "route_table_state" { default = "state" }

                ` + compartmentIdVariableStr + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "route_table_display_name" { default = "displayName2" }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr2 + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", "10.0.0.0/8"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "route_table_display_name" { default = "displayName2" }
variable "route_table_route_rules_cidr_block" { default = "10.0.0.0/8" }
variable "route_table_state" { default = "AVAILABLE" }

data "oci_core_route_tables" "test_route_tables" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.route_table_display_name}"
	state = "${var.route_table_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_route_table.test_route_table.id}"]
    }
}
                ` + compartmentIdVariableStr2 + RouteTableResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "route_tables.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.route_rules.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "route_tables.0.route_rules.0.cidr_block", "10.0.0.0/8"),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "route_tables.0.vcn_id"),
				),
			},
		},
	})
}

func TestCoreRouteTableResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_route_table.test_route_table"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + RouteTablePropertyVariables + compartmentIdVariableStr + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyRouteTable"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_route_rules_cidr_block" { default = "0.0.0.0/0" }
variable "route_table_state" { default = "state" }
				` + compartmentIdVariableStr2 + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyRouteTable"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter CompartmentId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
variable "route_table_display_name" { default = "MyRouteTable" }
variable "route_table_route_rules_cidr_block" { default = "0.0.0.0/0" }
variable "route_table_state" { default = "state" }
				` + compartmentIdVariableStr2 + RouteTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyRouteTable"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter VcnId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
