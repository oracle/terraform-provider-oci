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
	// Prakash - we need one ID defined here or send using TF_VAR, since we do NOT have a POST (create) for this

	DelegateAccessControlServiceProviderSingularDataSourceRepresentation = map[string]interface{}{
		// Prakash - use the service provider ocid defined above, that already exists
		"service_provider_id": acctest.Representation{RepType: acctest.Required, Create: `${var.svcProviderId}`},
	}

	DelegateAccessControlServiceProviderDataSourceRepresentation = map[string]interface{}{
		// Prakash - the service provider can only exist in the root compartment, so use tenancy ocid
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"service_provider_type":   acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_PROVIDED`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"supported_resource_type": acctest.Representation{RepType: acctest.Optional, Create: `VMCLUSTER`},
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
				// Prakash, this is the output of the list, it has an items [], but none of the individual attributes
				/*resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "service_provider_type", "ORACLE_PROVIDED"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "supported_resource_type", "VMCLUSTER"),*/

				resource.TestCheckResourceAttrSet(datasourceName, "service_provider_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_service_provider", "test_service_provider", acctest.Required, acctest.Create, DelegateAccessControlServiceProviderSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlServiceProviderResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_provider_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_provider_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_types.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "supported_resource_types.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
