// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	VirtualCircuitBandwidthShapeResourceConfig = VirtualCircuitBandwidthShapeResourceDependencies + `

`
	VirtualCircuitBandwidthShapePropertyVariables = `

`
	// @CODEGEN 07/2018: ProviderService is actually FastConnectProviderService
	VirtualCircuitBandwidthShapeResourceDependencies = FastConnectProviderServicePropertyVariables + FastConnectProviderServiceResourceConfig
)

func TestCoreVirtualCircuitBandwidthShapeResource_basic(t *testing.T) {
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
				Config: config + `

data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services" {
	#Required
	compartment_id = "${var.compartment_id}"

}

data "oci_core_virtual_circuit_bandwidth_shapes" "test_virtual_circuit_bandwidth_shapes" {
	#Required
	provider_service_id = "${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}"
}
                ` + compartmentIdVariableStr + VirtualCircuitBandwidthShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					//resource.TestCheckResourceAttrSet(datasourceName, "provider_service_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.0.name"),
					//resource.TestCheckResourceAttrSet(datasourceName, "virtual_circuit_bandwidth_shapes.0.provider_service_id"),
				),
			},
		},
	})
}
