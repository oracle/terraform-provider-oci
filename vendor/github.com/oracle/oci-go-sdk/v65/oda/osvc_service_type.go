// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// OsvcServiceTypeEnum Enum with underlying type: string
type OsvcServiceTypeEnum string

// Set of constants representing the allowable values for OsvcServiceTypeEnum
const (
	OsvcServiceTypeOsvc   OsvcServiceTypeEnum = "OSVC"
	OsvcServiceTypeFusion OsvcServiceTypeEnum = "FUSION"
)

var mappingOsvcServiceTypeEnum = map[string]OsvcServiceTypeEnum{
	"OSVC":   OsvcServiceTypeOsvc,
	"FUSION": OsvcServiceTypeFusion,
}

var mappingOsvcServiceTypeEnumLowerCase = map[string]OsvcServiceTypeEnum{
	"osvc":   OsvcServiceTypeOsvc,
	"fusion": OsvcServiceTypeFusion,
}

// GetOsvcServiceTypeEnumValues Enumerates the set of values for OsvcServiceTypeEnum
func GetOsvcServiceTypeEnumValues() []OsvcServiceTypeEnum {
	values := make([]OsvcServiceTypeEnum, 0)
	for _, v := range mappingOsvcServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOsvcServiceTypeEnumStringValues Enumerates the set of values in String for OsvcServiceTypeEnum
func GetOsvcServiceTypeEnumStringValues() []string {
	return []string{
		"OSVC",
		"FUSION",
	}
}

// GetMappingOsvcServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsvcServiceTypeEnum(val string) (OsvcServiceTypeEnum, bool) {
	enum, ok := mappingOsvcServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
