// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	DatabaseDbSystemOsPatchHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"os_patch_history_entry_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_system_os_patch_history_entries.test_db_system_os_patch_history_entries.db_system_os_patch_history_entry_collection.0.items.0.id}`},
	}

	DatabaseDbSystemOsPatchHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"action":       acctest.Representation{RepType: acctest.Optional, Create: `PRECHECK`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DbSystemResourceRepresentationForOsPatch = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Required, Create: `NORMAL`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.x86`},
		"compute_model":           acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"compute_count":           acctest.Representation{RepType: acctest.Required, Create: `4`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"domain":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myOracleDB`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemDbHomeGroupForOsPatch},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `tfDbSystemVmOsPatch`},
		"db_system_options":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DbSystemOptions},
		"os_patch_trigger":        acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"os_patch_action":         acctest.Representation{RepType: acctest.Optional, Create: `PRECHECK`, Update: `PRECHECK`},
	}

	DbSystemOptions = map[string]interface{}{
		"storage_management": acctest.Representation{RepType: acctest.Optional, Create: `LVM`},
	}

	DbSystemDbHomeGroupForOsPatch = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `tfDbHome`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemDatabaseGroupForOsPatch},
	}

	DbSystemDatabaseGroupForOsPatch = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: nil},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `tfDb`},
		"character_set":    acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"db_workload":      acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
		"pdb_name":         acctest.Representation{RepType: acctest.Required, Create: `tfPdb`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemDbBackupConfigGroupForOsPatch},
	}

	DbSystemDbBackupConfigGroupForOsPatch = map[string]interface{}{
		"auto_backup_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDbSystemOsPatchHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemOsPatchHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	resourceName := "oci_database_db_system.test_db_system"
	datasourceName := "data.oci_database_db_system_os_patch_history_entries.test_db_system_os_patch_history_entries"
	singularDatasourceName := "data.oci_database_db_system_os_patch_history_entry.test_db_system_os_patch_history_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Step 1: create db system (no patch yet)
		{
			Config: ResourceDatabaseBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Create, DbSystemResourceRepresentationForOsPatch),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_patch_trigger", "1"),
			),
		},
		// Step 2: update db system (trigger increases -> ExecuteDbSystemOsPatch runs)
		{
			Config: ResourceDatabaseBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Update, DbSystemResourceRepresentationForOsPatch),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB  System Resource tests
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_patch_trigger", "2"),
			),
		},

		// Step 3: verify list datasource
		{
			Config: ResourceDatabaseBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Update, DbSystemResourceRepresentationForOsPatch) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_database_db_system_os_patch_history_entries",
					"test_db_system_os_patch_history_entries",
					acctest.Required,
					acctest.Create,
					DatabaseDbSystemOsPatchHistoryEntryDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "db_system_os_patch_history_entry_collection.#", "1"),
				// Check first item in collection
				resource.TestCheckResourceAttr(datasourceName, "db_system_os_patch_history_entry_collection.0.items.0.action", "PRECHECK"),
				resource.TestCheckResourceAttr(datasourceName, "db_system_os_patch_history_entry_collection.0.items.0.state", "SUCCEEDED"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_os_patch_history_entry_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_os_patch_history_entry_collection.0.items.0.time_started"),
			),
		},

		// Step 4: verify singular datasource
		{
			Config: ResourceDatabaseBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Update, DbSystemResourceRepresentationForOsPatch) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_database_db_system_os_patch_history_entries",
					"test_db_system_os_patch_history_entries",
					acctest.Required,
					acctest.Create,
					DatabaseDbSystemOsPatchHistoryEntryDataSourceRepresentation,
				) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_database_db_system_os_patch_history_entry",
					"test_db_system_os_patch_history_entry",
					acctest.Required,
					acctest.Create,
					DatabaseDbSystemOsPatchHistoryEntrySingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_patch_history_entry_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action", "PRECHECK"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_patch_details.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "SUCCEEDED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	resource.AddTestSweepers("DatabaseDbSystem", &resource.Sweeper{
		Name:         "DatabaseDbSystem",
		Dependencies: acctest.DependencyGraph["dbSystem"],
		F:            sweepDatabaseDbSystemResource,
	})
}
