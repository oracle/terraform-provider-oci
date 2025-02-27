// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"strings"
)

// ResourcesSeveritiesEnum Enum with underlying type: string
type ResourcesSeveritiesEnum string

// Set of constants representing the allowable values for ResourcesSeveritiesEnum
const (
	ResourcesSeveritiesCritical ResourcesSeveritiesEnum = "CRITICAL"
	ResourcesSeveritiesHigh     ResourcesSeveritiesEnum = "HIGH"
	ResourcesSeveritiesMedium   ResourcesSeveritiesEnum = "MEDIUM"
	ResourcesSeveritiesLow      ResourcesSeveritiesEnum = "LOW"
	ResourcesSeveritiesInfo     ResourcesSeveritiesEnum = "INFO"
	ResourcesSeveritiesNone     ResourcesSeveritiesEnum = "NONE"
)

var mappingResourcesSeveritiesEnum = map[string]ResourcesSeveritiesEnum{
	"CRITICAL": ResourcesSeveritiesCritical,
	"HIGH":     ResourcesSeveritiesHigh,
	"MEDIUM":   ResourcesSeveritiesMedium,
	"LOW":      ResourcesSeveritiesLow,
	"INFO":     ResourcesSeveritiesInfo,
	"NONE":     ResourcesSeveritiesNone,
}

var mappingResourcesSeveritiesEnumLowerCase = map[string]ResourcesSeveritiesEnum{
	"critical": ResourcesSeveritiesCritical,
	"high":     ResourcesSeveritiesHigh,
	"medium":   ResourcesSeveritiesMedium,
	"low":      ResourcesSeveritiesLow,
	"info":     ResourcesSeveritiesInfo,
	"none":     ResourcesSeveritiesNone,
}

// GetResourcesSeveritiesEnumValues Enumerates the set of values for ResourcesSeveritiesEnum
func GetResourcesSeveritiesEnumValues() []ResourcesSeveritiesEnum {
	values := make([]ResourcesSeveritiesEnum, 0)
	for _, v := range mappingResourcesSeveritiesEnum {
		values = append(values, v)
	}
	return values
}

// GetResourcesSeveritiesEnumStringValues Enumerates the set of values in String for ResourcesSeveritiesEnum
func GetResourcesSeveritiesEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"INFO",
		"NONE",
	}
}

// GetMappingResourcesSeveritiesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourcesSeveritiesEnum(val string) (ResourcesSeveritiesEnum, bool) {
	enum, ok := mappingResourcesSeveritiesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
