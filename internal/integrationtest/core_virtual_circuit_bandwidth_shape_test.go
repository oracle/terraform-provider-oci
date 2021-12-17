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
	virtualCircuitBandwidthShapeDataSourceRepresentation = map[string]interface{}{
		"provider_service_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
	}

	VirtualCircuitBandwidthShapeResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_services", "test_fast_connect_provider_services", acctest.Required, acctest.Create, fastConnectProviderServiceDataSourceRepresentation)
)

// issue-routing-tag: core/default
func TestCoreVirtualCircuitBandwidthShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVirtualCircuitBandwidthShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_virtual_circuit_bandwidth_shapes", "test_virtual_circuit_bandwidth_shapes", acctest.Required, acctest.Create, virtualCircuitBandwidthShapeDataSourceRepresentation) +
				compartmentIdVariableStr + VirtualCircuitBandwidthShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.0.bandwidth_in_mbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.0.name"),
			),
		},
	})
}
