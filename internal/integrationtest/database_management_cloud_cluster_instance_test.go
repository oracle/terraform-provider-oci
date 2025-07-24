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
	DatabaseManagementCloudClusterInstanceRequiredOnlyResource = DatabaseManagementCloudClusterInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster_instance", "test_cloud_cluster_instance", acctest.Required, acctest.Create, DatabaseManagementCloudClusterInstanceRepresentation)

	DatabaseManagementCloudClusterInstanceResourceConfig = DatabaseManagementCloudClusterInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster_instance", "test_cloud_cluster_instance", acctest.Optional, acctest.Update, DatabaseManagementCloudClusterInstanceRepresentation)

	DatabaseManagementCloudClusterInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_cluster_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_cluster_instance.test_cloud_cluster_instance.id}`},
	}

	DatabaseManagementCloudClusterInstanceDataSourceRepresentation = map[string]interface{}{
		"cloud_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_clusters.test_cloud_clusters.cloud_cluster_collection.0.items.0.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseManagementCloudClusterInstanceRepresentation = map[string]interface{}{
		"cloud_cluster_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_cluster_instances.test_cloud_cluster_instances.cloud_cluster_instance_collection.0.items.0.id}`},
		"cloud_connector_id":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_cluster_instances.test_cloud_cluster_instances.cloud_cluster_instance_collection.0.items.0.cloud_connector_id}`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementCloudClusterInstanceResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_clusters", "test_cloud_clusters", acctest.Required, acctest.Create, DatabaseManagementCloudClusterDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_cluster_instances", "test_cloud_cluster_instances", acctest.Required, acctest.Create, DatabaseManagementCloudClusterInstanceDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudClusterInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudClusterInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_cloud_cluster_instance.test_cloud_cluster_instance"
	datasourceName := "data.oci_database_management_cloud_cluster_instances.test_cloud_cluster_instances"
	singularDatasourceName := "data.oci_database_management_cloud_cluster_instance.test_cloud_cluster_instance"

	dbaasDbsystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbaasDbsystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbaasDbsystemId)

	variableStr := compartmentIdVariableStr + dbaasDbsystemIdVariableStr
	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+DatabaseManagementCloudClusterInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster_instance", "test_cloud_cluster_instance", acctest.Optional, acctest.Create, DatabaseManagementCloudClusterInstanceRepresentation), "databasemanagement", "cloudClusterInstance", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + DatabaseManagementCloudClusterInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster_instance", "test_cloud_cluster_instance", acctest.Required, acctest.Create, DatabaseManagementCloudClusterInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_cluster_instance_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Create with optionals
		{
			Config: config + variableStr + DatabaseManagementCloudClusterInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster_instance", "test_cloud_cluster_instance", acctest.Optional, acctest.Create, DatabaseManagementCloudClusterInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_cluster_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "component_name"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + variableStr + DatabaseManagementCloudClusterInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_cluster_instance", "test_cloud_cluster_instance", acctest.Optional, acctest.Update, DatabaseManagementCloudClusterInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_cluster_id"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_cluster_instance_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_cluster_instance_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_cluster_instance", "test_cloud_cluster_instance", acctest.Required, acctest.Create, DatabaseManagementCloudClusterInstanceSingularDataSourceRepresentation) +
				variableStr + DatabaseManagementCloudClusterInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_cluster_instance_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "adr_home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_db_node_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "crs_base_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
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
			Config:            config + DatabaseManagementCloudClusterInstanceRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cloud_cluster_instance_id",
			},
			ResourceName: resourceName,
		},
	})
}
