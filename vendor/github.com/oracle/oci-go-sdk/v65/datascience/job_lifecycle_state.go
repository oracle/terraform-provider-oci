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

// JobLifecycleStateEnum Enum with underlying type: string
type JobLifecycleStateEnum string

// Set of constants representing the allowable values for JobLifecycleStateEnum
const (
	JobLifecycleStateCreating JobLifecycleStateEnum = "CREATING"
	JobLifecycleStateActive   JobLifecycleStateEnum = "ACTIVE"
	JobLifecycleStateDeleting JobLifecycleStateEnum = "DELETING"
	JobLifecycleStateFailed   JobLifecycleStateEnum = "FAILED"
	JobLifecycleStateDeleted  JobLifecycleStateEnum = "DELETED"
)

var mappingJobLifecycleStateEnum = map[string]JobLifecycleStateEnum{
	"CREATING": JobLifecycleStateCreating,
	"ACTIVE":   JobLifecycleStateActive,
	"DELETING": JobLifecycleStateDeleting,
	"FAILED":   JobLifecycleStateFailed,
	"DELETED":  JobLifecycleStateDeleted,
}

var mappingJobLifecycleStateEnumLowerCase = map[string]JobLifecycleStateEnum{
	"creating": JobLifecycleStateCreating,
	"active":   JobLifecycleStateActive,
	"deleting": JobLifecycleStateDeleting,
	"failed":   JobLifecycleStateFailed,
	"deleted":  JobLifecycleStateDeleted,
}

// GetJobLifecycleStateEnumValues Enumerates the set of values for JobLifecycleStateEnum
func GetJobLifecycleStateEnumValues() []JobLifecycleStateEnum {
	values := make([]JobLifecycleStateEnum, 0)
	for _, v := range mappingJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetJobLifecycleStateEnumStringValues Enumerates the set of values in String for JobLifecycleStateEnum
func GetJobLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"DELETED",
	}
}

// GetMappingJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobLifecycleStateEnum(val string) (JobLifecycleStateEnum, bool) {
	enum, ok := mappingJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
