// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	backendHealthSingularDataSourceRepresentation = map[string]interface{}{
		"backend_name":     Representation{RepType: Required, Create: `${oci_load_balancer_backend.test_backend.name}`},
		"backend_set_name": Representation{RepType: Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	BackendHealthResourceConfig = GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", Required, Create, backendSetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", Required, Create, backendRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", Required, Create, certificateRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerBackendHealthResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerBackendHealthResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_backend_health.test_backend_health"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_load_balancer_backend_health", "test_backend_health", Required, Create, backendHealthSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BackendHealthResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_set_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "health_check_results.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
			),
		},
	})
}
