package analytics

import (
	"fmt"

	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportAnalyticsAnalyticsInstanceResourceGroupHints.GetIdFn = getAnalyticsAnalyticsInstanceResourceGroupId
	tf_export.RegisterCompartmentGraphs("analytics", analyticsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getAnalyticsAnalyticsInstanceResourceGroupId(resource *tf_export.OCIResource) (string, error) {

	analyticsInstanceResourceGroupId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find analyticsInstanceResourceGroupId for Analytics AnalyticsInstanceResourceGroup")
	}
	return analyticsInstanceResourceGroupId, nil
}

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

var exportAnalyticsAnalyticsInstanceResourceGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_analytics_analytics_instance_resource_group",
	DatasourceClass:        "oci_analytics_analytics_instance_resource_groups",
	DatasourceItemsAttr:    "instance_resource_groups",
	ResourceAbbreviation:   "analytics_instance_resource_group",
	RequireResourceRefresh: true,
}

var analyticsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAnalyticsAnalyticsInstanceHints},
	},
	"oci_analytics_analytics_instance": {
		{
			TerraformResourceHints: exportAnalyticsAnalyticsInstanceResourceGroupHints,
			DatasourceQueryParams: map[string]string{
				"analytics_instance_id": "id",
			},
		},
	},
}
