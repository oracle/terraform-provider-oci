// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// PropertyUnitsEnum Enum with underlying type: string
type PropertyUnitsEnum string

// Set of constants representing the allowable values for PropertyUnitsEnum
const (
	PropertyUnitsPercentage PropertyUnitsEnum = "PERCENTAGE"
	PropertyUnitsMb         PropertyUnitsEnum = "MB"
)

var mappingPropertyUnitsEnum = map[string]PropertyUnitsEnum{
	"PERCENTAGE": PropertyUnitsPercentage,
	"MB":         PropertyUnitsMb,
}

var mappingPropertyUnitsEnumLowerCase = map[string]PropertyUnitsEnum{
	"percentage": PropertyUnitsPercentage,
	"mb":         PropertyUnitsMb,
}

// GetPropertyUnitsEnumValues Enumerates the set of values for PropertyUnitsEnum
func GetPropertyUnitsEnumValues() []PropertyUnitsEnum {
	values := make([]PropertyUnitsEnum, 0)
	for _, v := range mappingPropertyUnitsEnum {
		values = append(values, v)
	}
	return values
}

// GetPropertyUnitsEnumStringValues Enumerates the set of values in String for PropertyUnitsEnum
func GetPropertyUnitsEnumStringValues() []string {
	return []string{
		"PERCENTAGE",
		"MB",
	}
}

// GetMappingPropertyUnitsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPropertyUnitsEnum(val string) (PropertyUnitsEnum, bool) {
	enum, ok := mappingPropertyUnitsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
