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
	DatabaseExadbVmClusterUpdateSingularDataSourceRepresentation = map[string]interface{}{
		"exadb_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.exadb_vm_cluster_id}`},
		"update_id":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_exadb_vm_cluster_updates.test_exadb_vm_cluster_updates.exadb_vm_cluster_updates[0].id}`},
	}

	DatabaseExadbVmClusterUpdateDataSourceRepresentation = map[string]interface{}{
		"exadb_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.exadb_vm_cluster_id}`},
		"update_type":         acctest.Representation{RepType: acctest.Optional, Create: `GI_PATCH`},
	}

	// Note: set env variable TF_VAR_exadb_vm_cluster_id before running this test
	DatabaseExadbVmClusterUpdateResourceConfig = `variable "exadb_vm_cluster_id" {}`
)

// issue-routing-tag: database/default
func TestDatabaseExadbVmClusterUpdateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadbVmClusterUpdateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_database_exadb_vm_cluster_updates.test_exadb_vm_cluster_updates"
	singularDatasourceName := "data.oci_database_exadb_vm_cluster_update.test_exadb_vm_cluster_update"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadb_vm_cluster_updates", "test_exadb_vm_cluster_updates", acctest.Optional, acctest.Create, DatabaseExadbVmClusterUpdateDataSourceRepresentation) +
				DatabaseExadbVmClusterUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "update_type", "GI_PATCH"),

				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_updates.#"),
				resource.TestCheckResourceAttr(datasourceName, "exadb_vm_cluster_updates.0.available_actions.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_updates.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_updates.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_updates.0.time_released"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_updates.0.update_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadb_vm_cluster_updates.0.version"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadb_vm_cluster_updates", "test_exadb_vm_cluster_updates", acctest.Optional, acctest.Create, DatabaseExadbVmClusterUpdateDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_exadb_vm_cluster_update", "test_exadb_vm_cluster_update", acctest.Required, acctest.Create, DatabaseExadbVmClusterUpdateSingularDataSourceRepresentation) +
				DatabaseExadbVmClusterUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadb_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "available_actions.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
	})
}
