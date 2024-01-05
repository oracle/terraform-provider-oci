// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsLogAnalyticsLogSetsCountSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsLogSetsCountResourceConfig = ""
	LogAnalyticsLogSetsCountDependencies   = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogSetsCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogSetsCountResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	singularDatasourceName := "data.oci_log_analytics_log_sets_count.test_log_sets_count"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_sets_count", "test_log_sets_count", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsLogSetsCountSingularDataSourceRepresentation) +
				LogAnalyticsLogSetsCountDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_sets_count"),
			),
		},
	})
}
