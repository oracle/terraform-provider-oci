// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v58/databasemigration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	JobRequiredOnlyResource = JobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Required, acctest.Create, jobRepresentation)

	JobResourceConfig = //JobResourceDependencies +
	acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Optional, acctest.Update, jobRepresentation2)

	jobSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_job.test_job.id}`},
	}

	jobDataSourceRepresentation = map[string]interface{}{
		"migration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_migration.test_migration.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `TF_displayName`, Update: `TF_displayName2`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `Succeeded`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: jobDataSourceFilterRepresentation}}
	jobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `TF_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_job.test_job.id}`}},
	}

	jobRepresentation = map[string]interface{}{
		"job_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_job.test_job.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `TF_displayName`, Update: `TF_displayName2`},
	}

	jobRepresentation2 = map[string]interface{}{
		"job_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_job.test_job.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `TF_displayName`, Update: `TF_displayName2`},
	}

	JobResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", acctest.Required, acctest.Create, deploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", acctest.Required, acctest.Create, gatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, connectionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Required, acctest.Create, jobRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, migrationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, applicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, functionRepresentation) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobResource_basic(t *testing.T) {
	t.Skip("Skip this job creation is an independent operation after validating the migration")
	httpreplay.SetScenario("TestDatabaseMigrationJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_migration_job.test_job"
	datasourceName := "data.oci_database_migration_jobs.test_jobs"
	singularDatasourceName := "data.oci_database_migration_job.test_job"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Optional, acctest.Create, jobRepresentation), "databasemigration", "job", t)

	acctest.ResourceTest(t, testAccCheckDatabaseMigrationJobDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Required, acctest.Create, jobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "job_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "job_id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Optional, acctest.Create, jobRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "job_id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "job_id")
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
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Optional, acctest.Update, jobRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "job_id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "job_id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_jobs", "test_jobs", acctest.Optional, acctest.Update, jobDataSourceRepresentation) +
				compartmentIdVariableStr + //JobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Optional, acctest.Update, jobRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TF_displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "Succeeded"),
				resource.TestCheckResourceAttr(datasourceName, "job_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "job_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Required, acctest.Create, jobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + JobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + JobResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       false,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
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
					string(oci_database_migration.JobLifecycleStatesTerminated): true,
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
	jobIds, err := getJobIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &jobId, jobSweepWaitCondition, time.Duration(3*time.Minute),
				jobSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "JobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()

	listJobsRequest := oci_database_migration.ListJobsRequest{}
	migrationIds, error := getMigrationIds(compartment)
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

func jobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if jobResponse, ok := response.Response.(oci_database_migration.GetJobResponse); ok {
		return jobResponse.LifecycleState != oci_database_migration.JobLifecycleStatesTerminated
	}
	return false
}

func jobSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseMigrationClient().GetJob(context.Background(), oci_database_migration.GetJobRequest{
		JobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
