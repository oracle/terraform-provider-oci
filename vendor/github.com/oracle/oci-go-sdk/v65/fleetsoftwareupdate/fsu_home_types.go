// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// FsuHomeTypesEnum Enum with underlying type: string
type FsuHomeTypesEnum string

// Set of constants representing the allowable values for FsuHomeTypesEnum
const (
	FsuHomeTypesDbhome    FsuHomeTypesEnum = "DBHOME"
	FsuHomeTypesVmcluster FsuHomeTypesEnum = "VMCLUSTER"
)

var mappingFsuHomeTypesEnum = map[string]FsuHomeTypesEnum{
	"DBHOME":    FsuHomeTypesDbhome,
	"VMCLUSTER": FsuHomeTypesVmcluster,
}

var mappingFsuHomeTypesEnumLowerCase = map[string]FsuHomeTypesEnum{
	"dbhome":    FsuHomeTypesDbhome,
	"vmcluster": FsuHomeTypesVmcluster,
}

// GetFsuHomeTypesEnumValues Enumerates the set of values for FsuHomeTypesEnum
func GetFsuHomeTypesEnumValues() []FsuHomeTypesEnum {
	values := make([]FsuHomeTypesEnum, 0)
	for _, v := range mappingFsuHomeTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetFsuHomeTypesEnumStringValues Enumerates the set of values in String for FsuHomeTypesEnum
func GetFsuHomeTypesEnumStringValues() []string {
	return []string{
		"DBHOME",
		"VMCLUSTER",
	}
}

// GetMappingFsuHomeTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFsuHomeTypesEnum(val string) (FsuHomeTypesEnum, bool) {
	enum, ok := mappingFsuHomeTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
