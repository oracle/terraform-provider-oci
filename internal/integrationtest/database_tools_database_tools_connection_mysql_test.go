// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsConnectionMySqlRequiredOnlyResource = DatabaseToolsConnectionMySqlResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_mysql_connection", acctest.Required, acctest.Create, databaseToolsConnectionMySqlRepresentation)

	DatabaseToolsConnectionMySqlResourceConfig = DatabaseToolsConnectionMySqlResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_mysql_connection", acctest.Optional, acctest.Update, databaseToolsConnectionMySqlRepresentation)

	databaseToolsConnectionMySqlSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_mysql_connection.id}`},
	}

	databaseToolsConnectionMySqlDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `tf_mysql_connection_name`, Update: `displayNameMySql2`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`MYSQL`}},
		"runtime_support": acctest.Representation{RepType: acctest.Optional, Create: []string{`SUPPORTED`}},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsConnectionMySqlDataSourceFilterRepresentation}}

	databaseToolsConnectionMySqlDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_connection.test_database_tools_mysql_connection.id}`}},
	}

	databaseToolsConnectionMySqlRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `tf_mysql_connection_name`, Update: `displayNameMySql2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `MYSQL`},
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sslMode": "REQUIRED"}, Update: map[string]string{"sslMode": "VERIFY_CA"}},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `mysql://example.com:3306`, Update: `mysql://example2.com:3306`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionMySqlKeyStoresRepresentation},
		"related_resource":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionMySqlRelatedResourceRepresentation},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `user@example.com`, Update: `user2@example2.com`},
		"user_password":       acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsConnectionMySqlUserPasswordRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesDatabaseToolsConnectionMySqlRepresentation},
	}
	databaseToolsConnectionMySqlKeyStoresRepresentation = map[string]interface{}{
		"key_store_content": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionMySqlKeyStoresKeyStoreContentRepresentation},
		"key_store_type":    acctest.Representation{RepType: acctest.Optional, Create: `CA_CERTIFICATE_PEM`},
	}
	databaseToolsConnectionMySqlRelatedResourceRepresentation = map[string]interface{}{
		"entity_type": acctest.Representation{RepType: acctest.Required, Create: `MYSQLDBSYSTEM`},
		"identifier":  acctest.Representation{RepType: acctest.Required, Create: `${var.related_resource_id}`, Update: `identifier2`},
	}
	databaseToolsConnectionMySqlUserPasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
	}
	databaseToolsConnectionMySqlKeyStoresKeyStoreContentRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.secret_id}`},
	}

	ignoreChangesDatabaseToolsConnectionMySqlRepresentation = map[string]interface{}{ // This may vary depending on the tenancy settings
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DatabaseToolsConnectionMySqlResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsConnectionMySqlResource_basic(t *testing.T) {
	config := acctest.ProviderTestConfig()
	allVars := databaseToolsStandardVariables()
	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_mysql_connection"
	datasourceName := "data.oci_database_tools_database_tools_connections.test_database_tools_mysql_connections"
	singularDatasourceName := "data.oci_database_tools_database_tools_connection.test_database_tools_mysql_connection"
	secretId := utils.GetEnvSettingWithBlankDefault("secret_id")
	relatedResourceId := utils.GetEnvSettingWithBlankDefault("related_resource_id")

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		Providers:    databaseToolsOciProvider(),
		CheckDestroy: testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy,
		Steps: []resource.TestStep{
			// Step 1. Verify create
			{
				Config: config + allVars + DatabaseToolsConnectionMySqlResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_mysql_connection", acctest.Required, acctest.Create, databaseToolsConnectionMySqlRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_mysql_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "type", "MYSQL"),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "mysql://example.com:3306"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// Step 2. Delete before next create
			{
				Config: config + allVars + DatabaseToolsConnectionMySqlResourceDependencies,
			},
			//Step 3. Verify create with optionals
			{
				Config: config + allVars + DatabaseToolsConnectionMySqlResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_mysql_connection", acctest.Optional, acctest.Create, databaseToolsConnectionMySqlRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "mysql://example.com:3306"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_mysql_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "CA_CERTIFICATE_PEM"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "MYSQLDBSYSTEM"),
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", relatedResourceId),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "MYSQL"),
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

			// Step 4. Verify datasource after the update is done
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connections", "test_database_tools_mysql_connections", acctest.Optional, acctest.Update, databaseToolsConnectionMySqlDataSourceRepresentation) +
					allVars + DatabaseToolsConnectionMySqlResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_mysql_connection", acctest.Optional, acctest.Update, databaseToolsConnectionMySqlRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.key_stores.0.key_store_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.time_updated"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.advanced_properties.sslMode", "VERIFY_CA"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.type", "MYSQL"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.lifecycle_details", "Updated"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.connection_string", "mysql://example2.com:3306"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.runtime_support", "SUPPORTED"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_name", "user2@example2.com"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.display_name", "displayNameMySql2"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.0.secret_id", secretId),
				),
			},
			// Step 5. Verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_mysql_connection", acctest.Required, acctest.Create, databaseToolsConnectionMySqlSingularDataSourceRepresentation) +
					allVars + DatabaseToolsConnectionMySqlResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_string", "mysql://example2.com:3306"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayNameMySql2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_type", "CA_CERTIFICATE_PEM"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.entity_type", "MYSQLDBSYSTEM"),
					resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.identifier", "identifier2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "type", "MYSQL"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.0.value_type", "SECRETID"),
				),
			},

			// Step 6. Verify resource import
			{
				Config:                  config + DatabaseToolsConnectionMySqlRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func init() {
	databaseToolsInitDependencyGraphAndSweeper("DatabaseToolsDatabaseToolsConnectionMySql")
}
