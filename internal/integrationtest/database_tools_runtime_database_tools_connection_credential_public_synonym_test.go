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
	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymKeyStoresRepresentation = map[string]interface{}{
		"key_store_content": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymKeyStoreContentRepresentation},
		"key_store_type":    acctest.Representation{RepType: acctest.Required, Create: `SSO`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymKeyStoreContentRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.wallet_secret_id}`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceConfig = DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_public_synonym", "test_database_tools_connection_credential_public_synonym", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymRepresentation)

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymSingularDataSourceRepresentation = map[string]interface{}{
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_credential_key}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"public_synonym_key":           acctest.Representation{RepType: acctest.Required, Create: `${var.public_synonym_key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceRepresentation = map[string]interface{}{
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_credential_key}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceFilterRepresentation}}
	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.public_synonym_key}`}},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymRepresentation = map[string]interface{}{
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_credential_key}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `${var.public_synonym_key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceDependencies = ""
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	existingDatabaseToolsConnectionId := utils.GetEnvSettingWithDefault("existing_database_tools_connection_id",
		utils.GetEnvSettingWithDefault("database_tools_connection_id",
			utils.GetEnvSettingWithBlankDefault("database_tools_connection_ocid")))
	if existingDatabaseToolsConnectionId == "" {
		t.Skip("set existing_database_tools_connection_id, database_tools_connection_id, or database_tools_connection_ocid to run this test")
	}

	existingDatabaseToolsConnectionCredentialKey := utils.GetEnvSettingWithDefault("existing_database_tools_connection_credential_key",
		utils.GetEnvSettingWithDefault("database_tools_connection_credential_key",
			utils.GetEnvSettingWithDefault("database_tools_connection_credential_public_synonym_credential_key",
				utils.GetEnvSettingWithDefault("user_credential_key", "credentialKey"))))
	publicSynonymKey := utils.GetEnvSettingWithDefault("existing_database_tools_connection_credential_public_synonym_key",
		utils.GetEnvSettingWithDefault("database_tools_connection_credential_public_synonym_key",
			utils.GetEnvSettingWithDefault("public_synonym_key", "MY_PUBLIC_SYNONYM")))
	allVars := databaseToolsRuntimeExistingPublicSynonymVariables()

	resourceName := "oci_database_tools_runtime_database_tools_connection_credential_public_synonym.test_database_tools_connection_credential_public_synonym"
	datasourceName := "data.oci_database_tools_runtime_database_tools_connection_credential_public_synonyms.test_database_tools_connection_credential_public_synonyms"
	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_connection_credential_public_synonym.test_database_tools_connection_credential_public_synonym"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+allVars+DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_public_synonym", "test_database_tools_connection_credential_public_synonym", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymRepresentation), "databasetoolsruntime", "databaseToolsConnectionCredentialPublicSynonym", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_public_synonym", "test_database_tools_connection_credential_public_synonym", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credential_key", existingDatabaseToolsConnectionCredentialKey),
				resource.TestCheckResourceAttr(resourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(resourceName, "key", publicSynonymKey),

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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_public_synonyms", "test_database_tools_connection_credential_public_synonyms", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSourceRepresentation) +
				allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_public_synonym", "test_database_tools_connection_credential_public_synonym", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "credential_key", existingDatabaseToolsConnectionCredentialKey),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),

				resource.TestCheckResourceAttr(datasourceName, "credential_public_synonym_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "credential_public_synonym_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_public_synonym", "test_database_tools_connection_credential_public_synonym", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymSingularDataSourceRepresentation) +
				allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_key", existingDatabaseToolsConnectionCredentialKey),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(singularDatasourceName, "public_synonym_key", publicSynonymKey),

				resource.TestCheckResourceAttr(singularDatasourceName, "key", publicSynonymKey),
			),
		},
		// verify resource import
		{
			Config:            config + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateCheck: checkImportedDatabaseToolsRuntimeCompositeID(
				resourceName,
				parseDatabaseToolsRuntimeConnectionCredentialPublicSynonymCompositeIDToAttributes,
			),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func databaseToolsRuntimeExistingPublicSynonymVariables() string {
	return terraformStringVariable("existing_database_tools_connection_id",
		utils.GetEnvSettingWithDefault("existing_database_tools_connection_id",
			utils.GetEnvSettingWithDefault("database_tools_connection_id",
				utils.GetEnvSettingWithBlankDefault("database_tools_connection_ocid")))) +
		terraformStringVariable("existing_database_tools_connection_credential_key",
			utils.GetEnvSettingWithDefault("existing_database_tools_connection_credential_key",
				utils.GetEnvSettingWithDefault("database_tools_connection_credential_key",
					utils.GetEnvSettingWithDefault("database_tools_connection_credential_public_synonym_credential_key",
						utils.GetEnvSettingWithDefault("user_credential_key", "credentialKey"))))) +
		terraformStringVariable("public_synonym_key",
			utils.GetEnvSettingWithDefault("existing_database_tools_connection_credential_public_synonym_key",
				utils.GetEnvSettingWithDefault("database_tools_connection_credential_public_synonym_key",
					utils.GetEnvSettingWithDefault("public_synonym_key", "MY_PUBLIC_SYNONYM"))))
}

func testAccCheckDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsRuntimeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_runtime_database_tools_connection_credential_public_synonym" {
			noResourceFound = false
			request := oci_database_tools_runtime.GetCredentialPublicSynonymRequest{}

			if value, ok := rs.Primary.Attributes["credential_key"]; ok {
				request.CredentialKey = &value
			}

			if value, ok := rs.Primary.Attributes["database_tools_connection_id"]; ok {
				request.DatabaseToolsConnectionId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.PublicSynonymKey = &value
			}

			if request.CredentialKey == nil || request.DatabaseToolsConnectionId == nil || request.PublicSynonymKey == nil {
				credentialKey, databaseToolsConnectionId, publicSynonymKey, err := parseDatabaseToolsRuntimeConnectionCredentialPublicSynonymCompositeId(rs.Primary.ID)
				if err != nil {
					return err
				}
				request.CredentialKey = &credentialKey
				request.DatabaseToolsConnectionId = &databaseToolsConnectionId
				request.PublicSynonymKey = &publicSynonymKey
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")

			_, err := client.GetCredentialPublicSynonym(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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

func parseDatabaseToolsRuntimeConnectionCredentialPublicSynonymCompositeId(compositeId string) (credentialKey string, databaseToolsConnectionId string, publicSynonymKey string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("databaseToolsConnections/.*/credentials/.*/publicSynonyms/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	databaseToolsConnectionId, _ = url.PathUnescape(parts[1])
	credentialKey, _ = url.PathUnescape(parts[3])
	publicSynonymKey, _ = url.PathUnescape(parts[5])
	return
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonym") {
		resource.AddTestSweepers("DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonym", &resource.Sweeper{
			Name:         "DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonym",
			Dependencies: acctest.DependencyGraph["databaseToolsConnectionCredentialPublicSynonym"],
			F:            sweepDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResource,
		})
	}
}

func sweepDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResource(compartment string) error {
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()
	databaseToolsConnectionCredentialPublicSynonymIds, err := getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsConnectionCredentialPublicSynonymId := range databaseToolsConnectionCredentialPublicSynonymIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsConnectionCredentialPublicSynonymId]; !ok {
			deleteCredentialPublicSynonymRequest := oci_database_tools_runtime.DeleteCredentialPublicSynonymRequest{}

			deleteCredentialPublicSynonymRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")
			_, error := databaseToolsRuntimeClient.DeleteCredentialPublicSynonym(context.Background(), deleteCredentialPublicSynonymRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsConnectionCredentialPublicSynonym %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsConnectionCredentialPublicSynonymId, error)
				continue
			}
		}
	}
	return nil
}

func getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsConnectionCredentialPublicSynonymId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()

	listCredentialPublicSynonymsRequest := oci_database_tools_runtime.ListCredentialPublicSynonymsRequest{}

	databaseToolsConnectionCredentialIds, error := getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting credential ids required for DatabaseToolsConnectionCredentialPublicSynonym resource requests \n")
	}
	for _, databaseToolsConnectionCredentialId := range databaseToolsConnectionCredentialIds {
		parts := strings.SplitN(databaseToolsConnectionCredentialId, "/", 2)
		if len(parts) != 2 {
			continue
		}
		credentialKey := parts[1]
		listCredentialPublicSynonymsRequest.CredentialKey = &credentialKey

		databaseToolsConnectionIds, error := getDatabaseToolsConnectionIds(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting databaseToolsConnectionId required for DatabaseToolsConnectionCredentialPublicSynonym resource requests \n")
		}
		for _, databaseToolsConnectionId := range databaseToolsConnectionIds {
			listCredentialPublicSynonymsRequest.DatabaseToolsConnectionId = &databaseToolsConnectionId

			listCredentialPublicSynonymsResponse, err := databaseToolsRuntimeClient.ListCredentialPublicSynonyms(context.Background(), listCredentialPublicSynonymsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting DatabaseToolsConnectionCredentialPublicSynonym list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, databaseToolsConnectionCredentialPublicSynonym := range listCredentialPublicSynonymsResponse.Items {
				if databaseToolsConnectionCredentialPublicSynonym.Key == nil {
					continue
				}
				id := databaseToolsConnectionId + "/" + credentialKey + "/" + *databaseToolsConnectionCredentialPublicSynonym.Key
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsConnectionCredentialPublicSynonymId", id)
			}

		}
	}
	return resourceIds, nil
}
