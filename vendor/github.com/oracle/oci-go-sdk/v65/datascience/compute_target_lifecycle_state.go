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

// ComputeTargetLifecycleStateEnum Enum with underlying type: string
type ComputeTargetLifecycleStateEnum string

// Set of constants representing the allowable values for ComputeTargetLifecycleStateEnum
const (
	ComputeTargetLifecycleStateCreating ComputeTargetLifecycleStateEnum = "CREATING"
	ComputeTargetLifecycleStateActive   ComputeTargetLifecycleStateEnum = "ACTIVE"
	ComputeTargetLifecycleStateDeleting ComputeTargetLifecycleStateEnum = "DELETING"
	ComputeTargetLifecycleStateDeleted  ComputeTargetLifecycleStateEnum = "DELETED"
	ComputeTargetLifecycleStateFailed   ComputeTargetLifecycleStateEnum = "FAILED"
	ComputeTargetLifecycleStateUpdating ComputeTargetLifecycleStateEnum = "UPDATING"
)

var mappingComputeTargetLifecycleStateEnum = map[string]ComputeTargetLifecycleStateEnum{
	"CREATING": ComputeTargetLifecycleStateCreating,
	"ACTIVE":   ComputeTargetLifecycleStateActive,
	"DELETING": ComputeTargetLifecycleStateDeleting,
	"DELETED":  ComputeTargetLifecycleStateDeleted,
	"FAILED":   ComputeTargetLifecycleStateFailed,
	"UPDATING": ComputeTargetLifecycleStateUpdating,
}

var mappingComputeTargetLifecycleStateEnumLowerCase = map[string]ComputeTargetLifecycleStateEnum{
	"creating": ComputeTargetLifecycleStateCreating,
	"active":   ComputeTargetLifecycleStateActive,
	"deleting": ComputeTargetLifecycleStateDeleting,
	"deleted":  ComputeTargetLifecycleStateDeleted,
	"failed":   ComputeTargetLifecycleStateFailed,
	"updating": ComputeTargetLifecycleStateUpdating,
}

// GetComputeTargetLifecycleStateEnumValues Enumerates the set of values for ComputeTargetLifecycleStateEnum
func GetComputeTargetLifecycleStateEnumValues() []ComputeTargetLifecycleStateEnum {
	values := make([]ComputeTargetLifecycleStateEnum, 0)
	for _, v := range mappingComputeTargetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeTargetLifecycleStateEnumStringValues Enumerates the set of values in String for ComputeTargetLifecycleStateEnum
func GetComputeTargetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingComputeTargetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeTargetLifecycleStateEnum(val string) (ComputeTargetLifecycleStateEnum, bool) {
	enum, ok := mappingComputeTargetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
