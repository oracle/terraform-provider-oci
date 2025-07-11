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
	DatabaseManagementCloudClusterRequiredOnlyResource = DatabaseManagementCloudClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster", "test_cloud_cluster", acctest.Required, acctest.Create, DatabaseManagementCloudClusterRepresentation)

	DatabaseManagementCloudClusterResourceConfig = DatabaseManagementCloudClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster", "test_cloud_cluster", acctest.Optional, acctest.Update, DatabaseManagementCloudClusterRepresentation)

	DatabaseManagementCloudClusterSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_clusters.test_cloud_clusters.cloud_cluster_collection.0.items.0.id}`},
	}

	DatabaseManagementCloudClusterDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dbaas_dbsystem_id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseManagementCloudClusterRepresentation = map[string]interface{}{
		"cloud_cluster_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_clusters.test_cloud_clusters.cloud_cluster_collection.0.items.0.id}`},
		"cloud_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_clusters.test_cloud_clusters.cloud_cluster_collection.0.items.0.cloud_connector_id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseManagementCloudClusterResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_clusters", "test_cloud_clusters", acctest.Required, acctest.Create, DatabaseManagementCloudClusterDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_cloud_cluster.test_cloud_cluster"
	datasourceName := "data.oci_database_management_cloud_clusters.test_cloud_clusters"
	singularDatasourceName := "data.oci_database_management_cloud_cluster.test_cloud_cluster"

	dbaasDbsystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbaasDbsystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbaasDbsystemId)

	variableStr := compartmentIdVariableStr + dbaasDbsystemIdVariableStr
	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+DatabaseManagementCloudClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster", "test_cloud_cluster", acctest.Optional, acctest.Create, DatabaseManagementCloudClusterRepresentation), "databasemanagement", "cloudCluster", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + DatabaseManagementCloudClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster", "test_cloud_cluster", acctest.Required, acctest.Create, DatabaseManagementCloudClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_cluster_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Create with optionals
		{
			Config: config + variableStr + DatabaseManagementCloudClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster", "test_cloud_cluster", acctest.Optional, acctest.Create, DatabaseManagementCloudClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_connector_id"),
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

		// verify datasource
		{
			Config: config + variableStr + DatabaseManagementCloudClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster", "test_cloud_cluster", acctest.Optional, acctest.Update, DatabaseManagementCloudClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_system_id"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_cluster_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_cluster_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_cluster", "test_cloud_cluster", acctest.Required, acctest.Create, DatabaseManagementCloudClusterSingularDataSourceRepresentation) +
				variableStr + DatabaseManagementCloudClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_cluster_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
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
			Config:            config + DatabaseManagementCloudClusterRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cloud_cluster_id",
			},
			ResourceName: resourceName,
		},
	})
}
