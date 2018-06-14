// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	RemotePeeringConnectionRequiredOnlyResource = RemotePeeringConnectionResourceDependencies + `
resource "oci_core_remote_peering_connection" "test_remote_peering_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	drg_id = "${oci_core_drg.test_drg.id}"
}
`

	RemotePeeringConnectionResourceConfig = RemotePeeringConnectionResourceDependencies + `
resource "oci_core_remote_peering_connection" "test_remote_peering_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	drg_id = "${oci_core_drg.test_drg.id}"

	#Optional
	display_name = "${var.remote_peering_connection_display_name}"
}
`
	RemotePeeringConnectionPropertyVariables = `
variable "remote_peering_connection_display_name" { default = "displayName" }

`
	RemotePeeringConnectionResourceDependencies = DrgPropertyVariables + DrgResourceConfig
)

func TestCoreRemotePeeringConnectionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_remote_peering_connection.test_remote_peering_connection"
	datasourceName := "data.oci_core_remote_peering_connections.test_remote_peering_connections"

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
				Config:            config + RemotePeeringConnectionPropertyVariables + compartmentIdVariableStr + RemotePeeringConnectionRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RemotePeeringConnectionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + RemotePeeringConnectionPropertyVariables + compartmentIdVariableStr + RemotePeeringConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "remote_peering_connection_display_name" { default = "displayName2" }

                ` + compartmentIdVariableStr + RemotePeeringConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
variable "remote_peering_connection_display_name" { default = "displayName2" }

data "oci_core_remote_peering_connections" "test_remote_peering_connections" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	drg_id = "${oci_core_drg.test_drg.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_remote_peering_connection.test_remote_peering_connection.id}"]
    }
}
                ` + compartmentIdVariableStr + RemotePeeringConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),

					resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.drg_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.peering_status"),
					resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.time_created"),
				),
			},
		},
	})
}
