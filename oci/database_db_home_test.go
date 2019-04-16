// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DbHomeRequiredOnlyResource = DbHomeResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Required, Create, dbHomeRepresentationSourceNone)

	DbHomeResourceConfig = DbHomeResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Optional, Update, dbHomeRepresentationSourceNone)

	dbHomeSingularDataSourceRepresentation = map[string]interface{}{
		"db_home_id": Representation{repType: Required, create: `${oci_database_db_home.test_db_home_source_none.id}`},
	}

	dbHomeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"db_system_id":   Representation{repType: Required, create: `${oci_database_db_system.test_db_system.id}`},
		"display_name":   Representation{repType: Optional, create: `createdDbHome`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"filter":         RepresentationGroup{Required, dbHomeDataSourceFilterRepresentation}}
	dbHomeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_db_home.test_db_home_source_none.id}`}},
	}

	dbHomeRepresentationBase = map[string]interface{}{
		"db_system_id": Representation{repType: Required, create: `${oci_database_db_system.test_db_system.id}`},
		"display_name": Representation{repType: Optional, create: `createdDbHome`},
	}
	dbHomeRepresentationSourceNone = representationCopyWithNewProperties(dbHomeRepresentationBase, map[string]interface{}{
		"database":   RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceNone},
		"db_version": Representation{repType: Required, create: `12.1.0.2`},
		"source":     Representation{repType: Optional, create: `NONE`},
	})
	dbHomeDatabaseRepresentationSourceNone = map[string]interface{}{
		"admin_password":   Representation{repType: Required, create: `BEstrO0ng_#11`},
		"db_name":          Representation{repType: Required, create: `dbNone`},
		"character_set":    Representation{repType: Optional, create: `AL32UTF8`},
		"db_backup_config": RepresentationGroup{Optional, dbHomeDatabaseDbBackupConfigRepresentation},
		"db_workload":      Representation{repType: Optional, create: `OLTP`},
		"defined_tags":     Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    Representation{repType: Optional, create: map[string]string{"freeformTags": "freeformTags"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
		"ncharacter_set":   Representation{repType: Optional, create: `AL16UTF16`},
		"pdb_name":         Representation{repType: Optional, create: `pdbName`},
	}
	dbHomeRepresentationSourceNoneRequiredOnly = representationCopyWithNewProperties(dbHomeRepresentationSourceNone, map[string]interface{}{
		"database": RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceNoneRequiredOnly},
	})
	dbHomeDatabaseRepresentationSourceNoneRequiredOnly = representationCopyWithNewProperties(dbHomeDatabaseRepresentationSourceNone, map[string]interface{}{
		"db_name": Representation{repType: Required, create: `dbNone0`},
	})
	dbHomeDatabaseDbBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled": Representation{repType: Optional, create: `true`, update: `false`},
	}
	dbHomeRepresentationSourceDbBackup = representationCopyWithNewProperties(dbHomeRepresentationBase, map[string]interface{}{
		"database": RepresentationGroup{Required, dbHomeDatabaseRepresentationSourceDbBackup},
		"source":   Representation{repType: Required, create: `DB_BACKUP`},
	})
	dbHomeDatabaseRepresentationSourceDbBackup = map[string]interface{}{
		"admin_password":      Representation{repType: Required, create: `BEstrO0ng_#11`},
		"backup_id":           Representation{repType: Required, create: `${oci_database_backup.test_backup.id}`},
		"backup_tde_password": Representation{repType: Required, create: `BEstrO0ng_#11`},
		"db_name":             Representation{repType: Optional, create: `dbBackup`},
	}

	DbHomeResourceDependencies = BackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_backup", "test_backup", Required, Create, backupRepresentation)
)

func TestDatabaseDbHomeResource_basic(t *testing.T) {
	if httpreplay.ShouldRetryImmediately() {
		t.Skip("TestDatabaseDbHomeResource_basic test is flaky, tracked in TER-1274, skip this test in checkin test.")
	}

	httpreplay.SetScenario("TestDatabaseDbHomeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_home.test_db_home"
	datasourceName := "data.oci_database_db_homes.test_db_homes"
	singularDatasourceName := "data.oci_database_db_home.test_db_home"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDbHomeDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Required, Create, dbHomeRepresentationSourceNoneRequiredOnly) +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Required, Create, dbHomeRepresentationSourceDbBackup),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone0"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),

					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "database.0.backup_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "database.0.backup_tde_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "db_system_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "source", "DB_BACKUP"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Create, dbHomeRepresentationSourceNone) +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Create, dbHomeRepresentationSourceDbBackup),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHome"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
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
					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "display_name", "createdDbHome"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "id"),
					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "source", "DB_BACKUP"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "state"),
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DbHomeResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone) +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Update, dbHomeRepresentationSourceDbBackup),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "compartment_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.character_set", "AL32UTF8"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_backup_config.0.auto_backup_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_name", "dbNone"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.ncharacter_set", "AL16UTF16"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "database.0.pdb_name", "pdbName"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "db_system_id"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(resourceName+"_source_none", "display_name", "createdDbHome"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_none", "id"),
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
					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "display_name", "createdDbHome"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "id"),
					resource.TestCheckResourceAttr(resourceName+"_source_db_backup", "source", "DB_BACKUP"),
					resource.TestCheckResourceAttrSet(resourceName+"_source_db_backup", "state"),
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_homes", "test_db_homes", Optional, Update, dbHomeDataSourceRepresentation) +
					compartmentIdVariableStr + DbHomeResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone) +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Update, dbHomeRepresentationSourceDbBackup),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "createdDbHome"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "db_homes.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.db_system_id"),
					resource.TestCheckResourceAttr(datasourceName, "db_homes.0.db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(datasourceName, "db_homes.0.display_name", "createdDbHome"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_homes.0.state"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_home", "test_db_home", Required, Create, dbHomeSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DbHomeResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone) +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Update, dbHomeRepresentationSourceDbBackup),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_home_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "createdDbHome"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config +
					compartmentIdVariableStr + DbHomeResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_none", Optional, Update, dbHomeRepresentationSourceNone) +
					generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_source_db_backup", Optional, Update, dbHomeRepresentationSourceDbBackup),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"database.0.admin_password",
				},
				ResourceName: resourceName + "_source_none",
			},
		},
	})
}

func testAccCheckDatabaseDbHomeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_db_home" {
			noResourceFound = false
			request := oci_database.GetDbHomeRequest{}

			tmp := rs.Primary.ID
			request.DbHomeId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

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
