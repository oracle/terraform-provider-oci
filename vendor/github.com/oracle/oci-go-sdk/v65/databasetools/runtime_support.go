// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// RuntimeSupportEnum Enum with underlying type: string
type RuntimeSupportEnum string

// Set of constants representing the allowable values for RuntimeSupportEnum
const (
	RuntimeSupportSupported   RuntimeSupportEnum = "SUPPORTED"
	RuntimeSupportUnsupported RuntimeSupportEnum = "UNSUPPORTED"
)

var mappingRuntimeSupportEnum = map[string]RuntimeSupportEnum{
	"SUPPORTED":   RuntimeSupportSupported,
	"UNSUPPORTED": RuntimeSupportUnsupported,
}

var mappingRuntimeSupportEnumLowerCase = map[string]RuntimeSupportEnum{
	"supported":   RuntimeSupportSupported,
	"unsupported": RuntimeSupportUnsupported,
}

// GetRuntimeSupportEnumValues Enumerates the set of values for RuntimeSupportEnum
func GetRuntimeSupportEnumValues() []RuntimeSupportEnum {
	values := make([]RuntimeSupportEnum, 0)
	for _, v := range mappingRuntimeSupportEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeSupportEnumStringValues Enumerates the set of values in String for RuntimeSupportEnum
func GetRuntimeSupportEnumStringValues() []string {
	return []string{
		"SUPPORTED",
		"UNSUPPORTED",
	}
}

// GetMappingRuntimeSupportEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeSupportEnum(val string) (RuntimeSupportEnum, bool) {
	enum, ok := mappingRuntimeSupportEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
