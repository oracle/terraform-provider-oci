// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeNotebookSessionCreate     WorkRequestOperationTypeEnum = "NOTEBOOK_SESSION_CREATE"
	WorkRequestOperationTypeNotebookSessionDelete     WorkRequestOperationTypeEnum = "NOTEBOOK_SESSION_DELETE"
	WorkRequestOperationTypeNotebookSessionActivate   WorkRequestOperationTypeEnum = "NOTEBOOK_SESSION_ACTIVATE"
	WorkRequestOperationTypeNotebookSessionDeactivate WorkRequestOperationTypeEnum = "NOTEBOOK_SESSION_DEACTIVATE"
	WorkRequestOperationTypeModelDeploymentCreate     WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_CREATE"
	WorkRequestOperationTypeModelDeploymentDelete     WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_DELETE"
	WorkRequestOperationTypeModelDeploymentActivate   WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_ACTIVATE"
	WorkRequestOperationTypeModelDeploymentDeactivate WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_DEACTIVATE"
	WorkRequestOperationTypeModelDeploymentUpdate     WorkRequestOperationTypeEnum = "MODEL_DEPLOYMENT_UPDATE"
	WorkRequestOperationTypeProjectDelete             WorkRequestOperationTypeEnum = "PROJECT_DELETE"
	WorkRequestOperationTypeWorkrequestCancel         WorkRequestOperationTypeEnum = "WORKREQUEST_CANCEL"
	WorkRequestOperationTypeJobDelete                 WorkRequestOperationTypeEnum = "JOB_DELETE"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"NOTEBOOK_SESSION_CREATE":     WorkRequestOperationTypeNotebookSessionCreate,
	"NOTEBOOK_SESSION_DELETE":     WorkRequestOperationTypeNotebookSessionDelete,
	"NOTEBOOK_SESSION_ACTIVATE":   WorkRequestOperationTypeNotebookSessionActivate,
	"NOTEBOOK_SESSION_DEACTIVATE": WorkRequestOperationTypeNotebookSessionDeactivate,
	"MODEL_DEPLOYMENT_CREATE":     WorkRequestOperationTypeModelDeploymentCreate,
	"MODEL_DEPLOYMENT_DELETE":     WorkRequestOperationTypeModelDeploymentDelete,
	"MODEL_DEPLOYMENT_ACTIVATE":   WorkRequestOperationTypeModelDeploymentActivate,
	"MODEL_DEPLOYMENT_DEACTIVATE": WorkRequestOperationTypeModelDeploymentDeactivate,
	"MODEL_DEPLOYMENT_UPDATE":     WorkRequestOperationTypeModelDeploymentUpdate,
	"PROJECT_DELETE":              WorkRequestOperationTypeProjectDelete,
	"WORKREQUEST_CANCEL":          WorkRequestOperationTypeWorkrequestCancel,
	"JOB_DELETE":                  WorkRequestOperationTypeJobDelete,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		values = append(values, v)
	}
	return values
}
