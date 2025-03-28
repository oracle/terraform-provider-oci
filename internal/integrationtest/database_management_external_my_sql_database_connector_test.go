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
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalMySqlDatabaseConnectorResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_connector", "test_external_my_sql_database_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseConnectorRepresentation)

	DatabaseManagementExternalMySqlDatabaseConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"external_my_sql_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_my_sql_database_connector.test_external_my_sql_database_connector.id}`},
	}

	DatabaseManagementExternalMySqlDatabaseConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DatabaseManagementExternalMySqlDatabaseConnectorRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connector_details":               acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalMySqlDatabaseConnectorConnectorDetailsRepresentation},
		"is_test_connection_param":        acctest.Representation{RepType: acctest.Required, Create: `false`},
		"check_connection_status_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	DatabaseManagementExternalMySqlDatabaseConnectorConnectorDetailsRepresentation = map[string]interface{}{
		"credential_type":      acctest.Representation{RepType: acctest.Required, Create: `MYSQL_EXTERNAL_NON_SSL_CREDENTIALS`, Update: `MYSQL_EXTERNAL_SSL_CREDENTIALS`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `EXAMPLE-displayName1032-Value`, Update: `displayName2`},
		"external_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"host_name":            acctest.Representation{RepType: acctest.Required, Create: `hostName`, Update: `hostName2`},
		"macs_agent_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.macs_agent_id}`},
		"network_protocol":     acctest.Representation{RepType: acctest.Required, Create: `TCP`, Update: `TCPS`},
		"port":                 acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"ssl_secret_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.vault_secret_id}`},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalMySqlDatabaseConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalMySqlDatabaseConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	macs_agent_id := utils.GetEnvSettingWithBlankDefault("macs_agent_id")
	macsAgentIdVariableStr := fmt.Sprintf("variable \"macs_agent_id\" { default = \"%s\" }\n", macs_agent_id)

	test_managed_database_id := utils.GetEnvSettingWithBlankDefault("test_managed_database_id")
	testManagedDatabaseIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", test_managed_database_id)

	vault_secret_id := utils.GetEnvSettingWithBlankDefault("vault_secret_id")
	sslSecretIdVariableStr := fmt.Sprintf("variable \"vault_secret_id\" { default = \"%s\" }\n", vault_secret_id)

	resourceName := "oci_database_management_external_my_sql_database_connector.test_external_my_sql_database_connector"
	datasourceName := "data.oci_database_management_external_my_sql_database_connectors.test_external_my_sql_database_connectors"
	singularDatasourceName := "data.oci_database_management_external_my_sql_database_connector.test_external_my_sql_database_connector"
	//var resId string
	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+macsAgentIdVariableStr+testManagedDatabaseIdVariableStr+sslSecretIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_connector", "test_external_my_sql_database_connector", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseConnectorRepresentation), "databasemanagement", "externalMySqlDatabaseConnector", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementExternalMySqlDatabaseConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + macsAgentIdVariableStr + testManagedDatabaseIdVariableStr + sslSecretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_connector", "test_external_my_sql_database_connector", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.credential_type", "MYSQL_EXTERNAL_NON_SSL_CREDENTIALS"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.display_name", "EXAMPLE-displayName1032-Value"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_details.0.external_database_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.host_name", "hostName"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_details.0.macs_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.network_protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_details.0.ssl_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "is_test_connection_param", "false"),

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
			Config: config + compartmentIdVariableStr + macsAgentIdVariableStr + testManagedDatabaseIdVariableStr + sslSecretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_connector", "test_external_my_sql_database_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connector_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.credential_type", "MYSQL_EXTERNAL_SSL_CREDENTIALS"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_details.0.external_database_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.host_name", "hostName2"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_details.0.macs_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.network_protocol", "TCPS"),
				resource.TestCheckResourceAttr(resourceName, "connector_details.0.port", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_details.0.ssl_secret_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_test_connection_param", "false"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_my_sql_database_connectors", "test_external_my_sql_database_connectors", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + macsAgentIdVariableStr + testManagedDatabaseIdVariableStr + sslSecretIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database_connector", "test_external_my_sql_database_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "my_sql_connector_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_my_sql_database_connector", "test_external_my_sql_database_connector", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + macsAgentIdVariableStr + testManagedDatabaseIdVariableStr + sslSecretIdVariableStr + DatabaseManagementExternalMySqlDatabaseConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_my_sql_database_connector_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connector_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "credential_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "macs_agent_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_protocol"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "port"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_database"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_database_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ssl_secret_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ssl_secret_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_connection_status_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalMySqlDatabaseConnectorResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"connector_details",
				"is_test_connection_param",
				"check_connection_status_trigger",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementExternalMySqlDatabaseConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_external_my_sql_database_connector" {
			noResourceFound = false
			request := oci_database_management.GetExternalMySqlDatabaseConnectorRequest{}

			tmp := rs.Primary.ID
			request.ExternalMySqlDatabaseConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetExternalMySqlDatabaseConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.LifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
	if !acctest.InSweeperExcludeList("DatabaseManagementExternalMySqlDatabaseConnector") {
		resource.AddTestSweepers("DatabaseManagementExternalMySqlDatabaseConnector", &resource.Sweeper{
			Name:         "DatabaseManagementExternalMySqlDatabaseConnector",
			Dependencies: acctest.DependencyGraph["externalMySqlDatabaseConnector"],
			F:            sweepDatabaseManagementExternalMySqlDatabaseConnectorResource,
		})
	}
}

func sweepDatabaseManagementExternalMySqlDatabaseConnectorResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	externalMySqlDatabaseConnectorIds, err := getDatabaseManagementExternalMySqlDatabaseConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, externalMySqlDatabaseConnectorId := range externalMySqlDatabaseConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[externalMySqlDatabaseConnectorId]; !ok {
			deleteExternalMySqlDatabaseConnectorRequest := oci_database_management.DeleteExternalMySqlDatabaseConnectorRequest{}

			deleteExternalMySqlDatabaseConnectorRequest.ExternalMySqlDatabaseConnectorId = &externalMySqlDatabaseConnectorId

			deleteExternalMySqlDatabaseConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteExternalMySqlDatabaseConnector(context.Background(), deleteExternalMySqlDatabaseConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalMySqlDatabaseConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalMySqlDatabaseConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalMySqlDatabaseConnectorId, DatabaseManagementExternalMySqlDatabaseConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementExternalMySqlDatabaseConnectorSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementExternalMySqlDatabaseConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalMySqlDatabaseConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listMySqlDatabaseConnectorsRequest := oci_database_management.ListMySqlDatabaseConnectorsRequest{}
	listMySqlDatabaseConnectorsRequest.CompartmentId = &compartmentId
	listMySqlDatabaseConnectorsResponse, err := dbManagementClient.ListMySqlDatabaseConnectors(context.Background(), listMySqlDatabaseConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalMySqlDatabaseConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalMySqlDatabaseConnector := range listMySqlDatabaseConnectorsResponse.Items {
		id := *externalMySqlDatabaseConnector.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalMySqlDatabaseConnectorId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementExternalMySqlDatabaseConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalMySqlDatabaseConnectorResponse, ok := response.Response.(oci_database_management.GetExternalMySqlDatabaseConnectorResponse); ok {
		return externalMySqlDatabaseConnectorResponse.LifecycleState != oci_database_management.LifecycleStatesDeleted
	}
	return false
}

func DatabaseManagementExternalMySqlDatabaseConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetExternalMySqlDatabaseConnector(context.Background(), oci_database_management.GetExternalMySqlDatabaseConnectorRequest{
		ExternalMySqlDatabaseConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
