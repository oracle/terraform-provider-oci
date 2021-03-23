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
	"github.com/oracle/oci-go-sdk/v37/common"
	oci_events "github.com/oracle/oci-go-sdk/v37/events"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RuleRequiredOnlyResource = RuleResourceDependencies +
		generateResourceFromRepresentationMap("oci_events_rule", "test_rule", Required, Create, ruleRepresentation)

	RuleResourceConfig = RuleResourceDependencies +
		generateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Update, ruleRepresentation)

	ruleSingularDataSourceRepresentation = map[string]interface{}{
		"rule_id": Representation{repType: Required, create: `${oci_events_rule.test_rule.id}`},
	}

	ruleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `This rule sends a notification upon completion of DbaaS backup`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `INACTIVE`},
		"filter":         RepresentationGroup{Required, ruleDataSourceFilterRepresentation}}
	ruleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_events_rule.test_rule.id}`}},
	}

	ruleRepresentation = map[string]interface{}{
		"actions":        RepresentationGroup{Required, ruleActionsRepresentation},
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"condition":      Representation{repType: Required, create: `{\"eventType\":\"com.oraclecloud.databaseservice.autonomous.database.backup.end\"}`, update: `{}`},
		"display_name":   Representation{repType: Required, create: `This rule sends a notification upon completion of DbaaS backup`, update: `displayName2`},
		"is_enabled":     Representation{repType: Required, create: `true`, update: `false`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}
	ruleActionsRepresentation = map[string]interface{}{
		"actions": []RepresentationGroup{{Optional, ruleActionsOSSActionsRepresentation}},
	}
	ruleActionsUpdateRepresentation = map[string]interface{}{
		"actions": []RepresentationGroup{{Optional, ruleActionsOSSActionsRepresentation}, {Optional, ruleActionsONSActionsRepresentation}, {Optional, ruleActionsFAASActionsRepresentation}},
	}
	ruleActionsONSActionsRepresentation = map[string]interface{}{
		"action_type": Representation{repType: Required, create: `ONS`, update: `ONS`},
		"is_enabled":  Representation{repType: Required, create: `false`, update: `true`},
		"description": Representation{repType: Optional, create: `description`, update: `description2`},
		"topic_id":    Representation{repType: Optional, create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}
	ruleActionsOSSActionsRepresentation = map[string]interface{}{
		"action_type": Representation{repType: Required, create: `OSS`, update: `ONS`},
		"is_enabled":  Representation{repType: Required, create: `false`, update: `true`},
		"description": Representation{repType: Optional, create: `description`, update: `rule type updated`},
		"stream_id":   Representation{repType: Optional, create: `${oci_streaming_stream.test_stream.id}`, update: ``},
		"topic_id":    Representation{repType: Optional, create: ``, update: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}
	ruleActionsFAASActionsRepresentation = map[string]interface{}{
		"action_type": Representation{repType: Required, create: `FAAS`, update: `FAAS`},
		"is_enabled":  Representation{repType: Required, create: `false`, update: `true`},
		"description": Representation{repType: Optional, create: `description`, update: `description2`},
		"function_id": Representation{repType: Optional, create: `${oci_functions_function.test_function.id}`},
	}

	RuleResourceDependencies = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
		generateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation) +
		generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamRepresentation)
)

func TestEventsRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEventsRuleResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
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
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+RuleResourceDependencies+
		generateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Create, ruleRepresentation), "events", "rule", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckEventsRuleDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + imageVariableStr + RuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Create, ruleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "This rule sends a notification upon completion of DbaaS backup"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + imageVariableStr + compartmentIdUVariableStr + RuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Create,
						representationCopyWithNewProperties(ruleRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "This rule sends a notification upon completion of DbaaS backup"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Update,
						getUpdatedRepresentationCopy("actions", RepresentationGroup{Optional, ruleActionsUpdateRepresentation}, ruleRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_events_rules", "test_rules", Optional, Update, ruleDataSourceRepresentation) +
					compartmentIdVariableStr + imageVariableStr + RuleResourceDependencies +
					generateResourceFromRepresentationMap("oci_events_rule", "test_rule", Optional, Update, ruleRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "rules.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "rules.0.condition", "{}"),
					resource.TestCheckResourceAttr(datasourceName, "rules.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_events_rule", "test_rule", Required, Create, ruleSingularDataSourceRepresentation) +
					compartmentIdVariableStr + imageVariableStr + RuleResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "rule_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "actions.0.actions.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "condition", "{}"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "events")

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
	if !inSweeperExcludeList("EventsRule") {
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

			deleteRuleRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "events")
			_, error := eventsClient.DeleteRule(context.Background(), deleteRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting Rule %s %s, It is possible that the resource is already deleted. Please verify manually \n", ruleId, error)
				continue
			}
			waitTillCondition(testAccProvider, &ruleId, ruleSweepWaitCondition, time.Duration(3*time.Minute),
				ruleSweepResponseFetchOperation, "events", true)
		}
	}
	return nil
}

func getRuleIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "RuleId")
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
		addResourceIdToSweeperResourceIdMap(compartmentId, "RuleId", id)
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
