package dataintegration

import (
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("dataintegration", dataintegrationResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDataintegrationWorkspaceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataintegration_workspace",
	DatasourceClass:        "oci_dataintegration_workspaces",
	DatasourceItemsAttr:    "workspaces",
	ResourceAbbreviation:   "workspace",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dataintegration.WorkspaceLifecycleStateActive),
	},
}

var dataintegrationResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataintegrationWorkspaceHints},
	},
}
