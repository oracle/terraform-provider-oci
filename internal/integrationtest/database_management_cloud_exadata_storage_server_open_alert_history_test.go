// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementCloudExadataStorageServerOpenAlertHistorySingularDataSourceRepresentation = map[string]interface{}{
		"cloud_exadata_storage_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.storage_server_id}`},
	}

	DatabaseManagementCloudExadataStorageServerOpenAlertHistoryResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_server", "test_cloud_exadata_storage_server", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageServerRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudExadataStorageServerOpenAlertHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudExadataStorageServerOpenAlertHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	storageServerId := utils.GetEnvSettingWithBlankDefault("storage_server_id")
	storageServerIdVariableStr := fmt.Sprintf("variable \"storage_server_id\" { default = \"%s\" }\n", storageServerId)

	singularDatasourceName := "data.oci_database_management_cloud_exadata_storage_server_open_alert_history.test_cloud_exadata_storage_server_open_alert_history"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + storageServerIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_server_open_alert_history", "test_cloud_exadata_storage_server_open_alert_history", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageServerOpenAlertHistorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementCloudExadataStorageServerOpenAlertHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_exadata_storage_server_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "alerts.#", "1"),
			),
		},
	})
}
