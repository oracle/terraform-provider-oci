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
	CloudDbSystemCloudDatabaseManagementsManagementRequiredOnlyResource = CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_database_managements_management", "test_cloud_db_system_cloud_database_managements_management", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementRepresentation)

	DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementRepresentation = map[string]interface{}{
		"cloud_db_system_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.dbaas_dbsystem_id}`},
		"enable_cloud_database_management": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"is_enabled":                       acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		//"metadata":                         acctest.Representation{RepType: acctest.Optional, Create: `{}`, Update: `{}`},
	}

	CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_cloud_db_system_cloud_database_managements_management.test_cloud_db_system_cloud_database_managements_management"
	parentResourceName := "oci_database_management_cloud_db_system_cloud_database_managements_management.test_cloud_db_system_cloud_database_managements_management"
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_database_managements_management", "test_cloud_db_system_cloud_database_managements_management", acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementRepresentation), "databasemanagement", "cloudDbSystemCloudDatabaseManagementsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_database_managements_management", "test_cloud_db_system_cloud_database_managements_management", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_database_managements_management", "test_cloud_db_system_cloud_database_managements_management", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_cloud_database_management", "true"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_database_managements_management", "test_cloud_db_system_cloud_database_managements_management", acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
			),
		},
		// update to disable -> already disabled
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_database_managements_management", "test_cloud_db_system_cloud_database_managements_management", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + CloudDbSystemCloudDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_cloud_database_managements_management", "test_cloud_db_system_cloud_database_managements_management", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_cloud_database_management", "false"),
			),
		},
	})
}
