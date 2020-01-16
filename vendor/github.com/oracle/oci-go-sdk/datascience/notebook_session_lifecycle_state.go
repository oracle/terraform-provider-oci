// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

// NotebookSessionLifecycleStateEnum Enum with underlying type: string
type NotebookSessionLifecycleStateEnum string

// Set of constants representing the allowable values for NotebookSessionLifecycleStateEnum
const (
	NotebookSessionLifecycleStateCreating NotebookSessionLifecycleStateEnum = "CREATING"
	NotebookSessionLifecycleStateActive   NotebookSessionLifecycleStateEnum = "ACTIVE"
	NotebookSessionLifecycleStateDeleting NotebookSessionLifecycleStateEnum = "DELETING"
	NotebookSessionLifecycleStateDeleted  NotebookSessionLifecycleStateEnum = "DELETED"
	NotebookSessionLifecycleStateFailed   NotebookSessionLifecycleStateEnum = "FAILED"
	NotebookSessionLifecycleStateInactive NotebookSessionLifecycleStateEnum = "INACTIVE"
	NotebookSessionLifecycleStateUpdating NotebookSessionLifecycleStateEnum = "UPDATING"
)

var mappingNotebookSessionLifecycleState = map[string]NotebookSessionLifecycleStateEnum{
	"CREATING": NotebookSessionLifecycleStateCreating,
	"ACTIVE":   NotebookSessionLifecycleStateActive,
	"DELETING": NotebookSessionLifecycleStateDeleting,
	"DELETED":  NotebookSessionLifecycleStateDeleted,
	"FAILED":   NotebookSessionLifecycleStateFailed,
	"INACTIVE": NotebookSessionLifecycleStateInactive,
	"UPDATING": NotebookSessionLifecycleStateUpdating,
}

// GetNotebookSessionLifecycleStateEnumValues Enumerates the set of values for NotebookSessionLifecycleStateEnum
func GetNotebookSessionLifecycleStateEnumValues() []NotebookSessionLifecycleStateEnum {
	values := make([]NotebookSessionLifecycleStateEnum, 0)
	for _, v := range mappingNotebookSessionLifecycleState {
		values = append(values, v)
	}
	return values
}
