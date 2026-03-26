// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// OsNameEnum Enum with underlying type: string
type OsNameEnum string

// Set of constants representing the allowable values for OsNameEnum
const (
	OsNameOracleLinux   OsNameEnum = "ORACLE_LINUX"
	OsNameUbuntu        OsNameEnum = "UBUNTU"
	OsNameWindowsServer OsNameEnum = "WINDOWS_SERVER"
)

var mappingOsNameEnum = map[string]OsNameEnum{
	"ORACLE_LINUX":   OsNameOracleLinux,
	"UBUNTU":         OsNameUbuntu,
	"WINDOWS_SERVER": OsNameWindowsServer,
}

var mappingOsNameEnumLowerCase = map[string]OsNameEnum{
	"oracle_linux":   OsNameOracleLinux,
	"ubuntu":         OsNameUbuntu,
	"windows_server": OsNameWindowsServer,
}

// GetOsNameEnumValues Enumerates the set of values for OsNameEnum
func GetOsNameEnumValues() []OsNameEnum {
	values := make([]OsNameEnum, 0)
	for _, v := range mappingOsNameEnum {
		values = append(values, v)
	}
	return values
}

// GetOsNameEnumStringValues Enumerates the set of values in String for OsNameEnum
func GetOsNameEnumStringValues() []string {
	return []string{
		"ORACLE_LINUX",
		"UBUNTU",
		"WINDOWS_SERVER",
	}
}

// GetMappingOsNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsNameEnum(val string) (OsNameEnum, bool) {
	enum, ok := mappingOsNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
