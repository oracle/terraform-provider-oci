// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_resource_scheduler "github.com/oracle/oci-go-sdk/v65/resourcescheduler"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var (
	autonomousDatabaseOcid                                        = utils.GetEnvSettingWithBlankDefault("autonomousDatabase_ocid")
	computeInstanceOcid                                           = utils.GetEnvSettingWithBlankDefault("computeInstance_ocid")
	ResourceSchedulerScheduleWithResourceOcidRequiredOnlyResource = ResourceSchedulerScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, ResourceSchedulerScheduleWithResourceOcidRepresentation)

	ResourceSchedulerScheduleSingularDataSourceRepresentation = map[string]interface{}{
		"schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_scheduler_schedule.test_schedule.id}`},
	}

	ResourceSchedulerScheduleDataSourceRepresentation = map[string]interface{}{
		// must include at least one of `compartmentId` and `schedule_id`
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"schedule_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_resource_scheduler_schedule.test_schedule.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `provider displayName1`, Update: `provider displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ResourceSchedulerScheduleDataSourceFilterRepresentation}}
	ResourceSchedulerScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_resource_scheduler_schedule.test_schedule.id}`}},
	}

	ignoreChangesDefinedTagsResourceSchedulerRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	ResourceSchedulerScheduleWithResourceOcidRepresentation = map[string]interface{}{
		// Required
		"action":             acctest.Representation{RepType: acctest.Required, Create: `START_RESOURCE`, Update: `START_RESOURCE`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"recurrence_details": acctest.Representation{RepType: acctest.Required, Create: `FREQ=DAILY;INTERVAL=1`, Update: `FREQ=DAILY;INTERVAL=2`},
		"recurrence_type":    acctest.Representation{RepType: acctest.Required, Create: `ICAL`, Update: `ICAL`},
		// Must include either `resources` or `resource_filters` when creating schedules
		"resources": acctest.RepresentationGroup{RepType: acctest.Required, Group: ResourceSchedulerScheduleResourcesRepresentation},
		// Optionals
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `provider description1`, Update: `provider description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `provider displayName1`, Update: `provider displayName2`},
		"time_ends":     acctest.Representation{RepType: acctest.Optional, Create: `2024-06-22T00:00:00Z`, Update: `2024-06-24T00:00:00Z`},
		"time_starts":   acctest.Representation{RepType: acctest.Optional, Create: `2024-06-16T00:00:00Z`, Update: `2024-06-18T00:00:00Z`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"state":         acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
		"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreChangesDefinedTagsResourceSchedulerRepresentation},
	}

	ResourceSchedulerScheduleResourcesRepresentation = map[string]interface{}{
		"id":       acctest.Representation{RepType: acctest.Required, Create: autonomousDatabaseOcid, Update: computeInstanceOcid},
		"metadata": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"metadata": "metadata"}, Update: map[string]string{"metadata2": "metadata2"}},
	}

	ResourceSchedulerScheduleResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: resource_scheduler/default
func TestResourceSchedulerScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceSchedulerScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_resource_scheduler_schedule.test_schedule"
	singularDatasourceName := "data.oci_resource_scheduler_schedule.test_schedule"
	datasourceName := "data.oci_resource_scheduler_schedules.test_schedules"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ResourceSchedulerScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Optional, acctest.Create, ResourceSchedulerScheduleWithResourceOcidRepresentation), "resourcescheduler", "schedule", t)

	acctest.ResourceTest(t, testAccCheckResourceSchedulerScheduleDestroy, []resource.TestStep{
		//verify Create with Required - resourceOCID
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, ResourceSchedulerScheduleWithResourceOcidRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.id", autonomousDatabaseOcid),
				resource.TestCheckResourceAttr(resourceName, "resources.0.metadata.%", "0"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//delete before next Create
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies,
		},

		// verify create with optionals - resourceOCID
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Optional, acctest.Create, ResourceSchedulerScheduleWithResourceOcidRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),
				resource.TestCheckResourceAttr(resourceName, "description", "provider description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provider displayName1"),
				resource.TestCheckResourceAttr(resourceName, "time_ends", "2024-06-22T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "time_starts", "2024-06-16T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.id", autonomousDatabaseOcid),
				resource.TestCheckResourceAttr(resourceName, "resources.0.metadata.%", "1"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Optional, acctest.Update, ResourceSchedulerScheduleWithResourceOcidRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=2"),
				resource.TestCheckResourceAttr(resourceName, "recurrence_type", "ICAL"),
				resource.TestCheckResourceAttr(resourceName, "description", "provider description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "provider displayName2"),
				resource.TestCheckResourceAttr(resourceName, "time_ends", "2024-06-24T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "time_starts", "2024-06-18T00:00:00Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "resources.0.id", computeInstanceOcid),
				resource.TestCheckResourceAttr(resourceName, "resources.0.metadata.%", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					fmt.Printf("xiaotong printing resId and resId2, %s, %s", resId, resId2)
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, ResourceSchedulerScheduleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResourceSchedulerScheduleWithResourceOcidRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action", "START_RESOURCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "recurrence_details", "FREQ=DAILY;INTERVAL=1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recurrence_type", "ICAL"),

				resource.TestCheckResourceAttr(singularDatasourceName, "resources.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources.0.id", autonomousDatabaseOcid),
			),
		},

		// verify datasources
		{
			Config: config + compartmentIdVariableStr + ResourceSchedulerScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_resource_scheduler_schedule", "test_schedule", acctest.Required, acctest.Create, ResourceSchedulerScheduleWithResourceOcidRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_resource_scheduler_schedules", "test_schedules", acctest.Required, acctest.Create, ResourceSchedulerScheduleDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "schedule_collection.#", "1"),
			),
		},

		// verify resource import
		{
			Config:                  config + ResourceSchedulerScheduleWithResourceOcidRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"time_next_run"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckResourceSchedulerScheduleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ScheduleClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_resource_scheduler_schedule" {
			noResourceFound = false
			request := oci_resource_scheduler.GetScheduleRequest{}

			tmp := rs.Primary.ID
			request.ScheduleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resource_scheduler")

			response, err := client.GetSchedule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_resource_scheduler.ScheduleLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ResourceSchedulerSchedule") {
		resource.AddTestSweepers("ResourceSchedulerSchedule", &resource.Sweeper{
			Name:         "ResourceSchedulerSchedule",
			Dependencies: acctest.DependencyGraph["schedule"],
			F:            sweepResourceSchedulerScheduleResource,
		})
	}
}

func sweepResourceSchedulerScheduleResource(compartment string) error {
	scheduleClient := acctest.GetTestClients(&schema.ResourceData{}).ScheduleClient()
	scheduleIds, err := getResourceSchedulerScheduleIds(compartment)
	if err != nil {
		return err
	}
	for _, scheduleId := range scheduleIds {
		if ok := acctest.SweeperDefaultResourceId[scheduleId]; !ok {
			deleteScheduleRequest := oci_resource_scheduler.DeleteScheduleRequest{}

			deleteScheduleRequest.ScheduleId = &scheduleId

			deleteScheduleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "resource_scheduler")
			_, error := scheduleClient.DeleteSchedule(context.Background(), deleteScheduleRequest)
			if error != nil {
				fmt.Printf("Error deleting Schedule %s %s, It is possible that the resource is already deleted. Please verify manually \n", scheduleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &scheduleId, ResourceSchedulerScheduleSweepWaitCondition, time.Duration(3*time.Minute),
				ResourceSchedulerScheduleSweepResponseFetchOperation, "resource_scheduler", true)
		}
	}
	return nil
}

func getResourceSchedulerScheduleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ScheduleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	scheduleClient := acctest.GetTestClients(&schema.ResourceData{}).ScheduleClient()

	listSchedulesRequest := oci_resource_scheduler.ListSchedulesRequest{}
	listSchedulesRequest.CompartmentId = &compartmentId
	listSchedulesRequest.LifecycleState = oci_resource_scheduler.ScheduleLifecycleStateActive
	listSchedulesResponse, err := scheduleClient.ListSchedules(context.Background(), listSchedulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Schedule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, schedule := range listSchedulesResponse.Items {
		id := *schedule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ScheduleId", id)
	}
	return resourceIds, nil
}

func ResourceSchedulerScheduleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if scheduleResponse, ok := response.Response.(oci_resource_scheduler.GetScheduleResponse); ok {
		return scheduleResponse.LifecycleState != oci_resource_scheduler.ScheduleLifecycleStateDeleted
	}
	return false
}

func ResourceSchedulerScheduleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ScheduleClient().GetSchedule(context.Background(), oci_resource_scheduler.GetScheduleRequest{
		ScheduleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
