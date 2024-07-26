package generative_ai

import (
	"strings"

	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportGenerativeAiModelHints.ProcessDiscoveredResourcesFn = processExcludingBaseModels
	tf_export.RegisterCompartmentGraphs("generative_ai", generativeAiResourceGraph)
}

// Custom models are exposed to user but should not be part of their stack in resource discovery
func processExcludingBaseModels(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}
	for _, resource := range resources {
		modelType := resource.SourceAttributes["type"].(string)
		if strings.Compare(modelType, "BASE") != 0 {
			results = append(results, resource)
		}
	}
	return results, nil
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportGenerativeAiDedicatedAiClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_dedicated_ai_cluster",
	DatasourceClass:        "oci_generative_ai_dedicated_ai_clusters",
	DatasourceItemsAttr:    "dedicated_ai_cluster_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "dedicated_ai_cluster",
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
