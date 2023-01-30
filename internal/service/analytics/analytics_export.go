package analytics

import (
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("analytics", analyticsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAnalyticsAnalyticsInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_analytics_analytics_instance",
	DatasourceClass:        "oci_analytics_analytics_instances",
	DatasourceItemsAttr:    "analytics_instances",
	ResourceAbbreviation:   "analytics_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_analytics.AnalyticsInstanceLifecycleStateActive),
	},
}

var analyticsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAnalyticsAnalyticsInstanceHints},
	},
}
