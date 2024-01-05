// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"strings"
)

// AccessLevelEnum Enum with underlying type: string
type AccessLevelEnum string

// Set of constants representing the allowable values for AccessLevelEnum
const (
	AccessLevelAccessible AccessLevelEnum = "ACCESSIBLE"
	AccessLevelRestricted AccessLevelEnum = "RESTRICTED"
)

var mappingAccessLevelEnum = map[string]AccessLevelEnum{
	"ACCESSIBLE": AccessLevelAccessible,
	"RESTRICTED": AccessLevelRestricted,
}

var mappingAccessLevelEnumLowerCase = map[string]AccessLevelEnum{
	"accessible": AccessLevelAccessible,
	"restricted": AccessLevelRestricted,
}

// GetAccessLevelEnumValues Enumerates the set of values for AccessLevelEnum
func GetAccessLevelEnumValues() []AccessLevelEnum {
	values := make([]AccessLevelEnum, 0)
	for _, v := range mappingAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessLevelEnumStringValues Enumerates the set of values in String for AccessLevelEnum
func GetAccessLevelEnumStringValues() []string {
	return []string{
		"ACCESSIBLE",
		"RESTRICTED",
	}
}

// GetMappingAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessLevelEnum(val string) (AccessLevelEnum, bool) {
	enum, ok := mappingAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
