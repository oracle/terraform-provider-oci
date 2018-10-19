// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	FastConnectProviderServiceResourceConfig = FastConnectProviderServiceResourceDependencies + `

`
	FastConnectProviderServicePropertyVariables = `

`
	FastConnectProviderServiceResourceDependencies = ""
)

func TestCoreFastConnectProviderServiceResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services"
	singularDatasourceName := "data.oci_core_fast_connect_provider_service.test_fast_connect_provider_service"

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
                ` + compartmentIdVariableStr + FastConnectProviderServiceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.private_peering_bgp_management"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.provider_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.provider_service_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.public_peering_bgp_management"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.type"),
				),
			},
			// verify singular datasource
			{
				Config: config + `
data "oci_core_fast_connect_provider_services" "test_fast_connect_provider_services" {
	#Required
	compartment_id = "${var.compartment_id}"
}

data "oci_core_fast_connect_provider_service" "test_fast_connect_provider_service" {
	#Required
	provider_service_id = "${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}"
}
                ` + compartmentIdVariableStr + FastConnectProviderServiceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "provider_service_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_peering_bgp_management"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "provider_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "provider_service_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_peering_bgp_management"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "supported_virtual_circuit_types.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				),
			},
		},
	})
}
