// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
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
	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResourceConfig = DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_execute_grantee", "test_database_tools_connection_credential_execute_grantee", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeRepresentation)

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeSingularDataSourceRepresentation = map[string]interface{}{
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_credential_key}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"execute_grantee_key":          acctest.Representation{RepType: acctest.Required, Create: `${var.execute_grantee_key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceRepresentation = map[string]interface{}{
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_credential_key}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceFilterRepresentation}}
	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.execute_grantee_key}`}},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeRepresentation = map[string]interface{}{
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_credential_key}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `${var.execute_grantee_key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResourceDependencies = ""
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResource_basic")
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
	executeGranteeKey := utils.GetEnvSettingWithDefault("execute_grantee_key", utils.GetEnvSettingWithDefault("user_key", "APEX_240200"))
	allVars := databaseToolsRuntimeExistingExecuteGranteeVariables()

	resourceName := "oci_database_tools_runtime_database_tools_connection_credential_execute_grantee.test_database_tools_connection_credential_execute_grantee"
	datasourceName := "data.oci_database_tools_runtime_database_tools_connection_credential_execute_grantees.test_database_tools_connection_credential_execute_grantees"
	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_connection_credential_execute_grantee.test_database_tools_connection_credential_execute_grantee"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+allVars+DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_execute_grantee", "test_database_tools_connection_credential_execute_grantee", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeRepresentation), "databasetoolsruntime", "databaseToolsConnectionCredentialExecuteGrantee", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_execute_grantee", "test_database_tools_connection_credential_execute_grantee", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credential_key", existingDatabaseToolsConnectionCredentialKey),
				resource.TestCheckResourceAttr(resourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(resourceName, "key", executeGranteeKey),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_execute_grantees", "test_database_tools_connection_credential_execute_grantees", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSourceRepresentation) +
				allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_execute_grantee", "test_database_tools_connection_credential_execute_grantee", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "credential_key", existingDatabaseToolsConnectionCredentialKey),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),

				resource.TestCheckResourceAttr(datasourceName, "credential_execute_grantee_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "credential_execute_grantee_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential_execute_grantee", "test_database_tools_connection_credential_execute_grantee", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeSingularDataSourceRepresentation) +
				allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_key", existingDatabaseToolsConnectionCredentialKey),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(singularDatasourceName, "execute_grantee_key", executeGranteeKey),

				resource.TestCheckResourceAttr(singularDatasourceName, "key", executeGranteeKey),
			),
		},
		// verify resource import
		{
			Config:            config + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateCheck: checkImportedDatabaseToolsRuntimeCompositeID(
				resourceName,
				parseDatabaseToolsRuntimeConnectionCredentialExecuteGranteeCompositeIDToAttributes,
			),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func databaseToolsRuntimeExistingExecuteGranteeVariables() string {
	return terraformStringVariable("existing_database_tools_connection_id",
		utils.GetEnvSettingWithDefault("existing_database_tools_connection_id",
			utils.GetEnvSettingWithDefault("database_tools_connection_id",
				utils.GetEnvSettingWithBlankDefault("database_tools_connection_ocid")))) +
		terraformStringVariable("existing_database_tools_connection_credential_key",
			utils.GetEnvSettingWithDefault("existing_database_tools_connection_credential_key",
				utils.GetEnvSettingWithDefault("database_tools_connection_credential_key",
					utils.GetEnvSettingWithDefault("database_tools_connection_credential_public_synonym_credential_key",
						utils.GetEnvSettingWithDefault("user_credential_key", "credentialKey"))))) +
		terraformStringVariable("execute_grantee_key",
			utils.GetEnvSettingWithDefault("execute_grantee_key",
				utils.GetEnvSettingWithDefault("user_key", "APEX_240200")))
}

func testAccCheckDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsRuntimeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_runtime_database_tools_connection_credential_execute_grantee" {
			noResourceFound = false
			request := oci_database_tools_runtime.GetCredentialExecuteGranteeRequest{}

			if value, ok := rs.Primary.Attributes["credential_key"]; ok {
				request.CredentialKey = &value
			}

			if value, ok := rs.Primary.Attributes["database_tools_connection_id"]; ok {
				request.DatabaseToolsConnectionId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ExecuteGranteeKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")

			_, err := client.GetCredentialExecuteGrantee(context.Background(), request)

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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGrantee") {
		resource.AddTestSweepers("DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGrantee", &resource.Sweeper{
			Name:         "DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGrantee",
			Dependencies: acctest.DependencyGraph["databaseToolsConnectionCredentialExecuteGrantee"],
			F:            sweepDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResource,
		})
	}
}

func sweepDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResource(compartment string) error {
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()
	databaseToolsConnectionCredentialExecuteGranteeIds, err := getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsConnectionCredentialExecuteGranteeId := range databaseToolsConnectionCredentialExecuteGranteeIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsConnectionCredentialExecuteGranteeId]; !ok {
			deleteCredentialExecuteGranteeRequest := oci_database_tools_runtime.DeleteCredentialExecuteGranteeRequest{}

			deleteCredentialExecuteGranteeRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")
			_, error := databaseToolsRuntimeClient.DeleteCredentialExecuteGrantee(context.Background(), deleteCredentialExecuteGranteeRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsConnectionCredentialExecuteGrantee %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsConnectionCredentialExecuteGranteeId, error)
				continue
			}
		}
	}
	return nil
}

func getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsConnectionCredentialExecuteGranteeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()

	listCredentialExecuteGranteesRequest := oci_database_tools_runtime.ListCredentialExecuteGranteesRequest{}

	databaseToolsConnectionCredentialIds, error := getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting credential ids required for DatabaseToolsConnectionCredentialExecuteGrantee resource requests \n")
	}
	for _, databaseToolsConnectionCredentialId := range databaseToolsConnectionCredentialIds {
		parts := strings.SplitN(databaseToolsConnectionCredentialId, "/", 2)
		if len(parts) != 2 {
			continue
		}
		credentialKey := parts[1]
		listCredentialExecuteGranteesRequest.CredentialKey = &credentialKey

		databaseToolsConnectionIds, error := getDatabaseToolsConnectionIds(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting databaseToolsConnectionId required for DatabaseToolsConnectionCredentialExecuteGrantee resource requests \n")
		}
		for _, databaseToolsConnectionId := range databaseToolsConnectionIds {
			listCredentialExecuteGranteesRequest.DatabaseToolsConnectionId = &databaseToolsConnectionId

			listCredentialExecuteGranteesResponse, err := databaseToolsRuntimeClient.ListCredentialExecuteGrantees(context.Background(), listCredentialExecuteGranteesRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting DatabaseToolsConnectionCredentialExecuteGrantee list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, databaseToolsConnectionCredentialExecuteGrantee := range listCredentialExecuteGranteesResponse.Items {
				if databaseToolsConnectionCredentialExecuteGrantee.Key == nil {
					continue
				}
				id := databaseToolsConnectionId + "/" + credentialKey + "/" + *databaseToolsConnectionCredentialExecuteGrantee.Key
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsConnectionCredentialExecuteGranteeId", id)
			}

		}
	}
	return resourceIds, nil
}
