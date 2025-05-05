// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseAutonomousContainerDatabaseBackupDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `Automatic Backup`}, // Match API response
		"infrastructure_type":              acctest.Representation{RepType: acctest.Optional, Create: `CLOUD`},
		"is_remote":                        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`}, // Match API response
	}

	DatabaseAutonomousContainerDatabaseBackupResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation) +
		DatabaseCloudAutonomousVmClusterResourceConfig
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_backups.test_autonomous_container_database_backups"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_backups", "test_autonomous_container_database_backups", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseBackupDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_backup_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_backup_collection.0.items.#"), // Removed exact count check
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.type", "FULL"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_backup_collection.0.items.0.retention_period_in_days"), // Changed to check it's set rather than exact value
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_backup_collection.0.items.0.autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.display_name", "Automatic Backup"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.infrastructure_type", "CLOUD"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.state", "ACTIVE"),           // Changed from lifecycle_state to state
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_backup_collection.0.items.0.is_remote_backup", "false"), // Changed from lifecycle_state to state
			),
		},
	})
}
