// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementCloudDbNodeRequiredOnlyResource = DatabaseManagementCloudDbNodeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_node", "test_cloud_db_node", acctest.Required, acctest.Create, DatabaseManagementCloudDbNodeRepresentation)

	DatabaseManagementCloudDbNodeResourceConfig = DatabaseManagementCloudDbNodeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_node", "test_cloud_db_node", acctest.Optional, acctest.Update, DatabaseManagementCloudDbNodeRepresentation)

	DatabaseManagementCloudDbNodeSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_db_nodes.test_cloud_db_nodes.cloud_db_node_collection.0.items.0.id}`},
	}

	DatabaseManagementCloudDbNodeDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dbaas_dbsystem_id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseManagementCloudDbNodeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_db_node.test_cloud_db_node.id}`}},
	}

	DatabaseManagementCloudDbNodeRepresentation = map[string]interface{}{
		"cloud_db_node_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_db_nodes.test_cloud_db_nodes.cloud_db_node_collection.0.items.0.id}`},
		"cloud_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_db_nodes.test_cloud_db_nodes.cloud_db_node_collection.0.items.0.cloud_connector_id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseManagementCloudDbNodeResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_nodes", "test_cloud_db_nodes", acctest.Required, acctest.Create, DatabaseManagementCloudDbNodeDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudDbNodeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudDbNodeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_cloud_db_node.test_cloud_db_node"
	datasourceName := "data.oci_database_management_cloud_db_nodes.test_cloud_db_nodes"
	singularDatasourceName := "data.oci_database_management_cloud_db_node.test_cloud_db_node"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+DatabaseManagementCloudDbNodeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_node", "test_cloud_db_node", acctest.Optional, acctest.Create, DatabaseManagementCloudDbNodeRepresentation), "databasemanagement", "cloudDbNode", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_node", "test_cloud_db_node", acctest.Required, acctest.Create, DatabaseManagementCloudDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_node_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbNodeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_node", "test_cloud_db_node", acctest.Optional, acctest.Create, DatabaseManagementCloudDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_node_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "component_name"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_node", "test_cloud_db_node", acctest.Optional, acctest.Update, DatabaseManagementCloudDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_node_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "component_name"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_node", "test_cloud_db_node", acctest.Required, acctest.Create, DatabaseManagementCloudDbNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_db_node_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_node_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + dbSystemIdVariableStr + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_node", "test_cloud_db_node", acctest.Required, acctest.Create, DatabaseManagementCloudDbNodeSingularDataSourceRepresentation) +
				DatabaseManagementCloudDbNodeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_db_node_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dbaas_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudDbNodeRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cloud_db_node_id",
			},
			ResourceName: resourceName,
		},
	})
}
