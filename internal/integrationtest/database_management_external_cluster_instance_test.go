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
	DatabaseManagementExternalClusterInstanceRequiredOnlyResource = DatabaseManagementExternalClusterInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster_instance", "test_external_cluster_instance", acctest.Required, acctest.Create, DatabaseManagementExternalClusterInstanceRepresentation)

	DatabaseManagementExternalClusterInstanceResourceConfig = DatabaseManagementExternalClusterInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster_instance", "test_external_cluster_instance", acctest.Optional, acctest.Update, DatabaseManagementExternalClusterInstanceRepresentation)

	DatabaseManagementDatabaseManagementExternalClusterInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"external_cluster_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_cluster_instance.test_external_cluster_instance.id}`},
	}

	DatabaseManagementDatabaseManagementExternalClusterInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"external_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_clusters.test_external_clusters.external_cluster_collection.0.items.0.id}`},
	}
	DatabaseManagementExternalClusterInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_cluster_instance.test_external_cluster_instance.id}`}},
	}

	DatabaseManagementExternalClusterInstanceRepresentation = map[string]interface{}{
		"external_cluster_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_cluster_instances.test_external_cluster_instances.external_cluster_instance_collection.0.items.0.id}`},
		"external_connector_id":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_cluster_instances.test_external_cluster_instances.external_cluster_instance_collection.0.items.0.external_connector_id}`},
	}

	DatabaseManagementExternalClusterInstanceResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_clusters", "test_external_clusters", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalClusterDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_cluster_instances", "test_external_cluster_instances", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalClusterInstanceDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalClusterInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalClusterInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_cluster_instance.test_external_cluster_instance"
	datasourceName := "data.oci_database_management_external_cluster_instances.test_external_cluster_instances"
	singularDatasourceName := "data.oci_database_management_external_cluster_instance.test_external_cluster_instance"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementExternalClusterInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster_instance", "test_external_cluster_instance", acctest.Optional, acctest.Create, DatabaseManagementExternalClusterInstanceRepresentation), "databasemanagement", "externalClusterInstance", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalClusterInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster_instance", "test_external_cluster_instance", acctest.Required, acctest.Create, DatabaseManagementExternalClusterInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_cluster_instance_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalClusterInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_cluster_instance", "test_external_cluster_instance", acctest.Optional, acctest.Update, DatabaseManagementExternalClusterInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_cluster_id"),

				resource.TestCheckResourceAttr(datasourceName, "external_cluster_instance_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_cluster_instance_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_cluster_instance", "test_external_cluster_instance", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalClusterInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalClusterInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_cluster_instance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "adr_home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_node_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "node_role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalClusterInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_cluster_instance_id",
			},
			ResourceName: resourceName,
		},
	})
}
