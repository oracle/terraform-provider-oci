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
	DatabaseManagementDatabaseManagementExternalExadataStorageGridSingularDataSourceRepresentation = map[string]interface{}{
		"external_exadata_storage_grid_id": acctest.Representation{RepType: acctest.Required, Create: `${var.storage_server_grid}`},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalExadataStorageGridResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalExadataStorageGridResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	storageServerGridIdUri := utils.GetEnvSettingWithBlankDefault("storage_server_grid")
	storageServerGridIdUriVariableStr := fmt.Sprintf("variable \"storage_server_grid\" { default = \"%s\" }\n", storageServerGridIdUri)

	singularDatasourceName := "data.oci_database_management_external_exadata_storage_grid.test_external_exadata_storage_grid"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + storageServerGridIdUriVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_grid", "test_external_exadata_storage_grid", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalExadataStorageGridSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_exadata_storage_grid_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "server_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_servers.#", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
