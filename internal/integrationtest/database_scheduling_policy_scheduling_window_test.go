// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseSchedulingPolicySchedulingWindowRequiredOnlyResource = DatabaseSchedulingPolicySchedulingWindowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Required, acctest.Create, DatabaseSchedulingPolicySchedulingWindowRepresentation)

	DatabaseSchedulingPolicySchedulingWindowResourceConfig = DatabaseSchedulingPolicySchedulingWindowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Optional, acctest.Update, DatabaseSchedulingPolicySchedulingWindowRepresentation)

	DatabaseSchedulingPolicySchedulingWindowSingularDataSourceRepresentation = map[string]interface{}{
		"scheduling_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_policy.test_scheduling_policy.id}`},
		"scheduling_window_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_policy_scheduling_window.test_scheduling_policy_scheduling_window.id}`},
	}

	DatabaseSchedulingPolicySchedulingWindowDataSourceRepresentation = map[string]interface{}{
		"scheduling_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_policy.test_scheduling_policy.id}`},
		"compartment_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseSchedulingPolicySchedulingWindowDataSourceFilterRepresentation}}
	DatabaseSchedulingPolicySchedulingWindowDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_scheduling_policy_scheduling_window.test_scheduling_policy_scheduling_window.id}`}},
	}

	DatabaseSchedulingPolicySchedulingWindowRepresentation = map[string]interface{}{
		"scheduling_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_policy.test_scheduling_policy.id}`},
		"window_preference":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseSchedulingPolicySchedulingWindowWindowPreferenceRepresentation},
		"compartment_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	DatabaseSchedulingPolicySchedulingWindowWindowPreferenceRepresentation = map[string]interface{}{
		"days_of_week":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseSchedulingPolicySchedulingWindowWindowPreferenceDaysOfWeekRepresentation},
		"duration":             acctest.Representation{RepType: acctest.Required, Create: `180`, Update: `210`},
		"is_enforced_duration": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"start_time":           acctest.Representation{RepType: acctest.Required, Create: `10:00`, Update: `11:00`},
		"weeks_of_month":       acctest.Representation{RepType: acctest.Required, Create: []string{`1`}, Update: []string{`2`}},
		"months": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: DatabaseSchedulingPolicySchedulingWindowWindowPreferenceMonthsRepresentation},
			{RepType: acctest.Required, Group: DatabaseSchedulingPolicySchedulingWindowWindowPreferenceMonthsRepresentation1},
		},
	}
	DatabaseSchedulingPolicySchedulingWindowWindowPreferenceDaysOfWeekRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `MONDAY`, Update: `TUESDAY`},
	}
	DatabaseSchedulingPolicySchedulingWindowWindowPreferenceMonthsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `FEBRUARY`, Update: `MARCH`},
	}
	DatabaseSchedulingPolicySchedulingWindowWindowPreferenceMonthsRepresentation1 = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `AUGUST`, Update: `SEPTEMBER`},
	}

	DatabaseSchedulingPolicySchedulingWindowResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Required, acctest.Create, DatabaseSchedulingPolicyRepresentation) + DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseSchedulingPolicySchedulingWindowResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseSchedulingPolicySchedulingWindowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_scheduling_policy_scheduling_window.test_scheduling_policy_scheduling_window"
	datasourceName := "data.oci_database_scheduling_policy_scheduling_windows.test_scheduling_policy_scheduling_windows"
	singularDatasourceName := "data.oci_database_scheduling_policy_scheduling_window.test_scheduling_policy_scheduling_window"

	var resId, resId2, compositeId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseSchedulingPolicySchedulingWindowResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Optional, acctest.Create, DatabaseSchedulingPolicySchedulingWindowRepresentation), "database", "schedulingPolicySchedulingWindow", t)

	acctest.ResourceTest(t, testAccCheckDatabaseSchedulingPolicySchedulingWindowDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPolicySchedulingWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Required, acctest.Create, DatabaseSchedulingPolicySchedulingWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.days_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.duration", "180"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.is_enforced_duration", "false"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.start_time", "10:00"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.weeks_of_month.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPolicySchedulingWindowResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPolicySchedulingWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Optional, acctest.Create, DatabaseSchedulingPolicySchedulingWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.days_of_week.0.name", "MONDAY"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.duration", "180"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.is_enforced_duration", "false"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.months.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.months.0.name", "FEBRUARY"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.start_time", "10:00"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.weeks_of_month.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					schedulingPolicyId, _ := acctest.FromInstanceState(s, resourceName, "scheduling_policy_id")
					compositeId = "schedulingPolicies/" + schedulingPolicyId + "/schedulingWindows/" + resId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseSchedulingPolicySchedulingWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Optional, acctest.Update, DatabaseSchedulingPolicySchedulingWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.duration", "210"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.is_enforced_duration", "true"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.months.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.months.0.name", "MARCH"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.start_time", "11:00"),
				resource.TestCheckResourceAttr(resourceName, "window_preference.0.weeks_of_month.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_windows", "test_scheduling_policy_scheduling_windows", acctest.Optional, acctest.Update, DatabaseSchedulingPolicySchedulingWindowDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSchedulingPolicySchedulingWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Optional, acctest.Update, DatabaseSchedulingPolicySchedulingWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_windows.0.display_name"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_windows.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_windows.0.scheduling_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_windows.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_windows.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_windows.0.time_next_scheduling_window_starts"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_windows.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.0.duration", "210"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.0.is_enforced_duration", "true"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.0.months.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.0.months.0.name", "MARCH"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.0.start_time", "11:00"),
				resource.TestCheckResourceAttr(datasourceName, "scheduling_windows.0.window_preference.0.weeks_of_month.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Required, acctest.Create, DatabaseSchedulingPolicySchedulingWindowSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSchedulingPolicySchedulingWindowResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduling_window_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_next_scheduling_window_starts"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.0.days_of_week.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.0.days_of_week.0.name", "TUESDAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.0.duration", "210"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.0.is_enforced_duration", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.0.months.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.0.months.0.name", "MARCH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.0.start_time", "11:00"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_preference.0.weeks_of_month.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseSchedulingPolicySchedulingWindowRequiredOnlyResource,
			ImportState:             true,
			ImportStateIdFunc:       getSchedulingWindowImportId(resourceName),
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseSchedulingPolicySchedulingWindowDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_scheduling_policy_scheduling_window" {
			noResourceFound = false
			request := oci_database.GetSchedulingWindowRequest{}

			if value, ok := rs.Primary.Attributes["scheduling_policy_id"]; ok {
				request.SchedulingPolicyId = &value
			}

			tmp := rs.Primary.ID
			request.SchedulingWindowId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			log.Printf("[DEBUG]: Get call check destroy %v", rs.Primary.ID)
			response, err := client.GetSchedulingWindow(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.SchedulingWindowLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseSchedulingPolicySchedulingWindow") {
		resource.AddTestSweepers("DatabaseSchedulingPolicySchedulingWindow", &resource.Sweeper{
			Name:         "DatabaseSchedulingPolicySchedulingWindow",
			Dependencies: acctest.DependencyGraph["schedulingPolicySchedulingWindow"],
			F:            sweepDatabaseSchedulingPolicySchedulingWindowResource,
		})
	}
}

func sweepDatabaseSchedulingPolicySchedulingWindowResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	schedulingPolicySchedulingWindowIds, err := getDatabaseSchedulingPolicySchedulingWindowIds(compartment)
	if err != nil {
		return err
	}
	for _, schedulingPolicySchedulingWindowId := range schedulingPolicySchedulingWindowIds {
		if ok := acctest.SweeperDefaultResourceId[schedulingPolicySchedulingWindowId]; !ok {
			deleteSchedulingWindowRequest := oci_database.DeleteSchedulingWindowRequest{}

			deleteSchedulingWindowRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteSchedulingWindow(context.Background(), deleteSchedulingWindowRequest)
			if error != nil {
				fmt.Printf("Error deleting SchedulingPolicySchedulingWindow %s %s, It is possible that the resource is already deleted. Please verify manually \n", schedulingPolicySchedulingWindowId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &schedulingPolicySchedulingWindowId, DatabaseSchedulingPolicySchedulingWindowSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseSchedulingPolicySchedulingWindowSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseSchedulingPolicySchedulingWindowIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SchedulingPolicySchedulingWindowId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listSchedulingWindowsRequest := oci_database.ListSchedulingWindowsRequest{}
	listSchedulingWindowsRequest.CompartmentId = &compartmentId

	schedulingPolicyIds, error := getDatabaseSchedulingPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting schedulingPolicyId required for SchedulingPolicySchedulingWindow resource requests \n")
	}
	for _, schedulingPolicyId := range schedulingPolicyIds {
		listSchedulingWindowsRequest.SchedulingPolicyId = &schedulingPolicyId

		listSchedulingWindowsRequest.LifecycleState = oci_database.SchedulingWindowSummaryLifecycleStateAvailable
		listSchedulingWindowsResponse, err := databaseClient.ListSchedulingWindows(context.Background(), listSchedulingWindowsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting SchedulingPolicySchedulingWindow list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, schedulingPolicySchedulingWindow := range listSchedulingWindowsResponse.Items {
			id := *schedulingPolicySchedulingWindow.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SchedulingPolicySchedulingWindowId", id)
		}

	}
	return resourceIds, nil
}

func DatabaseSchedulingPolicySchedulingWindowSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if schedulingPolicySchedulingWindowResponse, ok := response.Response.(oci_database.GetSchedulingWindowResponse); ok {
		return schedulingPolicySchedulingWindowResponse.LifecycleState != oci_database.SchedulingWindowLifecycleStateDeleted
	}
	return false
}

func DatabaseSchedulingPolicySchedulingWindowSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetSchedulingWindow(context.Background(), oci_database.GetSchedulingWindowRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func getSchedulingWindowImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("schedulingPolicies/%s/schedulingWindows/%s", rs.Primary.Attributes["scheduling_policy_id"], rs.Primary.Attributes["id"]), nil
	}
}
