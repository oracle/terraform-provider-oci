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
	ignoreDbManagementDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatabaseManagementExternalDbSystemConnectorRequiredOnlyResource = DatabaseManagementExternalDbSystemConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_connector", "test_external_db_system_connector", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemConnectorRepresentation)

	DatabaseManagementExternalDbSystemConnectorResourceConfig = DatabaseManagementExternalDbSystemConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_connector", "test_external_db_system_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemConnectorRepresentation)

	DatabaseManagementDatabaseManagementExternalDbSystemConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"external_db_system_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_db_system_connector.test_external_db_system_connector.id}`},
	}

	DatabaseManagementDatabaseManagementExternalDbSystemConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `EXAMPLE-displayName-Value`},
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemConnectorDataSourceFilterRepresentation}}
	DatabaseManagementExternalDbSystemConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_db_system_connector.test_external_db_system_connector.id}`}},
	}

	DatabaseManagementExternalDbSystemConnectorRepresentation = map[string]interface{}{
		"connector_type":        acctest.Representation{RepType: acctest.Required, Create: `MACS`},
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `EXAMPLE-displayName-Value`},
		"agent_id":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_external_db_system.test_external_db_system.discovery_agent_id}`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}
	DatabaseManagementExternalDbSystemSingularDataSourceRepresentation = map[string]interface{}{
		"external_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.external_dbsystem_id}`},
	}

	DatabaseManagementExternalDbSystemConnectorResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_system", "test_external_db_system", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemSingularDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDbSystemConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDbSystemConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbmgmt_external_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"external_dbsystem_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_db_system_connector.test_external_db_system_connector"
	datasourceName := "data.oci_database_management_external_db_system_connectors.test_external_db_system_connectors"
	singularDatasourceName := "data.oci_database_management_external_db_system_connector.test_external_db_system_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+DatabaseManagementExternalDbSystemConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_connector", "test_external_db_system_connector", acctest.Optional, acctest.Create, DatabaseManagementExternalDbSystemConnectorRepresentation), "databasemanagement", "externalDbSystemConnector", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementExternalDbSystemConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbSystemConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_connector", "test_external_db_system_connector", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbSystemConnectorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbSystemConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_connector", "test_external_db_system_connector", acctest.Optional, acctest.Create, DatabaseManagementExternalDbSystemConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EXAMPLE-displayName-Value"),
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbSystemConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_connector", "test_external_db_system_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EXAMPLE-displayName-Value"),
				resource.TestCheckResourceAttrSet(resourceName, "external_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_system_connectors", "test_external_db_system_connectors", acctest.Optional, acctest.Update, DatabaseManagementDatabaseManagementExternalDbSystemConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbSystemConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_connector", "test_external_db_system_connector", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "EXAMPLE-displayName-Value"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "external_db_system_connector_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_db_system_connector_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_system_connector", "test_external_db_system_connector", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbSystemConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementExternalDbSystemConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_db_system_connector_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "EXAMPLE-displayName-Value"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseManagementExternalDbSystemConnectorRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseManagementExternalDbSystemConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_external_db_system_connector" {
			noResourceFound = false
			request := oci_database_management.GetExternalDbSystemConnectorRequest{}

			tmp := rs.Primary.ID
			request.ExternalDbSystemConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetExternalDbSystemConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.ExternalDbSystemConnectorLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementExternalDbSystemConnector") {
		resource.AddTestSweepers("DatabaseManagementExternalDbSystemConnector", &resource.Sweeper{
			Name:         "DatabaseManagementExternalDbSystemConnector",
			Dependencies: acctest.DependencyGraph["externalDbSystemConnector"],
			F:            sweepDatabaseManagementExternalDbSystemConnectorResource,
		})
	}
}

func sweepDatabaseManagementExternalDbSystemConnectorResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	externalDbSystemConnectorIds, err := getDatabaseManagementExternalDbSystemConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, externalDbSystemConnectorId := range externalDbSystemConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[externalDbSystemConnectorId]; !ok {
			deleteExternalDbSystemConnectorRequest := oci_database_management.DeleteExternalDbSystemConnectorRequest{}

			deleteExternalDbSystemConnectorRequest.ExternalDbSystemConnectorId = &externalDbSystemConnectorId

			deleteExternalDbSystemConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteExternalDbSystemConnector(context.Background(), deleteExternalDbSystemConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalDbSystemConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalDbSystemConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalDbSystemConnectorId, DatabaseManagementExternalDbSystemConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementExternalDbSystemConnectorSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementExternalDbSystemConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalDbSystemConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listExternalDbSystemConnectorsRequest := oci_database_management.ListExternalDbSystemConnectorsRequest{}
	listExternalDbSystemConnectorsRequest.CompartmentId = &compartmentId
	listExternalDbSystemConnectorsResponse, err := dbManagementClient.ListExternalDbSystemConnectors(context.Background(), listExternalDbSystemConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalDbSystemConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalDbSystemConnector := range listExternalDbSystemConnectorsResponse.Items {
		id := *externalDbSystemConnector.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalDbSystemConnectorId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementExternalDbSystemConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalDbSystemConnectorResponse, ok := response.Response.(oci_database_management.GetExternalDbSystemConnectorResponse); ok {
		return externalDbSystemConnectorResponse.GetLifecycleState() != oci_database_management.ExternalDbSystemConnectorLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementExternalDbSystemConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetExternalDbSystemConnector(context.Background(), oci_database_management.GetExternalDbSystemConnectorRequest{
		ExternalDbSystemConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
