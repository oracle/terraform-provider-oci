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

// CollectionServiceTypesEnum Enum with underlying type: string
type CollectionServiceTypesEnum string

// Set of constants representing the allowable values for CollectionServiceTypesEnum
const (
	CollectionServiceTypesExacs CollectionServiceTypesEnum = "EXACS"
	CollectionServiceTypesExacc CollectionServiceTypesEnum = "EXACC"
)

var mappingCollectionServiceTypesEnum = map[string]CollectionServiceTypesEnum{
	"EXACS": CollectionServiceTypesExacs,
	"EXACC": CollectionServiceTypesExacc,
}

var mappingCollectionServiceTypesEnumLowerCase = map[string]CollectionServiceTypesEnum{
	"exacs": CollectionServiceTypesExacs,
	"exacc": CollectionServiceTypesExacc,
}

// GetCollectionServiceTypesEnumValues Enumerates the set of values for CollectionServiceTypesEnum
func GetCollectionServiceTypesEnumValues() []CollectionServiceTypesEnum {
	values := make([]CollectionServiceTypesEnum, 0)
	for _, v := range mappingCollectionServiceTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetCollectionServiceTypesEnumStringValues Enumerates the set of values in String for CollectionServiceTypesEnum
func GetCollectionServiceTypesEnumStringValues() []string {
	return []string{
		"EXACS",
		"EXACC",
	}
}

// GetMappingCollectionServiceTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCollectionServiceTypesEnum(val string) (CollectionServiceTypesEnum, bool) {
	enum, ok := mappingCollectionServiceTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
