package oda

import (
	"fmt"

	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportOdaOdaPrivateEndpointScanProxyHints.GetIdFn = getOdaOdaPrivateEndpointScanProxyId
	exportOdaOdaPrivateEndpointScanProxyHints.ProcessDiscoveredResourcesFn = processPrivateEndpointScanProxy
	tf_export.RegisterCompartmentGraphs("oda", odaResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func getOdaOdaPrivateEndpointScanProxyId(resource *tf_export.OCIResource) (string, error) {

	odaPrivateEndpointId := resource.Parent.Id

	odaPrivateEndpointScanProxyId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find odaPrivateEndpointScanProxyId for Oda OdaPrivateEndpointScanProxy")
	}
	return GetOdaPrivateEndpointScanProxyCompositeId(odaPrivateEndpointId, odaPrivateEndpointScanProxyId), nil
}

func processPrivateEndpointScanProxy(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		odaPrivateEndpointScanProxyId := resource.Id
		odaPrivateEndpointId := resource.Parent.Id
		resource.ImportId = GetOdaPrivateEndpointScanProxyCompositeId(odaPrivateEndpointId, odaPrivateEndpointScanProxyId)
	}
	return resources, nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportOdaOdaInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_oda_oda_instance",
	DatasourceClass:      "oci_oda_oda_instances",
	DatasourceItemsAttr:  "oda_instances",
	ResourceAbbreviation: "oda_instance",
	DiscoverableLifecycleStates: []string{
		string(oci_oda.OdaInstanceLifecycleStateActive),
	},
}

var exportOdaOdaPrivateEndpointAttachmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_oda_oda_private_endpoint_attachment",
	DatasourceClass:        "oci_oda_oda_private_endpoint_attachments",
	DatasourceItemsAttr:    "oda_private_endpoint_attachment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oda_private_endpoint_attachment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_oda.OdaPrivateEndpointAttachmentLifecycleStateActive),
	},
}

var exportOdaOdaPrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_oda_oda_private_endpoint",
	DatasourceClass:        "oci_oda_oda_private_endpoints",
	DatasourceItemsAttr:    "oda_private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oda_private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_oda.OdaPrivateEndpointLifecycleStateActive),
	},
}

var exportOdaOdaPrivateEndpointScanProxyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_oda_oda_private_endpoint_scan_proxy",
	DatasourceClass:        "oci_oda_oda_private_endpoint_scan_proxies",
	DatasourceItemsAttr:    "oda_private_endpoint_scan_proxy_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "oda_private_endpoint_scan_proxy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_oda.OdaPrivateEndpointScanProxyLifecycleStateActive),
	},
}

var odaResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOdaOdaInstanceHints},
		{TerraformResourceHints: exportOdaOdaPrivateEndpointHints},
	},
	"oci_oda_oda_private_endpoint": {
		{
			TerraformResourceHints: exportOdaOdaPrivateEndpointAttachmentHints,
			DatasourceQueryParams: map[string]string{
				"oda_private_endpoint_id": "id",
			},
		},
		{
			TerraformResourceHints: exportOdaOdaPrivateEndpointScanProxyHints,
			DatasourceQueryParams: map[string]string{
				"oda_private_endpoint_id": "id",
			},
		},
	},
}
