// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var databaseVmClusterUpdateHistoryEntryDataSourceRepresentation = map[string]interface{}{
	"vm_cluster_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_id}`},
	"update_history_entry_id": acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_update_history_entry_id}`},
}

// issue-routing-tag: database/default
func TestDatabaseVmClusterUpdateHistoryEntryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterUpdateHistoryEntryResource_basic")
	defer httpreplay.SaveScenario()

	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterUpdateHistoryEntryResource_basic") {
		t.Skip("test requires an existing VM cluster update history entry")
	}
	vmClusterID := utils.GetEnvSettingWithBlankDefault("vm_cluster_id")
	updateHistoryEntryID := utils.GetEnvSettingWithBlankDefault("vm_cluster_update_history_entry_id")
	if vmClusterID == "" || updateHistoryEntryID == "" {
		t.Skip("test requires vm_cluster_id and vm_cluster_update_history_entry_id")
	}

	config := acctest.ProviderTestConfig()
	variableConfig := fmt.Sprintf(`variable "vm_cluster_id" { default = "%s" }
variable "vm_cluster_update_history_entry_id" { default = "%s" }
`, vmClusterID, updateHistoryEntryID)
	datasourceName := "data.oci_database_vm_cluster_update_history_entry.test_vm_cluster_update_history_entry"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_update_history_entry", "test_vm_cluster_update_history_entry", acctest.Required, acctest.Create, databaseVmClusterUpdateHistoryEntryDataSourceRepresentation) +
				variableConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "update_history_entry_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "state"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_completed"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "update_action"),
				resource.TestCheckResourceAttrSet(datasourceName, "update_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "update_type"),
			),
		},
	})
}
