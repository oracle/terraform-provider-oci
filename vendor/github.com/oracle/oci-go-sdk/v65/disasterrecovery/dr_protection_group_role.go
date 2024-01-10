// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// DrProtectionGroupRoleEnum Enum with underlying type: string
type DrProtectionGroupRoleEnum string

// Set of constants representing the allowable values for DrProtectionGroupRoleEnum
const (
	DrProtectionGroupRolePrimary      DrProtectionGroupRoleEnum = "PRIMARY"
	DrProtectionGroupRoleStandby      DrProtectionGroupRoleEnum = "STANDBY"
	DrProtectionGroupRoleUnconfigured DrProtectionGroupRoleEnum = "UNCONFIGURED"
)

var mappingDrProtectionGroupRoleEnum = map[string]DrProtectionGroupRoleEnum{
	"PRIMARY":      DrProtectionGroupRolePrimary,
	"STANDBY":      DrProtectionGroupRoleStandby,
	"UNCONFIGURED": DrProtectionGroupRoleUnconfigured,
}

var mappingDrProtectionGroupRoleEnumLowerCase = map[string]DrProtectionGroupRoleEnum{
	"primary":      DrProtectionGroupRolePrimary,
	"standby":      DrProtectionGroupRoleStandby,
	"unconfigured": DrProtectionGroupRoleUnconfigured,
}

// GetDrProtectionGroupRoleEnumValues Enumerates the set of values for DrProtectionGroupRoleEnum
func GetDrProtectionGroupRoleEnumValues() []DrProtectionGroupRoleEnum {
	values := make([]DrProtectionGroupRoleEnum, 0)
	for _, v := range mappingDrProtectionGroupRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDrProtectionGroupRoleEnumStringValues Enumerates the set of values in String for DrProtectionGroupRoleEnum
func GetDrProtectionGroupRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"UNCONFIGURED",
	}
}

// GetMappingDrProtectionGroupRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrProtectionGroupRoleEnum(val string) (DrProtectionGroupRoleEnum, bool) {
	enum, ok := mappingDrProtectionGroupRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
