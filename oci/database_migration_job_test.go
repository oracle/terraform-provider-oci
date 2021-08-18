// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v46/databasemigration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	JobRequiredOnlyResource = JobResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Required, Create, jobRepresentation)

	JobResourceConfig = //JobResourceDependencies +
	generateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Optional, Update, jobRepresentation2)

	jobSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": Representation{repType: Required, create: `${oci_database_migration_job.test_job.id}`},
	}

	jobDataSourceRepresentation = map[string]interface{}{
		"migration_id": Representation{repType: Required, create: `${oci_database_migration_job.test_job.id}`},
		"display_name": Representation{repType: Optional, create: `TF_displayName`, update: `TF_displayName2`},
		"state":        Representation{repType: Optional, create: `Succeeded`},
		"filter":       RepresentationGroup{Required, jobDataSourceFilterRepresentation}}
	jobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `TF_id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_migration_job.test_job.id}`}},
	}

	jobRepresentation = map[string]interface{}{
		"job_id":       Representation{repType: Required, create: `${oci_database_migration_job.test_job.id}`},
		"display_name": Representation{repType: Optional, create: `TF_displayName`, update: `TF_displayName2`},
	}

	jobRepresentation2 = map[string]interface{}{
		"job_id":       Representation{repType: Required, create: `${oci_database_migration_job.test_job.id}`},
		"display_name": Representation{repType: Optional, create: `TF_displayName`, update: `TF_displayName2`},
	}

	JobResourceDependencies = generateResourceFromRepresentationMap("oci_apigateway_deployment", "test_deployment", Required, Create, deploymentRepresentation) +
		generateResourceFromRepresentationMap("oci_apigateway_gateway", "test_gateway", Required, Create, gatewayRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Required, Create, connectionRepresentation) +
		generateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Required, Create, jobRepresentation) +
		generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Required, Create, migrationRepresentation) +
		generateResourceFromRepresentationMap("oci_functions_application", "test_application", Required, Create, applicationRepresentation) +
		generateResourceFromRepresentationMap("oci_functions_function", "test_function", Required, Create, functionRepresentation) +
		DefinedTagsDependencies +
		KeyResourceDependencyConfig +
		generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultRepresentation) +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobResource_basic(t *testing.T) {
	t.Skip("Skip this job creation is an independent operation after validating the migration")
	httpreplay.SetScenario("TestDatabaseMigrationJobResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_migration_job.test_job"
	datasourceName := "data.oci_database_migration_jobs.test_jobs"
	singularDatasourceName := "data.oci_database_migration_job.test_job"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+
		generateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Optional, Create, jobRepresentation), "databasemigration", "job", t)

	ResourceTest(t, testAccCheckDatabaseMigrationJobDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr +
				generateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Required, Create, jobRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "job_id"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "job_id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr +
				generateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Optional, Create, jobRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "job_id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "job_id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				generateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Optional, Update, jobRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "job_id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				func(s *terraform.State) (err error) {
					resId2, err = fromInstanceState(s, resourceName, "job_id")
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
				generateDataSourceFromRepresentationMap("oci_database_migration_jobs", "test_jobs", Optional, Update, jobDataSourceRepresentation) +
				compartmentIdVariableStr + //JobResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Optional, Update, jobRepresentation2),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TF_displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "Succeeded"),
				resource.TestCheckResourceAttr(datasourceName, "job_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "job_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_migration_job", "test_job", Required, Create, jobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + JobResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	client := testAccProvider.Meta().(*OracleClients).databaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_job" {
			noResourceFound = false
			request := oci_database_migration.GetJobRequest{}

			tmp := rs.Primary.ID
			request.JobId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_migration")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseMigrationJob") {
		resource.AddTestSweepers("DatabaseMigrationJob", &resource.Sweeper{
			Name:         "DatabaseMigrationJob",
			Dependencies: DependencyGraph["job"],
			F:            sweepDatabaseMigrationJobResource,
		})
	}
}

func sweepDatabaseMigrationJobResource(compartment string) error {
	databaseMigrationClient := GetTestClients(&schema.ResourceData{}).databaseMigrationClient()
	jobIds, err := getJobIds(compartment)
	if err != nil {
		return err
	}
	for _, jobId := range jobIds {
		if ok := SweeperDefaultResourceId[jobId]; !ok {
			deleteJobRequest := oci_database_migration.DeleteJobRequest{}

			deleteJobRequest.JobId = &jobId

			deleteJobRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteJob(context.Background(), deleteJobRequest)
			if error != nil {
				fmt.Printf("Error deleting Job %s %s, It is possible that the resource is already deleted. Please verify manually \n", jobId, error)
				continue
			}
			waitTillCondition(testAccProvider, &jobId, jobSweepWaitCondition, time.Duration(3*time.Minute),
				jobSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getJobIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "JobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := GetTestClients(&schema.ResourceData{}).databaseMigrationClient()

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
			addResourceIdToSweeperResourceIdMap(compartmentId, "JobId", id)
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

func jobSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseMigrationClient().GetJob(context.Background(), oci_database_migration.GetJobRequest{
		JobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
