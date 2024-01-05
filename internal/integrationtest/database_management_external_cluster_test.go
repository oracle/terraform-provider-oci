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
	DatabaseManagementExternalClusterRequiredOnlyResource = DatabaseManagementExternalClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster", "test_external_cluster", acctest.Required, acctest.Create, DatabaseManagementExternalClusterRepresentation)

	DatabaseManagementExternalClusterResourceConfig = DatabaseManagementExternalClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster", "test_external_cluster", acctest.Optional, acctest.Update, DatabaseManagementExternalClusterRepresentation)

	DatabaseManagementDatabaseManagementExternalClusterSingularDataSourceRepresentation = map[string]interface{}{
		"external_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_clusters.test_external_clusters.external_cluster_collection.0.items.0.id}`},
	}

	DatabaseManagementDatabaseManagementExternalClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
	}
	DatabaseManagementExternalClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_cluster.test_external_cluster.id}`}},
	}

	DatabaseManagementExternalClusterRepresentation = map[string]interface{}{
		"external_cluster_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_clusters.test_external_clusters.external_cluster_collection.0.items.0.id}`},
		"external_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_clusters.test_external_clusters.external_cluster_collection.0.items.0.external_connector_id}`},
	}

	DatabaseManagementExternalClusterResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_clusters", "test_external_clusters", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalClusterDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_cluster.test_external_cluster"
	datasourceName := "data.oci_database_management_external_clusters.test_external_clusters"
	singularDatasourceName := "data.oci_database_management_external_cluster.test_external_cluster"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+DatabaseManagementExternalClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster", "test_external_cluster", acctest.Optional, acctest.Create, DatabaseManagementExternalClusterRepresentation), "databasemanagement", "externalCluster", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster", "test_external_cluster", acctest.Required, acctest.Create, DatabaseManagementExternalClusterRepresentation),
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
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster", "test_external_cluster", acctest.Optional, acctest.Update, DatabaseManagementExternalClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_system_id"),

				resource.TestCheckResourceAttr(datasourceName, "external_cluster_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_cluster_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_cluster", "test_external_cluster", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grid_home"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_flex_cluster"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_configurations.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ocr_file_location"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scan_configurations.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vip_configurations.#"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalClusterRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_cluster_id",
			},
			ResourceName: resourceName,
		},
	})
}
