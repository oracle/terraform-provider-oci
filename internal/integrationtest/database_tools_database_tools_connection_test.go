// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v56/databasetools"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DatabaseToolsConnectionRequiredOnlyResource = DatabaseToolsConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, databaseToolsConnectionRepresentation)

	DatabaseToolsConnectionResourceConfig = DatabaseToolsConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, databaseToolsConnectionRepresentation)

	databaseToolsConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
	}

	databaseToolsConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tf_connection_name`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_DATABASE`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsConnectionDataSourceFilterRepresentation}}
	databaseToolsConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`}},
	}

	databaseToolsConnectionRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `tf_connection_name`, Update: `displayName2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle.jdbc.loginTimeout": "0"}, Update: map[string]string{"oracle.net.CONNECT_TIMEOUT": "0"}},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `tcps://adb.us-phoenix-1.oraclecloud.com:1522/u9adutfb2ba8x4d_db202103231111_low.adb.oraclecloud.com`, Update: `connectionString2`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionKeyStoresRepresentation},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_tools_database_tools_private_endpoint.test_private_endpoint.id}`},
		"related_resource":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionRelatedResourceRepresentation},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.name}`},
		"user_password":       acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsConnectionUserPasswordRepresentation},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesDatabaseToolsConnectionRepresentation},
	}
	databaseToolsConnectionKeyStoresRepresentation = map[string]interface{}{
		"key_store_content":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionKeyStoresKeyStoreContentRepresentation},
		"key_store_password": acctest.RepresentationGroup{RepType: acctest.Optional, Group: databaseToolsConnectionKeyStoresKeyStorePasswordRepresentation},
		"key_store_type":     acctest.Representation{RepType: acctest.Optional, Create: `JAVA_KEY_STORE`, Update: `JAVA_TRUST_STORE`},
	}
	databaseToolsConnectionRelatedResourceRepresentation = map[string]interface{}{
		"entity_type": acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
		"identifier":  acctest.Representation{RepType: acctest.Required, Create: `ocid1.database.oc1.phx.exampletksujfufl4bhe5sqkfgn7t7lcrkkpy7km5iwzvg6ycls7r5dlbx6q`, Update: `identifier2`},
	}
	databaseToolsConnectionUserPasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `ocid1.vaultsecret.dev.dev.amaaaaaaihuofciaie44ubvpggl6zrodrar7ils25hf53qyue3w5t3awtufa`},
	}
	databaseToolsConnectionKeyStoresKeyStoreContentRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Optional, Create: `ocid1.vaultsecret.dev.dev.amaaaaaaihuofciaie44ubvpggl6zrodrar7ils25hf53qyue3w5t3awtufa`},
	}
	databaseToolsConnectionKeyStoresKeyStorePasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Optional, Create: `ocid1.vaultsecret.dev.dev.amaaaaaaihuofciaie44ubvpggl6zrodrar7ils25hf53qyue3w5t3awtufa`},
	}

	ignoreChangesDatabaseToolsConnectionRepresentation = map[string]interface{}{ // This may vary depending on the tenancy settings
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DatabaseToolsConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, databaseToolsPrivateEndpointRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, userRepresentation)
)

func TestDatabaseToolsDatabaseToolsConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsConnectionResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	datasourceName := "data.oci_database_tools_database_tools_connections.test_database_tools_connections"
	singularDatasourceName := "data.oci_database_tools_database_tools_connection.test_database_tools_connection"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseToolsConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, databaseToolsConnectionRepresentation), "databasetools", "databaseToolsConnection", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy,
		Steps: []resource.TestStep{
			// 0. verify create
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, databaseToolsConnectionRepresentation),
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
						return err
					},
				),
			},

			// 1. delete before next create
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation),
			},
			// 2. verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, databaseToolsConnectionRepresentation),
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
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", "ocid1.database.oc1.phx.exampletksujfufl4bhe5sqkfgn7t7lcrkkpy7km5iwzvg6ycls7r5dlbx6q"),
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

			// 3. verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(databaseToolsConnectionRepresentation, map[string]interface{}{
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
					resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", "ocid1.database.oc1.phx.exampletksujfufl4bhe5sqkfgn7t7lcrkkpy7km5iwzvg6ycls7r5dlbx6q"),
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

			// 4. verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, databaseToolsConnectionRepresentation),
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
			// 5. verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connections", "test_database_tools_connections", acctest.Optional, acctest.Update, databaseToolsConnectionDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseToolsConnectionResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, databaseToolsConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.#", "1"),
				),
			},
			// 6. verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, databaseToolsConnectionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseToolsConnectionResourceConfig,
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
			// 7. remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsConnectionResourceConfig +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation),
			},
			// 8. verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_connection" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsConnectionRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.DatabaseToolsConnection.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.DatabaseToolsConnection.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseToolsDatabaseToolsConnection") {
		resource.AddTestSweepers("DatabaseToolsDatabaseToolsConnection", &resource.Sweeper{
			Name:         "DatabaseToolsDatabaseToolsConnection",
			Dependencies: acctest.DependencyGraph["databaseToolsConnection"],
			F:            sweepDatabaseToolsDatabaseToolsConnectionResource,
		})
	}
}

func sweepDatabaseToolsDatabaseToolsConnectionResource(compartment string) error {
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	databaseToolsConnectionIds, err := getDatabaseToolsConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsConnectionId := range databaseToolsConnectionIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsConnectionId]; !ok {
			deleteDatabaseToolsConnectionRequest := oci_database_tools.DeleteDatabaseToolsConnectionRequest{}

			deleteDatabaseToolsConnectionRequest.DatabaseToolsConnectionId = &databaseToolsConnectionId

			deleteDatabaseToolsConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsConnection(context.Background(), deleteDatabaseToolsConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsConnectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseToolsConnectionId, databaseToolsConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				databaseToolsConnectionSweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()

	listDatabaseToolsConnectionsRequest := oci_database_tools.ListDatabaseToolsConnectionsRequest{}
	listDatabaseToolsConnectionsRequest.CompartmentId = &compartmentId
	listDatabaseToolsConnectionsRequest.LifecycleState = oci_database_tools.ListDatabaseToolsConnectionsLifecycleStateActive
	listDatabaseToolsConnectionsResponse, err := databaseToolsClient.ListDatabaseToolsConnections(context.Background(), listDatabaseToolsConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsConnection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsConnection := range listDatabaseToolsConnectionsResponse.Items {
		id := *databaseToolsConnection.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsConnectionId", id)
	}
	return resourceIds, nil
}

func databaseToolsConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseToolsConnectionResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsConnectionResponse); ok {
		return databaseToolsConnectionResponse.DatabaseToolsConnection.GetLifecycleState() != oci_database_tools.LifecycleStateDeleted
	}
	return false
}

func databaseToolsConnectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseToolsClient().GetDatabaseToolsConnection(context.Background(), oci_database_tools.GetDatabaseToolsConnectionRequest{
		DatabaseToolsConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
