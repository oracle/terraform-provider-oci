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
	FleetAppsManagementSchedulerDefinitionRequiredOnlyResource = FleetAppsManagementSchedulerDefinitionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Required, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation)

	FleetAppsManagementSchedulerDefinitionResourceConfig = FleetAppsManagementSchedulerDefinitionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Update, FleetAppsManagementSchedulerDefinitionRepresentation)

	FleetAppsManagementSchedulerDefinitionSingularDataSourceRepresentation = map[string]interface{}{
		"scheduler_definition_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_scheduler_definition.test_scheduler_definition.id}`},
	}

	FleetAppsManagementSchedulerDefinitionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fleet_id":              acctest.Representation{RepType: acctest.Optional, Create: `${var.test_active_fleet}`},
		"maintenance_window_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_maintenance_window.test_maintenance_window.id}`},
		"product":               acctest.Representation{RepType: acctest.Optional, Create: `WEBLOGIC/JAVA`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementSchedulerDefinitionDataSourceFilterRepresentation}}
	FleetAppsManagementSchedulerDefinitionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_scheduler_definition.test_scheduler_definition.id}`}},
	}

	FleetAppsManagementSchedulerDefinitionRepresentation = map[string]interface{}{
		"action_groups":               acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementSchedulerDefinitionActionGroupsRepresentation},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"schedule":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementSchedulerDefinitionScheduleRepresentation},
		"activity_initiation_cut_off": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"run_books":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementSchedulerDefinitionRunBooksRepresentation},
	}
	FleetAppsManagementSchedulerDefinitionActionGroupsRepresentation = map[string]interface{}{
		"resource_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.test_active_fleet}`},
		"runbook_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.test_runbook_ocid}`},
		"lifecycle_operation": acctest.Representation{RepType: acctest.Required, Create: `PATCH`},
		"product":             acctest.Representation{RepType: acctest.Required, Create: `WEBLOGIC/JAVA`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `PRODUCT`},
	}
	FleetAppsManagementSchedulerDefinitionScheduleRepresentation = map[string]interface{}{
		"execution_startdate":   acctest.Representation{RepType: acctest.Required, Create: `2024-08-31T00:00:00.000Z`, Update: `2024-09-30T00:00:00.000Z`},
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`, Update: `MAINTENANCE_WINDOW`},
		"duration":              acctest.Representation{RepType: acctest.Required, Create: `PT2H`, Update: `PT3H`},
		"maintenance_window_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_fleet_apps_management_maintenance_window.test_maintenance_window.id}`},
	}
	FleetAppsManagementSchedulerDefinitionRunBooksRepresentation = map[string]interface{}{
		"id":               acctest.Representation{RepType: acctest.Required, Create: `${var.test_runbook_ocid}`},
		"input_parameters": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementSchedulerDefinitionRunBooksInputParametersRepresentation},
	}
	FleetAppsManagementSchedulerDefinitionRunBooksInputParametersRepresentation = map[string]interface{}{
		"step_name": acctest.Representation{RepType: acctest.Required, Create: `stepName`, Update: `stepName2`},
		"arguments": acctest.RepresentationGroup{RepType: acctest.Optional, Group: FleetAppsManagementSchedulerDefinitionRunBooksInputParametersArgumentsRepresentation},
	}
	FleetAppsManagementSchedulerDefinitionRunBooksInputParametersArgumentsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}

	FleetAppsManagementSchedulerDefinitionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, FleetAppsManagementMaintenanceWindowRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementSchedulerDefinitionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementSchedulerDefinitionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// Runbooks are currently created by Oracle, and read-only. There is no Create API.
	runbookId := utils.GetEnvSettingWithBlankDefault("test_runbook_ocid")
	testRunbookStr := fmt.Sprintf("variable \"test_runbook_ocid\" { default = \"%s\" }\n", runbookId)

	// Fleet in ACTIVE state. Fleets require a confirmation action call not supported by Terraform to go active.
	// Thus, this needs to be created and confirmed manually.
	activeFleetId := utils.GetEnvSettingWithBlankDefault("test_active_fleet")
	activeFleetStr := fmt.Sprintf("variable \"test_active_fleet\" { default = \"%s\" }\n", activeFleetId)

	resourceName := "oci_fleet_apps_management_scheduler_definition.test_scheduler_definition"
	datasourceName := "data.oci_fleet_apps_management_scheduler_definitions.test_scheduler_definitions"
	singularDatasourceName := "data.oci_fleet_apps_management_scheduler_definition.test_scheduler_definition"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementSchedulerDefinitionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation), "fleetappsmanagement", "schedulerDefinition", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementSchedulerDefinitionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + activeFleetStr + testRunbookStr + compartmentIdVariableStr + FleetAppsManagementSchedulerDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Required, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_groups.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.runbook_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.type", "CUSTOM"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + activeFleetStr + testRunbookStr + compartmentIdVariableStr + FleetAppsManagementSchedulerDefinitionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + activeFleetStr + testRunbookStr + compartmentIdVariableStr + FleetAppsManagementSchedulerDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Create, FleetAppsManagementSchedulerDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.lifecycle_operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.product", "WEBLOGIC/JAVA"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.runbook_id"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.type", "PRODUCT"),
				resource.TestCheckResourceAttr(resourceName, "activity_initiation_cut_off", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "run_books.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.id"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.step_name", "stepName"),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.duration", "PT2H"),
				resource.TestCheckResourceAttrSet(resourceName, "schedule.0.maintenance_window_id"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.type", "CUSTOM"),
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
			Config: config + activeFleetStr + testRunbookStr + compartmentIdVariableStr + FleetAppsManagementSchedulerDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Update, FleetAppsManagementSchedulerDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.lifecycle_operation", "PATCH"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.product", "WEBLOGIC/JAVA"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "action_groups.0.runbook_id"),
				resource.TestCheckResourceAttr(resourceName, "action_groups.0.type", "PRODUCT"),
				resource.TestCheckResourceAttr(resourceName, "activity_initiation_cut_off", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "run_books.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "run_books.0.id"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.arguments.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "run_books.0.input_parameters.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.duration", "PT3H"),
				resource.TestCheckResourceAttrSet(resourceName, "schedule.0.maintenance_window_id"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.type", "MAINTENANCE_WINDOW"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definitions", "test_scheduler_definitions", acctest.Optional, acctest.Update, FleetAppsManagementSchedulerDefinitionDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + testRunbookStr + FleetAppsManagementSchedulerDefinitionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Optional, acctest.Update, FleetAppsManagementSchedulerDefinitionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "fleet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_window_id"),
				resource.TestCheckResourceAttr(datasourceName, "product", "WEBLOGIC/JAVA"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "scheduler_definition_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "scheduler_definition_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_scheduler_definition", "test_scheduler_definition", acctest.Required, acctest.Create, FleetAppsManagementSchedulerDefinitionSingularDataSourceRepresentation) +
				activeFleetStr + compartmentIdVariableStr + testRunbookStr + FleetAppsManagementSchedulerDefinitionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduler_definition_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action_group_types.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_groups.0.lifecycle_operation", "PATCH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_groups.0.product", "WEBLOGIC/JAVA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_groups.0.type", "PRODUCT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "activity_initiation_cut_off", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "application_types.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "count_of_affected_action_groups"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "count_of_affected_resources"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "count_of_affected_targets"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifecycle_operations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "products.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_books.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.arguments.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "run_books.0.input_parameters.0.step_name", "stepName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule.0.duration", "PT3H"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule.0.execution_startdate"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule.0.type", "MAINTENANCE_WINDOW"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_next_run"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementSchedulerDefinitionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementSchedulerDefinitionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementOperationsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_scheduler_definition" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetSchedulerDefinitionRequest{}

			tmp := rs.Primary.ID
			request.SchedulerDefinitionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetSchedulerDefinition(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementSchedulerDefinition") {
		resource.AddTestSweepers("FleetAppsManagementSchedulerDefinition", &resource.Sweeper{
			Name:         "FleetAppsManagementSchedulerDefinition",
			Dependencies: acctest.DependencyGraph["schedulerDefinition"],
			F:            sweepFleetAppsManagementSchedulerDefinitionResource,
		})
	}
}

func sweepFleetAppsManagementSchedulerDefinitionResource(compartment string) error {
	fleetAppsManagementOperationsClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementOperationsClient()
	schedulerDefinitionIds, err := getFleetAppsManagementSchedulerDefinitionIds(compartment)
	if err != nil {
		return err
	}
	for _, schedulerDefinitionId := range schedulerDefinitionIds {
		if ok := acctest.SweeperDefaultResourceId[schedulerDefinitionId]; !ok {
			deleteSchedulerDefinitionRequest := oci_fleet_apps_management.DeleteSchedulerDefinitionRequest{}

			deleteSchedulerDefinitionRequest.SchedulerDefinitionId = &schedulerDefinitionId

			deleteSchedulerDefinitionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementOperationsClient.DeleteSchedulerDefinition(context.Background(), deleteSchedulerDefinitionRequest)
			if error != nil {
				fmt.Printf("Error deleting SchedulerDefinition %s %s, It is possible that the resource is already deleted. Please verify manually \n", schedulerDefinitionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &schedulerDefinitionId, FleetAppsManagementSchedulerDefinitionSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementSchedulerDefinitionSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementSchedulerDefinitionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SchedulerDefinitionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementOperationsClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementOperationsClient()

	listSchedulerDefinitionsRequest := oci_fleet_apps_management.ListSchedulerDefinitionsRequest{}
	listSchedulerDefinitionsRequest.CompartmentId = &compartmentId
	listSchedulerDefinitionsRequest.LifecycleState = oci_fleet_apps_management.SchedulerDefinitionLifecycleStateActive
	listSchedulerDefinitionsResponse, err := fleetAppsManagementOperationsClient.ListSchedulerDefinitions(context.Background(), listSchedulerDefinitionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SchedulerDefinition list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, schedulerDefinition := range listSchedulerDefinitionsResponse.Items {
		id := *schedulerDefinition.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SchedulerDefinitionId", id)
	}
	return resourceIds, nil
}

func FleetAppsManagementSchedulerDefinitionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if schedulerDefinitionResponse, ok := response.Response.(oci_fleet_apps_management.GetSchedulerDefinitionResponse); ok {
		return schedulerDefinitionResponse.LifecycleState != oci_fleet_apps_management.SchedulerDefinitionLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementSchedulerDefinitionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementOperationsClient().GetSchedulerDefinition(context.Background(), oci_fleet_apps_management.GetSchedulerDefinitionRequest{
		SchedulerDefinitionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
