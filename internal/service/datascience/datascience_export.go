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

var exportDatascienceModelVersionSetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_model_version_set",
	DatasourceClass:        "oci_datascience_model_version_sets",
	DatasourceItemsAttr:    "model_version_sets",
	ResourceAbbreviation:   "model_version_set",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.ModelVersionSetLifecycleStateActive),
	},
}

var exportDatasciencePipelineRunHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_pipeline_run",
	DatasourceClass:        "oci_datascience_pipeline_runs",
	DatasourceItemsAttr:    "pipeline_runs",
	ResourceAbbreviation:   "pipeline_run",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.PipelineRunLifecycleStateSucceeded),
	},
}

var exportDatasciencePipelineHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_pipeline",
	DatasourceClass:        "oci_datascience_pipelines",
	DatasourceItemsAttr:    "pipelines",
	ResourceAbbreviation:   "pipeline",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.PipelineLifecycleStateActive),
	},
}

var exportDatascienceDataSciencePrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_private_endpoint",
	DatasourceClass:        "oci_datascience_private_endpoints",
	DatasourceItemsAttr:    "data_science_private_endpoints",
	ResourceAbbreviation:   "data_science_private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.DataSciencePrivateEndpointLifecycleStateActive),
		string(oci_datascience.DataSciencePrivateEndpointLifecycleStateNeedsAttention),
	},
}

var exportDatascienceScheduleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_schedule",
	DatasourceClass:        "oci_datascience_schedules",
	DatasourceItemsAttr:    "schedules",
	ResourceAbbreviation:   "schedule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.ScheduleLifecycleStateActive),
	},
}

var exportDatascienceModelDefinedMetadataArtifactHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_datascience_model_defined_metadata_artifact",
	ResourceAbbreviation: "model_defined_metadata_artifact",
}

var exportDatascienceModelCustomMetadataArtifactHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_datascience_model_custom_metadata_artifact",
	ResourceAbbreviation: "model_custom_metadata_artifact",
}

var exportDatascienceMlApplicationImplementationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_ml_application_implementation",
	DatasourceClass:        "oci_datascience_ml_application_implementations",
	DatasourceItemsAttr:    "ml_application_implementation_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ml_application_implementation",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.MlApplicationImplementationLifecycleStateActive),
		string(oci_datascience.MlApplicationImplementationLifecycleStateNeedsAttention),
	},
}

var exportDatascienceMlApplicationInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_ml_application_instance",
	DatasourceClass:        "oci_datascience_ml_application_instances",
	DatasourceItemsAttr:    "ml_application_instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ml_application_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.MlApplicationInstanceLifecycleStateActive),
		string(oci_datascience.MlApplicationInstanceLifecycleStateNeedsAttention),
	},
}

var exportDatascienceMlApplicationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_datascience_ml_application",
	DatasourceClass:        "oci_datascience_ml_applications",
	DatasourceItemsAttr:    "ml_application_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ml_application",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_datascience.MlApplicationLifecycleStateActive),
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
		{TerraformResourceHints: exportDatascienceModelVersionSetHints},
		{TerraformResourceHints: exportDatasciencePipelineRunHints},
		{TerraformResourceHints: exportDatascienceModelVersionSetHints},
		{TerraformResourceHints: exportDatasciencePipelineHints},
		{TerraformResourceHints: exportDatascienceDataSciencePrivateEndpointHints},
		{TerraformResourceHints: exportDatascienceScheduleHints},
		{TerraformResourceHints: exportDatascienceMlApplicationImplementationHints},
		{TerraformResourceHints: exportDatascienceMlApplicationInstanceHints},
		{TerraformResourceHints: exportDatascienceMlApplicationHints},
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
