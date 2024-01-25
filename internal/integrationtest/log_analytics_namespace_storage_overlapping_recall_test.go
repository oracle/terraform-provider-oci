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
	LogAnalyticsLogAnalyticsNamespaceStorageOverlappingRecallDataSourceRepresentation = map[string]interface{}{
		"namespace":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"time_data_ended":   acctest.Representation{RepType: acctest.Optional, Create: `2020-06-25T00:00:00.000Z`},
		"time_data_started": acctest.Representation{RepType: acctest.Optional, Create: `2020-06-05T00:00:00.000Z`},
	}

	LogAnalyticsNamespaceStorageOverlappingRecallResourceConfig = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceStorageOverlappingRecallResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceStorageOverlappingRecallResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	requiredDatasourceName := "data.oci_log_analytics_namespace_storage_overlapping_recalls.test_namespace_storage_overlapping_recalls_required"
	optionalDatasourceName := "data.oci_log_analytics_namespace_storage_overlapping_recalls.test_namespace_storage_overlapping_recalls_optional"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource with required parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_storage_overlapping_recalls", "test_namespace_storage_overlapping_recalls_required", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceStorageOverlappingRecallDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceStorageOverlappingRecallResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "overlapping_recall_collection.0.items.0.status"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "overlapping_recall_collection.0.items.0.time_data_ended"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "overlapping_recall_collection.0.items.0.time_data_started"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "overlapping_recall_collection.0.items.0.time_started"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "overlapping_recall_collection.0.items.0.collection_id"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "overlapping_recall_collection.0.items.0.recall_id"),
			),
		},
		// verify datasource with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_storage_overlapping_recalls", "test_namespace_storage_overlapping_recalls_optional", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsNamespaceStorageOverlappingRecallDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceStorageOverlappingRecallResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "overlapping_recall_collection.0.items.0.status"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "overlapping_recall_collection.0.items.0.time_data_ended"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "overlapping_recall_collection.0.items.0.time_data_started"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "overlapping_recall_collection.0.items.0.time_started"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "overlapping_recall_collection.0.items.0.collection_id"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "overlapping_recall_collection.0.items.0.recall_id"),
			),
		},
	})
}
