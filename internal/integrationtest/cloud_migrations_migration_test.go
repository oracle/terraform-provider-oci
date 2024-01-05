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
	CloudMigrationsMigrationRequiredOnlyResource = CloudMigrationsMigrationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Required, acctest.Create, CloudMigrationsMigrationRepresentation)

	CloudMigrationsMigrationResourceConfig = CloudMigrationsMigrationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Optional, acctest.Update, CloudMigrationsMigrationRepresentation)

	CloudMigrationsCloudMigrationsMigrationSingularDataSourceRepresentation = map[string]interface{}{
		"migration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_migration.test_migration.id}`},
	}

	CloudMigrationsCloudMigrationsMigrationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"migration_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_migration.test_migration.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudMigrationsMigrationDataSourceFilterRepresentation}}
	CloudMigrationsMigrationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_migrations_migration.test_migration.id}`}},
	}

	CloudMigrationsMigrationRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"is_completed":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"replication_schedule_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_replication_schedule.test_replication_schedule.id}`},
	}

	CloudMigrationsMigrationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Required, acctest.Create, CloudMigrationsReplicationScheduleRepresentation)
)

// issue-routing-tag: cloud_migrations/default
func TestCloudMigrationsMigrationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudMigrationsMigrationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_migrations_migration.test_migration"
	datasourceName := "data.oci_cloud_migrations_migrations.test_migrations"
	singularDatasourceName := "data.oci_cloud_migrations_migration.test_migration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudMigrationsMigrationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Optional, acctest.Create, CloudMigrationsMigrationRepresentation), "cloudmigrations", "migration", t)

	acctest.ResourceTest(t, testAccCheckCloudMigrationsMigrationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudMigrationsMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Required, acctest.Create, CloudMigrationsMigrationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudMigrationsMigrationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudMigrationsMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Optional, acctest.Create, CloudMigrationsMigrationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_completed", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_schedule_id"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudMigrationsMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudMigrationsMigrationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_completed", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_schedule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + CloudMigrationsMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Optional, acctest.Update, CloudMigrationsMigrationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_completed", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_schedule_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_migrations", "test_migrations", acctest.Optional, acctest.Update, CloudMigrationsCloudMigrationsMigrationDataSourceRepresentation) +
				compartmentIdVariableStr + CloudMigrationsMigrationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Optional, acctest.Update, CloudMigrationsMigrationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "migration_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "migration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "migration_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Required, acctest.Create, CloudMigrationsCloudMigrationsMigrationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudMigrationsMigrationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "migration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_completed", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudMigrationsMigrationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudMigrationsMigrationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_migrations_migration" {
			noResourceFound = false
			request := oci_cloud_migrations.GetMigrationRequest{}

			tmp := rs.Primary.ID
			request.MigrationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")

			response, err := client.GetMigration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_migrations.MigrationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudMigrationsMigration") {
		resource.AddTestSweepers("CloudMigrationsMigration", &resource.Sweeper{
			Name:         "CloudMigrationsMigration",
			Dependencies: acctest.DependencyGraph["migration"],
			F:            sweepCloudMigrationsMigrationResource,
		})
	}
}

func sweepCloudMigrationsMigrationResource(compartment string) error {
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()
	migrationIds, err := getCloudMigrationsMigrationIds(compartment)
	if err != nil {
		return err
	}
	for _, migrationId := range migrationIds {
		if ok := acctest.SweeperDefaultResourceId[migrationId]; !ok {
			deleteMigrationRequest := oci_cloud_migrations.DeleteMigrationRequest{}

			deleteMigrationRequest.MigrationId = &migrationId

			deleteMigrationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")
			_, error := migrationClient.DeleteMigration(context.Background(), deleteMigrationRequest)
			if error != nil {
				fmt.Printf("Error deleting Migration %s %s, It is possible that the resource is already deleted. Please verify manually \n", migrationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &migrationId, CloudMigrationsMigrationSweepWaitCondition, time.Duration(3*time.Minute),
				CloudMigrationsMigrationSweepResponseFetchOperation, "cloud_migrations", true)
		}
	}
	return nil
}

func getCloudMigrationsMigrationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MigrationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()

	listMigrationsRequest := oci_cloud_migrations.ListMigrationsRequest{}
	listMigrationsRequest.CompartmentId = &compartmentId
	listMigrationsRequest.LifecycleState = oci_cloud_migrations.MigrationLifecycleStateActive
	listMigrationsResponse, err := migrationClient.ListMigrations(context.Background(), listMigrationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Migration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, migration := range listMigrationsResponse.Items {
		id := *migration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MigrationId", id)
	}
	return resourceIds, nil
}

func CloudMigrationsMigrationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if migrationResponse, ok := response.Response.(oci_cloud_migrations.GetMigrationResponse); ok {
		return migrationResponse.LifecycleState != oci_cloud_migrations.MigrationLifecycleStateDeleted
	}
	return false
}

func CloudMigrationsMigrationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MigrationClient().GetMigration(context.Background(), oci_cloud_migrations.GetMigrationRequest{
		MigrationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
