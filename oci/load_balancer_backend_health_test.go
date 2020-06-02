// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	backendHealthSingularDataSourceRepresentation = map[string]interface{}{
		"backend_name":     Representation{repType: Required, create: `${oci_load_balancer_backend.test_backend.name}`},
		"backend_set_name": Representation{repType: Required, create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	BackendHealthResourceConfig = generateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Required, Create, backendRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

func TestLoadBalancerBackendHealthResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerBackendHealthResource_basic")
	defer httpreplay.SaveScenario()

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
					resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_set_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "health_check_results.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				),
			},
		},
	})
}
