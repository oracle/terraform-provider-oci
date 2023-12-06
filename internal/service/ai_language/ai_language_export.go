package ai_language

import (
	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("ai_language", aiLanguageResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAiLanguageProjectHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_language_project",
	DatasourceClass:        "oci_ai_language_projects",
	DatasourceItemsAttr:    "project_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "project",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_language.ProjectLifecycleStateActive),
	},
}

var exportAiLanguageModelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_language_model",
	DatasourceClass:        "oci_ai_language_models",
	DatasourceItemsAttr:    "model_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "model",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_language.ModelLifecycleStateActive),
	},
}

var exportAiLanguageEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_language_endpoint",
	DatasourceClass:        "oci_ai_language_endpoints",
	DatasourceItemsAttr:    "endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_language.EndpointLifecycleStateActive),
	},
}

var aiLanguageResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAiLanguageProjectHints},
		{TerraformResourceHints: exportAiLanguageModelHints},
		{TerraformResourceHints: exportAiLanguageEndpointHints},
	},
}
