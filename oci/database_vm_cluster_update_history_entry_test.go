// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vmClusterUpdateHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"update_history_entry_id": Representation{repType: Required, create: `{}`},
		"vm_cluster_id":           Representation{repType: Required, create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	vmClusterUpdateHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": Representation{repType: Required, create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"state":         Representation{repType: Optional, create: `AVAILABLE`},
		"update_type":   Representation{repType: Optional, create: `GI_UPGRADE`},
	}

	VmClusterUpdateHistoryEntryResourceConfig = VmClusterNetworkValidatedResourceConfig +
		generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseVmClusterUpdateHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterUpdateHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterUpdateHistoryEntryResource_basic") {
		t.Skip("test not supported due to GI Update not supported in terraform which is pre-requisite for this test")
	}
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_vm_cluster_update_history_entries.test_vm_cluster_update_history_entries"
	singularDatasourceName := "data.oci_database_vm_cluster_update_history_entry.test_vm_cluster_update_history_entry"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_vm_cluster_update_history_entries", "test_vm_cluster_update_history_entries", Required, Create, vmClusterUpdateHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterUpdateHistoryEntryResourceConfig,
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
				generateDataSourceFromRepresentationMap("oci_database_vm_cluster_update_history_entry", "test_vm_cluster_update_history_entry", Required, Create, vmClusterUpdateHistoryEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterUpdateHistoryEntryResourceConfig,
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
