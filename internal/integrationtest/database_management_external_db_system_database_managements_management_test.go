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
	DatabaseManagementExternalDbSystemDatabaseManagementsManagementRepresentation = map[string]interface{}{
		"external_db_system_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
		"license_model":              acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"enable_database_management": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	ExternalDbSystemDatabaseManagementsManagementResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDbSystemDatabaseManagementsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDbSystemDatabaseManagementsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_db_system_database_managements_management.test_external_db_system_database_managements_management"
	parentResourceName := "oci_database_management_external_db_system_database_managements_management.test_external_db_system_database_managements_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+ExternalDbSystemDatabaseManagementsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_database_managements_management", "test_external_db_system_database_managements_management", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDatabaseManagementsManagementRepresentation), "databasemanagement", "externalDbSystemDatabaseManagementsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create and verify enable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_database_managements_management", "test_external_db_system_database_managements_management", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_database_managements_management", "test_external_db_system_database_managements_management", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_database_management", "true"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemDatabaseManagementsManagementResourceDependencies,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_database_managements_management", "test_external_db_system_database_managements_management", acctest.Optional, acctest.Create, DatabaseManagementExternalDbSystemDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
			),
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_database_managements_management", "test_external_db_system_database_managements_management", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + ExternalDbSystemDatabaseManagementsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_database_managements_management", "test_external_db_system_database_managements_management", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemDatabaseManagementsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_database_management", "false"),
			),
		},
	})
}
