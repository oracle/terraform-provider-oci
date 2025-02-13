// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
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
	DatabaseExecutionWindowRequiredOnlyResource = DatabaseExecutionWindowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Required, acctest.Create, DatabaseExecutionWindowRepresentation)

	DatabaseExecutionWindowResourceConfig = DatabaseExecutionWindowResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Optional, acctest.Update, DatabaseExecutionWindowRepresentation)

	DatabaseExecutionWindowSingularDataSourceRepresentation = map[string]interface{}{
		"execution_window_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_execution_window.test_execution_window.id}`},
	}

	DatabaseExecutionWindowDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"execution_resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_resource.test_resource.id}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExecutionWindowDataSourceFilterRepresentation}}
	DatabaseExecutionWindowDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_execution_window.test_execution_window.id}`}},
	}

	DatabaseExecutionWindowRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"execution_resource_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_maintenance_run.test_maintenance_run.id}`},
		"time_scheduled":          acctest.Representation{RepType: acctest.Required, Create: `2018-12-23T01:59:07.030Z`, Update: `timeScheduled2`},
		"window_duration_in_mins": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enforced_duration":    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	DatabaseExecutionWindowResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseExecutionWindowResource_basic(t *testing.T) {
	t.Skip("Skip this test as this cannot be tested. Exempted - ORM-141938")
	httpreplay.SetScenario("TestDatabaseExecutionWindowResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_execution_window.test_execution_window"
	datasourceName := "data.oci_database_execution_windows.test_execution_windows"
	singularDatasourceName := "data.oci_database_execution_window.test_execution_window"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExecutionWindowResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Optional, acctest.Create, DatabaseExecutionWindowRepresentation), "database", "executionWindow", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExecutionWindowDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExecutionWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Required, acctest.Create, DatabaseExecutionWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "execution_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "time_scheduled", "2018-12-23T01:59:07.030Z"),
				resource.TestCheckResourceAttr(resourceName, "window_duration_in_mins", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExecutionWindowResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseExecutionWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Optional, acctest.Create, DatabaseExecutionWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_duration", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "time_scheduled", "2018-12-23T01:59:07.030Z"),
				resource.TestCheckResourceAttr(resourceName, "window_duration_in_mins", "10"),

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
			Config: config + compartmentIdVariableStr + DatabaseExecutionWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Optional, acctest.Update, DatabaseExecutionWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enforced_duration", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "time_scheduled", "timeScheduled2"),
				resource.TestCheckResourceAttr(resourceName, "window_duration_in_mins", "11"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_execution_windows", "test_execution_windows", acctest.Optional, acctest.Update, DatabaseExecutionWindowDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExecutionWindowResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Optional, acctest.Update, DatabaseExecutionWindowRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "execution_windows.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "execution_windows.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.estimated_time_in_mins"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.execution_resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "execution_windows.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "execution_windows.0.is_enforced_duration", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.lifecycle_substate"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.time_ended"),
				resource.TestCheckResourceAttr(datasourceName, "execution_windows.0.time_scheduled", "timeScheduled2"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.total_time_taken_in_mins"),
				resource.TestCheckResourceAttr(datasourceName, "execution_windows.0.window_duration_in_mins", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_windows.0.window_type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Required, acctest.Create, DatabaseExecutionWindowSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExecutionWindowResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_window_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_time_in_mins"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enforced_duration", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_substate"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_scheduled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_time_taken_in_mins"),
				resource.TestCheckResourceAttr(singularDatasourceName, "window_duration_in_mins", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "window_type"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseExecutionWindowRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseExecutionWindowDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_execution_window" {
			noResourceFound = false
			request := oci_database.GetExecutionWindowRequest{}

			tmp := rs.Primary.ID
			request.ExecutionWindowId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetExecutionWindow(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExecutionWindowLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseExecutionWindow") {
		resource.AddTestSweepers("DatabaseExecutionWindow", &resource.Sweeper{
			Name:         "DatabaseExecutionWindow",
			Dependencies: acctest.DependencyGraph["executionWindow"],
			F:            sweepDatabaseExecutionWindowResource,
		})
	}
}

func sweepDatabaseExecutionWindowResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	executionWindowIds, err := getDatabaseExecutionWindowIds(compartment)
	if err != nil {
		return err
	}
	for _, executionWindowId := range executionWindowIds {
		if ok := acctest.SweeperDefaultResourceId[executionWindowId]; !ok {
			deleteExecutionWindowRequest := oci_database.DeleteExecutionWindowRequest{}

			deleteExecutionWindowRequest.ExecutionWindowId = &executionWindowId

			deleteExecutionWindowRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExecutionWindow(context.Background(), deleteExecutionWindowRequest)
			if error != nil {
				fmt.Printf("Error deleting ExecutionWindow %s %s, It is possible that the resource is already deleted. Please verify manually \n", executionWindowId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &executionWindowId, DatabaseExecutionWindowSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseExecutionWindowSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseExecutionWindowIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExecutionWindowId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listExecutionWindowsRequest := oci_database.ListExecutionWindowsRequest{}
	listExecutionWindowsRequest.CompartmentId = &compartmentId
	listExecutionWindowsRequest.LifecycleState = oci_database.ExecutionWindowSummaryLifecycleStateCreated
	listExecutionWindowsResponse, err := databaseClient.ListExecutionWindows(context.Background(), listExecutionWindowsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExecutionWindow list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, executionWindow := range listExecutionWindowsResponse.Items {
		id := *executionWindow.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExecutionWindowId", id)
	}
	return resourceIds, nil
}

func DatabaseExecutionWindowSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if executionWindowResponse, ok := response.Response.(oci_database.GetExecutionWindowResponse); ok {
		return executionWindowResponse.LifecycleState != oci_database.ExecutionWindowLifecycleStateDeleted
	}
	return false
}

func DatabaseExecutionWindowSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetExecutionWindow(context.Background(), oci_database.GetExecutionWindowRequest{
		ExecutionWindowId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
