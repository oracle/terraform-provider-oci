// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// ResponderDimensionEnum Enum with underlying type: string
type ResponderDimensionEnum string

// Set of constants representing the allowable values for ResponderDimensionEnum
const (
	ResponderDimensionResponderRuleType        ResponderDimensionEnum = "RESPONDER_RULE_TYPE"
	ResponderDimensionResponderExecutionStatus ResponderDimensionEnum = "RESPONDER_EXECUTION_STATUS"
)

var mappingResponderDimensionEnum = map[string]ResponderDimensionEnum{
	"RESPONDER_RULE_TYPE":        ResponderDimensionResponderRuleType,
	"RESPONDER_EXECUTION_STATUS": ResponderDimensionResponderExecutionStatus,
}

// GetResponderDimensionEnumValues Enumerates the set of values for ResponderDimensionEnum
func GetResponderDimensionEnumValues() []ResponderDimensionEnum {
	values := make([]ResponderDimensionEnum, 0)
	for _, v := range mappingResponderDimensionEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderDimensionEnumStringValues Enumerates the set of values in String for ResponderDimensionEnum
func GetResponderDimensionEnumStringValues() []string {
	return []string{
		"RESPONDER_RULE_TYPE",
		"RESPONDER_EXECUTION_STATUS",
	}
}

// GetMappingResponderDimensionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderDimensionEnum(val string) (ResponderDimensionEnum, bool) {
	mappingResponderDimensionEnumIgnoreCase := make(map[string]ResponderDimensionEnum)
	for k, v := range mappingResponderDimensionEnum {
		mappingResponderDimensionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingResponderDimensionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
