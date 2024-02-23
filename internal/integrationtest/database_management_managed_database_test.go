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
	DatabaseManagementDatabaseManagementManagedDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"deployment_type":                    acctest.Representation{RepType: acctest.Optional, Create: `ONPREMISE`},
		"external_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id}`},
		"id":                                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"management_option":                  acctest.Representation{RepType: acctest.Optional, Create: `BASIC`},
		"name":                               acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	DatabaseManagementManagedDatabaseResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	testManagedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	testManagedDatabaseIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", testManagedDatabaseId)

	datasourceName := "data.oci_database_management_managed_databases.test_managed_databases"
	singularDatasourceName := "data.oci_database_management_managed_database.test_managed_database"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementManagedDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.0.items.0.deployment_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.0.items.0.management_option"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database", "test_managed_database", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + DatabaseManagementManagedDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_sub_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cluster"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_groups.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_option"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
