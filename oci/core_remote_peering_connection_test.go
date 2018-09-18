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

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_remote_peering_connection.test_remote_peering_connection"
	datasourceName := "data.oci_core_remote_peering_connections.test_remote_peering_connections"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreRemotePeeringConnectionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + RemotePeeringConnectionPropertyVariables + compartmentIdVariableStr + RemotePeeringConnectionRequiredOnlyResource,
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

func testAccCheckCoreRemotePeeringConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_remote_peering_connection" {
			noResourceFound = false
			request := oci_core.GetRemotePeeringConnectionRequest{}

			tmp := rs.Primary.ID
			request.RemotePeeringConnectionId = &tmp

			response, err := client.GetRemotePeeringConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.RemotePeeringConnectionLifecycleStateTerminated): true,
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
