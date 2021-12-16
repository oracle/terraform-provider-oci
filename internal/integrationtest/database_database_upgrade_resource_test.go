package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	DatabasePrecheckResourceRep = acctest.GenerateResourceFromRepresentationMap("oci_database_database_upgrade", "test_database_upgrade", acctest.Optional, acctest.Update, databasePrecheckRep)
	DatabaseUpgradeResourceRep  = acctest.GenerateResourceFromRepresentationMap("oci_database_database_upgrade", "test_database_upgrade", acctest.Optional, acctest.Update, databaseUpgradeRep)

	//Database Software Image with Database version - 19.8.0.0 and Shape Family - Virtual Machine and Bare Metal Shapes needs to be pre-created on the tenancy
	databaseSoftwareImageId = utils.GetEnvSettingWithBlankDefault("database_software_image_id")

	databasePrecheckRep = map[string]interface{}{
		"action":                          acctest.Representation{RepType: acctest.Required, Create: `PRECHECK`},
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.id}`},
		"database_upgrade_source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databasePrecheckDatabaseUpgradeSourceDbSoftwareImageRep},
	}

	databasePrecheckDatabaseUpgradeSourceDbSoftwareImageRep = map[string]interface{}{
		"database_software_image_id": acctest.Representation{RepType: acctest.Optional, Create: databaseSoftwareImageId},
		"options":                    acctest.Representation{RepType: acctest.Optional, Create: `-upgradeTimezone false -keepEvents`},
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `DB_SOFTWARE_IMAGE`},
	}

	databaseUpgradeRep = map[string]interface{}{
		"action":                          acctest.Representation{RepType: acctest.Required, Create: `UPGRADE`},
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.id}`},
		"database_upgrade_source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseUpgradeDatabaseUpgradeSourceDbSoftwareImageRep},
	}

	databaseUpgradeDatabaseUpgradeSourceDbSoftwareImageRep = map[string]interface{}{
		"database_software_image_id": acctest.Representation{RepType: acctest.Optional, Create: databaseSoftwareImageId},
		"options":                    acctest.Representation{RepType: acctest.Optional, Create: `-upgradeTimezone false -keepEvents`},
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `DB_SOFTWARE_IMAGE`},
	}
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
			Config: ResourceDatabaseBaseConfig + dbSystemForDbUpgradeRepresentation,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.display_name", "-tf-db-home"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),

				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_home.t", "db_version", "12.2.0.1"),

				// Databases
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.#"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.character_set", "AL32UTF8"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					return err
				},
			),
		},
		// verify PRECHECK action on database with source=DB_SOFTWARE_IMAGE
		{
			Config: ResourceDatabaseBaseConfig + DatabasePrecheckResourceRep + dbSystemForDbUpgradeRepresentation,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_home.t", "db_version", "12.2.0.1"),

				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "database_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "compartment_id"),
			),
		},
		// verify upgrade history entries singular and plural datasources after PRECHECK action on database
		{
			Config: ResourceDatabaseBaseConfig + DatabasePrecheckResourceRep + dbSystemForDbUpgradeRepresentation + ResourceDatabaseTokenFn(`
				data "oci_database_database_upgrade_history_entries" "t" {
					database_id = "${data.oci_database_databases.t.databases.0.id}"
				}
				data "oci_database_database_upgrade_history_entry" "t" {
					database_id = "${data.oci_database_databases.t.databases.0.id}"
					upgrade_history_entry_id = "${data.oci_database_database_upgrade_history_entries.t.database_upgrade_history_entries.0.id}"
				}
				`, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_home.t", "db_version", "12.2.0.1"),

				// Databases
				resource.TestCheckResourceAttrSet("data.oci_database_databases.t", "databases.#"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.#", "1"),
				resource.TestCheckResourceAttr("data.oci_database_databases.t", "databases.0.character_set", "AL32UTF8"),

				//Upgrade history entry - plural datasource
				resource.TestCheckResourceAttrSet("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.#"),
				resource.TestCheckResourceAttrSet("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.0.id"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.0.action", "PRECHECK"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.0.state", "SUCCEEDED"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.0.source", "DB_SOFTWARE_IMAGE"),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "action", "PRECHECK"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "source", "DB_SOFTWARE_IMAGE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "state", "SUCCEEDED"),
			),
		},
		// verify UPGRADE action on database with source=DB_SOFTWARE_IMAGE
		{
			Config: ResourceDatabaseBaseConfig + DatabaseUpgradeResourceRep + dbSystemForDbUpgradeRepresentation,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "database_id"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
		// verify upgrade history entries singular and plural datasources after UPGRADE action on database
		{
			Config: ResourceDatabaseBaseConfig + DatabaseUpgradeResourceRep + dbSystemForDbUpgradeRepresentation + ResourceDatabaseTokenFn(`
				data "oci_database_database_upgrade_history_entries" "t" {
					database_id = "${data.oci_database_databases.t.databases.0.id}"
				}
				data "oci_database_database_upgrade_history_entry" "t" {
					database_id = "${data.oci_database_databases.t.databases.0.id}"
					upgrade_history_entry_id = "${data.oci_database_database_upgrade_history_entries.t.database_upgrade_history_entries.1.id}"
				}`, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "database_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "compartment_id"),

				//Upgrade history entry - plural datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.#", "2"),
				resource.TestCheckResourceAttrSet("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.id"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.state", "SUCCEEDED"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.source", "DB_SOFTWARE_IMAGE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.options", "-upgradeTimezone false -keepEvents"),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "source", "DB_SOFTWARE_IMAGE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "state", "SUCCEEDED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, "oci_database_db_system.t", "id")
					if resId != resId2 {
						return fmt.Errorf("expected same ocids, got different")
					}
					return err
				},
			),
		},
	})
}
