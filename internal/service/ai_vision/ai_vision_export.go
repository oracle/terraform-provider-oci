package ai_vision

import (
	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("ai_vision", aiVisionResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAiVisionProjectHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_vision_project",
	DatasourceClass:        "oci_ai_vision_projects",
	DatasourceItemsAttr:    "project_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "project",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_vision.ProjectLifecycleStateActive),
	},
}

var exportAiVisionModelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_vision_model",
	DatasourceClass:        "oci_ai_vision_models",
	DatasourceItemsAttr:    "model_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "model",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_vision.ModelLifecycleStateActive),
	},
}

var exportAiVisionVisionPrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_vision_vision_private_endpoint",
	DatasourceClass:        "oci_ai_vision_vision_private_endpoints",
	DatasourceItemsAttr:    "vision_private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "vision_private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_vision.VisionPrivateEndpointLifecycleStateActive),
	},
}

var exportAiVisionStreamSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_vision_stream_source",
	DatasourceClass:        "oci_ai_vision_stream_sources",
	DatasourceItemsAttr:    "stream_source_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "stream_source",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_vision.StreamSourceLifecycleStateActive),
	},
}

var exportAiVisionStreamGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_vision_stream_group",
	DatasourceClass:        "oci_ai_vision_stream_groups",
	DatasourceItemsAttr:    "stream_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "stream_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_vision.StreamGroupLifecycleStateActive),
	},
}

var exportAiVisionStreamJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_vision_stream_job",
	DatasourceClass:        "oci_ai_vision_stream_jobs",
	DatasourceItemsAttr:    "stream_job_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "stream_job",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_vision.StreamJobLifecycleStateActive),
		string(oci_ai_vision.StreamJobLifecycleStateNeedsAttention),
	},
}

var exportAiVisionDocumentJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_ai_vision_document_job",
	DatasourceClass:      "oci_ai_vision_document_job",
	ResourceAbbreviation: "document_job",
	DiscoverableLifecycleStates: []string{
		string(oci_ai_vision.DocumentJobLifecycleStateSucceeded),
	},
}

var exportAiVisionImageJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_ai_vision_image_job",
	DatasourceClass:      "oci_ai_vision_image_job",
	ResourceAbbreviation: "image_job",
	DiscoverableLifecycleStates: []string{
		string(oci_ai_vision.ImageJobLifecycleStateSucceeded),
	},
}

var aiVisionResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAiVisionProjectHints},
		{TerraformResourceHints: exportAiVisionModelHints},
		{TerraformResourceHints: exportAiVisionVisionPrivateEndpointHints},
		{TerraformResourceHints: exportAiVisionStreamSourceHints},
		{TerraformResourceHints: exportAiVisionStreamGroupHints},
		{TerraformResourceHints: exportAiVisionStreamJobHints},
	},
	"oci_ai_vision_document_job": {
		{
			TerraformResourceHints: exportAiVisionDocumentJobHints,
			DatasourceQueryParams: map[string]string{
				"document_job_id": "id",
			},
		},
	},
	"oci_ai_vision_image_job": {
		{
			TerraformResourceHints: exportAiVisionImageJobHints,
			DatasourceQueryParams: map[string]string{
				"image_job_id": "id",
			},
		},
	},
}
