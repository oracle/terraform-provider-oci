// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementCloudDbHomeRequiredOnlyResource = DatabaseManagementCloudDbHomeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_home", "test_cloud_db_home", acctest.Required, acctest.Create, DatabaseManagementCloudDbHomeRepresentation)

	DatabaseManagementCloudDbHomeResourceConfig = DatabaseManagementCloudDbHomeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_home", "test_cloud_db_home", acctest.Optional, acctest.Update, DatabaseManagementCloudDbHomeRepresentation)

	DatabaseManagementCloudDbHomeSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_db_homes.test_cloud_db_homes.cloud_db_home_collection.0.items.0.id}`},
	}

	DatabaseManagementCloudDbHomeDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dbaas_dbsystem_id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseManagementCloudDbHomeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_db_home.test_cloud_db_home.id}`}},
	}

	DatabaseManagementCloudDbHomeRepresentation = map[string]interface{}{
		"cloud_db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_db_homes.test_cloud_db_homes.cloud_db_home_collection.0.items.0.id}`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementCloudDbHomeResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_homes", "test_cloud_db_homes", acctest.Optional, acctest.Create, DatabaseManagementCloudDbHomeDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudDbHomeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudDbHomeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_cloud_db_home.test_cloud_db_home"
	datasourceName := "data.oci_database_management_cloud_db_homes.test_cloud_db_homes"
	singularDatasourceName := "data.oci_database_management_cloud_db_home.test_cloud_db_home"

	dbaasDbsystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbaasDbsystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbaasDbsystemId)

	variableStr := compartmentIdVariableStr + dbaasDbsystemIdVariableStr
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+DatabaseManagementCloudDbHomeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_home", "test_cloud_db_home", acctest.Required, acctest.Create, DatabaseManagementCloudDbHomeRepresentation), "databasemanagement", "cloudDbHome", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify Create
		{
			Config: config + variableStr + DatabaseManagementCloudDbHomeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_home", "test_cloud_db_home", acctest.Required, acctest.Create, DatabaseManagementCloudDbHomeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "5"),
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
		// verify Create with optionals
		{
			Config: config + variableStr + DatabaseManagementCloudDbHomeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_home", "test_cloud_db_home", acctest.Optional, acctest.Create, DatabaseManagementCloudDbHomeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "5"),
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

		// verify updates to updatable parameters
		{
			Config: config + variableStr + DatabaseManagementCloudDbHomeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_home", "test_cloud_db_home", acctest.Optional, acctest.Update, DatabaseManagementCloudDbHomeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
			),
		},

		// verify datasource
		{
			Config: config + variableStr + DatabaseManagementCloudDbHomeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_home", "test_cloud_db_home", acctest.Required, acctest.Create, DatabaseManagementCloudDbHomeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_home_collection.0.items.0.display_name"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_db_home_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_db_home_collection.0.items.#", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config + variableStr + DatabaseManagementCloudDbHomeResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_home", "test_cloud_db_home", acctest.Required, acctest.Create, DatabaseManagementCloudDbHomeSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_db_home_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "5"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudDbHomeRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cloud_db_home_id",
			},
			ResourceName: resourceName,
		},
	})
}
