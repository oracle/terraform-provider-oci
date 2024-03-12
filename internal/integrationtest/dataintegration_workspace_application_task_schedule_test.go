// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataintegrationWorkspaceApplicationTaskScheduleRequiredOnlyResource = DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedule", "test_workspace_application_task_schedule", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationTaskScheduleRepresentation)

	DataintegrationWorkspaceApplicationTaskScheduleResourceConfig = DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedule", "test_workspace_application_task_schedule", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationTaskScheduleRepresentation)

	DataintegrationWorkspaceApplicationTaskScheduleSingularDataSourceRepresentation = map[string]interface{}{
		"application_key":   acctest.Representation{RepType: acctest.Required, Create: `${var.application_key}`},
		"task_schedule_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_application_task_schedule.test_workspace_application_task_schedule.key}`},
		"workspace_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
	}

	DataintegrationWorkspaceApplicationTaskScheduleDataSourceRepresentation = map[string]interface{}{
		"application_key": acctest.Representation{RepType: acctest.Required, Create: `${var.application_key}`},
		"workspace_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
		"identifier":      acctest.Representation{RepType: acctest.Optional, Create: []string{`REST_20240207_111223_3365404`}},
		"is_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `REST_20240207_111223_3365404`, Update: `REST_20240207_111223_3365404`},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`type`}},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceApplicationTaskScheduleDataSourceFilterRepresentation},
	}

	DataintegrationWorkspaceApplicationTaskScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_application_task_schedule.test_workspace_application_task_schedule.name}`}},
	}

	DataintegrationWorkspaceApplicationTaskScheduleRepresentation = map[string]interface{}{
		"application_key":        acctest.Representation{RepType: acctest.Required, Create: `${var.application_key}`},
		"identifier":             acctest.Representation{RepType: acctest.Required, Create: `REST_20240207_111223_3365404`, Update: `REST_20240207_111223_3365404`},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `REST_20240207_111223_3365404`, Update: `REST_20240207_111223_3365404`},
		"workspace_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.workspace_id}`},
		"auth_mode":              acctest.Representation{RepType: acctest.Optional, Create: `RESOURCE_PRINCIPAL`, Update: `RESOURCE_PRINCIPAL`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description`},
		"expected_duration":      acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.0`},
		"end_time_millis":        acctest.Representation{RepType: acctest.Optional, Create: `1741201716000`, Update: `1741201716000`},
		"expected_duration_unit": acctest.Representation{RepType: acctest.Optional, Create: `SECONDS`, Update: `MINUTES`},
		"is_backfill_enabled":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_concurrent_allowed":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"model_version":          acctest.Representation{RepType: acctest.Optional, Create: `20210408`, Update: `20210408`},
		"number_of_retries":      acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `3`},
		"object_status":          acctest.Representation{RepType: acctest.Optional, Create: `8`, Update: `8`},
		"registry_metadata":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceApplicationTaskScheduleRegistryMetadataRepresentation},
		"retry_delay":            acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"retry_delay_unit":       acctest.Representation{RepType: acctest.Optional, Create: `SECONDS`, Update: `MINUTES`},
		"schedule_ref":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceApplicationTaskScheduleScheduleRefRepresentation},
	}

	DataintegrationWorkspaceApplicationTaskScheduleRegistryMetadataRepresentation = map[string]interface{}{
		"aggregator_key":   acctest.Representation{RepType: acctest.Required, Create: `d30b2adf-69e5-4b01-9600-3053413ade09`},
		"is_favorite":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"registry_version": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	DataintegrationWorkspaceApplicationTaskScheduleScheduleRefRepresentation = map[string]interface{}{
		"identifier": acctest.Representation{RepType: acctest.Optional, Create: `SCHEDULE_TEMP`, Update: `SCHEDULE_TEMP`},
		"key":        acctest.Representation{RepType: acctest.Required, Create: `17dfe4a1-180d-4e59-9315-73e40a14e3e5`, Update: `17dfe4a1-180d-4e59-9315-73e40a14e3e5`},
		"model_type": acctest.Representation{RepType: acctest.Required, Create: `SCHEDULE`, Update: `SCHEDULE`},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `SCHEDULE_TEMP`, Update: `SCHEDULE_TEMP`},
		"timezone":   acctest.Representation{RepType: acctest.Optional, Create: `GMT`, Update: `GMT`},
	}

	DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies = ""
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceApplicationTaskScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceApplicationTaskScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	workspaceId := utils.GetEnvSettingWithBlankDefault("workspace_id")
	workspaceIdVariableStr := fmt.Sprintf("variable \"workspace_id\" { default = \"%s\" }\n", workspaceId)

	applicationKey := utils.GetEnvSettingWithBlankDefault("application_key")
	applicationKeyVariableStr := fmt.Sprintf("variable \"application_key\" { default = \"%s\" }\n", applicationKey)

	aggregatorKey := utils.GetEnvSettingWithBlankDefault("aggregator_key")
	aggregatorKeyVariableStr := fmt.Sprintf("variable \"aggregator_key\" { default = \"%s\" }\n", aggregatorKey)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId) + workspaceIdVariableStr + applicationKeyVariableStr + aggregatorKeyVariableStr

	resourceName := "oci_dataintegration_workspace_application_task_schedule.test_workspace_application_task_schedule"
	datasourceName := "data.oci_dataintegration_workspace_application_task_schedules.test_workspace_application_task_schedules"
	singularDatasourceName := "data.oci_dataintegration_workspace_application_task_schedule.test_workspace_application_task_schedule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedule", "test_workspace_application_task_schedule", acctest.Optional, acctest.Create, DataintegrationWorkspaceApplicationTaskScheduleRepresentation), "dataintegration", "workspaceApplicationTaskSchedule", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceApplicationTaskScheduleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedule", "test_workspace_application_task_schedule", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationTaskScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_key"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttr(resourceName, "name", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedule", "test_workspace_application_task_schedule", acctest.Optional, acctest.Create, DataintegrationWorkspaceApplicationTaskScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_key"),
				resource.TestCheckResourceAttr(resourceName, "auth_mode", "RESOURCE_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "config_provider_delegate", "{\"bindings\":null,\"childProviders\":null}"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "expected_duration", "1"),
				resource.TestCheckResourceAttr(resourceName, "expected_duration_unit", "SECONDS"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttr(resourceName, "is_backfill_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_concurrent_allowed", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20210408"),
				resource.TestCheckResourceAttr(resourceName, "number_of_retries", "2"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "object_version", "2"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.aggregator_key"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.description", ""),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.aggregator_key"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "false"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.registry_version", "0"),
				resource.TestCheckResourceAttr(resourceName, "retry_delay", "1"),
				resource.TestCheckResourceAttr(resourceName, "retry_delay_unit", "SECONDS"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.is_daylight_adjustment_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.days.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.model_type", "MONTHLY"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.time.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.time.0.hour", "0"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.time.0.minute", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "start_time_millis"),
				resource.TestCheckResourceAttr(resourceName, "end_time_millis", "1741201716000"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.timezone", "GMT"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

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
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedule", "test_workspace_application_task_schedule", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationTaskScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_key"),
				resource.TestCheckResourceAttr(resourceName, "auth_mode", "RESOURCE_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(resourceName, "config_provider_delegate"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "expected_duration", "1"),
				resource.TestCheckResourceAttr(resourceName, "expected_duration_unit", "MINUTES"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttr(resourceName, "is_backfill_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_concurrent_allowed", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "object_version", "3"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20210408"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.description", ""),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.aggregator_key"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_metadata.0.aggregator_key"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.is_favorite", "true"),
				resource.TestCheckResourceAttr(resourceName, "registry_metadata.0.registry_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.days.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.interval", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.model_type", "MONTHLY"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.time.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.time.0.hour", "0"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.frequency_details.0.time.0.minute", "0"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.parent_ref.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "name", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.is_daylight_adjustment_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "schedule_ref.0.object_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "retry_delay", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedules", "test_workspace_application_task_schedules", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationTaskScheduleDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceApplicationTaskScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedule", "test_workspace_application_task_schedule", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationTaskScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "application_key"),
				resource.TestCheckResourceAttr(datasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "identifier.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "task_schedule_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "task_schedule_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_application_task_schedule", "test_workspace_application_task_schedule", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationTaskScheduleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceApplicationTaskScheduleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "task_schedule_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auth_mode", "RESOURCE_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "config_provider_delegate"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_version", "20210408"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_version", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retry_delay", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retry_delay_unit", "MINUTES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_ref.0.frequency_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_ref.0.frequency_details.0.days.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_ref.0.frequency_details.0.interval", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_ref.0.frequency_details.0.model_type", "MONTHLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_ref.0.frequency_details.0.time.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_ref.0.frequency_details.0.time.0.hour", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_ref.0.frequency_details.0.time.0.minute", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule_ref.0.timezone", "GMT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expected_duration", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "expected_duration_unit", "MINUTES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identifier", "REST_20240207_111223_3365404"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_backfill_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_concurrent_allowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "last_run_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "TASK_SCHEDULE"),
			),
		},
		// verify resource import
		{
			Config:            config + DataintegrationWorkspaceApplicationTaskScheduleRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"number_of_retries",
				"registry_metadata",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceApplicationTaskScheduleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_application_task_schedule" {
			noResourceFound = false
			request := oci_dataintegration.GetTaskScheduleRequest{}

			if value, ok := rs.Primary.Attributes["application_key"]; ok {
				request.ApplicationKey = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.TaskScheduleKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			_, err := client.GetTaskSchedule(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceApplicationTaskSchedule") {
		resource.AddTestSweepers("DataintegrationWorkspaceApplicationTaskSchedule", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceApplicationTaskSchedule",
			Dependencies: acctest.DependencyGraph["workspaceApplicationTaskSchedule"],
			F:            sweepDataintegrationWorkspaceApplicationTaskScheduleResource,
		})
	}
}

func sweepDataintegrationWorkspaceApplicationTaskScheduleResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceApplicationTaskScheduleIds, err := getDataintegrationWorkspaceApplicationTaskScheduleIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceApplicationTaskScheduleId := range workspaceApplicationTaskScheduleIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceApplicationTaskScheduleId]; !ok {
			deleteTaskScheduleRequest := oci_dataintegration.DeleteTaskScheduleRequest{}

			deleteTaskScheduleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteTaskSchedule(context.Background(), deleteTaskScheduleRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceApplicationTaskSchedule %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceApplicationTaskScheduleId, error)
				continue
			}
		}
	}
	return nil
}

func getDataintegrationWorkspaceApplicationTaskScheduleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceApplicationTaskScheduleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listTaskSchedulesRequest := oci_dataintegration.ListTaskSchedulesRequest{}

	workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceApplicationTaskSchedule resource requests \n")
	}
	for _, workspaceId := range workspaceIds {
		listTaskSchedulesRequest.WorkspaceId = &workspaceId

		listTaskSchedulesResponse, err := dataIntegrationClient.ListTaskSchedules(context.Background(), listTaskSchedulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting WorkspaceApplicationTaskSchedule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, workspaceApplicationTaskSchedule := range listTaskSchedulesResponse.Items {
			id := *workspaceApplicationTaskSchedule.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceApplicationTaskScheduleId", id)
		}

	}
	return resourceIds, nil
}
