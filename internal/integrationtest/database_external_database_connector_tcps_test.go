// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ExternalDatabaseTcpsConnectorRequiredOnlyResource = ExternalDatabaseTcpsConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorTcpsRepresentation)

	ExternalDatabaseTcpsConnectorResourceConfig = ExternalDatabaseTcpsConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Update, externalDatabaseConnectorTcpsRepresentation)

	externalDatabaseConnectorTcpsSingularDataSourceRepresentation = map[string]interface{}{
		"external_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
	}

	externalDatabaseConnectorTcpsDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `myTestConn`, Update: `displayName2`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: externalDatabaseConnectorTcpsDataSourceFilterRepresentation}}
	externalDatabaseConnectorTcpsDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_external_database_connector.test_external_database_connector.id}`}},
	}

	externalDatabaseConnectorTcpsRepresentation = map[string]interface{}{
		"connection_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: externalDatabaseConnectorTcpsConnectionCredentialsRepresentation},
		"connection_string":      acctest.RepresentationGroup{RepType: acctest.Required, Group: externalDatabaseConnectorTcpsConnectionStringRepresentation},
		"connector_agent_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `myTestConn`, Update: `displayName2`},
		"external_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"connector_type":         acctest.Representation{RepType: acctest.Optional, Create: `MACS`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	externalDatabaseConnectorTcpsConnectionCredentialsRepresentation = map[string]interface{}{
		"credential_name": acctest.Representation{RepType: acctest.Required, Create: `credential.name`},
		"credential_type": acctest.Representation{RepType: acctest.Required, Create: `SSL_DETAILS`},
		"password":        acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"role":            acctest.Representation{RepType: acctest.Required, Create: `SYSDBA`, Update: `NORMAL`},
		"username":        acctest.Representation{RepType: acctest.Required, Create: `testUser`, Update: `username2`},
		"ssl_secret_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.sslSecretId}`},
	}

	externalDatabaseConnectorTcpsConnectionStringRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `myHost.test`, Update: `hostname2`},
		"port":     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"protocol": acctest.Representation{RepType: acctest.Required, Create: `TCPS`},
		"service":  acctest.Representation{RepType: acctest.Required, Create: `testService`, Update: `service2`},
	}

	ExternalDatabaseTcpsConnectorResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Required, acctest.Create, DatabaseExternalNonContainerDatabaseRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseExternalDatabaseTcpsConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalDatabaseTcpsConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	agentId := utils.GetEnvSettingWithBlankDefault("connector_agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)

	sslSecretId := utils.GetEnvSettingWithBlankDefault("secret_id")
	sslSecretIdVariableStr := fmt.Sprintf("variable \"sslSecretId\" { default = \"%s\" }\n", sslSecretId)

	resourceName := "oci_database_external_database_connector.test_external_database_connector"
	datasourceName := "data.oci_database_external_database_connectors.test_external_database_connectors"
	singularDatasourceName := "data.oci_database_external_database_connector.test_external_database_connector"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+sslSecretIdVariableStr+ExternalDatabaseTcpsConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Create, externalDatabaseConnectorTcpsRepresentation), "database", "externalDatabaseConnector", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExternalDatabaseTcpsConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + sslSecretIdVariableStr + ExternalDatabaseTcpsConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorTcpsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.role", "SYSDBA"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_credentials.0.ssl_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.hostname", "myHost.test"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.protocol", "TCPS"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.service", "testService"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestConn"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + sslSecretIdVariableStr + ExternalDatabaseTcpsConnectorResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + sslSecretIdVariableStr + ExternalDatabaseTcpsConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Create, externalDatabaseConnectorTcpsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_type", "SSL_DETAILS"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.role", "SYSDBA"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_credentials.0.ssl_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.hostname", "myHost.test"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.protocol", "TCPS"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.service", "testService"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestConn"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + agentIdVariableStr + sslSecretIdVariableStr + ExternalDatabaseTcpsConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Update, externalDatabaseConnectorTcpsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_type", "SSL_DETAILS"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.role", "NORMAL"),
				resource.TestCheckResourceAttrSet(resourceName, "connection_credentials.0.ssl_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.port", "11"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.protocol", "TCPS"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.service", "service2"),
				resource.TestCheckResourceAttrSet(resourceName, "connector_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_external_database_connectors", "test_external_database_connectors", acctest.Optional, acctest.Update, externalDatabaseConnectorTcpsDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + sslSecretIdVariableStr + ExternalDatabaseTcpsConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Update, externalDatabaseConnectorTcpsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_connectors.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_credentials.0.credential_type", "SSL_DETAILS"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_credentials.0.role", "NORMAL"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_connectors.0.connection_credentials.0.ssl_secret_id"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.0.port", "11"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.0.protocol", "TCPS"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.0.service", "service2"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_connectors.0.connector_agent_id"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connector_type", "MACS"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_connectors.0.external_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_connectors.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_connectors.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_connectors.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorTcpsSingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + sslSecretIdVariableStr + ExternalDatabaseTcpsConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_database_connector_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_credentials.0.credential_type", "SSL_DETAILS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_credentials.0.role", "NORMAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.0.port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.0.protocol", "TCPS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.0.service", "service2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + sslSecretIdVariableStr + ExternalDatabaseTcpsConnectorResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"connection_credentials.0.username",
				"connection_credentials.0.password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseExternalDatabaseTcpsConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_external_database_connector" {
			noResourceFound = false
			request := oci_database.GetExternalDatabaseConnectorRequest{}

			tmp := rs.Primary.ID
			request.ExternalDatabaseConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetExternalDatabaseConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExternalDatabaseConnectorLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseExternalDatabaseTcpsConnector") {
		resource.AddTestSweepers("DatabaseExternalDatabaseTcpsConnector", &resource.Sweeper{
			Name:         "DatabaseExternalDatabaseTcpsConnector",
			Dependencies: acctest.DependencyGraph["externalDatabaseConnector"],
			F:            sweepDatabaseExternalDatabaseTcpsConnectorResource,
		})
	}
}

func sweepDatabaseExternalDatabaseTcpsConnectorResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	externalDatabaseConnectorIds, err := getExternalDatabaseTcpsConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, externalDatabaseConnectorId := range externalDatabaseConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[externalDatabaseConnectorId]; !ok {
			deleteExternalDatabaseConnectorRequest := oci_database.DeleteExternalDatabaseConnectorRequest{}

			deleteExternalDatabaseConnectorRequest.ExternalDatabaseConnectorId = &externalDatabaseConnectorId

			deleteExternalDatabaseConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExternalDatabaseConnector(context.Background(), deleteExternalDatabaseConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalDatabaseConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalDatabaseConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalDatabaseConnectorId, externalDatabaseTcpsConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				externalDatabaseTcpsConnectorSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExternalDatabaseTcpsConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalDatabaseConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listExternalDatabaseConnectorsRequest := oci_database.ListExternalDatabaseConnectorsRequest{}
	listExternalDatabaseConnectorsRequest.CompartmentId = &compartmentId

	externalDatabaseIds, error := getDatabaseDatabaseIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting externalDatabaseId required for ExternalDatabaseConnector resource requests \n")
	}
	for _, externalDatabaseId := range externalDatabaseIds {
		listExternalDatabaseConnectorsRequest.ExternalDatabaseId = &externalDatabaseId

		listExternalDatabaseConnectorsRequest.LifecycleState = oci_database.ExternalDatabaseConnectorLifecycleStateAvailable
		listExternalDatabaseConnectorsResponse, err := databaseClient.ListExternalDatabaseConnectors(context.Background(), listExternalDatabaseConnectorsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ExternalDatabaseConnector list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, externalDatabaseConnector := range listExternalDatabaseConnectorsResponse.Items {
			id := *externalDatabaseConnector.GetId()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalDatabaseConnectorId", id)
		}

	}
	return resourceIds, nil
}

func externalDatabaseTcpsConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalDatabaseConnectorResponse, ok := response.Response.(oci_database.GetExternalDatabaseConnectorResponse); ok {
		return externalDatabaseConnectorResponse.GetLifecycleState() != oci_database.ExternalDatabaseConnectorLifecycleStateTerminated
	}
	return false
}

func externalDatabaseTcpsConnectorSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetExternalDatabaseConnector(context.Background(), oci_database.GetExternalDatabaseConnectorRequest{
		ExternalDatabaseConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
