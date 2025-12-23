// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DelegateAccessControlServiceProviderSingularDataSourceRepresentation = map[string]interface{}{
		"service_provider_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_delegate_access_control_service_providers.test_service_providers.service_provider_summary_collection.0.items.0.id}`},
	}

	DelegateAccessControlServiceProviderDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	DelegateAccessControlServiceProviderResourceConfig = ""
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlServiceProviderResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlServiceProviderResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_delegate_access_control_service_providers.test_service_providers"
	singularDatasourceName := "data.oci_delegate_access_control_service_provider.test_service_provider"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_service_providers", "test_service_providers", acctest.Required, acctest.Create, DelegateAccessControlServiceProviderDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlServiceProviderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "service_provider_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_service_providers", "test_service_providers", acctest.Required, acctest.Create, DelegateAccessControlServiceProviderDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_service_provider", "test_service_provider", acctest.Required, acctest.Create, DelegateAccessControlServiceProviderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlServiceProviderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_provider_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_provider_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
