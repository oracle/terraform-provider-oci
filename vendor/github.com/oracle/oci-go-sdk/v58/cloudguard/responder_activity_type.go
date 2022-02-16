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

// ResponderActivityTypeEnum Enum with underlying type: string
type ResponderActivityTypeEnum string

// Set of constants representing the allowable values for ResponderActivityTypeEnum
const (
	ResponderActivityTypeStarted   ResponderActivityTypeEnum = "STARTED"
	ResponderActivityTypeCompleted ResponderActivityTypeEnum = "COMPLETED"
)

var mappingResponderActivityTypeEnum = map[string]ResponderActivityTypeEnum{
	"STARTED":   ResponderActivityTypeStarted,
	"COMPLETED": ResponderActivityTypeCompleted,
}

// GetResponderActivityTypeEnumValues Enumerates the set of values for ResponderActivityTypeEnum
func GetResponderActivityTypeEnumValues() []ResponderActivityTypeEnum {
	values := make([]ResponderActivityTypeEnum, 0)
	for _, v := range mappingResponderActivityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderActivityTypeEnumStringValues Enumerates the set of values in String for ResponderActivityTypeEnum
func GetResponderActivityTypeEnumStringValues() []string {
	return []string{
		"STARTED",
		"COMPLETED",
	}
}

// GetMappingResponderActivityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderActivityTypeEnum(val string) (ResponderActivityTypeEnum, bool) {
	mappingResponderActivityTypeEnumIgnoreCase := make(map[string]ResponderActivityTypeEnum)
	for k, v := range mappingResponderActivityTypeEnum {
		mappingResponderActivityTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingResponderActivityTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
