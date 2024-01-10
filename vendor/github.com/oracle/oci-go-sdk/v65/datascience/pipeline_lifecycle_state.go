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

// PipelineLifecycleStateEnum Enum with underlying type: string
type PipelineLifecycleStateEnum string

// Set of constants representing the allowable values for PipelineLifecycleStateEnum
const (
	PipelineLifecycleStateCreating PipelineLifecycleStateEnum = "CREATING"
	PipelineLifecycleStateActive   PipelineLifecycleStateEnum = "ACTIVE"
	PipelineLifecycleStateDeleting PipelineLifecycleStateEnum = "DELETING"
	PipelineLifecycleStateFailed   PipelineLifecycleStateEnum = "FAILED"
	PipelineLifecycleStateDeleted  PipelineLifecycleStateEnum = "DELETED"
)

var mappingPipelineLifecycleStateEnum = map[string]PipelineLifecycleStateEnum{
	"CREATING": PipelineLifecycleStateCreating,
	"ACTIVE":   PipelineLifecycleStateActive,
	"DELETING": PipelineLifecycleStateDeleting,
	"FAILED":   PipelineLifecycleStateFailed,
	"DELETED":  PipelineLifecycleStateDeleted,
}

var mappingPipelineLifecycleStateEnumLowerCase = map[string]PipelineLifecycleStateEnum{
	"creating": PipelineLifecycleStateCreating,
	"active":   PipelineLifecycleStateActive,
	"deleting": PipelineLifecycleStateDeleting,
	"failed":   PipelineLifecycleStateFailed,
	"deleted":  PipelineLifecycleStateDeleted,
}

// GetPipelineLifecycleStateEnumValues Enumerates the set of values for PipelineLifecycleStateEnum
func GetPipelineLifecycleStateEnumValues() []PipelineLifecycleStateEnum {
	values := make([]PipelineLifecycleStateEnum, 0)
	for _, v := range mappingPipelineLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineLifecycleStateEnumStringValues Enumerates the set of values in String for PipelineLifecycleStateEnum
func GetPipelineLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"DELETED",
	}
}

// GetMappingPipelineLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineLifecycleStateEnum(val string) (PipelineLifecycleStateEnum, bool) {
	enum, ok := mappingPipelineLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
