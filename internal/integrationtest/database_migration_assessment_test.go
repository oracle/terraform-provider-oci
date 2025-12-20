// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseMigrationAssessmentRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Required, acctest.Create, DatabaseMigrationAssessmentRepresentation)

	DatabaseMigrationAssessmentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Optional, acctest.Update, DatabaseMigrationAssessmentRepresentation)

	DatabaseMigrationAssessmentSingularDataSourceRepresentation = map[string]interface{}{
		"assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_assessment.test_assessment.id}`},
	}

	DatabaseMigrationAssessmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationAssessmentDataSourceFilterRepresentation}}
	DatabaseMigrationAssessmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_assessment.test_assessment.id}`}},
	}

	DatabaseMigrationAssessmentRepresentation = map[string]interface{}{
		"acceptable_downtime":              acctest.Representation{RepType: acctest.Required, Create: `LESS_THAN_10_MINUTES`, Update: `LESS_THAN_1_HOUR`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_combination":             acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"database_data_size":               acctest.Representation{RepType: acctest.Required, Create: `LESS_THAN_1GB`, Update: `GB_1_10`},
		"ddl_expectation":                  acctest.Representation{RepType: acctest.Required, Create: `DDL_EXPECTED`, Update: `DDL_NOT_EXPECTED`},
		"network_speed_megabit_per_second": acctest.Representation{RepType: acctest.Required, Create: `MBPS_10`, Update: `MBPS_100`},
		"source_database_connection":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationAssessmentSourceDatabaseConnectionRepresentation},
		"target_database_connection":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationAssessmentTargetDatabaseConnectionRepresentation},
		"creation_type":                    acctest.Representation{RepType: acctest.Required, Create: `CREATE_ONLY`},
		"description":                      acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":                     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"exclude_objects":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationAssessmentExcludeObjectsRepresentation},
	}
	DatabaseMigrationAssessmentSourceDatabaseConnectionRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_oracle_id}`},
	}
	DatabaseMigrationAssessmentTargetDatabaseConnectionRepresentation = map[string]interface{}{
		"connection_type":     acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"database_version":    acctest.Representation{RepType: acctest.Optional, Create: `databaseVersion`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${var.target_connection_oracle_id}`},
		"technology_sub_type": acctest.Representation{RepType: acctest.Optional, Create: `technologySubType`},
		"technology_type":     acctest.Representation{RepType: acctest.Optional, Create: `OCI_AUTONOMOUS_DATABASE`},
	}
	DatabaseMigrationAssessmentExcludeObjectsRepresentation = map[string]interface{}{
		"object": acctest.Representation{RepType: acctest.Required, Create: `.*`},
		"is_omit_excluded_table_from_replication": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"owner":  acctest.Representation{RepType: acctest.Optional, Create: `owner`},
		"schema": acctest.Representation{RepType: acctest.Optional, Create: `schema`},
		"type":   acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
	}
	DatabaseMigrationAssessmentIncludeObjectsRepresentation = map[string]interface{}{
		"object": acctest.Representation{RepType: acctest.Required, Create: `.*`},
		"is_omit_excluded_table_from_replication": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"owner":  acctest.Representation{RepType: acctest.Optional, Create: `owner`},
		"schema": acctest.Representation{RepType: acctest.Optional, Create: `schema`},
		"type":   acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
	}
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAssessmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationAssessmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	sourceConnectionOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_oracle_id")
	sourceConnectionOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_oracle_id\" { default = \"%s\" }\n", sourceConnectionOracleId)

	targetConnectionOracleId := utils.GetEnvSettingWithBlankDefault("target_connection_oracle_id")
	targetConnectionOracleIdVariableStr := fmt.Sprintf("variable \"target_connection_oracle_id\" { default = \"%s\" }\n", targetConnectionOracleId)

	resourceName := "oci_database_migration_assessment.test_assessment"
	datasourceName := "data.oci_database_migration_assessments.test_assessments"
	singularDatasourceName := "data.oci_database_migration_assessment.test_assessment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+sourceConnectionOracleIdVariableStr+targetConnectionOracleIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Optional, acctest.Create, DatabaseMigrationAssessmentRepresentation), "databasemigration", "assessment", t)

	acctest.ResourceTest(t, testAccCheckDatabaseMigrationAssessmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Required, acctest.Create, DatabaseMigrationAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "acceptable_downtime", "LESS_THAN_10_MINUTES"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "database_data_size", "LESS_THAN_1GB"),
				resource.TestCheckResourceAttr(resourceName, "ddl_expectation", "DDL_EXPECTED"),
				resource.TestCheckResourceAttr(resourceName, "network_speed_megabit_per_second", "MBPS_10"),
				resource.TestCheckResourceAttr(resourceName, "source_database_connection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection.0.id"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
			Config: config + compartmentIdVariableStr + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Optional, acctest.Create, DatabaseMigrationAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "acceptable_downtime", "LESS_THAN_10_MINUTES"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "creation_type", "CREATE_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "database_data_size", "LESS_THAN_1GB"),
				resource.TestCheckResourceAttr(resourceName, "ddl_expectation", "DDL_EXPECTED"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.is_omit_excluded_table_from_replication", "false"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.object", ".*"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.owner", "owner"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.schema", "schema"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.type", "ALL"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_speed_megabit_per_second", "MBPS_10"),
				resource.TestCheckResourceAttr(resourceName, "source_database_connection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.connection_type", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.database_version", "databaseVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection.0.id"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.technology_sub_type", "technologySubType"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.technology_type", "OCI_AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					time.Sleep(1 * time.Minute)
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseMigrationAssessmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "acceptable_downtime", "LESS_THAN_10_MINUTES"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "creation_type", "CREATE_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "database_data_size", "LESS_THAN_1GB"),
				resource.TestCheckResourceAttr(resourceName, "ddl_expectation", "DDL_EXPECTED"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.is_omit_excluded_table_from_replication", "false"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.object", ".*"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.owner", "owner"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.schema", "schema"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.type", "ALL"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_speed_megabit_per_second", "MBPS_10"),
				resource.TestCheckResourceAttr(resourceName, "source_database_connection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.connection_type", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.database_version", "databaseVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection.0.id"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.technology_sub_type", "technologySubType"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.technology_type", "OCI_AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Optional, acctest.Update, DatabaseMigrationAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "acceptable_downtime", "LESS_THAN_1_HOUR"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "creation_type", "CREATE_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "database_data_size", "GB_1_10"),
				resource.TestCheckResourceAttr(resourceName, "ddl_expectation", "DDL_NOT_EXPECTED"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.is_omit_excluded_table_from_replication", "false"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.object", ".*"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.owner", "owner"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.schema", "schema"),
				resource.TestCheckResourceAttr(resourceName, "exclude_objects.0.type", "ALL"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "network_speed_megabit_per_second", "MBPS_100"),
				resource.TestCheckResourceAttr(resourceName, "source_database_connection.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection.0.id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.connection_type", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.database_version", "databaseVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection.0.id"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.technology_sub_type", "technologySubType"),
				resource.TestCheckResourceAttr(resourceName, "target_database_connection.0.technology_type", "OCI_AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_assessments", "test_assessments", acctest.Optional, acctest.Update, DatabaseMigrationAssessmentDataSourceRepresentation) +
				compartmentIdVariableStr + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Optional, acctest.Update, DatabaseMigrationAssessmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "assessment_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Required, acctest.Create, DatabaseMigrationAssessmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseMigrationAssessmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "assessment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "acceptable_downtime", "LESS_THAN_1_HOUR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "creation_type", "CREATE_ONLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_data_size", "GB_1_10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ddl_expectation", "DDL_NOT_EXPECTED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_speed_megabit_per_second", "MBPS_100"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_database_connection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_database_connection.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_database_connection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_database_connection.0.connection_type", "ORACLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_database_connection.0.database_version", "databaseVersion"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_database_connection.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_database_connection.0.technology_sub_type", "technologySubType"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_database_connection.0.technology_type", "OCI_AUTONOMOUS_DATABASE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseMigrationAssessmentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"bulk_include_exclude_data",
				"exclude_objects",
				"include_objects",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseMigrationAssessmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_assessment" {
			noResourceFound = false
			request := oci_database_migration.GetAssessmentRequest{}

			tmp := rs.Primary.ID
			request.AssessmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")

			response, err := client.GetAssessment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_migration.AssessmentLifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("DatabaseMigrationAssessment") {
		resource.AddTestSweepers("DatabaseMigrationAssessment", &resource.Sweeper{
			Name:         "DatabaseMigrationAssessment",
			Dependencies: acctest.DependencyGraph["assessment"],
			F:            sweepDatabaseMigrationAssessmentResource,
		})
	}
}

func sweepDatabaseMigrationAssessmentResource(compartment string) error {
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()
	assessmentIds, err := getDatabaseMigrationAssessmentIds(compartment)
	if err != nil {
		return err
	}
	for _, assessmentId := range assessmentIds {
		if ok := acctest.SweeperDefaultResourceId[assessmentId]; !ok {
			deleteAssessmentRequest := oci_database_migration.DeleteAssessmentRequest{}

			deleteAssessmentRequest.AssessmentId = &assessmentId

			deleteAssessmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteAssessment(context.Background(), deleteAssessmentRequest)
			if error != nil {
				fmt.Printf("Error deleting Assessment %s %s, It is possible that the resource is already deleted. Please verify manually \n", assessmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &assessmentId, DatabaseMigrationAssessmentSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseMigrationAssessmentSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getDatabaseMigrationAssessmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AssessmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()

	listAssessmentsRequest := oci_database_migration.ListAssessmentsRequest{}
	listAssessmentsRequest.CompartmentId = &compartmentId
	listAssessmentsRequest.LifecycleState = oci_database_migration.ListAssessmentsLifecycleStateActive
	listAssessmentsResponse, err := databaseMigrationClient.ListAssessments(context.Background(), listAssessmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Assessment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, assessment := range listAssessmentsResponse.Items {
		id := *assessment.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AssessmentId", id)
	}
	return resourceIds, nil
}

func DatabaseMigrationAssessmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if assessmentResponse, ok := response.Response.(oci_database_migration.GetAssessmentResponse); ok {
		return assessmentResponse.GetLifecycleState() != oci_database_migration.AssessmentLifecycleStatesDeleted
	}
	return false
}

func DatabaseMigrationAssessmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseMigrationClient().GetAssessment(context.Background(), oci_database_migration.GetAssessmentRequest{
		AssessmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
