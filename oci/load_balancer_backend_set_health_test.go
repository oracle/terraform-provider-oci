// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	backendSetHealthSingularDataSourceRepresentation = map[string]interface{}{
		"backend_set_name": Representation{repType: Required, create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"depends_on":       Representation{repType: Required, create: []string{`oci_load_balancer_backend.test_backend`}},
	}

	BackendSetHealthResourceConfig = BackendRequiredOnlyResource
)

func TestLoadBalancerBackendSetHealthResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_backend_set_health.test_backend_set_health"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_backend_set_health", "test_backend_set_health", Required, Create, backendSetHealthSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BackendSetHealthResourceConfig,
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
