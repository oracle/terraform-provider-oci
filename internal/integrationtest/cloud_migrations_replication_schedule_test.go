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
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	execution_recurrences_create = "FREQ=DAILY;BYHOUR=5"
	execution_recurrences_update = "FREQ=DAILY;BYHOUR=6"

	CloudMigrationsReplicationScheduleRequiredOnlyResource = CloudMigrationsReplicationScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Required, acctest.Create, CloudMigrationsReplicationScheduleRepresentation)

	CloudMigrationsReplicationScheduleResourceConfig = CloudMigrationsReplicationScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Optional, acctest.Update, CloudMigrationsReplicationScheduleRepresentation)

	CloudMigrationsCloudMigrationsReplicationScheduleSingularDataSourceRepresentation = map[string]interface{}{
		"replication_schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_replication_schedule.test_replication_schedule.id}`},
	}

	CloudMigrationsCloudMigrationsReplicationScheduleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"replication_schedule_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_replication_schedule.test_replication_schedule.id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudMigrationsReplicationScheduleDataSourceFilterRepresentation}}
	CloudMigrationsReplicationScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_migrations_replication_schedule.test_replication_schedule.id}`}},
	}

	CloudMigrationsReplicationScheduleRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"execution_recurrences": acctest.Representation{RepType: acctest.Required, Create: execution_recurrences_create, Update: execution_recurrences_update},
	}

	CloudMigrationsReplicationScheduleResourceDependencies = ""
)

// issue-routing-tag: cloud_migrations/default
func TestCloudMigrationsReplicationScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudMigrationsReplicationScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_migrations_replication_schedule.test_replication_schedule"
	datasourceName := "data.oci_cloud_migrations_replication_schedules.test_replication_schedules"
	singularDatasourceName := "data.oci_cloud_migrations_replication_schedule.test_replication_schedule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudMigrationsReplicationScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Optional, acctest.Create, CloudMigrationsReplicationScheduleRepresentation), "cloudmigrations", "replicationSchedule", t)

	acctest.ResourceTest(t, testAccCheckCloudMigrationsReplicationScheduleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudMigrationsReplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Required, acctest.Create, CloudMigrationsReplicationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "execution_recurrences", execution_recurrences_create),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudMigrationsReplicationScheduleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudMigrationsReplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Optional, acctest.Create, CloudMigrationsReplicationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "execution_recurrences", execution_recurrences_create),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudMigrationsReplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudMigrationsReplicationScheduleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "execution_recurrences", execution_recurrences_create),
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
			Config: config + compartmentIdVariableStr + CloudMigrationsReplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Optional, acctest.Update, CloudMigrationsReplicationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "execution_recurrences", execution_recurrences_update),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_replication_schedules", "test_replication_schedules", acctest.Optional, acctest.Update, CloudMigrationsCloudMigrationsReplicationScheduleDataSourceRepresentation) +
				compartmentIdVariableStr + CloudMigrationsReplicationScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Optional, acctest.Update, CloudMigrationsReplicationScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "replication_schedule_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "replication_schedule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "replication_schedule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Required, acctest.Create, CloudMigrationsCloudMigrationsReplicationScheduleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudMigrationsReplicationScheduleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_schedule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "execution_recurrences", execution_recurrences_update),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudMigrationsReplicationScheduleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudMigrationsReplicationScheduleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_migrations_replication_schedule" {
			noResourceFound = false
			request := oci_cloud_migrations.GetReplicationScheduleRequest{}

			tmp := rs.Primary.ID
			request.ReplicationScheduleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")

			response, err := client.GetReplicationSchedule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_migrations.ReplicationScheduleLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudMigrationsReplicationSchedule") {
		resource.AddTestSweepers("CloudMigrationsReplicationSchedule", &resource.Sweeper{
			Name:         "CloudMigrationsReplicationSchedule",
			Dependencies: acctest.DependencyGraph["replicationSchedule"],
			F:            sweepCloudMigrationsReplicationScheduleResource,
		})
	}
}

func sweepCloudMigrationsReplicationScheduleResource(compartment string) error {
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()
	replicationScheduleIds, err := getCloudMigrationsReplicationScheduleIds(compartment)
	if err != nil {
		return err
	}
	for _, replicationScheduleId := range replicationScheduleIds {
		if ok := acctest.SweeperDefaultResourceId[replicationScheduleId]; !ok {
			deleteReplicationScheduleRequest := oci_cloud_migrations.DeleteReplicationScheduleRequest{}

			deleteReplicationScheduleRequest.ReplicationScheduleId = &replicationScheduleId

			deleteReplicationScheduleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")
			_, error := migrationClient.DeleteReplicationSchedule(context.Background(), deleteReplicationScheduleRequest)
			if error != nil {
				fmt.Printf("Error deleting ReplicationSchedule %s %s, It is possible that the resource is already deleted. Please verify manually \n", replicationScheduleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &replicationScheduleId, CloudMigrationsReplicationScheduleSweepWaitCondition, time.Duration(3*time.Minute),
				CloudMigrationsReplicationScheduleSweepResponseFetchOperation, "cloud_migrations", true)
		}
	}
	return nil
}

func getCloudMigrationsReplicationScheduleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ReplicationScheduleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()

	listReplicationSchedulesRequest := oci_cloud_migrations.ListReplicationSchedulesRequest{}
	listReplicationSchedulesRequest.CompartmentId = &compartmentId
	listReplicationSchedulesRequest.LifecycleState = oci_cloud_migrations.ReplicationScheduleLifecycleStateActive
	listReplicationSchedulesResponse, err := migrationClient.ListReplicationSchedules(context.Background(), listReplicationSchedulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ReplicationSchedule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, replicationSchedule := range listReplicationSchedulesResponse.Items {
		id := *replicationSchedule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ReplicationScheduleId", id)
	}
	return resourceIds, nil
}

func CloudMigrationsReplicationScheduleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if replicationScheduleResponse, ok := response.Response.(oci_cloud_migrations.GetReplicationScheduleResponse); ok {
		return replicationScheduleResponse.LifecycleState != oci_cloud_migrations.ReplicationScheduleLifecycleStateDeleted
	}
	return false
}

func CloudMigrationsReplicationScheduleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MigrationClient().GetReplicationSchedule(context.Background(), oci_cloud_migrations.GetReplicationScheduleRequest{
		ReplicationScheduleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
