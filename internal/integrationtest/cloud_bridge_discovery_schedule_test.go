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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

//fake

var (
	executionRecurrence                              = "FREQ=DAILY;BYHOUR=6"
	CloudBridgeDiscoveryScheduleRequiredOnlyResource = CloudBridgeDiscoveryScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Required, acctest.Create, CloudBridgeDiscoveryScheduleRepresentation)

	CloudBridgeDiscoveryScheduleResourceConfig = CloudBridgeDiscoveryScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Optional, acctest.Update, CloudBridgeDiscoveryScheduleRepresentation)

	CloudBridgeCloudBridgeDiscoveryScheduleSingularDataSourceRepresentation = map[string]interface{}{
		"discovery_schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id}`},
	}

	CloudBridgeCloudBridgeDiscoveryScheduleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"discovery_schedule_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeDiscoveryScheduleDataSourceFilterRepresentation}}
	CloudBridgeDiscoveryScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id}`}},
	}

	CloudBridgeDiscoveryScheduleRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"execution_recurrences": acctest.Representation{RepType: acctest.Required, Create: executionRecurrence, Update: executionRecurrence},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}

	CloudBridgeDiscoveryScheduleResourceDependencies = ""
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeDiscoveryScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeDiscoveryScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_bridge_discovery_schedule.test_discovery_schedule"
	datasourceName := "data.oci_cloud_bridge_discovery_schedules.test_discovery_schedules"
	singularDatasourceName := "data.oci_cloud_bridge_discovery_schedule.test_discovery_schedule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudBridgeDiscoveryScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Optional, acctest.Create, CloudBridgeDiscoveryScheduleRepresentation), "cloudbridge", "discoverySchedule", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeDiscoveryScheduleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeDiscoveryScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Required, acctest.Create, CloudBridgeDiscoveryScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "execution_recurrences", executionRecurrence),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeDiscoveryScheduleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudBridgeDiscoveryScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Optional, acctest.Create, CloudBridgeDiscoveryScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "execution_recurrences", executionRecurrence),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudBridgeDiscoveryScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudBridgeDiscoveryScheduleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "execution_recurrences", executionRecurrence),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CloudBridgeDiscoveryScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Optional, acctest.Update, CloudBridgeDiscoveryScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "execution_recurrences", executionRecurrence),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_discovery_schedules", "test_discovery_schedules", acctest.Optional, acctest.Update, CloudBridgeCloudBridgeDiscoveryScheduleDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeDiscoveryScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Optional, acctest.Update, CloudBridgeDiscoveryScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_schedule_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "discovery_schedule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "discovery_schedule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Required, acctest.Create, CloudBridgeCloudBridgeDiscoveryScheduleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeDiscoveryScheduleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovery_schedule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_recurrences", executionRecurrence),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudBridgeDiscoveryScheduleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudBridgeDiscoveryScheduleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DiscoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_bridge_discovery_schedule" {
			noResourceFound = false
			request := oci_cloud_bridge.GetDiscoveryScheduleRequest{}

			tmp := rs.Primary.ID
			request.DiscoveryScheduleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")

			response, err := client.GetDiscoverySchedule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_bridge.DiscoveryScheduleLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudBridgeDiscoverySchedule") {
		resource.AddTestSweepers("CloudBridgeDiscoverySchedule", &resource.Sweeper{
			Name:         "CloudBridgeDiscoverySchedule",
			Dependencies: acctest.DependencyGraph["discoverySchedule"],
			F:            sweepCloudBridgeDiscoveryScheduleResource,
		})
	}
}

func sweepCloudBridgeDiscoveryScheduleResource(compartment string) error {
	discoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DiscoveryClient()
	discoveryScheduleIds, err := getCloudBridgeDiscoveryScheduleIds(compartment)
	if err != nil {
		return err
	}
	for _, discoveryScheduleId := range discoveryScheduleIds {
		if ok := acctest.SweeperDefaultResourceId[discoveryScheduleId]; !ok {
			deleteDiscoveryScheduleRequest := oci_cloud_bridge.DeleteDiscoveryScheduleRequest{}

			deleteDiscoveryScheduleRequest.DiscoveryScheduleId = &discoveryScheduleId

			deleteDiscoveryScheduleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")
			_, error := discoveryClient.DeleteDiscoverySchedule(context.Background(), deleteDiscoveryScheduleRequest)
			if error != nil {
				fmt.Printf("Error deleting DiscoverySchedule %s %s, It is possible that the resource is already deleted. Please verify manually \n", discoveryScheduleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &discoveryScheduleId, CloudBridgeDiscoveryScheduleSweepWaitCondition, time.Duration(3*time.Minute),
				CloudBridgeDiscoveryScheduleSweepResponseFetchOperation, "cloud_bridge", true)
		}
	}
	return nil
}

func getCloudBridgeDiscoveryScheduleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DiscoveryScheduleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	discoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DiscoveryClient()

	listDiscoverySchedulesRequest := oci_cloud_bridge.ListDiscoverySchedulesRequest{}
	listDiscoverySchedulesRequest.CompartmentId = &compartmentId
	listDiscoverySchedulesRequest.LifecycleState = oci_cloud_bridge.ListDiscoverySchedulesLifecycleStateActive
	listDiscoverySchedulesResponse, err := discoveryClient.ListDiscoverySchedules(context.Background(), listDiscoverySchedulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DiscoverySchedule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, discoverySchedule := range listDiscoverySchedulesResponse.Items {
		id := *discoverySchedule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DiscoveryScheduleId", id)
	}
	return resourceIds, nil
}

func CloudBridgeDiscoveryScheduleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if discoveryScheduleResponse, ok := response.Response.(oci_cloud_bridge.GetDiscoveryScheduleResponse); ok {
		return discoveryScheduleResponse.LifecycleState != oci_cloud_bridge.DiscoveryScheduleLifecycleStateDeleted
	}
	return false
}

func CloudBridgeDiscoveryScheduleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DiscoveryClient().GetDiscoverySchedule(context.Background(), oci_cloud_bridge.GetDiscoveryScheduleRequest{
		DiscoveryScheduleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
