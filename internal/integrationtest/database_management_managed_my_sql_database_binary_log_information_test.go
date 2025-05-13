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
	DatabaseManagementManagedMySqlDatabaseBinaryLogInformationSingularDataSourceRepresentation = map[string]interface{}{
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id}`},
	}

	DatabaseManagementManagedMySqlDatabaseBinaryLogInformationResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseBinaryLogInformationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseBinaryLogInformationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_management_managed_my_sql_database_binary_log_information.test_managed_my_sql_database_binary_log_information"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database_binary_log_information", "test_managed_my_sql_database_binary_log_information", acctest.Required, acctest.Create, DatabaseManagementManagedMySqlDatabaseBinaryLogInformationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseBinaryLogInformationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "binary_log_compression"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "binary_log_compression_percent"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "binary_log_format"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "binary_log_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "binary_log_position"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "binary_logging"),
			),
		},
	})
}
