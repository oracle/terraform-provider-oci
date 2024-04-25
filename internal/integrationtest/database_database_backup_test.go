// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	databaseBackupRepresentation = map[string]interface{}{
		"database":   acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseBackupRepresentation},
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `DB_BACKUP`},
		"db_version": acctest.Representation{RepType: acctest.Optional, Create: `12.1.0.2`},
	}
	databaseDatabaseBackupRepresentation = map[string]interface{}{
		"admin_password":      acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":             acctest.Representation{RepType: acctest.Required, Create: `testDbBu`},
		"backup_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_database_backup.test_backup.id}`},
		"backup_tde_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"character_set":       acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_backup_config":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseDatabaseDbBackupBackupConfigRepresentation},
		"db_unique_name":      acctest.Representation{RepType: acctest.Optional, Create: `testDbBu_12`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":      acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":            acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
	}
	backupDatabaseRepresentation = map[string]interface{}{
		"database_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database.db.id}`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `Monthly Backup`},
	}
	databaseDatabaseDbBackupBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"auto_backup_window":      acctest.Representation{RepType: acctest.Optional, Create: `SLOT_TWO`, Update: `SLOT_THREE`},
		"recovery_window_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `30`},
	}
	databaseBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_home_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_home.test_db_home.id}`},
		"db_name":        acctest.Representation{RepType: acctest.Optional, Create: `testDbBu`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseBackupDataSourceFilterRepresentation}}
	databaseBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_database.test_database.id}`}},
	}
	databaseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database.test_database.id}`},
	}

	databaseDisabledBackupConf = map[string]interface{}{
		"auto_backup_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	databaseDatabaseBackupRepresentationCopy = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `DbBackup`},
		"character_set":    acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseDisabledBackupConf},
		"db_unique_name":   acctest.Representation{RepType: acctest.Optional, Create: `myTestDb_44`},
		"db_workload":      acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":         acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
		// "tde_wallet_password": acctest.Representation{RepType: acctest.Optional, Create: `tdeWalletPassword`},	exadata doesn't support it.
	}

	DatabaseDatabaseBackupRepresentation = map[string]interface{}{
		"database":         acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseBackupRepresentationCopy},
		"db_home_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":           acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"db_version":       acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"kms_key_id":       acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"kms_key_rotation": acctest.Representation{RepType: acctest.Optional, Update: `1`},
	}

	dbHomeDatabaseBackupRepresentationSourceNone = map[string]interface{}{
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		// "tde_wallet_password": acctest.Representation{RepType: acctest.Optional, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},  exadata doesn't support it.
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `dbNoneBT`},
		"character_set":    acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbHomeDatabaseDbBackupConfigRepresentation},
		"db_workload":      acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"pdb_name":         acctest.Representation{RepType: acctest.Optional, Create: `pdbName`},
	}

	dbHomeRepresentationSourceBackup = acctest.RepresentationCopyWithNewProperties(DatabaseDbHomeRepresentationBase, map[string]interface{}{
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeDatabaseBackupRepresentationSourceNone},
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"source":       acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeNone`},
	})

	DatabaseBackupResourceConfig = DatabaseBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, databaseBackupRepresentation)
	DatabaseBackupResourceDependencies = DatabaseDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "db", acctest.Optional, acctest.Create, DatabaseDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, backupDatabaseRepresentation)

	DatabaseBackupResourceDbHomeDependencies = ExaBaseDependencies + DefinedTagsDependencies + AvailabilityDomainConfig + KeyResourceDependencyConfig2 +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentationSourceBackup) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "db", acctest.Optional, acctest.Create, DatabaseDatabaseBackupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, backupDatabaseRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseDatabaseBackupResource_basic(t *testing.T) {
	// Skip the test because CreateDatabaseFromBackupDetails missing some parameters
	// https://confluence.oci.oraclecorp.com/display/TER/Support+ExaCS%3A+Create+DB+from+backup
	t.Skip("CreateDatabaseFromBackupDetails missing parameters")
	httpreplay.SetScenario("TestDatabaseDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_database.test_database"
	datasourceName := "data.oci_database_databases.test_databases"
	singularDatasourceName := "data.oci_database_database.test_database"

	var resId, resId2 string
	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DatabaseBackupResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, databaseBackupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "database.0.backup_id"),
					resource.TestCheckResourceAttr(resourceName, "database.0.backup_tde_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "testDbBu"),
					resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
					resource.TestCheckResourceAttr(resourceName, "source", "DB_BACKUP"),
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + DatabaseBackupResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + DatabaseBackupResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, databaseBackupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName, "db_name", "testDbBu"),
					resource.TestCheckResourceAttr(resourceName, "db_unique_name", "testDbBu_12"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
					resource.TestCheckResourceAttrSet(resourceName, "db_name"),
					resource.TestCheckResourceAttrSet(resourceName, "db_unique_name"),
					resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source", "DB_BACKUP"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_enabled", "true"),
					//resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_THREE"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "10"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DatabaseBackupResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, databaseBackupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName, "character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_enabled", "true"),
					//resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.auto_backup_window", "SLOT_THREE"),
					resource.TestCheckResourceAttr(resourceName, "db_backup_config.0.recovery_window_in_days", "30"),
					resource.TestCheckResourceAttr(resourceName, "db_name", "testDbBu"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName, "pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
					resource.TestCheckResourceAttrSet(resourceName, "db_name"),
					resource.TestCheckResourceAttrSet(resourceName, "db_unique_name"),
					resource.TestCheckResourceAttr(resourceName, "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "source", "DB_BACKUP"),
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
				Config: config + DatabaseBackupResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_databases", acctest.Optional, acctest.Update, databaseBackupDataSourceRepresentation) +
					compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, databaseBackupRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "db_home_id"),
					resource.TestCheckResourceAttr(datasourceName, "db_name", "testDbBu"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.character_set"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.compartment_id"),
					resource.TestCheckResourceAttr(datasourceName, "databases.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_home_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_system_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_unique_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.db_workload"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.ncharacter_set"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.pdb_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "databases.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config + DatabaseBackupResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, databaseBackupSingularDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, databaseBackupRepresentation) +
					compartmentIdVariableStr,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "character_set"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_backup_config.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_unique_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_workload"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ncharacter_set"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "pdb_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}
