// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// ObjectTypeEnum Enum with underlying type: string
type ObjectTypeEnum string

// Set of constants representing the allowable values for ObjectTypeEnum
const (
	ObjectTypeTable          ObjectTypeEnum = "TABLE"
	ObjectTypeEditioningView ObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingObjectTypeEnum = map[string]ObjectTypeEnum{
	"TABLE":           ObjectTypeTable,
	"EDITIONING_VIEW": ObjectTypeEditioningView,
}

var mappingObjectTypeEnumLowerCase = map[string]ObjectTypeEnum{
	"table":           ObjectTypeTable,
	"editioning_view": ObjectTypeEditioningView,
}

// GetObjectTypeEnumValues Enumerates the set of values for ObjectTypeEnum
func GetObjectTypeEnumValues() []ObjectTypeEnum {
	values := make([]ObjectTypeEnum, 0)
	for _, v := range mappingObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectTypeEnumStringValues Enumerates the set of values in String for ObjectTypeEnum
func GetObjectTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectTypeEnum(val string) (ObjectTypeEnum, bool) {
	enum, ok := mappingObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
