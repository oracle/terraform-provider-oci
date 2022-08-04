package optimizer

import (
	oci_optimizer "github.com/oracle/oci-go-sdk/v65/optimizer"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterTenancyGraphs("optimizer", optimizerResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOptimizerProfileHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_optimizer_profile",
	DatasourceClass:        "oci_optimizer_profiles",
	DatasourceItemsAttr:    "profile_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "profile",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_optimizer.LifecycleStateActive),
	},
}

var optimizerResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportOptimizerProfileHints},
	},
}
