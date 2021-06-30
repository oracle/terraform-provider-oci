// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	networkLoadBalancersProtocolDataSourceRepresentation = map[string]interface{}{}

	NetworkLoadBalancersProtocolResourceConfig = ""
)

func TestNetworkLoadBalancerNetworkLoadBalancersProtocolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancersProtocolResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	datasourceName := "data.oci_network_load_balancer_network_load_balancers_protocols.test_network_load_balancers_protocols"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_protocols", "test_network_load_balancers_protocols", Required, Create, networkLoadBalancersProtocolDataSourceRepresentation) +
					NetworkLoadBalancersProtocolResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "network_load_balancers_protocol_collection.#"),
					resource.TestCheckResourceAttr(datasourceName, "network_load_balancers_protocol_collection.0.items.#", "3"),
				),
			},
		},
	})
}
