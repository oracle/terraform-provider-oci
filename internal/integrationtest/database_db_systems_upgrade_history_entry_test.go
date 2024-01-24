// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseDbSystemsUpgradeHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"upgrade_history_entry_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
	}

	DatabaseDatabaseDbSystemsUpgradeHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"db_system_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"upgrade_action": acctest.Representation{RepType: acctest.Optional, Create: `PRECHECK`},
	}

	DatabaseDbSystemsUpgradeHistoryEntryResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", acctest.Required, acctest.Create, CoreVlanRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, DatabaseBackupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, DatabaseDbHomeRepresentation) +
		DatabaseBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationNFSRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": acctest.Representation{RepType: acctest.Optional, Update: activationFilePath}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, vmClusterNetworkValidateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Required, acctest.Create, DatabaseDatabaseDbSystemsUpgradeHistoryEntryDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseAutonomousExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create, IdentityDomainRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseDbSystemsUpgradeHistoryEntryResource_basic(t *testing.T) {
	t.Skip("Skip this test because upgrade history test is done in database_db_systems_upgrade_test.go")
	httpreplay.SetScenario("TestDatabaseDbSystemsUpgradeHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_systems_upgrade_history_entries.test_db_systems_upgrade_history_entries"
	singularDatasourceName := "data.oci_database_db_systems_upgrade_history_entry.test_db_systems_upgrade_history_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_systems_upgrade_history_entries", "test_db_systems_upgrade_history_entries", acctest.Required, acctest.Create, DatabaseDatabaseDbSystemsUpgradeHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbSystemsUpgradeHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "upgrade_action", "PRECHECK"),

				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.action"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.new_gi_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.new_os_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.old_gi_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.old_os_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.snapshot_retention_period_in_days"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.time_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_upgrade_history_entries.0.time_started"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_systems_upgrade_history_entry", "test_db_systems_upgrade_history_entry", acctest.Required, acctest.Create, DatabaseDatabaseDbSystemsUpgradeHistoryEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbSystemsUpgradeHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "upgrade_history_entry_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "new_gi_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "new_os_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "old_gi_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "old_os_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "snapshot_retention_period_in_days"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
	})
}
