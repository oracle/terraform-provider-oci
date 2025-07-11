// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
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
	DatabaseManagementCloudDbSystemConnectorRequiredOnlyResource = DatabaseManagementCloudDbSystemConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_connector", "test_cloud_db_system_connector", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemConnectorRepresentation)

	DatabaseManagementCloudDbSystemConnectorResourceConfig = DatabaseManagementCloudDbSystemConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_connector", "test_cloud_db_system_connector", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemConnectorRepresentation)

	DatabaseManagementCloudDbSystemConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_db_system_connector.test_cloud_db_system_connector.id}`},
	}

	DatabaseManagementCloudDbSystemConnectorDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.cloud_db_system_id}`},
		"compartment_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `EXAMPLE-displayName-Value`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementCloudDbSystemConnectorDataSourceFilterRepresentation}}
	DatabaseManagementCloudDbSystemConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_db_system_connector.test_cloud_db_system_connector.id}`}},
	}

	DatabaseManagementCloudDbSystemConnectorRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_db_system_id}`},
		"connector_type":     acctest.Representation{RepType: acctest.Required, Create: `MACS`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `EXAMPLE-displayName-Value`},
		"agent_id":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_cloud_db_system.test_cloud_db_system.discovery_agent_id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementDefinedTagsChangesRepresentation},
	}

	DatabaseManagementCloudDbSystemSingleDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cloud_db_system_id}`},
	}

	DatabaseManagementCloudDbSystemConnectorResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_system", "test_cloud_db_system", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemSingleDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudDbSystemConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudDbSystemConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("dbaas_dbsystem_id")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"cloud_db_system_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_cloud_db_system_connector.test_cloud_db_system_connector"
	datasourceName := "data.oci_database_management_cloud_db_system_connectors.test_cloud_db_system_connectors"
	singularDatasourceName := "data.oci_database_management_cloud_db_system_connector.test_cloud_db_system_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dbSystemIdVariableStr+DatabaseManagementCloudDbSystemConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_connector", "test_cloud_db_system_connector", acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemConnectorRepresentation), "databasemanagement", "cloudDbSystemConnector", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementCloudDbSystemConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbSystemConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_connector", "test_cloud_db_system_connector", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "MACS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbSystemConnectorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbSystemConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_connector", "test_cloud_db_system_connector", acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EXAMPLE-displayName-Value"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] Test step with optionals")

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
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbSystemConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_connector", "test_cloud_db_system_connector", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "MACS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "EXAMPLE-displayName-Value"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					log.Printf("[DEBUG] Test step with update")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_system_connectors", "test_cloud_db_system_connectors", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbSystemConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_connector", "test_cloud_db_system_connector", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cloud_db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "EXAMPLE-displayName-Value"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_db_system_connector_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_db_system_connector_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_system_connector", "test_cloud_db_system_connector", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbSystemIdVariableStr + DatabaseManagementCloudDbSystemConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cloud_db_system_connector_id"),

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
			Config:                  config + DatabaseManagementCloudDbSystemConnectorRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseManagementCloudDbSystemConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_cloud_db_system_connector" {
			noResourceFound = false
			request := oci_database_management.GetCloudDbSystemConnectorRequest{}

			tmp := rs.Primary.ID
			request.CloudDbSystemConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetCloudDbSystemConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.CloudDbSystemConnectorLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementCloudDbSystemConnector") {
		resource.AddTestSweepers("DatabaseManagementCloudDbSystemConnector", &resource.Sweeper{
			Name:         "DatabaseManagementCloudDbSystemConnector",
			Dependencies: acctest.DependencyGraph["cloudDbSystemConnector"],
			F:            sweepDatabaseManagementCloudDbSystemConnectorResource,
		})
	}
}

func sweepDatabaseManagementCloudDbSystemConnectorResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	cloudDbSystemConnectorIds, err := getDatabaseManagementCloudDbSystemConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudDbSystemConnectorId := range cloudDbSystemConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[cloudDbSystemConnectorId]; !ok {
			deleteCloudDbSystemConnectorRequest := oci_database_management.DeleteCloudDbSystemConnectorRequest{}

			deleteCloudDbSystemConnectorRequest.CloudDbSystemConnectorId = &cloudDbSystemConnectorId

			deleteCloudDbSystemConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteCloudDbSystemConnector(context.Background(), deleteCloudDbSystemConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudDbSystemConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudDbSystemConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cloudDbSystemConnectorId, DatabaseManagementCloudDbSystemConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementCloudDbSystemConnectorSweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementCloudDbSystemConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudDbSystemConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listCloudDbSystemConnectorsRequest := oci_database_management.ListCloudDbSystemConnectorsRequest{}
	listCloudDbSystemConnectorsRequest.CompartmentId = &compartmentId
	listCloudDbSystemConnectorsResponse, err := dbManagementClient.ListCloudDbSystemConnectors(context.Background(), listCloudDbSystemConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudDbSystemConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudDbSystemConnector := range listCloudDbSystemConnectorsResponse.Items {
		id := *cloudDbSystemConnector.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudDbSystemConnectorId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementCloudDbSystemConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudDbSystemConnectorResponse, ok := response.Response.(oci_database_management.GetCloudDbSystemConnectorResponse); ok {
		return cloudDbSystemConnectorResponse.GetLifecycleState() != oci_database_management.CloudDbSystemConnectorLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementCloudDbSystemConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetCloudDbSystemConnector(context.Background(), oci_database_management.GetCloudDbSystemConnectorRequest{
		CloudDbSystemConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
