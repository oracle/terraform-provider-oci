// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDbHomeRequiredOnlyResource = DatabaseDbHomeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Required, acctest.Create, dbHomeRepresentationSourceNone)

	DatabaseDbHomeResourceConfig = DatabaseDbHomeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Optional, acctest.Update, dbHomeRepresentationSourceNone)

	DatabaseDatabaseDbHomeSingularDataSourceRepresentation = map[string]interface{}{
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_source_none.id}`},
	}

	DatabaseDatabaseDbHomeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_system_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeNone`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDbHomeDataSourceFilterRepresentation}}

	DatabaseDbHomeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_db_home.test_db_home_source_none.id}`}},
	}

	DatabaseDbHomeRepresentation = map[string]interface{}{
		"database":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDbHomeDatabaseRepresentation},
		"database_software_image_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database_software_image.test_database_software_image.id}`},
		"db_system_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"db_version":                 acctest.Representation{RepType: acctest.Required, Create: `12.1.0.2`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `createdDbHome`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_desupported_version":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `DB_BACKUP`},
		"vm_cluster_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	DatabaseDbHomeDatabaseRepresentation = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `tfDbNam`},
	}
	DatabaseDbHomeRepresentationBase = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
	}
	dbHomeRepresentationSourceNone = acctest.RepresentationCopyWithNewProperties(DatabaseDbHomeRepresentationBase, map[string]interface{}{
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeDatabaseRepresentationSourceNone},
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `12.1.0.2`},
		"source":       acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeNone`},
	})
	dbHomeDatabaseRepresentationSourceNone = map[string]interface{}{
		"admin_password":      acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`}, // Update password not supported for exa
		"tde_wallet_password": acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`}, // Update password not supported for exa
		"db_name":             acctest.Representation{RepType: acctest.Required, Create: `dbNone`},
		"character_set":       acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_backup_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbHomeDatabaseDbBackupConfigRepresentation},
		"db_workload":         acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":      acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":            acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
	}
	dbHomeRepresentationSourceNoneRequiredOnly = acctest.RepresentationCopyWithNewProperties(dbHomeRepresentationSourceNone, map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeDatabaseRepresentationSourceNoneRequiredOnly},
	})
	dbHomeDatabaseRepresentationSourceNoneRequiredOnly = acctest.RepresentationCopyWithNewProperties(dbHomeDatabaseRepresentationSourceNone, map[string]interface{}{
		"db_name": acctest.Representation{RepType: acctest.Required, Create: `dbNone0`},
	})
	dbHomeDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"auto_backup_window":        acctest.Representation{RepType: acctest.Optional, Create: `SLOT_TWO`},
		"auto_full_backup_day":      acctest.Representation{RepType: acctest.Optional, Create: `SUNDAY`, Update: `MONDAY`},
		"auto_full_backup_window":   acctest.Representation{RepType: acctest.Optional, Create: `SLOT_ONE`, Update: `SLOT_THREE`},
		"recovery_window_in_days":   acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"run_immediate_full_backup": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	dbHomeRepresentationSourceDbBackup = acctest.RepresentationCopyWithNewProperties(DatabaseDbHomeRepresentationBase, map[string]interface{}{
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeDatabaseRepresentationSourceDbBackup},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `DB_BACKUP`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `createdDbHomeBackup`},
	})
	dbHomeDatabaseRepresentationSourceDbBackup = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"backup_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_database_backup.test_backup.id}`},
		// TDE wallet password is not required when backups are encrypted with customer-managed (Vault service) keys.
		// "backup_tde_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		// Modifying db_name as mandatory. If not mandatory test fails with error "The specified database name 'tfDbName' exists."
		// The test takes the backup of the DB created in the db_system which has the db_name=tfDbName.
		// When db_home is created with source as "DB_BACKUP" and db_name is not provided, Service uses the db_name from the backup which is causing this test to fail.
		"db_name": acctest.Representation{RepType: acctest.Required, Create: `dbBack`},
	}

	dbHomeRepresentationSourceVmClusterNew = map[string]interface{}{
		"database":      acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeDatabaseRepresentationSourceVmClusterNew},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeVm`},
		"source":        acctest.Representation{RepType: acctest.Required, Create: `VM_CLUSTER_NEW`},
		"db_version":    acctest.Representation{RepType: acctest.Required, Create: `12.1.0.2`},
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	dbHomeDatabaseRepresentationSourceVmClusterNew = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"character_set":    acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbHomeDatabaseDbBackupConfigVmClusterNewRepresentation},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `dbVMClus`},
		"db_workload":      acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":         acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
	}

	dbHomeDatabaseDbBackupConfigVmClusterNewRepresentation = map[string]interface{}{
		"auto_backup_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"auto_backup_window":         acctest.Representation{RepType: acctest.Optional, Create: `SLOT_TWO`},
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseDbHomeDatabaseDbBackupConfigBackupDestinationDetailsRepresentation},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	DatabaseDbHomeDatabaseDbBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `NFS`},
	}
	dbHomeRepresentationSourceDatabase = acctest.RepresentationCopyWithNewProperties(DatabaseDbHomeRepresentationBase, map[string]interface{}{
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeDatabaseRepresentationSourceDatabase},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeDatabase`},
	})
	dbHomeDatabaseRepresentationSourceDatabase = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		// TDE wallet password is not required when backups are encrypted with customer-managed (Vault service) keys.
		// "backup_tde_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database.db.id}`},
		"db_name":     acctest.Representation{RepType: acctest.Required, Create: `dbDb`},
	}

	dbHomeRepresentationSourceVmClusterDatabase = map[string]interface{}{
		"database":      acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeDatabaseRepresentationSourceVmClusterDatabase},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeVmClusterDatabase`},
		"source":        acctest.Representation{RepType: acctest.Required, Create: `VM_CLUSTER_DATABASE`},
		"db_version":    acctest.Representation{RepType: acctest.Required, Create: `12.1.0.2`},
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	dbHomeDatabaseRepresentationSourceVmClusterDatabase = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"character_set":  acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		//"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbHomeDatabaseDbBackupConfigVmClusterDatabaseRepresentation},
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `dbVMClusDb`},
		"db_workload":    acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set": acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":       acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
	}

	dbHomeRepresentationSourceVmCluster = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
		"source":        acctest.Representation{RepType: acctest.Required, Create: `VM_CLUSTER_NEW`},
		"db_version":    acctest.Representation{RepType: acctest.Required, Create: `12.1.0.2`},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `TFTestDbHome1`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"freeformTags": "freeformTags"}},
	}

	DatabaseDatabaseExacsRepresentation = map[string]interface{}{
		"database":         acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation},
		"db_home_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_vm_cluster_no_db.id}`},
		"source":           acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"db_version":       acctest.Representation{RepType: acctest.Optional, Create: `12.1.0.2`},
		"kms_key_id":       acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_rotation": acctest.Representation{RepType: acctest.Optional, Update: `1`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
	}

	nesstedDatabaseRepresentationSourceNone = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`}, // Update password not supported for exa
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `nestDb`},
	}

	nesstedDatabaseRepresentationSourceNone2 = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`}, // Update password not supported for exa
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `nestDb2`},
	}

	createManualDbBackup = acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup_for_db", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(backupDatabaseRepresentation, map[string]interface{}{"display_name": acctest.Representation{RepType: acctest.Required, Create: `MonthlyBackup2`}}))

	DatabaseDbHomeResourceDependencies = DatabaseBackupResourceDbHomeDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationNFSRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": acctest.Representation{RepType: acctest.Optional, Update: activationFilePath}})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, DatabaseDatabaseDbServerDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, vmClusterNetworkValidateRepresentation)

	//DbHomeResourceVmClusterDependencies = ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig + KeyResourceDependencyConfig

	DbHomeResourceVmClusterDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudVmClusterRepresentation) +
		AvailabilityDomainConfig + DatabaseCloudVmClusterResourceDependencies + DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseDbHomeTdeWalletPassword(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomeTdeWalletPassword")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_home.test_db_home"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDbHomeDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Required, acctest.Create, dbHomeRepresentationSourceNoneRequiredOnly),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone0"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Optional, acctest.Create, dbHomeRepresentationSourceNone),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.tde_wallet_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_window", "SLOT_TWO"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_full_backup_day", "SUNDAY"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_full_backup_window", "SLOT_ONE"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.run_immediate_full_backup", "false"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHomeNone"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "source", "NONE"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "state"),
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Optional, acctest.Update, dbHomeRepresentationSourceNone),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.tde_wallet_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_full_backup_day", "MONDAY"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_full_backup_window", "SLOT_THREE"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.run_immediate_full_backup", "true"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHomeNone"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "source", "NONE"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "state"),
				),
			},
		},
	})
}

// issue-routing-tag: database/default
func TestDatabaseDbHomeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_home.test_db_home"
	datasourceName := "data.oci_database_db_homes.test_db_homes"
	singularDatasourceName := "data.oci_database_db_home.get_test_db_home"

	//var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+kmsKeyIdVariableStr+vaultIdVariableStr+kmsKeyVersionIdVariableStr+DatabaseDbHomeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Optional, acctest.Create, DatabaseDbHomeRepresentation), "database", "dbHome", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDbHomeDestroy, []resource.TestStep{
		// Create all dependencies first because test_db_home_source_database can trigger at the same time as the backup creation
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Required, acctest.Create, dbHomeRepresentationSourceDbBackup),
		},
		// Create second backup for create db_home source database
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies + createManualDbBackup +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Required, acctest.Create, dbHomeRepresentationSourceDbBackup),
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies + createManualDbBackup +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Required, acctest.Create, dbHomeRepresentationSourceNoneRequiredOnly) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Required, acctest.Create, dbHomeRepresentationSourceDbBackup) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_vm_cluster_new", acctest.Required, acctest.Create, dbHomeRepresentationSourceVmClusterNew) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", acctest.Required, acctest.Create, dbHomeRepresentationSourceDatabase),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone0"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),

				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "database.0.backup_id"),
				//resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "source", "DB_BACKUP"),

				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "source", "VM_CLUSTER_NEW"),

				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.admin_password", "BEstrO0ng_#11"),
				//resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "database.0.database_id"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "source", "DATABASE"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies,
		},
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Optional, acctest.Create, dbHomeRepresentationSourceDbBackup),
		},
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies + createManualDbBackup +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Required, acctest.Create, dbHomeRepresentationSourceDbBackup),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies + createManualDbBackup +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Optional, acctest.Create, dbHomeRepresentationSourceNone) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Optional, acctest.Create, dbHomeRepresentationSourceDbBackup) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_vm_cluster_new", acctest.Optional, acctest.Create, dbHomeRepresentationSourceVmClusterNew) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", acctest.Optional, acctest.Create, dbHomeRepresentationSourceDatabase),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttr(resourceName, "is_desupported_version", "false"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_window", "SLOT_TWO"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "database.0.backup_id"),
				//resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.db_name", "dbBack"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "display_name", "createdDbHomeBackup"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "source", "DB_BACKUP"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.auto_backup_window", "SLOT_TWO"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_name", "dbVMClus"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "display_name", "createdDbHomeVm"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "source", "VM_CLUSTER_NEW"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.admin_password", "BEstrO0ng_#11"),
				//resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "database.0.database_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.db_name", "dbDb"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "display_name", "createdDbHomeDatabase"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "state"),
				//resource.TestCheckResourceAttr(resourceName, "is_desupported_version", "false"),

				//func(s *terraform.State) (err error) {
				//	resId, err = acctest.FromInstanceState(s, resourceName, "id")
				//	if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
				//		if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
				//			return errExport
				//		}
				//	}
				//	return err
				//},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies + createManualDbBackup +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Optional, acctest.Update, dbHomeRepresentationSourceNone) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Optional, acctest.Update, dbHomeRepresentationSourceDbBackup) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_vm_cluster_new", acctest.Optional, acctest.Update, dbHomeRepresentationSourceVmClusterNew) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", acctest.Optional, acctest.Update, dbHomeRepresentationSourceDatabase),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				//resource.TestCheckResourceAttr(resourceName, "is_desupported_version", "false"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "database.0.backup_id"),
				//resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.db_name", "dbBack"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "display_name", "createdDbHomeBackup"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "source", "DB_BACKUP"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.auto_backup_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_name", "dbVMClus"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "display_name", "createdDbHomeVm"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "source", "VM_CLUSTER_NEW"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.admin_password", "BEstrO0ng_#11"),
				//resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "database.0.database_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.db_name", "dbDb"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "display_name", "createdDbHomeDatabase"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "state"),
				//resource.TestCheckResourceAttr(resourceName, "is_desupported_version", "true"),
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "test_db_homes", acctest.Optional, acctest.Update, DatabaseDatabaseDbHomeDataSourceRepresentation) +
				compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies + createManualDbBackup +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Optional, acctest.Update, dbHomeRepresentationSourceNone) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Optional, acctest.Update, dbHomeRepresentationSourceDbBackup) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", acctest.Optional, acctest.Update, dbHomeRepresentationSourceDatabase),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "db_homes.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "db_homes.0.db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(datasourceName, "db_homes.0.display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_home", "get_test_db_home", acctest.Required, acctest.Create, DatabaseDatabaseDbHomeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr + DatabaseDbHomeResourceDependencies + createManualDbBackup +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", acctest.Optional, acctest.Update, dbHomeRepresentationSourceNone) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", acctest.Optional, acctest.Update, dbHomeRepresentationSourceDbBackup) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", acctest.Optional, acctest.Update, dbHomeRepresentationSourceDatabase),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseDbHomeRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"database",
				"is_desupported_version",
				"database.0.admin_password",
				"kms_key_version_id",
			},
			ResourceName: resourceName + "_source_none",
		},
	})
}

// issue-routing-tag: database/default
func TestDatabaseDbHomeResource_exacs(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomeResource_exacs")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_home.test_db_home"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDbHomeDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_no_db", acctest.Optional, acctest.Create, dbHomeRepresentationSourceVmCluster),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "source", "VM_CLUSTER_NEW"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "display_name", "TFTestDbHome1"),
					resource.TestCheckResourceAttrSet(resourceName+"_vm_cluster_no_db", "vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "db_version", "12.1.0.2"),
				),
			},

			// Create DB outside of dbHome
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_no_db", acctest.Optional, acctest.Create, dbHomeRepresentationSourceVmCluster) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_db_vm_cluster_no_db", acctest.Required, acctest.Create, DatabaseDatabaseExacsRepresentation),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "source", "VM_CLUSTER_NEW"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "display_name", "TFTestDbHome1"),
					resource.TestCheckResourceAttrSet(resourceName+"_vm_cluster_no_db", "vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "db_version", "12.1.0.2"),
					resource.TestCheckNoResourceAttr(resourceName+"_vm_cluster_no_db", "database"),
				),
			},

			//Update DB home
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_no_db", acctest.Optional, acctest.Update, dbHomeRepresentationSourceVmCluster) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_db_vm_cluster_no_db", acctest.Required, acctest.Create, DatabaseDatabaseExacsRepresentation),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "source", "VM_CLUSTER_NEW"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "display_name", "TFTestDbHome1"),
					resource.TestCheckResourceAttrSet(resourceName+"_vm_cluster_no_db", "vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "db_version", "12.1.0.2"),
					resource.TestCheckNoResourceAttr(resourceName+"_vm_cluster_no_db", "database"),
				),
			},
			// Create DB inside of dbHome
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_no_db", acctest.Optional, acctest.Update,
						acctest.RepresentationCopyWithNewProperties(dbHomeRepresentationSourceVmCluster, map[string]interface{}{
							"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: nesstedDatabaseRepresentationSourceNone},
						})) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_db_vm_cluster_no_db", acctest.Required, acctest.Create, DatabaseDatabaseExacsRepresentation),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "source", "VM_CLUSTER_NEW"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "display_name", "TFTestDbHome1"),
					resource.TestCheckResourceAttrSet(resourceName+"_vm_cluster_no_db", "vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "database.#", "1"),
					func(s *terraform.State) (err error) {
						time.Sleep(3 * time.Minute)
						return
					},
				),
			},

			// Delete DB inside of dbHome
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_no_db", acctest.Optional, acctest.Update, dbHomeRepresentationSourceVmCluster) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_db_vm_cluster_no_db", acctest.Required, acctest.Create, DatabaseDatabaseExacsRepresentation),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "source", "VM_CLUSTER_NEW"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "display_name", "TFTestDbHome1"),
					resource.TestCheckResourceAttrSet(resourceName+"_vm_cluster_no_db", "vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "db_version", "12.1.0.2"),
					resource.TestCheckNoResourceAttr(resourceName+"_vm_cluster_no_db", "database"),
				),
			},

			// Create DB inside of dbHome
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_no_db", acctest.Optional, acctest.Update,
						acctest.RepresentationCopyWithNewProperties(dbHomeRepresentationSourceVmCluster, map[string]interface{}{
							"database": acctest.RepresentationGroup{RepType: acctest.Optional, Group: nesstedDatabaseRepresentationSourceNone2},
						})) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_db_vm_cluster_no_db", acctest.Required, acctest.Create, DatabaseDatabaseExacsRepresentation),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "source", "VM_CLUSTER_NEW"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "display_name", "TFTestDbHome1"),
					resource.TestCheckResourceAttrSet(resourceName+"_vm_cluster_no_db", "vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "database.#", "1"),

					// Added wait time to go into console to terminate the resource to catch the 404
					func(s *terraform.State) (err error) {
						time.Sleep(3 * time.Minute)
						return
					},
				),
			},

			// Delete with enabled_database_delete
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceVmClusterDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster_no_db", acctest.Optional, acctest.Update,
						acctest.RepresentationCopyWithNewProperties(dbHomeRepresentationSourceVmCluster, map[string]interface{}{
							"enable_database_delete": acctest.Representation{RepType: acctest.Required, Create: `true`},
						})) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_db_vm_cluster_no_db", acctest.Required, acctest.Create, DatabaseDatabaseExacsRepresentation),

				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "source", "VM_CLUSTER_NEW"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "display_name", "TFTestDbHome1"),
					resource.TestCheckResourceAttrSet(resourceName+"_vm_cluster_no_db", "vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName+"_vm_cluster_no_db", "db_version", "12.1.0.2"),
					resource.TestCheckNoResourceAttr(resourceName+"_vm_cluster_no_db", "database"),
				),
			},
		},
	})
}

func testAccCheckDatabaseDbHomeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_db_home" {
			noResourceFound = false
			request := oci_database.GetDbHomeRequest{}

			tmp := rs.Primary.ID
			request.DbHomeId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetDbHome(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DbHomeLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseDbHome") {
		resource.AddTestSweepers("DatabaseDbHome", &resource.Sweeper{
			Name:         "DatabaseDbHome",
			Dependencies: acctest.DependencyGraph["dbHome"],
			F:            sweepDatabaseDbHomeResource,
		})
	}
}

func sweepDatabaseDbHomeResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	dbHomeIds, err := getDatabaseDbHomeIds(compartment)
	if err != nil {
		return err
	}
	for _, dbHomeId := range dbHomeIds {
		if ok := acctest.SweeperDefaultResourceId[dbHomeId]; !ok {
			deleteDbHomeRequest := oci_database.DeleteDbHomeRequest{}

			deleteDbHomeRequest.DbHomeId = &dbHomeId

			deleteDbHomeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDbHome(context.Background(), deleteDbHomeRequest)
			if error != nil {
				fmt.Printf("Error deleting DbHome %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbHomeId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dbHomeId, DatabaseDbHomeSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseDbHomeSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseDbHomeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DbHomeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listDbHomesRequest := oci_database.ListDbHomesRequest{}
	listDbHomesRequest.CompartmentId = &compartmentId

	// Terminate the newest database first to make sure we delete any standby databases created by Data Guard Associations first
	listDbHomesRequest.SortBy = oci_database.ListDbHomesSortByTimecreated
	listDbHomesRequest.SortOrder = oci_database.ListDbHomesSortOrderDesc

	dbSystemIds, err := getDbSystemIds(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting dbSystemId required for DbHome resource requests \n")
	}
	for _, dbSystemId := range dbSystemIds {
		listDbHomesRequest.DbSystemId = &dbSystemId

		listDbHomesRequest.LifecycleState = oci_database.DbHomeSummaryLifecycleStateAvailable
		listDbHomesResponse, err := databaseClient.ListDbHomes(context.Background(), listDbHomesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DbHome list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, dbHome := range listDbHomesResponse.Items {
			id := *dbHome.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DbHomeId", id)
		}

	}
	listDbHomesRequest.DbSystemId = nil
	vmClusterIds, err := getDatabaseVmClusterIds(compartment)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting vmClusterId required for DbHome resource requests \n")
	}
	for _, vmClusterId := range vmClusterIds {
		listDbHomesRequest.VmClusterId = &vmClusterId

		listDbHomesRequest.LifecycleState = oci_database.DbHomeSummaryLifecycleStateAvailable
		listDbHomesResponse, err := databaseClient.ListDbHomes(context.Background(), listDbHomesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DbHome list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, dbHome := range listDbHomesResponse.Items {
			id := *dbHome.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DbHomeId", id)
		}

	}
	return resourceIds, nil
}

func DatabaseDbHomeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dbHomeResponse, ok := response.Response.(oci_database.GetDbHomeResponse); ok {
		return dbHomeResponse.LifecycleState != oci_database.DbHomeLifecycleStateTerminated
	}
	return false
}

func DatabaseDbHomeSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetDbHome(context.Background(), oci_database.GetDbHomeRequest{
		DbHomeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
