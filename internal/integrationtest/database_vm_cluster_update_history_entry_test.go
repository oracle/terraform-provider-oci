// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseVmClusterUpdateHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"update_history_entry_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"vm_cluster_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	DatabaseDatabaseVmClusterUpdateHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"state":         acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"update_type":   acctest.Representation{RepType: acctest.Optional, Create: `GI_UPGRADE`},
	}

	DatabaseVmClusterUpdateHistoryEntryResourceConfig = VmClusterNetworkValidatedResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseVmClusterUpdateHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterUpdateHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterUpdateHistoryEntryResource_basic") {
		t.Skip("test not supported due to GI Update not supported in terraform which is pre-requisite for this test")
	}
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_vm_cluster_update_history_entries.test_vm_cluster_update_history_entries"
	singularDatasourceName := "data.oci_database_vm_cluster_update_history_entry.test_vm_cluster_update_history_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_update_history_entries", "test_vm_cluster_update_history_entries", acctest.Required, acctest.Create, DatabaseDatabaseVmClusterUpdateHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseVmClusterUpdateHistoryEntryResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(datasourceName, "update_type", "GI_UPGRADE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_update_history_entries.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_update_history_entries.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_update_history_entries.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_update_history_entries.0.time_completed"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_update_history_entries.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_update_history_entries.0.update_action"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_update_history_entries.0.update_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_update_history_entries.0.update_type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_update_history_entry", "test_vm_cluster_update_history_entry", acctest.Required, acctest.Create, DatabaseDatabaseVmClusterUpdateHistoryEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseVmClusterUpdateHistoryEntryResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_history_entry_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_completed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_type"),
			),
		},
	})
}
