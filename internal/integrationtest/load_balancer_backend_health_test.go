// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	backendHealthSingularDataSourceRepresentation = map[string]interface{}{
		"backend_name":     acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend.test_backend.name}`},
		"backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
	}

	BackendHealthResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", acctest.Required, acctest.Create, backendRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerBackendHealthResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerBackendHealthResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_backend_health.test_backend_health"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_backend_health", "test_backend_health", acctest.Required, acctest.Create, backendHealthSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BackendHealthResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_set_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "health_check_results.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
			),
		},
	})
}
