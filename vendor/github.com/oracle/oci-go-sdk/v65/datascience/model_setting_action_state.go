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

// ModelSettingActionStateEnum Enum with underlying type: string
type ModelSettingActionStateEnum string

// Set of constants representing the allowable values for ModelSettingActionStateEnum
const (
	ModelSettingActionStatePending   ModelSettingActionStateEnum = "PENDING"
	ModelSettingActionStateFailed    ModelSettingActionStateEnum = "FAILED"
	ModelSettingActionStateSucceeded ModelSettingActionStateEnum = "SUCCEEDED"
)

var mappingModelSettingActionStateEnum = map[string]ModelSettingActionStateEnum{
	"PENDING":   ModelSettingActionStatePending,
	"FAILED":    ModelSettingActionStateFailed,
	"SUCCEEDED": ModelSettingActionStateSucceeded,
}

var mappingModelSettingActionStateEnumLowerCase = map[string]ModelSettingActionStateEnum{
	"pending":   ModelSettingActionStatePending,
	"failed":    ModelSettingActionStateFailed,
	"succeeded": ModelSettingActionStateSucceeded,
}

// GetModelSettingActionStateEnumValues Enumerates the set of values for ModelSettingActionStateEnum
func GetModelSettingActionStateEnumValues() []ModelSettingActionStateEnum {
	values := make([]ModelSettingActionStateEnum, 0)
	for _, v := range mappingModelSettingActionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetModelSettingActionStateEnumStringValues Enumerates the set of values in String for ModelSettingActionStateEnum
func GetModelSettingActionStateEnumStringValues() []string {
	return []string{
		"PENDING",
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingModelSettingActionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelSettingActionStateEnum(val string) (ModelSettingActionStateEnum, bool) {
	enum, ok := mappingModelSettingActionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
