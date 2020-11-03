// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	databaseUpgradeHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"database_id":              Representation{repType: Required, create: `${data.oci_database_databases.db.databases.0.id}`},
		"upgrade_history_entry_id": Representation{repType: Required, create: `{}`},
	}

	databaseUpgradeHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"database_id":    Representation{repType: Required, create: `${data.oci_database_databases.db.databases.0.id}`},
		"state":          Representation{repType: Optional, create: `AVAILABLE`},
		"upgrade_action": Representation{repType: Optional, create: `PRECHECK`},
	}

	DatabaseUpgradeHistoryEntryResourceConfig = generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vlan", "test_vlan", Required, Create, vlanRepresentation) +
		generateResourceFromRepresentationMap("oci_database_backup", "test_backup", Required, Create, backupRepresentation) +
		generateResourceFromRepresentationMap("oci_database_database", "test_database", Required, Create, databaseRepresentation) +
		generateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", Required, Create, dbHomeRepresentationBase) +
		BackupResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Create, backupDestinationNFSRepresentation) +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, representationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{repType: Optional, update: activationFilePath}})) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation) +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureRepresentation) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create, vmClusterNetworkRepresentation) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation) +
		AvailabilityDomainConfig
)

func TestDatabaseDatabaseUpgradeHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDatabaseUpgradeHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_database_upgrade_history_entries.test_database_upgrade_history_entries"
	singularDatasourceName := "data.oci_database_database_upgrade_history_entry.test_database_upgrade_history_entry"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_database_upgrade_history_entries", "test_database_upgrade_history_entries", Required, Create, databaseUpgradeHistoryEntryDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseUpgradeHistoryEntryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(datasourceName, "upgrade_action", "PRECHECK"),

					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.action"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.source"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.source_db_home_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.target_db_version"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.target_database_software_image_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.target_db_home_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.time_ended"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_upgrade_history_entries.0.time_started"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_database_upgrade_history_entry", "test_database_upgrade_history_entry", Required, Create, databaseUpgradeHistoryEntrySingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseUpgradeHistoryEntryResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "upgrade_history_entry_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "action"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "source"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "source_db_home_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "target_db_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "target_database_software_image_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "target_db_home_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				),
			},
		},
	})
}
