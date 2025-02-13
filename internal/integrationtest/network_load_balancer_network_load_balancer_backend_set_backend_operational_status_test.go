// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusSingularDataSourceRepresentation = map[string]interface{}{
		"backend_name":             acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend.test_backend.name}`},
		"backend_set_name":         acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, NetworkLoadBalancerBackendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend", "test_backend", acctest.Required, acctest.Create, NetworkLoadBalancerBackendRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation)
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_network_load_balancer_network_load_balancer_backend_set_backend_operational_status.test_network_load_balancer_backend_set_backend_operational_status"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer_backend_set_backend_operational_status", "test_network_load_balancer_backend_set_backend_operational_status", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_set_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_load_balancer_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
			),
		},
	})
}
