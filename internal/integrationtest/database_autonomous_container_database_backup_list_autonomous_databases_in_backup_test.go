// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"time_stamp_requested":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups.autonomous_container_database_backup_collection.0.items.0.time_ended}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"depends_on":                       acctest.Representation{RepType: acctest.Required, Create: []string{`data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups`}},
	}

	DatabaseAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backups", "test_autonomous_container_database_backups", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseBackupDataSourceRepresentation) +
		DatabaseAutonomousContainerDatabaseBackupResourceConfig

	DatabaseExaccAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"time_stamp_requested":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups.autonomous_container_database_backup_collection.0.items.0.time_ended}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"depends_on":                       acctest.Representation{RepType: acctest.Required, Create: []string{`data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups`}},
	}

	DatabaseExaccAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backups", "test_autonomous_container_database_backups", acctest.Required, acctest.Create, DatabaseExaccAutonomousContainerDatabaseBackupDataSourceRepresentation) +
		DatabaseExaccAutonomousContainerDatabaseBackupResourceConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	backupDatasourceName := "data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups"
	datasourceName := "data.oci_database_autonomous_container_database_backup_list_autonomous_databases_in_backups.test_autonomous_container_database_backup_list_autonomous_databases_in_backups"
	autonomousDatabaseResourceName := "oci_database_autonomous_database.test_autonomous_database"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource for Cloud ADBs in an ACD backup at the current timestamp
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backup_list_autonomous_databases_in_backups", "test_autonomous_container_database_backup_list_autonomous_databases_in_backups", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(backupDatasourceName, "autonomous_container_database_backup_collection.#"),
				testCheckResourceAttrCountAtLeast(backupDatasourceName, "autonomous_container_database_backup_collection.0.items.#", 1),
				testCheckAnyListItemNestedResourceAttrCountAtLeast(backupDatasourceName, "autonomous_container_database_backup_collection.0.items", "autonomous_databases.#", 1),
				testCheckAnyNestedListContainsResourceAttrValue(backupDatasourceName, "autonomous_container_database_backup_collection.0.items", "autonomous_databases", "display_name", autonomousDatabaseResourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_in_backup_collection.#"),
				testCheckResourceAttrCountAtLeast(datasourceName, "autonomous_database_in_backup_collection.0.items.#", 1),
				testCheckListContainsResourceAttrValue(datasourceName, "autonomous_database_in_backup_collection.0.items", "display_name", autonomousDatabaseResourceName, "display_name"),
			),
		},
	})
}

// issue-routing-tag: database/ExaCC
func TestDatabaseExaccAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResource_basic(t *testing.T) {
	shouldSkipEXACCtest := utils.GetEnvSettingWithDefault("TF_VAR_should_skip_exacc_test", "false")

	if shouldSkipEXACCtest == "true" {
		t.Skip("Skipping TestDatabaseExaccAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResource_basic test.\n" + "Current TF_VAR_should_skip_exacc_test=" + shouldSkipEXACCtest)
	}

	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	backupDatasourceName := "data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups"
	datasourceName := "data.oci_database_autonomous_container_database_backup_list_autonomous_databases_in_backups.test_autonomous_container_database_backup_list_autonomous_databases_in_backups"
	autonomousDatabaseResourceName := "oci_database_autonomous_database.test_autonomous_database"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource for ExaCC ADBs in an ACD backup at the current timestamp
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backup_list_autonomous_databases_in_backups", "test_autonomous_container_database_backup_list_autonomous_databases_in_backups", acctest.Required, acctest.Create, DatabaseExaccAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExaccAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(backupDatasourceName, "autonomous_container_database_backup_collection.#"),
				testCheckResourceAttrCountAtLeast(backupDatasourceName, "autonomous_container_database_backup_collection.0.items.#", 1),
				testCheckAnyListItemNestedResourceAttrCountAtLeast(backupDatasourceName, "autonomous_container_database_backup_collection.0.items", "autonomous_databases.#", 1),
				testCheckAnyNestedListContainsResourceAttrValue(backupDatasourceName, "autonomous_container_database_backup_collection.0.items", "autonomous_databases", "display_name", autonomousDatabaseResourceName, "display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_database_in_backup_collection.#"),
				testCheckResourceAttrCountAtLeast(datasourceName, "autonomous_database_in_backup_collection.0.items.#", 1),
				testCheckListContainsResourceAttrValue(datasourceName, "autonomous_database_in_backup_collection.0.items", "display_name", autonomousDatabaseResourceName, "display_name"),
			),
		},
	})
}

func testCheckListContainsResourceAttrValue(name, listKey, attrKey, resourceName, resourceAttr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectedValue, err := acctest.FromInstanceState(s, resourceName, resourceAttr)
		if err != nil {
			return err
		}

		countValue, err := acctest.FromInstanceState(s, name, listKey+".#")
		if err != nil {
			return err
		}

		count, err := strconv.Atoi(countValue)
		if err != nil {
			return fmt.Errorf("%s.%s.# value %q is not a number: %w", name, listKey, countValue, err)
		}

		for i := 0; i < count; i++ {
			actualValue, err := acctest.FromInstanceState(s, name, fmt.Sprintf("%s.%d.%s", listKey, i, attrKey))
			if err != nil {
				continue
			}

			if actualValue == expectedValue {
				return nil
			}
		}

		return fmt.Errorf("%s.%s contains no item with %s matching %s.%s value %q", name, listKey, attrKey, resourceName, resourceAttr, expectedValue)
	}
}

func testCheckAnyNestedListContainsResourceAttrValue(name, listKey, nestedListKey, attrKey, resourceName, resourceAttr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectedValue, err := acctest.FromInstanceState(s, resourceName, resourceAttr)
		if err != nil {
			return err
		}

		countValue, err := acctest.FromInstanceState(s, name, listKey+".#")
		if err != nil {
			return err
		}

		count, err := strconv.Atoi(countValue)
		if err != nil {
			return fmt.Errorf("%s.%s.# value %q is not a number: %w", name, listKey, countValue, err)
		}

		for i := 0; i < count; i++ {
			nestedCountValue, err := acctest.FromInstanceState(s, name, fmt.Sprintf("%s.%d.%s.#", listKey, i, nestedListKey))
			if err != nil {
				continue
			}

			nestedCount, err := strconv.Atoi(nestedCountValue)
			if err != nil {
				return fmt.Errorf("%s.%s.%d.%s.# value %q is not a number: %w", name, listKey, i, nestedListKey, nestedCountValue, err)
			}

			for j := 0; j < nestedCount; j++ {
				actualValue, err := acctest.FromInstanceState(s, name, fmt.Sprintf("%s.%d.%s.%d.%s", listKey, i, nestedListKey, j, attrKey))
				if err != nil {
					continue
				}

				if actualValue == expectedValue {
					return nil
				}
			}
		}

		return fmt.Errorf("%s.%s contains no nested %s item with %s matching %s.%s value %q", name, listKey, nestedListKey, attrKey, resourceName, resourceAttr, expectedValue)
	}
}
