package limits

import (
	oci_limits "github.com/oracle/oci-go-sdk/v65/limits"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterTenancyGraphs("limits", limitsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportLimitsQuotaHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_limits_quota",
	DatasourceClass:        "oci_limits_quotas",
	DatasourceItemsAttr:    "quotas",
	ResourceAbbreviation:   "quota",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_limits.QuotaLifecycleStateActive),
	},
}

var limitsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportLimitsQuotaHints},
	},
}
