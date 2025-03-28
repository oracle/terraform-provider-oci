// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsLogAnalyticsLogAnalyticsEntitiesSummarySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsEntitiesSummaryResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation)

	LogAnalyticsLogAnalyticsEntitiesSummaryResourceConfig = ""
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsEntitiesSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntitiesSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_log_analytics_log_analytics_entities_summary.test_log_analytics_entities_summary"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsEntitiesSummaryResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entities_summary", "test_log_analytics_entities_summary", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsLogAnalyticsEntitiesSummarySingularDataSourceRepresentation) +
				LogAnalyticsLogAnalyticsEntitiesSummaryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
