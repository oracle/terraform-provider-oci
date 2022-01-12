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
	"github.com/oracle/oci-go-sdk/v55/common"
	oci_events "github.com/oracle/oci-go-sdk/v55/events"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RuleRequiredOnlyResource = RuleResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", Required, Create, ruleRepresentation)

	RuleResourceConfig = RuleResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Update, ruleRepresentation)

	ruleSingularDataSourceRepresentation = map[string]interface{}{
		"rule_id": Representation{RepType: Required, Create: `${oci_events_rule.test_rule.id}`},
	}

	ruleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `This rule sends a notification upon completion of DbaaS backup`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `INACTIVE`},
		"filter":         RepresentationGroup{Required, ruleDataSourceFilterRepresentation}}
	ruleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_events_rule.test_rule.id}`}},
	}

	ruleRepresentation = map[string]interface{}{
		"actions":        RepresentationGroup{Required, ruleActionsRepresentation},
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"condition":      Representation{RepType: Required, Create: `{\"eventType\":\"com.oraclecloud.databaseservice.autonomous.database.backup.end\"}`, Update: `{}`},
		"display_name":   Representation{RepType: Required, Create: `This rule sends a notification upon completion of DbaaS backup`, Update: `displayName2`},
		"is_enabled":     Representation{RepType: Required, Create: `true`, Update: `false`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	ruleActionsRepresentation = map[string]interface{}{
		"actions": []RepresentationGroup{{Optional, ruleActionsOSSActionsRepresentation}},
	}
	ruleActionsUpdateRepresentation = map[string]interface{}{
		"actions": []RepresentationGroup{{Optional, ruleActionsOSSActionsRepresentation}, {Optional, ruleActionsONSActionsRepresentation}, {Optional, ruleActionsFAASActionsRepresentation}},
	}
	ruleActionsONSActionsRepresentation = map[string]interface{}{
		"action_type": Representation{RepType: Required, Create: `ONS`, Update: `ONS`},
		"is_enabled":  Representation{RepType: Required, Create: `false`, Update: `true`},
		"description": Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"topic_id":    Representation{RepType: Optional, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}
	ruleActionsOSSActionsRepresentation = map[string]interface{}{
		"action_type": Representation{RepType: Required, Create: `OSS`, Update: `ONS`},
		"is_enabled":  Representation{RepType: Required, Create: `false`, Update: `true`},
		"description": Representation{RepType: Optional, Create: `description`, Update: `rule type updated`},
		"stream_id":   Representation{RepType: Optional, Create: `${oci_streaming_stream.test_stream.id}`, Update: ``},
		"topic_id":    Representation{RepType: Optional, Create: ``, Update: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}
	ruleActionsFAASActionsRepresentation = map[string]interface{}{
		"action_type": Representation{RepType: Required, Create: `FAAS`, Update: `FAAS`},
		"is_enabled":  Representation{RepType: Required, Create: `false`, Update: `true`},
		"description": Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"function_id": Representation{RepType: Optional, Create: `${oci_functions_function.test_function.id}`},
	}

	RuleResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
		GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation) +
		GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamRepresentation)
)

// issue-routing-tag: events/default
func TestEventsRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEventsRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	image := getEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	resourceName := "oci_events_rule.test_rule"
	datasourceName := "data.oci_events_rules.test_rules"
	singularDatasourceName := "data.oci_events_rule.test_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+RuleResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Create, ruleRepresentation), "events", "rule", t)

	ResourceTest(t, testAccCheckEventsRuleDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + RuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Create, ruleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.actions.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
					"action_type": "OSS",
					"description": "description",
					"is_enabled":  "false",
				},
					[]string{
						"id",
						"state",
						"stream_id",
					}),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "condition", "{\"eventType\":\"com.oraclecloud.databaseservice.autonomous.database.backup.end\"}"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "This rule sends a notification upon completion of DbaaS backup"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + compartmentIdUVariableStr + RuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Create,
					RepresentationCopyWithNewProperties(ruleRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.actions.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
					"action_type": "OSS",
					"description": "description",
					"is_enabled":  "false",
				},
					[]string{
						"id",
						"state",
						"stream_id",
					}),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "condition", "{\"eventType\":\"com.oraclecloud.databaseservice.autonomous.database.backup.end\"}"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "This rule sends a notification upon completion of DbaaS backup"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + RuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Update,
					GetUpdatedRepresentationCopy("actions", RepresentationGroup{Optional, ruleActionsUpdateRepresentation}, ruleRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.actions.#", "3"),
				CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
					"action_type": "ONS",
					"description": "rule type updated",
					"is_enabled":  "true",
				},
					[]string{
						"id",
						"state",
						"topic_id",
					}),
				CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
					"action_type": "ONS",
					"description": "description2",
					"is_enabled":  "true",
				},
					[]string{
						"id",
						"state",
						"topic_id",
					}),
				CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
					"action_type": "FAAS",
					"description": "description2",
					"is_enabled":  "true",
				},
					[]string{
						"function_id",
						"id",
						"state",
					}),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "condition", "{}"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_events_rules", "test_rules", Optional, Update, ruleDataSourceRepresentation) +
				compartmentIdVariableStr + imageVariableStr + RuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Update, ruleRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.condition", "{}"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "rules.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "rules.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "rules.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "rules.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_events_rule", "test_rule", Required, Create, ruleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + imageVariableStr + RuleResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "rule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition", "{}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + RuleResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckEventsRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).eventsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_events_rule" {
			noResourceFound = false
			request := oci_events.GetRuleRequest{}

			tmp := rs.Primary.ID
			request.RuleId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "events")

			response, err := client.GetRule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_events.RuleLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("EventsRule") {
		resource.AddTestSweepers("EventsRule", &resource.Sweeper{
			Name:         "EventsRule",
			Dependencies: DependencyGraph["rule"],
			F:            sweepEventsRuleResource,
		})
	}
}

func sweepEventsRuleResource(compartment string) error {
	eventsClient := GetTestClients(&schema.ResourceData{}).eventsClient()
	ruleIds, err := getRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, ruleId := range ruleIds {
		if ok := SweeperDefaultResourceId[ruleId]; !ok {
			deleteRuleRequest := oci_events.DeleteRuleRequest{}

			deleteRuleRequest.RuleId = &ruleId

			deleteRuleRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "events")
			_, error := eventsClient.DeleteRule(context.Background(), deleteRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting Rule %s %s, It is possible that the resource is already deleted. Please verify manually \n", ruleId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &ruleId, ruleSweepWaitCondition, time.Duration(3*time.Minute),
				ruleSweepResponseFetchOperation, "events", true)
		}
	}
	return nil
}

func getRuleIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "RuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	eventsClient := GetTestClients(&schema.ResourceData{}).eventsClient()

	listRulesRequest := oci_events.ListRulesRequest{}
	listRulesRequest.CompartmentId = &compartmentId
	listRulesRequest.LifecycleState = oci_events.RuleLifecycleStateActive
	listRulesResponse, err := eventsClient.ListRules(context.Background(), listRulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Rule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, rule := range listRulesResponse.Items {
		id := *rule.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "RuleId", id)
	}
	return resourceIds, nil
}

func ruleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ruleResponse, ok := response.Response.(oci_events.GetRuleResponse); ok {
		return ruleResponse.LifecycleState != oci_events.RuleLifecycleStateDeleted
	}
	return false
}

func ruleSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.eventsClient().GetRule(context.Background(), oci_events.GetRuleRequest{
		RuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
