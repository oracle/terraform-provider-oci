// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	DatabaseToolsConnectionGenericJdbcRequiredOnlyResource = DatabaseToolsConnectionGenericJdbcResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_generic_jdbc_connection", acctest.Required, acctest.Create, databaseToolsConnectionGenericJdbcRepresentation)

	DatabaseToolsConnectionGenericJdbcResourceConfig = DatabaseToolsConnectionGenericJdbcResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_database_tools_database_tools_connection",
			"test_database_tools_generic_jdbc_connection",
			acctest.Optional,
			acctest.Update,
			databaseToolsConnectionGenericJdbcRepresentation)

	databaseToolsConnectionGenericJdbcSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_generic_jdbc_connection.id}`},
	}

	databaseToolsConnectionGenericJdbcDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `tf_generic_jdbc_connection_name`, Update: `displayNameGenericJdbc2`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`GENERIC_JDBC`}},
		"runtime_support": acctest.Representation{RepType: acctest.Optional, Create: []string{`UNSUPPORTED`}},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsConnectionGenericJdbcDataSourceFilterRepresentation},
	}

	databaseToolsConnectionGenericJdbcDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_connection.test_database_tools_generic_jdbc_connection.id}`}},
	}

	databaseToolsConnectionGenericJdbcRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `tf_generic_jdbc_connection_name`, Update: `displayNameGenericJdbc2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `GENERIC_JDBC`},
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sslMode": "REQUIRED"}, Update: map[string]string{"sslMode": "VERIFY_CA"}},
		"url":                 acctest.Representation{RepType: acctest.Required, Create: `jdbc:mysql://example.com:3306`, Update: `jdbc:mysql://example2.com:3306`},
		"runtime_support":     acctest.Representation{RepType: acctest.Required, Create: `UNSUPPORTED`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"foo": "value"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionGenericJdbcKeyStoresRepresentation},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `user@example.com`, Update: `user2@example2.com`},
		"user_password":       acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsConnectionGenericJdbcUserPasswordRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesDatabaseToolsConnectionGenericJdbcRepresentation},
	}
	databaseToolsConnectionGenericJdbcKeyStoresRepresentation = map[string]interface{}{
		"key_store_content": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionGenericJdbcKeyStoresKeyStoreContentRepresentation},
		"key_store_type":    acctest.Representation{RepType: acctest.Optional, Create: `CA_CERTIFICATE_PEM`},
	}
	databaseToolsConnectionGenericJdbcUserPasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
	}
	databaseToolsConnectionGenericJdbcKeyStoresKeyStoreContentRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.secret_id}`},
	}

	ignoreChangesDatabaseToolsConnectionGenericJdbcRepresentation = map[string]interface{}{ // This may vary depending on the tenancy settings
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DatabaseToolsConnectionGenericJdbcResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsConnectionGenericJdbcResource_basic(t *testing.T) {
	config := acctest.ProviderTestConfig()
	allVars := databaseToolsStandardVariables()
	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_generic_jdbc_connection"
	datasourceName := "data.oci_database_tools_database_tools_connections.test_database_tools_generic_jdbc_connections"
	singularDatasourceName := "data.oci_database_tools_database_tools_connection.test_database_tools_generic_jdbc_connection"
	secretId := utils.GetEnvSettingWithBlankDefault("secret_id")

	var resourceId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		Providers:    databaseToolsOciProvider(),
		CheckDestroy: testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy,
		Steps: []resource.TestStep{
			//Step 1. Verify create
			{
				Config: config + allVars + DatabaseToolsConnectionGenericJdbcResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap(
						"oci_database_tools_database_tools_connection",
						"test_database_tools_generic_jdbc_connection",
						acctest.Required,
						acctest.Create,
						databaseToolsConnectionGenericJdbcRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_generic_jdbc_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "type", "GENERIC_JDBC"),
					resource.TestCheckResourceAttr(resourceName, "runtime_support", "UNSUPPORTED"),
					resource.TestCheckResourceAttr(resourceName, "url", "jdbc:mysql://example.com:3306"),
					resource.TestCheckResourceAttr(resourceName, "user_name", "user@example.com"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.secret_id", secretId),

					func(s *terraform.State) (err error) {
						resourceId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// Step 2. Delete the resource
			{
				Config: config + allVars + DatabaseToolsConnectionGenericJdbcResourceDependencies,
			},

			//Step 3. Verify create with optionals
			{
				Config: config + allVars + DatabaseToolsConnectionGenericJdbcResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap(
						"oci_database_tools_database_tools_connection",
						"test_database_tools_generic_jdbc_connection",
						acctest.Optional,
						acctest.Create,
						databaseToolsConnectionGenericJdbcRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "url", "jdbc:mysql://example.com:3306"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_generic_jdbc_connection_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "CA_CERTIFICATE_PEM"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "GENERIC_JDBC"),
					resource.TestCheckResourceAttr(resourceName, "runtime_support", "UNSUPPORTED"),
					resource.TestCheckResourceAttr(resourceName, "user_name", "user@example.com"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.secret_id", secretId),

					func(s *terraform.State) (err error) {
						resourceId, err = acctest.FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resourceId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			//Step 4. Verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap(
						"oci_database_tools_database_tools_connections",
						"test_database_tools_generic_jdbc_connections",
						acctest.Optional,
						acctest.Update,
						databaseToolsConnectionGenericJdbcDataSourceRepresentation) +
					allVars + DatabaseToolsConnectionGenericJdbcResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap(
						"oci_database_tools_database_tools_connection",
						"test_database_tools_generic_jdbc_connection",
						acctest.Optional,
						acctest.Update,
						databaseToolsConnectionGenericJdbcRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.key_stores.0.key_store_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.time_updated"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.advanced_properties.sslMode", "VERIFY_CA"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.type", "GENERIC_JDBC"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.lifecycle_details", "Updated"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.url", "jdbc:mysql://example2.com:3306"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.runtime_support", "UNSUPPORTED"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_name", "user2@example2.com"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.display_name", "displayNameGenericJdbc2"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.0.secret_id", secretId),
				),
			},

			//Step 5. Verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap(
						"oci_database_tools_database_tools_connection",
						"test_database_tools_generic_jdbc_connection",
						acctest.Required,
						acctest.Create,
						databaseToolsConnectionGenericJdbcSingularDataSourceRepresentation) +
					allVars + DatabaseToolsConnectionGenericJdbcResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "url", "jdbc:mysql://example2.com:3306"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayNameGenericJdbc2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_type", "CA_CERTIFICATE_PEM"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "type", "GENERIC_JDBC"),
					resource.TestCheckResourceAttr(singularDatasourceName, "runtime_support", "UNSUPPORTED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.0.value_type", "SECRETID"),
				),
			},

			// Step 6. Verify resource import
			{
				Config:                  config + DatabaseToolsConnectionGenericJdbcRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func init() {
	databaseToolsInitDependencyGraphAndSweeper("DatabaseToolsDatabaseToolsConnectionGenericJdbc")
}
