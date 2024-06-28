// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	DatabaseExadbVmClusterUpdateHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"exadb_vm_cluster_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.exadb_vm_cluster_id}`},
		"update_history_entry_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_exadb_vm_cluster_update_history_entries.test_exadb_vm_cluster_update_history_entries.exadb_vm_cluster_update_history_entries[0].id}`},
	}

	DatabaseExadbVmClusterUpdateHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"exadb_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.exadb_vm_cluster_id}`},
		"update_type":         acctest.Representation{RepType: acctest.Optional, Create: `OS_UPDATE`},
	}

	// Note: set env variable TF_VAR_exadb_vm_cluster_id before running this test
	DatabaseExadbVmClusterUpdateHistoryEntryResourceConfig = `variable "exadb_vm_cluster_id" {}`
)

// issue-routing-tag: database/default
func TestDatabaseExadbVmClusterUpdateHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadbVmClusterUpdateHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_database_exadb_vm_cluster_update_history_entries.test_exadb_vm_cluster_update_history_entries"
	singularDatasourceName := "data.oci_database_exadb_vm_cluster_update_history_entry.test_exadb_vm_cluster_update_history_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadb_vm_cluster_update_history_entries", "test_exadb_vm_cluster_update_history_entries", acctest.Optional, acctest.Create, DatabaseExadbVmClusterUpdateHistoryEntryDataSourceRepresentation) +
				DatabaseExadbVmClusterUpdateHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "update_type", "OS_UPDATE"),

				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.0.lifecycle_details"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.0.time_completed"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.0.update_action"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.0.update_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_update_history_entries.0.update_type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadb_vm_cluster_update_history_entries", "test_exadb_vm_cluster_update_history_entries", acctest.Optional, acctest.Create, DatabaseExadbVmClusterUpdateHistoryEntryDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadb_vm_cluster_update_history_entry", "test_exadb_vm_cluster_update_history_entry", acctest.Required, acctest.Create, DatabaseExadbVmClusterUpdateHistoryEntrySingularDataSourceRepresentation) +
				DatabaseExadbVmClusterUpdateHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadb_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_details"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_completed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_history_entry_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_type"),
			),
		},
	})
}
