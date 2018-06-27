// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	BackendHealthResourceConfig = BackendHealthResourceDependencies + `

`
	BackendHealthPropertyVariables = `
variable "backend_health_backend_name" { default = "backendName" }
variable "backend_health_backend_set_name" { default = "backendSetName" }
variable "backend_health_load_balancer_id" { default = "loadBalancerId" }

`
	BackendHealthResourceDependencies = BackendRequiredOnlyResource + BackendPropertyVariables
)

func TestLoadBalancerBackendHealthResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_backend_health.test_backend_health"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
variable "backend_health_backend_name" { default = "backendName" }
variable "backend_health_backend_set_name" { default = "backendSetName" }
variable "backend_health_load_balancer_id" { default = "loadBalancerId" }

data "oci_load_balancer_backend_health" "test_backend_health" {
	#Required
	backend_name = "${oci_load_balancer_backend.test_backend.name}"
	backend_set_name = "${oci_load_balancer_backend_set.test_backend_set.name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
                ` + compartmentIdVariableStr + BackendHealthResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "backend_name", "10.0.0.3:10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "backend_set_name", "backendSet1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "health_check_results.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				),
			},
		},
	})
}
