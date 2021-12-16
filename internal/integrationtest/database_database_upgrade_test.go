// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
	DatabasePrecheckResourceRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_database_database_upgrade", "test_database_upgrade", acctest.Optional, acctest.Update, databasePrecheckRepresentation)
	DatabaseUpgradeResourceRepresentation  = acctest.GenerateResourceFromRepresentationMap("oci_database_database_upgrade", "test_database_upgrade", acctest.Optional, acctest.Update, databaseUpgradeRepresentation)

	databasePrecheckRepresentation = map[string]interface{}{
		"action":                          acctest.Representation{RepType: acctest.Required, Create: `PRECHECK`},
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.id}`},
		"database_upgrade_source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databasePrecheckDatabaseUpgradeSourceDbVersionRepresentation},
	}

	databasePrecheckDatabaseUpgradeSourceDbVersionRepresentation = map[string]interface{}{
		"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"source":     acctest.Representation{RepType: acctest.Optional, Create: `DB_VERSION`},
	}

	databaseUpgradeRepresentation = map[string]interface{}{
		"action":                          acctest.Representation{RepType: acctest.Required, Create: `UPGRADE`},
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.t.databases.0.id}`},
		"database_upgrade_source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseUpgradeDatabaseUpgradeSourceDbVersionRepresentation},
	}

	databaseUpgradeDatabaseUpgradeSourceDbVersionRepresentation = map[string]interface{}{
		"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"options":    acctest.Representation{RepType: acctest.Optional, Create: `-upgradeTimezone false -keepEvents`},
		"source":     acctest.Representation{RepType: acctest.Optional, Create: `DB_VERSION`},
	}

	dbSystemForDbUpgradeRepresentation = `
		resource "oci_database_db_system" "t" {
			availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
			compartment_id = "${var.compartment_id}"
			subnet_id = "${oci_core_subnet.t.id}"
			database_edition = "ENTERPRISE_EDITION"
			disk_redundancy = "NORMAL"
			shape = "VM.Standard1.1"
			ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
			display_name = "-tf-dbSystem-001"
			domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
			hostname = "myOracleDB" // this will be lowercased server side
			data_storage_size_in_gb = "256"
			license_model = "LICENSE_INCLUDED"
			node_count = "1"
			fault_domains = ["FAULT-DOMAIN-1"]
			db_home {
				db_version = "12.2.0.1"
				display_name = "-tf-db-home"
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
func TestDatabaseDatabaseUpgradeResource_basic(t *testing.T) {
	if strings.Contains(utils.GetEnvSettingWithBlankDefault("suppressed_tests"), "Database_upgrade") {
		t.Skip("Skipping suppressed upgrade_tests")
	}

	httpreplay.SetScenario("TestDatabaseDatabaseUpgradeResource_basic")
	defer httpreplay.SaveScenario()

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(ResourceDatabaseBaseConfig+dbSystemForDbUpgradeRepresentation, "database", "databaseUpgrade", t)

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
		// verify PRECHECK action on database with source=DB_VERSION
		{
			Config: ResourceDatabaseBaseConfig + DatabasePrecheckResourceRepresentation + dbSystemForDbUpgradeRepresentation,
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
			Config: ResourceDatabaseBaseConfig + DatabasePrecheckResourceRepresentation + dbSystemForDbUpgradeRepresentation + ResourceDatabaseTokenFn(`
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
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.0.source", "DB_VERSION"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.0.target_db_version", "19.0.0.0"),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "target_db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "action", "PRECHECK"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "source", "DB_VERSION"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "state", "SUCCEEDED"),
			),
		},
		// verify UPGRADE action on database with source=DB_VERSION
		{
			Config: ResourceDatabaseBaseConfig + DatabaseUpgradeResourceRepresentation + dbSystemForDbUpgradeRepresentation,
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
			Config: ResourceDatabaseBaseConfig + DatabaseUpgradeResourceRepresentation + dbSystemForDbUpgradeRepresentation + ResourceDatabaseTokenFn(`
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

				resource.TestCheckResourceAttr("data.oci_database_db_homes.t", "db_homes.0.db_version", "19.0.0.0"),

				//Upgrade history entry - plural datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.#", "2"),
				resource.TestCheckResourceAttrSet("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.id"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.state", "SUCCEEDED"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.source", "DB_VERSION"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.target_db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.t", "database_upgrade_history_entries.1.options", "-upgradeTimezone false -keepEvents"),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "target_db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.t", "source", "DB_VERSION"),
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
