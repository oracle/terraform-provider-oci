// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseExadbVmClusterDbNodeSingularDataSourceRepresentation = map[string]interface{}{
		"db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}`},
	}

	DatabaseExadbVmClusterDbNodeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vm_cluster_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.exadb_vm_cluster_id}`},
	}

	// Note: set env variable TF_VAR_exadb_vm_cluster_id before running this test
	DatabaseExadbVmClusterDbNodeResourceConfig = `variable "exadb_vm_cluster_id" {}`
)

// issue-routing-tag: database/ExaCS
func TestDatabaseExaDbVmClusterDbNodeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaDbVmClusterDbNodeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_nodes.test_db_nodes"
	singularDatasourceName := "data.oci_database_db_node.test_db_node"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_nodes", "test_db_nodes", acctest.Optional, acctest.Create, DatabaseExadbVmClusterDbNodeDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExadbVmClusterDbNodeResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.cpu_core_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.db_node_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.db_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.hostname"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.total_cpu_core_count"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node", "test_db_node", acctest.Optional, acctest.Create, DatabaseExadbVmClusterDbNodeSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_nodes", "test_db_nodes", acctest.Optional, acctest.Create, DatabaseExadbVmClusterDbNodeDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExadbVmClusterDbNodeResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_core_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_storage_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hostname"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_cpu_core_count"),
			),
		},
	})
}
