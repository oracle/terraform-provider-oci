// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsNamespaceIngestTimeRuleRequiredOnlyResource = LogAnalyticsNamespaceIngestTimeRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation)

	LogAnalyticsNamespaceIngestTimeRuleResourceConfig = LogAnalyticsNamespaceIngestTimeRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Optional, acctest.Update, LogAnalyticsNamespaceIngestTimeRuleRepresentation)

	LogAnalyticsLogAnalyticsNamespaceIngestTimeRuleSingularDataSourceRepresentation = map[string]interface{}{
		"ingest_time_rule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule.id}`},
		"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsLogAnalyticsNamespaceIngestTimeRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"condition_kind": acctest.Representation{RepType: acctest.Optional, Create: `FIELD`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"field_name":     acctest.Representation{RepType: acctest.Optional, Create: `mtag`},
		"field_value":    acctest.Representation{RepType: acctest.Optional, Create: `cveexploitattempt`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`, Update: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceIngestTimeRuleDataSourceFilterRepresentation}}
	LogAnalyticsNamespaceIngestTimeRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule.id}`}},
	}

	LogAnalyticsNamespaceIngestTimeRuleRepresentation = map[string]interface{}{
		"actions":        acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceIngestTimeRuleActionsRepresentation},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"conditions":     acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceIngestTimeRuleConditionsRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	LogAnalyticsNamespaceIngestTimeRuleActionsRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"metric_name":    acctest.Representation{RepType: acctest.Required, Create: `tfmetriccve`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `tfmetricnamespace`, Update: `tfmetricnamespace`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `METRIC_EXTRACTION`},
		"dimensions":     acctest.Representation{RepType: acctest.Optional, Create: []string{`SOURCE_NAME`, `event`}, Update: []string{`SOURCE_NAME`, `event`}},
		"resource_group": acctest.Representation{RepType: acctest.Optional, Create: `critical`, Update: `critical`},
	}
	LogAnalyticsNamespaceIngestTimeRuleConditionsRepresentation = map[string]interface{}{
		"field_name":            acctest.Representation{RepType: acctest.Required, Create: `mtag`, Update: `mtag`},
		"field_operator":        acctest.Representation{RepType: acctest.Required, Create: `EQUAL`},
		"field_value":           acctest.Representation{RepType: acctest.Required, Create: `cveexploitattempt`, Update: `cveexploitattempt`},
		"kind":                  acctest.Representation{RepType: acctest.Required, Create: `FIELD`},
		"additional_conditions": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: LogAnalyticsNamespaceIngestTimeRuleConditionsAdditionalConditionsRepresentation1}, {RepType: acctest.Optional, Group: LogAnalyticsNamespaceIngestTimeRuleConditionsAdditionalConditionsRepresentation2}},
	}
	LogAnalyticsNamespaceIngestTimeRuleConditionsAdditionalConditionsRepresentation1 = map[string]interface{}{
		"condition_field":    acctest.Representation{RepType: acctest.Required, Create: `SOURCE_NAME`, Update: `SOURCE_NAME`},
		"condition_operator": acctest.Representation{RepType: acctest.Required, Create: `EQUAL`, Update: `EQUAL`},
		"condition_value":    acctest.Representation{RepType: acctest.Required, Create: `omc_ociAuditLogSource`, Update: `omc_ociAuditLogSource`},
	}
	LogAnalyticsNamespaceIngestTimeRuleConditionsAdditionalConditionsRepresentation2 = map[string]interface{}{
		"condition_field":    acctest.Representation{RepType: acctest.Required, Create: `mtgttype`, Update: `mtgttype`},
		"condition_operator": acctest.Representation{RepType: acctest.Required, Create: `EQUAL`, Update: `EQUAL`},
		"condition_value":    acctest.Representation{RepType: acctest.Required, Create: `omc_host_linux`, Update: `omc_host_linux`},
	}

	LogAnalyticsNamespaceIngestTimeRuleResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceIngestTimeRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceIngestTimeRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule"
	datasourceName := "data.oci_log_analytics_namespace_ingest_time_rules.test_namespace_ingest_time_rules"
	singularDatasourceName := "data.oci_log_analytics_namespace_ingest_time_rule.test_namespace_ingest_time_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsNamespaceIngestTimeRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Optional, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation), "loganalytics", "namespaceIngestTimeRule", t)

	acctest.ResourceTest(t, testAccCheckLogAnalyticsNamespaceIngestTimeRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceIngestTimeRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "actions.0.metric_name", "tfmetriccve"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.namespace", "tfmetricnamespace"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "METRIC_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_name", "mtag"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_operator", "EQUAL"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_value", "cveexploitattempt"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.kind", "FIELD"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceIngestTimeRuleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceIngestTimeRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Optional, acctest.Create, LogAnalyticsNamespaceIngestTimeRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "actions.0.dimensions.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.metric_name", "tfmetriccve"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.namespace", "tfmetricnamespace"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.resource_group", "critical"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "METRIC_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.additional_conditions.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_name", "mtag"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_operator", "EQUAL"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_value", "cveexploitattempt"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.kind", "FIELD"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LogAnalyticsNamespaceIngestTimeRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LogAnalyticsNamespaceIngestTimeRuleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "actions.0.dimensions.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.metric_name", "tfmetriccve"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.namespace", "tfmetricnamespace"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.resource_group", "critical"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "METRIC_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.additional_conditions.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_name", "mtag"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_operator", "EQUAL"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_value", "cveexploitattempt"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.kind", "FIELD"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceIngestTimeRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Optional, acctest.Update, LogAnalyticsNamespaceIngestTimeRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "actions.0.dimensions.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.metric_name", "tfmetriccve"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.namespace", "tfmetricnamespace"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.resource_group", "critical"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.type", "METRIC_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.additional_conditions.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_name", "mtag"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_operator", "EQUAL"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.field_value", "cveexploitattempt"),
				resource.TestCheckResourceAttr(resourceName, "conditions.0.kind", "FIELD"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rules", "test_namespace_ingest_time_rules", acctest.Optional, acctest.Update, LogAnalyticsLogAnalyticsNamespaceIngestTimeRuleDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceIngestTimeRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Optional, acctest.Update, LogAnalyticsNamespaceIngestTimeRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "condition_kind", "FIELD"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "field_name", "mtag"),
				resource.TestCheckResourceAttr(datasourceName, "field_value", "cveexploitattempt"),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "ingest_time_rule_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ingest_time_rule_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_ingest_time_rule", "test_namespace_ingest_time_rule", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceIngestTimeRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceIngestTimeRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ingest_time_rule_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.dimensions.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.namespace", "tfmetricnamespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.resource_group", "critical"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.type", "METRIC_EXTRACTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.additional_conditions.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.field_name", "mtag"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.field_operator", "EQUAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.field_value", "cveexploitattempt"),
				resource.TestCheckResourceAttr(singularDatasourceName, "conditions.0.kind", "FIELD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + LogAnalyticsNamespaceIngestTimeRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getLogAnalyticsNamespaceIngestTimeRulesEndpointImportId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLogAnalyticsNamespaceIngestTimeRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LogAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_namespace_ingest_time_rule" {
			noResourceFound = false
			request := oci_log_analytics.GetIngestTimeRuleRequest{}

			if value, ok := rs.Primary.Attributes["ingest_time_rule_id"]; ok {
				request.IngestTimeRuleId = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

			response, err := client.GetIngestTimeRule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_log_analytics.ConfigLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("LogAnalyticsNamespaceIngestTimeRule") {
		resource.AddTestSweepers("LogAnalyticsNamespaceIngestTimeRule", &resource.Sweeper{
			Name:         "LogAnalyticsNamespaceIngestTimeRule",
			Dependencies: acctest.DependencyGraph["namespaceIngestTimeRule"],
			F:            sweepLogAnalyticsNamespaceIngestTimeRuleResource,
		})
	}
}

func sweepLogAnalyticsNamespaceIngestTimeRuleResource(compartment string) error {
	logAnalyticsClient := acctest.GetTestClients(&schema.ResourceData{}).LogAnalyticsClient()
	namespaceIngestTimeRuleIds, err := getLogAnalyticsNamespaceIngestTimeRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, namespaceIngestTimeRuleId := range namespaceIngestTimeRuleIds {
		if ok := acctest.SweeperDefaultResourceId[namespaceIngestTimeRuleId]; !ok {
			deleteIngestTimeRuleRequest := oci_log_analytics.DeleteIngestTimeRuleRequest{}

			deleteIngestTimeRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")
			_, error := logAnalyticsClient.DeleteIngestTimeRule(context.Background(), deleteIngestTimeRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting NamespaceIngestTimeRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", namespaceIngestTimeRuleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &namespaceIngestTimeRuleId, LogAnalyticsNamespaceIngestTimeRuleSweepWaitCondition, time.Duration(3*time.Minute),
				LogAnalyticsNamespaceIngestTimeRuleSweepResponseFetchOperation, "log_analytics", true)
		}
	}
	return nil
}

func getLogAnalyticsNamespaceIngestTimeRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NamespaceIngestTimeRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	logAnalyticsClient := acctest.GetTestClients(&schema.ResourceData{}).LogAnalyticsClient()

	listIngestTimeRulesRequest := oci_log_analytics.ListIngestTimeRulesRequest{}
	listIngestTimeRulesRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for NamespaceIngestTimeRule resource requests \n")
	}
	for _, namespace := range namespaces {
		listIngestTimeRulesRequest.NamespaceName = &namespace

		listIngestTimeRulesRequest.LifecycleState = oci_log_analytics.ListIngestTimeRulesLifecycleStateActive
		listIngestTimeRulesResponse, err := logAnalyticsClient.ListIngestTimeRules(context.Background(), listIngestTimeRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NamespaceIngestTimeRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, namespaceIngestTimeRule := range listIngestTimeRulesResponse.Items {
			id := *namespaceIngestTimeRule.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NamespaceIngestTimeRuleId", id)
		}

	}
	return resourceIds, nil
}

func LogAnalyticsNamespaceIngestTimeRuleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if namespaceIngestTimeRuleResponse, ok := response.Response.(oci_log_analytics.GetIngestTimeRuleResponse); ok {
		return namespaceIngestTimeRuleResponse.LifecycleState != oci_log_analytics.ConfigLifecycleStateDeleted
	}
	return false
}

func LogAnalyticsNamespaceIngestTimeRuleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LogAnalyticsClient().GetIngestTimeRule(context.Background(), oci_log_analytics.GetIngestTimeRuleRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func getLogAnalyticsNamespaceIngestTimeRulesEndpointImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("namespaces/" + rs.Primary.Attributes["namespace"] + "/ingestTimeRules/" + rs.Primary.Attributes["id"]), nil
	}
}
