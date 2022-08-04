package identity_data_plane

import (
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("identity_data_plane", identityDataPlaneResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportIdentityDataPlaneGenerateScopedAccessTokenHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_identity_data_plane_generate_scoped_access_token",
	ResourceAbbreviation: "generate_scoped_access_token",
}

var identityDataPlaneResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}
