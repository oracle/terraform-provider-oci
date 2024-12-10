// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_datascience_containers", DatascienceContainersDataSource())
	tfresource.RegisterDatasource("oci_datascience_private_endpoint", DatasciencePrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_datascience_private_endpoints", DatasciencePrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_datascience_fast_launch_job_configs", DatascienceFastLaunchJobConfigsDataSource())
	tfresource.RegisterDatasource("oci_datascience_job", DatascienceJobDataSource())
	tfresource.RegisterDatasource("oci_datascience_job_run", DatascienceJobRunDataSource())
	tfresource.RegisterDatasource("oci_datascience_job_runs", DatascienceJobRunsDataSource())
	tfresource.RegisterDatasource("oci_datascience_job_shapes", DatascienceJobShapesDataSource())
	tfresource.RegisterDatasource("oci_datascience_jobs", DatascienceJobsDataSource())
	tfresource.RegisterDatasource("oci_datascience_ml_application", DatascienceMlApplicationDataSource())
	tfresource.RegisterDatasource("oci_datascience_ml_application_implementation", DatascienceMlApplicationImplementationDataSource())
	tfresource.RegisterDatasource("oci_datascience_ml_application_implementations", DatascienceMlApplicationImplementationsDataSource())
	tfresource.RegisterDatasource("oci_datascience_ml_application_instance", DatascienceMlApplicationInstanceDataSource())
	tfresource.RegisterDatasource("oci_datascience_ml_application_instances", DatascienceMlApplicationInstancesDataSource())
	tfresource.RegisterDatasource("oci_datascience_ml_applications", DatascienceMlApplicationsDataSource())
	tfresource.RegisterDatasource("oci_datascience_model", DatascienceModelDataSource())
	tfresource.RegisterDatasource("oci_datascience_model_custom_metadata_artifact_content", DatascienceModelCustomMetadataArtifactContentDataSource())
	tfresource.RegisterDatasource("oci_datascience_model_defined_metadata_artifact_content", DatascienceModelDefinedMetadataArtifactContentDataSource())
	tfresource.RegisterDatasource("oci_datascience_model_deployment", DatascienceModelDeploymentDataSource())
	tfresource.RegisterDatasource("oci_datascience_model_deployment_shapes", DatascienceModelDeploymentShapesDataSource())
	tfresource.RegisterDatasource("oci_datascience_model_deployments", DatascienceModelDeploymentsDataSource())
	tfresource.RegisterDatasource("oci_datascience_model_provenance", DatascienceModelProvenanceDataSource())
	tfresource.RegisterDatasource("oci_datascience_model_version_set", DatascienceModelVersionSetDataSource())
	tfresource.RegisterDatasource("oci_datascience_model_version_sets", DatascienceModelVersionSetsDataSource())
	tfresource.RegisterDatasource("oci_datascience_models", DatascienceModelsDataSource())
	tfresource.RegisterDatasource("oci_datascience_notebook_session", DatascienceNotebookSessionDataSource())
	tfresource.RegisterDatasource("oci_datascience_notebook_session_shapes", DatascienceNotebookSessionShapesDataSource())
	tfresource.RegisterDatasource("oci_datascience_notebook_sessions", DatascienceNotebookSessionsDataSource())
	tfresource.RegisterDatasource("oci_datascience_pipeline", DatasciencePipelineDataSource())
	tfresource.RegisterDatasource("oci_datascience_pipeline_run", DatasciencePipelineRunDataSource())
	tfresource.RegisterDatasource("oci_datascience_pipeline_runs", DatasciencePipelineRunsDataSource())
	tfresource.RegisterDatasource("oci_datascience_pipelines", DatasciencePipelinesDataSource())
	tfresource.RegisterDatasource("oci_datascience_project", DatascienceProjectDataSource())
	tfresource.RegisterDatasource("oci_datascience_projects", DatascienceProjectsDataSource())
	tfresource.RegisterDatasource("oci_datascience_schedule", DatascienceScheduleDataSource())
	tfresource.RegisterDatasource("oci_datascience_schedules", DatascienceSchedulesDataSource())
}
