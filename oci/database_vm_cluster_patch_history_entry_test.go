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
	vmClusterPatchHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"patch_history_entry_id": Representation{RepType: Required, Create: `{}`},
		"vm_cluster_id":          Representation{RepType: Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	vmClusterPatchHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": Representation{RepType: Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	VmClusterPatchHistoryEntryResourceConfig = GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create, vmClusterNetworkRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseVmClusterPatchHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterPatchHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(GetEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterPatchHistoryEntryResource_basic") {
		t.Skip("test not supported due to GI Patching not supported in terraform which is pre-requisite for this test")
	}
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_vm_cluster_patch_history_entries.test_vm_cluster_patch_history_entries"
	singularDatasourceName := "data.oci_database_vm_cluster_patch_history_entry.test_vm_cluster_patch_history_entry"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_patch_history_entries", "test_vm_cluster_patch_history_entries", Required, Create, vmClusterPatchHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterPatchHistoryEntryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.action"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.patch_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.time_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_history_entries.0.time_started"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_patch_history_entry", "test_vm_cluster_patch_history_entry", Required, Create, vmClusterPatchHistoryEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterPatchHistoryEntryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_history_entry_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
	})
}
