// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
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
	WorkRequestOperationTypeProjectDelete             WorkRequestOperationTypeEnum = "PROJECT_DELETE"
	WorkRequestOperationTypeWorkrequestCancel         WorkRequestOperationTypeEnum = "WORKREQUEST_CANCEL"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"NOTEBOOK_SESSION_CREATE":     WorkRequestOperationTypeNotebookSessionCreate,
	"NOTEBOOK_SESSION_DELETE":     WorkRequestOperationTypeNotebookSessionDelete,
	"NOTEBOOK_SESSION_ACTIVATE":   WorkRequestOperationTypeNotebookSessionActivate,
	"NOTEBOOK_SESSION_DEACTIVATE": WorkRequestOperationTypeNotebookSessionDeactivate,
	"PROJECT_DELETE":              WorkRequestOperationTypeProjectDelete,
	"WORKREQUEST_CANCEL":          WorkRequestOperationTypeWorkrequestCancel,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		values = append(values, v)
	}
	return values
}
