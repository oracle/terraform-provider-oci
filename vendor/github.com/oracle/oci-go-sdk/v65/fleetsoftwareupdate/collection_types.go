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

// CollectionTypesEnum Enum with underlying type: string
type CollectionTypesEnum string

// Set of constants representing the allowable values for CollectionTypesEnum
const (
	CollectionTypesDb CollectionTypesEnum = "DB"
	CollectionTypesGi CollectionTypesEnum = "GI"
)

var mappingCollectionTypesEnum = map[string]CollectionTypesEnum{
	"DB": CollectionTypesDb,
	"GI": CollectionTypesGi,
}

var mappingCollectionTypesEnumLowerCase = map[string]CollectionTypesEnum{
	"db": CollectionTypesDb,
	"gi": CollectionTypesGi,
}

// GetCollectionTypesEnumValues Enumerates the set of values for CollectionTypesEnum
func GetCollectionTypesEnumValues() []CollectionTypesEnum {
	values := make([]CollectionTypesEnum, 0)
	for _, v := range mappingCollectionTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetCollectionTypesEnumStringValues Enumerates the set of values in String for CollectionTypesEnum
func GetCollectionTypesEnumStringValues() []string {
	return []string{
		"DB",
		"GI",
	}
}

// GetMappingCollectionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCollectionTypesEnum(val string) (CollectionTypesEnum, bool) {
	enum, ok := mappingCollectionTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
