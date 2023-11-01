// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	DatabaseManagementDatabaseManagementManagedMySqlDatabaseConfigurationDataDataSourceRepresentation = map[string]interface{}{
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id}`},
	}

	DatabaseManagementManagedMySqlDatabaseConfigurationDataResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseConfigurationDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseConfigurationDataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_my_sql_database_configuration_data.test_managed_my_sql_database_configuration_data"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database_configuration_data", "test_managed_my_sql_database_configuration_data", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseConfigurationDataDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseConfigurationDataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_sql_configuration_data_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "my_sql_configuration_data_collection.0.items.#"),
			),
		},
	})
}
