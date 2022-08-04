package resourcemanager

import (
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("resourcemanager", resourcemanagerResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportResourcemanagerPrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_resourcemanager_private_endpoint",
	DatasourceClass:        "oci_resourcemanager_private_endpoints",
	DatasourceItemsAttr:    "private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_resourcemanager.PrivateEndpointLifecycleStateActive),
	},
}

var resourcemanagerResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportResourcemanagerPrivateEndpointHints},
	},
}
