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
	DatabaseManagementDatabaseManagementExternalDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
		"external_database_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_external_database.test_external_database.id}`},
	}

	DatabaseManagementExternalDatabaseResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	datasourceName := "data.oci_database_management_external_databases.test_external_databases"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_databases", "test_external_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_collection.0.items.#"),
			),
		},
	})
}
