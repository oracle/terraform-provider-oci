// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	nlbBackendSetHealthSingularDataSourceRepresentation = map[string]interface{}{
		"backend_set_name":         Representation{repType: Required, create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"network_load_balancer_id": Representation{repType: Required, create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	NlbBackendSetHealthResourceConfig = generateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", Required, Create, nlbBackendSetRepresentation) +
		generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Required, Create, networkLoadBalancerRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation)
)

func TestNetworkLoadBalancerBackendSetHealthResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerBackendSetHealthResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_network_load_balancer_backend_set_health.test_backend_set_health"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_network_load_balancer_backend_set_health", "test_backend_set_health", Required, Create, nlbBackendSetHealthSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NlbBackendSetHealthResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_set_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "network_load_balancer_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "critical_state_backend_names.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "total_backend_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "unknown_state_backend_names.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "warning_state_backend_names.#"),
				),
			},
		},
	})
}
