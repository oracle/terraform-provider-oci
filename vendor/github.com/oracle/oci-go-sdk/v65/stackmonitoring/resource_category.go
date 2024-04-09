// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// ResourceCategoryEnum Enum with underlying type: string
type ResourceCategoryEnum string

// Set of constants representing the allowable values for ResourceCategoryEnum
const (
	ResourceCategoryApplication ResourceCategoryEnum = "APPLICATION"
	ResourceCategoryDatabase    ResourceCategoryEnum = "DATABASE"
	ResourceCategoryMiddleware  ResourceCategoryEnum = "MIDDLEWARE"
	ResourceCategoryUnknown     ResourceCategoryEnum = "UNKNOWN"
)

var mappingResourceCategoryEnum = map[string]ResourceCategoryEnum{
	"APPLICATION": ResourceCategoryApplication,
	"DATABASE":    ResourceCategoryDatabase,
	"MIDDLEWARE":  ResourceCategoryMiddleware,
	"UNKNOWN":     ResourceCategoryUnknown,
}

var mappingResourceCategoryEnumLowerCase = map[string]ResourceCategoryEnum{
	"application": ResourceCategoryApplication,
	"database":    ResourceCategoryDatabase,
	"middleware":  ResourceCategoryMiddleware,
	"unknown":     ResourceCategoryUnknown,
}

// GetResourceCategoryEnumValues Enumerates the set of values for ResourceCategoryEnum
func GetResourceCategoryEnumValues() []ResourceCategoryEnum {
	values := make([]ResourceCategoryEnum, 0)
	for _, v := range mappingResourceCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceCategoryEnumStringValues Enumerates the set of values in String for ResourceCategoryEnum
func GetResourceCategoryEnumStringValues() []string {
	return []string{
		"APPLICATION",
		"DATABASE",
		"MIDDLEWARE",
		"UNKNOWN",
	}
}

// GetMappingResourceCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceCategoryEnum(val string) (ResourceCategoryEnum, bool) {
	enum, ok := mappingResourceCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
