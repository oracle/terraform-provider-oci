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
	DatabaseManagementExternalDbNodeRequiredOnlyResource = DatabaseManagementExternalDbNodeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_node", "test_external_db_node", acctest.Required, acctest.Create, DatabaseManagementExternalDbNodeRepresentation)

	DatabaseManagementExternalDbNodeResourceConfig = DatabaseManagementExternalDbNodeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_node", "test_external_db_node", acctest.Optional, acctest.Update, DatabaseManagementExternalDbNodeRepresentation)

	DatabaseManagementDatabaseManagementExternalDbNodeSingularDataSourceRepresentation = map[string]interface{}{
		"external_db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_db_nodes.test_external_db_nodes.external_db_node_collection.0.items.0.id}`},
	}

	DatabaseManagementDatabaseManagementExternalDbNodeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
	}
	DatabaseManagementExternalDbNodeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_db_node.test_external_db_node.id}`}},
	}

	DatabaseManagementExternalDbNodeRepresentation = map[string]interface{}{
		"external_db_node_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_db_nodes.test_external_db_nodes.external_db_node_collection.0.items.1.id}`},
		"external_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_db_nodes.test_external_db_nodes.external_db_node_collection.0.items.1.external_connector_id}`},
	}

	DatabaseManagementExternalDbNodeResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_nodes", "test_external_db_nodes", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbNodeDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDbNodeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDbNodeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_db_node.test_external_db_node"
	datasourceName := "data.oci_database_management_external_db_nodes.test_external_db_nodes"
	singularDatasourceName := "data.oci_database_management_external_db_node.test_external_db_node"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+DatabaseManagementExternalDbNodeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_node", "test_external_db_node", acctest.Optional, acctest.Create, DatabaseManagementExternalDbNodeRepresentation), "databasemanagement", "externalDbNode", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_node", "test_external_db_node", acctest.Required, acctest.Create, DatabaseManagementExternalDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_node", "test_external_db_node", acctest.Optional, acctest.Update, DatabaseManagementExternalDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_system_id"),

				resource.TestCheckResourceAttr(datasourceName, "external_db_node_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_node_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_node", "test_external_db_node", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbNodeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbNodeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_node_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalDbNodeRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_db_node_id",
			},
			ResourceName: resourceName,
		},
	})
}
