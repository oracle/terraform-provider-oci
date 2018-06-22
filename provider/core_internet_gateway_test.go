// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	InternetGatewayRequiredOnlyResource = InternetGatewayResourceDependencies + `
resource "oci_core_internet_gateway" "test_internet_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	enabled = "${var.internet_gateway_enabled}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
`

	InternetGatewayResourceConfig = InternetGatewayResourceDependencies + `
resource "oci_core_internet_gateway" "test_internet_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	enabled = "${var.internet_gateway_enabled}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.internet_gateway_defined_tags_value}")}"
	display_name = "${var.internet_gateway_display_name}"
	freeform_tags = "${var.internet_gateway_freeform_tags}"
}
`
	InternetGatewayPropertyVariables = `
variable "internet_gateway_defined_tags_value" { default = "value" }
variable "internet_gateway_display_name" { default = "MyInternetGateway" }
variable "internet_gateway_enabled" { default = false }
variable "internet_gateway_freeform_tags" { default = {"Department"= "Finance"} }
variable "internet_gateway_state" { default = "AVAILABLE" }

`
	InternetGatewayResourceDependencies = VcnPropertyVariables + VcnResourceConfig
)

func TestCoreInternetGatewayResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_internet_gateway.test_internet_gateway"
	datasourceName := "data.oci_core_internet_gateways.test_internet_gateways"

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
				Config:            config + InternetGatewayPropertyVariables + compartmentIdVariableStr + InternetGatewayRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + InternetGatewayResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + InternetGatewayPropertyVariables + compartmentIdVariableStr + InternetGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyInternetGateway"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
variable "internet_gateway_defined_tags_value" { default = "updatedValue" }
variable "internet_gateway_display_name" { default = "displayName2" }
variable "internet_gateway_enabled" { default = true }
variable "internet_gateway_freeform_tags" { default = {"Department"= "Accounting"} }
variable "internet_gateway_state" { default = "AVAILABLE" }

                ` + compartmentIdVariableStr + InternetGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
variable "internet_gateway_defined_tags_value" { default = "updatedValue" }
variable "internet_gateway_display_name" { default = "displayName2" }
variable "internet_gateway_enabled" { default = true }
variable "internet_gateway_freeform_tags" { default = {"Department"= "Accounting"} }
variable "internet_gateway_state" { default = "AVAILABLE" }

data "oci_core_internet_gateways" "test_internet_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.internet_gateway_display_name}"
	state = "${var.internet_gateway_state}"

    filter {
    	name = "id"
    	values = ["${oci_core_internet_gateway.test_internet_gateway.id}"]
    }
}
                ` + compartmentIdVariableStr + InternetGatewayResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

					resource.TestCheckResourceAttr(datasourceName, "gateways.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.enabled", "true"),
					resource.TestCheckResourceAttr(datasourceName, "gateways.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "gateways.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "gateways.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "gateways.0.vcn_id"),
				),
			},
		},
	})
}
