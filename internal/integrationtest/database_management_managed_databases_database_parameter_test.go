// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseManagementDatabaseManagementManagedDatabasesDatabaseParameterSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"is_allowed_values_included": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: "open_cursors"},
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `CURRENT`},
	}

	DatabaseManagementDatabaseManagementManagedDatabasesDatabaseParameterDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"is_allowed_values_included": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `open_cursors`},
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `CURRENT`},
		"opc_named_credential_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.named_credential_id}`},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesDatabaseParameterResource_basic(t *testing.T) {
	// t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesDatabaseParameterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	datasourceName := "data.oci_database_management_managed_databases_database_parameters.test_managed_databases_database_parameters"
	singularDatasourceName := "data.oci_database_management_managed_databases_database_parameter.test_managed_databases_database_parameter"

	namedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	namedCredentialIdVariableStr := fmt.Sprintf("variable \"named_credential_id\" { default = \"%s\" }\n", namedCredentialId)

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + managedDatabaseIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_database_parameters", "test_managed_databases_database_parameters", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesDatabaseParameterDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.0.database_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.0.database_sub_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.0.database_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.0.database_version"),
				resource.TestCheckResourceAttr(datasourceName, "database_parameters_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + managedDatabaseIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_database_parameter", "test_managed_databases_database_parameter", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesDatabaseParameterSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_sub_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
		// verify datasource with named credential
		{
			Config: config + compartmentIdVariableStr + managedDatabaseIdVariableStr + namedCredentialIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_database_parameters", "test_managed_databases_database_parameters", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesDatabaseParameterDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.0.database_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.0.database_sub_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.0.database_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_parameters_collection.0.database_version"),
				resource.TestCheckResourceAttr(datasourceName, "database_parameters_collection.0.items.#", "1"),
			),
		},
	})
}
