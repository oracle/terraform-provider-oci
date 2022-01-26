// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ExternalDatabaseConnectorRequiredOnlyResource = ExternalDatabaseConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorRepresentation)

	ExternalDatabaseConnectorResourceConfig = ExternalDatabaseConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Update, externalDatabaseConnectorRepresentation)

	externalDatabaseConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"external_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
	}

	externalDatabaseConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `myTestConn`, Update: `displayName2`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":               acctest.RepresentationGroup{RepType: acctest.Required, Group: externalDatabaseConnectorDataSourceFilterRepresentation}}
	externalDatabaseConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_external_database_connector.test_external_database_connector.id}`}},
	}

	externalDatabaseConnectorRepresentation = map[string]interface{}{
		"connection_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: externalDatabaseConnectorConnectionCredentialsRepresentation},
		"connection_string":      acctest.RepresentationGroup{RepType: acctest.Required, Group: externalDatabaseConnectorConnectionStringRepresentation},
		"connector_agent_id":     acctest.Representation{RepType: acctest.Required, Create: `ocid1.managementagent.oc1.phx.amaaaaaajobtc3iaes4ijczgekzqigoji25xocsny7yundummydummydummy`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `myTestConn`, Update: `displayName2`},
		"external_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"connector_type":         acctest.Representation{RepType: acctest.Optional, Create: `MACS`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	externalDatabaseConnectorConnectionCredentialsRepresentation = map[string]interface{}{
		"credential_name": acctest.Representation{RepType: acctest.Required, Create: `credential.name`},
		"credential_type": acctest.Representation{RepType: acctest.Optional, Create: `DETAILS`},
		"password":        acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"role":            acctest.Representation{RepType: acctest.Required, Create: `SYSDBA`, Update: `NORMAL`},
		"username":        acctest.Representation{RepType: acctest.Required, Create: `testUser`, Update: `username2`},
	}
	externalDatabaseConnectorConnectionStringRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `myHost.test`, Update: `hostname2`},
		"port":     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"protocol": acctest.Representation{RepType: acctest.Required, Create: `TCP`},
		"service":  acctest.Representation{RepType: acctest.Required, Create: `testService`, Update: `service2`},
	}

	ExternalDatabaseConnectorResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Required, acctest.Create, externalNonContainerDatabaseRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseExternalDatabaseConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalDatabaseConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_database_connector.test_external_database_connector"
	datasourceName := "data.oci_database_external_database_connectors.test_external_database_connectors"
	singularDatasourceName := "data.oci_database_external_database_connector.test_external_database_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExternalDatabaseConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Create, externalDatabaseConnectorRepresentation), "database", "externalDatabaseConnector", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExternalDatabaseConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.role", "SYSDBA"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.hostname", "myHost.test"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.protocol", "TCP"),
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
			Config: config + compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Create, externalDatabaseConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_type", "DETAILS"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.role", "SYSDBA"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.hostname", "myHost.test"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.protocol", "TCP"),
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
			Config: config + compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Update, externalDatabaseConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.credential_type", "DETAILS"),
				resource.TestCheckResourceAttr(resourceName, "connection_credentials.0.role", "NORMAL"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.port", "11"),
				resource.TestCheckResourceAttr(resourceName, "connection_string.0.protocol", "TCP"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_external_database_connectors", "test_external_database_connectors", acctest.Optional, acctest.Update, externalDatabaseConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Optional, acctest.Update, externalDatabaseConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_database_connectors.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_credentials.0.credential_type", "DETAILS"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_credentials.0.role", "NORMAL"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.0.port", "11"),
				resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.connection_string.0.protocol", "TCP"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ExternalDatabaseConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_database_connector_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_credentials.0.credential_name", "credential.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_credentials.0.credential_type", "DETAILS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_credentials.0.role", "NORMAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.0.port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string.0.protocol", "TCP"),
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
			Config: config + compartmentIdVariableStr + ExternalDatabaseConnectorResourceConfig,
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

func testAccCheckDatabaseExternalDatabaseConnectorDestroy(s *terraform.State) error {
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
	if !acctest.InSweeperExcludeList("DatabaseExternalDatabaseConnector") {
		resource.AddTestSweepers("DatabaseExternalDatabaseConnector", &resource.Sweeper{
			Name:         "DatabaseExternalDatabaseConnector",
			Dependencies: acctest.DependencyGraph["externalDatabaseConnector"],
			F:            sweepDatabaseExternalDatabaseConnectorResource,
		})
	}
}

func sweepDatabaseExternalDatabaseConnectorResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	externalDatabaseConnectorIds, err := getExternalDatabaseConnectorIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalDatabaseConnectorId, externalDatabaseConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				externalDatabaseConnectorSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExternalDatabaseConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalDatabaseConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listExternalDatabaseConnectorsRequest := oci_database.ListExternalDatabaseConnectorsRequest{}
	listExternalDatabaseConnectorsRequest.CompartmentId = &compartmentId

	externalDatabaseIds, error := getDatabaseIds(compartment)
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

func externalDatabaseConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalDatabaseConnectorResponse, ok := response.Response.(oci_database.GetExternalDatabaseConnectorResponse); ok {
		return externalDatabaseConnectorResponse.GetLifecycleState() != oci_database.ExternalDatabaseConnectorLifecycleStateTerminated
	}
	return false
}

func externalDatabaseConnectorSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetExternalDatabaseConnector(context.Background(), oci_database.GetExternalDatabaseConnectorRequest{
		ExternalDatabaseConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
