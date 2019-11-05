// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	fastConnectProviderServiceSingularDataSourceRepresentation = map[string]interface{}{
		"provider_service_id": Representation{repType: Required, create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
	}

	fastConnectProviderServiceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	FastConnectProviderServiceResourceConfig = ""
)

func TestCoreFastConnectProviderServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreFastConnectProviderServiceResource_basic")
	defer httpreplay.SaveScenario()

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
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_services", "test_fast_connect_provider_services", Required, Create, fastConnectProviderServiceDataSourceRepresentation) +
					compartmentIdVariableStr + FastConnectProviderServiceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.bandwith_shape_management"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.customer_asn_management"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.private_peering_bgp_management"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.provider_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.provider_service_key_management"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.provider_service_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.public_peering_bgp_management"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.required_total_cross_connects"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.supported_virtual_circuit_types.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "fast_connect_provider_services.0.type"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_services", "test_fast_connect_provider_services", Required, Create, fastConnectProviderServiceDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_service", "test_fast_connect_provider_service", Required, Create, fastConnectProviderServiceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + FastConnectProviderServiceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "provider_service_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "bandwith_shape_management"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "customer_asn_management"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_peering_bgp_management"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "provider_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "provider_service_key_management"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "provider_service_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "public_peering_bgp_management"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "required_total_cross_connects"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "supported_virtual_circuit_types.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				),
			},
		},
	})
}
