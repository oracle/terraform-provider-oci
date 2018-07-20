// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

const (
	BackendRequiredOnlyResource = BackendResourceDependencies + `
resource "oci_load_balancer_backend" "test_backend" {
	#Required
	backendset_name = "${oci_load_balancer_backend_set.test_backend_set.name}"
	ip_address = "${var.backend_ip_address}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	port = "${var.backend_port}"
}
`

	BackendResourceConfig = BackendResourceDependencies + `
resource "oci_load_balancer_backend" "test_backend" {
	#Required
	backendset_name = "${oci_load_balancer_backend_set.test_backend_set.name}"
	ip_address = "${var.backend_ip_address}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	port = "${var.backend_port}"

	#Optional
	backup = "${var.backend_backup}"
	drain = "${var.backend_drain}"
	offline = "${var.backend_offline}"
	weight = "${var.backend_weight}"
}
`
	BackendPropertyVariables = `
variable "backend_backendset_name" { default = "backendSet1" }
variable "backend_backup" { default = false }
variable "backend_drain" { default = false }
variable "backend_ip_address" { default = "10.0.0.3" }
variable "backend_offline" { default = false }
variable "backend_port" { default = 10 }
variable "backend_weight" { default = 10 }

`
	BackendResourceDependencies = BackendSetRequiredOnlyResource + BackendSetPropertyVariables
)

func TestLoadBalancerBackendResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_backend.test_backend"
	datasourceName := "data.oci_load_balancer_backends.test_backends"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerBackendDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + BackendPropertyVariables + compartmentIdVariableStr + BackendRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "backendset_name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BackendResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + BackendPropertyVariables + compartmentIdVariableStr + BackendResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "backendset_name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "backup", "false"),
					resource.TestCheckResourceAttr(resourceName, "drain", "false"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttr(resourceName, "offline", "false"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "weight", "10"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "backend_backendset_name" { default = "backendSet1" }
variable "backend_backup" { default = false }
variable "backend_drain" { default = false }
variable "backend_ip_address" { default = "10.0.0.3" }
variable "backend_offline" { default = false }
variable "backend_port" { default = 10 }
variable "backend_weight" { default = 10 }

                ` + compartmentIdVariableStr + BackendResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "backendset_name", "backendSet1"),
					resource.TestCheckResourceAttr(resourceName, "backup", "false"),
					resource.TestCheckResourceAttr(resourceName, "drain", "false"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttr(resourceName, "offline", "false"),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "weight", "10"),

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
variable "backend_backendset_name" { default = "backendSet1" }
variable "backend_backup" { default = false }
variable "backend_drain" { default = false }
variable "backend_ip_address" { default = "10.0.0.3" }
variable "backend_offline" { default = false }
variable "backend_port" { default = 10 }
variable "backend_weight" { default = 10 }

data "oci_load_balancer_backends" "test_backends" {
	#Required
	backendset_name = "${var.backend_backendset_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
                ` + compartmentIdVariableStr + BackendResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "backendset_name", "backendSet1"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "backends.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.backup", "false"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.drain", "false"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttrSet(datasourceName, "backends.0.name"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.offline", "false"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.port", "10"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.weight", "10"),
				),
			},
		},
	})
}

func testAccCheckLoadBalancerBackendDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_backend" {
			noResourceFound = false
			request := oci_load_balancer.GetBackendRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.BackendName = &value
			}

			if value, ok := rs.Primary.Attributes["backendset_name"]; ok {
				request.BackendSetName = &value
			}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			_, err := client.GetBackend(context.Background(), request)

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
