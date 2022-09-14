package containerengine

import (
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportContainerengineNodePoolHints.ProcessDiscoveredResourcesFn = processContainerengineNodePool
	tf_export.RegisterCompartmentGraphs("containerengine", containerengineResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processContainerengineNodePool(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, nodePool := range resources {
		// subnet_ids and quantity_per_subnet are deprecated and conflict with node_config_details
		if _, exists := nodePool.SourceAttributes["node_config_details"]; exists {
			if _, ok := nodePool.SourceAttributes["subnet_ids"]; ok {
				delete(nodePool.SourceAttributes, "subnet_ids")
			}
			if _, ok := nodePool.SourceAttributes["quantity_per_subnet"]; ok {
				delete(nodePool.SourceAttributes, "quantity_per_subnet")
			}
		}
	}
	return resources, nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportContainerengineClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_containerengine_cluster",
	DatasourceClass:        "oci_containerengine_clusters",
	DatasourceItemsAttr:    "clusters",
	ResourceAbbreviation:   "cluster",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_containerengine.ClusterLifecycleStateActive),
	},
}

var exportContainerengineNodePoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_containerengine_node_pool",
	DatasourceClass:        "oci_containerengine_node_pools",
	DatasourceItemsAttr:    "node_pools",
	ResourceAbbreviation:   "node_pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_containerengine.NodePoolLifecycleStateActive),
		string(oci_containerengine.NodePoolLifecycleStateNeedsAttention),
	},
}

var containerengineResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportContainerengineClusterHints},
		{TerraformResourceHints: exportContainerengineNodePoolHints},
	},
}
