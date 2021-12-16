// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	dbNodeSingularDataSourceRepresentation = map[string]interface{}{
		"db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}`},
	}

	dbNodeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vm_cluster_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id}`},
	}

	DbNodeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_vm_cluster", "test_cloud_vm_cluster", acctest.Required, acctest.Create, cloudVmClusterRepresentation) +
		AvailabilityDomainConfig +
		CloudVmClusterResourceDependencies +
		DefinedTagsDependencies
)

func TestDatabaseDbNodeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_db_nodes.test_db_nodes"
	singularDatasourceName := "data.oci_database_db_node.test_db_node"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_nodes", "test_db_nodes", acctest.Required, acctest.Create, dbNodeDataSourceRepresentation) +
					compartmentIdVariableStr + DbNodeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.backup_ip_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.backup_vnic2id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.host_ip_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "db_nodes.0.vnic2id"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node", "test_db_node", acctest.Required, acctest.Create, dbNodeSingularDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_nodes", "test_db_nodes", acctest.Required, acctest.Create, dbNodeDataSourceRepresentation) +
					compartmentIdVariableStr + DbNodeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_ip_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_vnic2id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "host_ip_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vnic2id"),
				),
			},
		},
	})
}
