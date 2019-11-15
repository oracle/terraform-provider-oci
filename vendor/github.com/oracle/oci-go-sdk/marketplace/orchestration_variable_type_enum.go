// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// OrchestrationVariableTypeEnumEnum Enum with underlying type: string
type OrchestrationVariableTypeEnumEnum string

// Set of constants representing the allowable values for OrchestrationVariableTypeEnumEnum
const (
	OrchestrationVariableTypeEnumString  OrchestrationVariableTypeEnumEnum = "STRING"
	OrchestrationVariableTypeEnumInteger OrchestrationVariableTypeEnumEnum = "INTEGER"
)

var mappingOrchestrationVariableTypeEnum = map[string]OrchestrationVariableTypeEnumEnum{
	"STRING":  OrchestrationVariableTypeEnumString,
	"INTEGER": OrchestrationVariableTypeEnumInteger,
}

// GetOrchestrationVariableTypeEnumEnumValues Enumerates the set of values for OrchestrationVariableTypeEnumEnum
func GetOrchestrationVariableTypeEnumEnumValues() []OrchestrationVariableTypeEnumEnum {
	values := make([]OrchestrationVariableTypeEnumEnum, 0)
	for _, v := range mappingOrchestrationVariableTypeEnum {
		values = append(values, v)
	}
	return values
}
