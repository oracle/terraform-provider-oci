// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// KeyGenerationTypeEnum Enum with underlying type: string
type KeyGenerationTypeEnum string

// Set of constants representing the allowable values for KeyGenerationTypeEnum
const (
	KeyGenerationTypeSystem KeyGenerationTypeEnum = "SYSTEM"
	KeyGenerationTypeByok   KeyGenerationTypeEnum = "BYOK"
)

var mappingKeyGenerationTypeEnum = map[string]KeyGenerationTypeEnum{
	"SYSTEM": KeyGenerationTypeSystem,
	"BYOK":   KeyGenerationTypeByok,
}

var mappingKeyGenerationTypeEnumLowerCase = map[string]KeyGenerationTypeEnum{
	"system": KeyGenerationTypeSystem,
	"byok":   KeyGenerationTypeByok,
}

// GetKeyGenerationTypeEnumValues Enumerates the set of values for KeyGenerationTypeEnum
func GetKeyGenerationTypeEnumValues() []KeyGenerationTypeEnum {
	values := make([]KeyGenerationTypeEnum, 0)
	for _, v := range mappingKeyGenerationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyGenerationTypeEnumStringValues Enumerates the set of values in String for KeyGenerationTypeEnum
func GetKeyGenerationTypeEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"BYOK",
	}
}

// GetMappingKeyGenerationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyGenerationTypeEnum(val string) (KeyGenerationTypeEnum, bool) {
	enum, ok := mappingKeyGenerationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
