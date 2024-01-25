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
	LogAnalyticsLogAnalyticsNamespaceRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"kind":           acctest.Representation{RepType: acctest.Optional, Create: `INGEST_TIME`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"target_service": acctest.Representation{RepType: acctest.Optional, Create: `MONITORING`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceRulesDataSourceFilterRepresentation},
	}
	LogAnalyticsNamespaceRulesDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule.id}`}},
	}

	LogAnalyticsNamespaceRuleResourceConfig = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule"
	requiredDatasourceName := "data.oci_log_analytics_namespace_rules.test_namespace_rules_required"
	optionalDatasourceName := "data.oci_log_analytics_namespace_rules.test_namespace_rules_optional"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create a rule
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceRuleResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ingest_time_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},
		// verify required input
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceRuleResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_rules", "test_namespace_rules_required", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceRuleDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "rule_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "rule_summary_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "rule_summary_collection.0.items.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "rule_summary_collection.0.items.0.kind", "INGEST_TIME"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "rule_summary_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(requiredDatasourceName, "rule_summary_collection.0.items.0.target_service", "MONITORING"),
			),
		},
		// delete for next test
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceRuleResourceConfig,
		},
		// Recreate the rule
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceRuleResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ingest_time_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},
		// verify optional input
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceRuleResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_rules", "test_namespace_rules_optional", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsNamespaceRuleDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "rule_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "rule_summary_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "rule_summary_collection.0.items.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "rule_summary_collection.0.items.0.kind", "INGEST_TIME"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "rule_summary_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "rule_summary_collection.0.items.0.target_service", "MONITORING"),
			),
		},
		// delete before exit
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceRuleResourceConfig,
		},
	})
}
