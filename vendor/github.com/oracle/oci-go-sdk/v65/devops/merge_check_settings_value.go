// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"strings"
)

// MergeCheckSettingsValueEnum Enum with underlying type: string
type MergeCheckSettingsValueEnum string

// Set of constants representing the allowable values for MergeCheckSettingsValueEnum
const (
	MergeCheckSettingsValueEnabled  MergeCheckSettingsValueEnum = "ENABLED"
	MergeCheckSettingsValueDisabled MergeCheckSettingsValueEnum = "DISABLED"
)

var mappingMergeCheckSettingsValueEnum = map[string]MergeCheckSettingsValueEnum{
	"ENABLED":  MergeCheckSettingsValueEnabled,
	"DISABLED": MergeCheckSettingsValueDisabled,
}

var mappingMergeCheckSettingsValueEnumLowerCase = map[string]MergeCheckSettingsValueEnum{
	"enabled":  MergeCheckSettingsValueEnabled,
	"disabled": MergeCheckSettingsValueDisabled,
}

// GetMergeCheckSettingsValueEnumValues Enumerates the set of values for MergeCheckSettingsValueEnum
func GetMergeCheckSettingsValueEnumValues() []MergeCheckSettingsValueEnum {
	values := make([]MergeCheckSettingsValueEnum, 0)
	for _, v := range mappingMergeCheckSettingsValueEnum {
		values = append(values, v)
	}
	return values
}

// GetMergeCheckSettingsValueEnumStringValues Enumerates the set of values in String for MergeCheckSettingsValueEnum
func GetMergeCheckSettingsValueEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingMergeCheckSettingsValueEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMergeCheckSettingsValueEnum(val string) (MergeCheckSettingsValueEnum, bool) {
	enum, ok := mappingMergeCheckSettingsValueEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
