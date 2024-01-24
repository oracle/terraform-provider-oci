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
	LogAnalyticsNamespaceIngestTimeRulesManagementRepresentation = map[string]interface{}{
		"ingest_time_rule_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule.id}`},
		"namespace":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"enable_ingest_time_rule": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	NamespaceIngestTimeRulesManagementResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceIngestTimeRulesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceIngestTimeRulesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_namespace_ingest_time_rules_management.test_namespace_ingest_time_rules_management"
	parentResourceName := "oci_log_analytics_namespace_ingest_time_rules_management.test_namespace_ingest_time_rules_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NamespaceIngestTimeRulesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation)+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rules_management", "test_namespace_ingest_time_rules_management", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRulesManagementRepresentation), "loganalytics", "namespaceIngestTimeRulesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + NamespaceIngestTimeRulesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rules_management", "test_namespace_ingest_time_rules_management", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRulesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ingest_time_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + NamespaceIngestTimeRulesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rules_management", "test_namespace_ingest_time_rules_management", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRulesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_ingest_time_rule", "true"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NamespaceIngestTimeRulesManagementResourceDependencies,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + NamespaceIngestTimeRulesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rules_management", "test_namespace_ingest_time_rules_management", acctest.Optional, acctest.Create, LogAnalyticsNamespaceIngestTimeRulesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ingest_time_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + NamespaceIngestTimeRulesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rules_management", "test_namespace_ingest_time_rules_management", acctest.Optional, acctest.Update, LogAnalyticsNamespaceIngestTimeRulesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ingest_time_rule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + NamespaceIngestTimeRulesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rules_management", "test_namespace_ingest_time_rules_management", acctest.Optional, acctest.Update, LogAnalyticsNamespaceIngestTimeRulesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_ingest_time_rule", "false"),
			),
		},
		// delete before exit
		{
			Config: config + compartmentIdVariableStr + NamespaceIngestTimeRulesManagementResourceDependencies,
		},
	})
}
