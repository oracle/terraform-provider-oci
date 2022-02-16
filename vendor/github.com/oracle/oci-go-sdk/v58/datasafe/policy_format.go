// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// PolicyFormatEnum Enum with underlying type: string
type PolicyFormatEnum string

// Set of constants representing the allowable values for PolicyFormatEnum
const (
	PolicyFormatXml PolicyFormatEnum = "XML"
)

var mappingPolicyFormatEnum = map[string]PolicyFormatEnum{
	"XML": PolicyFormatXml,
}

// GetPolicyFormatEnumValues Enumerates the set of values for PolicyFormatEnum
func GetPolicyFormatEnumValues() []PolicyFormatEnum {
	values := make([]PolicyFormatEnum, 0)
	for _, v := range mappingPolicyFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetPolicyFormatEnumStringValues Enumerates the set of values in String for PolicyFormatEnum
func GetPolicyFormatEnumStringValues() []string {
	return []string{
		"XML",
	}
}

// GetMappingPolicyFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPolicyFormatEnum(val string) (PolicyFormatEnum, bool) {
	mappingPolicyFormatEnumIgnoreCase := make(map[string]PolicyFormatEnum)
	for k, v := range mappingPolicyFormatEnum {
		mappingPolicyFormatEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPolicyFormatEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
