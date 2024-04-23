package cluster_placement_groups

import (
	oci_cluster_placement_groups "github.com/oracle/oci-go-sdk/v65/clusterplacementgroups"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("cluster_placement_groups", clusterPlacementGroupsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportClusterPlacementGroupsClusterPlacementGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cluster_placement_groups_cluster_placement_group",
	DatasourceClass:        "oci_cluster_placement_groups_cluster_placement_groups",
	DatasourceItemsAttr:    "cluster_placement_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "cluster_placement_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cluster_placement_groups.ClusterPlacementGroupLifecycleStateActive),
	},
}

var clusterPlacementGroupsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportClusterPlacementGroupsClusterPlacementGroupHints},
	},
}
