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
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DisasterRecoveryDrPlanRequiredOnlyResource = DisasterRecoveryDrPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Required, acctest.Create, DisasterRecoveryDrPlanRepresentation)

	DisasterRecoveryDrPlanResourceConfig = DisasterRecoveryDrPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Optional, acctest.Update, DisasterRecoveryDrPlanRepresentation)

	DisasterRecoveryDrPlanResourceConfigWithoutTriggers = DisasterRecoveryDrPlanResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Optional, acctest.Update, DisasterRecoveryDrPlanRepresentationWithoutTriggers)

	DisasterRecoveryDisasterRecoveryDrPlanSingularDataSourceRepresentation = map[string]interface{}{
		"dr_plan_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_dr_plan.test_dr_plan.id}`},
	}

	//Dr Plan Data source
	DisasterRecoveryDisasterRecoveryDrPlanDataSourceRepresentation = map[string]interface{}{
		"dr_protection_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_dr_protection_group.test_peer.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `Switchover from PHX to IAD`, Update: `displayName2`},
		"dr_plan_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_disaster_recovery_dr_plan.test_dr_plan.id}`},
		"dr_plan_type":           acctest.Representation{RepType: acctest.Optional, Create: `SWITCHOVER`},
		"lifecycle_sub_state":    acctest.Representation{RepType: acctest.Optional, Create: `NEEDS_REFRESH`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DisasterRecoveryDrPlanDataSourceFilterRepresentation}}

	DisasterRecoveryDrPlanDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_disaster_recovery_dr_plan.test_dr_plan.id}`}},
	}

	DisasterRecoveryDrPlanRepresentation = map[string]interface{}{
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `Switchover from PHX to IAD`, Update: `displayName2`},
		"dr_protection_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_dr_protection_group.test_peer.id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `SWITCHOVER`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//"source_plan_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_disaster_recovery_source_plan.test_source_plan.id}`},
		"refresh_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"verify_trigger":  acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DefinedTagsIgnoreRepresentation},
	}

	DisasterRecoveryDrPlanRepresentationWithoutTriggers = map[string]interface{}{
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `Switchover from PHX to IAD`, Update: `displayName2`},
		"dr_protection_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_disaster_recovery_dr_protection_group.test_peer.id}`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `SWITCHOVER`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//"source_plan_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_disaster_recovery_source_plan.test_source_plan.id}`},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DefinedTagsIgnoreRepresentation},
	}

	DisasterRecoveryDrPlanConfigWithRefreshTrigger = `
	resource "oci_disaster_recovery_dr_plan" "test_dr_plan" {
	  	#Required
	  	display_name   = "Switchover from PHX to IAD"
		dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_peer.id
	  	type = "SWITCHOVER"
		refresh_trigger = 1
	}
	`
	DisasterRecoveryDrPlanConfigWithVerifyTrigger = `
	resource "oci_disaster_recovery_dr_plan" "test_dr_plan" {
	  	#Required
	  	display_name   = "Switchover from PHX to IAD"
		dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_peer.id
	  	type = "SWITCHOVER"
		refresh_trigger = 1
		verify_trigger = 1
	}
	`

	DisasterRecoveryDrPlanResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_protection_group", "test_peer", acctest.Optional, acctest.Create, DisasterRecoveryPeerDrProtectionGroupRepresentation) +
		OKEClusterDependencyConfig +
		ObjectStorageBucketDependencyConfig +
		VolumeGroupDependencyConfig +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: disaster_recovery/default
func TestDisasterRecoveryDrPlanResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDisasterRecoveryDrPlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_disaster_recovery_dr_plan.test_dr_plan"
	datasourceName := "data.oci_disaster_recovery_dr_plans.test_dr_plans"
	singularDatasourceName := "data.oci_disaster_recovery_dr_plan.test_dr_plan"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DisasterRecoveryDrPlanResourceDependencies+
		DrProtectionGroupConfig+
		acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Optional, acctest.Create, DisasterRecoveryDrPlanRepresentation), "disasterrecovery", "drPlan", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create Dependencies
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				DrProtectionGroupConfig,
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				DrProtectionGroupConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Required, acctest.Create, DisasterRecoveryDrPlanRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "Switchover from PHX to IAD"),
				resource.TestCheckResourceAttrSet(resourceName, "dr_protection_group_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "SWITCHOVER"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// move DRPlan to needs refresh
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_disaster_recovery_dr_protection_group",
					"test_dr_protection_group",
					acctest.Required,
					acctest.Create,
					acctest.RepresentationCopyWithRemovedProperties(DisasterRecoveryDrProtectionGroupRepresentation, []string{"members"}),
				) +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_disaster_recovery_dr_plan",
					"test_dr_plan",
					acctest.Required,
					acctest.Create,
					DisasterRecoveryDrPlanRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "Switchover from PHX to IAD"),
				resource.TestCheckResourceAttrSet(resourceName, "dr_protection_group_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "SWITCHOVER"),
				func(s *terraform.State) error {
					client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DisasterRecoveryClient()
					drPlanID, _ := acctest.FromInstanceState(s, resourceName, "id")
					return waitForDrPlanState(*client, drPlanID, "NEEDS_ATTENTION", "NEEDS_REFRESH", 10, 10*time.Second)
				},
			),
		},

		// refresh DRPlan
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_disaster_recovery_dr_protection_group",
					"test_dr_protection_group",
					acctest.Required,
					acctest.Create,
					acctest.RepresentationCopyWithRemovedProperties(DisasterRecoveryDrProtectionGroupRepresentation, []string{"members"}),
				) +
				DisasterRecoveryDrPlanConfigWithRefreshTrigger,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "refresh_trigger", "1"),

				func(s *terraform.State) error {
					client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DisasterRecoveryClient()
					drPlanID, _ := acctest.FromInstanceState(s, resourceName, "id")
					return waitForDrPlanState(*client, drPlanID, "NEEDS_ATTENTION", "NEEDS_VERIFICATION", 10, 10*time.Second)
				},
			),
		},

		// verify DRPlan
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_disaster_recovery_dr_protection_group",
					"test_dr_protection_group",
					acctest.Required,
					acctest.Create,
					acctest.RepresentationCopyWithRemovedProperties(DisasterRecoveryDrProtectionGroupRepresentation, []string{"members"}),
				) +
				DisasterRecoveryDrPlanConfigWithVerifyTrigger,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "verify_trigger", "1"),

				func(s *terraform.State) error {
					client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DisasterRecoveryClient()
					drPlanID, _ := acctest.FromInstanceState(s, resourceName, "id")
					return waitForDrPlanState(*client, drPlanID, "ACTIVE", "", 10, 10*time.Second)
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				DrProtectionGroupConfig,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				DrProtectionGroupConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Optional, acctest.Create, DisasterRecoveryDrPlanRepresentationWithoutTriggers),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Switchover from PHX to IAD"),
				resource.TestCheckResourceAttrSet(resourceName, "dr_protection_group_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_dr_protection_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_region"),
				//resource.TestCheckResourceAttrSet(resourceName, "source_plan_id"),
				resource.TestCheckResourceAttr(resourceName, "plan_groups.#", "4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "SWITCHOVER"),

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
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				DrProtectionGroupConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Optional, acctest.Update, DisasterRecoveryDrPlanRepresentationWithoutTriggers),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "dr_protection_group_id"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_dr_protection_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_region"),
				//resource.TestCheckResourceAttrSet(resourceName, "source_plan_id"),
				resource.TestCheckResourceAttr(resourceName, "plan_groups.#", "4"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "SWITCHOVER"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_disaster_recovery_dr_plans", "test_dr_plans", acctest.Optional, acctest.Update, DisasterRecoveryDisasterRecoveryDrPlanDataSourceRepresentation) +
				compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
				DrProtectionGroupConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Optional, acctest.Update, DisasterRecoveryDrPlanRepresentationWithoutTriggers),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "dr_plan_id"),
				resource.TestCheckResourceAttr(datasourceName, "dr_plan_type", "SWITCHOVER"),
				resource.TestCheckResourceAttrSet(datasourceName, "dr_protection_group_id"),
				//resource.TestCheckResourceAttr(datasourceName, "lifecycle_sub_state", "NEEDS_REFRESH"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "dr_plan_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dr_plan_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_disaster_recovery_dr_plan", "test_dr_plan", acctest.Required, acctest.Create, DisasterRecoveryDisasterRecoveryDrPlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DisasterRecoveryDrPlanResourceConfigWithoutTriggers +
				DrProtectionGroupConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dr_plan_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Accounting"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_dr_protection_group_id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_region"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "plan_groups.#", "3"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "life_cycle_details"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_sub_state"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "type", "SWITCHOVER"),
			),
		},
		// verify resource import
		{
			Config:                  config + DisasterRecoveryDrPlanRequiredOnlyResource + DrProtectionGroupConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
		// Disassociate DrProtectionGroup
		{
			Config: config + compartmentIdVariableStr + DisasterRecoveryDrPlanResourceDependencies +
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

func testAccCheckDisasterRecoveryDrPlanDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DisasterRecoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_disaster_recovery_dr_plan" {
			noResourceFound = false
			request := oci_disaster_recovery.GetDrPlanRequest{}

			tmp := rs.Primary.ID
			request.DrPlanId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "disaster_recovery")

			response, err := client.GetDrPlan(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_disaster_recovery.DrPlanLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DisasterRecoveryDrPlan") {
		resource.AddTestSweepers("DisasterRecoveryDrPlan", &resource.Sweeper{
			Name:         "DisasterRecoveryDrPlan",
			Dependencies: acctest.DependencyGraph["drPlan"],
			F:            sweepDisasterRecoveryDrPlanResource,
		})
	}
}

func sweepDisasterRecoveryDrPlanResource(compartment string) error {
	disasterRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DisasterRecoveryClient()
	drPlanIds, err := getDisasterRecoveryDrPlanIds(compartment)
	if err != nil {
		return err
	}
	for _, drPlanId := range drPlanIds {
		if ok := acctest.SweeperDefaultResourceId[drPlanId]; !ok {
			deleteDrPlanRequest := oci_disaster_recovery.DeleteDrPlanRequest{}

			deleteDrPlanRequest.DrPlanId = &drPlanId

			deleteDrPlanRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "disaster_recovery")
			_, error := disasterRecoveryClient.DeleteDrPlan(context.Background(), deleteDrPlanRequest)
			if error != nil {
				fmt.Printf("Error deleting DrPlan %s %s, It is possible that the resource is already deleted. Please verify manually \n", drPlanId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &drPlanId, DisasterRecoveryDrPlanSweepWaitCondition, time.Duration(3*time.Minute),
				DisasterRecoveryDrPlanSweepResponseFetchOperation, "disaster_recovery", true)
		}
	}
	return nil
}

func getDisasterRecoveryDrPlanIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DrPlanId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	disasterRecoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DisasterRecoveryClient()

	listDrPlansRequest := oci_disaster_recovery.ListDrPlansRequest{}
	//listDrPlansRequest.CompartmentId = &compartmentId

	drProtectionGroupIds, error := getDisasterRecoveryDrProtectionGroupIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting drProtectionGroupId required for DrPlan resource requests \n")
	}
	for _, drProtectionGroupId := range drProtectionGroupIds {
		listDrPlansRequest.DrProtectionGroupId = &drProtectionGroupId

		listDrPlansRequest.LifecycleState = oci_disaster_recovery.ListDrPlansLifecycleStateActive
		listDrPlansResponse, err := disasterRecoveryClient.ListDrPlans(context.Background(), listDrPlansRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DrPlan list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, drPlan := range listDrPlansResponse.Items {
			id := *drPlan.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DrPlanId", id)
		}

	}
	return resourceIds, nil
}

func waitForDrPlanState(client oci_disaster_recovery.DisasterRecoveryClient, drPlanID, desiredState, desiredSubState string, maxRetries int, delayBetweenRetries time.Duration) error {
	// get DR Plan details to get the correct state it is in while running the acceptance test
	for i := 0; i < maxRetries; i++ {
		request := oci_disaster_recovery.GetDrPlanRequest{
			DrPlanId: common.String(drPlanID),
			RequestMetadata: common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(false, "disaster_recovery"),
			},
		}

		response, err := client.GetDrPlan(context.Background(), request)
		if err != nil {
			return fmt.Errorf("error fetching DR Plan state: %v", err)
		}

		currentState := string(response.LifecycleState)
		currentSubState := string(response.LifecycleSubState)

		if currentState == desiredState && currentSubState == desiredSubState {
			log.Printf("DR Plan reached desired state '%s' and sub-state '%s'", desiredState, desiredSubState)
			return nil
		}
		log.Printf("Attempt %d: Current state is '%s' and lifecycle_sub_state is '%s'. Retrying...\n", i+1, currentState, currentSubState)
		time.Sleep(delayBetweenRetries)
	}
	return fmt.Errorf("timeout waiting for DR Plan to reach state '%s' and sub-state '%s'", desiredState, desiredSubState)
}

func DisasterRecoveryDrPlanSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if drPlanResponse, ok := response.Response.(oci_disaster_recovery.GetDrPlanResponse); ok {
		return drPlanResponse.LifecycleState != oci_disaster_recovery.DrPlanLifecycleStateDeleted
	}
	return false
}

func DisasterRecoveryDrPlanSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DisasterRecoveryClient().GetDrPlan(context.Background(), oci_disaster_recovery.GetDrPlanRequest{
		DrPlanId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
