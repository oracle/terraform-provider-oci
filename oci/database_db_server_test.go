// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dbServerSingularDataSourceRepresentation = map[string]interface{}{
		"db_server_id":              Representation{RepType: Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`},
		"exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
	}

	dbServerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{RepType: Required, Create: `${var.compartment_id}`},
		"exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"display_name":              Representation{RepType: Optional, Create: `displayName`},
		"state":                     Representation{RepType: Optional, Create: `AVAILABLE`},
	}

	DbServerResourceConfig = GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create, exadataInfrastructureRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseDbServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbServerResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_servers.test_db_servers"
	singularDatasourceName := "data.oci_database_db_server.test_db_server"

	SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
					compartmentIdVariableStr + DbServerResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.cpu_core_count"),
					resource.TestCheckResourceAttr(datasourceName, "db_servers.0.db_node_ids.#", "0"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_servers.0.db_node_storage_size_in_gbs"),
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
					GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
					GenerateDataSourceFromRepresentationMap("oci_database_db_server", "test_db_server", Required, Create, dbServerSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DbServerResourceConfig,

				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_server_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_core_count"),
					resource.TestCheckResourceAttr(singularDatasourceName, "db_node_ids.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_storage_size_in_gbs"),
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
