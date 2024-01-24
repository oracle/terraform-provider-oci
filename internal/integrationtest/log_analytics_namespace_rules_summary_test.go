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
	LogAnalyticsLogAnalyticsNamespaceRulesSummarySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsNamespaceRulesSummaryResourceConfig = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceRulesSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceRulesSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule"
	singularDatasourceName := "data.oci_log_analytics_namespace_rules_summary.test_namespace_rules_summary"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create a rule
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceRulesSummaryResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ingest_time_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_rules_summary", "test_namespace_rules_summary", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceRulesSummarySingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceRulesSummaryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "ingest_time_rules_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "saved_search_rules_count", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "total_count", "1"),
			),
		},
		// delete before exit
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceRulesSummaryResourceConfig,
		},
	})
}
