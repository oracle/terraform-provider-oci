// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/oci-go-sdk/v65/core"
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

	RouteTableScenarioTestDependencies = CoreVcnResourceConfig + VcnResourceDependencies + ObjectStorageCoreService +
		acctest.GenerateResourceFromRepresentationMap("oci_core_local_peering_gateway", "test_local_peering_gateway", acctest.Required, acctest.Create, CoreLocalPeeringGatewayRepresentation) +
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
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_internet_gateway.test_internet_gateway.id}`},
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`, Update: `10.0.0.0/8`},
	}
	routeTableRouteRulesRepresentationWithServiceCidr = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_service_gateway.test_service_gateway.id}`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `SERVICE_CIDR_BLOCK`},
	}
	routeTableRouteRulesRepresentationWithServiceCidrAddingCidrBlock = map[string]interface{}{
		"network_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_service_gateway.test_service_gateway.id}`},
		"cidr_block":        acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination":       acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination_type":  acctest.Representation{RepType: acctest.Required, Create: `SERVICE_CIDR_BLOCK`},
	}
	routeTableRepresentationWithServiceCidr = acctest.GetUpdatedRepresentationCopy("route_rules", []acctest.RepresentationGroup{
		{RepType: acctest.Required, Group: routeTableRouteRulesRepresentationWithServiceCidr},
		{RepType: acctest.Required, Group: routeTableRouteRulesRepresentationWithCidrBlock}},
		routeTableRepresentationWithRouteRulesReqired,
	)
	routeTableRepresentationWithServiceCidrAddingCidrBlock = acctest.GetUpdatedRepresentationCopy("route_rules", []acctest.RepresentationGroup{
		{RepType: acctest.Required, Group: routeTableRouteRulesRepresentationWithServiceCidrAddingCidrBlock},
		{RepType: acctest.Required, Group: routeTableRouteRulesRepresentationWithCidrBlock}},
		routeTableRepresentationWithRouteRulesReqired,
	)
	routeTableRepresentationWithRouteRulesReqired = acctest.RepresentationCopyWithNewProperties(CoreRouteTableRepresentation, map[string]interface{}{
		"route_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: routeTableRouteRulesRepresentationWithCidrBlock},
	})
)

// We needed to add a lot of special code to handle this case because of the terraform deficiency on differentiating values from statefile and from the config
// We test all the edge cases for that code here.
// issue-routing-tag: core/virtualNetwork
func TestResourceCoreRouteTable_deprecatedCidrBlock(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreRouteTable_deprecatedCidrBlock")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_route_table.test_route_table"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckCoreRouteTableDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, routeTableRepresentationWithRouteRulesReqired),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"cidr_block": "0.0.0.0/0",
				},
					[]string{
						"network_entity_id",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Update to deprecated cidr_block
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Update, routeTableRepresentationWithRouteRulesReqired),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"cidr_block": "10.0.0.0/8"}, []string{"network_entity_id"}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify Update to network_id
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("route_rules.network_entity_id", acctest.Representation{RepType: acctest.Required, Create: `${oci_core_local_peering_gateway.test_local_peering_gateway.id}`},
						routeTableRepresentationWithRouteRulesReqired,
					)),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"cidr_block": "10.0.0.0/8"}, []string{"network_entity_id"}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify Create with destination_type
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, routeTableRepresentationWithServiceCidr),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "0.0.0.0/0"}, []string{"network_entity_id"}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
			),
		},
		// verify Update after having a destination_type rule
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Update, routeTableRepresentationWithServiceCidr),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "10.0.0.0/8"}, []string{"network_entity_id"}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
			),
		},
		// verify adding cidr_block to a rule that has destination already
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Update, routeTableRepresentationWithServiceCidrAddingCidrBlock),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "SERVICE_CIDR_BLOCK"}, []string{"network_entity_id", "destination"}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination_type": "CIDR_BLOCK", "destination": "10.0.0.0/8"}, []string{"network_entity_id"}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
			),
		},
		// We need to test that updating network entity also works when specifying destination instead of cidr_block
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies,
		},
		//Create with optionals and destination
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Optional, acctest.Update, CoreRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destination":      "10.0.0.0/8",
					"destination_type": "CIDR_BLOCK",
				},
					[]string{
						"network_entity_id",
					}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to network entity when using destination
		{
			Config: config + compartmentIdVariableStr + RouteTableScenarioTestDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("route_rules.network_entity_id", acctest.Representation{RepType: acctest.Required, Create: `${oci_core_local_peering_gateway.test_local_peering_gateway.id}`},
						CoreRouteTableRepresentation,
					)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{"destination": "10.0.0.0/8", "destination_type": "CIDR_BLOCK"}, []string{"network_entity_id"}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestResourceCoreRouteTable_defaultResource(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreRouteTable_defaultResource")
	defer httpreplay.SaveScenario()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig() + compartmentIdVariableStr + `
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

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	resourceName := "oci_core_route_table.t"
	defaultResourceName := "oci_core_default_route_table.default"

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify Create without rules
			{
				Config: config + `
					resource "oci_core_route_table" "t" {
						compartment_id = "${var.compartment_id}"
						vcn_id = "${oci_core_virtual_network.t.id}"
					}

					resource "oci_core_default_route_table" "default" {
						manage_default_resource_id = "${oci_core_virtual_network.t.default_route_table_id}"
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttrSet(defaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet(defaultResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(defaultResourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(defaultResourceName, "display_name"),
					resource.TestCheckResourceAttr(defaultResourceName, "route_rules.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(defaultResourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
				),
			},
			// verify Update
			{
				Config: compartmentIdUVariableStr + config + `
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
						compartment_id = "${var.compartment_id_for_update}"
						route_rules {
							cidr_block = "0.0.0.0/0"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
						route_rules {
							cidr_block = "10.0.0.0/8"
							network_entity_id = "${oci_core_internet_gateway.internet-gateway1.id}"
						}
					}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-route-table"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
						"cidr_block": "10.0.0.0/8",
					},
						[]string{
							"network_entity_id",
						}),
					resource.TestCheckResourceAttr(resourceName, "state", string(core.RouteTableLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(defaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttr(defaultResourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(defaultResourceName, "display_name", "default-tf-route-table"),
					resource.TestCheckResourceAttr(defaultResourceName, "route_rules.#", "2"),
					acctest.CheckResourceSetContainsElementWithProperties(defaultResourceName, "route_rules", map[string]string{
						"cidr_block": "0.0.0.0/0",
					},
						[]string{
							"network_entity_id",
						}),
					acctest.CheckResourceSetContainsElementWithProperties(defaultResourceName, "route_rules", map[string]string{
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(defaultResourceName, "manage_default_resource_id"),
					resource.TestCheckResourceAttrSet(defaultResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(defaultResourceName, "compartment_id"),
					resource.TestCheckResourceAttr(defaultResourceName, "route_rules.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(defaultResourceName, "route_rules", map[string]string{
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

func TestResourceCoreRouteTable_resourceDiscovery_crossCompartment(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreDHCPOptions_resourceDiscovery_crossCompartment")
	defer httpreplay.SaveScenario()

	var resId string
	var resId2 string

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	// Config Setup is as follows:
	// ---------------------------
	//  Compartment One                     Compartment Two
	//  -------------------                 -------------------
	//  |                  |                |                  |
	//  |                  |                |                  |
	//  |      VCN 1       |                |       VCN 2      |
	//  |        ^         |                |                  |
	//  |        |         |                |                  |
	//  |        |         |                |                  |
	//  |        |         |                |                  |
	//  |      RT 1        |                |                  |
	//  |                  |                |                  |
	//  |      RT 2        |  ---------->   |                  |
	//  |                  |                |                  |
	//  |                  |                |                  |
	//  |                  |                |                  |
	//  -------------------                 -------------------
	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig() + compartmentIdVariableStr + compartmentIdUVariableStr + `
    	resource "oci_core_virtual_network" "vcn1" {
    		cidr_block = "10.0.0.0/16"
    		compartment_id = "${var.compartment_id}"
    		display_name = "network_name"
    	}

    	resource "oci_core_virtual_network" "vcn2" {
    		cidr_block = "10.0.0.0/16"
    		compartment_id = "${var.compartment_id_for_update}"
    		display_name = "network_name2"
    	}

    	resource "oci_core_route_table" "rt1" {
            compartment_id = "${var.compartment_id}"
            vcn_id = "${oci_core_virtual_network.vcn1.id}"
        }

    	resource "oci_core_route_table" "rt2" {
            compartment_id = "${var.compartment_id}"
            vcn_id = "${oci_core_virtual_network.vcn2.id}"
        }`

	resourceName := "oci_core_route_table.rt1"
	resourceName2 := "oci_core_route_table.rt2"

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, "oci_core_route_table.rt1", "id")
						resId2, err = acctest.FromInstanceState(s, "oci_core_route_table.rt2", "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}

							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId2, &compartmentId, resourceName2); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
		},
	})
}
