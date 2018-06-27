// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	LoadBalancerHealthResourceConfig = LoadBalancerHealthResourceDependencies + `

`
	LoadBalancerHealthPropertyVariables = `
variable "load_balancer_health_load_balancer_id" { default = "loadBalancerId" }

`
	LoadBalancerHealthResourceDependencies = BackendRequiredOnlyResource + BackendPropertyVariables
)

func TestLoadBalancerLoadBalancerHealthResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_health.test_load_balancer_health"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
variable "load_balancer_health_load_balancer_id" { default = "loadBalancerId" }

data "oci_load_balancer_health" "test_load_balancer_health" {
	depends_on = ["oci_load_balancer_backend.test_backend"]

	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
                ` + compartmentIdVariableStr + LoadBalancerHealthResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "critical_state_backend_set_names.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
					resource.TestCheckResourceAttr(singularDatasourceName, "total_backend_set_count", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "unknown_state_backend_set_names.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "unknown_state_backend_set_names.0", "backendSet1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "warning_state_backend_set_names.#", "0"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
