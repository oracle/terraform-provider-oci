// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseMigrationJobRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Required, acctest.Create, DatabaseMigrationJobRepresentation)

	DatabaseMigrationJobResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Optional, acctest.Update, DatabaseMigrationJobRepresentation)

	DatabaseMigrationJobSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": acctest.Representation{RepType: acctest.Required, Create: `${var.job_new_api_id}`},
	}

	DatabaseMigrationjobSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": acctest.Representation{RepType: acctest.Required, Create: `${var.job_new_api_id}`},
	}

	DatabaseMigrationJobDataSourceRepresentation = map[string]interface{}{
		"migration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_migration.test_migration.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationJobDataSourceFilterRepresentation}}
	DatabaseMigrationJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_job.test_job.id}`}},
	}

	DatabaseMigrationJobRepresentation = map[string]interface{}{
		"job_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.job_new_api_id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	jobId := utils.GetEnvSettingWithBlankDefault("job_new_api_id")
	jobIdVariableStr := fmt.Sprintf("variable \"job_new_api_id\" { default = \"%s\" }\n", jobId)

	singularDatasourceName := "data.oci_database_migration_job.test_job"

	acctest.ResourceTest(t, testAccCheckDatabaseMigrationJobDestroy, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + jobIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Required, acctest.Create, DatabaseMigrationjobSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "job-20240516171512"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "progress.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
	})
}

func testAccCheckDatabaseMigrationJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_job" {
			noResourceFound = false
			request := oci_database_migration.GetJobRequest{}

			tmp := rs.Primary.ID
			request.JobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")

			response, err := client.GetJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					// Because of the nature of this test, the Job is expected to finish in a Failed State
					string(oci_database_migration.JobLifecycleStatesFailed): true,
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
	if !acctest.InSweeperExcludeList("DatabaseMigrationJob") {
		resource.AddTestSweepers("DatabaseMigrationJob", &resource.Sweeper{
			Name:         "DatabaseMigrationJob",
			Dependencies: acctest.DependencyGraph["job"],
			F:            sweepDatabaseMigrationJobResource,
		})
	}
}

func sweepDatabaseMigrationJobResource(compartment string) error {
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()
	jobIds, err := getDatabaseMigrationJobIds(compartment)
	if err != nil {
		return err
	}
	for _, jobId := range jobIds {
		if ok := acctest.SweeperDefaultResourceId[jobId]; !ok {
			deleteJobRequest := oci_database_migration.DeleteJobRequest{}

			deleteJobRequest.JobId = &jobId

			deleteJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteJob(context.Background(), deleteJobRequest)
			if error != nil {
				fmt.Printf("Error deleting Job %s %s, It is possible that the resource is already deleted. Please verify manually \n", jobId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &jobId, DatabaseMigrationJobSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseMigrationJobSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getDatabaseMigrationJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "JobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()

	listJobsRequest := oci_database_migration.ListJobsRequest{}
	//listJobsRequest.CompartmentId = &compartmentId

	migrationIds, error := getDatabaseMigrationMigrationIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting migrationId required for Job resource requests \n")
	}
	for _, migrationId := range migrationIds {
		listJobsRequest.MigrationId = &migrationId

		listJobsRequest.LifecycleState = oci_database_migration.ListJobsLifecycleStateSucceeded
		listJobsResponse, err := databaseMigrationClient.ListJobs(context.Background(), listJobsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Job list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, job := range listJobsResponse.Items {
			id := *job.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "JobId", id)
		}

	}
	return resourceIds, nil
}

func DatabaseMigrationJobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if jobResponse, ok := response.Response.(oci_database_migration.GetJobResponse); ok {
		return jobResponse.LifecycleState != oci_database_migration.JobLifecycleStatesTerminated
	}
	return false
}

func DatabaseMigrationJobSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseMigrationClient().GetJob(context.Background(), oci_database_migration.GetJobRequest{
		JobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
