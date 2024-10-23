// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// ScopeEnum Enum with underlying type: string
type ScopeEnum string

// Set of constants representing the allowable values for ScopeEnum
const (
	ScopeTaxonomy       ScopeEnum = "TAXONOMY"
	ScopePlatformConfig ScopeEnum = "PLATFORM_CONFIG"
)

var mappingScopeEnum = map[string]ScopeEnum{
	"TAXONOMY":        ScopeTaxonomy,
	"PLATFORM_CONFIG": ScopePlatformConfig,
}

var mappingScopeEnumLowerCase = map[string]ScopeEnum{
	"taxonomy":        ScopeTaxonomy,
	"platform_config": ScopePlatformConfig,
}

// GetScopeEnumValues Enumerates the set of values for ScopeEnum
func GetScopeEnumValues() []ScopeEnum {
	values := make([]ScopeEnum, 0)
	for _, v := range mappingScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetScopeEnumStringValues Enumerates the set of values in String for ScopeEnum
func GetScopeEnumStringValues() []string {
	return []string{
		"TAXONOMY",
		"PLATFORM_CONFIG",
	}
}

// GetMappingScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScopeEnum(val string) (ScopeEnum, bool) {
	enum, ok := mappingScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
