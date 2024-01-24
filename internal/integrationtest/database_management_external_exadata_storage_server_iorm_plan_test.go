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
	DatabaseManagementDatabaseManagementExternalExadataStorageServerIormPlanSingularDataSourceRepresentation = map[string]interface{}{
		"external_exadata_storage_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.storage_server_id}`},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalExadataStorageServerIormPlanResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalExadataStorageServerIormPlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	storageServerId := utils.GetEnvSettingWithBlankDefault("storage_server_id")
	storageServerIdVariableStr := fmt.Sprintf("variable \"storage_server_id\" { default = \"%s\" }\n", storageServerId)

	singularDatasourceName := "data.oci_database_management_external_exadata_storage_server_iorm_plan.test_external_exadata_storage_server_iorm_plan"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + storageServerIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_server_iorm_plan", "test_external_exadata_storage_server_iorm_plan", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalExadataStorageServerIormPlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_exadata_storage_server_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "db_plan.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_objective"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_status"),
			),
		},
	})
}
