// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalDbHomeRequiredOnlyResource = DatabaseManagementExternalDbHomeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_home", "test_external_db_home", acctest.Required, acctest.Create, DatabaseManagementExternalDbHomeRepresentation)

	DatabaseManagementExternalDbHomeResourceConfig = DatabaseManagementExternalDbHomeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_home", "test_external_db_home", acctest.Optional, acctest.Update, DatabaseManagementExternalDbHomeRepresentation)

	DatabaseManagementDatabaseManagementExternalDbHomeSingularDataSourceRepresentation = map[string]interface{}{
		"external_db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_db_homes.test_external_db_homes.external_db_home_collection.0.items.0.id}`},
	}

	DatabaseManagementDatabaseManagementExternalDbHomeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
	}

	DatabaseManagementExternalDbHomeRepresentation = map[string]interface{}{
		"external_db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_db_homes.test_external_db_homes.external_db_home_collection.0.items.0.id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementExternalDbHomeResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_homes", "test_external_db_homes", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbHomeDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDbHomeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDbHomeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_external_db_home.test_external_db_home"
	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	datasourceName := "data.oci_database_management_external_db_homes.test_external_db_homes"
	singularDatasourceName := "data.oci_database_management_external_db_home.test_external_db_home"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+DatabaseManagementExternalDbHomeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_home", "test_external_db_home", acctest.Optional, acctest.Create, DatabaseManagementExternalDbHomeRepresentation), "databasemanagement", "externalDbHome", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbHomeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_home", "test_external_db_home", acctest.Optional, acctest.Create, DatabaseManagementExternalDbHomeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "external_db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbHomeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_home", "test_external_db_home", acctest.Optional, acctest.Update, DatabaseManagementExternalDbHomeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_system_id"),

				resource.TestCheckResourceAttr(datasourceName, "external_db_home_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_home_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_home", "test_external_db_home", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbHomeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbHomeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_home_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalDbHomeRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_db_home_id",
			},
			ResourceName: resourceName,
		},
	})
}
