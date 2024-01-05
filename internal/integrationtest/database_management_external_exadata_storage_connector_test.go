// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	DatabaseManagementExternalExadataStorageConnectorRequiredOnlyResource = DatabaseManagementExternalExadataStorageConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_connector", "test_external_exadata_storage_connector", acctest.Required, acctest.Create, DatabaseManagementExternalExadataStorageConnectorRepresentation)
	DatabaseManagementExternalExadataStorageConnectorResourceConfig = DatabaseManagementExternalExadataStorageConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_connector", "test_external_exadata_storage_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalExadataStorageConnectorRepresentation)

	DatabaseManagementDatabaseManagementExternalExadataStorageConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"external_exadata_storage_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_exadata_storage_connector.test_external_exadata_storage_connector.id}`},
	}

	DatabaseManagementDatabaseManagementExternalExadataStorageConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_exadata_infra_id}`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `slcm21celadm01-connUpdate`}}

	DatabaseManagementExternalExadataStorageConnectorRepresentation = map[string]interface{}{
		"agent_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.connector_agent_id}`},
		"connection_uri":    acctest.Representation{RepType: acctest.Required, Create: `${var.connector_connection_uri}`},
		"connector_name":    acctest.Representation{RepType: acctest.Required, Create: `slcm21celadm01-conn`, Update: `slcm21celadm01-connUpdate`},
		"credential_info":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalExadataStorageConnectorCredentialInfoRepresentation},
		"storage_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_storage_server_id}`},
	}
	DatabaseManagementExternalExadataStorageConnectorCredentialInfoRepresentation = map[string]interface{}{
		"password":                 acctest.Representation{RepType: acctest.Required, Create: `${var.connector_password}`},
		"username":                 acctest.Representation{RepType: acctest.Required, Create: `${var.connector_username}`},
		"ssl_trust_store_location": acctest.Representation{RepType: acctest.Optional, Create: `${var.connector_ssl_trust_store_location}`},
		"ssl_trust_store_password": acctest.Representation{RepType: acctest.Optional, Create: `${var.connector_ssl_trust_store_password}`},
		"ssl_trust_store_type":     acctest.Representation{RepType: acctest.Optional, Create: `${var.connector_ssl_trust_store_type}`},
	}
	DatabaseManagementExternalExadataStorageConnectorResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalExadataStorageConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalExadataStorageConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	connectorPassword := utils.GetEnvSettingWithBlankDefault("connector_password")
	connectorPasswordVariableStr := fmt.Sprintf("variable \"connector_password\" { default = \"%s\" }\n", connectorPassword)

	connectorAgentId := utils.GetEnvSettingWithBlankDefault("connector_agent_id")
	connectorAgentIdVariableStr := fmt.Sprintf("variable \"connector_agent_id\" { default = \"%s\" }\n", connectorAgentId)

	connectorUserName := utils.GetEnvSettingWithBlankDefault("connector_username")
	connectorUserNameVariableStr := fmt.Sprintf("variable \"connector_username\" { default = \"%s\" }\n", connectorUserName)

	connectorSslTrustStoreLocation := utils.GetEnvSettingWithBlankDefault("connector_ssl_trust_store_location")
	connectorSslTrustStoreLocationVariableStr := fmt.Sprintf("variable \"connector_ssl_trust_store_location\" { default = \"%s\" }\n", connectorSslTrustStoreLocation)

	connectorSslTrustStorePassword := utils.GetEnvSettingWithBlankDefault("connector_ssl_trust_store_password")
	connectorSslTrustStorePasswordVariableStr := fmt.Sprintf("variable \"connector_ssl_trust_store_password\" { default = \"%s\" }\n", connectorSslTrustStorePassword)

	connectorSslTrustStoreType := utils.GetEnvSettingWithBlankDefault("connector_ssl_trust_store_type")
	connectorSslTrustStoreTypeVariableStr := fmt.Sprintf("variable \"connector_ssl_trust_store_type\" { default = \"%s\" }\n", connectorSslTrustStoreType)

	connectorConnectionUri := utils.GetEnvSettingWithBlankDefault("connector_connection_uri")
	connectorConnectionUriVariableStr := fmt.Sprintf("variable \"connector_connection_uri\" { default = \"%s\" }\n", connectorConnectionUri)

	connectorStorageServerId := utils.GetEnvSettingWithBlankDefault("connector_storage_server_id")
	connectorStorageServerIdVariableStr := fmt.Sprintf("variable \"connector_storage_server_id\" { default = \"%s\" }\n", connectorStorageServerId)

	connectorExadataInfraId := utils.GetEnvSettingWithBlankDefault("connector_exadata_infra_id")
	connectorExadataInfraIdVariableStr := fmt.Sprintf("variable \"connector_exadata_infra_id\" { default = \"%s\" }\n", connectorExadataInfraId)

	resourceName := "oci_database_management_external_exadata_storage_connector.test_external_exadata_storage_connector"
	datasourceName := "data.oci_database_management_external_exadata_storage_connectors.test_external_exadata_storage_connectors"
	singularDatasourceName := "data.oci_database_management_external_exadata_storage_connector.test_external_exadata_storage_connector"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementExternalExadataStorageConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_connector", "test_external_exadata_storage_connector", acctest.Optional, acctest.Create, DatabaseManagementExternalExadataStorageConnectorRepresentation), "databasemanagement", "externalExadataStorageConnector", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementExternalExadataStorageConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				connectorStorageServerIdVariableStr +
				connectorPasswordVariableStr +
				connectorAgentIdVariableStr +
				connectorUserNameVariableStr +
				connectorSslTrustStoreLocationVariableStr +
				connectorSslTrustStorePasswordVariableStr +
				connectorSslTrustStoreTypeVariableStr +
				connectorConnectionUriVariableStr +
				DatabaseManagementExternalExadataStorageConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_connector", "test_external_exadata_storage_connector", acctest.Optional, acctest.Create, DatabaseManagementExternalExadataStorageConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_uri"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_name"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.#"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.password"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.username"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_server_id"),

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
			Config: config + compartmentIdVariableStr +
				connectorExadataInfraIdVariableStr +
				connectorStorageServerIdVariableStr +
				connectorPasswordVariableStr +
				connectorAgentIdVariableStr +
				connectorUserNameVariableStr +
				connectorSslTrustStoreLocationVariableStr +
				connectorSslTrustStorePasswordVariableStr +
				connectorSslTrustStoreTypeVariableStr +
				connectorConnectionUriVariableStr +
				DatabaseManagementExternalExadataStorageConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_connector", "test_external_exadata_storage_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalExadataStorageConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_uri"),
				resource.TestCheckResourceAttr(resourceName, "connector_name", "slcm21celadm01-connUpdate"),
				resource.TestCheckResourceAttr(resourceName, "credential_info.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_server_id"),

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
				connectorStorageServerIdVariableStr +
				connectorExadataInfraIdVariableStr +
				connectorPasswordVariableStr +
				connectorAgentIdVariableStr +
				connectorUserNameVariableStr +
				connectorSslTrustStoreLocationVariableStr +
				connectorSslTrustStorePasswordVariableStr +
				connectorSslTrustStoreTypeVariableStr +
				connectorConnectionUriVariableStr +
				DatabaseManagementExternalExadataStorageConnectorResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_connectors", "test_external_exadata_storage_connectors", acctest.Optional, acctest.Update, DatabaseManagementDatabaseManagementExternalExadataStorageConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalExadataStorageConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_storage_connector", "test_external_exadata_storage_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalExadataStorageConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "slcm21celadm01-connUpdate"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_exadata_infrastructure_id"),

				resource.TestCheckResourceAttr(datasourceName, "external_exadata_storage_connector_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_exadata_storage_connector_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				connectorStorageServerIdVariableStr +
				connectorExadataInfraIdVariableStr +
				connectorPasswordVariableStr +
				connectorAgentIdVariableStr +
				connectorUserNameVariableStr +
				connectorSslTrustStoreLocationVariableStr +
				connectorSslTrustStorePasswordVariableStr +
				connectorSslTrustStoreTypeVariableStr +
				connectorConnectionUriVariableStr +
				DatabaseManagementExternalExadataStorageConnectorResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_exadata_storage_connector", "test_external_exadata_storage_connector", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalExadataStorageConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalExadataStorageConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_exadata_storage_connector_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_uri"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalExadataStorageConnectorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"connector_name",
				"credential_info",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementExternalExadataStorageConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_external_exadata_storage_connector" {
			noResourceFound = false
			request := oci_database_management.GetExternalExadataStorageConnectorRequest{}

			tmp := rs.Primary.ID
			request.ExternalExadataStorageConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetExternalExadataStorageConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.DbmResourceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementExternalExadataStorageConnector") {
		resource.AddTestSweepers("DatabaseManagementExternalExadataStorageConnector", &resource.Sweeper{
			Name:         "DatabaseManagementExternalExadataStorageConnector",
			Dependencies: acctest.DependencyGraph["externalExadataStorageConnector"],
			F:            sweepDatabaseManagementExternalExadataStorageConnectorResource,
		})
	}
}

func sweepDatabaseManagementExternalExadataStorageConnectorResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	externalExadataStorageConnectorIds, err := getDatabaseManagementExternalExadataStorageConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, externalExadataStorageConnectorId := range externalExadataStorageConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[externalExadataStorageConnectorId]; !ok {
			deleteExternalExadataStorageConnectorRequest := oci_database_management.DeleteExternalExadataStorageConnectorRequest{}

			deleteExternalExadataStorageConnectorRequest.ExternalExadataStorageConnectorId = &externalExadataStorageConnectorId

			deleteExternalExadataStorageConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteExternalExadataStorageConnector(context.Background(), deleteExternalExadataStorageConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalExadataStorageConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalExadataStorageConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalExadataStorageConnectorId, DatabaseManagementExternalExadataStorageConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementExternalExadataStorageConnectorSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementExternalExadataStorageConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalExadataStorageConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listExternalExadataStorageConnectorsRequest := oci_database_management.ListExternalExadataStorageConnectorsRequest{}
	listExternalExadataStorageConnectorsRequest.CompartmentId = &compartmentId

	externalExadataInfrastructureIds, error := getDatabaseManagementExternalExadataInfrastructureIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting externalExadataInfrastructureId required for ExternalExadataStorageConnector resource requests \n")
	}
	for _, externalExadataInfrastructureId := range externalExadataInfrastructureIds {
		listExternalExadataStorageConnectorsRequest.ExternalExadataInfrastructureId = &externalExadataInfrastructureId

		listExternalExadataStorageConnectorsResponse, err := dbManagementClient.ListExternalExadataStorageConnectors(context.Background(), listExternalExadataStorageConnectorsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ExternalExadataStorageConnector list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, externalExadataStorageConnector := range listExternalExadataStorageConnectorsResponse.Items {
			id := *externalExadataStorageConnector.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalExadataStorageConnectorId", id)
		}

	}
	return resourceIds, nil
}

func DatabaseManagementExternalExadataStorageConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalExadataStorageConnectorResponse, ok := response.Response.(oci_database_management.GetExternalExadataStorageConnectorResponse); ok {
		return externalExadataStorageConnectorResponse.GetLifecycleState() != oci_database_management.DbmResourceLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementExternalExadataStorageConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetExternalExadataStorageConnector(context.Background(), oci_database_management.GetExternalExadataStorageConnectorRequest{
		ExternalExadataStorageConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
