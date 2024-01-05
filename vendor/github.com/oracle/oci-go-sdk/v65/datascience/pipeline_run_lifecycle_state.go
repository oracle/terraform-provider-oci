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

// PipelineRunLifecycleStateEnum Enum with underlying type: string
type PipelineRunLifecycleStateEnum string

// Set of constants representing the allowable values for PipelineRunLifecycleStateEnum
const (
	PipelineRunLifecycleStateAccepted   PipelineRunLifecycleStateEnum = "ACCEPTED"
	PipelineRunLifecycleStateInProgress PipelineRunLifecycleStateEnum = "IN_PROGRESS"
	PipelineRunLifecycleStateFailed     PipelineRunLifecycleStateEnum = "FAILED"
	PipelineRunLifecycleStateSucceeded  PipelineRunLifecycleStateEnum = "SUCCEEDED"
	PipelineRunLifecycleStateCanceling  PipelineRunLifecycleStateEnum = "CANCELING"
	PipelineRunLifecycleStateCanceled   PipelineRunLifecycleStateEnum = "CANCELED"
	PipelineRunLifecycleStateDeleting   PipelineRunLifecycleStateEnum = "DELETING"
	PipelineRunLifecycleStateDeleted    PipelineRunLifecycleStateEnum = "DELETED"
)

var mappingPipelineRunLifecycleStateEnum = map[string]PipelineRunLifecycleStateEnum{
	"ACCEPTED":    PipelineRunLifecycleStateAccepted,
	"IN_PROGRESS": PipelineRunLifecycleStateInProgress,
	"FAILED":      PipelineRunLifecycleStateFailed,
	"SUCCEEDED":   PipelineRunLifecycleStateSucceeded,
	"CANCELING":   PipelineRunLifecycleStateCanceling,
	"CANCELED":    PipelineRunLifecycleStateCanceled,
	"DELETING":    PipelineRunLifecycleStateDeleting,
	"DELETED":     PipelineRunLifecycleStateDeleted,
}

var mappingPipelineRunLifecycleStateEnumLowerCase = map[string]PipelineRunLifecycleStateEnum{
	"accepted":    PipelineRunLifecycleStateAccepted,
	"in_progress": PipelineRunLifecycleStateInProgress,
	"failed":      PipelineRunLifecycleStateFailed,
	"succeeded":   PipelineRunLifecycleStateSucceeded,
	"canceling":   PipelineRunLifecycleStateCanceling,
	"canceled":    PipelineRunLifecycleStateCanceled,
	"deleting":    PipelineRunLifecycleStateDeleting,
	"deleted":     PipelineRunLifecycleStateDeleted,
}

// GetPipelineRunLifecycleStateEnumValues Enumerates the set of values for PipelineRunLifecycleStateEnum
func GetPipelineRunLifecycleStateEnumValues() []PipelineRunLifecycleStateEnum {
	values := make([]PipelineRunLifecycleStateEnum, 0)
	for _, v := range mappingPipelineRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineRunLifecycleStateEnumStringValues Enumerates the set of values in String for PipelineRunLifecycleStateEnum
func GetPipelineRunLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingPipelineRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineRunLifecycleStateEnum(val string) (PipelineRunLifecycleStateEnum, bool) {
	enum, ok := mappingPipelineRunLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
