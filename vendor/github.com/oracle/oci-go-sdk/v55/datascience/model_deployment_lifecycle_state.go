// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

// ModelDeploymentLifecycleStateEnum Enum with underlying type: string
type ModelDeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for ModelDeploymentLifecycleStateEnum
const (
	ModelDeploymentLifecycleStateCreating       ModelDeploymentLifecycleStateEnum = "CREATING"
	ModelDeploymentLifecycleStateActive         ModelDeploymentLifecycleStateEnum = "ACTIVE"
	ModelDeploymentLifecycleStateDeleting       ModelDeploymentLifecycleStateEnum = "DELETING"
	ModelDeploymentLifecycleStateFailed         ModelDeploymentLifecycleStateEnum = "FAILED"
	ModelDeploymentLifecycleStateInactive       ModelDeploymentLifecycleStateEnum = "INACTIVE"
	ModelDeploymentLifecycleStateUpdating       ModelDeploymentLifecycleStateEnum = "UPDATING"
	ModelDeploymentLifecycleStateDeleted        ModelDeploymentLifecycleStateEnum = "DELETED"
	ModelDeploymentLifecycleStateNeedsAttention ModelDeploymentLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingModelDeploymentLifecycleState = map[string]ModelDeploymentLifecycleStateEnum{
	"CREATING":        ModelDeploymentLifecycleStateCreating,
	"ACTIVE":          ModelDeploymentLifecycleStateActive,
	"DELETING":        ModelDeploymentLifecycleStateDeleting,
	"FAILED":          ModelDeploymentLifecycleStateFailed,
	"INACTIVE":        ModelDeploymentLifecycleStateInactive,
	"UPDATING":        ModelDeploymentLifecycleStateUpdating,
	"DELETED":         ModelDeploymentLifecycleStateDeleted,
	"NEEDS_ATTENTION": ModelDeploymentLifecycleStateNeedsAttention,
}

// GetModelDeploymentLifecycleStateEnumValues Enumerates the set of values for ModelDeploymentLifecycleStateEnum
func GetModelDeploymentLifecycleStateEnumValues() []ModelDeploymentLifecycleStateEnum {
	values := make([]ModelDeploymentLifecycleStateEnum, 0)
	for _, v := range mappingModelDeploymentLifecycleState {
		values = append(values, v)
	}
	return values
}
