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
	ignoreDatabaseManagementCloudExadataInfrastructureConnectorRunDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `system_tags`, `freeform_tags`}},
	}

	DatabaseManagementCloudExadataStorageConnectorRequiredOnlyResource = DatabaseManagementCloudExadataStorageConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connector", "test_cloud_exadata_storage_connector", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageConnectorRepresentation)

	DatabaseManagementCloudExadataStorageConnectorResourceConfig = DatabaseManagementCloudExadataStorageConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connector", "test_cloud_exadata_storage_connector", acctest.Optional, acctest.Update, DatabaseManagementCloudExadataStorageConnectorRepresentation)

	DatabaseManagementCloudExadataStorageConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_exadata_storage_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_exadata_storage_connector.test_cloud_exadata_storage_connector.id}`},
	}

	DatabaseManagementCloudExadataStorageConnectorDataSourceRepresentation = map[string]interface{}{
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_exadata_infra_id}`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `slcqan01celadm03-conn`},
		"filter":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementCloudExadataStorageConnectorDataSourceFilterRepresentation}}
	DatabaseManagementCloudExadataStorageConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_exadata_storage_connector.test_cloud_exadata_storage_connector.id}`}},
	}

	DatabaseManagementCloudExadataStorageConnectorRepresentation = map[string]interface{}{
		"agent_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.connector_agent_id}`},
		"connection_uri":    acctest.Representation{RepType: acctest.Required, Create: `${var.connector_connection_uri}`},
		"credential_info":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementCloudExadataStorageConnectorCredentialInfoRepresentation},
		"storage_server_id": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_storage_server_id}`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `slcqan01celadm03-conn`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDatabaseManagementCloudExadataInfrastructureConnectorRunDefinedTagsRepresentation},
	}
	DatabaseManagementCloudExadataStorageConnectorCredentialInfoRepresentation = map[string]interface{}{
		"password":                 acctest.Representation{RepType: acctest.Required, Create: `${var.connector_password}`},
		"username":                 acctest.Representation{RepType: acctest.Required, Create: `${var.connector_username}`},
		"ssl_trust_store_location": acctest.Representation{RepType: acctest.Required, Create: `/etc/pki/ca-trust/extracted/java/cacerts`},
		"ssl_trust_store_password": acctest.Representation{RepType: acctest.Required, Create: `${var.connector_ssl_trust_store_password}`},
		"ssl_trust_store_type":     acctest.Representation{RepType: acctest.Required, Create: `${var.connector_ssl_trust_store_type}`},
	}

	DatabaseManagementCloudExadataStorageConnectorResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudExadataStorageConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudExadataStorageConnectorResource_basic")
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

	resourceName := "oci_database_management_cloud_exadata_storage_connector.test_cloud_exadata_storage_connector"
	datasourceName := "data.oci_database_management_cloud_exadata_storage_connectors.test_cloud_exadata_storage_connectors"
	singularDatasourceName := "data.oci_database_management_cloud_exadata_storage_connector.test_cloud_exadata_storage_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementCloudExadataStorageConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connector", "test_cloud_exadata_storage_connector", acctest.Optional, acctest.Create, DatabaseManagementCloudExadataStorageConnectorRepresentation), "databasemanagement", "cloudExadataStorageConnector", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementCloudExadataStorageConnectorDestroy, []resource.TestStep{
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
				DatabaseManagementCloudExadataStorageConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connector", "test_cloud_exadata_storage_connector", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_uri", "https://slcqan01celadm03.us.oracle.com/MS/RESTService/"),
				resource.TestCheckResourceAttr(resourceName, "credential_info.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credential_info.0.password", "Ka7#Rn9-B0w-H2n#EX_"),
				resource.TestCheckResourceAttr(resourceName, "credential_info.0.username", "cloud_user_cluclu03711"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_server_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementCloudExadataStorageConnectorResourceDependencies,
		},
		// verify Create with optionals
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
				DatabaseManagementCloudExadataStorageConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connector", "test_cloud_exadata_storage_connector", acctest.Optional, acctest.Create, DatabaseManagementCloudExadataStorageConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_uri"),
				resource.TestCheckResourceAttr(resourceName, "credential_info.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.password"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.ssl_trust_store_location"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.ssl_trust_store_password"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.ssl_trust_store_type"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.username"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				DatabaseManagementCloudExadataStorageConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connector", "test_cloud_exadata_storage_connector", acctest.Optional, acctest.Update, DatabaseManagementCloudExadataStorageConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_uri"),
				resource.TestCheckResourceAttr(resourceName, "credential_info.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.password"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.ssl_trust_store_location"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.ssl_trust_store_password"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.ssl_trust_store_type"),
				resource.TestCheckResourceAttrSet(resourceName, "credential_info.0.username"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connectors", "test_cloud_exadata_storage_connectors", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageConnectorDataSourceRepresentation) +
				compartmentIdVariableStr +
				connectorExadataInfraIdVariableStr +
				connectorStorageServerIdVariableStr +
				connectorPasswordVariableStr +
				connectorAgentIdVariableStr +
				connectorUserNameVariableStr +
				connectorSslTrustStoreLocationVariableStr +
				connectorSslTrustStorePasswordVariableStr +
				connectorSslTrustStoreTypeVariableStr +
				connectorConnectionUriVariableStr +
				DatabaseManagementCloudExadataStorageConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connector", "test_cloud_exadata_storage_connector", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_storage_connector_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_exadata_storage_connector_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_exadata_storage_connector", "test_cloud_exadata_storage_connector", acctest.Required, acctest.Create, DatabaseManagementCloudExadataStorageConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr +
				connectorExadataInfraIdVariableStr +
				connectorStorageServerIdVariableStr +
				connectorPasswordVariableStr +
				connectorAgentIdVariableStr +
				connectorUserNameVariableStr +
				connectorSslTrustStoreLocationVariableStr +
				connectorSslTrustStorePasswordVariableStr +
				connectorSslTrustStoreTypeVariableStr +
				connectorConnectionUriVariableStr +
				DatabaseManagementCloudExadataStorageConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_exadata_storage_connector_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_uri"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudExadataStorageConnectorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"credential_info", "defined_tags", "defined_tags.example-tag-namespace-all.example-tag",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementCloudExadataStorageConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_cloud_exadata_storage_connector" {
			noResourceFound = false
			request := oci_database_management.GetCloudExadataStorageConnectorRequest{}

			tmp := rs.Primary.ID
			request.CloudExadataStorageConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetCloudExadataStorageConnector(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DatabaseManagementCloudExadataStorageConnector") {
		resource.AddTestSweepers("DatabaseManagementCloudExadataStorageConnector", &resource.Sweeper{
			Name:         "DatabaseManagementCloudExadataStorageConnector",
			Dependencies: acctest.DependencyGraph["cloudExadataStorageConnector"],
			F:            sweepDatabaseManagementCloudExadataStorageConnectorResource,
		})
	}
}

func sweepDatabaseManagementCloudExadataStorageConnectorResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	cloudExadataStorageConnectorIds, err := getDatabaseManagementCloudExadataStorageConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudExadataStorageConnectorId := range cloudExadataStorageConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[cloudExadataStorageConnectorId]; !ok {
			deleteCloudExadataStorageConnectorRequest := oci_database_management.DeleteCloudExadataStorageConnectorRequest{}

			deleteCloudExadataStorageConnectorRequest.CloudExadataStorageConnectorId = &cloudExadataStorageConnectorId

			deleteCloudExadataStorageConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteCloudExadataStorageConnector(context.Background(), deleteCloudExadataStorageConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudExadataStorageConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudExadataStorageConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cloudExadataStorageConnectorId, DatabaseManagementCloudExadataStorageConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementCloudExadataStorageConnectorSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementCloudExadataStorageConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudExadataStorageConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listCloudExadataStorageConnectorsRequest := oci_database_management.ListCloudExadataStorageConnectorsRequest{}
	listCloudExadataStorageConnectorsRequest.CompartmentId = &compartmentId

	cloudExadataInfrastructureIds, error := getDatabaseManagementCloudExadataInfrastructureIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting cloudExadataInfrastructureId required for CloudExadataStorageConnector resource requests \n")
	}
	for _, cloudExadataInfrastructureId := range cloudExadataInfrastructureIds {
		listCloudExadataStorageConnectorsRequest.CloudExadataInfrastructureId = &cloudExadataInfrastructureId

		listCloudExadataStorageConnectorsResponse, err := dbManagementClient.ListCloudExadataStorageConnectors(context.Background(), listCloudExadataStorageConnectorsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting CloudExadataStorageConnector list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, cloudExadataStorageConnector := range listCloudExadataStorageConnectorsResponse.Items {
			id := *cloudExadataStorageConnector.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudExadataStorageConnectorId", id)
		}

	}
	return resourceIds, nil
}

func DatabaseManagementCloudExadataStorageConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudExadataStorageConnectorResponse, ok := response.Response.(oci_database_management.GetCloudExadataStorageConnectorResponse); ok {
		return cloudExadataStorageConnectorResponse.GetLifecycleState() != oci_database_management.DbmResourceLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementCloudExadataStorageConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetCloudExadataStorageConnector(context.Background(), oci_database_management.GetCloudExadataStorageConnectorRequest{
		CloudExadataStorageConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
