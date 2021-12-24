// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	networkLoadBalancersProtocolDataSourceRepresentation = map[string]interface{}{}

	NetworkLoadBalancersProtocolResourceConfig = ""
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerNetworkLoadBalancersProtocolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancersProtocolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_network_load_balancer_network_load_balancers_protocols.test_network_load_balancers_protocols"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_protocols", "test_network_load_balancers_protocols", acctest.Required, acctest.Create, networkLoadBalancersProtocolDataSourceRepresentation) +
				NetworkLoadBalancersProtocolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_load_balancers_protocol_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "network_load_balancers_protocol_collection.0.items.#", "3"),
			),
		},
	})
}
