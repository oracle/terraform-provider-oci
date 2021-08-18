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
	loadBalancerProtocolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	LoadBalancerProtocolResourceConfig = ""
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerLoadBalancerProtocolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerLoadBalancerProtocolResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_load_balancer_protocols.test_load_balancer_protocols"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_load_balancer_protocols", "test_load_balancer_protocols", Required, Create, loadBalancerProtocolDataSourceRepresentation) +
				compartmentIdVariableStr + LoadBalancerProtocolResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "protocols.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "protocols.0.name"),
			),
		},
	})
}
