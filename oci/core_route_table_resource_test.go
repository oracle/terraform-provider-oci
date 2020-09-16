// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/oracle/oci-go-sdk/v25/core"
)

var (
	defaultRouteTable = `
resource "oci_core_default_route_table" "default" {
	manage_default_resource_id = "${oci_core_virtual_network.t.default_route_table_id}"
	route_rules {
		cidr_block = "0.0.0.0/0"
		network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
	}
}
`

	RouteTableScenarioTestDependencies = VcnResourceConfig + VcnResourceDependencies + ObjectStorageCoreService +
		generateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", Required, Create, localPeeringGatewayRepresentation) +
		`
	resource "oci_core_internet_gateway" "test_internet_gateway" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_vcn.id}"
		display_name = "-tf-internet-gateway"
	}

	resource "oci_core_service_gateway" "test_service_gateway" {
		#Required
		compartment_id = "${var.compartment_id}"
		services {
			service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
		}
		vcn_id = "${oci_core_vcn.test_vcn.id}"
	}`

	ObjectStorageCoreService = `data "oci_core_services" "test_services" {
  		filter {
    		name   = "name"
    		values = ["OCI .* Object Storage"]
			regex  = true
  		}
	}
	`

	routeTableRouteRulesRepresentationWithCidrBlock = map[string]interface{}{
		"network_entity_id": Representation{repType: Required, create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"cidr_block":        Representation{repType: Required, create: `0.0.0.0/0`, update: `10.0.0.0/8`},
	}
	routeTableRouteRulesRepresentationWithServiceCidr = map[string]interface{}{
		"network_entity_id": Representation{repType: Required, create: `${oci_core_service_gateway.test_service_gateway.id}`},
		"destination":       Representation{repType: Required, create: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination_type":  Representation{repType: Required, create: `SERVICE_CIDR_BLOCK`},
	}
	routeTableRouteRulesRepresentationWithServiceCidrAddingCidrBlock = map[string]interface{}{
		"network_entity_id": Representation{repType: Required, create: `${oci_core_service_gateway.test_service_gateway.id}`},
		"cidr_block":        Representation{repType: Required, create: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination":       Representation{repType: Required, create: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination_type":  Representation{repType: Required, create: `SERVICE_CIDR_BLOCK`},
	}
	routeTableRepresentationWithServiceCidr = getUpdatedRepresentationCopy("route_rules", []RepresentationGroup{
		{Required, routeTableRouteRulesRepresentationWithServiceCidr},
		{Required, routeTableRouteRulesRepresentationWithCidrBlock}},
		routeTableRepresentationWithRouteRulesReqired,
	)
	routeTableRepresentationWithServiceCidrAddingCidrBlock = getUpdatedRepresentationCopy("route_rules", []RepresentationGroup{
		{Required, routeTableRouteRulesRepresentationWithServiceCidrAddingCidrBlock},
		{Required, routeTableRouteRulesRepresentationWithCidrBlock}},
		routeTableRepresentationWithRouteRulesReqired,
	)
	routeTableRepresentationWithRouteRulesReqired = representationCopyWithNewProperties(routeTableRepresentation, map[string]interface{}{
		"route_rules": RepresentationGroup{Required, routeTableRouteRulesRepresentationWithCidrBlock},
	})
)

// We needed to add a lot of special code to handle this case because of the terraform deficiency on differentiating values from statefile and from the config
// We test all the edge cases for that code here.
func TestResourceCoreRouteTable_deprecatedCidrBlock(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreRouteTable_deprecatedCidrBlock")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_route_table.test_route_table"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreRouteTableDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentationWithRouteRulesReqired),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify update to deprecated cidr_block
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Update, routeTableRepresentationWithRouteRulesReqired),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"cidr_block": "10.0.0.0/8"}, []string{"network_entity_id"}),
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
			// verify update to network_id
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Update,
						getUpdatedRepresentationCopy("route_rules.network_entity_id", Representation{repType: Required, create: `${oci_core_local_peering_gateway.test_local_peering_gateway.id}`},
							routeTableRepresentationWithRouteRulesReqired,
						)),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"cidr_block": "10.0.0.0/8"}, []string{"network_entity_id"}),
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
			// verify create with destination_type
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Create, routeTableRepresentationWithServiceCidr),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "0.0.0.0/0"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				),
			},
			// verify update after having a destination_type rule
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Update, routeTableRepresentationWithServiceCidr),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "10.0.0.0/8"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				),
			},
			// verify adding cidr_block to a rule that has destination already
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Required, Update, routeTableRepresentationWithServiceCidrAddingCidrBlock),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "10.0.0.0/8"}, []string{"network_entity_id"}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				),
			},
			// We need to test that updating network entity also works when specifying destination instead of cidr_block
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies,
			},
			//create with optionals and destination
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Update, routeTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"destination":      "10.0.0.0/8",
						"destination_type": "CIDR_BLOCK",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to network entity when using destination
			{
				Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
					generateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", Optional, Update,
						getUpdatedRepresentationCopy("route_rules.network_entity_id", Representation{repType: Required, create: `${oci_core_local_peering_gateway.test_local_peering_gateway.id}`},
							routeTableRepresentation,
						)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination": "10.0.0.0/8", "destination_type": "CIDR_BLOCK"}, []string{"network_entity_id"}),
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
		},
	})
}

func TestResourceCoreRouteTable_defaultResource(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreRouteTable_defaultResource")
	defer httpreplay.SaveScenario()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	provider := testAccProvider
	config := testProviderConfig() + compartmentIdVariableStr + `
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

	resourceName := "oci_core_route_table.t"
	defaultResourceName := "oci_core_default_route_table.default"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create without rules
			{
				Config: config + `
					resource "oci_core_route_table" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}

					resource "oci_core_default_route_table" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_route_table_id}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "0"),
					resource.TestCheckResourceAttrSet(defaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet(defaultResourceName, "display_name"),
					resource.TestCheckResourceAttr(defaultResourceName, "route_rules.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
				),
			},
			// verify add rule
			{
				Config: config + `
					resource "oci_core_route_table" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
						route_rules {
							cidr_block = "0.0.0.0/0"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
					}` + defaultRouteTable,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(defaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttr(defaultResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(defaultResourceName, "display_name"),
					resource.TestCheckResourceAttr(defaultResourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(defaultResourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
				),
			},
			// verify update
			{
				Config: config + `
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
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-route-table"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block": "10.0.0.0/8",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttr(resourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(defaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttr(defaultResourceName, "display_name", "default-tf-route-table"),
					resource.TestCheckResourceAttr(defaultResourceName, "route_rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(defaultResourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					CheckResourceSetContainsElementWithProperties(defaultResourceName, "route_rules", map[string]string{
						"cidr_block": "10.0.0.0/8",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttr(defaultResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
				),
			},
			// verify default resource delete
			{
				Config: config,
				Check:  nil,
			},
			// verify adding the default resource back to the config
			{
				Config: config + defaultRouteTable,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(defaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet(defaultResourceName, "display_name"),
					resource.TestCheckResourceAttr(defaultResourceName, "route_rules.#", "1"),
					CheckResourceSetContainsElementWithProperties(defaultResourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttr(defaultResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
				),
			},
		},
	})
}
