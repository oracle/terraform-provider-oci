// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_events "github.com/oracle/oci-go-sdk/v65/events"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	EventsRuleRequiredOnlyResource = EventsRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Required, acctest.Create, ruleRepresentation)

	EventsRuleResourceConfig = EventsRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Optional, acctest.Update, ruleRepresentation)

	EventsruleSingularDataSourceRepresentation = map[string]interface{}{
		"rule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_events_rule.test_rule.id}`},
	}

	EventsruleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `This rule sends a notification upon completion of DbaaS backup`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ruleDataSourceFilterRepresentation}}
	ruleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_events_rule.test_rule.id}`}},
	}

	ruleRepresentation = map[string]interface{}{
		"actions":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ruleActionsRepresentation},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"condition":      acctest.Representation{RepType: acctest.Required, Create: `{\"eventType\":\"com.oraclecloud.databaseservice.autonomous.database.backup.end\"}`, Update: `{}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `This rule sends a notification upon completion of DbaaS backup`, Update: `displayName2`},
		"is_enabled":     acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	ruleActionsRepresentation = map[string]interface{}{
		"actions": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: ruleActionsOSSActionsRepresentation}},
	}
	ruleActionsUpdateRepresentation = map[string]interface{}{
		"actions": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: ruleActionsOSSActionsRepresentation}, {RepType: acctest.Optional, Group: ruleActionsONSActionsRepresentation}, {RepType: acctest.Optional, Group: ruleActionsFAASActionsRepresentation}},
	}
	ruleActionsONSActionsRepresentation = map[string]interface{}{
		"action_type": acctest.Representation{RepType: acctest.Required, Create: `ONS`, Update: `ONS`},
		"is_enabled":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"topic_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}
	ruleActionsOSSActionsRepresentation = map[string]interface{}{
		"action_type": acctest.Representation{RepType: acctest.Required, Create: `OSS`, Update: `ONS`},
		"is_enabled":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `rule type updated`},
		"stream_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_streaming_stream.test_stream.id}`, Update: ``},
		"topic_id":    acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}
	ruleActionsFAASActionsRepresentation = map[string]interface{}{
		"action_type": acctest.Representation{RepType: acctest.Required, Create: `FAAS`, Update: `FAAS`},
		"is_enabled":  acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"function_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_functions_function.test_function.id}`},
	}

	EventsRuleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation)
)

// issue-routing-tag: events/default
func TestEventsRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEventsRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	resourceName := "oci_events_rule.test_rule"
	datasourceName := "data.oci_events_rules.test_rules"
	singularDatasourceName := "data.oci_events_rule.test_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+EventsRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Optional, acctest.Create, ruleRepresentation), "events", "rule", t)

	acctest.ResourceTest(t, testAccCheckEventsRuleDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + imageVariableStr + EventsRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Optional, acctest.Create, ruleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.actions.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
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
			Config: config + compartmentIdVariableStr + imageVariableStr + compartmentIdUVariableStr + EventsRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ruleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.actions.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
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
			Config: config + compartmentIdVariableStr + imageVariableStr + EventsRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("actions", acctest.RepresentationGroup{RepType: acctest.Optional, Group: ruleActionsUpdateRepresentation}, ruleRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "actions.0.actions.#", "3"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
					"action_type": "ONS",
					"description": "rule type updated",
					"is_enabled":  "true",
				},
					[]string{
						"id",
						"state",
						"topic_id",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
					"action_type": "ONS",
					"description": "description2",
					"is_enabled":  "true",
				},
					[]string{
						"id",
						"state",
						"topic_id",
					}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "actions.0.actions", map[string]string{
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_events_rules", "test_rules", acctest.Optional, acctest.Update, EventsruleDataSourceRepresentation) +
				compartmentIdVariableStr + imageVariableStr + EventsRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Optional, acctest.Update, ruleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_events_rule", "test_rule", acctest.Required, acctest.Create, EventsruleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + imageVariableStr + EventsRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		// verify resource import
		{
			Config:                  config + EventsRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckEventsRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EventsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_events_rule" {
			noResourceFound = false
			request := oci_events.GetRuleRequest{}

			tmp := rs.Primary.ID
			request.RuleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "events")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("EventsRule") {
		resource.AddTestSweepers("EventsRule", &resource.Sweeper{
			Name:         "EventsRule",
			Dependencies: acctest.DependencyGraph["rule"],
			F:            sweepEventsRuleResource,
		})
	}
}

func sweepEventsRuleResource(compartment string) error {
	eventsClient := acctest.GetTestClients(&schema.ResourceData{}).EventsClient()
	ruleIds, err := getEventsRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, ruleId := range ruleIds {
		if ok := acctest.SweeperDefaultResourceId[ruleId]; !ok {
			deleteRuleRequest := oci_events.DeleteRuleRequest{}

			deleteRuleRequest.RuleId = &ruleId

			deleteRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "events")
			_, error := eventsClient.DeleteRule(context.Background(), deleteRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting Rule %s %s, It is possible that the resource is already deleted. Please verify manually \n", ruleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ruleId, EventsrulesSweepWaitCondition, time.Duration(3*time.Minute),
				EventsrulesSweepResponseFetchOperation, "events", true)
		}
	}
	return nil
}

func getEventsRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	eventsClient := acctest.GetTestClients(&schema.ResourceData{}).EventsClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RuleId", id)
	}
	return resourceIds, nil
}

func EventsrulesSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ruleResponse, ok := response.Response.(oci_events.GetRuleResponse); ok {
		return ruleResponse.LifecycleState != oci_events.RuleLifecycleStateDeleted
	}
	return false
}

func EventsrulesSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EventsClient().GetRule(context.Background(), oci_events.GetRuleRequest{
		RuleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
