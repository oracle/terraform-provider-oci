package zpr

import (
	oci_zpr "github.com/oracle/oci-go-sdk/v65/zpr"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterTenancyGraphs("zpr", zprResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
// Hints for discovering and exporting this resource to configuration and state files
var exportZprConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_zpr_configuration",
	DatasourceClass:      "oci_zpr_configuration",
	ResourceAbbreviation: "configuration",
	DiscoverableLifecycleStates: []string{
		string(oci_zpr.ConfigurationLifecycleStateActive),
	},
}

var exportZprZprPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_zpr_zpr_policy",
	DatasourceClass:        "oci_zpr_zpr_policies",
	DatasourceItemsAttr:    "zpr_policies",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "zpr_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_zpr.ZprPolicyLifecycleStateActive),
		string(oci_zpr.ZprPolicyLifecycleStateNeedsAttention),
	},
}

var zprResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportZprConfigurationHints},
		{TerraformResourceHints: exportZprZprPolicyHints},
	},
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportZprZprPolicyHints,
			DatasourceQueryParams:  map[string]string{"compartment_id": "id"},
		},
	},
}
