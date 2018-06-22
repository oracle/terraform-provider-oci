// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + LocalPeeringGatewayPropertyVariables + compartmentIdVariableStr + LocalPeeringGatewayRequiredOnlyResource,
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
