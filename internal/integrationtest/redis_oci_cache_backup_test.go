// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisOciCacheBackupResource = RedisOciCacheBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Create, RedisOciCacheBackupRepresentation)

	RedisOciCacheBackupResourceConfig = RedisOciCacheBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Update, RedisOciCacheBackupRepresentation)

	RedisOciCacheBackupSingularDataSourceRepresentation = map[string]interface{}{
		"oci_cache_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_oci_cache_backup.test_oci_cache_backup.id}`},
	}

	RedisOciCacheBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"oci_cache_backup_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_redis_oci_cache_backup.test_oci_cache_backup.id}`},
		"source_cluster_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_redis_redis_cluster.test_redis_cluster_non_sharded.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheBackupDataSourceFilterRepresentation}}
	RedisOciCacheBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_oci_cache_backup.test_oci_cache_backup.id}`}},
	}

	RedisOciCacheBackupRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"source_cluster_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_redis_cluster.test_redis_cluster_non_sharded.id}`},
		"backup_source":            acctest.Representation{RepType: acctest.Optional, Create: `REPLICA`},
		"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"retention_period_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	RedisOciCacheBackupShardedRepresentation = acctest.RepresentationCopyWithNewProperties(
		RedisOciCacheBackupRepresentation,
		map[string]interface{}{
			"display_name": acctest.Representation{RepType: acctest.Required, Create: `displayNameShardedBackup`},
			"source_cluster_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  `${oci_redis_redis_cluster.test_redis_cluster_sharded.id}`,
			},
		},
	)

	RedisOciCacheBackupShardedResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_redis_oci_cache_backup",
		"test_oci_cache_backup_sharded",
		acctest.Optional,
		acctest.Create,
		RedisOciCacheBackupShardedRepresentation,
	)

	RedisOciCacheBackupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "redis_security_list", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSecurityListRepresentation, map[string]interface{}{
				"display_name": acctest.Representation{RepType: acctest.Required, Create: `redis-security-list`},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
				"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.redis_security_list.id}`}},
			})) +
		acctest.GenerateDataSourceFromRepresentationMap(
			"oci_objectstorage_namespace",
			"test_namespace",
			acctest.Optional,
			acctest.Create,
			ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation,
		) +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_objectstorage_bucket",
			"test_bucket",
			acctest.Required,
			acctest.Create,
			ObjectStorageBucketRepresentation,
		) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster_non_sharded", acctest.Required, acctest.Create, RedisRedisClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster_sharded", acctest.Required, acctest.Create, RedisRedisShardedClusterRepresentation)

	// deterministic export prefixes so object names are predictable
	RedisOciCacheBackupExportPrefixNonSharded = "tf-acc-export-nonsharded"
	RedisOciCacheBackupExportPrefixSharded    = "tf-acc-export-sharded"

	RedisOciCacheBackupExportToObjectStorageNonShardedRepresentation = map[string]interface{}{
		"bucket":              acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"oci_cache_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_oci_cache_backup.test_oci_cache_backup.id}`},
		"prefix":              acctest.Representation{RepType: acctest.Optional, Create: RedisOciCacheBackupExportPrefixNonSharded},
	}

	RedisOciCacheBackupExportToObjectStorageShardedRepresentation = map[string]interface{}{
		"bucket":              acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"oci_cache_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_oci_cache_backup.test_oci_cache_backup_sharded.id}`},
		"prefix":              acctest.Representation{RepType: acctest.Optional, Create: RedisOciCacheBackupExportPrefixSharded},
	}

	RedisOciCacheBackupExportToObjectStorageNonShardedResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_redis_oci_cache_backup_export_to_object_storage",
		"test_export_non_sharded",
		acctest.Required,
		acctest.Create,
		RedisOciCacheBackupExportToObjectStorageNonShardedRepresentation,
	)

	RedisOciCacheBackupExportToObjectStorageShardedResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_redis_oci_cache_backup_export_to_object_storage",
		"test_export_sharded",
		acctest.Required,
		acctest.Create,
		RedisOciCacheBackupExportToObjectStorageShardedRepresentation,
	)

	RedisRestoredFromBackupNonShardedHCL = `resource "oci_redis_redis_cluster" "test_restored_from_backup_non_sharded" {
  		compartment_id     = var.compartment_id
  		display_name       = "displayNameRestoredFromBackupNonSharded"
  		cluster_mode       = "NONSHARDED"
  		node_count         = 3
  		node_memory_in_gbs = 2
  		software_version   = "REDIS_7_0"
  		subnet_id          = oci_core_subnet.test_subnet.id

  		backup_id = oci_redis_oci_cache_backup.test_oci_cache_backup.id

  		lifecycle {
    		ignore_changes = [defined_tags, freeform_tags, system_tags]
	  	}
	}
`
	RedisRestoredFromBackupShardedHCL = `resource "oci_redis_redis_cluster" "test_restored_from_backup_sharded" {
  		compartment_id     = var.compartment_id
  		display_name       = "displayNameRestoredFromBackupSharded"
  		cluster_mode       = "SHARDED"
  		node_count         = 2
  		node_memory_in_gbs = 2
  		shard_count        = 3
  		software_version   = "REDIS_7_0"
  		subnet_id          = oci_core_subnet.test_subnet.id

		backup_id = oci_redis_oci_cache_backup.test_oci_cache_backup_sharded.id

  		lifecycle {
    		ignore_changes = [defined_tags, freeform_tags, system_tags]
  		}
	}
`

	RedisImportedClusterNonShardedHCL = `resource "oci_redis_redis_cluster" "test_imported_cluster_non_sharded" {
  		compartment_id     = var.compartment_id
  		display_name       = "displayNameImportedNonSharded"
  		cluster_mode       = "NONSHARDED"
  		node_count         = 3
  		node_memory_in_gbs = 2
  		software_version   = "REDIS_7_0"
  		subnet_id          = oci_core_subnet.test_subnet.id

  		import_from_object_storage_details {
    		bucket    = oci_objectstorage_bucket.test_bucket.name
    		namespace = data.oci_objectstorage_namespace.test_namespace.namespace

    		objects {
      			object = "dump.rdb"
    		}
  		}

		lifecycle {
    		ignore_changes = [
      			import_from_object_storage_details,
      			defined_tags,
      			freeform_tags,
      			system_tags,
    		]
  		}
	}

`

	RedisImportedClusterShardedHCL = `resource "oci_redis_redis_cluster" "test_imported_cluster_sharded" {
		compartment_id     = var.compartment_id
  		display_name       = "displayNameImportedSharded"
  		cluster_mode       = "SHARDED"
  		node_count         = 2
  		node_memory_in_gbs = 2
  		shard_count        = 3
  		software_version   = "REDIS_7_0"
  		subnet_id          = oci_core_subnet.test_subnet.id

  		import_from_object_storage_details {
    		bucket    = oci_objectstorage_bucket.test_bucket.name
			namespace = data.oci_objectstorage_namespace.test_namespace.namespace

			objects { object = "dump.001.rdb" }
			objects { object = "dump.002.rdb" }
			objects { object = "dump.003.rdb" }
		}

		lifecycle {
    		ignore_changes = [
      			import_from_object_storage_details,
      			defined_tags,
      			freeform_tags,
      			system_tags,
    		]
  		}

	}

`
)

// issue-routing-tag: redis/default
func TestRedisOciCacheBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisOciCacheBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_redis_oci_cache_backup.test_oci_cache_backup"
	shardedBackupResourceName := "oci_redis_oci_cache_backup.test_oci_cache_backup_sharded"
	datasourceName := "data.oci_redis_oci_cache_backups.test_oci_cache_backups"
	singularDatasourceName := "data.oci_redis_oci_cache_backup.test_oci_cache_backup"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisOciCacheBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Create, RedisOciCacheBackupRepresentation), "redis", "ociCacheBackup", t)

	acctest.ResourceTest(t, testAccCheckRedisOciCacheBackupDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Create, RedisOciCacheBackupRepresentation) +
				RedisOciCacheBackupShardedResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(shardedBackupResourceName, "id"),
				resource.TestCheckResourceAttr(shardedBackupResourceName, "display_name", "displayNameShardedBackup"),
				resource.TestCheckResourceAttrSet(shardedBackupResourceName, "source_cluster_id"),

				resource.TestCheckResourceAttrSet(resourceName, "backup_size_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "backup_source", "REPLICA"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_mode"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "software_version"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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

		{
			Config: config + compartmentIdVariableStr + RedisOciCacheBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Create, RedisOciCacheBackupRepresentation) +
				RedisOciCacheBackupShardedResource +
				RedisOciCacheBackupExportToObjectStorageNonShardedResource +
				RedisOciCacheBackupExportToObjectStorageShardedResource,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet("oci_redis_oci_cache_backup_export_to_object_storage.test_export_non_sharded", "id"),
				resource.TestCheckResourceAttrSet("oci_redis_oci_cache_backup_export_to_object_storage.test_export_sharded", "id"),
			),
		},

		{
			Config: config + compartmentIdVariableStr + RedisOciCacheBackupResourceDependencies +
				// Keep backups + export resources present
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Create, RedisOciCacheBackupRepresentation) +
				RedisOciCacheBackupShardedResource +
				RedisOciCacheBackupExportToObjectStorageNonShardedResource +
				RedisOciCacheBackupExportToObjectStorageShardedResource +
				RedisRestoredFromBackupNonShardedHCL +
				RedisRestoredFromBackupShardedHCL +
				// Now create imported clusters
				RedisImportedClusterNonShardedHCL + "\n" +
				RedisImportedClusterShardedHCL,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet("oci_redis_redis_cluster.test_imported_cluster_non_sharded", "id"),
				resource.TestCheckResourceAttr("oci_redis_redis_cluster.test_imported_cluster_non_sharded", "cluster_mode", "NONSHARDED"),
				resource.TestCheckResourceAttrSet("oci_redis_redis_cluster.test_imported_cluster_sharded", "id"),
				resource.TestCheckResourceAttr("oci_redis_redis_cluster.test_imported_cluster_sharded", "cluster_mode", "SHARDED"),
				func(s *terraform.State) error {
					return emptyTestBucketObjects(s)
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RedisOciCacheBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RedisOciCacheBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})) +
				RedisOciCacheBackupShardedResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backup_size_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "backup_source", "REPLICA"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_mode"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "software_version"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + RedisOciCacheBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Update, RedisOciCacheBackupRepresentation) +
				RedisOciCacheBackupShardedResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backup_size_in_gbs"),
				resource.TestCheckResourceAttr(resourceName, "backup_source", "REPLICA"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_mode"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period_in_days", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "software_version"),
				resource.TestCheckResourceAttrSet(resourceName, "source_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_backups", "test_oci_cache_backups", acctest.Optional, acctest.Update, RedisOciCacheBackupDataSourceRepresentation) +
				compartmentIdVariableStr + RedisOciCacheBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Update, RedisOciCacheBackupRepresentation) +
				RedisOciCacheBackupShardedResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "oci_cache_backup_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "oci_cache_backup_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oci_cache_backup_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Required, acctest.Create, RedisOciCacheBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RedisOciCacheBackupResourceConfig +
				RedisOciCacheBackupShardedResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oci_cache_backup_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_size_in_gbs"),
				resource.TestCheckResourceAttr(singularDatasourceName, "backup_source", "REPLICA"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_mode"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retention_period_in_days", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + RedisOciCacheBackupResource + RedisOciCacheBackupShardedResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckRedisOciCacheBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OciCacheBackupClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_redis_oci_cache_backup" {
			noResourceFound = false
			request := oci_redis.GetOciCacheBackupRequest{}

			tmp := rs.Primary.ID
			request.OciCacheBackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "redis")

			response, err := client.GetOciCacheBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_redis.OciCacheBackupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("RedisOciCacheBackup") {
		resource.AddTestSweepers("RedisOciCacheBackup", &resource.Sweeper{
			Name:         "RedisOciCacheBackup",
			Dependencies: acctest.DependencyGraph["ociCacheBackup"],
			F:            sweepRedisOciCacheBackupResource,
		})
	}
}

func sweepRedisOciCacheBackupResource(compartment string) error {
	ociCacheBackupClient := acctest.GetTestClients(&schema.ResourceData{}).OciCacheBackupClient()
	ociCacheBackupIds, err := getRedisOciCacheBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, ociCacheBackupId := range ociCacheBackupIds {
		if ok := acctest.SweeperDefaultResourceId[ociCacheBackupId]; !ok {
			deleteOciCacheBackupRequest := oci_redis.DeleteOciCacheBackupRequest{}

			deleteOciCacheBackupRequest.OciCacheBackupId = &ociCacheBackupId

			deleteOciCacheBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "redis")
			_, error := ociCacheBackupClient.DeleteOciCacheBackup(context.Background(), deleteOciCacheBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting OciCacheBackup %s %s, It is possible that the resource is already deleted. Please verify manually \n", ociCacheBackupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ociCacheBackupId, RedisOciCacheBackupSweepWaitCondition, time.Duration(3*time.Minute),
				RedisOciCacheBackupSweepResponseFetchOperation, "redis", true)
		}
	}
	return nil
}

func getRedisOciCacheBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OciCacheBackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	ociCacheBackupClient := acctest.GetTestClients(&schema.ResourceData{}).OciCacheBackupClient()

	listOciCacheBackupsRequest := oci_redis.ListOciCacheBackupsRequest{}
	listOciCacheBackupsRequest.CompartmentId = &compartmentId
	listOciCacheBackupsRequest.LifecycleState = oci_redis.OciCacheBackupLifecycleStateActive
	listOciCacheBackupsResponse, err := ociCacheBackupClient.ListOciCacheBackups(context.Background(), listOciCacheBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OciCacheBackup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ociCacheBackup := range listOciCacheBackupsResponse.Items {
		id := *ociCacheBackup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OciCacheBackupId", id)
	}
	return resourceIds, nil
}

func RedisOciCacheBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ociCacheBackupResponse, ok := response.Response.(oci_redis.GetOciCacheBackupResponse); ok {
		return ociCacheBackupResponse.LifecycleState != oci_redis.OciCacheBackupLifecycleStateDeleted
	}
	return false
}

func RedisOciCacheBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OciCacheBackupClient().GetOciCacheBackup(context.Background(), oci_redis.GetOciCacheBackupRequest{
		OciCacheBackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func emptyTestBucketObjects(s *terraform.State) error {
	namespace, err := acctest.FromInstanceState(s, "data.oci_objectstorage_namespace.test_namespace", "namespace")
	if err != nil {
		return err
	}

	bucketName, err := acctest.FromInstanceState(s, "oci_objectstorage_bucket.test_bucket", "name")
	if err != nil {
		return err
	}

	objectStorageClient := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ObjectStorageClient()

	listReq := oci_object_storage.ListObjectsRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketName,
	}

	listResp, err := objectStorageClient.ListObjects(context.Background(), listReq)
	if err != nil {
		return err
	}

	for _, obj := range listResp.ListObjects.Objects {
		if obj.Name == nil || *obj.Name == "" {
			continue
		}

		delReq := oci_object_storage.DeleteObjectRequest{
			NamespaceName: &namespace,
			BucketName:    &bucketName,
			ObjectName:    obj.Name,
		}

		if _, err := objectStorageClient.DeleteObject(context.Background(), delReq); err != nil {
			return err
		}
	}

	return nil
}
