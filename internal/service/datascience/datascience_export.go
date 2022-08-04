package datascience

import (
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDatascienceModelProvenanceHints.GetIdFn = getDatascienceModelProvenanceId
	exportDatascienceModelHints.DefaultValuesForMissingAttributes = map[string]interface{}{
		"artifact_content_length": "0",
	}
	tf_export.RegisterCompartmentGraphs("datascience", datascienceResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getDatascienceModelProvenanceId(resource *tf_export.OCIResource) (string, error) {

	modelId := resource.Parent.Id
	return GetModelProvenanceCompositeId(modelId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDatascienceProjectHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_datascience_project",
	DatasourceClass:      "oci_datascience_projects",
	DatasourceItemsAttr:  "projects",
	ResourceAbbreviation: "project",
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.ProjectLifecycleStateActive),
	},
}

var exportDatascienceNotebookSessionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_notebook_session",
	DatasourceClass:        "oci_datascience_notebook_sessions",
	DatasourceItemsAttr:    "notebook_sessions",
	ResourceAbbreviation:   "notebook_session",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.NotebookSessionLifecycleStateActive),
	},
}

var exportDatascienceModelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_model",
	DatasourceClass:        "oci_datascience_models",
	DatasourceItemsAttr:    "models",
	ResourceAbbreviation:   "model",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.ModelLifecycleStateActive),
	},
}

var exportDatascienceModelProvenanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_datascience_model_provenance",
	DatasourceClass:      "oci_datascience_model_provenance",
	ResourceAbbreviation: "model_provenance",
}

var exportDatascienceModelDeploymentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_datascience_model_deployment",
	DatasourceClass:      "oci_datascience_model_deployments",
	DatasourceItemsAttr:  "model_deployments",
	ResourceAbbreviation: "model_deployment",
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.ModelDeploymentLifecycleStateActive),
		string(oci_datascience.ModelDeploymentLifecycleStateNeedsAttention),
	},
}

var exportDatascienceJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_job",
	DatasourceClass:        "oci_datascience_jobs",
	DatasourceItemsAttr:    "jobs",
	ResourceAbbreviation:   "job",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.JobLifecycleStateActive),
	},
}

var exportDatascienceJobRunHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_job_run",
	DatasourceClass:        "oci_datascience_job_runs",
	DatasourceItemsAttr:    "job_runs",
	ResourceAbbreviation:   "job_run",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.JobRunLifecycleStateSucceeded),
		string(oci_datascience.JobRunLifecycleStateNeedsAttention),
	},
}

var datascienceResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatascienceProjectHints},
		{TerraformResourceHints: exportDatascienceNotebookSessionHints},
		{TerraformResourceHints: exportDatascienceModelHints},
		{TerraformResourceHints: exportDatascienceModelDeploymentHints},
		{TerraformResourceHints: exportDatascienceJobHints},
		{TerraformResourceHints: exportDatascienceJobRunHints},
	},
	"oci_datascience_model": {
		{
			TerraformResourceHints: exportDatascienceModelProvenanceHints,
			DatasourceQueryParams: map[string]string{
				"model_id": "id",
			},
		},
	},
}
