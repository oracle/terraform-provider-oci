// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsDatabaseToolsMcpServerRequiredOnlyResource = DatabaseToolsDatabaseToolsMcpServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpServerRepresentation)

	DatabaseToolsDatabaseToolsMcpServerResourceConfig = DatabaseToolsDatabaseToolsMcpServerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpServerRepresentation)

	DatabaseToolsDatabaseToolsMcpServerSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_mcp_server_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_mcp_server.test_database_tools_mcp_server.id}`},
	}

	DatabaseToolsDatabaseToolsMcpServerDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.database_tools_connection_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `McpServer1`, Update: `displayName2`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`DEFAULT`}},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsMcpServerDataSourceFilterRepresentation}}
	DatabaseToolsDatabaseToolsMcpServerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_mcp_server.test_database_tools_mcp_server.id}`}},
	}

	DatabaseToolsDatabaseToolsMcpServerRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_tools_connection_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_connection_id}`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `McpServer1`, Update: `displayName2`},
		"domain_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.domain_id}`},
		"storage":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsMcpServerStorageRepresentation},
		"type":                            acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"access_token_expiry_in_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `61`},
		"custom_roles":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsMcpServerCustomRolesRepresentation},
		"description":                     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"refresh_token_expiry_in_seconds": acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `61`},
		"runtime_identity":                acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATED_PRINCIPAL`},
	}
	DatabaseToolsDatabaseToolsMcpServerStorageRepresentation = map[string]interface{}{
		// Provider schema for storage.type allows only NONE / OBJECT_STORAGE.
		// Keep this test minimal: do not require any bucket configuration.
		"type": acctest.Representation{RepType: acctest.Required, Create: `NONE`, Update: `NONE`},
	}
	DatabaseToolsDatabaseToolsMcpServerCustomRolesRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `McpServer1`, Update: `displayName2`},
	}
	DatabaseToolsDatabaseToolsMcpServerStorageBucketRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Optional, Create: `bucket`, Update: `bucket2`},
		"namespace": acctest.Representation{RepType: acctest.Optional, Create: `namespace`, Update: `namespace2`},
	}

	DatabaseToolsDatabaseToolsMcpServerResourceDependencies = ""
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsMcpServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsMcpServerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	databaseToolsConnectionId := utils.GetEnvSettingWithBlankDefault("database_tools_connection_ocid")
	databaseToolsConnectionIdVariableStr := fmt.Sprintf("variable \"database_tools_connection_id\" { default = \"%s\" }\n", databaseToolsConnectionId)

	domainId := utils.GetEnvSettingWithBlankDefault("domain_ocid")
	domainIdVariableStr := fmt.Sprintf("variable \"domain_id\" { default = \"%s\" }\n", domainId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_mcp_server.test_database_tools_mcp_server"
	datasourceName := "data.oci_database_tools_database_tools_mcp_servers.test_database_tools_mcp_servers"
	singularDatasourceName := "data.oci_database_tools_database_tools_mcp_server.test_database_tools_mcp_server"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+databaseToolsConnectionIdVariableStr+domainIdVariableStr+DatabaseToolsDatabaseToolsMcpServerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpServerRepresentation), "databasetools", "databaseToolsMcpServer", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsMcpServerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsConnectionIdVariableStr + domainIdVariableStr + DatabaseToolsDatabaseToolsMcpServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "McpServer1"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage.0.bucket.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "storage.0.type", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsConnectionIdVariableStr + domainIdVariableStr + DatabaseToolsDatabaseToolsMcpServerResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + databaseToolsConnectionIdVariableStr + domainIdVariableStr + DatabaseToolsDatabaseToolsMcpServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsMcpServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_token_expiry_in_seconds", "60"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.0.display_name", "McpServer1"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "McpServer1"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "refresh_token_expiry_in_seconds", "60"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage.0.bucket.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "storage.0.type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + databaseToolsConnectionIdVariableStr + domainIdVariableStr + DatabaseToolsDatabaseToolsMcpServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsMcpServerRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_token_expiry_in_seconds", "60"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.0.display_name", "McpServer1"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "McpServer1"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "refresh_token_expiry_in_seconds", "60"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage.0.bucket.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "storage.0.type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + databaseToolsConnectionIdVariableStr + domainIdVariableStr + DatabaseToolsDatabaseToolsMcpServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_token_expiry_in_seconds", "61"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "custom_roles.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "refresh_token_expiry_in_seconds", "61"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "storage.0.bucket.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "storage.0.type", "NONE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_servers", "test_database_tools_mcp_servers", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpServerDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsConnectionIdVariableStr + domainIdVariableStr + DatabaseToolsDatabaseToolsMcpServerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsMcpServerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_server_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_mcp_server_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_mcp_server", "test_database_tools_mcp_server", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsMcpServerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsConnectionIdVariableStr + domainIdVariableStr + DatabaseToolsDatabaseToolsMcpServerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_mcp_server_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "access_token_expiry_in_seconds", "61"),
				resource.TestCheckResourceAttr(singularDatasourceName, "built_in_roles.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_roles.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_roles.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_roles.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_app_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "refresh_token_expiry_in_seconds", "61"),
				resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage.0.bucket.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage.0.type", "NONE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "DEFAULT"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + databaseToolsConnectionIdVariableStr + domainIdVariableStr + DatabaseToolsDatabaseToolsMcpServerRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseToolsDatabaseToolsMcpServerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_mcp_server" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsMcpServerRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsMcpServerId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsMcpServer(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.DatabaseToolsMcpServerLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("DatabaseToolsDatabaseToolsMcpServer") {
		resource.AddTestSweepers("DatabaseToolsDatabaseToolsMcpServer", &resource.Sweeper{
			Name:         "DatabaseToolsDatabaseToolsMcpServer",
			Dependencies: acctest.DependencyGraph["databaseToolsMcpServer"],
			F:            sweepDatabaseToolsDatabaseToolsMcpServerResource,
		})
	}
}

func sweepDatabaseToolsDatabaseToolsMcpServerResource(compartment string) error {
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	databaseToolsMcpServerIds, err := getDatabaseToolsDatabaseToolsMcpServerIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsMcpServerId := range databaseToolsMcpServerIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsMcpServerId]; !ok {
			deleteDatabaseToolsMcpServerRequest := oci_database_tools.DeleteDatabaseToolsMcpServerRequest{}

			deleteDatabaseToolsMcpServerRequest.DatabaseToolsMcpServerId = &databaseToolsMcpServerId

			deleteDatabaseToolsMcpServerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsMcpServer(context.Background(), deleteDatabaseToolsMcpServerRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsMcpServer %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsMcpServerId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseToolsMcpServerId, DatabaseToolsDatabaseToolsMcpServerSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseToolsDatabaseToolsMcpServerSweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsDatabaseToolsMcpServerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsMcpServerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()

	listDatabaseToolsMcpServersRequest := oci_database_tools.ListDatabaseToolsMcpServersRequest{}
	listDatabaseToolsMcpServersRequest.CompartmentId = &compartmentId
	listDatabaseToolsMcpServersRequest.LifecycleState = oci_database_tools.ListDatabaseToolsMcpServersLifecycleStateActive
	listDatabaseToolsMcpServersResponse, err := databaseToolsClient.ListDatabaseToolsMcpServers(context.Background(), listDatabaseToolsMcpServersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsMcpServer list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsMcpServer := range listDatabaseToolsMcpServersResponse.Items {
		id := *databaseToolsMcpServer.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsMcpServerId", id)
	}
	return resourceIds, nil
}

func DatabaseToolsDatabaseToolsMcpServerSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseToolsMcpServerResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsMcpServerResponse); ok {
		return databaseToolsMcpServerResponse.GetLifecycleState() != oci_database_tools.DatabaseToolsMcpServerLifecycleStateDeleted
	}
	return false
}

func DatabaseToolsDatabaseToolsMcpServerSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseToolsClient().GetDatabaseToolsMcpServer(context.Background(), oci_database_tools.GetDatabaseToolsMcpServerRequest{
		DatabaseToolsMcpServerId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
