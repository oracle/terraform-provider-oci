// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseVmClusterUpdateSingularDataSourceRepresentation = map[string]interface{}{
		"update_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_update_id}`},
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	DatabaseDatabaseVmClusterUpdateDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseVmClusterUpdateDataSourceFilterRepresentation},
	}
	DatabaseVmClusterUpdateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.vm_cluster_update_id}`}},
	}

	DatabaseVmClusterUpdateResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, DatabaseDatabaseDbServerDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithNewProperties(
				acctest.RepresentationCopyWithRemovedProperties(exadataInfrastructureActivateRepresentation, []string{"defined_tags"}),
				map[string]interface{}{
					"activation_file":    acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
					"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
				},
			)) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithRemovedProperties(vmClusterNetworkValidateRepresentation, []string{"defined_tags"})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithRemovedProperties(DatabaseVmClusterRepresentation, []string{"cloud_automation_update_details", "defined_tags"}))
)

// issue-routing-tag: database/default
func TestDatabaseVmClusterUpdateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterUpdateResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(utils.GetEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterUpdateResource_basic") {
		t.Skip("test not supported due to GI Update not supported in terraform which is pre-requisite for this test")
	}
	updateId := utils.GetEnvSettingWithBlankDefault("vm_cluster_update_id")
	if updateId == "" {
		t.Skip("test requires vm_cluster_update_id")
	}
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	updateIdVariableStr := fmt.Sprintf("variable \"vm_cluster_update_id\" { default = \"%s\" }\n", updateId)

	datasourceName := "data.oci_database_vm_cluster_updates.test_vm_cluster_updates"
	singularDatasourceName := "data.oci_database_vm_cluster_update.test_vm_cluster_update"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_updates", "test_vm_cluster_updates", acctest.Required, acctest.Create, DatabaseDatabaseVmClusterUpdateDataSourceRepresentation) +
				compartmentIdVariableStr + updateIdVariableStr + DatabaseVmClusterUpdateResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.#"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_updates.0.id", updateId),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.available_actions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.available_update_modes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.oracle_linux_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.time_released"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.update_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.version"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_update", "test_vm_cluster_update", acctest.Required, acctest.Create, DatabaseDatabaseVmClusterUpdateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + updateIdVariableStr + DatabaseVmClusterUpdateResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(singularDatasourceName, "update_id", updateId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_actions.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_update_modes.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_linux_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
	})
}
