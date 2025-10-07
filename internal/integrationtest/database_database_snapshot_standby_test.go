// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	aDBName1 = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	aDBName2 = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	DatabaseDatabaseSnapshotStandybyRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot_standby", "test_database_snapshot_standby", acctest.Required, acctest.Create, DatabaseDatabaseSnapshotStandybyRepresentation)

	DatabaseDatabaseSnapshotStandybyRepresentation = map[string]interface{}{
		"database_admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"database_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database.test_database_2.id}`},
		"standby_conversion_type":   acctest.Representation{RepType: acctest.Required, Create: `SNAPSHOT`},
		"snapshot_duration_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	DatabaseSnapShotStandbyPrimaryDbRepresentation = map[string]interface{}{
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19.24.0.0`},
		"lifecycle":  acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	DatabaseSnapShotStandbyDbRepresentation = map[string]interface{}{
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_2.id}`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `DATAGUARD`},
		"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19.24.0.0`},
		"lifecycle":  acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	DatabaseExaccSnapShotStandbyPrimaryDbRepresentation = map[string]interface{}{
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_vm_cluster.id}`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19.23.0.0`},
	}

	DatabaseExaccSnapShotStandbyStandbyDbRepresentation = map[string]interface{}{
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_vm_cluster_2.id}`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `DATAGUARD`},
	}

	databaseExaccStandbyDatabaseRepresentation = map[string]interface{}{
		"database_admin_password":    acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"db_name":                    acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
		"character_set":              acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"db_unique_name":             acctest.Representation{RepType: acctest.Optional, Create: aDBName2},
		"db_workload":                acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"ncharacter_set":             acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"protection_mode":            acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`},
		"transport_type":             acctest.Representation{RepType: acctest.Optional, Create: `ASYNC`},
		"source_tde_wallet_password": acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`},
		"sid_prefix":                 acctest.Representation{RepType: acctest.Optional, Create: `myTestDb72`},
	}

	databaseExaccDatabaseRepresentation = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
		"character_set":  acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"db_unique_name": acctest.Representation{RepType: acctest.Optional, Create: aDBName1},
		"db_workload":    acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set": acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"sid_prefix":     acctest.Representation{RepType: acctest.Optional, Create: `myTestDb71`},
	}

	databaseSnapShotDb1Representation = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `myTestDb`},
		"character_set":  acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"db_unique_name": acctest.Representation{RepType: acctest.Optional, Create: aDBName1},
		"ncharacter_set": acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
	}
	dbHomeExaccRepresentationSourceVmCluster = map[string]interface{}{
		"vm_cluster_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"source":                      acctest.Representation{RepType: acctest.Required, Create: `VM_CLUSTER_NEW`},
		"db_version":                  acctest.Representation{RepType: acctest.Required, Create: `19.25.0.0`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `TFTestDbHome1`},
		"is_unified_auditing_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"freeformTags": "freeformTags"}},
	}

	ExaBaseDependenciesForSnapshotStandby = ad_subnet_security +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseCloudExadataInfrastructureRepresentation, []string{"display_name"}), map[string]interface{}{
			"display_name": acctest.Representation{RepType: acctest.Required, Create: `tstExaInfra2`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseCloudVmClusterRepresentation, []string{"cloud_exadata_infrastructure_id", "file_system_configuration_details"}), map[string]interface{}{
			"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure_2.id}`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabaseDbHomeRepresentationBase3, map[string]interface{}{
			"db_version":   acctest.Representation{RepType: acctest.Required, Create: `19.24.0.0`},
			"source":       acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
			"display_name": acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeNone2`},
		}))

	ExaccDependenciesForSnapshotStandby = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "peer_exadata_infrastructure", acctest.Required, acctest.Create, peerExadataInfraRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "peer_vm_cluster_network", acctest.Required, acctest.Create, DatabasePeerVmClusterNetworkRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "peer_db_servers", acctest.Required, acctest.Create, DatabaseDatabasePeerExaInfraDbServerDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseVmClusterRepresentation, []string{"display_name", "exadata_infrastructure_id", "vm_cluster_network_id", "db_servers", "cpu_core_count"}), map[string]interface{}{
			"display_name":              acctest.Representation{RepType: acctest.Required, Create: `vmCluster2`},
			"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.peer_exadata_infrastructure.id}`},
			"vm_cluster_network_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.peer_vm_cluster_network.id}`},
			"db_servers":                acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_database_db_servers.peer_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.peer_db_servers.db_servers.1.id}`}},
			"cpu_core_count":            acctest.Representation{RepType: acctest.Required, Create: `16`, Update: `20`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(dbHomeExaccRepresentationSourceVmCluster, []string{"display_name", "vm_cluster_id", "db_version"}), map[string]interface{}{
			"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster_2.id}`},
			"display_name":  acctest.Representation{RepType: acctest.Required, Create: `TFTestDbHome2`},
			"db_version":    acctest.Representation{RepType: acctest.Required, Create: `19.27.0.0`},
		}))

	ExaccDatabaseResourceDependencies = DatabaseVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseVmClusterRepresentation, []string{"cpu_core_count"}), map[string]interface{}{
			"cpu_core_count": acctest.Representation{RepType: acctest.Required, Create: `16`, Update: `20`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseKeyStoreRepresentation) + KmsVaultIdVariableStr + OkvSecretVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster", acctest.Required, acctest.Create, dbHomeRepresentationSourceVmClusterExacc)
)

// issue-routing-tag: database/default
func TestDatabaseDatabaseSnapshotStandybyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSnapshotStandybyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	primaryDatabase := "oci_database_database.test_database"
	standbyDatabase := "oci_database_database.test_database_2"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDatabaseResourceDependencies+kmsKeyIdVariableStr+vaultIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, DatabaseDatabaseRepresentation), "database", "database", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseDestroy, []resource.TestStep{
		// verify Add Standby
		{
			Config: config + compartmentIdVariableStr + ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentationSourceNone2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabaseSnapShotStandbyPrimaryDbRepresentation, map[string]interface{}{
					"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseSnapShotDb1Representation},
				})) +
				ExaBaseDependenciesForSnapshotStandby +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database_2", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabaseSnapShotStandbyDbRepresentation, map[string]interface{}{
					"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(databaseMultipleStandbyDb2Representation, map[string]interface{}{
						"source_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.id}`},
					})},
				})) +
				DatabaseDatabaseSnapshotStandybyRequiredOnlyResource,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(primaryDatabase, "database.#", "1"),
				resource.TestCheckResourceAttr(primaryDatabase, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(primaryDatabase, "db_home_id"),
				resource.TestCheckResourceAttr(primaryDatabase, "source", "NONE"),
				resource.TestCheckResourceAttr(standbyDatabase, "database.#", "1"),
				resource.TestCheckResourceAttr(standbyDatabase, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(standbyDatabase, "db_home_id"),
				resource.TestCheckResourceAttr(standbyDatabase, "source", "DATAGUARD"),
				resource.TestCheckResourceAttrSet(standbyDatabase, "data_guard_group.#"),
				resource.TestCheckResourceAttr(standbyDatabase, "data_guard_group.0.protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttr(standbyDatabase, "data_guard_group.0.members.0.role", "PRIMARY"),
				resource.TestCheckResourceAttr(standbyDatabase, "data_guard_group.0.members.1.role", "STANDBY"),
			),
		},
		{
			Config: config + compartmentIdVariableStr + ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentationSourceNone2) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabaseSnapShotStandbyPrimaryDbRepresentation, map[string]interface{}{
					"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseSnapShotDb1Representation},
				})) +
				ExaBaseDependenciesForSnapshotStandby +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database_2", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabaseSnapShotStandbyDbRepresentation, map[string]interface{}{
					"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(databaseMultipleStandbyDb2Representation, map[string]interface{}{
						"source_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.id}`},
					})},
				})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot_standby", "test_database_snapshot_standby", acctest.Required, acctest.Update, acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseSnapshotStandybyRepresentation, map[string]interface{}{
					"source_database_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.id}`},
					"standby_conversion_type": acctest.Representation{RepType: acctest.Required, Update: `PHYSICAL`},
				})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(primaryDatabase, "database.#", "1"),
				resource.TestCheckResourceAttr(primaryDatabase, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(primaryDatabase, "db_home_id"),
				resource.TestCheckResourceAttr(primaryDatabase, "source", "NONE"),
				resource.TestCheckResourceAttr(standbyDatabase, "database.#", "1"),
				resource.TestCheckResourceAttr(standbyDatabase, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(standbyDatabase, "db_home_id"),
				resource.TestCheckResourceAttr(standbyDatabase, "source", "DATAGUARD"),
				resource.TestCheckResourceAttrSet(standbyDatabase, "data_guard_group.#"),
				resource.TestCheckResourceAttr(standbyDatabase, "data_guard_group.0.protection_mode", "MAXIMUM_PERFORMANCE"),
			),
		},
	})
}

func TestEXACCDatabaseSnapshotStandybyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseSnapshotStandybyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	primaryDatabase := "oci_database_database.test_database"
	standbyDatabase := "oci_database_database.test_database_2"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExaccRequiredOnlyResource+vaultIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, DatabaseExaccDatabaseRepresentation), "database", "database", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseDestroy, []resource.TestStep{
		// verify convert Standby
		{
			Config: config + compartmentIdVariableStr + ExaccDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabaseExaccSnapShotStandbyPrimaryDbRepresentation, map[string]interface{}{
					"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseExaccDatabaseRepresentation},
				})) +
				ExaccDependenciesForSnapshotStandby +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database_2", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DatabaseExaccSnapShotStandbyStandbyDbRepresentation, map[string]interface{}{
					"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(databaseExaccStandbyDatabaseRepresentation, map[string]interface{}{
						"source_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.id}`},
					})},
				})) +
				DatabaseDatabaseSnapshotStandybyRequiredOnlyResource,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(primaryDatabase, "database.#", "1"),
				resource.TestCheckResourceAttr(primaryDatabase, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(primaryDatabase, "db_home_id"),
				resource.TestCheckResourceAttr(primaryDatabase, "source", "NONE"),
				resource.TestCheckResourceAttr(standbyDatabase, "database.#", "1"),
				resource.TestCheckResourceAttr(standbyDatabase, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(standbyDatabase, "db_home_id"),
				resource.TestCheckResourceAttr(standbyDatabase, "source", "DATAGUARD"),
				resource.TestCheckResourceAttrSet(standbyDatabase, "data_guard_group.#"),
				resource.TestCheckResourceAttr(standbyDatabase, "data_guard_group.0.protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttr(standbyDatabase, "data_guard_group.0.members.0.role", "PRIMARY"),
				resource.TestCheckResourceAttr(standbyDatabase, "data_guard_group.0.members.1.role", "STANDBY"),
			),
		},
		{
			Config: config + compartmentIdVariableStr + DatabaseExaccDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DatabaseExaccSnapShotStandbyPrimaryDbRepresentation, map[string]interface{}{
					"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseExaccDatabaseRepresentation},
				})) +
				ExaccDependenciesForSnapshotStandby +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database_2", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(DatabaseExaccSnapShotStandbyStandbyDbRepresentation, map[string]interface{}{
					"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(databaseExaccStandbyDatabaseRepresentation, map[string]interface{}{
						"source_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.id}`},
					})},
				})) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot_standby", "test_database_snapshot_standby", acctest.Required, acctest.Update, acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseSnapshotStandybyRepresentation, map[string]interface{}{
					"source_database_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.id}`},
					"standby_conversion_type": acctest.Representation{RepType: acctest.Required, Update: `PHYSICAL`},
				})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(primaryDatabase, "database.#", "1"),
				resource.TestCheckResourceAttr(primaryDatabase, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(primaryDatabase, "db_home_id"),
				resource.TestCheckResourceAttr(primaryDatabase, "source", "NONE"),
				resource.TestCheckResourceAttr(standbyDatabase, "database.#", "1"),
				resource.TestCheckResourceAttr(standbyDatabase, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(standbyDatabase, "db_home_id"),
				resource.TestCheckResourceAttr(standbyDatabase, "source", "DATAGUARD"),
				resource.TestCheckResourceAttrSet(standbyDatabase, "data_guard_group.#"),
				resource.TestCheckResourceAttr(standbyDatabase, "data_guard_group.0.protection_mode", "MAXIMUM_PERFORMANCE"),
			),
		},
	})
}
