// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"strings"
)

// OrchestrationVariableTypeEnumEnum Enum with underlying type: string
type OrchestrationVariableTypeEnumEnum string

// Set of constants representing the allowable values for OrchestrationVariableTypeEnumEnum
const (
	OrchestrationVariableTypeEnumString  OrchestrationVariableTypeEnumEnum = "STRING"
	OrchestrationVariableTypeEnumInteger OrchestrationVariableTypeEnumEnum = "INTEGER"
)

var mappingOrchestrationVariableTypeEnumEnum = map[string]OrchestrationVariableTypeEnumEnum{
	"STRING":  OrchestrationVariableTypeEnumString,
	"INTEGER": OrchestrationVariableTypeEnumInteger,
}

// GetOrchestrationVariableTypeEnumEnumValues Enumerates the set of values for OrchestrationVariableTypeEnumEnum
func GetOrchestrationVariableTypeEnumEnumValues() []OrchestrationVariableTypeEnumEnum {
	values := make([]OrchestrationVariableTypeEnumEnum, 0)
	for _, v := range mappingOrchestrationVariableTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetOrchestrationVariableTypeEnumEnumStringValues Enumerates the set of values in String for OrchestrationVariableTypeEnumEnum
func GetOrchestrationVariableTypeEnumEnumStringValues() []string {
	return []string{
		"STRING",
		"INTEGER",
	}
}

// GetMappingOrchestrationVariableTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOrchestrationVariableTypeEnumEnum(val string) (OrchestrationVariableTypeEnumEnum, bool) {
	mappingOrchestrationVariableTypeEnumEnumIgnoreCase := make(map[string]OrchestrationVariableTypeEnumEnum)
	for k, v := range mappingOrchestrationVariableTypeEnumEnum {
		mappingOrchestrationVariableTypeEnumEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOrchestrationVariableTypeEnumEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
