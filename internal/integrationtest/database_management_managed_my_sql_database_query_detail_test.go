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
	DatabaseManagementManagedMySqlDatabaseQueryDetailSingularDataSourceRepresentation = map[string]interface{}{
		"digest":                     acctest.Representation{RepType: acctest.Required, Create: `digest`},
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id}`},
	}

	DatabaseManagementManagedMySqlDatabaseQueryDetailResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseQueryDetailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseQueryDetailResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_management_managed_my_sql_database_query_detail.test_managed_my_sql_database_query_detail"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database_query_detail", "test_managed_my_sql_database_query_detail", acctest.Required, acctest.Create, DatabaseManagementManagedMySqlDatabaseQueryDetailSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseQueryDetailResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "digest", "digest"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "query_explain_plan.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_messages.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_sample_details.#", "1"),
			),
		},
	})
}
