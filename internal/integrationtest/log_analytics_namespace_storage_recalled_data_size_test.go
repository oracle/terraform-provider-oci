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
	LogAnalyticsLogAnalyticsNamespaceStorageRecalledDataSizeSingularDataSourceRepresentation = map[string]interface{}{
		"namespace":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"time_data_ended":   acctest.Representation{RepType: acctest.Optional, Create: `2023-01-31T23:59:59.000Z`},
		"time_data_started": acctest.Representation{RepType: acctest.Optional, Create: `2023-01-01T00:00:01.000Z`},
	}

	LogAnalyticsNamespaceStorageRecalledDataSizeResourceConfig = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceStorageRecalledDataSizeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceStorageRecalledDataSizeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	requiredDatasourceName := "data.oci_log_analytics_namespace_storage_recalled_data_size.test_namespace_storage_recalled_data_size_required"
	optionalDatasourceName := "data.oci_log_analytics_namespace_storage_recalled_data_size.test_namespace_storage_recalled_data_size_optional"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource with required parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_storage_recalled_data_size", "test_namespace_storage_recalled_data_size_required", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceStorageRecalledDataSizeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceStorageRecalledDataSizeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "not_recalled_data_in_bytes"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "recalled_data_in_bytes"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "time_data_ended"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "time_data_started"),
			),
		},
		// verify datasource with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_storage_recalled_data_size", "test_namespace_storage_recalled_data_size_optional", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsNamespaceStorageRecalledDataSizeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceStorageRecalledDataSizeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "not_recalled_data_in_bytes"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "recalled_data_in_bytes"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "time_data_ended"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "time_data_started"),
			),
		},
	})
}
