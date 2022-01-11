// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

// ModelLifecycleStateEnum Enum with underlying type: string
type ModelLifecycleStateEnum string

// Set of constants representing the allowable values for ModelLifecycleStateEnum
const (
	ModelLifecycleStateActive   ModelLifecycleStateEnum = "ACTIVE"
	ModelLifecycleStateDeleted  ModelLifecycleStateEnum = "DELETED"
	ModelLifecycleStateFailed   ModelLifecycleStateEnum = "FAILED"
	ModelLifecycleStateInactive ModelLifecycleStateEnum = "INACTIVE"
)

var mappingModelLifecycleState = map[string]ModelLifecycleStateEnum{
	"ACTIVE":   ModelLifecycleStateActive,
	"DELETED":  ModelLifecycleStateDeleted,
	"FAILED":   ModelLifecycleStateFailed,
	"INACTIVE": ModelLifecycleStateInactive,
}

// GetModelLifecycleStateEnumValues Enumerates the set of values for ModelLifecycleStateEnum
func GetModelLifecycleStateEnumValues() []ModelLifecycleStateEnum {
	values := make([]ModelLifecycleStateEnum, 0)
	for _, v := range mappingModelLifecycleState {
		values = append(values, v)
	}
	return values
}
