// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDatabaseVmClusterPatchSingularDataSourceRepresentation = map[string]interface{}{
		"patch_id":      acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	DatabaseDatabaseVmClusterPatchDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	DatabaseVmClusterPatchResourceConfig = VmClusterNetworkValidatedResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseVmClusterPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterPatchResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterPatchResource_basic") {
		t.Skip("test not supported due to GI Patching not supported in terraform which is pre-requisite for this test")
	}
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_vm_cluster_patches.test_vm_cluster_patches"
	singularDatasourceName := "data.oci_database_vm_cluster_patch.test_vm_cluster_patch"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_patches", "test_vm_cluster_patches", acctest.Required, acctest.Create, DatabaseDatabaseVmClusterPatchDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseVmClusterPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "patches.#"),
				resource.TestCheckResourceAttr(datasourceName, "patches.0.available_actions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.last_action"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.time_released"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.version"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_patch", "test_vm_cluster_patch", acctest.Required, acctest.Create, DatabaseDatabaseVmClusterPatchSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseVmClusterPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "available_actions.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
	})
}
