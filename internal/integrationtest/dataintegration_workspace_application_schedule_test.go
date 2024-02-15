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
	DataintegrationWorkspaceApplicationScheduleRequiredOnlyResource = DataintegrationWorkspaceApplicationScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_schedule", "test_workspace_application_schedule", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationScheduleRepresentation)

	DataintegrationWorkspaceApplicationScheduleResourceConfig = DataintegrationWorkspaceApplicationScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_schedule", "test_workspace_application_schedule", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationScheduleRepresentation)

	DataintegrationWorkspaceApplicationScheduleSingularDataSourceRepresentation = map[string]interface{}{
		"workspace_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"application_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_application.test_workspace_application.key}`},
		"schedule_key":    acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_application_schedule.test_workspace_application_schedule.key}`},
	}

	DataintegrationWorkspaceApplicationScheduleDataSourceRepresentation = map[string]interface{}{
		"application_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_application.test_workspace_application.key}`},
		"workspace_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"identifier":      acctest.Representation{RepType: acctest.Optional, Create: []string{`TERSI_TEST_SCHEDULE`}},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `TERSI_TEST_SCHEDULE`, Update: `TERSI_TEST_SCHEDULE_2`},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`SCHEDULE`}},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceApplicationScheduleDataSourceFilterRepresentation}}
	DataintegrationWorkspaceApplicationScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataintegration_workspace_application_schedule.test_workspace_application_schedule.name}`}},
	}

	DataintegrationWorkspaceApplicationScheduleRepresentation = map[string]interface{}{
		"application_key":                acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace_application.test_workspace_application.key}`},
		"identifier":                     acctest.Representation{RepType: acctest.Required, Create: `TERSI_TEST_SCHEDULE`},
		"name":                           acctest.Representation{RepType: acctest.Required, Create: `TERSI_TEST_SCHEDULE`, Update: `TERSI_TEST_SCHEDULE_2`},
		"workspace_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_dataintegration_workspace.test_workspace.id}`},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"frequency_details":              acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceApplicationScheduleFrequencyDetailsRepresentation},
		"is_daylight_adjustment_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"model_version":                  acctest.Representation{RepType: acctest.Optional, Create: `20210408`},
		"object_status":                  acctest.Representation{RepType: acctest.Optional, Create: `8`},
		"timezone":                       acctest.Representation{RepType: acctest.Required, Create: `UTC`, Update: `Asia/Kolkata`},
	}
	DataintegrationWorkspaceApplicationScheduleFrequencyDetailsRepresentation = map[string]interface{}{
		"model_type": acctest.Representation{RepType: acctest.Required, Create: `HOURLY`, Update: `DAILY`},
		"interval":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"time":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DataintegrationWorkspaceApplicationScheduleFrequencyDetailsTimeRepresentation},
	}
	DataintegrationWorkspaceApplicationScheduleFrequencyDetailsTimeRepresentation = map[string]interface{}{
		"hour":   acctest.Representation{RepType: acctest.Required, Create: `0`, Update: `11`},
		"minute": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	DataintegrationWorkspaceApplicationScheduleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace", "test_workspace", acctest.Required, acctest.Create, DataintegrationWorkspaceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application", "test_workspace_application", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataintegration/default
func TestDataintegrationWorkspaceApplicationScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataintegrationWorkspaceApplicationScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_dataintegration_workspace_application_schedule.test_workspace_application_schedule"
	datasourceName := "data.oci_dataintegration_workspace_application_schedules.test_workspace_application_schedules"
	singularDatasourceName := "data.oci_dataintegration_workspace_application_schedule.test_workspace_application_schedule"
	fmt.Printf("value  isss")
	fmt.Printf(compartmentIdVariableStr)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataintegrationWorkspaceApplicationScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_schedule", "test_workspace_application_schedule", acctest.Optional, acctest.Create, DataintegrationWorkspaceApplicationScheduleRepresentation), "dataintegration", "workspaceApplicationSchedule", t)

	acctest.ResourceTest(t, testAccCheckDataintegrationWorkspaceApplicationScheduleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_schedule", "test_workspace_application_schedule", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_key"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "TERSI_TEST_SCHEDULE"),
				resource.TestCheckResourceAttr(resourceName, "name", "TERSI_TEST_SCHEDULE"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationScheduleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_schedule", "test_workspace_application_schedule", acctest.Optional, acctest.Create, DataintegrationWorkspaceApplicationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_key"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.interval", "10"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.model_type", "HOURLY"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.time.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.time.0.hour", "0"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.time.0.minute", "10"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "TERSI_TEST_SCHEDULE"),
				resource.TestCheckResourceAttr(resourceName, "is_daylight_adjustment_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "SCHEDULE"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20210408"),
				resource.TestCheckResourceAttr(resourceName, "name", "TERSI_TEST_SCHEDULE"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "UTC"),
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
			Config: config + compartmentIdVariableStr + DataintegrationWorkspaceApplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_schedule", "test_workspace_application_schedule", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "application_key"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.interval", "11"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.model_type", "DAILY"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.time.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.time.0.hour", "11"),
				resource.TestCheckResourceAttr(resourceName, "frequency_details.0.time.0.minute", "11"),
				resource.TestCheckResourceAttr(resourceName, "identifier", "TERSI_TEST_SCHEDULE"),
				resource.TestCheckResourceAttr(resourceName, "is_daylight_adjustment_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "20210408"),
				resource.TestCheckResourceAttr(resourceName, "name", "TERSI_TEST_SCHEDULE_2"),
				resource.TestCheckResourceAttr(resourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(resourceName, "object_version", "2"),
				resource.TestCheckResourceAttr(resourceName, "timezone", "Asia/Kolkata"),
				resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Printf("update success")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_application_schedules", "test_workspace_application_schedules", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationScheduleDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceApplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataintegration_workspace_application_schedule", "test_workspace_application_schedule", acctest.Optional, acctest.Update, DataintegrationWorkspaceApplicationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "application_key"),
				resource.TestCheckResourceAttr(datasourceName, "identifier.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "name", "TERSI_TEST_SCHEDULE_2"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(datasourceName, "schedule_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schedule_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataintegration_workspace_application_schedule", "test_workspace_application_schedule", acctest.Required, acctest.Create, DataintegrationWorkspaceApplicationScheduleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataintegrationWorkspaceApplicationScheduleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "workspace_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "frequency_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "frequency_details.0.interval", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "frequency_details.0.model_type", "DAILY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "frequency_details.0.time.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "frequency_details.0.time.0.hour", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "frequency_details.0.time.0.minute", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identifier", "TERSI_TEST_SCHEDULE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_daylight_adjustment_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "SCHEDULE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_version", "20210408"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TERSI_TEST_SCHEDULE_2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_status", "8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object_version", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timezone", "Asia/Kolkata"),
			),
		},
		// verify resource import
		{
			Config:            config + DataintegrationWorkspaceApplicationScheduleRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"registry_metadata",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDataintegrationWorkspaceApplicationScheduleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataIntegrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataintegration_workspace_application_schedule" {
			noResourceFound = false
			request := oci_dataintegration.GetScheduleRequest{}

			if value, ok := rs.Primary.Attributes["application_key"]; ok {
				request.ApplicationKey = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ScheduleKey = &value
			}

			if value, ok := rs.Primary.Attributes["workspace_id"]; ok {
				request.WorkspaceId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")

			_, err := client.GetSchedule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DataintegrationWorkspaceApplicationSchedule") {
		resource.AddTestSweepers("DataintegrationWorkspaceApplicationSchedule", &resource.Sweeper{
			Name:         "DataintegrationWorkspaceApplicationSchedule",
			Dependencies: acctest.DependencyGraph["workspaceApplicationSchedule"],
			F:            sweepDataintegrationWorkspaceApplicationScheduleResource,
		})
	}
}

func sweepDataintegrationWorkspaceApplicationScheduleResource(compartment string) error {
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()
	workspaceApplicationScheduleIds, err := getDataintegrationWorkspaceApplicationScheduleIds(compartment)
	if err != nil {
		return err
	}
	for _, workspaceApplicationScheduleId := range workspaceApplicationScheduleIds {
		if ok := acctest.SweeperDefaultResourceId[workspaceApplicationScheduleId]; !ok {
			deleteScheduleRequest := oci_dataintegration.DeleteScheduleRequest{}

			deleteScheduleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataintegration")
			_, error := dataIntegrationClient.DeleteSchedule(context.Background(), deleteScheduleRequest)
			if error != nil {
				fmt.Printf("Error deleting WorkspaceApplicationSchedule %s %s, It is possible that the resource is already deleted. Please verify manually \n", workspaceApplicationScheduleId, error)
				continue
			}
		}
	}
	return nil
}

func getDataintegrationWorkspaceApplicationScheduleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WorkspaceApplicationScheduleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataIntegrationClient := acctest.GetTestClients(&schema.ResourceData{}).DataIntegrationClient()

	listSchedulesRequest := oci_dataintegration.ListSchedulesRequest{}

	applicationKeys, error := getDataintegrationWorkspaceApplicationIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting applicationKey required for WorkspaceApplicationSchedule resource requests \n")
	}
	for _, applicationKey := range applicationKeys {
		listSchedulesRequest.ApplicationKey = &applicationKey

		workspaceIds, error := getDataintegrationWorkspaceIds(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting workspaceId required for WorkspaceApplicationSchedule resource requests \n")
		}
		for _, workspaceId := range workspaceIds {
			listSchedulesRequest.WorkspaceId = &workspaceId

			listSchedulesResponse, err := dataIntegrationClient.ListSchedules(context.Background(), listSchedulesRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting WorkspaceApplicationSchedule list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, workspaceApplicationSchedule := range listSchedulesResponse.Items {
				id := *workspaceApplicationSchedule.Key
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WorkspaceApplicationScheduleId", id)
			}

		}
	}
	return resourceIds, nil
}
