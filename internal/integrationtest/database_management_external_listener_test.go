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
	DatabaseManagementExternalListenerRequiredOnlyResource = DatabaseManagementExternalListenerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_listener", "test_external_listener", acctest.Required, acctest.Create, DatabaseManagementExternalListenerRepresentation)

	DatabaseManagementExternalListenerResourceConfig = DatabaseManagementExternalListenerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_listener", "test_external_listener", acctest.Optional, acctest.Update, DatabaseManagementExternalListenerRepresentation)

	DatabaseManagementDatabaseManagementExternalListenerSingularDataSourceRepresentation = map[string]interface{}{
		"external_listener_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_listeners.test_external_listeners.external_listener_collection.0.items.0.id}`},
	}

	DatabaseManagementDatabaseManagementExternalListenerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
	}
	DatabaseManagementExternalListenerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_listener.test_external_listener.id}`}},
	}

	DatabaseManagementExternalListenerRepresentation = map[string]interface{}{
		"external_listener_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_listeners.test_external_listeners.external_listener_collection.0.items.0.id}`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"external_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_listeners.test_external_listeners.external_listener_collection.0.items.0.external_connector_id}`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementExternalListenerResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_listeners", "test_external_listeners", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalListenerDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalListenerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalListenerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_listener.test_external_listener"
	datasourceName := "data.oci_database_management_external_listeners.test_external_listeners"
	singularDatasourceName := "data.oci_database_management_external_listener.test_external_listener"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+DatabaseManagementExternalListenerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_listener", "test_external_listener", acctest.Optional, acctest.Create, DatabaseManagementExternalListenerRepresentation), "databasemanagement", "externalListener", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_listener", "test_external_listener", acctest.Required, acctest.Create, DatabaseManagementExternalListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_listener_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_listener", "test_external_listener", acctest.Optional, acctest.Create, DatabaseManagementExternalListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "component_name"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "external_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_listener_id"),
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
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_listener", "test_external_listener", acctest.Optional, acctest.Update, DatabaseManagementExternalListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_system_id"),

				resource.TestCheckResourceAttr(datasourceName, "external_listener_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_listener_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_listener", "test_external_listener", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalListenerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalListenerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_listener_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "endpoints.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_node_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_alias"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_ora_location"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_home"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "serviced_asms.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "serviced_databases.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trace_directory"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalListenerRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_listener_id",
			},
			ResourceName: resourceName,
		},
	})
}
