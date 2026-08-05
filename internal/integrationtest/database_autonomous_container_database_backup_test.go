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
	DatabaseAutonomousContainerDatabaseBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `Automatic Backup`}, // Match API response
		"infrastructure_type":              acctest.Representation{RepType: acctest.Optional, Create: `CLOUD`},
		"is_remote":                        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`}, // Match API response
		"depends_on":                       acctest.Representation{RepType: acctest.Required, Create: []string{`oci_database_autonomous_database.test_autonomous_database`}},
	}

	DatabaseAutonomousContainerDatabaseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups.autonomous_container_database_backup_collection.0.items.0.id}`},
	}

	DatabaseAutonomousContainerDatabaseBackupResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation) +
		DatabaseCloudAutonomousVmClusterRequiredOnlyResource +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseForAcdBackupRepresentation)

	DatabaseAutonomousDatabaseForAcdBackupRepresentation = acctest.RepresentationCopyWithNewProperties(autonomousDatabaseDedicatedRepresentation, map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: DbaasIgnoreDefinedTagsRepresentation},
	})

	DatabaseExaccAutonomousContainerDatabaseBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"infrastructure_type":              acctest.Representation{RepType: acctest.Optional, Create: `CLOUD_AT_CUSTOMER`},
		"is_remote":                        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"depends_on":                       acctest.Representation{RepType: acctest.Required, Create: []string{`oci_database_autonomous_database.test_autonomous_database`}},
	}

	DatabaseExaccAutonomousContainerDatabaseBackupSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_backup_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups.autonomous_container_database_backup_collection.0.items.0.id}`},
	}

	DatabaseExaccAutonomousContainerDatabaseBackupResourceConfig = ExaccDatabaseAutonomousContainerDatabaseResourceDependency +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create, ExaccAutonomousContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, DatabaseExaccAutonomousDatabaseForAcdBackupRepresentation)

	DatabaseExaccAutonomousDatabaseForAcdBackupRepresentation = acctest.RepresentationCopyWithNewProperties(basicAutonomousDatabaseRepresentation, map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: DbaasIgnoreDefinedTagsRepresentation},
	})
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups"
	singularDatasourceName := "data.oci_database_autonomous_container_database_backup.test_autonomous_container_database_backup"
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource for Cloud ACD backups with ADBs in backup
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backups", "test_autonomous_container_database_backups", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseBackupDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backup", "test_autonomous_container_database_backup", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_backup_collection.#"),
				testCheckResourceAttrCountAtLeast(datasourceName, "autonomous_container_database_backup_collection.0.items.#", 1), // there should be at least one ACD Backup
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_backup_collection.0.items.0.autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.infrastructure_type", "CLOUD"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.state", "ACTIVE"),
				testCheckAnyListItemNestedResourceAttrCountAtLeast(datasourceName, "autonomous_container_database_backup_collection.0.items", "autonomous_databases.#", 1),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_backup_id"), // GET backup call validations start from here.
				acctest.TestCheckResourceAttributesEqual(singularDatasourceName, "id", datasourceName, "autonomous_container_database_backup_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_type", "CLOUD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				testCheckResourceAttrCountAtLeast(singularDatasourceName, "autonomous_databases.#", 1),
			),
		},
	})
}

// issue-routing-tag: database/ExaCC
func TestDatabaseExaccAutonomousContainerDatabaseBackupResource_basic(t *testing.T) {
	shouldSkipEXACCtest := utils.GetEnvSettingWithDefault("TF_VAR_should_skip_exacc_test", "false")

	if shouldSkipEXACCtest == "true" {
		t.Skip("Skipping TestDatabaseExaccAutonomousContainerDatabaseBackupResource_basic test.\n" + "Current TF_VAR_should_skip_exacc_test=" + shouldSkipEXACCtest)
	}

	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups"
	singularDatasourceName := "data.oci_database_autonomous_container_database_backup.test_autonomous_container_database_backup"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource for ExaCC ACD backups with ADBs in backup
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backups", "test_autonomous_container_database_backups", acctest.Required, acctest.Create, DatabaseExaccAutonomousContainerDatabaseBackupDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backup", "test_autonomous_container_database_backup", acctest.Required, acctest.Create, DatabaseExaccAutonomousContainerDatabaseBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExaccAutonomousContainerDatabaseBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_backup_collection.#"),
				testCheckResourceAttrCountAtLeast(datasourceName, "autonomous_container_database_backup_collection.0.items.#", 1), // there should be at least one ACD Backup
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_backup_collection.0.items.0.autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.infrastructure_type", "CLOUD_AT_CUSTOMER"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.state", "ACTIVE"),
				testCheckAnyListItemNestedResourceAttrCountAtLeast(datasourceName, "autonomous_container_database_backup_collection.0.items", "autonomous_databases.#", 1),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_backup_id"), // GET backup call validations start from here.
				acctest.TestCheckResourceAttributesEqual(singularDatasourceName, "id", datasourceName, "autonomous_container_database_backup_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_type", "CLOUD_AT_CUSTOMER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				testCheckResourceAttrCountAtLeast(singularDatasourceName, "autonomous_databases.#", 1),
			),
		},
	})
}

/*
testCheckResourceAttrCountAtLeast reads a Terraform state attribute that
represents a numeric count, such as a list/set "#"-count attribute, and
verifies that the parsed value is greater than or equal to the expected
minimum.
*/
func testCheckResourceAttrCountAtLeast(name, key string, min int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		value, err := acctest.FromInstanceState(s, name, key)
		if err != nil {
			return err
		}

		count, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("%s.%s value %q is not a number: %w", name, key, value, err)
		}

		if count < min {
			return fmt.Errorf("%s.%s count = %d, expected at least %d", name, key, count, min)
		}

		return nil
	}
}

/*
testCheckAnyListItemNestedResourceAttrCountAtLeast iterates over every item
in a Terraform state list and checks a nested count attribute under each
item. The check passes when at least one list item has the nested count
greater than or equal to the expected minimum.
*/
func testCheckAnyListItemNestedResourceAttrCountAtLeast(name, listKey, nestedKey string, min int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		value, err := acctest.FromInstanceState(s, name, listKey+".#")
		if err != nil {
			return err
		}

		itemCount, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("%s.%s.# value %q is not a number: %w", name, listKey, value, err)
		}

		for i := 0; i < itemCount; i++ {
			countValue, err := acctest.FromInstanceState(s, name, fmt.Sprintf("%s.%d.%s", listKey, i, nestedKey))
			if err != nil {
				continue
			}

			count, err := strconv.Atoi(countValue)
			if err != nil {
				return fmt.Errorf("%s.%s.%d.%s value %q is not a number: %w", name, listKey, i, nestedKey, countValue, err)
			}

			if count >= min {
				return nil
			}
		}

		return fmt.Errorf("%s.%s contains no item with %s count at least %d", name, listKey, nestedKey, min)
	}
}
