// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	logSetsCountSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogSetsCountResourceConfig           = ""
	LogAnalyticsLogSetsCountDependencies = GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogSetsCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogSetsCountResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	singularDatasourceName := "data.oci_log_analytics_log_sets_count.test_log_sets_count"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_sets_count", "test_log_sets_count", Required, Create, logSetsCountSingularDataSourceRepresentation) +
				LogAnalyticsLogSetsCountDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_sets_count"),
			),
		},
	})
}
