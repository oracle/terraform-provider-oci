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
	DatabaseManagementDatabaseManagementManagedMySqlDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"managed_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id}`},
	}

	DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatabaseManagementManagedMySqlDatabaseResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedMySqlDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedMySqlDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases"
	singularDatasourceName := "data.oci_database_management_managed_my_sql_database.test_managed_my_sql_database"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_my_sql_database_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_databases", "test_managed_my_sql_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_my_sql_database", "test_managed_my_sql_database", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedMySqlDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedMySqlDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_my_sql_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "heat_wave_memory_size"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "heat_wave_node_shape"),
				resource.TestCheckResourceAttr(singularDatasourceName, "heat_wave_nodes.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_heat_wave_active"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_heat_wave_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_lakehouse_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created_heat_wave"),
			),
		},
	})
}
