// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
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
		destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
		destination_type = "SERVICE_CIDR_BLOCK"
		network_entity_id = "${oci_core_service_gateway.test_service_gateway.id}"
	}
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
	RouteTableResourceConfigWithAdditionalRule = RouteTableResourceDependencies + `
resource "oci_core_route_table" "test_route_table" {
	#Required
	compartment_id = "${var.compartment_id}"
	route_rules {
		#Required
		cidr_block = "10.0.0.0/8"
		network_entity_id = "${oci_core_internet_gateway.test_network_entity.id}"
	}
    route_rules {
		#Required
		destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
		destination_type = "SERVICE_CIDR_BLOCK"
		network_entity_id = "${oci_core_service_gateway.test_service_gateway.id}"
	}
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

	resource "oci_core_service_gateway" "test_service_gateway" {
	  #Required
	  compartment_id = "${var.compartment_id}"
	
	  services {
		service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
	  }
	
	  vcn_id = "${oci_core_vcn.test_vcn.id}"
	}
	
	data "oci_core_services" "test_services" {
	  filter {
		name   = "name"
		values = [".*Object.*Storage"]
		regex  = true
	  }
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
				Config: config + RouteTablePropertyVariables + compartmentIdVariableStr + RouteTableRequiredOnlyResource,
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
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.destination_type", "SERVICE_CIDR_BLOCK"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.0.network_entity_id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.1.cidr_block", "0.0.0.0/0"),
					resource.TestCheckResourceAttrSet(resourceName, "route_rules.1.network_entity_id"),
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

                ` + compartmentIdVariableStr + RouteTableResourceConfigWithAdditionalRule,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", "10.0.0.0/8"),
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
					resource.TestCheckResourceAttr(resourceName, "route_rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "route_rules.0.cidr_block", ""),
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
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreRouteTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_route_table" {
			noResourceFound = false
			request := oci_core.GetRouteTableRequest{}

			tmp := rs.Primary.ID
			request.RtId = &tmp

			response, err := client.GetRouteTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.RouteTableLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
