// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentResourceConfig = ""
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	databaseToolsDatabaseApiGatewayConfigId := utils.GetEnvSettingWithDefault("database_tools_cloud_api_gateway_config_id",
		utils.GetEnvSettingWithBlankDefault("database_tools_database_api_gateway_config_id"))
	if databaseToolsDatabaseApiGatewayConfigId == "" {
		t.Skip("set database_tools_cloud_api_gateway_config_id or database_tools_database_api_gateway_config_id to run this test")
	}
	databaseToolsDatabaseApiGatewayConfigIdVariableStr := terraformStringVariable("database_tools_cloud_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId)

	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_content.test_database_tools_database_api_gateway_config_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_content", "test_database_tools_database_api_gateway_config_content", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentSingularDataSourceRepresentation) +
				databaseToolsDatabaseApiGatewayConfigIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId),
			),
		},
	})
}
