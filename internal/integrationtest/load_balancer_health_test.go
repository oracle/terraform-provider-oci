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
	loadBalancerHealthSingularDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"depends_on":       acctest.Representation{RepType: acctest.Required, Create: []string{`oci_load_balancer_backend.test_backend`}},
	}

	LoadBalancerHealthResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", acctest.Required, acctest.Create, backendRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerLoadBalancerHealthResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerLoadBalancerHealthResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_load_balancer_health.test_load_balancer_health"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_health", "test_load_balancer_health", acctest.Required, acctest.Create, loadBalancerHealthSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LoadBalancerHealthResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "critical_state_backend_set_names.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_backend_set_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "unknown_state_backend_set_names.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "warning_state_backend_set_names.#"),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
