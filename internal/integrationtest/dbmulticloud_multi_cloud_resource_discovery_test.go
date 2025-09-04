// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	// "strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	// "github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudMultiCloudResourceDiscoveryRequiredOnlyResource = DbmulticloudMultiCloudResourceDiscoveryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Required, acctest.Create, DbmulticloudMultiCloudResourceDiscoveryRepresentation)

	DbmulticloudMultiCloudResourceDiscoveryResourceConfig = DbmulticloudMultiCloudResourceDiscoveryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Optional, acctest.Update, DbmulticloudMultiCloudResourceDiscoveryRepresentation)

	DbmulticloudMultiCloudResourceDiscoverySingularDataSourceRepresentation = map[string]interface{}{
		"multi_cloud_resource_discovery_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_multi_cloud_resource_discovery.test_multi_cloud_resource_discovery.id}`},
	}

	DbmulticloudMultiCloudResourceDiscoveryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                      acctest.Representation{RepType: acctest.Optional, Create: `Tersi_Discover_Test`, Update: `Tersi_Discover_Test`},
		"multi_cloud_resource_discovery_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dbmulticloud_multi_cloud_resource_discovery.test_multi_cloud_resource_discovery.id}`},
		"oracle_db_azure_connector_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
		"resource_type":                     acctest.Representation{RepType: acctest.Optional, Create: `VAULTS`},
		//"resources_filter":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"keyVault": "resourcesFilter"}, Update: map[string]string{"keyVault": "resourcesFilter2"}},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudMultiCloudResourceDiscoveryDataSourceFilterRepresentation}}

	DbmulticloudMultiCloudResourceDiscoveryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_multi_cloud_resource_discovery.test_multi_cloud_resource_discovery.id}`}},
	}

	DbmulticloudMultiCloudResourceDiscoveryRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `Tersi_Discover_Test`, Update: `Tersi_Discover_Test`},
		"oracle_db_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
		"resource_type":          acctest.Representation{RepType: acctest.Required, Create: `VAULTS`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//"resources_filter":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"keyVault": "resourcesFilter"}, Update: map[string]string{"keyVault": "resourcesFilter2"}},
	}

	DbmulticloudMultiCloudResourceDiscoveryResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureConnectorRepresentation)
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudMultiCloudResourceDiscoveryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudMultiCloudResourceDiscoveryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_multi_cloud_resource_discovery.test_multi_cloud_resource_discovery"
	datasourceName := "data.oci_dbmulticloud_multi_cloud_resource_discoveries.test_multi_cloud_resource_discoveries"
	singularDatasourceName := "data.oci_dbmulticloud_multi_cloud_resource_discovery.test_multi_cloud_resource_discovery"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudMultiCloudResourceDiscoveryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Optional, acctest.Create, DbmulticloudMultiCloudResourceDiscoveryRepresentation), "dbmulticloud", "multiCloudResourceDiscovery", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudMultiCloudResourceDiscoveryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudMultiCloudResourceDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Required, acctest.Create, DbmulticloudMultiCloudResourceDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Tersi_Discover_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "VAULTS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudMultiCloudResourceDiscoveryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbmulticloudMultiCloudResourceDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Optional, acctest.Create, DbmulticloudMultiCloudResourceDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Tersi_Discover_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "VAULTS"),
				resource.TestCheckResourceAttr(resourceName, "resources_filter.%", "0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					// 	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					// 		return errExport
					// 	}
					// }
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudMultiCloudResourceDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudMultiCloudResourceDiscoveryRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Tersi_Discover_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "VAULTS"),
				resource.TestCheckResourceAttr(resourceName, "resources_filter.%", "0"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudMultiCloudResourceDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Optional, acctest.Update, DbmulticloudMultiCloudResourceDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Tersi_Discover_Test"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "VAULTS"),
				resource.TestCheckResourceAttr(resourceName, "resources_filter.%", "0"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discoveries", "test_multi_cloud_resource_discoveries", acctest.Optional, acctest.Update, DbmulticloudMultiCloudResourceDiscoveryDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudMultiCloudResourceDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Optional, acctest.Update, DbmulticloudMultiCloudResourceDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Tersi_Discover_Test"),
				resource.TestCheckResourceAttrSet(datasourceName, "multi_cloud_resource_discovery_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_azure_connector_id"),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "VAULTS"),
				resource.TestCheckResourceAttr(datasourceName, "resources_filter.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

				resource.TestCheckResourceAttr(datasourceName, "multi_cloud_resource_discovery_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "multi_cloud_resource_discovery_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_multi_cloud_resource_discovery", "test_multi_cloud_resource_discovery", acctest.Required, acctest.Create, DbmulticloudMultiCloudResourceDiscoverySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudMultiCloudResourceDiscoveryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_connector_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Tersi_Discover_Test"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type", "VAULTS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources_filter.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DbmulticloudMultiCloudResourceDiscoveryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDbmulticloudMultiCloudResourceDiscoveryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MultiCloudResourceDiscoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_multi_cloud_resource_discovery" {
			noResourceFound = false
			request := oci_dbmulticloud.GetMultiCloudResourceDiscoveryRequest{}

			tmp := rs.Primary.ID
			request.MultiCloudResourceDiscoveryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			response, err := client.GetMultiCloudResourceDiscovery(context.Background(), request)

			// if err == nil {
			// 	return fmt.Errorf("resource still exists")
			// }

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateCanceled):       true,
					string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateFailed):         true,
					string(oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateNeedsAttention): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudMultiCloudResourceDiscovery") {
		resource.AddTestSweepers("DbmulticloudMultiCloudResourceDiscovery", &resource.Sweeper{
			Name:         "DbmulticloudMultiCloudResourceDiscovery",
			Dependencies: acctest.DependencyGraph["multiCloudResourceDiscovery"],
			F:            sweepDbmulticloudMultiCloudResourceDiscoveryResource,
		})
	}
}

func sweepDbmulticloudMultiCloudResourceDiscoveryResource(compartment string) error {
	multiCloudResourceDiscoveryClient := acctest.GetTestClients(&schema.ResourceData{}).MultiCloudResourceDiscoveryClient()
	multiCloudResourceDiscoveryIds, err := getDbmulticloudMultiCloudResourceDiscoveryIds(compartment)
	if err != nil {
		return err
	}
	for _, multiCloudResourceDiscoveryId := range multiCloudResourceDiscoveryIds {
		if ok := acctest.SweeperDefaultResourceId[multiCloudResourceDiscoveryId]; !ok {
			deleteMultiCloudResourceDiscoveryRequest := oci_dbmulticloud.DeleteMultiCloudResourceDiscoveryRequest{}

			deleteMultiCloudResourceDiscoveryRequest.MultiCloudResourceDiscoveryId = &multiCloudResourceDiscoveryId

			deleteMultiCloudResourceDiscoveryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := multiCloudResourceDiscoveryClient.DeleteMultiCloudResourceDiscovery(context.Background(), deleteMultiCloudResourceDiscoveryRequest)
			if error != nil {
				fmt.Printf("Error deleting MultiCloudResourceDiscovery %s %s, It is possible that the resource is already deleted. Please verify manually \n", multiCloudResourceDiscoveryId, error)
				continue
			}
		}
	}
	return nil
}

func getDbmulticloudMultiCloudResourceDiscoveryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MultiCloudResourceDiscoveryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	oracleDbAzureConnectorId := "oracle_db_connector_id"
	multiCloudResourceDiscoveryClient := acctest.GetTestClients(&schema.ResourceData{}).MultiCloudResourceDiscoveryClient()

	listMultiCloudResourceDiscoveriesRequest := oci_dbmulticloud.ListMultiCloudResourceDiscoveriesRequest{}
	listMultiCloudResourceDiscoveriesRequest.CompartmentId = &compartmentId
	listMultiCloudResourceDiscoveriesRequest.ResourceType = oci_dbmulticloud.MultiCloudResourceDiscoveryResourceTypeVaults
	listMultiCloudResourceDiscoveriesRequest.OracleDbAzureConnectorId = &oracleDbAzureConnectorId
	listMultiCloudResourceDiscoveriesResponse, err := multiCloudResourceDiscoveryClient.ListMultiCloudResourceDiscoveries(context.Background(), listMultiCloudResourceDiscoveriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MultiCloudResourceDiscovery list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, multiCloudResourceDiscovery := range listMultiCloudResourceDiscoveriesResponse.Items {
		id := *multiCloudResourceDiscovery.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MultiCloudResourceDiscoveryId", id)
	}
	return resourceIds, nil
}
