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
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FleetAppsManagementFleetRequiredOnlyResource = FleetAppsManagementFleetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation)

	FleetAppsManagementFleetResourceConfig = FleetAppsManagementFleetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Update, FleetAppsManagementFleetRepresentation)

	FleetAppsManagementFleetSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
	}

	FleetAppsManagementFleetDataSourceRepresentation = map[string]interface{}{
		"application_type": acctest.Representation{RepType: acctest.Optional, Create: `applicationType`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"environment_type": acctest.Representation{RepType: acctest.Optional, Create: `environmentType`},
		"fleet_type":       acctest.Representation{RepType: acctest.Optional, Create: `GENERIC`},
		"id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `NEEDS_ATTENTION`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetDataSourceFilterRepresentation}}
	FleetAppsManagementFleetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_fleet.test_fleet.id}`}},
	}

	FleetAppsManagementFleetRepresentation = map[string]interface{}{
		"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: fleetIgnoreChangesRecipeRepresentation},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"fleet_type":               acctest.Representation{RepType: acctest.Required, Create: `GENERIC`},
		"application_type":         acctest.Representation{RepType: acctest.Optional, Create: `applicationType`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"environment_type":         acctest.Representation{RepType: acctest.Optional, Create: `environmentType`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"group_type":               acctest.Representation{RepType: acctest.Optional, Create: `ENVIRONMENT`},
		"is_target_auto_confirm":   acctest.Representation{RepType: acctest.Required, Create: `true`},
		"notification_preferences": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementFleetNotificationPreferencesRepresentation},
		"products":                 acctest.Representation{RepType: acctest.Optional, Create: []string{"OS(COMPUTE)"}},
		"resource_selection_type":  acctest.Representation{RepType: acctest.Required, Create: `MANUAL`},
		"rule_selection_criteria":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementFleetRuleSelectionCriteriaRepresentation},
	}
	FleetAppsManagementFleetNotificationPreferencesRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"topic_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"preferences":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementFleetNotificationPreferencesPreferencesRepresentation},
	}
	FleetAppsManagementFleetRuleSelectionCriteriaRepresentation = map[string]interface{}{
		"match_condition": acctest.Representation{RepType: acctest.Required, Create: `MATCH_ALL`, Update: `ANY`},
		"rules":           acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetRuleSelectionCriteriaRulesRepresentation},
	}
	FleetAppsManagementFleetNotificationPreferencesPreferencesRepresentation = map[string]interface{}{
		"on_job_failure":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_topology_modification": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"on_upcoming_schedule":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	FleetAppsManagementFleetRuleSelectionCriteriaRulesRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"conditions":              acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementFleetRuleSelectionCriteriaRulesConditionsRepresentation},
		"resource_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"basis":                   acctest.Representation{RepType: acctest.Optional, Create: `basis`, Update: `basis2`},
	}
	FleetAppsManagementFleetRuleSelectionCriteriaRulesConditionsRepresentation = map[string]interface{}{
		"attr_group": acctest.Representation{RepType: acctest.Required, Create: `attrGroup`, Update: `attrGroup2`},
		"attr_key":   acctest.Representation{RepType: acctest.Required, Create: `attrKey`, Update: `attrKey2`},
		"attr_value": acctest.Representation{RepType: acctest.Required, Create: `attrValue`, Update: `attrValue2`},
	}

	FleetAppsManagementFleetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)

	fleetIgnoreChangesRecipeRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementFleetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementFleetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_fleet.test_fleet"
	datasourceName := "data.oci_fleet_apps_management_fleets.test_fleets"
	singularDatasourceName := "data.oci_fleet_apps_management_fleet.test_fleet"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementFleetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Create, FleetAppsManagementFleetRepresentation), "fleetappsmanagement", "fleet", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementFleetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "fleet_type", "GENERIC"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Create, FleetAppsManagementFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_type", "applicationType"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "environment_type", "environmentType"),
				resource.TestCheckResourceAttr(resourceName, "fleet_type", "GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_type", "ENVIRONMENT"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_target_auto_confirm", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_failure", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_topology_modification", "false"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_upcoming_schedule", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_preferences.0.topic_id"),
				resource.TestCheckResourceAttr(resourceName, "products.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection_type", "MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.match_condition", "MATCH_ALL"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.basis", "basis"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_group", "attrGroup"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_key", "attrKey"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_value", "attrValue"),
				resource.TestCheckResourceAttrSet(resourceName, "rule_selection_criteria.0.rules.0.resource_compartment_id"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Update, FleetAppsManagementFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_type", "applicationType"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "environment_type", "environmentType"),
				resource.TestCheckResourceAttr(resourceName, "fleet_type", "GENERIC"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_type", "ENVIRONMENT"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_target_auto_confirm", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_job_failure", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_topology_modification", "true"),
				resource.TestCheckResourceAttr(resourceName, "notification_preferences.0.preferences.0.on_upcoming_schedule", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_preferences.0.topic_id"),
				resource.TestCheckResourceAttr(resourceName, "products.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resource_selection_type", "MANUAL"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.match_condition", "ANY"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.basis", "basis2"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_group", "attrGroup2"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_key", "attrKey2"),
				resource.TestCheckResourceAttr(resourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_value", "attrValue2"),
				resource.TestCheckResourceAttrSet(resourceName, "rule_selection_criteria.0.rules.0.resource_compartment_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleets", "test_fleets", acctest.Optional, acctest.Update, FleetAppsManagementFleetDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Optional, acctest.Update, FleetAppsManagementFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "application_type", "applicationType"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "environment_type", "environmentType"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_type", "GENERIC"),
				resource.TestCheckResourceAttr(datasourceName, "state", "NEEDS_ATTENTION"),

				resource.TestCheckResourceAttr(datasourceName, "fleet_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementFleetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "application_type", "applicationType"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "environment_type", "environmentType"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleet_type", "GENERIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "group_type", "ENVIRONMENT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_target_auto_confirm", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_job_failure", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_topology_modification", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_preferences.0.preferences.0.on_upcoming_schedule", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "products.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_selection_type", "MANUAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.0.match_condition", "ANY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.0.rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.0.rules.0.basis", "basis2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.0.rules.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.0.rules.0.conditions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_group", "attrGroup2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_key", "attrKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rule_selection_criteria.0.rules.0.conditions.0.attr_value", "attrValue2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementFleetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementFleetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_fleet" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetFleetRequest{}

			tmp := rs.Primary.ID
			request.FleetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetFleet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.FleetLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementFleet") {
		resource.AddTestSweepers("FleetAppsManagementFleet", &resource.Sweeper{
			Name:         "FleetAppsManagementFleet",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepFleetAppsManagementFleetResource,
		})
	}
}

func sweepFleetAppsManagementFleetResource(compartment string) error {
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()
	fleetIds, err := getFleetAppsManagementFleetIds(compartment)
	if err != nil {
		return err
	}
	for _, fleetId := range fleetIds {
		if ok := acctest.SweeperDefaultResourceId[fleetId]; !ok {
			deleteFleetRequest := oci_fleet_apps_management.DeleteFleetRequest{}

			deleteFleetRequest.FleetId = &fleetId

			deleteFleetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementClient.DeleteFleet(context.Background(), deleteFleetRequest)
			if error != nil {
				fmt.Printf("Error deleting Fleet %s %s, It is possible that the resource is already deleted. Please verify manually \n", fleetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fleetId, FleetAppsManagementFleetSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementFleetSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementFleetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FleetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementClient()

	listFleetsRequest := oci_fleet_apps_management.ListFleetsRequest{}
	listFleetsRequest.CompartmentId = &compartmentId
	listFleetsRequest.LifecycleState = oci_fleet_apps_management.FleetLifecycleStateActive
	listFleetsResponse, err := fleetAppsManagementClient.ListFleets(context.Background(), listFleetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Fleet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, fleet := range listFleetsResponse.Items {
		id := *fleet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FleetId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementFleetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fleetResponse, ok := response.Response.(oci_fleet_apps_management.GetFleetResponse); ok {
		return fleetResponse.LifecycleState != oci_fleet_apps_management.FleetLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementFleetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementClient().GetFleet(context.Background(), oci_fleet_apps_management.GetFleetRequest{
		FleetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
