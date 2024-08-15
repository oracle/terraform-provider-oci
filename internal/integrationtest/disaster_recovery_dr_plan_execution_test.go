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
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DisasterRecoveryDrPlanExecutionRequiredOnlyResource = DisasterRecoveryDrPlanExecutionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan_execution", "test_dr_plan_execution", acctest.Required, acctest.Create, DisasterRecoveryDrPlanExecutionRepresentation)

	DisasterRecoveryDrPlanExecutionResourceConfig = DisasterRecoveryDrPlanExecutionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan_execution", "test_dr_plan_execution", acctest.Optional, acctest.Update, DisasterRecoveryDrPlanExecutionRepresentation)

	DisasterRecoveryDisasterRecoveryDrPlanExecutionSingularDataSourceRepresentation = map[string]interface{}{
		"dr_plan_execution_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_dr_plan_execution.test_dr_plan_execution.id}`},
	}

	//Dr Plan Data source
	DisasterRecoveryDisasterRecoveryDrPlanExecutionDataSourceRepresentation = map[string]interface{}{
		"dr_protection_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_dr_protection_group.test_peer.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `Precheck Switchover from PHX to IAD`, Update: `displayName2`},
		"dr_plan_execution_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_disaster_recovery_dr_plan_execution.test_dr_plan_execution.id}`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryDrPlanExecutionDataSourceFilterRepresentation}}
	DisasterRecoveryDrPlanExecutionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_disaster_recovery_dr_plan_execution.test_dr_plan_execution.id}`}},
	}

	DisasterRecoveryDrPlanExecutionRepresentation = map[string]interface{}{
		"execution_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryDrPlanExecutionExecutionOptionsRepresentation},
		"plan_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_dr_plan.test_dr_plan.id}`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `Precheck Switchover from PHX to IAD`, Update: `displayName2`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DefinedTagsIgnoreRepresentation},
	}
	DisasterRecoveryDrPlanExecutionExecutionOptionsRepresentation = map[string]interface{}{
		"plan_execution_type":   acctest.Representation{RepType: acctest.Required, Create: `SWITCHOVER_PRECHECK`},
		"are_prechecks_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"are_warnings_ignored":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DisasterRecoveryDrPlanExecutionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_peer", acctest.Optional, acctest.Create, DisasterRecoveryPeerDrProtectionGroupRepresentation) +
		ObjectStorageBucketDependencyConfig +
		VolumeGroupDependencyConfig +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	DrProtectionGroupConfig = acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Create, DisasterRecoveryDrProtectionGroupRepresentation)

	DrProtectionGroupWithDisassociateTriggerConfig = acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_dr_protection_group", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DisasterRecoveryDrProtectionGroupRepresentation, map[string]interface{}{
		"disassociate_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}))

	DrPlanConfig = acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Required, acctest.Create, DisasterRecoveryDrPlanRepresentation)
)

// issue-routing-tag: disaster_recovery/default
func TestDisasterRecoveryDrPlanExecutionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDisasterRecoveryDrPlanExecutionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_disaster_recovery_dr_plan_execution.test_dr_plan_execution"
	datasourceName := "data.oci_disaster_recovery_dr_plan_executions.test_dr_plan_executions"
	singularDatasourceName := "data.oci_disaster_recovery_dr_plan_execution.test_dr_plan_execution"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DisasterRecoveryDrPlanExecutionResourceDependencies+
		DrProtectionGroupConfig+DrPlanConfig+
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan_execution", "test_dr_plan_execution", acctest.Optional, acctest.Create, DisasterRecoveryDrPlanExecutionRepresentation), "disasterrecovery", "drPlanExecution", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create Dependencies
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceDependencies +
				DrProtectionGroupConfig,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceDependencies +
				DrProtectionGroupConfig + DrPlanConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan_execution", "test_dr_plan_execution", acctest.Required, acctest.Create, DisasterRecoveryDrPlanExecutionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "execution_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.0.plan_execution_type", "SWITCHOVER_PRECHECK"),
				resource.TestCheckResourceAttrSet(resourceName, "plan_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceDependencies +
				DrProtectionGroupConfig + DrPlanConfig,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceDependencies +
				DrProtectionGroupConfig + DrPlanConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan_execution", "test_dr_plan_execution", acctest.Optional, acctest.Create, DisasterRecoveryDrPlanExecutionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Precheck Switchover from PHX to IAD"),
				resource.TestCheckResourceAttrSet(resourceName, "dr_protection_group_id"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.0.are_prechecks_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.0.are_warnings_ignored", "false"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.0.plan_execution_type", "SWITCHOVER_PRECHECK"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_executions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_dr_protection_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_region"),
				resource.TestCheckResourceAttrSet(resourceName, "plan_execution_type"),
				resource.TestCheckResourceAttrSet(resourceName, "plan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceDependencies +
				DrProtectionGroupConfig + DrPlanConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan_execution", "test_dr_plan_execution", acctest.Optional, acctest.Update, DisasterRecoveryDrPlanExecutionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "dr_protection_group_id"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.0.are_prechecks_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.0.are_warnings_ignored", "false"),
				resource.TestCheckResourceAttr(resourceName, "execution_options.0.plan_execution_type", "SWITCHOVER_PRECHECK"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_executions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_dr_protection_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_region"),
				resource.TestCheckResourceAttrSet(resourceName, "plan_execution_type"),
				resource.TestCheckResourceAttrSet(resourceName, "plan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceDependencies +
				DrProtectionGroupConfig + DrPlanConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_disaster_recovery_dr_plan_executions", "test_dr_plan_executions", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(DisasterRecoveryDisasterRecoveryDrPlanExecutionDataSourceRepresentation, []string{"state"})) +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan_execution", "test_dr_plan_execution", acctest.Optional, acctest.Update, DisasterRecoveryDrPlanExecutionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "dr_plan_execution_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dr_protection_group_id"),
				//resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

				resource.TestCheckResourceAttr(datasourceName, "dr_plan_execution_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dr_plan_execution_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceConfig +
				DrProtectionGroupConfig + DrPlanConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_disaster_recovery_dr_plan_execution", "test_dr_plan_execution", acctest.Required, acctest.Create, DisasterRecoveryDisasterRecoveryDrPlanExecutionSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dr_plan_execution_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_duration_in_sec"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_options.0.are_prechecks_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_options.0.are_warnings_ignored", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_options.0.plan_execution_type", "SWITCHOVER_PRECHECK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "group_executions.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "life_cycle_details"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_location.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_dr_protection_group_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_region"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_execution_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionRequiredOnlyResource +
				DrProtectionGroupConfig + DrPlanConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
		// delete dr plan and dr plan execution
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceDependencies +
				DrProtectionGroupConfig,
		},
		// Disassociate DrProtectionGroup
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanExecutionResourceDependencies +
				DrProtectionGroupWithDisassociateTriggerConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(2 * time.Minute)
					return
				},
			),
		},
	})
}

func testAccCheckDisasterRecoveryDrPlanExecutionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DisasterRecoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_disaster_recovery_dr_plan_execution" {
			noResourceFound = false
			request := oci_disaster_recovery.GetDrPlanExecutionRequest{}

			tmp := rs.Primary.ID
			request.DrPlanExecutionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "disaster_recovery")

			response, err := client.GetDrPlanExecution(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_disaster_recovery.DrPlanExecutionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DisasterRecoveryDrPlanExecution") {
		resource.AddTestSweepers("DisasterRecoveryDrPlanExecution", &resource.Sweeper{
			Name:         "DisasterRecoveryDrPlanExecution",
			Dependencies: acctest.DependencyGraph["drPlanExecution"],
			F:            sweepDisasterRecoveryDrPlanExecutionResource,
		})
	}
}

func sweepDisasterRecoveryDrPlanExecutionResource(compartment string) error {
	disasterRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DisasterRecoveryClient()
	drPlanExecutionIds, err := getDisasterRecoveryDrPlanExecutionIds(compartment)
	if err != nil {
		return err
	}
	for _, drPlanExecutionId := range drPlanExecutionIds {
		if ok := acctest.SweeperDefaultResourceId[drPlanExecutionId]; !ok {
			deleteDrPlanExecutionRequest := oci_disaster_recovery.DeleteDrPlanExecutionRequest{}

			deleteDrPlanExecutionRequest.DrPlanExecutionId = &drPlanExecutionId

			deleteDrPlanExecutionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "disaster_recovery")
			_, error := disasterRecoveryClient.DeleteDrPlanExecution(context.Background(), deleteDrPlanExecutionRequest)
			if error != nil {
				fmt.Printf("Error deleting DrPlanExecution %s %s, It is possible that the resource is already deleted. Please verify manually \n", drPlanExecutionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &drPlanExecutionId, DisasterRecoveryDrPlanExecutionSweepWaitCondition, time.Duration(3*time.Minute),
				DisasterRecoveryDrPlanExecutionSweepResponseFetchOperation, "disaster_recovery", true)
		}
	}
	return nil
}

func getDisasterRecoveryDrPlanExecutionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DrPlanExecutionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	disasterRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DisasterRecoveryClient()

	listDrPlanExecutionsRequest := oci_disaster_recovery.ListDrPlanExecutionsRequest{}
	//listDrPlanExecutionsRequest.CompartmentId = &compartmentId

	drProtectionGroupIds, error := getDisasterRecoveryDrProtectionGroupIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting drProtectionGroupId required for DrPlanExecution resource requests \n")
	}
	for _, drProtectionGroupId := range drProtectionGroupIds {
		listDrPlanExecutionsRequest.DrProtectionGroupId = &drProtectionGroupId

		listDrPlanExecutionsRequest.LifecycleState = oci_disaster_recovery.ListDrPlanExecutionsLifecycleStateSucceeded
		listDrPlanExecutionsResponse, err := disasterRecoveryClient.ListDrPlanExecutions(context.Background(), listDrPlanExecutionsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DrPlanExecution list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, drPlanExecution := range listDrPlanExecutionsResponse.Items {
			id := *drPlanExecution.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DrPlanExecutionId", id)
		}

	}
	return resourceIds, nil
}

func DisasterRecoveryDrPlanExecutionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drPlanExecutionResponse, ok := response.Response.(oci_disaster_recovery.GetDrPlanExecutionResponse); ok {
		return drPlanExecutionResponse.LifecycleState != oci_disaster_recovery.DrPlanExecutionLifecycleStateDeleted
	}
	return false
}

func DisasterRecoveryDrPlanExecutionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DisasterRecoveryClient().GetDrPlanExecution(context.Background(), oci_disaster_recovery.GetDrPlanExecutionRequest{
		DrPlanExecutionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
