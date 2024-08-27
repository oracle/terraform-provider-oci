package desktops

import (
	oci_desktops "github.com/oracle/oci-go-sdk/v65/desktops"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("desktops", desktopsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDesktopsDesktopPoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_desktops_desktop_pool",
	DatasourceClass:        "oci_desktops_desktop_pools",
	DatasourceItemsAttr:    "desktop_pool_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "desktop_pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_desktops.LifecycleStateActive),
	},
}

var desktopsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDesktopsDesktopPoolHints},
	},
}
