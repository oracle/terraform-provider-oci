// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	networkLoadBalancerHealthSingularDataSourceRepresentation = map[string]interface{}{
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	NetworkLoadBalancerHealthResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, networkLoadBalancerRepresentation)
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerNetworkLoadBalancerHealthResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancerHealthResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_network_load_balancer_network_load_balancer_health.test_network_load_balancer_health"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer_health", "test_network_load_balancer_health", acctest.Required, acctest.Create, networkLoadBalancerHealthSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkLoadBalancerHealthResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_load_balancer_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "critical_state_backend_set_names.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_backend_set_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "unknown_state_backend_set_names.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "warning_state_backend_set_names.#", "0"),
			),
		},
	})
}
