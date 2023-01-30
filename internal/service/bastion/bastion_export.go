package bastion

import (
	oci_bastion "github.com/oracle/oci-go-sdk/v65/bastion"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("bastion", bastionResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportBastionBastionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_bastion_bastion",
	DatasourceClass:        "oci_bastion_bastions",
	DatasourceItemsAttr:    "bastions",
	ResourceAbbreviation:   "bastion",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_bastion.BastionLifecycleStateActive),
	},
}

var exportBastionSessionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_bastion_session",
	DatasourceClass:        "oci_bastion_sessions",
	DatasourceItemsAttr:    "sessions",
	ResourceAbbreviation:   "session",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_bastion.SessionLifecycleStateActive),
	},
}

var bastionResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBastionBastionHints},
	},
	"oci_bastion_bastion": {
		{
			TerraformResourceHints: exportBastionSessionHints,
			DatasourceQueryParams: map[string]string{
				"bastion_id": "id",
			},
		},
	},
}
