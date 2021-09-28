// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v48/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v48/loganalytics"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsObjectCollectionRuleRequiredOnlyResource = LogAnalyticsObjectCollectionRuleResourceDependencies +
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Required, Create, logAnalyticsObjectCollectionRuleRepresentation)

	LogAnalyticsObjectCollectionRuleResourceConfig = LogAnalyticsObjectCollectionRuleResourceDependencies +
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Optional, Update, logAnalyticsObjectCollectionRuleRepresentation)

	logAnalyticsObjectCollectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"log_analytics_object_collection_rule_id": Representation{repType: Required, create: `${oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule.id}`},
		"namespace": Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	logAnalyticsObjectCollectionRuleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"namespace":      Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"name":           Representation{repType: Optional, create: `test_terraform_rule`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, logAnalyticsObjectCollectionRuleDataSourceFilterRepresentation}}
	logAnalyticsObjectCollectionRuleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule.id}`}},
	}

	logAnalyticsObjectCollectionRulePropertyOverridesRepresentation = map[string]interface{}{
		"match_type":     Representation{repType: Optional, create: `contains`, update: `contains`},
		"match_value":    Representation{repType: Optional, create: `db`, update: `db`},
		"property_name":  Representation{repType: Optional, create: `charEncoding`, update: `charEncoding`},
		"property_value": Representation{repType: Optional, create: `utf-8`, update: `utf-16`},
	}

	logAnalyticsObjectCollectionRuleRepresentation = map[string]interface{}{
		"compartment_id":  Representation{repType: Required, create: `${var.compartment_id}`},
		"log_group_id":    Representation{repType: Required, create: `${oci_log_analytics_log_analytics_log_group.test_log_analytics_log_group.id}`},
		"log_source_name": Representation{repType: Required, create: `LinuxSyslogSource`, update: `LinuxSyslogSource`},
		"name":            Representation{repType: Required, create: `test_terraform_rule`},
		"namespace":       Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"os_bucket_name":  Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"os_namespace":    Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"char_encoding":   Representation{repType: Optional, create: `utf-8`, update: `utf-16`},
		"collection_type": Representation{repType: Optional, create: `HISTORIC`},
		"defined_tags":    Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":     Representation{repType: Optional, create: `test terraform rule create`, update: `test terraform rule update`},
		"entity_id":       Representation{repType: Optional, create: `${oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id}`},
		"freeform_tags":   Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"overrides":       RepresentationGroup{Optional, logAnalyticsObjectCollectionRulePropertyOverridesRepresentation},
		"poll_since":      Representation{repType: Optional, create: `2020-04-01T00:00:00.000Z`},
		"poll_till":       Representation{repType: Optional, create: `2021-04-01T00:00:00.000Z`},
	}

	LogAnalyticsObjectCollectionRuleResourceDependencies = DefinedTagsDependencies +
		generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_log_group", "test_log_analytics_log_group", Required, Create, logAnalyticsLogGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity", "test_log_analytics_entity", Optional, Create, logAnalyticsEntityRepresentation)
)

func TestLogAnalyticsLogAnalyticsObjectCollectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsObjectCollectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	managementAgentId := getEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule"
	datasourceName := "data.oci_log_analytics_log_analytics_object_collection_rules.test_log_analytics_object_collection_rules"
	singularDatasourceName := "data.oci_log_analytics_log_analytics_object_collection_rule.test_log_analytics_object_collection_rule"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+managementAgentIdVariableStr+LogAnalyticsObjectCollectionRuleResourceDependencies+
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Optional, Create, logAnalyticsObjectCollectionRuleRepresentation), "loganalytics", "logAnalyticsObjectCollectionRule", t)

	ResourceTest(t, testAccCheckLogAnalyticsLogAnalyticsObjectCollectionRuleDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceDependencies +
				generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Required, Create, logAnalyticsObjectCollectionRuleRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "os_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "os_namespace"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceDependencies +
				generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Optional, Create, logAnalyticsObjectCollectionRuleRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "char_encoding", "utf-8"),
				resource.TestCheckResourceAttr(resourceName, "collection_type", "HISTORIC"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "test terraform rule create"),
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "os_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "os_namespace"),
				resource.TestCheckResourceAttr(resourceName, "overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_name", "charEncoding"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_value", "utf-8"),
				resource.TestCheckResourceAttr(resourceName, "poll_since", "2020-04-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(resourceName, "poll_till", "2021-04-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceDependencies +
				generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Optional, Create,
					representationCopyWithNewProperties(logAnalyticsObjectCollectionRuleRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "char_encoding", "utf-8"),
				resource.TestCheckResourceAttr(resourceName, "collection_type", "HISTORIC"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "test terraform rule create"),
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "os_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "os_namespace"),
				resource.TestCheckResourceAttr(resourceName, "overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_name", "charEncoding"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_value", "utf-8"),
				resource.TestCheckResourceAttr(resourceName, "poll_since", "2020-04-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(resourceName, "poll_till", "2021-04-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						Logf("TestLogAnalyticsLogAnalyticsObjectCollectionRuleResource:: resource Ids not matching \n%s\n%s", fmt.Sprintf(resId), fmt.Sprintf(resId2))
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceDependencies +
				generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Optional, Update, logAnalyticsObjectCollectionRuleRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(resourceName, "char_encoding", "utf-16"),
				resource.TestCheckResourceAttr(resourceName, "collection_type", "HISTORIC"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "test terraform rule update"),
				resource.TestCheckResourceAttrSet(resourceName, "entity_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(resourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "os_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "os_namespace"),
				resource.TestCheckResourceAttr(resourceName, "overrides.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_name", "charEncoding"),
				resource.TestCheckResourceAttr(resourceName, "overrides.0.property_value", "utf-16"),
				resource.TestCheckResourceAttr(resourceName, "poll_since", "2020-04-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(resourceName, "poll_till", "2021-04-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						Logf("TestLogAnalyticsLogAnalyticsObjectCollectionRuleResource:: resource Ids not matching \n%s\n%s", fmt.Sprintf(resId), fmt.Sprintf(resId2))
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rules", "test_log_analytics_object_collection_rules", Optional, Update, logAnalyticsObjectCollectionRuleDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceDependencies +
				generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Optional, Update, logAnalyticsObjectCollectionRuleRepresentation),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "log_analytics_object_collection_rule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "log_analytics_object_collection_rule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_object_collection_rule", "test_log_analytics_object_collection_rule", Required, Create, logAnalyticsObjectCollectionRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_analytics_object_collection_rule_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttr(singularDatasourceName, "char_encoding", "utf-16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "collection_type", "HISTORIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "test terraform rule update"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_source_name", "LinuxSyslogSource"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "test_terraform_rule"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "overrides.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "overrides.0.property_name", "charEncoding"),
				resource.TestCheckResourceAttr(singularDatasourceName, "overrides.0.property_value", "utf-16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "poll_since", "2020-04-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "poll_till", "2021-04-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceConfig,
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsObjectCollectionRuleResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getLogAnalyticsObjectCollectionRulesEndpointImportId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLogAnalyticsLogAnalyticsObjectCollectionRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).logAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_log_analytics_object_collection_rule" {
			noResourceFound = false
			request := oci_log_analytics.GetLogAnalyticsObjectCollectionRuleRequest{}

			tmp := rs.Primary.ID
			request.LogAnalyticsObjectCollectionRuleId = &tmp

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "log_analytics")

			response, err := client.GetLogAnalyticsObjectCollectionRule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_log_analytics.ObjectCollectionRuleLifecycleStatesDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("LogAnalyticsLogAnalyticsObjectCollectionRule") {
		resource.AddTestSweepers("LogAnalyticsLogAnalyticsObjectCollectionRule", &resource.Sweeper{
			Name:         "LogAnalyticsLogAnalyticsObjectCollectionRule",
			Dependencies: DependencyGraph["logAnalyticsObjectCollectionRule"],
			F:            sweepLogAnalyticsLogAnalyticsObjectCollectionRuleResource,
		})
	}
}

func sweepLogAnalyticsLogAnalyticsObjectCollectionRuleResource(compartment string) error {
	logAnalyticsClient := GetTestClients(&schema.ResourceData{}).logAnalyticsClient()
	logAnalyticsObjectCollectionRuleIds, err := getLogAnalyticsObjectCollectionRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, logAnalyticsObjectCollectionRuleId := range logAnalyticsObjectCollectionRuleIds {
		if ok := SweeperDefaultResourceId[logAnalyticsObjectCollectionRuleId]; !ok {
			deleteLogAnalyticsObjectCollectionRuleRequest := oci_log_analytics.DeleteLogAnalyticsObjectCollectionRuleRequest{}

			deleteLogAnalyticsObjectCollectionRuleRequest.LogAnalyticsObjectCollectionRuleId = &logAnalyticsObjectCollectionRuleId

			deleteLogAnalyticsObjectCollectionRuleRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "log_analytics")
			_, error := logAnalyticsClient.DeleteLogAnalyticsObjectCollectionRule(context.Background(), deleteLogAnalyticsObjectCollectionRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting LogAnalyticsObjectCollectionRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", logAnalyticsObjectCollectionRuleId, error)
				continue
			}
			waitTillCondition(testAccProvider, &logAnalyticsObjectCollectionRuleId, logAnalyticsObjectCollectionRuleSweepWaitCondition, time.Duration(3*time.Minute),
				logAnalyticsObjectCollectionRuleSweepResponseFetchOperation, "log_analytics", true)
		}
	}
	return nil
}

func getLogAnalyticsObjectCollectionRuleIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "LogAnalyticsObjectCollectionRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	logAnalyticsClient := GetTestClients(&schema.ResourceData{}).logAnalyticsClient()

	listLogAnalyticsObjectCollectionRulesRequest := oci_log_analytics.ListLogAnalyticsObjectCollectionRulesRequest{}
	listLogAnalyticsObjectCollectionRulesRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for LogAnalyticsObjectCollectionRule resource requests \n")
	}
	for _, namespace := range namespaces {
		listLogAnalyticsObjectCollectionRulesRequest.NamespaceName = &namespace

		listLogAnalyticsObjectCollectionRulesRequest.LifecycleState = oci_log_analytics.ListLogAnalyticsObjectCollectionRulesLifecycleStateActive
		listLogAnalyticsObjectCollectionRulesResponse, err := logAnalyticsClient.ListLogAnalyticsObjectCollectionRules(context.Background(), listLogAnalyticsObjectCollectionRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting LogAnalyticsObjectCollectionRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, logAnalyticsObjectCollectionRule := range listLogAnalyticsObjectCollectionRulesResponse.Items {
			id := *logAnalyticsObjectCollectionRule.Id
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "LogAnalyticsObjectCollectionRuleId", id)
		}

	}
	return resourceIds, nil
}

func logAnalyticsObjectCollectionRuleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if logAnalyticsObjectCollectionRuleResponse, ok := response.Response.(oci_log_analytics.GetLogAnalyticsObjectCollectionRuleResponse); ok {
		return logAnalyticsObjectCollectionRuleResponse.LifecycleState != oci_log_analytics.ObjectCollectionRuleLifecycleStatesDeleted
	}
	return false
}

func logAnalyticsObjectCollectionRuleSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.logAnalyticsClient().GetLogAnalyticsObjectCollectionRule(context.Background(), oci_log_analytics.GetLogAnalyticsObjectCollectionRuleRequest{
		LogAnalyticsObjectCollectionRuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getLogAnalyticsObjectCollectionRulesEndpointImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("namespaces/" + rs.Primary.Attributes["namespace"] + "/logAnalyticsObjectCollectionRules/" + rs.Primary.Attributes["id"]), nil
	}
}
