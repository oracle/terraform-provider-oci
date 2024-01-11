package generative_ai

import (
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("generative_ai", generativeAiResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportGenerativeAiDedicatedAiClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_dedicated_ai_cluster",
	DatasourceClass:        "oci_generative_ai_dedicated_ai_clusters",
	DatasourceItemsAttr:    "dedicated_ai_cluster_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "dedicated_ai_cluster",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_generative_ai.DedicatedAiClusterLifecycleStateActive),
		string(oci_generative_ai.DedicatedAiClusterLifecycleStateNeedsAttention),
	},
}

var exportGenerativeAiEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_endpoint",
	DatasourceClass:        "oci_generative_ai_endpoints",
	DatasourceItemsAttr:    "endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_generative_ai.EndpointLifecycleStateActive),
	},
}

var exportGenerativeAiModelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_model",
	DatasourceClass:        "oci_generative_ai_models",
	DatasourceItemsAttr:    "model_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "model",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_generative_ai.ModelLifecycleStateActive),
	},
}

var generativeAiResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportGenerativeAiDedicatedAiClusterHints},
		{TerraformResourceHints: exportGenerativeAiEndpointHints},
		{TerraformResourceHints: exportGenerativeAiModelHints},
	},
}
