// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var (
	DatabaseToolsDatabaseToolsConnectionRequiredOnlyResource = DatabaseToolsDatabaseToolsConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionRepresentation)

	DatabaseToolsDatabaseToolsConnectionResourceConfig = DatabaseToolsDatabaseToolsConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionRepresentation)

	DatabaseToolsDatabaseToolsDatabaseToolsConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
	}

	DatabaseToolsDatabaseToolsDatabaseToolsConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `tf_connection_name`, Update: `displayName2`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_DATABASE`}},
		"runtime_support": acctest.Representation{RepType: acctest.Optional, Create: []string{`SUPPORTED`}},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionDataSourceFilterRepresentation}}

	DatabaseToolsDatabaseToolsConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`}},
	}

	DatabaseToolsUsernameProxyClientRepresentation = map[string]interface{}{
		"proxy_authentication_type": acctest.Representation{RepType: acctest.Required, Create: `USER_NAME`},
		"user_name":                 acctest.Representation{RepType: acctest.Required, Create: `johndoe`},
		"user_password":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation},
	}

	DatabaseToolsDatabaseToolsConnectionRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `tf_connection_name`, Update: `displayName2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle.jdbc.loginTimeout": "0"}, Update: map[string]string{"oracle.net.CONNECT_TIMEOUT": "0"}},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com`, Update: `connectionString2`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"proxy_client":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsUsernameProxyClientRepresentation},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_tools_database_tools_private_endpoint.test_private_endpoint.id}`},
		"related_resource":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `user@example.com`},
		"user_password":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreChangesDatabaseToolsConnectionRepresentation},
	}
	DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation = map[string]interface{}{
		"key_store_content":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresKeyStoreContentRepresentation},
		"key_store_password": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresKeyStorePasswordRepresentation},
		"key_store_type":     acctest.Representation{RepType: acctest.Optional, Create: `JAVA_KEY_STORE`, Update: `JAVA_TRUST_STORE`},
	}
	DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation = map[string]interface{}{
		"entity_type": acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
		"identifier":  acctest.Representation{RepType: acctest.Required, Create: `${var.related_resource_id}`, Update: `identifier2`},
	}
	DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
	}
	DatabaseToolsDatabaseToolsConnectionKeyStoresKeyStoreContentRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.secret_id}`},
	}
	DatabaseToolsDatabaseToolsConnectionKeyStoresKeyStorePasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.secret_id}`},
	}

	DatabaseToolsIgnoreChangesDatabaseToolsConnectionRepresentation = map[string]interface{}{ // This may vary depending on the tenancy settings
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DatabaseToolsDatabaseToolsConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsPrivateEndpointRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsConnectionOracleResource_basic(t *testing.T) {
	config := acctest.ProviderTestConfig()
	allVars := databaseToolsStandardVariables()
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	datasourceName := "data.oci_database_tools_database_tools_connections.test_database_tools_connections"
	singularDatasourceName := "data.oci_database_tools_database_tools_connection.test_database_tools_connection"
	relatedResourceId := utils.GetEnvSettingWithBlankDefault("related_resource_id")

	var resId string
	var resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		Providers:    databaseToolsOciProvider(),
		CheckDestroy: testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy,
		Steps: []resource.TestStep{
			// Step 1. Verify create
			{
				Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						print(resId)
						return err
					},
				),
			},

			// Step 2. Delete before next create
			{
				Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseToolsEndpointServiceDataSourceRepresentation),
			},
			//// Step 3. Verify create with optionals
			{
				Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
					resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", relatedResourceId),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

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
			//
			// Step 4. Verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + allVars + compartmentIdUVariableStr + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsConnectionRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
					resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", relatedResourceId),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// Step 5. Verify updates to updatable parameters
			{
				Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "connectionString2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
					resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", "identifier2"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Step 6. Verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connections", "test_database_tools_connections", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsDatabaseToolsConnectionDataSourceRepresentation) +
					allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.#", "1"),
				),
			},
			// Step 7. Verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseToolsConnectionSingularDataSourceRepresentation) +
					allVars + DatabaseToolsDatabaseToolsConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_string", "connectionString2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_password.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.entity_type", "DATABASE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.identifier", "identifier2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "type", "ORACLE_DATABASE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.0.value_type", "SECRETID"),
				),
			},
			// Step 8. Verify resource import
			{
				Config:                  config + DatabaseToolsDatabaseToolsConnectionRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func init() {
	databaseToolsInitDependencyGraphAndSweeper("DatabaseToolsDatabaseToolsConnection")
}
