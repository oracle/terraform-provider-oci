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
	DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementRepresentation = map[string]interface{}{
		"connector_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.external_my_sql_database_connector_id}`},
		"external_my_sql_database_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"enable_external_mysql_database": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	external_my_sql_database_connector_id := utils.GetEnvSettingWithBlankDefault("external_my_sql_database_connector_id")
	externalMySqlDatabaseconnectorIdVariableStr := fmt.Sprintf("variable \"external_my_sql_database_connector_id\" { default = \"%s\" }\n", external_my_sql_database_connector_id)

	test_managed_database_id := utils.GetEnvSettingWithBlankDefault("test_managed_database_id")
	testManagedDatabaseIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", test_managed_database_id)

	resourceName := "oci_database_management_external_my_sql_database_external_mysql_databases_management.test_external_my_sql_database_external_mysql_databases_management"
	parentResourceName := "oci_database_management_external_my_sql_database_external_mysql_databases_management.test_external_my_sql_database_external_mysql_databases_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+testManagedDatabaseIdVariableStr+externalMySqlDatabaseconnectorIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_external_mysql_databases_management", "test_external_my_sql_database_external_mysql_databases_management", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementRepresentation), "databasemanagement", "externalMySqlDatabaseExternalMysqlDatabasesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + testManagedDatabaseIdVariableStr + externalMySqlDatabaseconnectorIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_external_mysql_databases_management", "test_external_my_sql_database_external_mysql_databases_management", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_my_sql_database_id"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + testManagedDatabaseIdVariableStr + externalMySqlDatabaseconnectorIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_external_mysql_databases_management", "test_external_my_sql_database_external_mysql_databases_management", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_external_mysql_database", "true"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + testManagedDatabaseIdVariableStr + externalMySqlDatabaseconnectorIdVariableStr,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + testManagedDatabaseIdVariableStr + externalMySqlDatabaseconnectorIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_external_mysql_databases_management", "test_external_my_sql_database_external_mysql_databases_management", acctest.Optional, acctest.Create, DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_my_sql_database_id"),
			),
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + testManagedDatabaseIdVariableStr + externalMySqlDatabaseconnectorIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_external_mysql_databases_management", "test_external_my_sql_database_external_mysql_databases_management", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_my_sql_database_id"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + testManagedDatabaseIdVariableStr + externalMySqlDatabaseconnectorIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_external_mysql_databases_management", "test_external_my_sql_database_external_mysql_databases_management", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_external_mysql_database", "false"),
			),
		},
	})
}
