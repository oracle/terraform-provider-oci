// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_datascience_private_endpoint", DatasciencePrivateEndpointResource())
	tfresource.RegisterResource("oci_datascience_job", DatascienceJobResource())
	tfresource.RegisterResource("oci_datascience_job_run", DatascienceJobRunResource())
	tfresource.RegisterResource("oci_datascience_ml_application", DatascienceMlApplicationResource())
	tfresource.RegisterResource("oci_datascience_ml_application_implementation", DatascienceMlApplicationImplementationResource())
	tfresource.RegisterResource("oci_datascience_ml_application_instance", DatascienceMlApplicationInstanceResource())
	tfresource.RegisterResource("oci_datascience_model", DatascienceModelResource())
	tfresource.RegisterResource("oci_datascience_model_custom_metadata_artifact", DatascienceModelCustomMetadataArtifactResource())
	tfresource.RegisterResource("oci_datascience_model_defined_metadata_artifact", DatascienceModelDefinedMetadataArtifactResource())
	tfresource.RegisterResource("oci_datascience_model_deployment", DatascienceModelDeploymentResource())
	tfresource.RegisterResource("oci_datascience_model_provenance", DatascienceModelProvenanceResource())
	tfresource.RegisterResource("oci_datascience_model_version_set", DatascienceModelVersionSetResource())
	tfresource.RegisterResource("oci_datascience_model_artifact_export", DatascienceModelArtifactExportResource())
	tfresource.RegisterResource("oci_datascience_model_artifact_import", DatascienceModelArtifactImportResource())
	tfresource.RegisterResource("oci_datascience_notebook_session", DatascienceNotebookSessionResource())
	tfresource.RegisterResource("oci_datascience_pipeline", DatasciencePipelineResource())
	tfresource.RegisterResource("oci_datascience_pipeline_run", DatasciencePipelineRunResource())
	tfresource.RegisterResource("oci_datascience_project", DatascienceProjectResource())
	tfresource.RegisterResource("oci_datascience_schedule", DatascienceScheduleResource())
}
