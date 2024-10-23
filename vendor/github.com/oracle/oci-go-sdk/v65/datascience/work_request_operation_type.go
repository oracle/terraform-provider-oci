// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeNotebookSessionCreate     WorkRequestOperationTypeEnum = "NOTEBOOK_SESSION_CREATE"
	WorkRequestOperationTypeNotebookSessionDelete     WorkRequestOperationTypeEnum = "NOTEBOOK_SESSION_DELETE"
	WorkRequestOperationTypeNotebookSessionActivate   WorkRequestOperationTypeEnum = "NOTEBOOK_SESSION_ACTIVATE"
	WorkRequestOperationTypeNotebookSessionDeactivate WorkRequestOperationTypeEnum = "NOTEBOOK_SESSION_DEACTIVATE"
	WorkRequestOperationTypeModelversionsetDelete     WorkRequestOperationTypeEnum = "MODELVERSIONSET_DELETE"
	WorkRequestOperationTypeExportModelArtifact       WorkRequestOperationTypeEnum = "EXPORT_MODEL_ARTIFACT"
	WorkRequestOperationTypeImportModelArtifact       WorkRequestOperationTypeEnum = "IMPORT_MODEL_ARTIFACT"
	WorkRequestOperationTypeModelDeploymentCreate     WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_CREATE"
	WorkRequestOperationTypeModelDeploymentDelete     WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_DELETE"
	WorkRequestOperationTypeModelDeploymentActivate   WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_ACTIVATE"
	WorkRequestOperationTypeModelDeploymentDeactivate WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_DEACTIVATE"
	WorkRequestOperationTypeModelDeploymentUpdate     WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_UPDATE"
	WorkRequestOperationTypeProjectDelete             WorkRequestOperationTypeEnum = "PROJECT_DELETE"
	WorkRequestOperationTypeWorkrequestCancel         WorkRequestOperationTypeEnum = "WORKREQUEST_CANCEL"
	WorkRequestOperationTypeJobDelete                 WorkRequestOperationTypeEnum = "JOB_DELETE"
	WorkRequestOperationTypePipelineCreate            WorkRequestOperationTypeEnum = "PIPELINE_CREATE"
	WorkRequestOperationTypePipelineDelete            WorkRequestOperationTypeEnum = "PIPELINE_DELETE"
	WorkRequestOperationTypePipelineRunCreate         WorkRequestOperationTypeEnum = "PIPELINE_RUN_CREATE"
	WorkRequestOperationTypePipelineRunCancel         WorkRequestOperationTypeEnum = "PIPELINE_RUN_CANCEL"
	WorkRequestOperationTypePipelineRunDelete         WorkRequestOperationTypeEnum = "PIPELINE_RUN_DELETE"
	WorkRequestOperationTypePrivateEndpointCreate     WorkRequestOperationTypeEnum = "PRIVATE_ENDPOINT_CREATE"
	WorkRequestOperationTypePrivateEndpointDelete     WorkRequestOperationTypeEnum = "PRIVATE_ENDPOINT_DELETE"
	WorkRequestOperationTypePrivateEndpointMove       WorkRequestOperationTypeEnum = "PRIVATE_ENDPOINT_MOVE"
	WorkRequestOperationTypePrivateEndpointUpdate     WorkRequestOperationTypeEnum = "PRIVATE_ENDPOINT_UPDATE"
	WorkRequestOperationTypeRestoreArchivedModel      WorkRequestOperationTypeEnum = "RESTORE_ARCHIVED_MODEL"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"NOTEBOOK_SESSION_CREATE":     WorkRequestOperationTypeNotebookSessionCreate,
	"NOTEBOOK_SESSION_DELETE":     WorkRequestOperationTypeNotebookSessionDelete,
	"NOTEBOOK_SESSION_ACTIVATE":   WorkRequestOperationTypeNotebookSessionActivate,
	"NOTEBOOK_SESSION_DEACTIVATE": WorkRequestOperationTypeNotebookSessionDeactivate,
	"MODELVERSIONSET_DELETE":      WorkRequestOperationTypeModelversionsetDelete,
	"EXPORT_MODEL_ARTIFACT":       WorkRequestOperationTypeExportModelArtifact,
	"IMPORT_MODEL_ARTIFACT":       WorkRequestOperationTypeImportModelArtifact,
	"MODEL_DEPLOYMENT_CREATE":     WorkRequestOperationTypeModelDeploymentCreate,
	"MODEL_DEPLOYMENT_DELETE":     WorkRequestOperationTypeModelDeploymentDelete,
	"MODEL_DEPLOYMENT_ACTIVATE":   WorkRequestOperationTypeModelDeploymentActivate,
	"MODEL_DEPLOYMENT_DEACTIVATE": WorkRequestOperationTypeModelDeploymentDeactivate,
	"MODEL_DEPLOYMENT_UPDATE":     WorkRequestOperationTypeModelDeploymentUpdate,
	"PROJECT_DELETE":              WorkRequestOperationTypeProjectDelete,
	"WORKREQUEST_CANCEL":          WorkRequestOperationTypeWorkrequestCancel,
	"JOB_DELETE":                  WorkRequestOperationTypeJobDelete,
	"PIPELINE_CREATE":             WorkRequestOperationTypePipelineCreate,
	"PIPELINE_DELETE":             WorkRequestOperationTypePipelineDelete,
	"PIPELINE_RUN_CREATE":         WorkRequestOperationTypePipelineRunCreate,
	"PIPELINE_RUN_CANCEL":         WorkRequestOperationTypePipelineRunCancel,
	"PIPELINE_RUN_DELETE":         WorkRequestOperationTypePipelineRunDelete,
	"PRIVATE_ENDPOINT_CREATE":     WorkRequestOperationTypePrivateEndpointCreate,
	"PRIVATE_ENDPOINT_DELETE":     WorkRequestOperationTypePrivateEndpointDelete,
	"PRIVATE_ENDPOINT_MOVE":       WorkRequestOperationTypePrivateEndpointMove,
	"PRIVATE_ENDPOINT_UPDATE":     WorkRequestOperationTypePrivateEndpointUpdate,
	"RESTORE_ARCHIVED_MODEL":      WorkRequestOperationTypeRestoreArchivedModel,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"notebook_session_create":     WorkRequestOperationTypeNotebookSessionCreate,
	"notebook_session_delete":     WorkRequestOperationTypeNotebookSessionDelete,
	"notebook_session_activate":   WorkRequestOperationTypeNotebookSessionActivate,
	"notebook_session_deactivate": WorkRequestOperationTypeNotebookSessionDeactivate,
	"modelversionset_delete":      WorkRequestOperationTypeModelversionsetDelete,
	"export_model_artifact":       WorkRequestOperationTypeExportModelArtifact,
	"import_model_artifact":       WorkRequestOperationTypeImportModelArtifact,
	"model_deployment_create":     WorkRequestOperationTypeModelDeploymentCreate,
	"model_deployment_delete":     WorkRequestOperationTypeModelDeploymentDelete,
	"model_deployment_activate":   WorkRequestOperationTypeModelDeploymentActivate,
	"model_deployment_deactivate": WorkRequestOperationTypeModelDeploymentDeactivate,
	"model_deployment_update":     WorkRequestOperationTypeModelDeploymentUpdate,
	"project_delete":              WorkRequestOperationTypeProjectDelete,
	"workrequest_cancel":          WorkRequestOperationTypeWorkrequestCancel,
	"job_delete":                  WorkRequestOperationTypeJobDelete,
	"pipeline_create":             WorkRequestOperationTypePipelineCreate,
	"pipeline_delete":             WorkRequestOperationTypePipelineDelete,
	"pipeline_run_create":         WorkRequestOperationTypePipelineRunCreate,
	"pipeline_run_cancel":         WorkRequestOperationTypePipelineRunCancel,
	"pipeline_run_delete":         WorkRequestOperationTypePipelineRunDelete,
	"private_endpoint_create":     WorkRequestOperationTypePrivateEndpointCreate,
	"private_endpoint_delete":     WorkRequestOperationTypePrivateEndpointDelete,
	"private_endpoint_move":       WorkRequestOperationTypePrivateEndpointMove,
	"private_endpoint_update":     WorkRequestOperationTypePrivateEndpointUpdate,
	"restore_archived_model":      WorkRequestOperationTypeRestoreArchivedModel,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"NOTEBOOK_SESSION_CREATE",
		"NOTEBOOK_SESSION_DELETE",
		"NOTEBOOK_SESSION_ACTIVATE",
		"NOTEBOOK_SESSION_DEACTIVATE",
		"MODELVERSIONSET_DELETE",
		"EXPORT_MODEL_ARTIFACT",
		"IMPORT_MODEL_ARTIFACT",
		"MODEL_DEPLOYMENT_CREATE",
		"MODEL_DEPLOYMENT_DELETE",
		"MODEL_DEPLOYMENT_ACTIVATE",
		"MODEL_DEPLOYMENT_DEACTIVATE",
		"MODEL_DEPLOYMENT_UPDATE",
		"PROJECT_DELETE",
		"WORKREQUEST_CANCEL",
		"JOB_DELETE",
		"PIPELINE_CREATE",
		"PIPELINE_DELETE",
		"PIPELINE_RUN_CREATE",
		"PIPELINE_RUN_CANCEL",
		"PIPELINE_RUN_DELETE",
		"PRIVATE_ENDPOINT_CREATE",
		"PRIVATE_ENDPOINT_DELETE",
		"PRIVATE_ENDPOINT_MOVE",
		"PRIVATE_ENDPOINT_UPDATE",
		"RESTORE_ARCHIVED_MODEL",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
