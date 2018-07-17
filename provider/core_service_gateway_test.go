// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ServiceGatewayRequiredOnlyResource = ServiceGatewayResourceDependencies + `
resource "oci_core_service_gateway" "test_service_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	services {
		#Required
		service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	ServiceGatewayResourceConfig = ServiceGatewayResourceDependencies + `
resource "oci_core_service_gateway" "test_service_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	services {
		service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.service_gateway_defined_tags_value}")}"
	display_name = "${var.service_gateway_display_name}"
	freeform_tags = "${var.service_gateway_freeform_tags}"
}
`
	ServiceGatewayPropertyVariables = `
variable "service_gateway_defined_tags_value" { default = "value" }
variable "service_gateway_display_name" { default = "displayName" }
variable "service_gateway_freeform_tags" { default = {"bar-key"= "value"} }
variable "service_gateway_state" { default = "AVAILABLE" }

`
	ServiceGatewayResourceDependencies = VcnPropertyVariables + VcnResourceConfig + `
data "oci_core_services" "test_services" {
}
`
)

func TestCoreServiceGatewayResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_service_gateway.test_service_gateway"
	datasourceName := "data.oci_core_service_gateways.test_service_gateways"

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
				Config:            config + ServiceGatewayPropertyVariables + compartmentIdVariableStr + ServiceGatewayRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
						[]string{
							"service_id",
						}),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ServiceGatewayResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + ServiceGatewayPropertyVariables + compartmentIdVariableStr + ServiceGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "block_traffic"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
						[]string{
							"service_id",
							"service_name",
						}),
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
variable "service_gateway_defined_tags_value" { default = "updatedValue" }
variable "service_gateway_display_name" { default = "displayName2" }
variable "service_gateway_freeform_tags" { default = {"Department"= "Accounting"} }
variable "service_gateway_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + ServiceGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "block_traffic"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "services.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "services", map[string]string{},
						[]string{
							"service_id",
							"service_name",
						}),
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
			// verify datasource
			{
				Config: config + `
variable "service_gateway_defined_tags_value" { default = "updatedValue" }
variable "service_gateway_display_name" { default = "displayName2" }
variable "service_gateway_freeform_tags" { default = {"Department"= "Accounting"} }
variable "service_gateway_state" { default = "AVAILABLE" }

data "oci_core_service_gateways" "test_service_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	state = "${var.service_gateway_state}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_service_gateway.test_service_gateway.id}"]
    }
}
                ` + compartmentIdVariableStr + ServiceGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "service_gateways.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.block_traffic"),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "service_gateways.0.services.#", "1"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "service_gateways.0.services", map[string]string{},
						[]string{
							"service_id",
							"service_name",
						}),
					resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "service_gateways.0.vcn_id"),
				),
			},
		},
	})
}
