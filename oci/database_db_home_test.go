// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v51/common"
	oci_database "github.com/oracle/oci-go-sdk/v51/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DbHomeRequiredOnlyResource = DbHomeResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Required, Create, dbHomeRepresentationSourceNone)

	DbHomeResourceConfig = DbHomeResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Optional, Update, dbHomeRepresentationSourceNone)

	dbHomeSingularDataSourceRepresentation = map[string]interface{}{
		"db_home_id": Representation{RepType: Required, Create: `${oci_database_db_home.test_db_home_source_none.id}`},
	}

	dbHomeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"db_system_id":   Representation{RepType: Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"display_name":   Representation{RepType: Optional, Create: `createdDbHomeNone`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, dbHomeDataSourceFilterRepresentation}}

	dbHomeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_db_home.test_db_home_source_none.id}`}},
	}

	dbHomeRepresentation = map[string]interface{}{
		"database":                   RepresentationGroup{Required, dbHomeDatabaseRepresentation},
		"database_software_image_id": Representation{RepType: Optional, Create: `${oci_database_database_software_image.test_database_software_image.id}`},
		"db_system_id":               Representation{RepType: Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"db_version":                 Representation{RepType: Required, Create: `12.1.0.2`},
		"defined_tags":               Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":               Representation{RepType: Optional, Create: `createdDbHome`},
		"freeform_tags":              Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_desupported_version":     Representation{RepType: Optional, Create: `false`},
		"kms_key_id":                 Representation{RepType: Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_version_id":         Representation{RepType: Optional, Create: `${oci_kms_key_version.test_key_version.id}`},
		"source":                     Representation{RepType: Optional, Create: `DB_BACKUP`},
		"vm_cluster_id":              Representation{RepType: Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	dbHomeDatabaseRepresentation = map[string]interface{}{
		"admin_password": Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"db_name":        Representation{RepType: Required, Create: `tfDbNam`},
	}
	dbHomeRepresentationBase = map[string]interface{}{
		"db_system_id": Representation{RepType: Required, Create: `${oci_database_db_system.test_db_system.id}`},
	}
	dbHomeRepresentationSourceNone = RepresentationCopyWithNewProperties(dbHomeRepresentationBase, map[string]interface{}{
		"database":     RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceNone},
		"db_version":   Representation{RepType: Required, Create: `12.1.0.2`},
		"source":       Representation{RepType: Optional, Create: `NONE`},
		"display_name": Representation{RepType: Optional, Create: `createdDbHomeNone`},
	})
	dbHomeDatabaseRepresentationSourceNone = map[string]interface{}{
		"admin_password":      Representation{RepType: Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"tde_wallet_password": Representation{RepType: Optional, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"db_name":             Representation{RepType: Required, Create: `dbNone`},
		"character_set":       Representation{RepType: Optional, Create: `AL32UTF8`},
		"db_backup_config":    RepresentationGroup{Optional, dbHomeDatabaseDbBackupConfigRepresentation},
		"db_workload":         Representation{RepType: Optional, Create: `OLTP`},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":      Representation{RepType: Optional, Create: `AL16UTF16`},
		"pdb_name":            Representation{RepType: Optional, Create: `pdbName`},
	}
	dbHomeRepresentationSourceNoneRequiredOnly = RepresentationCopyWithNewProperties(dbHomeRepresentationSourceNone, map[string]interface{}{
		"database": RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceNoneRequiredOnly},
	})
	dbHomeDatabaseRepresentationSourceNoneRequiredOnly = RepresentationCopyWithNewProperties(dbHomeDatabaseRepresentationSourceNone, map[string]interface{}{
		"db_name": Representation{RepType: Required, Create: `dbNone0`},
	})
	dbHomeDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":     Representation{RepType: Optional, Create: `true`},
		"auto_backup_window":      Representation{RepType: Optional, Create: `SLOT_TWO`},
		"recovery_window_in_days": Representation{RepType: Optional, Create: `10`},
	}
	dbHomeRepresentationSourceDbBackup = RepresentationCopyWithNewProperties(dbHomeRepresentationBase, map[string]interface{}{
		"database":     RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceDbBackup},
		"source":       Representation{RepType: Required, Create: `DB_BACKUP`},
		"display_name": Representation{RepType: Required, Create: `createdDbHomeBackup`},
	})
	dbHomeDatabaseRepresentationSourceDbBackup = map[string]interface{}{
		"admin_password":      Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"backup_id":           Representation{RepType: Required, Create: `${oci_database_backup.test_backup.id}`},
		"backup_tde_password": Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		// Modifying db_name as mandatory. If not mandatory test fails with error "The specified database name 'tfDbName' exists."
		// The test takes the backup of the DB created in the db_system which has the db_name=tfDbName.
		// When db_home is created with source as "DB_BACKUP" and db_name is not provided, Service uses the db_name from the backup which is causing this test to fail.
		"db_name": Representation{RepType: Required, Create: `dbBackup`},
	}

	dbHomeRepresentationSourceVmClusterNew = map[string]interface{}{
		"database":      RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceVmClusterNew},
		"display_name":  Representation{RepType: Optional, Create: `createdDbHomeVm`},
		"source":        Representation{RepType: Required, Create: `VM_CLUSTER_NEW`},
		"db_version":    Representation{RepType: Required, Create: `12.1.0.2`},
		"vm_cluster_id": Representation{RepType: Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"defined_tags":  Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	dbHomeDatabaseRepresentationSourceVmClusterNew = map[string]interface{}{
		"admin_password":   Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"character_set":    Representation{RepType: Optional, Create: `AL32UTF8`},
		"db_backup_config": RepresentationGroup{Optional, dbHomeDatabaseDbBackupConfigVmClusterNewRepresentation},
		"db_name":          Representation{RepType: Required, Create: `dbVMClus`},
		"db_workload":      Representation{RepType: Optional, Create: `OLTP`},
		"defined_tags":     Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   Representation{RepType: Optional, Create: `AL16UTF16`},
		"pdb_name":         Representation{RepType: Optional, Create: `pdbName`},
	}

	dbHomeDatabaseDbBackupConfigVmClusterNewRepresentation = map[string]interface{}{
		"auto_backup_enabled":        Representation{RepType: Optional, Create: `true`, Update: `false`},
		"auto_backup_window":         Representation{RepType: Optional, Create: `SLOT_TWO`},
		"backup_destination_details": RepresentationGroup{Optional, dbHomeDatabaseDbBackupConfigBackupDestinationDetails2Representation},
		"recovery_window_in_days":    Representation{RepType: Optional, Create: `10`},
	}

	dbHomeDatabaseDbBackupConfigBackupDestinationDetails2Representation = map[string]interface{}{
		"id":   Representation{RepType: Optional, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
		"type": Representation{RepType: Required, Create: `NFS`},
	}
	dbHomeRepresentationSourceDatabase = RepresentationCopyWithNewProperties(dbHomeRepresentationBase, map[string]interface{}{
		"database":     RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceDatabase},
		"source":       Representation{RepType: Required, Create: `DATABASE`},
		"display_name": Representation{RepType: Optional, Create: `createdDbHomeDatabase`},
	})
	dbHomeDatabaseRepresentationSourceDatabase = map[string]interface{}{
		"admin_password":      Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"backup_tde_password": Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"database_id":         Representation{RepType: Required, Create: `${data.oci_database_databases.db.databases.0.id}`},
		"db_name":             Representation{RepType: Required, Create: `dbDb`},
	}

	dbHomeRepresentationSourceVmClusterDatabase = map[string]interface{}{
		"database":      RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceVmClusterDatabase},
		"display_name":  Representation{RepType: Optional, Create: `createdDbHomeVmClusterDatabase`},
		"source":        Representation{RepType: Required, Create: `VM_CLUSTER_DATABASE`},
		"db_version":    Representation{RepType: Required, Create: `12.1.0.2`},
		"vm_cluster_id": Representation{RepType: Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	dbHomeDatabaseRepresentationSourceVmClusterDatabase = map[string]interface{}{
		"admin_password": Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"character_set":  Representation{RepType: Optional, Create: `AL32UTF8`},
		//"db_backup_config": RepresentationGroup{Optional, dbHomeDatabaseDbBackupConfigVmClusterDatabaseRepresentation},
		"db_name":        Representation{RepType: Required, Create: `dbVMClusDb`},
		"db_workload":    Representation{RepType: Optional, Create: `OLTP`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set": Representation{RepType: Optional, Create: `AL16UTF16`},
		"pdb_name":       Representation{RepType: Optional, Create: `pdbName`},
	}

	DbHomeResourceDependencies = BackupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Create, backupDestinationNFSRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{RepType: Optional, Update: activationFilePath}})) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", Required, Create, backupRepresentation) +
		KeyResourceDependencyConfig +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseDbHomeTdeWalletPassword(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbHomeTdeWalletPassword")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_home.test_db_home"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDbHomeDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DbSystemResourceConfig +
					GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Required, Create, dbHomeRepresentationSourceNoneRequiredOnly),

				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone0"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestMatchResourceAttr(resourceName+"_source_none", "db_version", regexp.MustCompile(`^12\.1\.0\.2\.[0-9]+$`)),
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + DbSystemResourceConfig,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + DbSystemResourceConfig +
					GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Create, dbHomeRepresentationSourceNone),

				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.tde_wallet_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_window", "SLOT_TWO"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestMatchResourceAttr(resourceName+"_source_none", "db_version", regexp.MustCompile(`^12\.1\.0\.2(\.[0-9]+)?$`)),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHomeNone"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "source", "NONE"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "state"),
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DbSystemResourceConfig +
					GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.tde_wallet_password", "BEstrO0ng_#12"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestMatchResourceAttr(resourceName+"_source_none", "db_version", regexp.MustCompile(`^12\.1\.0\.2(\.[0-9]+)?$`)),
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

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_home.test_db_home"
	datasourceName := "data.oci_database_db_homes.test_db_homes"
	singularDatasourceName := "data.oci_database_db_home.test_db_home"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DbHomeResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Optional, Create, dbHomeRepresentation), "database", "dbHome", t)

	ResourceTest(t, testAccCheckDatabaseDbHomeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbHomeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Required, Create, dbHomeRepresentationSourceNoneRequiredOnly) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Required, Create, dbHomeRepresentationSourceDbBackup) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_vm_cluster_new", Required, Create, dbHomeRepresentationSourceVmClusterNew) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", Required, Create, dbHomeRepresentationSourceDatabase),

			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone0"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
				resource.TestMatchResourceAttr(resourceName+"_source_none", "db_version", regexp.MustCompile(`^12\.1\.0\.2\.[0-9]+$`)),

				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "database.0.backup_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "source", "DB_BACKUP"),

				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "vm_cluster_id"),
				resource.TestMatchResourceAttr(resourceName+"_source_vm_cluster_new", "db_version", regexp.MustCompile(`^12\.1\.0\.2\.[0-9]+$`)),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "source", "VM_CLUSTER_NEW"),

				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "database.0.database_id"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "source", "DATABASE"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbHomeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbHomeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Create, dbHomeRepresentationSourceNone) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Create, dbHomeRepresentationSourceDbBackup) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_vm_cluster_new", Optional, Create, dbHomeRepresentationSourceVmClusterNew) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", Optional, Create, dbHomeRepresentationSourceDatabase),

			Check: ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(resourceName, "is_desupported_version", "false"),
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
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
				resource.TestMatchResourceAttr(resourceName+"_source_none", "db_version", regexp.MustCompile(`^12\.1\.0\.2(\.[0-9]+)?$`)),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "database.0.backup_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.db_name", "dbBackup"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "db_version", "12.1.0.2"),
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
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "vm_cluster_id"),
				resource.TestMatchResourceAttr(resourceName+"_source_vm_cluster_new", "db_version", regexp.MustCompile(`^12\.1\.0\.2(\.[0-9]+)?$`)),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "display_name", "createdDbHomeVm"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "source", "VM_CLUSTER_NEW"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "database.0.database_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.db_name", "dbDb"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "display_name", "createdDbHomeDatabase"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "state"),
				resource.TestCheckResourceAttr(resourceName, "is_desupported_version", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DbHomeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Update, dbHomeRepresentationSourceDbBackup) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_vm_cluster_new", Optional, Update, dbHomeRepresentationSourceVmClusterNew) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", Optional, Update, dbHomeRepresentationSourceDatabase),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_desupported_version", "false"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
				resource.TestMatchResourceAttr(resourceName+"_source_none", "db_version", regexp.MustCompile(`^12\.1\.0\.2(\.[0-9]+)?$`)),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "kms_key_version_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_none", "source", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_none", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "database.0.backup_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.db_name", "dbBackup"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "db_version", "12.1.0.2"),
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
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.ncharacter_set", "AL16UTF16"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "database.0.pdb_name", "pdbName"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "vm_cluster_id"),
				resource.TestMatchResourceAttr(resourceName+"_source_vm_cluster_new", "db_version", regexp.MustCompile(`^12\.1\.0\.2(\.[0-9]+)?$`)),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "display_name", "createdDbHomeVm"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_vm_cluster_new", "source", "VM_CLUSTER_NEW"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_vm_cluster_new", "state"),

				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "compartment_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.backup_tde_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "database.0.database_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "database.0.db_name", "dbDb"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "db_system_id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "db_version", "12.1.0.2"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "display_name", "createdDbHomeDatabase"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "id"),
				resource.TestCheckResourceAttr(resourceName+"_source_database", "source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName+"_source_database", "state"),
				resource.TestCheckResourceAttr(resourceName, "is_desupported_version", "true"),
			),
		},

		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "test_db_homes", Optional, Update, dbHomeDataSourceRepresentation) +
				compartmentIdVariableStr + DbHomeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Update, dbHomeRepresentationSourceDbBackup) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", Optional, Update, dbHomeRepresentationSourceDatabase),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "db_homes.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.db_system_id"),
				resource.TestMatchResourceAttr(datasourceName, "db_homes.0.db_version", regexp.MustCompile(`^12\.1\.0\.2(\.[0-9]+)?$`)),
				resource.TestCheckResourceAttr(datasourceName, "db_homes.0.display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.kms_key_id"),
				//resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.kms_key_version_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_db_home", "test_db_home", Required, Create, dbHomeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbHomeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Update, dbHomeRepresentationSourceDbBackup) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_database", Optional, Update, dbHomeRepresentationSourceDatabase),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestMatchResourceAttr(singularDatasourceName, "db_version", regexp.MustCompile(`^12\.1\.0\.2(\.[0-9]+)?$`)),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "createdDbHomeNone"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config +
				compartmentIdVariableStr + DbHomeResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone) +
				GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Update, dbHomeRepresentationSourceDbBackup),
		},
		// verify resource import
		{
			Config:            config,
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

func testAccCheckDatabaseDbHomeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_db_home" {
			noResourceFound = false
			request := oci_database.GetDbHomeRequest{}

			tmp := rs.Primary.ID
			request.DbHomeId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DatabaseDbHome") {
		resource.AddTestSweepers("DatabaseDbHome", &resource.Sweeper{
			Name:         "DatabaseDbHome",
			Dependencies: DependencyGraph["dbHome"],
			F:            sweepDatabaseDbHomeResource,
		})
	}
}

func sweepDatabaseDbHomeResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	dbHomeIds, err := getDbHomeIds(compartment)
	if err != nil {
		return err
	}
	for _, dbHomeId := range dbHomeIds {
		if ok := SweeperDefaultResourceId[dbHomeId]; !ok {
			deleteDbHomeRequest := oci_database.DeleteDbHomeRequest{}

			deleteDbHomeRequest.DbHomeId = &dbHomeId

			deleteDbHomeRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDbHome(context.Background(), deleteDbHomeRequest)
			if error != nil {
				fmt.Printf("Error deleting DbHome %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbHomeId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &dbHomeId, dbHomeSweepWaitCondition, time.Duration(3*time.Minute),
				dbHomeSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDbHomeIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DbHomeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

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
			AddResourceIdToSweeperResourceIdMap(compartmentId, "DbHomeId", id)
		}

	}
	listDbHomesRequest.DbSystemId = nil
	vmClusterIds, err := getVmClusterIds(compartment)

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
			AddResourceIdToSweeperResourceIdMap(compartmentId, "DbHomeId", id)
		}

	}
	return resourceIds, nil
}

func dbHomeSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dbHomeResponse, ok := response.Response.(oci_database.GetDbHomeResponse); ok {
		return dbHomeResponse.LifecycleState != oci_database.DbHomeLifecycleStateTerminated
	}
	return false
}

func dbHomeSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetDbHome(context.Background(), oci_database.GetDbHomeRequest{
		DbHomeId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
