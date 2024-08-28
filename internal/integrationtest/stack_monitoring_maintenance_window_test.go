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
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringMaintenanceWindowRequiredOnlyResource = StackMonitoringMaintenanceWindowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, StackMonitoringMaintenanceWindowRepresentation)

	StackMonitoringMaintenanceWindowResourceConfig = StackMonitoringMaintenanceWindowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Update, StackMonitoringMaintenanceWindowRepresentation)

	StackMonitoringMaintenanceWindowSingularDataSourceRepresentation = map[string]interface{}{
		"maintenance_window_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_maintenance_window.test_maintenance_window.id}`},
	}

	StackMonitoringMaintenanceWindowDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `TFMaintenanceWindowsTest`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMaintenanceWindowDataSourceFilterRepresentation}}
	StackMonitoringMaintenanceWindowDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_maintenance_window.test_maintenance_window.id}`}},
	}

	StackMonitoringMaintenanceWindowRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `TFMaintenanceWindowsTest`},
		"resources":      acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMaintenanceWindowResourcesRepresentation},
		"schedule":       acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMaintenanceWindowScheduleRepresentation},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}

	StackMonitoringMaintenanceWindowResourcesRepresentation = map[string]interface{}{
		"resource_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.resource_id}`},
		"are_members_included": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
	}

	StackMonitoringMaintenanceWindowScheduleRepresentation = map[string]interface{}{
		"schedule_type":                  acctest.Representation{RepType: acctest.Required, Create: `ONE_TIME`, Update: `RECURRENT`},
		"time_maintenance_window_start":  acctest.Representation{RepType: acctest.Required, Create: `${var.start_date}`, Update: `2024-10-17T10:47:01.001Z`},
		"time_maintenance_window_end":    acctest.Representation{RepType: acctest.Required, Create: `${var.end_date}`, Update: `2024-10-28T10:47:01.001Z`},
		"maintenance_window_duration":    acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `PT1H`},
		"maintenance_window_recurrences": acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `FREQ=DAILY;BYHOUR=10`},
	}

	StackMonitoringMaintenanceWindowResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMaintenanceWindowResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMaintenanceWindowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	resourceId := utils.GetEnvSettingWithBlankDefault("resource_id_for_maintwin")
	resourceIdVariableStr := fmt.Sprintf("variable \"resource_id\" { default = \"%s\" }\n", resourceId)

	defaultSartTime := time.Now().Add(time.Hour * 24 * 2)
	maintWinStartDate := utils.GetEnvSettingWithDefault("start_date_for_maintwin", defaultSartTime.Format(time.RFC3339Nano))
	maintWinStartDateVariableStr := fmt.Sprintf("variable \"start_date\" { default = \"%s\" }\n", maintWinStartDate)

	defaultEndTime := defaultSartTime.Add(time.Hour * 24)
	maintWinEndDate := utils.GetEnvSettingWithDefault("end_date_for_maintwin", defaultEndTime.Format(time.RFC3339Nano))
	maintWinEndDateVariableStr := fmt.Sprintf("variable \"end_date\" { default = \"%s\" }\n", maintWinEndDate)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_stack_monitoring_maintenance_window.test_maintenance_window"
	datasourceName := "data.oci_stack_monitoring_maintenance_windows.test_maintenance_windows"
	singularDatasourceName := "data.oci_stack_monitoring_maintenance_window.test_maintenance_window"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringMaintenanceWindowResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Create, StackMonitoringMaintenanceWindowRepresentation), "stackmonitoring", "maintenanceWindow", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringMaintenanceWindowDestroy, []resource.TestStep{
		// verify Create
		// one time schedule
		{
			Config: config + compartmentIdVariableStr + resourceIdVariableStr + maintWinStartDateVariableStr + maintWinEndDateVariableStr + StackMonitoringMaintenanceWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, StackMonitoringMaintenanceWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "TFMaintenanceWindowsTest"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "resources.0.resource_id"),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.schedule_type", "ONE_TIME"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMaintenanceWindowResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + resourceIdVariableStr + maintWinStartDateVariableStr +
				maintWinEndDateVariableStr + StackMonitoringMaintenanceWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Create, StackMonitoringMaintenanceWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "TFMaintenanceWindowsTest"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.are_members_included", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "resources.0.resource_id"),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.schedule_type", "ONE_TIME"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.time_maintenance_window_start", maintWinStartDate),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.time_maintenance_window_end", maintWinEndDate),

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
			Config: config + compartmentIdVariableStr + resourceIdVariableStr + maintWinStartDateVariableStr +
				maintWinEndDateVariableStr + StackMonitoringMaintenanceWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Update, StackMonitoringMaintenanceWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "TFMaintenanceWindowsTest"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.are_members_included", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "resources.0.resource_id"),
				resource.TestCheckResourceAttr(resourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.schedule_type", "RECURRENT"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.time_maintenance_window_start", "2024-10-17T10:47:01.001Z"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.time_maintenance_window_end", "2024-10-28T10:47:01.001Z"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.maintenance_window_duration", "PT1H"),
				resource.TestCheckResourceAttr(resourceName, "schedule.0.maintenance_window_recurrences", "FREQ=DAILY;BYHOUR=10"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_maintenance_windows", "test_maintenance_windows", acctest.Optional, acctest.Update, StackMonitoringMaintenanceWindowDataSourceRepresentation) +
				compartmentIdVariableStr + resourceIdVariableStr + maintWinStartDateVariableStr +
				maintWinEndDateVariableStr + StackMonitoringMaintenanceWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Update, StackMonitoringMaintenanceWindowRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "TFMaintenanceWindowsTest"),

				resource.TestCheckResourceAttr(datasourceName, "maintenance_window_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_window_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, StackMonitoringMaintenanceWindowSingularDataSourceRepresentation) +
				compartmentIdVariableStr + resourceIdVariableStr + maintWinStartDateVariableStr + maintWinEndDateVariableStr +
				StackMonitoringMaintenanceWindowResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TFMaintenanceWindowsTest"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources.0.are_members_included", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schedule.0.schedule_type", "RECURRENT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule.0.time_maintenance_window_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule.0.time_maintenance_window_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule.0.maintenance_window_duration"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule.0.maintenance_window_recurrences"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + StackMonitoringMaintenanceWindowRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckStackMonitoringMaintenanceWindowDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_maintenance_window" {
			noResourceFound = false
			request := oci_stack_monitoring.GetMaintenanceWindowRequest{}

			tmp := rs.Primary.ID
			request.MaintenanceWindowId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetMaintenanceWindow(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.MaintenanceWindowLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("StackMonitoringMaintenanceWindow") {
		resource.AddTestSweepers("StackMonitoringMaintenanceWindow", &resource.Sweeper{
			Name:         "StackMonitoringMaintenanceWindow",
			Dependencies: acctest.DependencyGraph["maintenanceWindow"],
			F:            sweepStackMonitoringMaintenanceWindowResource,
		})
	}
}

func sweepStackMonitoringMaintenanceWindowResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	maintenanceWindowIds, err := getStackMonitoringMaintenanceWindowIds(compartment)
	if err != nil {
		return err
	}
	for _, maintenanceWindowId := range maintenanceWindowIds {
		if ok := acctest.SweeperDefaultResourceId[maintenanceWindowId]; !ok {
			deleteMaintenanceWindowRequest := oci_stack_monitoring.DeleteMaintenanceWindowRequest{}

			deleteMaintenanceWindowRequest.MaintenanceWindowId = &maintenanceWindowId

			deleteMaintenanceWindowRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteMaintenanceWindow(context.Background(), deleteMaintenanceWindowRequest)
			if error != nil {
				fmt.Printf("Error deleting MaintenanceWindow %s %s, It is possible that the resource is already deleted. Please verify manually \n", maintenanceWindowId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &maintenanceWindowId, StackMonitoringMaintenanceWindowSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringMaintenanceWindowSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringMaintenanceWindowIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MaintenanceWindowId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listMaintenanceWindowsRequest := oci_stack_monitoring.ListMaintenanceWindowsRequest{}
	listMaintenanceWindowsRequest.CompartmentId = &compartmentId
	listMaintenanceWindowsResponse, err := stackMonitoringClient.ListMaintenanceWindows(context.Background(), listMaintenanceWindowsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MaintenanceWindow list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, maintenanceWindow := range listMaintenanceWindowsResponse.Items {
		id := *maintenanceWindow.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MaintenanceWindowId", id)
	}
	return resourceIds, nil
}

func StackMonitoringMaintenanceWindowSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if maintenanceWindowResponse, ok := response.Response.(oci_stack_monitoring.GetMaintenanceWindowResponse); ok {
		return maintenanceWindowResponse.LifecycleState != oci_stack_monitoring.MaintenanceWindowLifecycleStateDeleted
	}
	return false
}

func StackMonitoringMaintenanceWindowSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetMaintenanceWindow(context.Background(), oci_stack_monitoring.GetMaintenanceWindowRequest{
		MaintenanceWindowId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
