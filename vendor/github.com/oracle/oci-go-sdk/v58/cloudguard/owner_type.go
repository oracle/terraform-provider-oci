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

// OwnerTypeEnum Enum with underlying type: string
type OwnerTypeEnum string

// Set of constants representing the allowable values for OwnerTypeEnum
const (
	OwnerTypeCustomer OwnerTypeEnum = "CUSTOMER"
	OwnerTypeOracle   OwnerTypeEnum = "ORACLE"
)

var mappingOwnerTypeEnum = map[string]OwnerTypeEnum{
	"CUSTOMER": OwnerTypeCustomer,
	"ORACLE":   OwnerTypeOracle,
}

// GetOwnerTypeEnumValues Enumerates the set of values for OwnerTypeEnum
func GetOwnerTypeEnumValues() []OwnerTypeEnum {
	values := make([]OwnerTypeEnum, 0)
	for _, v := range mappingOwnerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOwnerTypeEnumStringValues Enumerates the set of values in String for OwnerTypeEnum
func GetOwnerTypeEnumStringValues() []string {
	return []string{
		"CUSTOMER",
		"ORACLE",
	}
}

// GetMappingOwnerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOwnerTypeEnum(val string) (OwnerTypeEnum, bool) {
	mappingOwnerTypeEnumIgnoreCase := make(map[string]OwnerTypeEnum)
	for k, v := range mappingOwnerTypeEnum {
		mappingOwnerTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOwnerTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
