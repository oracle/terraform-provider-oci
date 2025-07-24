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
	DatabaseManagementCloudListenerRequiredOnlyResource = DatabaseManagementCloudListenerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_listener", "test_cloud_listener", acctest.Required, acctest.Create, DatabaseManagementCloudListenerRepresentation)

	DatabaseManagementCloudListenerResourceConfig = DatabaseManagementCloudListenerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_listener", "test_cloud_listener", acctest.Optional, acctest.Update, DatabaseManagementCloudListenerRepresentation)

	DatabaseManagementCloudListenerSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_listener_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_listeners.test_cloud_listeners.cloud_listener_collection.0.items.0.id}`},
	}

	DatabaseManagementCloudListenerDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dbaas_dbsystem_id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatabaseManagementCloudListenerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_listener.test_cloud_listener.id}`}},
	}

	DatabaseManagementCloudListenerRepresentation = map[string]interface{}{
		"cloud_listener_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_listeners.test_cloud_listeners.cloud_listener_collection.0.items.0.id}`},
		"cloud_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_listeners.test_cloud_listeners.cloud_listener_collection.0.items.0.cloud_connector_id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseManagementCloudListenerResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_listeners", "test_cloud_listeners", acctest.Required, acctest.Create, DatabaseManagementCloudListenerDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudListenerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudListenerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"dbaas_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_cloud_listener.test_cloud_listener"
	datasourceName := "data.oci_database_management_cloud_listeners.test_cloud_listeners"
	singularDatasourceName := "data.oci_database_management_cloud_listener.test_cloud_listener"

	variableStr := compartmentIdVariableStr + dbSystemIdVariableStr

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+DatabaseManagementCloudListenerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_listener", "test_cloud_listener", acctest.Optional, acctest.Create, DatabaseManagementCloudListenerRepresentation), "databasemanagement", "cloudListener", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + DatabaseManagementCloudListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_listener", "test_cloud_listener", acctest.Required, acctest.Create, DatabaseManagementCloudListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_listener_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Create with optionals
		{
			Config: config + variableStr + DatabaseManagementCloudListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_listener", "test_cloud_listener", acctest.Optional, acctest.Create, DatabaseManagementCloudListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_listener_id"),
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
			Config: config + variableStr + DatabaseManagementCloudListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_listener", "test_cloud_listener", acctest.Optional, acctest.Update, DatabaseManagementCloudListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_listener_id"),
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
			Config: config + variableStr + DatabaseManagementCloudListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_listener", "test_cloud_listener", acctest.Required, acctest.Create, DatabaseManagementCloudListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "cloud_listener_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + variableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_listener", "test_cloud_listener", acctest.Required, acctest.Create, DatabaseManagementCloudListenerSingularDataSourceRepresentation) +
				DatabaseManagementCloudListenerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_listener_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "adr_home_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_db_node_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_alias"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_ora_location"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_home"),
				resource.TestCheckResourceAttr(singularDatasourceName, "serviced_asms.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "serviced_databases.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trace_directory"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudListenerRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cloud_listener_id",
			},
			ResourceName: resourceName,
		},
	})
}
