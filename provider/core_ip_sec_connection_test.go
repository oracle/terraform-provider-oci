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
resource "oci_core_ip_sec_connection" "test_ip_sec_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"
	static_routes = "${var.ip_sec_connection_static_routes}"
}
`

	IpSecConnectionResourceConfig = IpSecConnectionResourceDependencies + `
resource "oci_core_ip_sec_connection" "test_ip_sec_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"
	static_routes = "${var.ip_sec_connection_static_routes}"

	#Optional
	display_name = "${var.ip_sec_connection_display_name}"
}
`
	IpSecConnectionPropertyVariables = `
variable "ip_sec_connection_display_name" { default = "MyIPSecConnection" }
variable "ip_sec_connection_static_routes" { default = [] }

`
	IpSecConnectionResourceDependencies = CpePropertyVariables + CpeResourceConfig + DrgPropertyVariables + DrgResourceConfig
)

func TestCoreIpSecConnectionResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_ip_sec_connection.test_ip_sec_connection"
	datasourceName := "data.oci_core_ip_sec_connections.test_ip_sec_connections"

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
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "0"),

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
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyIPSecConnection"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "0"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "ip_sec_connection_display_name" { default = "displayName2" }
variable "ip_sec_connection_static_routes" { default = [] }

                ` + compartmentIdVariableStr + IpSecConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cpe_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "0"),

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
variable "ip_sec_connection_display_name" { default = "displayName2" }
variable "ip_sec_connection_static_routes" { default = [] }

data "oci_core_ip_sec_connections" "test_ip_sec_connections" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_ip_sec_connection.test_ip_sec_connection.id}"]
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
					resource.TestCheckResourceAttr(datasourceName, "connections.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.drg_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "connections.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "connections.0.static_routes.#", "0"),
				),
			},
		},
	})
}
