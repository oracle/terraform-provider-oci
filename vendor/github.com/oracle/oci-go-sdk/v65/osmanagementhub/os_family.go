// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// OsFamilyEnum Enum with underlying type: string
type OsFamilyEnum string

// Set of constants representing the allowable values for OsFamilyEnum
const (
	OsFamilyOracleLinux9 OsFamilyEnum = "ORACLE_LINUX_9"
	OsFamilyOracleLinux8 OsFamilyEnum = "ORACLE_LINUX_8"
	OsFamilyOracleLinux7 OsFamilyEnum = "ORACLE_LINUX_7"
)

var mappingOsFamilyEnum = map[string]OsFamilyEnum{
	"ORACLE_LINUX_9": OsFamilyOracleLinux9,
	"ORACLE_LINUX_8": OsFamilyOracleLinux8,
	"ORACLE_LINUX_7": OsFamilyOracleLinux7,
}

var mappingOsFamilyEnumLowerCase = map[string]OsFamilyEnum{
	"oracle_linux_9": OsFamilyOracleLinux9,
	"oracle_linux_8": OsFamilyOracleLinux8,
	"oracle_linux_7": OsFamilyOracleLinux7,
}

// GetOsFamilyEnumValues Enumerates the set of values for OsFamilyEnum
func GetOsFamilyEnumValues() []OsFamilyEnum {
	values := make([]OsFamilyEnum, 0)
	for _, v := range mappingOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetOsFamilyEnumStringValues Enumerates the set of values in String for OsFamilyEnum
func GetOsFamilyEnumStringValues() []string {
	return []string{
		"ORACLE_LINUX_9",
		"ORACLE_LINUX_8",
		"ORACLE_LINUX_7",
	}
}

// GetMappingOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsFamilyEnum(val string) (OsFamilyEnum, bool) {
	enum, ok := mappingOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
