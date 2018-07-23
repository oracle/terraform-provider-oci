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
	LocalPeeringGatewayRequiredOnlyResource = LocalPeeringGatewayResourceDependencies + `
resource "oci_core_local_peering_gateway" "test_local_peering_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	LocalPeeringGatewayResourceConfig = LocalPeeringGatewayResourceDependencies + `
resource "oci_core_local_peering_gateway" "test_local_peering_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.local_peering_gateway_defined_tags_value}")}"
	display_name = "${var.local_peering_gateway_display_name}"
	freeform_tags = "${var.local_peering_gateway_freeform_tags}"
}
`
	LocalPeeringGatewayPropertyVariables = `
variable "local_peering_gateway_defined_tags_value" { default = "value" }
variable "local_peering_gateway_display_name" { default = "displayName" }
variable "local_peering_gateway_freeform_tags" { default = {"Department"= "Finance"} }

`
	LocalPeeringGatewayResourceDependencies = VcnPropertyVariables + VcnResourceConfig

	secondLocalPeeringGatewayWithPeerId = `
variable "vcn_cidr_block2" { default = "10.1.0.0/16" }
variable "vcn_display_name2" { default = "displayName2" }
variable "vcn_dns_label2" { default = "dnslabel2" }

resource "oci_core_vcn" "test_vcn2" {
	#Required
	cidr_block = "${var.vcn_cidr_block2}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.vcn_display_name2}"
	dns_label = "${var.vcn_dns_label2}"
}

variable "local_peering_gateway_display_name2" { default = "requestor" }

resource "oci_core_local_peering_gateway" "test_local_peering_gateway2" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn2.id}"

	#Optional
	display_name = "${var.local_peering_gateway_display_name2}"
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.local_peering_gateway_defined_tags_value}")}"
	freeform_tags = "${var.local_peering_gateway_freeform_tags}"
	peer_id = "${oci_core_local_peering_gateway.test_local_peering_gateway.id}"
}
`
)

func TestCoreLocalPeeringGatewayResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_local_peering_gateway.test_local_peering_gateway"
	datasourceName := "data.oci_core_local_peering_gateways.test_local_peering_gateways"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreLocalPeeringGatewayDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + LocalPeeringGatewayPropertyVariables + compartmentIdVariableStr + LocalPeeringGatewayRequiredOnlyResource,
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
				Config: config + compartmentIdVariableStr + LocalPeeringGatewayResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + LocalPeeringGatewayPropertyVariables + compartmentIdVariableStr + LocalPeeringGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
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
variable "local_peering_gateway_defined_tags_value" { default = "updatedValue" }
variable "local_peering_gateway_display_name" { default = "displayName2" }
variable "local_peering_gateway_freeform_tags" { default = {"Department"= "Accounting"} }

                ` + compartmentIdVariableStr + LocalPeeringGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
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
variable "local_peering_gateway_defined_tags_value" { default = "updatedValue" }
variable "local_peering_gateway_display_name" { default = "displayName2" }
variable "local_peering_gateway_freeform_tags" { default = {"Department"= "Accounting"} }

data "oci_core_local_peering_gateways" "test_local_peering_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_local_peering_gateway.test_local_peering_gateway.id}"]
    }
}
                ` + compartmentIdVariableStr + LocalPeeringGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "local_peering_gateways.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.peering_status"),
					resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "local_peering_gateways.0.vcn_id"),
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
			// verify connect functionality
			{
				Config: config + `
variable "local_peering_gateway_defined_tags_value" { default = "updatedValue" }
variable "local_peering_gateway_display_name" { default = "displayName2" }
variable "local_peering_gateway_freeform_tags" { default = {"Department"= "Accounting"} }

			` + compartmentIdVariableStr + LocalPeeringGatewayResourceConfig + secondLocalPeeringGatewayWithPeerId,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
					resource.TestCheckResourceAttr(resourceName+"2", "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName+"2", "display_name", "requestor"),
					resource.TestCheckResourceAttrSet(resourceName+"2", "id"),
					resource.TestCheckResourceAttrSet(resourceName+"2", "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName+"2", "peer_id"),
					resource.TestCheckResourceAttr(resourceName+"2", "peering_status", string(oci_core.LocalPeeringGatewayPeeringStatusPeered)),
					resource.TestCheckResourceAttr(resourceName+"2", "state", string(oci_core.LocalPeeringGatewayLifecycleStateAvailable)),
					resource.TestCheckResourceAttrSet(resourceName+"2", "time_created"),
					resource.TestCheckResourceAttrSet(resourceName+"2", "vcn_id"),
				),
			},
		},
	})
}

func testAccCheckCoreLocalPeeringGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_local_peering_gateway" {
			noResourceFound = false
			request := oci_core.GetLocalPeeringGatewayRequest{}

			tmp := rs.Primary.ID
			request.LocalPeeringGatewayId = &tmp

			response, err := client.GetLocalPeeringGateway(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.LocalPeeringGatewayLifecycleStateTerminated): true,
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
