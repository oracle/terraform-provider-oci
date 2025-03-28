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
	LogAnalyticsLogAnalyticsLogAnalyticsLogGroupsSummarySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsLogGroupsSummaryResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation)

	LogAnalyticsLogAnalyticsLogGroupsSummaryResourceConfig = ""
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsLogGroupsSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsLogGroupsSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_log_analytics_log_analytics_log_groups_summary.test_log_analytics_log_groups_summary"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				compartmentIdVariableStr +
				LogAnalyticsLogGroupsSummaryResourceDependencies +
				LogAnalyticsLogAnalyticsLogGroupsSummaryResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_log_groups_summary", "test_log_analytics_log_groups_summary", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsLogAnalyticsLogGroupsSummarySingularDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_group_count"),
			),
		},
	})
}
