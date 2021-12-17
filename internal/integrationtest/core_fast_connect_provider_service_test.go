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
	fastConnectProviderServiceSingularDataSourceRepresentation = map[string]interface{}{
		"provider_service_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
	}

	fastConnectProviderServiceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	FastConnectProviderServiceResourceConfig = ""
)

// issue-routing-tag: core/default
func TestCoreFastConnectProviderServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreFastConnectProviderServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services"
	singularDatasourceName := "data.oci_core_fast_connect_provider_service.test_fast_connect_provider_service"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_services", "test_fast_connect_provider_services", acctest.Required, acctest.Create, fastConnectProviderServiceDataSourceRepresentation) +
				compartmentIdVariableStr + FastConnectProviderServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_services", "test_fast_connect_provider_services", acctest.Required, acctest.Create, fastConnectProviderServiceDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_service", "test_fast_connect_provider_service", acctest.Required, acctest.Create, fastConnectProviderServiceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FastConnectProviderServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}
