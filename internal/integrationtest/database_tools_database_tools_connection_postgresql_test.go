// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var (
	DatabaseToolsConnectionPostgresqlRequiredOnlyResource = DatabaseToolsConnectionPostgresqlResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_postgresql_connection", acctest.Required, acctest.Create, databaseToolsConnectionPostgresqlRepresentation)

	DatabaseToolsConnectionPostgresqlResourceConfig = acctest.GenerateResourceFromRepresentationMap(
		"oci_database_tools_database_tools_connection",
		"test_database_tools_postgresql_connection",
		acctest.Optional,
		acctest.Update,
		databaseToolsConnectionPostgresqlRepresentation)

	databaseToolsConnectionPostgresqlSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_postgresql_connection.id}`},
	}

	databaseToolsConnectionPostgresqlDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `tf_postgresql_connection_name`, Update: `displayNamePostgresql2`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"runtime_support": acctest.Representation{RepType: acctest.Optional, Create: []string{`UNSUPPORTED`}},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`POSTGRESQL`}},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsConnectionPostgresqlDataSourceFilterRepresentation},
	}

	databaseToolsConnectionPostgresqlDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_connection.test_database_tools_postgresql_connection.id}`}},
	}

	databaseToolsConnectionPostgresqlRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `tf_postgresql_connection_name`, Update: `displayNamePostgresql2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `POSTGRESQL`},
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"sslMode": "REQUIRED"}, Update: map[string]string{"sslMode": "VERIFY_CA"}},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `postgresql://example.com:3306`, Update: `postgresql://example2.com:3306`},
		"runtime_support":     acctest.Representation{RepType: acctest.Required, Create: `UNSUPPORTED`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"foo": "value"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionPostgresqlKeyStoresRepresentation},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `user@example.com`, Update: `user2@example2.com`},
		"user_password":       acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsConnectionPostgresqlUserPasswordRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesDatabaseToolsConnectionPostgresqlRepresentation},
	}
	databaseToolsConnectionPostgresqlKeyStoresRepresentation = map[string]interface{}{
		"key_store_content": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionPostgresqlKeyStoresKeyStoreContentRepresentation},
		"key_store_type":    acctest.Representation{RepType: acctest.Optional, Create: `CA_CERTIFICATE_PEM`},
	}
	databaseToolsConnectionPostgresqlUserPasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
	}
	databaseToolsConnectionPostgresqlKeyStoresKeyStoreContentRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.secret_id}`},
	}

	ignoreChangesDatabaseToolsConnectionPostgresqlRepresentation = map[string]interface{}{ // This may vary depending on the tenancy settings
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DatabaseToolsConnectionPostgresqlResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsConnectionPostgresqlResource_basic(t *testing.T) {
	config := acctest.ProviderTestConfig()
	allVars := databaseToolsStandardVariables()
	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_postgresql_connection"
	datasourceName := "data.oci_database_tools_database_tools_connections.test_database_tools_postgresql_connections"
	singularDatasourceName := "data.oci_database_tools_database_tools_connection.test_database_tools_postgresql_connection"
	secretId := utils.GetEnvSettingWithBlankDefault("secret_id")

	var resourceId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		Providers:    databaseToolsOciProvider(),
		CheckDestroy: testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy,
		Steps: []resource.TestStep{
			// Step 1. Verify create
			{
				Config: config + allVars + DatabaseToolsConnectionPostgresqlResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap(
						"oci_database_tools_database_tools_connection",
						"test_database_tools_postgresql_connection",
						acctest.Required,
						acctest.Create,
						databaseToolsConnectionPostgresqlRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_postgresql_connection_name"),
					resource.TestCheckResourceAttr(resourceName, "type", "POSTGRESQL"),
					resource.TestCheckResourceAttr(resourceName, "runtime_support", "UNSUPPORTED"),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "postgresql://example.com:3306"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

					func(s *terraform.State) (err error) {
						resourceId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// Step 2. Delete the resource
			{
				Config: config + allVars + DatabaseToolsConnectionPostgresqlResourceDependencies,
			},

			// Step 3. Verify create with optionals
			{
				Config: config + allVars + DatabaseToolsConnectionPostgresqlResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap(
						"oci_database_tools_database_tools_connection",
						"test_database_tools_postgresql_connection",
						acctest.Optional,
						acctest.Create,
						databaseToolsConnectionPostgresqlRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "connection_string", "postgresql://example.com:3306"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tf_postgresql_connection_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "CA_CERTIFICATE_PEM"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
					resource.TestCheckResourceAttr(resourceName, "type", "POSTGRESQL"),
					resource.TestCheckResourceAttr(resourceName, "runtime_support", "UNSUPPORTED"),
					resource.TestCheckResourceAttrSet(resourceName, "user_name"),
					resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
					resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),

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

			//Step 4. Verify datasource after the update is done
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap(
						"oci_database_tools_database_tools_connections",
						"test_database_tools_postgresql_connections",
						acctest.Optional,
						acctest.Update,
						databaseToolsConnectionPostgresqlDataSourceRepresentation) +
					allVars + DatabaseToolsConnectionPostgresqlResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap(
						"oci_database_tools_database_tools_connection",
						"test_database_tools_postgresql_connection",
						acctest.Optional,
						acctest.Update,
						databaseToolsConnectionPostgresqlRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.key_stores.0.key_store_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_collection.0.items.0.time_updated"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.advanced_properties.sslMode", "VERIFY_CA"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.type", "POSTGRESQL"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.lifecycle_details", "Updated"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.connection_string", "postgresql://example2.com:3306"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.runtime_support", "UNSUPPORTED"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_name", "user2@example2.com"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.display_name", "displayNamePostgresql2"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.0.user_password.0.secret_id", secretId),
				),
			},
			// Step 5. Verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap(
						"oci_database_tools_database_tools_connection",
						"test_database_tools_postgresql_connection",
						acctest.Required,
						acctest.Create,
						databaseToolsConnectionPostgresqlSingularDataSourceRepresentation) + DatabaseToolsConnectionPostgresqlResourceDependencies +
					allVars + DatabaseToolsConnectionPostgresqlResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "advanced_properties.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "connection_string", "postgresql://example2.com:3306"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayNamePostgresql2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
					resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_type", "CA_CERTIFICATE_PEM"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "type", "POSTGRESQL"),
					resource.TestCheckResourceAttr(singularDatasourceName, "runtime_support", "UNSUPPORTED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "user_password.0.value_type", "SECRETID"),
				),
			},

			// Step 6. Verify resource import
			{
				Config:                  config + DatabaseToolsConnectionPostgresqlRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})

	println(resourceId)
}

func init() {
	databaseToolsInitDependencyGraphAndSweeper("DatabaseToolsDatabaseToolsConnectionPostgresql")
}
