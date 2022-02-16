// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"strings"
)

// StrategyParameterTypeEnum Enum with underlying type: string
type StrategyParameterTypeEnum string

// Set of constants representing the allowable values for StrategyParameterTypeEnum
const (
	StrategyParameterTypeString   StrategyParameterTypeEnum = "STRING"
	StrategyParameterTypeBoolean  StrategyParameterTypeEnum = "BOOLEAN"
	StrategyParameterTypeNumber   StrategyParameterTypeEnum = "NUMBER"
	StrategyParameterTypeDatetime StrategyParameterTypeEnum = "DATETIME"
)

var mappingStrategyParameterTypeEnum = map[string]StrategyParameterTypeEnum{
	"STRING":   StrategyParameterTypeString,
	"BOOLEAN":  StrategyParameterTypeBoolean,
	"NUMBER":   StrategyParameterTypeNumber,
	"DATETIME": StrategyParameterTypeDatetime,
}

// GetStrategyParameterTypeEnumValues Enumerates the set of values for StrategyParameterTypeEnum
func GetStrategyParameterTypeEnumValues() []StrategyParameterTypeEnum {
	values := make([]StrategyParameterTypeEnum, 0)
	for _, v := range mappingStrategyParameterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStrategyParameterTypeEnumStringValues Enumerates the set of values in String for StrategyParameterTypeEnum
func GetStrategyParameterTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"BOOLEAN",
		"NUMBER",
		"DATETIME",
	}
}

// GetMappingStrategyParameterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStrategyParameterTypeEnum(val string) (StrategyParameterTypeEnum, bool) {
	mappingStrategyParameterTypeEnumIgnoreCase := make(map[string]StrategyParameterTypeEnum)
	for k, v := range mappingStrategyParameterTypeEnum {
		mappingStrategyParameterTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingStrategyParameterTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
