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
	NatGatewayRequiredOnlyResource = NatGatewayResourceDependencies + `
resource "oci_core_nat_gateway" "test_nat_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	NatGatewayResourceConfig = NatGatewayResourceDependencies + `
resource "oci_core_nat_gateway" "test_nat_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	block_traffic = "${var.nat_gateway_block_traffic}"
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.nat_gateway_defined_tags_value}")}"
	display_name = "${var.nat_gateway_display_name}"
	freeform_tags = "${var.nat_gateway_freeform_tags}"
}
`
	NatGatewayPropertyVariables = `
variable "nat_gateway_block_traffic" { default = false }
variable "nat_gateway_defined_tags_value" { default = "value" }
variable "nat_gateway_display_name" { default = "displayName" }
variable "nat_gateway_freeform_tags" { default = {"Department"= "Finance"} }
variable "nat_gateway_state" { default = "AVAILABLE" }

`
	NatGatewayResourceDependencies = VcnPropertyVariables + VcnResourceConfig
)

func TestCoreNatGatewayResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_nat_gateway.test_nat_gateway"
	datasourceName := "data.oci_core_nat_gateways.test_nat_gateways"
	singularDatasourceName := "data.oci_core_nat_gateway.test_nat_gateway"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreNatGatewayDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + NatGatewayPropertyVariables + compartmentIdVariableStr + NatGatewayRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + NatGatewayResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + NatGatewayPropertyVariables + compartmentIdVariableStr + NatGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "block_traffic", "false"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "nat_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
variable "nat_gateway_block_traffic" { default = true }
variable "nat_gateway_defined_tags_value" { default = "updatedValue" }
variable "nat_gateway_display_name" { default = "displayName2" }
variable "nat_gateway_freeform_tags" { default = {"Department"= "Accounting"} }
variable "nat_gateway_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + NatGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "block_traffic", "true"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "nat_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
			// verify datasource
			{
				Config: config + `
variable "nat_gateway_block_traffic" { default = true }
variable "nat_gateway_defined_tags_value" { default = "updatedValue" }
variable "nat_gateway_display_name" { default = "displayName2" }
variable "nat_gateway_freeform_tags" { default = {"Department"= "Accounting"} }
variable "nat_gateway_state" { default = "AVAILABLE" }

data "oci_core_nat_gateways" "test_nat_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.nat_gateway_display_name}"
	state = "${var.nat_gateway_state}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_nat_gateway.test_nat_gateway.id}"]
    }
}
                ` + compartmentIdVariableStr + NatGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.block_traffic", "true"),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "nat_gateways.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.nat_ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "nat_gateways.0.vcn_id"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
variable "nat_gateway_block_traffic" { default = true }
variable "nat_gateway_defined_tags_value" { default = "updatedValue" }
variable "nat_gateway_display_name" { default = "displayName2" }
variable "nat_gateway_freeform_tags" { default = {"Department"= "Accounting"} }
variable "nat_gateway_state" { default = "AVAILABLE" }

data "oci_core_nat_gateway" "test_nat_gateway" {
	#Required
	nat_gateway_id = "${oci_core_nat_gateway.test_nat_gateway.id}"
}
                ` + compartmentIdVariableStr + NatGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "nat_gateway_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "block_traffic", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "nat_ip"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + `
variable "nat_gateway_block_traffic" { default = true }
variable "nat_gateway_defined_tags_value" { default = "updatedValue" }
variable "nat_gateway_display_name" { default = "displayName2" }
variable "nat_gateway_freeform_tags" { default = {"Department"= "Accounting"} }
variable "nat_gateway_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + NatGatewayResourceConfig,
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

func testAccCheckCoreNatGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_nat_gateway" {
			noResourceFound = false
			request := oci_core.GetNatGatewayRequest{}

			tmp := rs.Primary.ID
			request.NatGatewayId = &tmp

			response, err := client.GetNatGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.NatGatewayLifecycleStateTerminated): true,
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
