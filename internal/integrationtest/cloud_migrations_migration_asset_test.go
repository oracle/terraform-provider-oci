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

// fake
var (
	CloudMigrationsBucketName = "test_bucket"
	CloudMigrationsAssetId    = `${var.inventoryAssetId}`
	CloudMigrationsAssetAD    = "oQNt:US-ASHBURN-AD-1"

	CloudMigrationsMigrationAssetRequiredOnlyResource = CloudMigrationsMigrationAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_asset", "test_migration_asset", acctest.Required, acctest.Create, CloudMigrationsMigrationAssetRepresentation)

	CloudMigrationsMigrationAssetResourceConfig = CloudMigrationsMigrationAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_asset", "test_migration_asset", acctest.Optional, acctest.Update, CloudMigrationsMigrationAssetRepresentation)

	CloudMigrationsCloudMigrationsMigrationAssetSingularDataSourceRepresentation = map[string]interface{}{
		"migration_asset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_migration_asset.test_migration_asset.id}`},
	}

	CloudMigrationsCloudMigrationsMigrationAssetDataSourceRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"migration_asset_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_migration_asset.test_migration_asset.id}`},
		"migration_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_migration.test_migration.id}`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudMigrationsMigrationAssetDataSourceFilterRepresentation}}
	CloudMigrationsMigrationAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_migrations_migration_asset.test_migration_asset.id}`}},
	}

	CloudMigrationsMigrationAssetRepresentation = map[string]interface{}{
		"availability_domain":        acctest.Representation{RepType: acctest.Required, Create: CloudMigrationsAssetAD},
		"inventory_asset_id":         acctest.Representation{RepType: acctest.Required, Create: CloudMigrationsAssetId},
		"migration_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_migrations_migration.test_migration.id}`},
		"replication_compartment_id": acctest.Representation{RepType: acctest.Required, Create: compartmentId},
		"snap_shot_bucket_name":      acctest.Representation{RepType: acctest.Required, Create: CloudMigrationsBucketName},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"replication_schedule_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_migrations_replication_schedule.test_replication_schedule.id}`},
	}

	CloudMigrationsMigrationAssetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration", "test_migration", acctest.Required, acctest.Create, CloudMigrationsMigrationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_replication_schedule", "test_replication_schedule", acctest.Required, acctest.Create, CloudMigrationsReplicationScheduleRepresentation)
)

// issue-routing-tag: cloud_migrations/default
func TestCloudMigrationsMigrationAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudMigrationsMigrationAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	inventoryAssetId := utils.GetEnvSettingWithBlankDefault("inventoryAssetId")
	inventoryAssetIdVariableStr := fmt.Sprintf("variable \"inventoryAssetId\" { default = \"%s\" }\n", inventoryAssetId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	variableStr := inventoryAssetIdVariableStr + compartmentIdVariableStr

	resourceName := "oci_cloud_migrations_migration_asset.test_migration_asset"
	datasourceName := "data.oci_cloud_migrations_migration_assets.test_migration_assets"
	singularDatasourceName := "data.oci_cloud_migrations_migration_asset.test_migration_asset"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+CloudMigrationsMigrationAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_asset", "test_migration_asset", acctest.Optional, acctest.Create, CloudMigrationsMigrationAssetRepresentation), "cloudmigrations", "migrationAsset", t)

	acctest.ResourceTest(t, testAccCheckCloudMigrationsMigrationAssetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + CloudMigrationsMigrationAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_asset", "test_migration_asset", acctest.Required, acctest.Create, CloudMigrationsMigrationAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_asset_id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "snap_shot_bucket_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variableStr + CloudMigrationsMigrationAssetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variableStr + CloudMigrationsMigrationAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_asset", "test_migration_asset", acctest.Optional, acctest.Create, CloudMigrationsMigrationAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_asset_id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_schedule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "snap_shot_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

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
			Config: config + variableStr + CloudMigrationsMigrationAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_asset", "test_migration_asset", acctest.Optional, acctest.Update, CloudMigrationsMigrationAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_asset_id"),
				resource.TestCheckResourceAttrSet(resourceName, "migration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_schedule_id"),
				resource.TestCheckResourceAttrSet(resourceName, "snap_shot_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_migration_assets", "test_migration_assets", acctest.Optional, acctest.Update, CloudMigrationsCloudMigrationsMigrationAssetDataSourceRepresentation) +
				variableStr + CloudMigrationsMigrationAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_migrations_migration_asset", "test_migration_asset", acctest.Optional, acctest.Update, CloudMigrationsMigrationAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "migration_asset_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "migration_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "migration_asset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "migration_asset_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_migrations_migration_asset", "test_migration_asset", acctest.Required, acctest.Create, CloudMigrationsCloudMigrationsMigrationAssetSingularDataSourceRepresentation) +
				variableStr + CloudMigrationsMigrationAssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "migration_asset_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notifications.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_asset_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// verify resource import
		{
			Config:            config + CloudMigrationsMigrationAssetRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"inventory_asset_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCloudMigrationsMigrationAssetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_migrations_migration_asset" {
			noResourceFound = false
			request := oci_cloud_migrations.GetMigrationAssetRequest{}

			tmp := rs.Primary.ID
			request.MigrationAssetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")

			response, err := client.GetMigrationAsset(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_migrations.MigrationAssetLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudMigrationsMigrationAsset") {
		resource.AddTestSweepers("CloudMigrationsMigrationAsset", &resource.Sweeper{
			Name:         "CloudMigrationsMigrationAsset",
			Dependencies: acctest.DependencyGraph["migrationAsset"],
			F:            sweepCloudMigrationsMigrationAssetResource,
		})
	}
}

func sweepCloudMigrationsMigrationAssetResource(compartment string) error {
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()
	migrationAssetIds, err := getCloudMigrationsMigrationAssetIds(compartment)
	if err != nil {
		return err
	}
	for _, migrationAssetId := range migrationAssetIds {
		if ok := acctest.SweeperDefaultResourceId[migrationAssetId]; !ok {
			deleteMigrationAssetRequest := oci_cloud_migrations.DeleteMigrationAssetRequest{}

			deleteMigrationAssetRequest.MigrationAssetId = &migrationAssetId

			deleteMigrationAssetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_migrations")
			_, error := migrationClient.DeleteMigrationAsset(context.Background(), deleteMigrationAssetRequest)
			if error != nil {
				fmt.Printf("Error deleting MigrationAsset %s %s, It is possible that the resource is already deleted. Please verify manually \n", migrationAssetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &migrationAssetId, CloudMigrationsMigrationAssetSweepWaitCondition, time.Duration(3*time.Minute),
				CloudMigrationsMigrationAssetSweepResponseFetchOperation, "cloud_migrations", true)
		}
	}
	return nil
}

func getCloudMigrationsMigrationAssetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MigrationAssetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	migrationClient := acctest.GetTestClients(&schema.ResourceData{}).MigrationClient()

	listMigrationAssetsRequest := oci_cloud_migrations.ListMigrationAssetsRequest{}
	listMigrationAssetsRequest.MigrationId = &compartmentId //TODO WRONG
	listMigrationAssetsRequest.LifecycleState = oci_cloud_migrations.MigrationAssetLifecycleStateActive
	listMigrationAssetsResponse, err := migrationClient.ListMigrationAssets(context.Background(), listMigrationAssetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MigrationAsset list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, migrationAsset := range listMigrationAssetsResponse.Items {
		id := *migrationAsset.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MigrationAssetId", id)
	}
	return resourceIds, nil
}

func CloudMigrationsMigrationAssetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if migrationAssetResponse, ok := response.Response.(oci_cloud_migrations.GetMigrationAssetResponse); ok {
		return migrationAssetResponse.LifecycleState != oci_cloud_migrations.MigrationAssetLifecycleStateDeleted
	}
	return false
}

func CloudMigrationsMigrationAssetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MigrationClient().GetMigrationAsset(context.Background(), oci_cloud_migrations.GetMigrationAssetRequest{
		MigrationAssetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
