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
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var (
	DatabaseScheduledActionRequiredOnlyResource = DatabaseScheduledActionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduled_action", "test_scheduled_action", acctest.Required, acctest.Create, DatabaseScheduledActionRepresentation)

	DatabaseScheduledActionResourceConfig = DatabaseScheduledActionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduled_action", "test_scheduled_action", acctest.Optional, acctest.Update, DatabaseScheduledActionRepresentation)

	DatabaseScheduledActionSingularDataSourceRepresentation = map[string]interface{}{
		"scheduled_action_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduled_action.test_scheduled_action.id}`},
	}

	DatabaseScheduledActionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_scheduled_action.test_scheduled_action.id}`},
		"scheduling_plan_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_scheduling_plan.test_scheduling_plan.id}`},
		"service_type":       acctest.Representation{RepType: acctest.Optional, Create: `serviceType`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseScheduledActionDataSourceFilterRepresentation}}
	DatabaseScheduledActionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_scheduled_action.test_scheduled_action.id}`}},
	}

	DatabaseScheduledActionRepresentation = map[string]interface{}{
		"action_type":          acctest.Representation{RepType: acctest.Required, Create: `DB_SERVER_FULL_SOFTWARE_UPDATE`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"scheduling_plan_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_plan.test_scheduling_plan.id}`},
		"scheduling_window_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_policy_scheduling_window.test_scheduling_policy_scheduling_window.id}`},
		//"action_members":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseScheduledActionActionMembersRepresentation},
		//"action_params":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"count": "0"}},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	//DatabaseScheduledActionActionMembersRepresentation = map[string]interface{}{
	//	"member_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_database_member.test_member.id}`},
	//	"member_order":           acctest.Representation{RepType: acctest.Required, Create: `1`},
	//	"estimated_time_in_mins": acctest.Representation{RepType: acctest.Optional, Create: `90`},
	//}

	DatabaseScheduledActionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Required, acctest.Create, DatabaseSchedulingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Required, acctest.Create, DatabaseSchedulingPolicySchedulingWindowRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_plan", "test_scheduling_plan", acctest.Required, acctest.Create, DatabaseSchedulingPlanRepresentation) + DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseScheduledActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseScheduledActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_scheduled_action.test_scheduled_action"
	datasourceName := "data.oci_database_scheduled_actions.test_scheduled_actions"
	singularDatasourceName := "data.oci_database_scheduled_action.test_scheduled_action"

	var resId, resId2 string
	//Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseScheduledActionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduled_action", "test_scheduled_action", acctest.Optional, acctest.Create, DatabaseScheduledActionRepresentation), "database", "scheduledAction", t)

	acctest.ResourceTest(t, testAccCheckDatabaseScheduledActionDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseScheduledActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduled_action", "test_scheduled_action", acctest.Required, acctest.Create, DatabaseScheduledActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_plan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_window_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseScheduledActionResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseScheduledActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduled_action", "test_scheduled_action", acctest.Optional, acctest.Create, DatabaseScheduledActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_members.#", "0"),
				//resource.TestCheckResourceAttr(resourceName, "action_members.0.estimated_time_in_mins", "90"),
				//resource.TestCheckResourceAttrSet(resourceName, "action_members.0.member_id"),
				//resource.TestCheckResourceAttr(resourceName, "action_members.0.member_order", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "action_order"),
				//resource.TestCheckResourceAttr(resourceName, "action_params.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_plan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_window_id"),
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
			Config: config + compartmentIdVariableStr + DatabaseScheduledActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduled_action", "test_scheduled_action", acctest.Optional, acctest.Update, DatabaseScheduledActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action_members.#", "0"),
				//resource.TestCheckResourceAttr(resourceName, "action_members.0.estimated_time_in_mins", "90"),
				//resource.TestCheckResourceAttrSet(resourceName, "action_members.0.member_id"),
				//resource.TestCheckResourceAttr(resourceName, "action_members.0.member_order", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "action_order"),
				//resource.TestCheckResourceAttr(resourceName, "action_params.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_plan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "scheduling_window_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduled_actions", "test_scheduled_actions", acctest.Optional, acctest.Update, DatabaseScheduledActionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseScheduledActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_scheduled_action", "test_scheduled_action", acctest.Optional, acctest.Update, DatabaseScheduledActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_plan_id"),
				resource.TestCheckResourceAttr(datasourceName, "service_type", "serviceType"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "scheduled_action_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "scheduled_action_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduled_action", "test_scheduled_action", acctest.Required, acctest.Create, DatabaseScheduledActionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseScheduledActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_action_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action_members.#", "0"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "action_members.0.estimated_time_in_mins", "11"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "action_members.0.member_order", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "action_order"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "action_params.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_time_in_mins"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseScheduledActionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseScheduledActionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_scheduled_action" {
			noResourceFound = false
			request := oci_database.GetScheduledActionRequest{}

			tmp := rs.Primary.ID
			request.ScheduledActionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetScheduledAction(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ScheduledActionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseScheduledAction") {
		resource.AddTestSweepers("DatabaseScheduledAction", &resource.Sweeper{
			Name:         "DatabaseScheduledAction",
			Dependencies: acctest.DependencyGraph["scheduledAction"],
			F:            sweepDatabaseScheduledActionResource,
		})
	}
}

func sweepDatabaseScheduledActionResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	scheduledActionIds, err := getDatabaseScheduledActionIds(compartment)
	if err != nil {
		return err
	}
	for _, scheduledActionId := range scheduledActionIds {
		if ok := acctest.SweeperDefaultResourceId[scheduledActionId]; !ok {
			deleteScheduledActionRequest := oci_database.DeleteScheduledActionRequest{}

			deleteScheduledActionRequest.ScheduledActionId = &scheduledActionId

			deleteScheduledActionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteScheduledAction(context.Background(), deleteScheduledActionRequest)
			if error != nil {
				fmt.Printf("Error deleting ScheduledAction %s %s, It is possible that the resource is already deleted. Please verify manually \n", scheduledActionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &scheduledActionId, DatabaseScheduledActionSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseScheduledActionSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseScheduledActionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ScheduledActionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listScheduledActionsRequest := oci_database.ListScheduledActionsRequest{}
	listScheduledActionsRequest.CompartmentId = &compartmentId
	listScheduledActionsRequest.LifecycleState = oci_database.ScheduledActionSummaryLifecycleStateAvailable
	listScheduledActionsResponse, err := databaseClient.ListScheduledActions(context.Background(), listScheduledActionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ScheduledAction list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, scheduledAction := range listScheduledActionsResponse.Items {
		id := *scheduledAction.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ScheduledActionId", id)
	}
	return resourceIds, nil
}

func DatabaseScheduledActionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if scheduledActionResponse, ok := response.Response.(oci_database.GetScheduledActionResponse); ok {
		return scheduledActionResponse.LifecycleState != oci_database.ScheduledActionLifecycleStateDeleted
	}
	return false
}

func DatabaseScheduledActionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetScheduledAction(context.Background(), oci_database.GetScheduledActionRequest{
		ScheduledActionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
