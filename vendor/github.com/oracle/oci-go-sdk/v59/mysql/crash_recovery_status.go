// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// CrashRecoveryStatusEnum Enum with underlying type: string
type CrashRecoveryStatusEnum string

// Set of constants representing the allowable values for CrashRecoveryStatusEnum
const (
	CrashRecoveryStatusEnabled  CrashRecoveryStatusEnum = "ENABLED"
	CrashRecoveryStatusDisabled CrashRecoveryStatusEnum = "DISABLED"
)

var mappingCrashRecoveryStatusEnum = map[string]CrashRecoveryStatusEnum{
	"ENABLED":  CrashRecoveryStatusEnabled,
	"DISABLED": CrashRecoveryStatusDisabled,
}

var mappingCrashRecoveryStatusEnumLowerCase = map[string]CrashRecoveryStatusEnum{
	"enabled":  CrashRecoveryStatusEnabled,
	"disabled": CrashRecoveryStatusDisabled,
}

// GetCrashRecoveryStatusEnumValues Enumerates the set of values for CrashRecoveryStatusEnum
func GetCrashRecoveryStatusEnumValues() []CrashRecoveryStatusEnum {
	values := make([]CrashRecoveryStatusEnum, 0)
	for _, v := range mappingCrashRecoveryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCrashRecoveryStatusEnumStringValues Enumerates the set of values in String for CrashRecoveryStatusEnum
func GetCrashRecoveryStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCrashRecoveryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCrashRecoveryStatusEnum(val string) (CrashRecoveryStatusEnum, bool) {
	enum, ok := mappingCrashRecoveryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
