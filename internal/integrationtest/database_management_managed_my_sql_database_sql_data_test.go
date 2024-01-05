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
	DatabaseManagementDatabaseManagementManagedMySqlDatabaseSqlDataDataSourceRepresentation = map[string]interface{}{
		"end_time":                   acctest.Representation{RepType: acctest.Required, Create: `${replace(timestamp(), "/Z/", ".000Z")}`},
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id}`},
		"start_time":                 acctest.Representation{RepType: acctest.Required, Create: `${replace(timeadd(timestamp(), "-2h"), "/Z/", ".000Z")}`},
		"filter_column":              acctest.Representation{RepType: acctest.Required, Create: `COUNT_STAR`},
	}

	DatabaseManagementManagedMySqlDatabaseSqlDataResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseSqlDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseSqlDataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_my_sql_database_sql_data.test_managed_my_sql_database_sql_data"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database_sql_data", "test_managed_my_sql_database_sql_data", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseSqlDataDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseSqlDataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "filter_column", "COUNT_STAR"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_sql_data_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_sql_data_collection.0.items.#"),
			),
			// Non empty plan expected because the data source input relies on interpolation syntax
			ExpectNonEmptyPlan: true,
		},
	})
}
