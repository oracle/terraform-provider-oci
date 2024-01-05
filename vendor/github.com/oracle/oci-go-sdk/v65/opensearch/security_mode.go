// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// SecurityModeEnum Enum with underlying type: string
type SecurityModeEnum string

// Set of constants representing the allowable values for SecurityModeEnum
const (
	SecurityModeDisabled   SecurityModeEnum = "DISABLED"
	SecurityModePermissive SecurityModeEnum = "PERMISSIVE"
	SecurityModeEnforcing  SecurityModeEnum = "ENFORCING"
)

var mappingSecurityModeEnum = map[string]SecurityModeEnum{
	"DISABLED":   SecurityModeDisabled,
	"PERMISSIVE": SecurityModePermissive,
	"ENFORCING":  SecurityModeEnforcing,
}

var mappingSecurityModeEnumLowerCase = map[string]SecurityModeEnum{
	"disabled":   SecurityModeDisabled,
	"permissive": SecurityModePermissive,
	"enforcing":  SecurityModeEnforcing,
}

// GetSecurityModeEnumValues Enumerates the set of values for SecurityModeEnum
func GetSecurityModeEnumValues() []SecurityModeEnum {
	values := make([]SecurityModeEnum, 0)
	for _, v := range mappingSecurityModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityModeEnumStringValues Enumerates the set of values in String for SecurityModeEnum
func GetSecurityModeEnumStringValues() []string {
	return []string{
		"DISABLED",
		"PERMISSIVE",
		"ENFORCING",
	}
}

// GetMappingSecurityModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityModeEnum(val string) (SecurityModeEnum, bool) {
	enum, ok := mappingSecurityModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
