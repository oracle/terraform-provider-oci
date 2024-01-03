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

// ModelVersionSetLifecycleStateEnum Enum with underlying type: string
type ModelVersionSetLifecycleStateEnum string

// Set of constants representing the allowable values for ModelVersionSetLifecycleStateEnum
const (
	ModelVersionSetLifecycleStateActive   ModelVersionSetLifecycleStateEnum = "ACTIVE"
	ModelVersionSetLifecycleStateDeleting ModelVersionSetLifecycleStateEnum = "DELETING"
	ModelVersionSetLifecycleStateDeleted  ModelVersionSetLifecycleStateEnum = "DELETED"
	ModelVersionSetLifecycleStateFailed   ModelVersionSetLifecycleStateEnum = "FAILED"
)

var mappingModelVersionSetLifecycleStateEnum = map[string]ModelVersionSetLifecycleStateEnum{
	"ACTIVE":   ModelVersionSetLifecycleStateActive,
	"DELETING": ModelVersionSetLifecycleStateDeleting,
	"DELETED":  ModelVersionSetLifecycleStateDeleted,
	"FAILED":   ModelVersionSetLifecycleStateFailed,
}

var mappingModelVersionSetLifecycleStateEnumLowerCase = map[string]ModelVersionSetLifecycleStateEnum{
	"active":   ModelVersionSetLifecycleStateActive,
	"deleting": ModelVersionSetLifecycleStateDeleting,
	"deleted":  ModelVersionSetLifecycleStateDeleted,
	"failed":   ModelVersionSetLifecycleStateFailed,
}

// GetModelVersionSetLifecycleStateEnumValues Enumerates the set of values for ModelVersionSetLifecycleStateEnum
func GetModelVersionSetLifecycleStateEnumValues() []ModelVersionSetLifecycleStateEnum {
	values := make([]ModelVersionSetLifecycleStateEnum, 0)
	for _, v := range mappingModelVersionSetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetModelVersionSetLifecycleStateEnumStringValues Enumerates the set of values in String for ModelVersionSetLifecycleStateEnum
func GetModelVersionSetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingModelVersionSetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelVersionSetLifecycleStateEnum(val string) (ModelVersionSetLifecycleStateEnum, bool) {
	enum, ok := mappingModelVersionSetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
