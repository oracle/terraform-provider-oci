// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
variable "backend_backendset_name" { default = "backendsetName" }
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
	t.Skip("Skipping generated test for now as it has not been worked on.")
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
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + BackendPropertyVariables + compartmentIdVariableStr + BackendRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "backendset_name", "backendsetName"),
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
					resource.TestCheckResourceAttr(resourceName, "backendset_name", "backendsetName"),
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
variable "backend_backendset_name" { default = "backendsetName" }
variable "backend_backup" { default = false }
variable "backend_drain" { default = false }
variable "backend_ip_address" { default = "10.0.0.3" }
variable "backend_offline" { default = false }
variable "backend_port" { default = 10 }
variable "backend_weight" { default = 10 }

                ` + compartmentIdVariableStr + BackendResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "backendset_name", "backendsetName"),
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
variable "backend_backendset_name" { default = "backendsetName" }
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

    filter {
    	name = "id"
    	values = ["${oci_load_balancer_backend.test_backend.id}"]
    }
}
                ` + compartmentIdVariableStr + BackendResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "backendset_name", "backendsetName"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "backends.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.backup", "false"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.drain", "false"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttrSet(datasourceName, "backends.0.load_balancer_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "backends.0.name"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.offline", "false"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.port", "10"),
					resource.TestCheckResourceAttr(datasourceName, "backends.0.weight", "10"),
				),
			},
		},
	})
}
