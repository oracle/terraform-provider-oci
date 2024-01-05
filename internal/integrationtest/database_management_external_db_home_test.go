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
	DatabaseManagementDatabaseManagementExternalDbHomeSingularDataSourceRepresentation = map[string]interface{}{
		"external_db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_db_homes.test_external_db_homes.external_db_home_collection.0.items.0.id}`},
	}

	DatabaseManagementDatabaseManagementExternalDbHomeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
	}

	DatabaseManagementExternalDbHomeResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDbHomeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDbHomeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	datasourceName := "data.oci_database_management_external_db_homes.test_external_db_homes"
	singularDatasourceName := "data.oci_database_management_external_db_home.test_external_db_home"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_homes", "test_external_db_homes", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbHomeDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbHomeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_system_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "external_db_home_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_homes", "test_external_db_homes", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbHomeDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_home", "test_external_db_home", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbHomeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbHomeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_home_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
