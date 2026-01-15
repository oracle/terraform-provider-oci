// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ModelGroupVersionHistoryLifecycleStateEnum Enum with underlying type: string
type ModelGroupVersionHistoryLifecycleStateEnum string

// Set of constants representing the allowable values for ModelGroupVersionHistoryLifecycleStateEnum
const (
	ModelGroupVersionHistoryLifecycleStateActive   ModelGroupVersionHistoryLifecycleStateEnum = "ACTIVE"
	ModelGroupVersionHistoryLifecycleStateDeleted  ModelGroupVersionHistoryLifecycleStateEnum = "DELETED"
	ModelGroupVersionHistoryLifecycleStateFailed   ModelGroupVersionHistoryLifecycleStateEnum = "FAILED"
	ModelGroupVersionHistoryLifecycleStateDeleting ModelGroupVersionHistoryLifecycleStateEnum = "DELETING"
)

var mappingModelGroupVersionHistoryLifecycleStateEnum = map[string]ModelGroupVersionHistoryLifecycleStateEnum{
	"ACTIVE":   ModelGroupVersionHistoryLifecycleStateActive,
	"DELETED":  ModelGroupVersionHistoryLifecycleStateDeleted,
	"FAILED":   ModelGroupVersionHistoryLifecycleStateFailed,
	"DELETING": ModelGroupVersionHistoryLifecycleStateDeleting,
}

var mappingModelGroupVersionHistoryLifecycleStateEnumLowerCase = map[string]ModelGroupVersionHistoryLifecycleStateEnum{
	"active":   ModelGroupVersionHistoryLifecycleStateActive,
	"deleted":  ModelGroupVersionHistoryLifecycleStateDeleted,
	"failed":   ModelGroupVersionHistoryLifecycleStateFailed,
	"deleting": ModelGroupVersionHistoryLifecycleStateDeleting,
}

// GetModelGroupVersionHistoryLifecycleStateEnumValues Enumerates the set of values for ModelGroupVersionHistoryLifecycleStateEnum
func GetModelGroupVersionHistoryLifecycleStateEnumValues() []ModelGroupVersionHistoryLifecycleStateEnum {
	values := make([]ModelGroupVersionHistoryLifecycleStateEnum, 0)
	for _, v := range mappingModelGroupVersionHistoryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetModelGroupVersionHistoryLifecycleStateEnumStringValues Enumerates the set of values in String for ModelGroupVersionHistoryLifecycleStateEnum
func GetModelGroupVersionHistoryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
		"DELETING",
	}
}

// GetMappingModelGroupVersionHistoryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelGroupVersionHistoryLifecycleStateEnum(val string) (ModelGroupVersionHistoryLifecycleStateEnum, bool) {
	enum, ok := mappingModelGroupVersionHistoryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
