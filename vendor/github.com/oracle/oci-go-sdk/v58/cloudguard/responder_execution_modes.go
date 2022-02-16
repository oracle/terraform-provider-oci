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

// ResponderExecutionModesEnum Enum with underlying type: string
type ResponderExecutionModesEnum string

// Set of constants representing the allowable values for ResponderExecutionModesEnum
const (
	ResponderExecutionModesManual    ResponderExecutionModesEnum = "MANUAL"
	ResponderExecutionModesAutomated ResponderExecutionModesEnum = "AUTOMATED"
	ResponderExecutionModesAll       ResponderExecutionModesEnum = "ALL"
)

var mappingResponderExecutionModesEnum = map[string]ResponderExecutionModesEnum{
	"MANUAL":    ResponderExecutionModesManual,
	"AUTOMATED": ResponderExecutionModesAutomated,
	"ALL":       ResponderExecutionModesAll,
}

// GetResponderExecutionModesEnumValues Enumerates the set of values for ResponderExecutionModesEnum
func GetResponderExecutionModesEnumValues() []ResponderExecutionModesEnum {
	values := make([]ResponderExecutionModesEnum, 0)
	for _, v := range mappingResponderExecutionModesEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderExecutionModesEnumStringValues Enumerates the set of values in String for ResponderExecutionModesEnum
func GetResponderExecutionModesEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTOMATED",
		"ALL",
	}
}

// GetMappingResponderExecutionModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderExecutionModesEnum(val string) (ResponderExecutionModesEnum, bool) {
	mappingResponderExecutionModesEnumIgnoreCase := make(map[string]ResponderExecutionModesEnum)
	for k, v := range mappingResponderExecutionModesEnum {
		mappingResponderExecutionModesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingResponderExecutionModesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
