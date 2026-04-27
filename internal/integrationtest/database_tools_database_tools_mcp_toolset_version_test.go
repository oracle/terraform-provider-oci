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
	DatabaseToolsDatabaseToolsMcpToolsetVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.database_tools_mcp_server_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	DatabaseToolsDatabaseToolsMcpToolsetVersionResourceConfig = ""
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsMcpToolsetVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsMcpToolsetVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	mcpServerId := utils.GetEnvSettingWithBlankDefault("database_tools_mcp_server_ocid")
	mcpServerIdVariableStr := fmt.Sprintf("variable \"database_tools_mcp_server_id\" { default = \"%s\" }\n", mcpServerId)

	datasourceName := "data.oci_database_tools_database_tools_mcp_toolset_versions.test_database_tools_mcp_toolset_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_toolset_versions", "test_database_tools_mcp_toolset_versions", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpToolsetVersionDataSourceRepresentation) +
				compartmentIdVariableStr + mcpServerIdVariableStr + DatabaseToolsDatabaseToolsMcpToolsetVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_mcp_toolset_version_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_toolset_version_collection.0.items.#", "4"),
			),
		},
	})
}
