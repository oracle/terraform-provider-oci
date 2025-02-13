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
	DatabaseExecutionActionRequiredOnlyResource = DatabaseExecutionActionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_execution_action", "test_execution_action", acctest.Required, acctest.Create, DatabaseExecutionActionRepresentation)

	DatabaseExecutionActionResourceConfig = DatabaseExecutionActionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_execution_action", "test_execution_action", acctest.Optional, acctest.Update, DatabaseExecutionActionRepresentation)

	DatabaseExecutionActionSingularDataSourceRepresentation = map[string]interface{}{
		"execution_action_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_execution_action.test_execution_action.id}`},
	}

	DatabaseExecutionActionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"execution_window_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_execution_window.test_execution_window.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExecutionActionDataSourceFilterRepresentation}}
	DatabaseExecutionActionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_execution_action.test_execution_action.id}`}},
	}

	DatabaseExecutionActionRepresentation = map[string]interface{}{
		"action_type":         acctest.Representation{RepType: acctest.Required, Create: `DB_SERVER_FULL_SOFTWARE_UPDATE`},
		"execution_window_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_execution_window.test_execution_window.id}`},
		"action_members":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseExecutionActionActionMembersRepresentation},
		"action_params":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"count": "3"}},
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	DatabaseExecutionActionActionMembersRepresentation = map[string]interface{}{
		"member_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_database_member.test_member.id}`},
		"member_order":             acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `11`},
		"estimated_time_in_mins":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"status":                   acctest.Representation{RepType: acctest.Optional, Create: `status`, Update: `status2`},
		"total_time_taken_in_mins": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	DatabaseExecutionActionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_execution_window", "test_execution_window", acctest.Required, acctest.Create, DatabaseExecutionWindowRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseExecutionActionResource_basic(t *testing.T) {
	t.Skip("Skip this test as this cannot be tested. Exempted - ORM-141938")
	httpreplay.SetScenario("TestDatabaseExecutionActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_execution_action.test_execution_action"
	datasourceName := "data.oci_database_execution_actions.test_execution_actions"
	singularDatasourceName := "data.oci_database_execution_action.test_execution_action"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExecutionActionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_execution_action", "test_execution_action", acctest.Optional, acctest.Create, DatabaseExecutionActionRepresentation), "database", "executionAction", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExecutionActionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExecutionActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_execution_action", "test_execution_action", acctest.Required, acctest.Create, DatabaseExecutionActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_window_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExecutionActionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseExecutionActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_execution_action", "test_execution_action", acctest.Optional, acctest.Create, DatabaseExecutionActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_members.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_members.0.estimated_time_in_mins", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "action_members.0.member_id"),
				resource.TestCheckResourceAttr(resourceName, "action_members.0.member_order", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_members.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "action_members.0.total_time_taken_in_mins", "10"),
				resource.TestCheckResourceAttr(resourceName, "action_params.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_window_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + DatabaseExecutionActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_execution_action", "test_execution_action", acctest.Optional, acctest.Update, DatabaseExecutionActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_members.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_members.0.estimated_time_in_mins", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "action_members.0.member_id"),
				resource.TestCheckResourceAttr(resourceName, "action_members.0.member_order", "11"),
				resource.TestCheckResourceAttr(resourceName, "action_members.0.status", "status2"),
				resource.TestCheckResourceAttr(resourceName, "action_members.0.total_time_taken_in_mins", "11"),
				resource.TestCheckResourceAttr(resourceName, "action_params.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "execution_window_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_execution_actions", "test_execution_actions", acctest.Optional, acctest.Update, DatabaseExecutionActionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExecutionActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_execution_action", "test_execution_action", acctest.Optional, acctest.Update, DatabaseExecutionActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_window_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "execution_actions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.action_members.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.action_members.0.estimated_time_in_mins", "11"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.action_members.0.member_id"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.action_members.0.member_order", "11"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.action_members.0.status", "status2"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.action_members.0.total_time_taken_in_mins", "11"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.action_params.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.estimated_time_in_mins"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.execution_action_order"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.execution_window_id"),
				resource.TestCheckResourceAttr(datasourceName, "execution_actions.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.lifecycle_substate"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "execution_actions.0.total_time_taken_in_mins"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_execution_action", "test_execution_action", acctest.Required, acctest.Create, DatabaseExecutionActionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExecutionActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_action_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action_members.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_members.0.estimated_time_in_mins", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_members.0.member_order", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_members.0.status", "status2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_members.0.total_time_taken_in_mins", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_params.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_time_in_mins"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_action_order"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_substate"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_time_taken_in_mins"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseExecutionActionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseExecutionActionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_execution_action" {
			noResourceFound = false
			request := oci_database.GetExecutionActionRequest{}

			tmp := rs.Primary.ID
			request.ExecutionActionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetExecutionAction(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExecutionActionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseExecutionAction") {
		resource.AddTestSweepers("DatabaseExecutionAction", &resource.Sweeper{
			Name:         "DatabaseExecutionAction",
			Dependencies: acctest.DependencyGraph["executionAction"],
			F:            sweepDatabaseExecutionActionResource,
		})
	}
}

func sweepDatabaseExecutionActionResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	executionActionIds, err := getDatabaseExecutionActionIds(compartment)
	if err != nil {
		return err
	}
	for _, executionActionId := range executionActionIds {
		if ok := acctest.SweeperDefaultResourceId[executionActionId]; !ok {
			deleteExecutionActionRequest := oci_database.DeleteExecutionActionRequest{}

			deleteExecutionActionRequest.ExecutionActionId = &executionActionId

			deleteExecutionActionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExecutionAction(context.Background(), deleteExecutionActionRequest)
			if error != nil {
				fmt.Printf("Error deleting ExecutionAction %s %s, It is possible that the resource is already deleted. Please verify manually \n", executionActionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &executionActionId, DatabaseExecutionActionSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseExecutionActionSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseExecutionActionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExecutionActionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listExecutionActionsRequest := oci_database.ListExecutionActionsRequest{}
	listExecutionActionsRequest.CompartmentId = &compartmentId
	listExecutionActionsRequest.LifecycleState = oci_database.ExecutionActionSummaryLifecycleStateSucceeded
	listExecutionActionsResponse, err := databaseClient.ListExecutionActions(context.Background(), listExecutionActionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExecutionAction list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, executionAction := range listExecutionActionsResponse.Items {
		id := *executionAction.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExecutionActionId", id)
	}
	return resourceIds, nil
}

func DatabaseExecutionActionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if executionActionResponse, ok := response.Response.(oci_database.GetExecutionActionResponse); ok {
		return executionActionResponse.LifecycleState != oci_database.ExecutionActionLifecycleStateDeleted
	}
	return false
}

func DatabaseExecutionActionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetExecutionAction(context.Background(), oci_database.GetExecutionActionRequest{
		ExecutionActionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
