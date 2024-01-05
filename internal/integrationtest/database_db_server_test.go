// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseDbServerSingularDataSourceRepresentation = map[string]interface{}{
		"db_server_id":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
	}

	DatabaseDatabaseDbServerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DatabaseDatabasePeerExaInfraDbServerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.peer_exadata_infrastructure.id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DatabaseDbServerResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseDbServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbServerResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_servers.test_db_servers"
	singularDatasourceName := "data.oci_database_db_server.test_db_server"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, DatabaseDatabaseDbServerDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseDbServerResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.#"),
					resource.TestCheckResourceAttr(datasourceName, "db_servers.0.autonomous_virtual_machine_ids.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "db_servers.0.autonomous_vm_cluster_ids.#", "0"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.cpu_core_count"),
					resource.TestCheckResourceAttr(datasourceName, "db_servers.0.db_node_ids.#", "0"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.db_node_storage_size_in_gbs"),
					resource.TestCheckResourceAttr(datasourceName, "db_servers.0.db_server_patching_details.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.exadata_infrastructure_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.max_cpu_count"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.max_db_node_storage_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.max_memory_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.memory_size_in_gbs"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "db_servers.0.vm_cluster_ids.#", "0"),
				),
			},
			//verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, DatabaseDatabaseDbServerDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_server", "test_db_server", acctest.Required, acctest.Create, DatabaseDatabaseDbServerSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseDbServerResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_server_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_virtual_machine_ids.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "autonomous_vm_cluster_ids.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_core_count"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_node_ids.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_storage_size_in_gbs"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_server_patching_details.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "max_cpu_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "max_db_node_storage_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "max_memory_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_size_in_gbs"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "vm_cluster_ids.#", "0"),
				),
			},
		},
	})
}
