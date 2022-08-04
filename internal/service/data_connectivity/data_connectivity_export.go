package data_connectivity

import (
	"fmt"

	oci_data_connectivity "github.com/oracle/oci-go-sdk/v65/dataconnectivity"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDataConnectivityRegistryConnectionHints.GetIdFn = getDataConnectivityRegistryConnectionId
	exportDataConnectivityRegistryDataAssetHints.GetIdFn = getDataConnectivityRegistryDataAssetId
	exportDataConnectivityRegistryFolderHints.GetIdFn = getDataConnectivityRegistryFolderId
	tf_export.RegisterCompartmentGraphs("data_connectivity", dataConnectivityResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getDataConnectivityRegistryConnectionId(resource *tf_export.OCIResource) (string, error) {

	connectionKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find connectionKey for DataConnectivity RegistryConnection")
	}
	registryId := resource.Parent.SourceAttributes["registry_id"].(string)
	return GetRegistryConnectionCompositeId(connectionKey, registryId), nil
}

func getDataConnectivityRegistryDataAssetId(resource *tf_export.OCIResource) (string, error) {

	dataAssetKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for DataConnectivity RegistryDataAsset")
	}
	registryId := resource.Parent.Id
	return GetRegistryDataAssetCompositeId(dataAssetKey, registryId), nil
}

func getDataConnectivityRegistryFolderId(resource *tf_export.OCIResource) (string, error) {

	folderKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find folderKey for DataConnectivity RegistryFolder")
	}
	registryId := resource.Parent.Id
	return GetRegistryFolderCompositeId(folderKey, registryId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDataConnectivityRegistryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_connectivity_registry",
	DatasourceClass:        "oci_data_connectivity_registries",
	DatasourceItemsAttr:    "registry_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "registry",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_data_connectivity.RegistryLifecycleStateActive),
	},
}

var exportDataConnectivityRegistryConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_connectivity_registry_connection",
	DatasourceClass:        "oci_data_connectivity_registry_connections",
	DatasourceItemsAttr:    "connection_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "registry_connection",
	RequireResourceRefresh: true,
}

var exportDataConnectivityRegistryDataAssetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_connectivity_registry_data_asset",
	DatasourceClass:        "oci_data_connectivity_registry_data_assets",
	DatasourceItemsAttr:    "data_asset_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "registry_data_asset",
	RequireResourceRefresh: true,
}

var exportDataConnectivityRegistryFolderHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_data_connectivity_registry_folder",
	DatasourceClass:        "oci_data_connectivity_registry_folders",
	DatasourceItemsAttr:    "folder_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "registry_folder",
	RequireResourceRefresh: true,
}

var dataConnectivityResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataConnectivityRegistryHints},
	},
	"oci_data_connectivity_registry": {
		{
			TerraformResourceHints: exportDataConnectivityRegistryConnectionHints,
			DatasourceQueryParams: map[string]string{
				"data_asset_key": "data_asset_key",
				"registry_id":    "id",
			},
		},
		{
			TerraformResourceHints: exportDataConnectivityRegistryDataAssetHints,
			DatasourceQueryParams: map[string]string{
				"registry_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDataConnectivityRegistryFolderHints,
			DatasourceQueryParams: map[string]string{
				"registry_id": "id",
			},
		},
	},
	"oci_data_connectivity_registry_data_asset": {
		{
			TerraformResourceHints: exportDataConnectivityRegistryConnectionHints,
			DatasourceQueryParams: map[string]string{
				"data_asset_key": "key",
				"registry_id":    "registry_id",
			},
		},
	},
}
