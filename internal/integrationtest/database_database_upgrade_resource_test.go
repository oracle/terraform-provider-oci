package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var (
	dbPrecheckFromDatabaseSoftwareImageRepresentation = map[string]interface{}{
		"action":                          acctest.Representation{RepType: acctest.Required, Create: `PRECHECK`},
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_db_system_for_upgrade.databases.0.id}`},
		"database_upgrade_source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbPrecheckFromDatabaseSoftwareImageGroup},
	}

	dbPrecheckFromDatabaseSoftwareImageGroup = map[string]interface{}{
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `DB_SOFTWARE_IMAGE`},
		"options":                    acctest.Representation{RepType: acctest.Optional, Create: `-upgradeTimezone false -keepEvents`},
		"database_software_image_id": acctest.Representation{RepType: acctest.Optional, Create: databaseSoftwareImageId},
	}

	dbUpgradeFromDatabaseSoftwareImageRepresentation = map[string]interface{}{
		"action":                          acctest.Representation{RepType: acctest.Required, Create: `UPGRADE`},
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_db_system_for_upgrade.databases.0.id}`},
		"database_upgrade_source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbUpgradeFromDatabaseSoftwareImageGroup},
	}

	dbUpgradeFromDatabaseSoftwareImageGroup = map[string]interface{}{
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `DB_SOFTWARE_IMAGE`},
		"options":                    acctest.Representation{RepType: acctest.Optional, Create: `-upgradeTimezone false -keepEvents`},
		"database_software_image_id": acctest.Representation{RepType: acctest.Optional, Create: databaseSoftwareImageId},
	}

	//Database Software Image with Database version - 19.24.0.0 and Shape Family - Virtual Machine Shape need to be pre-created on the tenancy
	databaseSoftwareImageId = utils.GetEnvSettingWithBlankDefault("database_software_image_id")

	dbPrecheckResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_database_upgrade", "test_database_precheck", acctest.Optional, acctest.Update, dbPrecheckFromDatabaseSoftwareImageRepresentation)
	dbUpgradeResourceConfig  = acctest.GenerateResourceFromRepresentationMap("oci_database_database_upgrade", "test_database_upgrade", acctest.Optional, acctest.Update, dbUpgradeFromDatabaseSoftwareImageRepresentation)
)

// TestDatabaseDatabaseUpgradeResource_basic tests Database using Virtual Machines.
// issue-routing-tag: database/default
func TestDatabaseDatabaseUpgradeResource_DbSoftwareImage(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "Database_upgrade") {
		t.Skip("Skipping suppressed upgrade_tests")
	}

	httpreplay.SetScenario("TestDatabaseDatabaseUpgradeResource_DbSoftwareImage")
	defer httpreplay.SaveScenario()

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: dbSystemBaseConfig + DbSystemForUpgradeConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.display_name", "tfDbHome"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_name", "tfDb"),

				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.test_db_system_for_upgrade", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.test_db_system_for_upgrade", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_home.test_db_system_for_upgrade", "db_version", "12.2.0.1"),

				// Databases
				resource.TestCheckResourceAttrSet("data.oci_database_databases.test_db_system_for_upgrade", "databases.#"),
				resource.TestCheckResourceAttr("data.oci_database_databases.test_db_system_for_upgrade", "databases.#", "1"),
				resource.TestCheckResourceAttr("data.oci_database_databases.test_db_system_for_upgrade", "databases.0.character_set", "AL32UTF8"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, "oci_database_db_system.test_db_system_for_upgrade", "id")
					return err
				},
			),
		},
		// verify PRECHECK action on database with source=DB_SOFTWARE_IMAGE
		{
			Config: dbSystemBaseConfig + dbPrecheckResourceConfig + DbSystemForUpgradeConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.test_db_system_for_upgrade", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.test_db_system_for_upgrade", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_home.test_db_system_for_upgrade", "db_version", "12.2.0.1"),

				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "database_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_db_system_for_upgrade", "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "compartment_id"),
			),
		},
		// verify upgrade history entries singular and plural datasources after PRECHECK action on database
		{
			Config: dbSystemBaseConfig + dbPrecheckResourceConfig + DbSystemForUpgradeConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_upgrade_history_entries", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbUpgradeHistoryEntriesDatasourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_upgrade_history_entry", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbUpgradeHistoryFirstEntryDatasourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.test_db_system_for_upgrade", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.test_db_system_for_upgrade", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_home.test_db_system_for_upgrade", "db_version", "12.2.0.1"),

				// Databases
				resource.TestCheckResourceAttrSet("data.oci_database_databases.test_db_system_for_upgrade", "databases.#"),
				resource.TestCheckResourceAttr("data.oci_database_databases.test_db_system_for_upgrade", "databases.#", "1"),
				resource.TestCheckResourceAttr("data.oci_database_databases.test_db_system_for_upgrade", "databases.0.character_set", "AL32UTF8"),

				//Upgrade history entry - plural datasource
				resource.TestCheckResourceAttrSet("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.#"),
				resource.TestCheckResourceAttrSet("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.0.id"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.0.action", "PRECHECK"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.0.state", "SUCCEEDED"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.0.source", "DB_SOFTWARE_IMAGE"),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "action", "PRECHECK"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "source", "DB_SOFTWARE_IMAGE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "state", "SUCCEEDED"),
			),
		},
		// verify UPGRADE action on database with source=DB_SOFTWARE_IMAGE
		{
			Config: dbSystemBaseConfig + dbUpgradeResourceConfig + DbSystemForUpgradeConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "database_id"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.test_db_system_for_upgrade", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		// verify upgrade history entries singular and plural datasources after UPGRADE action on database
		{
			Config: dbSystemBaseConfig + dbUpgradeResourceConfig + DbSystemForUpgradeConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_upgrade_history_entries", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbUpgradeHistoryEntriesDatasourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_upgrade_history_entry", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbUpgradeHistorySecondEntryDatasourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "database_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_db_system_for_upgrade", "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "compartment_id"),

				//Upgrade history entry - plural datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.#", "2"),
				resource.TestCheckResourceAttrSet("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.id"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.state", "SUCCEEDED"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.source", "DB_SOFTWARE_IMAGE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.options", "-upgradeTimezone false -keepEvents"),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "source", "DB_SOFTWARE_IMAGE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "state", "SUCCEEDED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.test_db_system_for_upgrade", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
	})
}
