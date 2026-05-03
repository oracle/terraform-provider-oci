// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialConnectionRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
			"advanced_properties",
			"defined_tags",
			"proxy_client",
			"freeform_tags",
			"key_stores",
			"private_endpoint_id",
			"related_resource",
			"runtime_identity",
			"runtime_support",
			"lifecycle",
		}, DatabaseToolsDatabaseToolsConnectionRepresentation),
		map[string]interface{}{
			"connection_string": acctest.Representation{RepType: acctest.Required, Create: `${var.connection_string}`},
			"key_stores":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymKeyStoresRepresentation},
			"runtime_identity":  acctest.Representation{RepType: acctest.Required, Create: `RESOURCE_PRINCIPAL`},
			"runtime_support":   acctest.Representation{RepType: acctest.Required, Create: `SUPPORTED`},
			"user_name":         acctest.Representation{RepType: acctest.Required, Create: `admin`},
		},
	)

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceConfig = DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential", "test_database_tools_connection_credential", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialRepresentation)

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_connection_credential.test_database_tools_connection_credential.key}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsConnectionCredentialDataSourceFilterRepresentation}}
	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_runtime_database_tools_connection_credential.test_database_tools_connection_credential.key}`}},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `MY_TEST_CREDENTIAL4`},
		"password":                     acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `BASIC`},
		"user_name":                    acctest.Representation{RepType: acctest.Required, Create: `admin`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialConnectionRepresentation)
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsConnectionCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsConnectionCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	allVars := databaseToolsRuntimeBaseVariables()

	resourceName := "oci_database_tools_runtime_database_tools_connection_credential.test_database_tools_connection_credential"
	datasourceName := "data.oci_database_tools_runtime_database_tools_connection_credentials.test_database_tools_connection_credentials"
	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_connection_credential.test_database_tools_connection_credential"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+allVars+DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential", "test_database_tools_connection_credential", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialRepresentation), "databasetoolsruntime", "databaseToolsConnectionCredential", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsRuntimeDatabaseToolsConnectionCredentialDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential", "test_database_tools_connection_credential", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "key", "MY_TEST_CREDENTIAL4"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),

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
			Config: config + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential", "test_database_tools_connection_credential", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "enabled"),
				resource.TestCheckResourceAttr(resourceName, "key", "MY_TEST_CREDENTIAL4"),
				resource.TestCheckResourceAttrSet(resourceName, "owner"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),

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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credentials", "test_database_tools_connection_credentials", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialDataSourceRepresentation) +
				allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential", "test_database_tools_connection_credential", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_id"),

				resource.TestCheckResourceAttr(datasourceName, "credential_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "credential_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential", "test_database_tools_connection_credential", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialSingularDataSourceRepresentation) +
				allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_key", "MY_TEST_CREDENTIAL4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "enabled"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key", "MY_TEST_CREDENTIAL4"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner"),
				resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_name", "admin"),
			),
		},
		// verify resource import
		{
			Config:            config + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateCheck: checkImportedDatabaseToolsRuntimeCompositeID(
				resourceName,
				parseDatabaseToolsRuntimeConnectionCredentialCompositeIDToAttributes,
			),
			ImportStateVerifyIgnore: []string{
				"password",
				"type",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseToolsRuntimeDatabaseToolsConnectionCredentialDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsRuntimeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_runtime_database_tools_connection_credential" {
			noResourceFound = false
			request := oci_database_tools_runtime.GetCredentialRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.CredentialKey = &value
			}

			if value, ok := rs.Primary.Attributes["database_tools_connection_id"]; ok {
				request.DatabaseToolsConnectionId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")

			_, err := client.GetCredential(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			if request.CredentialKey == nil || request.DatabaseToolsConnectionId == nil {
				credentialKey, databaseToolsConnectionId, err := parseDatabaseToolsRuntimeConnectionCredentialCompositeId(rs.Primary.ID)
				if err != nil {
					return err
				}
				request.CredentialKey = &credentialKey
				request.DatabaseToolsConnectionId = &databaseToolsConnectionId
			}

			if failure, isServiceError := common.IsServiceError(err); isServiceError {
				if failure.GetHTTPStatusCode() == 404 {
					continue
				}
				if failure.GetHTTPStatusCode() == 400 && strings.Contains(failure.GetMessage(), "Connection is inactive") {
					continue
				}
			}
			return err
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
	if !acctest.InSweeperExcludeList("DatabaseToolsRuntimeDatabaseToolsConnectionCredential") {
		resource.AddTestSweepers("DatabaseToolsRuntimeDatabaseToolsConnectionCredential", &resource.Sweeper{
			Name:         "DatabaseToolsRuntimeDatabaseToolsConnectionCredential",
			Dependencies: acctest.DependencyGraph["databaseToolsConnectionCredential"],
			F:            sweepDatabaseToolsRuntimeDatabaseToolsConnectionCredentialResource,
		})
	}
}

func sweepDatabaseToolsRuntimeDatabaseToolsConnectionCredentialResource(compartment string) error {
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()
	databaseToolsConnectionCredentialIds, err := getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsConnectionCredentialId := range databaseToolsConnectionCredentialIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsConnectionCredentialId]; !ok {
			deleteCredentialRequest := oci_database_tools_runtime.DeleteCredentialRequest{}

			deleteCredentialRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")
			_, error := databaseToolsRuntimeClient.DeleteCredential(context.Background(), deleteCredentialRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsConnectionCredential %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsConnectionCredentialId, error)
				continue
			}
		}
	}
	return nil
}

func getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsConnectionCredentialId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()

	listCredentialsRequest := oci_database_tools_runtime.ListCredentialsRequest{}

	databaseToolsConnectionIds, error := getDatabaseToolsConnectionIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting databaseToolsConnectionId required for DatabaseToolsConnectionCredential resource requests \n")
	}
	for _, databaseToolsConnectionId := range databaseToolsConnectionIds {
		listCredentialsRequest.DatabaseToolsConnectionId = &databaseToolsConnectionId

		listCredentialsResponse, err := databaseToolsRuntimeClient.ListCredentials(context.Background(), listCredentialsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DatabaseToolsConnectionCredential list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, databaseToolsConnectionCredential := range listCredentialsResponse.Items {
			if databaseToolsConnectionCredential.Key == nil {
				continue
			}
			id := databaseToolsConnectionId + "/" + *databaseToolsConnectionCredential.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsConnectionCredentialId", id)
		}

	}
	return resourceIds, nil
}

func parseDatabaseToolsRuntimeConnectionCredentialCompositeId(compositeId string) (credentialKey string, databaseToolsConnectionId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsConnections/.*/credentials/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	databaseToolsConnectionId, _ = url.PathUnescape(parts[1])
	credentialKey, _ = url.PathUnescape(parts[3])
	return
}
