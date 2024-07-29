// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// CycleTypesEnum Enum with underlying type: string
type CycleTypesEnum string

// Set of constants representing the allowable values for CycleTypesEnum
const (
	CycleTypesPatch CycleTypesEnum = "PATCH"
)

var mappingCycleTypesEnum = map[string]CycleTypesEnum{
	"PATCH": CycleTypesPatch,
}

var mappingCycleTypesEnumLowerCase = map[string]CycleTypesEnum{
	"patch": CycleTypesPatch,
}

// GetCycleTypesEnumValues Enumerates the set of values for CycleTypesEnum
func GetCycleTypesEnumValues() []CycleTypesEnum {
	values := make([]CycleTypesEnum, 0)
	for _, v := range mappingCycleTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetCycleTypesEnumStringValues Enumerates the set of values in String for CycleTypesEnum
func GetCycleTypesEnumStringValues() []string {
	return []string{
		"PATCH",
	}
}

// GetMappingCycleTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCycleTypesEnum(val string) (CycleTypesEnum, bool) {
	enum, ok := mappingCycleTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
