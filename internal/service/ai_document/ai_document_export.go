package ai_document

import (
	oci_ai_document "github.com/oracle/oci-go-sdk/v65/aidocument"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("ai_document", aiDocumentResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAiDocumentProcessorJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_ai_document_processor_job",
	DatasourceClass:      "oci_ai_document_processor_job",
	ResourceAbbreviation: "processor_job",
	DiscoverableLifecycleStates: []string{
		string(oci_ai_document.ProcessorJobLifecycleStateSucceeded),
	},
}

var exportAiDocumentProjectHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_document_project",
	DatasourceClass:        "oci_ai_document_projects",
	DatasourceItemsAttr:    "project_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "project",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_document.ProjectLifecycleStateActive),
	},
}

var exportAiDocumentModelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_document_model",
	DatasourceClass:        "oci_ai_document_models",
	DatasourceItemsAttr:    "model_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "model",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_document.ModelLifecycleStateActive),
	},
}

var aiDocumentResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAiDocumentProjectHints},
		{TerraformResourceHints: exportAiDocumentModelHints},
	},
	"oci_ai_document_processor_job": {
		{
			TerraformResourceHints: exportAiDocumentProcessorJobHints,
			DatasourceQueryParams: map[string]string{
				"processor_job_id": "id",
			},
		},
	},
}
