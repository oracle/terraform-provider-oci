// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisOciCacheBackupExportToObjectStorageRequiredOnlyResource = RedisOciCacheBackupExportToObjectStorageResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup_export_to_object_storage", "test_oci_cache_backup_export_to_object_storage", acctest.Required, acctest.Create, RedisOciCacheBackupExportToObjectStorageRepresentation)

	RedisOciCacheBackupExportToObjectStorageRepresentation = map[string]interface{}{
		"bucket":              acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"oci_cache_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_oci_cache_backup.test_oci_cache_backup.id}`},
		"prefix":              acctest.Representation{RepType: acctest.Optional, Create: `prefix`},
	}

	RedisOciCacheBackupExportTestBackupRepresentation = acctest.RepresentationCopyWithNewProperties(
		RedisOciCacheBackupRepresentation,
		map[string]interface{}{
			"source_cluster_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_redis_cluster.test_redis_cluster_non_sharded.id}`},
			"backup_source":            acctest.Representation{RepType: acctest.Optional, Create: `REPLICA`},
			"description":              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
			"retention_period_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		},
	)

	RedisOciCacheBackupExportToObjectStorageResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
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
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster_non_sharded", acctest.Required, acctest.Create, RedisRedisClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup", "test_oci_cache_backup", acctest.Optional, acctest.Create, RedisOciCacheBackupExportTestBackupRepresentation)
)

// issue-routing-tag: redis/default
func TestRedisOciCacheBackupExportToObjectStorageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisOciCacheBackupExportToObjectStorageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_redis_oci_cache_backup_export_to_object_storage.test_oci_cache_backup_export_to_object_storage"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisOciCacheBackupExportToObjectStorageResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup_export_to_object_storage", "test_oci_cache_backup_export_to_object_storage", acctest.Optional, acctest.Create, RedisOciCacheBackupExportToObjectStorageRepresentation), "redis", "ociCacheBackupExportToObjectStorage", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheBackupExportToObjectStorageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup_export_to_object_storage", "test_oci_cache_backup_export_to_object_storage", acctest.Required, acctest.Create, RedisOciCacheBackupExportToObjectStorageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "oci_cache_backup_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheBackupExportToObjectStorageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheBackupExportToObjectStorageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_backup_export_to_object_storage", "test_oci_cache_backup_export_to_object_storage", acctest.Optional, acctest.Create, RedisOciCacheBackupExportToObjectStorageRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bucket"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "oci_cache_backup_id"),
				resource.TestCheckResourceAttr(resourceName, "prefix", "prefix"),

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
			Config: config + compartmentIdVariableStr + RedisOciCacheBackupExportToObjectStorageResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_redis_oci_cache_backup_export_to_object_storage",
					"test_oci_cache_backup_export_to_object_storage",
					acctest.Optional,
					acctest.Create,
					RedisOciCacheBackupExportToObjectStorageRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) error {
					return emptyTestBucketObjects(s)
				},
			),
		},
	})
}
