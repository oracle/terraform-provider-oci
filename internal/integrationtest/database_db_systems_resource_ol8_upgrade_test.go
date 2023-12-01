// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	DbSystemPrecheckResourceRepresentationOl8 = acctest.GenerateResourceFromRepresentationMap("oci_database_db_systems_upgrade", "test_db_system_upgrade", acctest.Optional, acctest.Update, dbSystemPrecheckRepresentationOl8)
	DbSystemUpgradeResourceRepresentationOl8  = acctest.GenerateResourceFromRepresentationMap("oci_database_db_systems_upgrade", "test_db_system_upgrade", acctest.Optional, acctest.Update, dbSystemUpgradeRepresentationOl8)

	dbSystemPrecheckRepresentationOl8 = map[string]interface{}{
		"action":       acctest.Representation{RepType: acctest.Required, Create: `PRECHECK`},
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_systems.t.db_systems.0.id}`},
		"is_snapshot_retention_days_force_updated": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"new_os_version":                    acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithDefault("new_os_version", "oracle linux server release 8")},
		"snapshot_retention_period_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	dbSystemUpgradeRepresentationOl8 = map[string]interface{}{
		"action":       acctest.Representation{RepType: acctest.Required, Create: `UPGRADE`},
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_systems.t.db_systems.0.id}`},
		"is_snapshot_retention_days_force_updated": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"new_os_version":                    acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithDefault("new_os_version", "oracle linux server release 8")},
		"snapshot_retention_period_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	dbSystemForDbSystemUpgradeRepresentationOl8 = `
		resource "oci_database_db_system" "t" {
			availability_domain = "${oci_core_subnet.t.availability_domain}"
			compartment_id = "${var.compartment_id}"
			subnet_id = "${oci_core_subnet.t.id}"
			database_edition = "ENTERPRISE_EDITION"
			disk_redundancy = "NORMAL"
			shape = "VM.Standard2.2"
			ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
			display_name = "tfDbSystemTest"
			domain = "${oci_core_subnet.t.subnet_domain_name}"
			hostname = "myOracleDB" // this will be lowercased server side
			data_storage_size_in_gb = "256"
			license_model = "LICENSE_INCLUDED"
			node_count = "1"
			db_home {
				db_version = "19.0.0.0"
				display_name = "dbHome1"
				database {
					admin_password = "BEstrO0ng_#11"
					db_name = "aTFdb"
					character_set = "AL32UTF8"
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					ncharacter_set = "AL16UTF16"
					db_workload = "OLTP"
					pdb_name = "pdbName"
				}
			}
			db_system_options {
				storage_management = "LVM"
			}
			defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
			freeform_tags = {"Department" = "Finance"}
			nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
			lifecycle {
				ignore_changes = [
					db_home.0.db_version,
					defined_tags,
					db_home.0.database.0.defined_tags,
				]
			}
		}
		data "oci_database_db_systems" "t" {
			compartment_id = "${var.compartment_id}"
			filter {
				name   = "id"
				values = ["${oci_database_db_system.t.id}"]
			}
		}
		data "oci_database_db_homes" "t" {
			compartment_id = "${var.compartment_id}"
			db_system_id = "${oci_database_db_system.t.id}"
			filter {
				name   = "db_system_id"
				values = ["${oci_database_db_system.t.id}"]
			}
		}
		data "oci_database_db_home" "t" {
			db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
		}
		data "oci_database_databases" "t" {
			compartment_id = "${var.compartment_id}"
			db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
			filter {
				name   = "db_name"
				values = ["${oci_database_db_system.t.db_home.0.database.0.db_name}"]
			}
		}
		data "oci_database_database" "t" {
			  database_id = "${data.oci_database_databases.t.databases.0.id}"
		}`
)

// issue-routing-tag: database/default
func TestDatabaseDbSystemsUpgradeResourceOl8(t *testing.T) {

	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "Db_system_upgrade") {
		t.Skip("Skipping suppressed upgrade_tests")
	}

	httpreplay.SetScenario("TestDatabaseDbSystemsUpgradeResourceOl8")
	defer httpreplay.SaveScenario()

	var resId, resId2 string

	var newOsVersion = utils.GetEnvSettingWithDefault("new_os_version", "oracle linux server release 8")

	acctest.SaveConfigContent(ResourceDatabaseBaseConfig+dbSystemForDbSystemUpgradeRepresentationOl8, "database", "dbSystem", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: ResourceDatabaseBaseConfig + dbSystemForDbSystemUpgradeRepresentationOl8,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.display_name", "dbHome1"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),

				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "compartment_id"),

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
		// Precheck
		{
			Config: ResourceDatabaseBaseConfig + DbSystemPrecheckResourceRepresentationOl8 + dbSystemForDbSystemUpgradeRepresentationOl8,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DBHome
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "db_home_id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_home.t", "compartment_id"),

				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "database_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "compartment_id"),
			),
		},
		// Upgrade
		{
			Config: ResourceDatabaseBaseConfig + DbSystemUpgradeResourceRepresentationOl8 + dbSystemForDbSystemUpgradeRepresentationOl8,
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
		// Upgrade History
		{
			Config: ResourceDatabaseBaseConfig + DbSystemUpgradeResourceRepresentationOl8 + dbSystemForDbSystemUpgradeRepresentationOl8 + ResourceDatabaseTokenFn(`
				data "oci_database_db_systems_upgrade_history_entries" "t" {
					db_system_id = "${data.oci_database_db_systems.t.db_systems.0.id}"
				}
				data "oci_database_db_systems_upgrade_history_entry" "t" {
					db_system_id = "${data.oci_database_db_systems.t.db_systems.0.id}"
					upgrade_history_entry_id = "${data.oci_database_db_systems_upgrade_history_entries.t.db_system_upgrade_history_entries.0.id}"
				}`, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "database_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.t", "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.t", "compartment_id"),

				//Upgrade history entry - plural datasource
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.#", "2"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.0.id"),
				resource.TestCheckResourceAttrSet("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.1.id"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.0.action", "PRECHECK"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.0.state", "SUCCEEDED"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.0.new_os_version", newOsVersion),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.0.old_os_version", ""),

				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.1.action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.1.state", "SUCCEEDED"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.1.new_os_version", newOsVersion),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entries.t", "db_system_upgrade_history_entries.1.old_os_version", ""),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entry.t", "new_os_version", newOsVersion),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entry.t", "old_os_version", ""),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entry.t", "action", "PRECHECK"),
				resource.TestCheckResourceAttr("data.oci_database_db_systems_upgrade_history_entry.t", "state", "SUCCEEDED"),

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
