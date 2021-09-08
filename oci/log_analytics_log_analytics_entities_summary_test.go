// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	logAnalyticsEntitiesSummarySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"namespace":      Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsEntitiesSummaryResourceDependencies = generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)

	LogAnalyticsEntitiesSummaryResourceConfig = ""
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsEntitiesSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntitiesSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_log_analytics_log_analytics_entities_summary.test_log_analytics_entities_summary"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsEntitiesSummaryResourceDependencies +
				generateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entities_summary", "test_log_analytics_entities_summary", Required, Create, logAnalyticsEntitiesSummarySingularDataSourceRepresentation) +
				LogAnalyticsEntitiesSummaryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "active_entities_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entities_with_has_logs_collected_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entities_with_management_agent_count"),
			),
		},
	})
}
