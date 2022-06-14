// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v65/dataconnectivity"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"
	tf_client "terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/resourcediscovery"
	"terraform-provider-oci/internal/tfresource"
	"terraform-provider-oci/internal/utils"
)

var (
	RegistryDataAssetRequiredOnlyResource = RegistryDataAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Required, acctest.Create, registryDataAssetRepresentation)

	RegistryDataAssetResourceConfig = RegistryDataAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Optional, acctest.Update, registryDataAssetRepresentation)

	registryDataAssetSingularDataSourceRepresentation = map[string]interface{}{
		"data_asset_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry_data_asset.test_registry_data_asset.key}`},
		"registry_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
	}

	registryDataAssetDataSourceRepresentation = map[string]interface{}{
		"registry_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`}}
	registryDataAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_connectivity_registry_data_asset.test_registry_data_asset.name}`}},
	}

	registryDataAssetRepresentation = map[string]interface{}{
		"identifier":         acctest.Representation{RepType: acctest.Required, Create: `IDENTIFIER`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `dataassetname`},
		"properties":         acctest.Representation{RepType: acctest.Required, Create: map[string]string{"dataAssetType": "POSTGRESQL", "host": "host", "port": "port"}, Update: map[string]string{"dataAssetType": "POSTGRESQL", "host": "host", "port": "port"}},
		"registry_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
		"type":               acctest.Representation{RepType: acctest.Required, Create: `GENERIC_JDBC_DATA_ASSET`},
		"asset_properties":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"assetProperties": "assetProperties"}, Update: map[string]string{"assetProperties2": "assetProperties2"}},
		"default_connection": acctest.RepresentationGroup{RepType: acctest.Required, Group: registryDataAssetDefaultConnectionRepresentation},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"model_type":         acctest.Representation{RepType: acctest.Required, Create: `GENERIC_DATA_ASSET`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: dcmsDataAssetignoreChangesRepresentation},
	}
	registryDataAssetDefaultConnectionRepresentation = map[string]interface{}{
		"identifier":  acctest.Representation{RepType: acctest.Required, Create: `IDENTIFIER`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `defconnname`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"is_default":  acctest.Representation{RepType: acctest.Required, Create: `true`},
		"model_type":  acctest.Representation{RepType: acctest.Required, Create: `GENERIC_CONNECTION`},
		"properties":  acctest.Representation{RepType: acctest.Required, Create: map[string]string{"username": "username1", "password": "password1"}, Update: map[string]string{"username": "username2", "password": "password2"}},
		"type":        acctest.Representation{RepType: acctest.Required, Create: `GENERIC_JDBC_CONNECTION`},
	}
	dcmsDataAssetignoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`default_connection`}},
	}

	RegistryDataAssetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry", "test_registry", acctest.Required, acctest.Create, registryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, userRepresentation)
)

// issue-routing-tag: data_connectivity/default
func TestDataConnectivityRegistryDataAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataConnectivityRegistryDataAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_connectivity_registry_data_asset.test_registry_data_asset"
	datasourceName := "data.oci_data_connectivity_registry_data_assets.test_registry_data_assets"
	singularDatasourceName := "data.oci_data_connectivity_registry_data_asset.test_registry_data_asset"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RegistryDataAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Optional, acctest.Create, registryDataAssetRepresentation), "dataconnectivity", "registryDataAsset", t)

	acctest.ResourceTest(t, testAccCheckDataConnectivityRegistryDataAssetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RegistryDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Required, acctest.Create, registryDataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identifier", "IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "name", "dataassetname"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "GENERIC_JDBC_DATA_ASSET"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RegistryDataAssetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RegistryDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Optional, acctest.Create, registryDataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "default_connection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "default_connection.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "default_connection.0.identifier", "IDENTIFIER"),
				resource.TestCheckResourceAttr(resourceName, "default_connection.0.is_default", "true"),
				resource.TestCheckResourceAttr(resourceName, "default_connection.0.properties.%", "3"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "GENERIC_DATA_ASSET"),
				resource.TestCheckResourceAttr(resourceName, "name", "dataassetname"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "GENERIC_JDBC_DATA_ASSET"),

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
			Config: config + compartmentIdVariableStr + RegistryDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Optional, acctest.Update, registryDataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "default_connection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "properties.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "registry_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_data_assets", "test_registry_data_assets", acctest.Optional, acctest.Update, registryDataAssetDataSourceRepresentation) +
				compartmentIdVariableStr + RegistryDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Optional, acctest.Update, registryDataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "registry_id"),

				resource.TestCheckResourceAttr(datasourceName, "data_asset_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_asset_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_data_asset", "test_registry_data_asset", acctest.Required, acctest.Create, registryDataAssetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RegistryDataAssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "registry_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "default_connection.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "default_connection.0.properties.%", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "GENERIC_DATA_ASSET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "GENERIC_JDBC_DATA_ASSET"),
			),
		},
		// verify resource import
		{
			Config:                  config + RegistryDataAssetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"registry_metadata"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataConnectivityRegistryDataAssetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataConnectivityManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_connectivity_registry_data_asset" {
			noResourceFound = false
			request := oci_data_connectivity.GetDataAssetRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.DataAssetKey = &value
			}

			if value, ok := rs.Primary.Attributes["registry_id"]; ok {
				request.RegistryId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_connectivity")

			_, err := client.GetDataAsset(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DataConnectivityRegistryDataAsset") {
		resource.AddTestSweepers("DataConnectivityRegistryDataAsset", &resource.Sweeper{
			Name:         "DataConnectivityRegistryDataAsset",
			Dependencies: acctest.DependencyGraph["registryDataAsset"],
			F:            sweepDataConnectivityRegistryDataAssetResource,
		})
	}
}

func sweepDataConnectivityRegistryDataAssetResource(compartment string) error {
	dataConnectivityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DataConnectivityManagementClient()
	registryDataAssetIds, err := getRegistryDataAssetIds(compartment)
	if err != nil {
		return err
	}
	for _, registryDataAssetId := range registryDataAssetIds {
		if ok := acctest.SweeperDefaultResourceId[registryDataAssetId]; !ok {
			deleteDataAssetRequest := oci_data_connectivity.DeleteDataAssetRequest{}

			deleteDataAssetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_connectivity")
			_, error := dataConnectivityManagementClient.DeleteDataAsset(context.Background(), deleteDataAssetRequest)
			if error != nil {
				fmt.Printf("Error deleting RegistryDataAsset %s %s, It is possible that the resource is already deleted. Please verify manually \n", registryDataAssetId, error)
				continue
			}
		}
	}
	return nil
}

func getRegistryDataAssetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RegistryDataAssetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataConnectivityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DataConnectivityManagementClient()

	listDataAssetsRequest := oci_data_connectivity.ListDataAssetsRequest{}

	registryIds, error := getRegistryIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting registryId required for RegistryDataAsset resource requests \n")
	}
	for _, registryId := range registryIds {
		listDataAssetsRequest.RegistryId = &registryId

		listDataAssetsResponse, err := dataConnectivityManagementClient.ListDataAssets(context.Background(), listDataAssetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting RegistryDataAsset list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, registryDataAsset := range listDataAssetsResponse.Items {
			id := *registryDataAsset.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RegistryDataAssetId", id)

		}

	}
	return resourceIds, nil
}
