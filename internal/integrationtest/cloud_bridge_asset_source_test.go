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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// fake
var (
	vcenterEndpoint = `https://11.0.11.130/sdk`
	vaultSecretId   = `${var.vaultId}`
	inventoryId     = `${oci_cloud_bridge_inventory.test_inventory.id}`

	CloudBridgeAssetSourceRequiredOnlyResource = CloudBridgeAssetSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Required, acctest.Create, CloudBridgeAssetSourceRepresentation)

	CloudBridgeAssetSourceResourceConfig = CloudBridgeAssetSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Optional, acctest.Update, CloudBridgeAssetSourceRepresentation)

	CloudBridgeCloudBridgeAssetSourceSingularDataSourceRepresentation = map[string]interface{}{
		"asset_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_asset_source.test_asset_source.id}`},
	}

	CloudBridgeCloudBridgeAssetSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"asset_source_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_bridge_asset_source.test_asset_source.id}`},
	}

	CloudBridgeAssetSourceRepresentation = map[string]interface{}{
		"assets_compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"discovery_credentials":            acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetSourceDiscoveryCredentialsRepresentation},
		"environment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_environment.test_environment.id}`},
		"inventory_id":                     acctest.Representation{RepType: acctest.Required, Create: inventoryId},
		"type":                             acctest.Representation{RepType: acctest.Required, Create: `VMWARE`},
		"vcenter_endpoint":                 acctest.Representation{RepType: acctest.Required, Create: vcenterEndpoint, Update: vcenterEndpoint},
		"are_historical_metrics_collected": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"are_realtime_metrics_collected":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"discovery_schedule_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"replication_credentials":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetSourceReplicationCredentialsRepresentation},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}
	CloudBridgeAssetSourceDiscoveryCredentialsRepresentation = map[string]interface{}{
		"secret_id": acctest.Representation{RepType: acctest.Required, Create: vaultSecretId},
		"type":      acctest.Representation{RepType: acctest.Required, Create: `BASIC`},
	}
	CloudBridgeAssetSourceReplicationCredentialsRepresentation = map[string]interface{}{
		"secret_id": acctest.Representation{RepType: acctest.Required, Create: vaultSecretId},
		"type":      acctest.Representation{RepType: acctest.Required, Create: `BASIC`},
	}

	CloudBridgeAssetSourceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_discovery_schedule", "test_discovery_schedule", acctest.Required, acctest.Create, CloudBridgeDiscoveryScheduleRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, CloudBridgeEnvironmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Required, acctest.Create, CloudBridgeInventoryRepresentation)

	//acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Required, acctest.Create, CloudBridgeInventoryRepresentation)
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeAssetSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeAssetSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("vaultId")
	vaultSecretIdVariableStr := fmt.Sprintf("variable \"vaultId\" { default = \"%s\" }\n", vaultSecretId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	variableStr := compartmentIdVariableStr + vaultSecretIdVariableStr

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_bridge_asset_source.test_asset_source"
	datasourceName := "data.oci_cloud_bridge_asset_sources.test_asset_sources"
	singularDatasourceName := "data.oci_cloud_bridge_asset_source.test_asset_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+CloudBridgeAssetSourceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Optional, acctest.Create, CloudBridgeAssetSourceRepresentation), "cloudbridge", "assetSource", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeAssetSourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + CloudBridgeAssetSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Required, acctest.Create, CloudBridgeAssetSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "assets_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "discovery_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "discovery_credentials.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "discovery_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "VMWARE"),
				resource.TestCheckResourceAttr(resourceName, "vcenter_endpoint", vcenterEndpoint),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variableStr + CloudBridgeAssetSourceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variableStr + CloudBridgeAssetSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Optional, acctest.Create, CloudBridgeAssetSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_historical_metrics_collected", "false"),
				resource.TestCheckResourceAttr(resourceName, "are_realtime_metrics_collected", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "assets_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "discovery_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "discovery_credentials.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "discovery_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "discovery_schedule_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_credentials.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "VMWARE"),
				resource.TestCheckResourceAttr(resourceName, "vcenter_endpoint", vcenterEndpoint),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + variableStr + compartmentIdUVariableStr + CloudBridgeAssetSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudBridgeAssetSourceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_historical_metrics_collected", "false"),
				resource.TestCheckResourceAttr(resourceName, "are_realtime_metrics_collected", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "assets_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "discovery_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "discovery_credentials.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "discovery_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "discovery_schedule_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_credentials.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "VMWARE"),
				resource.TestCheckResourceAttr(resourceName, "vcenter_endpoint", vcenterEndpoint),

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
			Config: config + variableStr + CloudBridgeAssetSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Optional, acctest.Update, CloudBridgeAssetSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "are_historical_metrics_collected", "true"),
				resource.TestCheckResourceAttr(resourceName, "are_realtime_metrics_collected", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "assets_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "discovery_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "discovery_credentials.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "discovery_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "discovery_schedule_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_credentials.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "replication_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "VMWARE"),
				resource.TestCheckResourceAttr(resourceName, "vcenter_endpoint", vcenterEndpoint),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_asset_sources", "test_asset_sources", acctest.Optional, acctest.Update, CloudBridgeCloudBridgeAssetSourceDataSourceRepresentation) +
				variableStr + CloudBridgeAssetSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Optional, acctest.Update, CloudBridgeAssetSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "asset_source_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "asset_source_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Required, acctest.Create, CloudBridgeCloudBridgeAssetSourceSingularDataSourceRepresentation) +
				variableStr + CloudBridgeAssetSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "asset_source_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "are_historical_metrics_collected", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "are_realtime_metrics_collected", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_credentials.0.type", "BASIC"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "VMWARE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vcenter_endpoint", vcenterEndpoint),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudBridgeAssetSourceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudBridgeAssetSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DiscoveryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_bridge_asset_source" {
			noResourceFound = false
			request := oci_cloud_bridge.GetAssetSourceRequest{}

			tmp := rs.Primary.ID
			request.AssetSourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")

			response, err := client.GetAssetSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_bridge.AssetSourceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudBridgeAssetSource") {
		resource.AddTestSweepers("CloudBridgeAssetSource", &resource.Sweeper{
			Name:         "CloudBridgeAssetSource",
			Dependencies: acctest.DependencyGraph["assetSource"],
			F:            sweepCloudBridgeAssetSourceResource,
		})
	}
}

func sweepCloudBridgeAssetSourceResource(compartment string) error {
	discoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DiscoveryClient()
	assetSourceIds, err := getCloudBridgeAssetSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, assetSourceId := range assetSourceIds {
		if ok := acctest.SweeperDefaultResourceId[assetSourceId]; !ok {
			deleteAssetSourceRequest := oci_cloud_bridge.DeleteAssetSourceRequest{}

			deleteAssetSourceRequest.AssetSourceId = &assetSourceId

			deleteAssetSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")
			_, error := discoveryClient.DeleteAssetSource(context.Background(), deleteAssetSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting AssetSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", assetSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &assetSourceId, CloudBridgeAssetSourceSweepWaitCondition, time.Duration(3*time.Minute),
				CloudBridgeAssetSourceSweepResponseFetchOperation, "cloud_bridge", true)
		}
	}
	return nil
}

func getCloudBridgeAssetSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AssetSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	discoveryClient := acctest.GetTestClients(&schema.ResourceData{}).DiscoveryClient()

	listAssetSourcesRequest := oci_cloud_bridge.ListAssetSourcesRequest{}
	listAssetSourcesRequest.CompartmentId = &compartmentId
	listAssetSourcesRequest.LifecycleState = oci_cloud_bridge.ListAssetSourcesLifecycleStateActive
	listAssetSourcesResponse, err := discoveryClient.ListAssetSources(context.Background(), listAssetSourcesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AssetSource list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, assetSource := range listAssetSourcesResponse.Items {
		id := *assetSource.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AssetSourceId", id)
	}
	return resourceIds, nil
}

func CloudBridgeAssetSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if assetSourceResponse, ok := response.Response.(oci_cloud_bridge.GetAssetSourceResponse); ok {
		return assetSourceResponse.GetLifecycleState() != oci_cloud_bridge.AssetSourceLifecycleStateDeleted
	}
	return false
}

func CloudBridgeAssetSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DiscoveryClient().GetAssetSource(context.Background(), oci_cloud_bridge.GetAssetSourceRequest{
		AssetSourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
