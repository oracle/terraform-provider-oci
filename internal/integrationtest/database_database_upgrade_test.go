// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var (

	// RESOURCE(S)

	// DbSystem for Upgrade Representation 19
	dbSystemForUpgradeRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDbSystemForUpgrade`},
		"database_edition":        acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"fault_domains":           acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]interface{}{"example-tag-namespace-all.example-tag": "originalValue"}},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]interface{}{"Department": "Finance", "Author": "Esteban Cabrera"}},
		"nsg_ids":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_nsg.id}`}},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`}},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `tfOracleDb`},
		"db_system_options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbSystemOptions},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeGroup},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: definedTagsIgnoreGroup},
	}

	dbSystemOptions = map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Optional, Create: `LVM`},
	}

	dbHomeGroup = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `tfDbHome`},
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `12.2.0.1`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseGroup},
	}

	databaseGroup = map[string]interface{}{
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `tfDb`},
		"pdb_name":       acctest.Representation{RepType: acctest.Required, Create: `tfPdb`},
		"character_set":  acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"ncharacter_set": acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"db_workload":    acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"admin_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
	}

	definedTagsIgnoreGroup = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{
			`db_home[0].db_version`,
			`defined_tags`,
		}},
	}

	// Database Upgrade with [Precheck] Representation
	dbPrecheckRepresentation = map[string]interface{}{
		"action":                          acctest.Representation{RepType: acctest.Required, Create: `PRECHECK`},
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_db_system_for_upgrade.databases.0.id}`},
		"database_upgrade_source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbPrecheckDatabaseUpgradeSourceDetailsGroup},
	}

	dbPrecheckDatabaseUpgradeSourceDetailsGroup = map[string]interface{}{
		"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19.24.0.0`},
		"source":     acctest.Representation{RepType: acctest.Optional, Create: `DB_VERSION`},
	}

	// Database [Upgrade] Representation
	dbUpgradeRepresentation = map[string]interface{}{
		"action":                          acctest.Representation{RepType: acctest.Required, Create: `UPGRADE`},
		"database_id":                     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_db_system_for_upgrade.databases.0.id}`},
		"database_upgrade_source_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: dbUpgradeDatabaseUpgradeSourceDetailsGroup},
	}

	dbUpgradeDatabaseUpgradeSourceDetailsGroup = map[string]interface{}{
		"db_version": acctest.Representation{RepType: acctest.Optional, Create: `19.24.0.0`},
		"options":    acctest.Representation{RepType: acctest.Optional, Create: `-upgradeTimezone false -keepEvents`},
		"source":     acctest.Representation{RepType: acctest.Optional, Create: `DB_VERSION`},
	}

	// DATASOURCE(S)

	// DbSystemForUpgrade [DbSystem] Datasource Representation
	dbSystemForUpgradeDatasourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: datasourceFilterGroup},
	}

	datasourceFilterGroup = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_database_db_system.test_db_system_for_upgrade.id}`}},
	}

	// DbSystemForUpgrade [DbHomes] Datasource Representation
	dbSystemForUpgradeDbHomesDatasourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_system_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_db_system.test_db_system_for_upgrade.id}`},
	}

	// DbSystemForUpgrade [DbHome] Datasource Representation
	dbSystemForUpgradeDbHomeDatasourceRepresentation = map[string]interface{}{
		"db_home_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_db_homes.test_db_system_for_upgrade.db_homes.0.db_home_id}`},
	}

	// DbSystemForUpgrade [Databases] Datasource Representation
	dbSystemForUpgradeDatabasesDatasourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_home_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_db_homes.test_db_system_for_upgrade.db_homes.0.db_home_id}`},
	}

	// DbSystemForUpgrade [Database] Datasource Representation
	dbSystemForUpgradeDatabaseDatasourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_db_system_for_upgrade.databases.0.id}`},
	}

	// Database Upgrade History Entries Datasource Representation
	dbUpgradeHistoryEntriesDatasourceRepresentation = map[string]interface{}{
		"database_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_databases.test_db_system_for_upgrade.databases.0.id}`},
	}

	// Database Upgrade History First Entry Datasource Representation
	dbUpgradeHistoryFirstEntryDatasourceRepresentation = map[string]interface{}{
		"database_id":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_db_system_for_upgrade.databases.0.id}`},
		"upgrade_history_entry_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade.database_upgrade_history_entries.0.id}`},
	}

	// Database Upgrade History Second Entry Datasource Representation
	dbUpgradeHistorySecondEntryDatasourceRepresentation = map[string]interface{}{
		"database_id":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_databases.test_db_system_for_upgrade.databases.0.id}`},
		"upgrade_history_entry_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade.database_upgrade_history_entries.1.id}`},
	}

	// RESOURCE CONFIG(S)
	DbSystemForUpgradeConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbSystemForUpgradeRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_systems", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbSystemForUpgradeDatasourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_homes", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbSystemForUpgradeDbHomesDatasourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_home", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbSystemForUpgradeDbHomeDatasourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_databases", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbSystemForUpgradeDatabasesDatasourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_database", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbSystemForUpgradeDatabaseDatasourceRepresentation)
	dbUpgradeConfig  = acctest.GenerateResourceFromRepresentationMap("oci_database_database_upgrade", "test_database_upgrade", acctest.Optional, acctest.Create, dbUpgradeRepresentation)
	dbPrecheckConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_database_upgrade", "test_database_precheck", acctest.Optional, acctest.Create, dbPrecheckRepresentation)

	// Upstream Resource Dependencies
	// NSG Representation
	nsgRepresentation = map[string]interface{}{
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `tfNsg`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
	}

	// NSG Config
	nsgConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_nsg", acctest.Optional, acctest.Create, nsgRepresentation)

	dbSystemBaseConfig = acctest.LegacyTestProviderConfig() + nsgConfig + DbSystemBaseConfig

	// DbSystemForUpgrade Resource Name
	resourceName = "oci_database_db_system.test_db_system_for_upgrade"

	DatabaseSystemForDbUpgradeRepresentation = DbSystemForUpgradeConfig // TODO: escabrer, Remove later
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
	//acctest.SaveConfigContent(dbSystemBaseConfig+DbSystemForUpgradeConfig, "database", "databaseUpgrade", t)

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
		// verify PRECHECK action on database with source=DB_VERSION
		{
			Config: dbSystemBaseConfig + dbPrecheckConfig + DbSystemForUpgradeConfig,
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
			Config: dbSystemBaseConfig + dbPrecheckConfig + DbSystemForUpgradeConfig +
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
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.0.source", "DB_VERSION"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.0.target_db_version", "19.24.0.0"),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "target_db_version", "19.24.0.0"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "action", "PRECHECK"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "source", "DB_VERSION"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "state", "SUCCEEDED"),
			),
		},
		// verify UPGRADE action on database with source=DB_VERSION
		{
			Config: dbSystemBaseConfig + dbUpgradeConfig + DbSystemForUpgradeConfig,
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
			Config: dbSystemBaseConfig + dbUpgradeConfig + DbSystemForUpgradeConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_upgrade_history_entries", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbUpgradeHistoryEntriesDatasourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_upgrade_history_entry", "test_db_system_for_upgrade", acctest.Optional, acctest.Create, dbUpgradeHistorySecondEntryDatasourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// Database
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "id"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "database_id"),
				resource.TestCheckResourceAttr("data.oci_database_database.test_db_system_for_upgrade", "character_set", "AL32UTF8"),
				resource.TestCheckResourceAttrSet("data.oci_database_database.test_db_system_for_upgrade", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_database_db_homes.test_db_system_for_upgrade", "db_homes.0.db_version", "19.24.0.0"),

				//Upgrade history entry - plural datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.#", "2"),
				resource.TestCheckResourceAttrSet("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.id"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.state", "SUCCEEDED"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.source", "DB_VERSION"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entries.test_db_system_for_upgrade", "database_upgrade_history_entries.1.target_db_version", "19.24.0.0"),

				//Upgrade history entry - singular datasource
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "target_db_version", "19.24.0.0"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "action", "UPGRADE"),
				resource.TestCheckResourceAttr("data.oci_database_database_upgrade_history_entry.test_db_system_for_upgrade", "source", "DB_VERSION"),
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

// Both inputs execute a genuine PRECHECK against an owned BaseDB source DB.
// ExaCS requires DB_HOME upgrade sources, which this resource does not expose.
func TestDatabaseDatabaseUpgradeResourceOlderVersionAcknowledgement(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseUpgradeResourceOlderVersionAcknowledgement")
	defer httpreplay.SaveScenario()
	databaseAckRequiredInput(t, "ssh_public_key")
	databaseAckRequiredInput(t, "ack_test_admin_password")
	sourceVersion := databaseAckRequiredInput(t, "ack_test_db_home_version")
	target := databaseAckRequiredInput(t, "ack_test_upgrade_target_version")
	shape := databaseAckRequiredInput(t, "ack_test_upgrade_db_shape")
	if shape != "BM.DenseIO2.52" {
		t.Fatal("the live PRECHECK baseline for this focused test validates BM.DenseIO2.52/ASM")
	}
	if sourceVersion == target {
		t.Fatal("PRECHECK target must differ from the selected source version")
	}
	compartment := databaseAckRequiredInput(t, "compartment_ocid")
	config := acctest.ProviderTestConfig() + fmt.Sprintf("variable \"compartment_id\" { default = %q }\n", compartment) + `
variable "ack_test_db_home_version" { type = string }
variable "ack_test_upgrade_target_version" { type = string }
variable "ack_test_upgrade_db_shape" { type = string }
variable "ack_test_admin_password" {
 type = string
 sensitive = true
}
resource "oci_core_vcn" "ack_upgrade_vcn" {
 compartment_id = var.compartment_id
 cidr_block = "10.253.0.0/16"
 display_name = "TERSI-5426-ack-precheck-vcn"
 dns_label = "ackupgvcn"
}
resource "oci_core_nat_gateway" "ack_upgrade_nat" {
 compartment_id = var.compartment_id
 vcn_id = oci_core_vcn.ack_upgrade_vcn.id
 display_name = "TERSI-5426-ack-precheck-nat"
}
data "oci_identity_availability_domains" "ack_upgrade_ads" {
 compartment_id = var.compartment_id
}
resource "oci_core_security_list" "ack_upgrade_security" {
 compartment_id = var.compartment_id
 vcn_id = oci_core_vcn.ack_upgrade_vcn.id
 display_name = "TERSI-5426-ack-precheck-security"
 ingress_security_rules {
  source = oci_core_vcn.ack_upgrade_vcn.cidr_block
  protocol = "all"
 }
 egress_security_rules {
  destination = "0.0.0.0/0"
  protocol = "all"
 }
}
resource "oci_core_route_table" "ack_upgrade_routes" {
 compartment_id = var.compartment_id
 vcn_id = oci_core_vcn.ack_upgrade_vcn.id
 route_rules {
  destination = "0.0.0.0/0"
  destination_type = "CIDR_BLOCK"
  network_entity_id = oci_core_nat_gateway.ack_upgrade_nat.id
 }
}
resource "oci_core_subnet" "ack_upgrade_subnet" {
 compartment_id = var.compartment_id
 vcn_id = oci_core_vcn.ack_upgrade_vcn.id
 cidr_block = cidrsubnet(oci_core_vcn.ack_upgrade_vcn.cidr_block, 8, 72)
 display_name = "TERSI-5426-ack-precheck-subnet"
 dns_label = "ackupg72"
 prohibit_public_ip_on_vnic = true
 dhcp_options_id = oci_core_vcn.ack_upgrade_vcn.default_dhcp_options_id
 route_table_id = oci_core_route_table.ack_upgrade_routes.id
 security_list_ids = [oci_core_security_list.ack_upgrade_security.id]
}
resource "oci_database_db_system" "ack_upgrade_source" {
 availability_domain = data.oci_identity_availability_domains.ack_upgrade_ads.availability_domains[0].name
 compartment_id = var.compartment_id
 database_edition = "ENTERPRISE_EDITION"
 shape = var.ack_test_upgrade_db_shape
 cpu_core_count = 2
 data_storage_percentage = 80
 disk_redundancy = "NORMAL"
 license_model = "LICENSE_INCLUDED"
 node_count = 1
 hostname = "ackprecheck"
 domain = oci_core_subnet.ack_upgrade_subnet.subnet_domain_name
 display_name = "TERSI-5426-ack-precheck-system"
 subnet_id = oci_core_subnet.ack_upgrade_subnet.id
 ssh_public_keys = [var.ssh_public_key]
 db_home {
  db_version = var.ack_test_db_home_version
  display_name = "TERSI-5426-ack-precheck-home"
  database {
   admin_password = var.ack_test_admin_password
   db_name = "ACKUPG01"
   db_backup_config { auto_backup_enabled = false }
  }
 }
 lifecycle {
  ignore_changes = [db_home[0].db_version]
 }
}
data "oci_database_db_homes" "ack_upgrade_source_homes" {
 compartment_id = var.compartment_id
 db_system_id = oci_database_db_system.ack_upgrade_source.id
}
data "oci_database_databases" "ack_upgrade_source_databases" {
 compartment_id = var.compartment_id
 db_home_id = data.oci_database_db_homes.ack_upgrade_source_homes.db_homes[0].db_home_id
}
`
	name := "oci_database_database_upgrade.ack_precheck"
	historyName := "data.oci_database_database_upgrade_history_entries.ack_precheck_history"
	configFor := func(ack bool) string {
		return config + fmt.Sprintf(`
resource "oci_database_database_upgrade" "ack_precheck" {
 action = "PRECHECK"
 database_id = data.oci_database_databases.ack_upgrade_source_databases.databases[0].id
 database_upgrade_source_details {
  source = "DB_VERSION"
  db_version = var.ack_test_upgrade_target_version
  is_older_version_acknowledged = %t
 }
}
data "oci_database_database_upgrade_history_entries" "ack_precheck_history" {
 database_id = data.oci_database_databases.ack_upgrade_source_databases.databases[0].id
 upgrade_action = "PRECHECK"
 state = "SUCCEEDED"
 depends_on = [oci_database_database_upgrade.ack_precheck]
}
`, ack)
	}
	historyIDs := map[string]bool{}
	databaseAcknowledgementResourceTest(t, databaseAckBaseDBDestroy, []resource.TestStep{
		{Config: configFor(true), Check: acctest.ComposeAggregateTestCheckFuncWrapper(
			resource.TestCheckResourceAttr(name, "action", "PRECHECK"),
			resource.TestCheckResourceAttr(name, "database_upgrade_source_details.0.is_older_version_acknowledged", "true"),
			resource.TestCheckResourceAttr("oci_database_db_system.ack_upgrade_source", "state", "AVAILABLE"),
			databaseAckRequireNewPrecheckHistory(historyName, target, historyIDs),
		)},
		{Config: configFor(true), PlanOnly: true},
		{Config: configFor(false), Check: acctest.ComposeAggregateTestCheckFuncWrapper(
			resource.TestCheckResourceAttr(name, "database_upgrade_source_details.0.is_older_version_acknowledged", "false"),
			resource.TestCheckResourceAttr("oci_database_db_system.ack_upgrade_source", "state", "AVAILABLE"),
			databaseAckRequireNewPrecheckHistory(historyName, target, historyIDs),
		)},
	})
}
