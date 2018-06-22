// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	IpSecConnectionRequiredOnlyResource = IpSecConnectionResourceDependencies + `
resource "oci_core_ipsec" "test_ip_sec_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"
	static_routes = "${var.ip_sec_connection_static_routes}"
}
`

	IpSecConnectionResourceConfig = IpSecConnectionResourceDependencies + `
resource "oci_core_ipsec" "test_ip_sec_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"
	static_routes = "${var.ip_sec_connection_static_routes}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.ip_sec_connection_defined_tags_value}")}"
	display_name = "${var.ip_sec_connection_display_name}"
	freeform_tags = "${var.ip_sec_connection_freeform_tags}"
}
`
	IpSecConnectionPropertyVariables = `
variable "ip_sec_connection_defined_tags_value" { default = "value" }
variable "ip_sec_connection_display_name" { default = "MyIPSecConnection" }
variable "ip_sec_connection_freeform_tags" { default = {"Department"= "Finance"} }
variable "ip_sec_connection_static_routes" { default = ["10.0.0.0/16"] }

`
	IpSecConnectionResourceDependencies = DefinedTagsDependencies + CpePropertyVariables + CpeRequiredOnlyResource + DrgPropertyVariables + DrgRequiredOnlyResource
)

func TestCoreIpSecConnectionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ipsec.test_ip_sec_connection"
	datasourceName := "data.oci_core_ipsec_connections.test_ip_sec_connections"

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
				Config:            config + IpSecConnectionPropertyVariables + compartmentIdVariableStr + IpSecConnectionRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + IpSecConnectionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + IpSecConnectionPropertyVariables + compartmentIdVariableStr + IpSecConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnection"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "ip_sec_connection_defined_tags_value" { default = "updatedValue" }
variable "ip_sec_connection_display_name" { default = "displayName2" }
variable "ip_sec_connection_freeform_tags" { default = {"Department"= "Accounting"} }
variable "ip_sec_connection_static_routes" { default = ["10.0.0.0/16"] }

                ` + compartmentIdVariableStr + IpSecConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),

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
variable "ip_sec_connection_defined_tags_value" { default = "updatedValue" }
variable "ip_sec_connection_display_name" { default = "displayName2" }
variable "ip_sec_connection_freeform_tags" { default = {"Department"= "Accounting"} }
variable "ip_sec_connection_static_routes" { default = ["10.0.0.0/32"] }

data "oci_core_ipsec_connections" "test_ip_sec_connections" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_ipsec.test_ip_sec_connection.id}"]
    }
}
                ` + compartmentIdVariableStr + IpSecConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "cpe_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),

					resource.TestCheckResourceAttr(datasourceName, "connections.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.cpe_id"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.drg_id"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.static_routes.#", "1"),
				),
			},
		},
	})
}
