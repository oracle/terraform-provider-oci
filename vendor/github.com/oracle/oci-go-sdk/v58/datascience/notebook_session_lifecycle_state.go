// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

var mappingNotebookSessionLifecycleStateEnum = map[string]NotebookSessionLifecycleStateEnum{
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
	for _, v := range mappingNotebookSessionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNotebookSessionLifecycleStateEnumStringValues Enumerates the set of values in String for NotebookSessionLifecycleStateEnum
func GetNotebookSessionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
		"UPDATING",
	}
}

// GetMappingNotebookSessionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotebookSessionLifecycleStateEnum(val string) (NotebookSessionLifecycleStateEnum, bool) {
	mappingNotebookSessionLifecycleStateEnumIgnoreCase := make(map[string]NotebookSessionLifecycleStateEnum)
	for k, v := range mappingNotebookSessionLifecycleStateEnum {
		mappingNotebookSessionLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingNotebookSessionLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
