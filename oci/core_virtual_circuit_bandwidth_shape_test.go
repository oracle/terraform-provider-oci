// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	virtualCircuitBandwidthShapeDataSourceRepresentation = map[string]interface{}{
		"provider_service_id": Representation{repType: Required, create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
	}

	VirtualCircuitBandwidthShapeResourceConfig = generateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_services", "test_fast_connect_provider_services", Required, Create, fastConnectProviderServiceDataSourceRepresentation)
)

func TestCoreVirtualCircuitBandwidthShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVirtualCircuitBandwidthShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_virtual_circuit_bandwidth_shapes.test_virtual_circuit_bandwidth_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_virtual_circuit_bandwidth_shapes", "test_virtual_circuit_bandwidth_shapes", Required, Create, virtualCircuitBandwidthShapeDataSourceRepresentation) +
					compartmentIdVariableStr + VirtualCircuitBandwidthShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.0.bandwidth_in_mbps"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.0.name"),
				),
			},
		},
	})
}
