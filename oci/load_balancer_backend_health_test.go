// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	backendHealthSingularDataSourceRepresentation = map[string]interface{}{
		"backend_name":     Representation{repType: Required, create: `${oci_load_balancer_backend.test_backend.name}`},
		"backend_set_name": Representation{repType: Required, create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	BackendHealthResourceConfig = BackendRequiredOnlyResource
)

func TestLoadBalancerBackendHealthResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_backend_health.test_backend_health"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_backend_health", "test_backend_health", Required, Create, backendHealthSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BackendHealthResourceConfig,
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
