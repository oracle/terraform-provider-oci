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
	DatabaseManagementDatabaseManagementExternalExadataStorageServerSingularDataSourceRepresentation = map[string]interface{}{
		"external_exadata_storage_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.storage_server_id}`},
	}

	DatabaseManagementExternalExadataStorageServerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_exadata_infra_id}`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalExadataStorageServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalExadataStorageServerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	storageServerId := utils.GetEnvSettingWithBlankDefault("storage_server_id")
	storageServerIdVariableStr := fmt.Sprintf("variable \"storage_server_id\" { default = \"%s\" }\n", storageServerId)

	connectorExadataInfraId := utils.GetEnvSettingWithBlankDefault("connector_exadata_infra_id")
	connectorExadataInfraIdVariableStr := fmt.Sprintf("variable \"connector_exadata_infra_id\" { default = \"%s\" }\n", connectorExadataInfraId)

	datasourceName := "data.oci_database_management_external_exadata_storage_servers.test_external_exadata_storage_servers"
	singularDatasourceName := "data.oci_database_management_external_exadata_storage_server.test_external_exadata_storage_server"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + storageServerIdVariableStr + connectorExadataInfraIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_servers", "test_external_exadata_storage_servers", acctest.Required, acctest.Create, DatabaseManagementExternalExadataStorageServerDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "external_exadata_infrastructure_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "external_exadata_storage_server_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + storageServerIdVariableStr + connectorExadataInfraIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_server", "test_external_exadata_storage_server", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalExadataStorageServerSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_exadata_storage_server_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "connector.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "internal_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "make_model"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_flash_disk_iops"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_flash_disk_throughput"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_hard_disk_iops"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "max_hard_disk_throughput"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_gb"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
	})
}
