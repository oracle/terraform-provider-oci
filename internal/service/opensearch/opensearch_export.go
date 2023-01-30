package opensearch

import (
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("opensearch", opensearchResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOpensearchOpensearchClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opensearch_opensearch_cluster",
	DatasourceClass:        "oci_opensearch_opensearch_clusters",
	DatasourceItemsAttr:    "opensearch_cluster_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "opensearch_cluster",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opensearch.OpensearchClusterLifecycleStateActive),
	},
}

var opensearchResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOpensearchOpensearchClusterHints},
	},
}
