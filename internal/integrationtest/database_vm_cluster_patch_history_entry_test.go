// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	vmClusterPatchHistoryEntrySingularDataSourceRepresentation = map[string]interface{}{
		"patch_history_entry_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"vm_cluster_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	vmClusterPatchHistoryEntryDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	VmClusterPatchHistoryEntryResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, exadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, vmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, vmClusterRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseVmClusterPatchHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterPatchHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterPatchHistoryEntryResource_basic") {
		t.Skip("test not supported due to GI Patching not supported in terraform which is pre-requisite for this test")
	}
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_vm_cluster_patch_history_entries.test_vm_cluster_patch_history_entries"
	singularDatasourceName := "data.oci_database_vm_cluster_patch_history_entry.test_vm_cluster_patch_history_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_patch_history_entries", "test_vm_cluster_patch_history_entries", acctest.Required, acctest.Create, vmClusterPatchHistoryEntryDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterPatchHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_patch_history_entry", "test_vm_cluster_patch_history_entry", acctest.Required, acctest.Create, vmClusterPatchHistoryEntrySingularDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterPatchHistoryEntryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
