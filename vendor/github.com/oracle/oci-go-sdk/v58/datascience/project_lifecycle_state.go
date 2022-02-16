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

// ProjectLifecycleStateEnum Enum with underlying type: string
type ProjectLifecycleStateEnum string

// Set of constants representing the allowable values for ProjectLifecycleStateEnum
const (
	ProjectLifecycleStateActive   ProjectLifecycleStateEnum = "ACTIVE"
	ProjectLifecycleStateDeleting ProjectLifecycleStateEnum = "DELETING"
	ProjectLifecycleStateDeleted  ProjectLifecycleStateEnum = "DELETED"
)

var mappingProjectLifecycleStateEnum = map[string]ProjectLifecycleStateEnum{
	"ACTIVE":   ProjectLifecycleStateActive,
	"DELETING": ProjectLifecycleStateDeleting,
	"DELETED":  ProjectLifecycleStateDeleted,
}

// GetProjectLifecycleStateEnumValues Enumerates the set of values for ProjectLifecycleStateEnum
func GetProjectLifecycleStateEnumValues() []ProjectLifecycleStateEnum {
	values := make([]ProjectLifecycleStateEnum, 0)
	for _, v := range mappingProjectLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProjectLifecycleStateEnumStringValues Enumerates the set of values in String for ProjectLifecycleStateEnum
func GetProjectLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingProjectLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProjectLifecycleStateEnum(val string) (ProjectLifecycleStateEnum, bool) {
	mappingProjectLifecycleStateEnumIgnoreCase := make(map[string]ProjectLifecycleStateEnum)
	for k, v := range mappingProjectLifecycleStateEnum {
		mappingProjectLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingProjectLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
