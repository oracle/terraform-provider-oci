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
	LogAnalyticsLogAnalyticsNamespaceStorageRecallCountSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsNamespaceStorageRecallCountResourceConfig = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceStorageRecallCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceStorageRecallCountResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_log_analytics_namespace_storage_recall_count.test_namespace_storage_recall_count"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_storage_recall_count", "test_namespace_storage_recall_count", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceStorageRecallCountSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceStorageRecallCountResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "recall_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recall_failed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recall_limit"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recall_pending"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recall_succeeded"),
			),
		},
	})
}
