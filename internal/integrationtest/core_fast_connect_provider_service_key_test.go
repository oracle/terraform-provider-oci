// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCoreFastConnectProviderServiceKeySingularDataSourceRepresentation = map[string]interface{}{
		"provider_service_id":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_fast_connect_provider_services.test_fast_connect_provider_services.fast_connect_provider_services.0.id}`},
		"provider_service_key_name": acctest.Representation{RepType: acctest.Required, Create: `d8f7a443-28c2-4dcf-996c-286351908c58`},
	}

	CoreFastConnectProviderServiceKeyResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_services", "test_fast_connect_provider_services", acctest.Required, acctest.Create, CoreCoreFastConnectProviderServiceDataSourceRepresentation)
)

// issue-routing-tag: core/default
func TestCoreFastConnectProviderServiceKeyResource_basic(t *testing.T) {
	if httpreplay.ShouldRetryImmediately() {
		t.Skip("FastConnect failed in dxterraformtest tenancy IAD region with a known issue: TER-1232")
	}

	httpreplay.SetScenario("TestCoreFastConnectProviderServiceKeyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_fast_connect_provider_service_key.test_fast_connect_provider_service_key"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_fast_connect_provider_service_key", "test_fast_connect_provider_service_key", acctest.Required, acctest.Create, CoreCoreFastConnectProviderServiceKeySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreFastConnectProviderServiceKeyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "provider_service_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provider_service_key_name", "d8f7a443-28c2-4dcf-996c-286351908c58"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "bandwidth_shape_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peering_location"),
			),
		},
	})
}
