// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	databaseBackupRepresentation = map[string]interface{}{
		"database":   RepresentationGroup{Required, databaseDatabaseBackupRepresentation},
		"db_home_id": Representation{RepType: Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":     Representation{RepType: Required, Create: `DB_BACKUP`},
		"db_version": Representation{RepType: Optional, Create: `12.1.0.2`},
	}
	databaseDatabaseBackupRepresentation = map[string]interface{}{
		"admin_password":      Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"db_name":             Representation{RepType: Required, Create: `testDbBu`},
		"backup_id":           Representation{RepType: Required, Create: `${oci_database_backup.test_backup.id}`},
		"backup_tde_password": Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"character_set":       Representation{RepType: Optional, Create: `AL32UTF8`},
		"db_backup_config":    RepresentationGroup{Optional, databaseDatabaseDbBackupBackupConfigRepresentation},
		"db_unique_name":      Representation{RepType: Optional, Create: `testDbBu_12`},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       Representation{RepType: Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":      Representation{RepType: Optional, Create: `AL16UTF16`},
		"pdb_name":            Representation{RepType: Optional, Create: `pdbName`},
	}
	backupDatabaseRepresentation = map[string]interface{}{
		"database_id":  Representation{RepType: Required, Create: `${oci_database_database.db.id}`},
		"display_name": Representation{RepType: Required, Create: `Monthly Backup`},
	}
	databaseDatabaseDbBackupBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":     Representation{RepType: Optional, Create: `true`},
		"auto_backup_window":      Representation{RepType: Optional, Create: `SLOT_TWO`, Update: `SLOT_THREE`},
		"recovery_window_in_days": Representation{RepType: Optional, Create: `10`, Update: `30`},
	}
	databaseBackupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"db_home_id":     Representation{RepType: Optional, Create: `${oci_database_db_home.test_db_home.id}`},
		"db_name":        Representation{RepType: Optional, Create: `testDbBu`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, databaseBackupDataSourceFilterRepresentation}}
	databaseBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_database_database.test_database.id}`}},
	}
	databaseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"database_id": Representation{RepType: Required, Create: `${oci_database_database.test_database.id}`},
	}
	DatabaseBackupResourceConfig = DatabaseBackupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseBackupRepresentation)
	DatabaseBackupResourceDependencies = DatabaseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_database", "db", Optional, Create, databaseRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", Required, Create, backupDatabaseRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseDatabaseBackupResource_basic(t *testing.T) {
	// Skip the test because CreateDatabaseFromBackupDetails missing some parameters
	// https://confluence.oci.oraclecorp.com/display/TER/Support+ExaCS%3A+Create+DB+from+backup
	t.Skip("CreateDatabaseFromBackupDetails missing parameters")
	httpreplay.SetScenario("TestDatabaseDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_database.test_database"
	datasourceName := "data.oci_database_databases.test_databases"
	singularDatasourceName := "data.oci_database_database.test_database"

	var resId, resId2 string
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDatabaseDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DatabaseBackupResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseBackupRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Create, databaseBackupRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DatabaseBackupResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseBackupRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId2, err = FromInstanceState(s, resourceName, "id")
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
					GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_databases", Optional, Update, databaseBackupDataSourceRepresentation) +
					compartmentIdVariableStr +
					GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseBackupRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					GenerateDataSourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseBackupSingularDataSourceRepresentation) +
					GenerateResourceFromRepresentationMap("oci_database_database", "test_database", Optional, Update, databaseBackupRepresentation) +
					compartmentIdVariableStr,
				Check: ComposeAggregateTestCheckFuncWrapper(
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
