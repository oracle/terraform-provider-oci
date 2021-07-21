// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v45/common"
	oci_database "github.com/oracle/oci-go-sdk/v45/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ExternalDatabaseConnectorRequiredOnlyResource = ExternalDatabaseConnectorResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalDatabaseConnectorRepresentation)

	ExternalDatabaseConnectorResourceConfig = ExternalDatabaseConnectorResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Optional, Update, externalDatabaseConnectorRepresentation)

	externalDatabaseConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"external_database_connector_id": Representation{repType: Required, create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
	}

	externalDatabaseConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"external_database_id": Representation{repType: Required, create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"display_name":         Representation{repType: Optional, create: `myTestConn`, update: `displayName2`},
		"state":                Representation{repType: Optional, create: `AVAILABLE`},
		"filter":               RepresentationGroup{Required, externalDatabaseConnectorDataSourceFilterRepresentation}}
	externalDatabaseConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_external_database_connector.test_external_database_connector.id}`}},
	}

	externalDatabaseConnectorRepresentation = map[string]interface{}{
		"connection_credentials": RepresentationGroup{Required, externalDatabaseConnectorConnectionCredentialsRepresentation},
		"connection_string":      RepresentationGroup{Required, externalDatabaseConnectorConnectionStringRepresentation},
		"connector_agent_id":     Representation{repType: Required, create: `ocid1.managementagent.oc1.phx.amaaaaaajobtc3iaes4ijczgekzqigoji25xocsny7yundummydummydummy`},
		"display_name":           Representation{repType: Required, create: `myTestConn`, update: `displayName2`},
		"external_database_id":   Representation{repType: Required, create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"connector_type":         Representation{repType: Optional, create: `MACS`},
		"defined_tags":           Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":          Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	externalDatabaseConnectorConnectionCredentialsRepresentation = map[string]interface{}{
		"credential_name": Representation{repType: Required, create: `credential.name`},
		"credential_type": Representation{repType: Optional, create: `DETAILS`},
		"password":        Representation{repType: Required, create: `BEstrO0ng_#11`, update: `BEstrO0ng_#12`},
		"role":            Representation{repType: Required, create: `SYSDBA`, update: `NORMAL`},
		"username":        Representation{repType: Required, create: `testUser`, update: `username2`},
	}
	externalDatabaseConnectorConnectionStringRepresentation = map[string]interface{}{
		"hostname": Representation{repType: Required, create: `myHost.test`, update: `hostname2`},
		"port":     Representation{repType: Required, create: `10`, update: `11`},
		"protocol": Representation{repType: Required, create: `TCP`},
		"service":  Representation{repType: Required, create: `testService`, update: `service2`},
	}

	ExternalDatabaseConnectorResourceDependencies = generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Required, Create, externalNonContainerDatabaseRepresentation) +
		DefinedTagsDependencies
)

func TestDatabaseExternalDatabaseConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalDatabaseConnectorResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_database_connector.test_external_database_connector"
	datasourceName := "data.oci_database_external_database_connectors.test_external_database_connectors"
	singularDatasourceName := "data.oci_database_external_database_connector.test_external_database_connector"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ExternalDatabaseConnectorResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Optional, Create, externalDatabaseConnectorRepresentation), "database", "externalDatabaseConnector", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseExternalDatabaseConnectorDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalDatabaseConnectorRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies,
			},

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Optional, Create, externalDatabaseConnectorRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "myTestConn"),
					resource.TestCheckResourceAttrSet(resourceName, "external_database_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Optional, Update, externalDatabaseConnectorRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "external_database_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_external_database_connectors", "test_external_database_connectors", Optional, Update, externalDatabaseConnectorDataSourceRepresentation) +
					compartmentIdVariableStr + ExternalDatabaseConnectorResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Optional, Update, externalDatabaseConnectorRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					resource.TestCheckResourceAttr(datasourceName, "external_database_connectors.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalDatabaseConnectorSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ExternalDatabaseConnectorResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		},
	})
}

func testAccCheckDatabaseExternalDatabaseConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_external_database_connector" {
			noResourceFound = false
			request := oci_database.GetExternalDatabaseConnectorRequest{}

			tmp := rs.Primary.ID
			request.ExternalDatabaseConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseExternalDatabaseConnector") {
		resource.AddTestSweepers("DatabaseExternalDatabaseConnector", &resource.Sweeper{
			Name:         "DatabaseExternalDatabaseConnector",
			Dependencies: DependencyGraph["externalDatabaseConnector"],
			F:            sweepDatabaseExternalDatabaseConnectorResource,
		})
	}
}

func sweepDatabaseExternalDatabaseConnectorResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	externalDatabaseConnectorIds, err := getExternalDatabaseConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, externalDatabaseConnectorId := range externalDatabaseConnectorIds {
		if ok := SweeperDefaultResourceId[externalDatabaseConnectorId]; !ok {
			deleteExternalDatabaseConnectorRequest := oci_database.DeleteExternalDatabaseConnectorRequest{}

			deleteExternalDatabaseConnectorRequest.ExternalDatabaseConnectorId = &externalDatabaseConnectorId

			deleteExternalDatabaseConnectorRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExternalDatabaseConnector(context.Background(), deleteExternalDatabaseConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalDatabaseConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalDatabaseConnectorId, error)
				continue
			}
			waitTillCondition(testAccProvider, &externalDatabaseConnectorId, externalDatabaseConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				externalDatabaseConnectorSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExternalDatabaseConnectorIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ExternalDatabaseConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

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
			addResourceIdToSweeperResourceIdMap(compartmentId, "ExternalDatabaseConnectorId", id)
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

func externalDatabaseConnectorSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetExternalDatabaseConnector(context.Background(), oci_database.GetExternalDatabaseConnectorRequest{
		ExternalDatabaseConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
