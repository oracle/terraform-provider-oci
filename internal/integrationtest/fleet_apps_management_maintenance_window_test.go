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
	FleetAppsManagementMaintenanceWindowRequiredOnlyResource = FleetAppsManagementMaintenanceWindowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, FleetAppsManagementMaintenanceWindowRepresentation)

	FleetAppsManagementMaintenanceWindowResourceConfig = FleetAppsManagementMaintenanceWindowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Update, FleetAppsManagementMaintenanceWindowRepresentation)

	FleetAppsManagementMaintenanceWindowSingularDataSourceRepresentation = map[string]interface{}{
		"maintenance_window_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_maintenance_window.test_maintenance_window.id}`},
	}

	FleetAppsManagementMaintenanceWindowDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.tenancy_ocid}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: FleetAppsManagementMaintenanceWindowDataSourceFilterRepresentation}}
	FleetAppsManagementMaintenanceWindowDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_fleet_apps_management_maintenance_window.test_maintenance_window.id}`}},
	}

	FleetAppsManagementMaintenanceWindowRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"duration":                acctest.Representation{RepType: acctest.Required, Create: `PT2H`, Update: `PT1H`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `MWTT20240808`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_outage":               acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_recurring":            acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"maintenance_window_type": acctest.Representation{RepType: acctest.Required, Create: `OPEN_ENDED`},
		"recurrences":             acctest.Representation{RepType: acctest.Optional, Create: `recurrences`, Update: `recurrences2`},
		"task_initiation_cutoff":  acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `11`},
		"time_schedule_start":     acctest.Representation{RepType: acctest.Required, Create: `2024-08-08T17:43:32.292Z`, Update: `2024-12-09T17:43:32.292Z`},
	}

	FleetAppsManagementMaintenanceWindowResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementMaintenanceWindowResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementMaintenanceWindowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_fleet_apps_management_maintenance_window.test_maintenance_window"
	datasourceName := "data.oci_fleet_apps_management_maintenance_windows.test_maintenance_windows"
	singularDatasourceName := "data.oci_fleet_apps_management_maintenance_window.test_maintenance_window"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FleetAppsManagementMaintenanceWindowResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Create, FleetAppsManagementMaintenanceWindowRepresentation), "fleetappsmanagement", "maintenanceWindow", t)

	acctest.ResourceTest(t, testAccCheckFleetAppsManagementMaintenanceWindowDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementMaintenanceWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, FleetAppsManagementMaintenanceWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "duration", "PT2H"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementMaintenanceWindowResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FleetAppsManagementMaintenanceWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Create, FleetAppsManagementMaintenanceWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MWTT20240808"),
				resource.TestCheckResourceAttr(resourceName, "duration", "PT2H"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_outage", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_recurring", "false"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_type", "OPEN_ENDED"),
				resource.TestCheckResourceAttr(resourceName, "recurrences", "recurrences"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "task_initiation_cutoff", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_schedule_start", "2024-08-08T17:43:32.292Z"),

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
			Config: config + compartmentIdVariableStr + FleetAppsManagementMaintenanceWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Update, FleetAppsManagementMaintenanceWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "duration", "PT1H"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_outage", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_recurring", "true"),
				resource.TestCheckResourceAttr(resourceName, "maintenance_window_type", "OPEN_ENDED"),
				resource.TestCheckResourceAttr(resourceName, "recurrences", "recurrences2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "task_initiation_cutoff", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_schedule_start", "2024-12-09T17:43:32.292Z"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_maintenance_windows", "test_maintenance_windows", acctest.Optional, acctest.Update, FleetAppsManagementMaintenanceWindowDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementMaintenanceWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Optional, acctest.Update, FleetAppsManagementMaintenanceWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "maintenance_window_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_window_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_maintenance_window", "test_maintenance_window", acctest.Required, acctest.Create, FleetAppsManagementMaintenanceWindowSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementMaintenanceWindowResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_window_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "duration", "PT1H"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_outage", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_recurring", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "maintenance_window_type", "OPEN_ENDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recurrences", "recurrences2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "task_initiation_cutoff", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_schedule_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + FleetAppsManagementMaintenanceWindowRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFleetAppsManagementMaintenanceWindowDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FleetAppsManagementMaintenanceWindowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_fleet_apps_management_maintenance_window" {
			noResourceFound = false
			request := oci_fleet_apps_management.GetMaintenanceWindowRequest{}

			tmp := rs.Primary.ID
			request.MaintenanceWindowId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")

			response, err := client.GetMaintenanceWindow(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_fleet_apps_management.MaintenanceWindowLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FleetAppsManagementMaintenanceWindow") {
		resource.AddTestSweepers("FleetAppsManagementMaintenanceWindow", &resource.Sweeper{
			Name:         "FleetAppsManagementMaintenanceWindow",
			Dependencies: acctest.DependencyGraph["maintenanceWindow"],
			F:            sweepFleetAppsManagementMaintenanceWindowResource,
		})
	}
}

func sweepFleetAppsManagementMaintenanceWindowResource(compartment string) error {
	fleetAppsManagementMaintenanceWindowClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementMaintenanceWindowClient()
	maintenanceWindowIds, err := getFleetAppsManagementMaintenanceWindowIds(compartment)
	if err != nil {
		return err
	}
	for _, maintenanceWindowId := range maintenanceWindowIds {
		if ok := acctest.SweeperDefaultResourceId[maintenanceWindowId]; !ok {
			deleteMaintenanceWindowRequest := oci_fleet_apps_management.DeleteMaintenanceWindowRequest{}

			deleteMaintenanceWindowRequest.MaintenanceWindowId = &maintenanceWindowId

			deleteMaintenanceWindowRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "fleet_apps_management")
			_, error := fleetAppsManagementMaintenanceWindowClient.DeleteMaintenanceWindow(context.Background(), deleteMaintenanceWindowRequest)
			if error != nil {
				fmt.Printf("Error deleting MaintenanceWindow %s %s, It is possible that the resource is already deleted. Please verify manually \n", maintenanceWindowId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &maintenanceWindowId, FleetAppsManagementMaintenanceWindowSweepWaitCondition, time.Duration(3*time.Minute),
				FleetAppsManagementMaintenanceWindowSweepResponseFetchOperation, "fleet_apps_management", true)
		}
	}
	return nil
}

func getFleetAppsManagementMaintenanceWindowIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MaintenanceWindowId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fleetAppsManagementMaintenanceWindowClient := acctest.GetTestClients(&schema.ResourceData{}).FleetAppsManagementMaintenanceWindowClient()

	listMaintenanceWindowsRequest := oci_fleet_apps_management.ListMaintenanceWindowsRequest{}
	listMaintenanceWindowsRequest.CompartmentId = &compartmentId
	listMaintenanceWindowsRequest.LifecycleState = oci_fleet_apps_management.MaintenanceWindowLifecycleStateActive
	listMaintenanceWindowsResponse, err := fleetAppsManagementMaintenanceWindowClient.ListMaintenanceWindows(context.Background(), listMaintenanceWindowsRequest)

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

func FleetAppsManagementMaintenanceWindowSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if maintenanceWindowResponse, ok := response.Response.(oci_fleet_apps_management.GetMaintenanceWindowResponse); ok {
		return maintenanceWindowResponse.LifecycleState != oci_fleet_apps_management.MaintenanceWindowLifecycleStateDeleted
	}
	return false
}

func FleetAppsManagementMaintenanceWindowSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FleetAppsManagementMaintenanceWindowClient().GetMaintenanceWindow(context.Background(), oci_fleet_apps_management.GetMaintenanceWindowRequest{
		MaintenanceWindowId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
