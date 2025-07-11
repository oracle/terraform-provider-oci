// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
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
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ignoreDbManagementCloudDbSystemDiscoveryDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatabaseManagementCloudDbSystemDiscoveryRequiredOnlyResource = DatabaseManagementCloudDbSystemDiscoveryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryRepresentation)

	DatabaseManagementCloudDbSystemDiscoveryResourceConfig = DatabaseManagementCloudDbSystemDiscoveryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemDiscoveryRepresentation)

	DatabaseManagementCloudDbSystemDiscoverySingularDataSourceRepresentation = map[string]interface{}{
		"cloud_db_system_discovery_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_db_system_discovery.test_cloud_db_system_discovery.id}`},
	}

	DatabaseManagementCloudDbSystemDiscoveryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.disc_compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `${var.disc_dbaas_dbsystem_name}`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementCloudDbSystemDiscoveryDataSourceFilterRepresentation}}
	DatabaseManagementCloudDbSystemDiscoveryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_cloud_db_system_discovery.test_cloud_db_system_discovery.id}`}},
	}

	DatabaseManagementCloudDbSystemDiscoveryRepresentation = map[string]interface{}{
		"agent_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.disc_compartment_id}`},
		"dbaas_parent_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${var.disc_dbaas_dbsystem_id}`},
		"deployment_type":                acctest.Representation{RepType: acctest.Required, Create: `VM`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `${var.disc_dbaas_dbsystem_name}`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDbManagementCloudDbSystemDiscoveryDefinedTagsChangesRepresentation},
	}

	DatabaseManagementCloudDbSystemDiscoveryResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudDbSystemDiscoveryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudDbSystemDiscoveryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("disc_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"disc_compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbaasDbsystemId := utils.GetEnvSettingWithBlankDefault("disc_dbaas_dbsystem_id")
	dbaasDbsystemIdVariableStr := fmt.Sprintf("variable \"disc_dbaas_dbsystem_id\" { default = \"%s\" }\n", dbaasDbsystemId)

	dbaasDbsystemName := utils.GetEnvSettingWithBlankDefault("disc_dbaas_dbsystem_name")
	dbaasDbsystemNameVariableStr := fmt.Sprintf("variable \"disc_dbaas_dbsystem_name\" { default = \"%s\" }\n", dbaasDbsystemName)

	cloudAgentId := utils.GetEnvSettingWithBlankDefault("disc_cloud_agent_id")
	cloudAgentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", cloudAgentId)

	resourceName := "oci_database_management_cloud_db_system_discovery.test_cloud_db_system_discovery"
	datasourceName := "data.oci_database_management_cloud_db_system_discoveries.test_cloud_db_system_discoveries"
	singularDatasourceName := "data.oci_database_management_cloud_db_system_discovery.test_cloud_db_system_discovery"

	var variableStr = compartmentIdVariableStr + dbaasDbsystemIdVariableStr + cloudAgentIdVariableStr + dbaasDbsystemNameVariableStr

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+DatabaseManagementCloudDbSystemDiscoveryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery", acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryRepresentation), "databasemanagement", "cloudDbSystemDiscovery", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementCloudDbSystemDiscoveryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + DatabaseManagementCloudDbSystemDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dbaas_parent_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "VM"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variableStr + DatabaseManagementCloudDbSystemDiscoveryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variableStr + DatabaseManagementCloudDbSystemDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery", acctest.Optional, acctest.Create, DatabaseManagementCloudDbSystemDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dbaas_parent_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "VM"),
				resource.TestCheckResourceAttr(resourceName, "display_name", dbaasDbsystemName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + variableStr + DatabaseManagementCloudDbSystemDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dbaas_parent_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "deployment_type", "VM"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_system_discoveries", "test_cloud_db_system_discoveries", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemDiscoveryDataSourceRepresentation) +
				variableStr + DatabaseManagementCloudDbSystemDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery", acctest.Optional, acctest.Update, DatabaseManagementCloudDbSystemDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "cloud_db_system_discovery_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cloud_db_system_discovery_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_cloud_db_system_discovery", "test_cloud_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementCloudDbSystemDiscoverySingularDataSourceRepresentation) +
				variableStr + DatabaseManagementCloudDbSystemDiscoveryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_type", "VM"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovered_components.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grid_home"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementCloudDbSystemDiscoveryRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"patch_operations",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementCloudDbSystemDiscoveryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_cloud_db_system_discovery" {
			noResourceFound = false
			request := oci_database_management.GetCloudDbSystemDiscoveryRequest{}

			tmp := rs.Primary.ID
			request.CloudDbSystemDiscoveryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetCloudDbSystemDiscovery(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.CloudDbSystemDiscoveryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementCloudDbSystemDiscovery") {
		resource.AddTestSweepers("DatabaseManagementCloudDbSystemDiscovery", &resource.Sweeper{
			Name:         "DatabaseManagementCloudDbSystemDiscovery",
			Dependencies: acctest.DependencyGraph["cloudDbSystemDiscovery"],
			F:            sweepDatabaseManagementCloudDbSystemDiscoveryResource,
		})
	}
}

func sweepDatabaseManagementCloudDbSystemDiscoveryResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	cloudDbSystemDiscoveryIds, err := getDatabaseManagementCloudDbSystemDiscoveryIds(compartment)
	if err != nil {
		return err
	}
	for _, cloudDbSystemDiscoveryId := range cloudDbSystemDiscoveryIds {
		if ok := acctest.SweeperDefaultResourceId[cloudDbSystemDiscoveryId]; !ok {
			deleteCloudDbSystemDiscoveryRequest := oci_database_management.DeleteCloudDbSystemDiscoveryRequest{}

			deleteCloudDbSystemDiscoveryRequest.CloudDbSystemDiscoveryId = &cloudDbSystemDiscoveryId

			deleteCloudDbSystemDiscoveryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteCloudDbSystemDiscovery(context.Background(), deleteCloudDbSystemDiscoveryRequest)
			if error != nil {
				fmt.Printf("Error deleting CloudDbSystemDiscovery %s %s, It is possible that the resource is already deleted. Please verify manually \n", cloudDbSystemDiscoveryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cloudDbSystemDiscoveryId, DatabaseManagementCloudDbSystemDiscoverySweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementCloudDbSystemDiscoverySweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementCloudDbSystemDiscoveryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CloudDbSystemDiscoveryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listCloudDbSystemDiscoveriesRequest := oci_database_management.ListCloudDbSystemDiscoveriesRequest{}
	listCloudDbSystemDiscoveriesRequest.CompartmentId = &compartmentId
	listCloudDbSystemDiscoveriesResponse, err := dbManagementClient.ListCloudDbSystemDiscoveries(context.Background(), listCloudDbSystemDiscoveriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CloudDbSystemDiscovery list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cloudDbSystemDiscovery := range listCloudDbSystemDiscoveriesResponse.Items {
		id := *cloudDbSystemDiscovery.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CloudDbSystemDiscoveryId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementCloudDbSystemDiscoverySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cloudDbSystemDiscoveryResponse, ok := response.Response.(oci_database_management.GetCloudDbSystemDiscoveryResponse); ok {
		return cloudDbSystemDiscoveryResponse.LifecycleState != oci_database_management.CloudDbSystemDiscoveryLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementCloudDbSystemDiscoverySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetCloudDbSystemDiscovery(context.Background(), oci_database_management.GetCloudDbSystemDiscoveryRequest{
		CloudDbSystemDiscoveryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
