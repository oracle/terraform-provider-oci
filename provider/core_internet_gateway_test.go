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
		CheckDestroy: testAccCheckCoreInternetGatewayDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + InternetGatewayPropertyVariables + compartmentIdVariableStr + InternetGatewayRequiredOnlyResource,
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

func testAccCheckCoreInternetGatewayDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_internet_gateway" {
			noResourceFound = false
			request := oci_core.GetInternetGatewayRequest{}

			tmp := rs.Primary.ID
			request.IgId = &tmp

			_, err := client.GetInternetGateway(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
