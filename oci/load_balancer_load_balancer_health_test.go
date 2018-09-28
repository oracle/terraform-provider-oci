// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	loadBalancerHealthSingularDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"depends_on":       Representation{repType: Required, create: []string{`oci_load_balancer_backend.test_backend`}},
	}

	LoadBalancerHealthResourceConfig = BackendRequiredOnlyResource
)

func TestLoadBalancerLoadBalancerHealthResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_health.test_load_balancer_health"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_health", "test_load_balancer_health", Required, Create, loadBalancerHealthSingularDataSourceRepresentation) +
					compartmentIdVariableStr + LoadBalancerHealthResourceConfig,
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
