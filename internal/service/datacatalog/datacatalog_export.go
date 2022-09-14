package datacatalog

import (
	"fmt"

	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDatacatalogDataAssetHints.GetIdFn = getDatacatalogDataAssetId
	exportDatacatalogConnectionHints.GetIdFn = getDatacatalogConnectionId
	tf_export.RegisterCompartmentGraphs("datacatalog", datacatalogResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getDatacatalogDataAssetId(resource *tf_export.OCIResource) (string, error) {

	catalogId := resource.Parent.Id
	dataAssetKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for Datacatalog DataAsset")
	}
	return GetDataAssetCompositeId(catalogId, dataAssetKey), nil
}

func getDatacatalogConnectionId(resource *tf_export.OCIResource) (string, error) {

	catalogId, ok := resource.Parent.SourceAttributes["catalog_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find catalogId for Datacatalog Connection")
	}
	connectionKey, ok := resource.SourceAttributes["key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find connectionKey for Datacatalog Connection")
	}
	dataAssetKey, ok := resource.SourceAttributes["data_asset_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataAssetKey for Datacatalog Connection")
	}
	return GetConnectionCompositeId(catalogId, connectionKey, dataAssetKey), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDatacatalogCatalogHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_datacatalog_catalog",
	DatasourceClass:      "oci_datacatalog_catalogs",
	DatasourceItemsAttr:  "catalogs",
	ResourceAbbreviation: "catalog",
	DiscoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDatacatalogDataAssetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datacatalog_data_asset",
	DatasourceClass:        "oci_datacatalog_data_assets",
	DatasourceItemsAttr:    "data_asset_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "data_asset",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDatacatalogConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datacatalog_connection",
	DatasourceClass:        "oci_datacatalog_connections",
	DatasourceItemsAttr:    "connection_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "connection",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDatacatalogCatalogPrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_datacatalog_catalog_private_endpoint",
	DatasourceClass:      "oci_datacatalog_catalog_private_endpoints",
	DatasourceItemsAttr:  "catalog_private_endpoints",
	ResourceAbbreviation: "catalog_private_endpoint",
	DiscoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var exportDatacatalogMetastoreHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datacatalog_metastore",
	DatasourceClass:        "oci_datacatalog_metastores",
	DatasourceItemsAttr:    "metastores",
	ResourceAbbreviation:   "metastore",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datacatalog.LifecycleStateActive),
	},
}

var datacatalogResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatacatalogCatalogHints},
		{TerraformResourceHints: exportDatacatalogCatalogPrivateEndpointHints},
		{TerraformResourceHints: exportDatacatalogMetastoreHints},
	},
	"oci_datacatalog_catalog": {
		{
			TerraformResourceHints: exportDatacatalogConnectionHints,
			DatasourceQueryParams: map[string]string{
				"catalog_id":     "id",
				"data_asset_key": "data_asset_key",
			},
		},
		{
			TerraformResourceHints: exportDatacatalogDataAssetHints,
			DatasourceQueryParams: map[string]string{
				"catalog_id": "id",
			},
		},
	},
	"oci_datacatalog_data_asset": {
		{
			TerraformResourceHints: exportDatacatalogConnectionHints,
			DatasourceQueryParams: map[string]string{
				"data_asset_key": "key",
				"catalog_id":     "catalog_id",
			},
		},
	},
}
