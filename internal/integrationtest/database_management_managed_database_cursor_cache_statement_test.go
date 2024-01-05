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
	DatabaseManagementDatabaseManagementManagedDatabaseCursorCacheStatementDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"sql_text":            acctest.Representation{RepType: acctest.Optional, Create: `sqlText`},
		"limit":               acctest.Representation{RepType: acctest.Required, Create: `1000`},
	}

	DatabaseManagementManagedDatabaseCursorCacheStatementResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseCursorCacheStatementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseCursorCacheStatementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDbId := utils.GetEnvSettingWithBlankDefault("test_managed_database_id")
	managedDbIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", managedDbId)

	datasourceName := "data.oci_database_management_managed_database_cursor_cache_statements.test_managed_database_cursor_cache_statements"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_cursor_cache_statements", "test_managed_database_cursor_cache_statements", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseCursorCacheStatementDataSourceRepresentation) +
				compartmentIdVariableStr + managedDbIdVariableStr + DatabaseManagementManagedDatabaseCursorCacheStatementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "cursor_cache_statement_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "cursor_cache_statement_collection.0.items.#"),
			),
		},
	})
}
