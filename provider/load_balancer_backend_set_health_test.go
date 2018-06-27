// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	BackendSetHealthResourceConfig = BackendSetHealthResourceDependencies + `

`
	BackendSetHealthPropertyVariables = `
variable "backend_set_health_backend_set_name" { default = "backendSetName" }
variable "backend_set_health_load_balancer_id" { default = "loadBalancerId" }

`
	BackendSetHealthResourceDependencies = BackendRequiredOnlyResource + BackendPropertyVariables
)

func TestLoadBalancerBackendSetHealthResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_backend_set_health.test_backend_set_health"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
variable "backend_set_health_backend_set_name" { default = "backendSetName" }
variable "backend_set_health_load_balancer_id" { default = "loadBalancerId" }

data "oci_load_balancer_backend_set_health" "test_backend_set_health" {
	depends_on = ["oci_load_balancer_backend.test_backend"]

	#Required
	backend_set_name = "${oci_load_balancer_backend_set.test_backend_set.name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
                ` + compartmentIdVariableStr + BackendSetHealthResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "backend_set_name", "backendSet1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "critical_state_backend_names.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
					resource.TestCheckResourceAttr(singularDatasourceName, "total_backend_count", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "unknown_state_backend_names.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "unknown_state_backend_names.0", "10.0.0.3:10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "warning_state_backend_names.#", "0"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
