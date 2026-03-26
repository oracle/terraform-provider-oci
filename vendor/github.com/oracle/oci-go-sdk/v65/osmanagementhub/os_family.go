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

// OsFamilyEnum Enum with underlying type: string
type OsFamilyEnum string

// Set of constants representing the allowable values for OsFamilyEnum
const (
	OsFamilyOracleLinux10     OsFamilyEnum = "ORACLE_LINUX_10"
	OsFamilyOracleLinux9      OsFamilyEnum = "ORACLE_LINUX_9"
	OsFamilyOracleLinux8      OsFamilyEnum = "ORACLE_LINUX_8"
	OsFamilyOracleLinux7      OsFamilyEnum = "ORACLE_LINUX_7"
	OsFamilyOracleLinux6      OsFamilyEnum = "ORACLE_LINUX_6"
	OsFamilyWindowsServer2016 OsFamilyEnum = "WINDOWS_SERVER_2016"
	OsFamilyWindowsServer2019 OsFamilyEnum = "WINDOWS_SERVER_2019"
	OsFamilyWindowsServer2022 OsFamilyEnum = "WINDOWS_SERVER_2022"
	OsFamilyWindowsServer2025 OsFamilyEnum = "WINDOWS_SERVER_2025"
	OsFamilyWindows11         OsFamilyEnum = "WINDOWS_11"
	OsFamilyAll               OsFamilyEnum = "ALL"
	OsFamilyUbuntu2004        OsFamilyEnum = "UBUNTU_20_04"
	OsFamilyUbuntu2204        OsFamilyEnum = "UBUNTU_22_04"
	OsFamilyUbuntu2404        OsFamilyEnum = "UBUNTU_24_04"
)

var mappingOsFamilyEnum = map[string]OsFamilyEnum{
	"ORACLE_LINUX_10":     OsFamilyOracleLinux10,
	"ORACLE_LINUX_9":      OsFamilyOracleLinux9,
	"ORACLE_LINUX_8":      OsFamilyOracleLinux8,
	"ORACLE_LINUX_7":      OsFamilyOracleLinux7,
	"ORACLE_LINUX_6":      OsFamilyOracleLinux6,
	"WINDOWS_SERVER_2016": OsFamilyWindowsServer2016,
	"WINDOWS_SERVER_2019": OsFamilyWindowsServer2019,
	"WINDOWS_SERVER_2022": OsFamilyWindowsServer2022,
	"WINDOWS_SERVER_2025": OsFamilyWindowsServer2025,
	"WINDOWS_11":          OsFamilyWindows11,
	"ALL":                 OsFamilyAll,
	"UBUNTU_20_04":        OsFamilyUbuntu2004,
	"UBUNTU_22_04":        OsFamilyUbuntu2204,
	"UBUNTU_24_04":        OsFamilyUbuntu2404,
}

var mappingOsFamilyEnumLowerCase = map[string]OsFamilyEnum{
	"oracle_linux_10":     OsFamilyOracleLinux10,
	"oracle_linux_9":      OsFamilyOracleLinux9,
	"oracle_linux_8":      OsFamilyOracleLinux8,
	"oracle_linux_7":      OsFamilyOracleLinux7,
	"oracle_linux_6":      OsFamilyOracleLinux6,
	"windows_server_2016": OsFamilyWindowsServer2016,
	"windows_server_2019": OsFamilyWindowsServer2019,
	"windows_server_2022": OsFamilyWindowsServer2022,
	"windows_server_2025": OsFamilyWindowsServer2025,
	"windows_11":          OsFamilyWindows11,
	"all":                 OsFamilyAll,
	"ubuntu_20_04":        OsFamilyUbuntu2004,
	"ubuntu_22_04":        OsFamilyUbuntu2204,
	"ubuntu_24_04":        OsFamilyUbuntu2404,
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
		"ORACLE_LINUX_10",
		"ORACLE_LINUX_9",
		"ORACLE_LINUX_8",
		"ORACLE_LINUX_7",
		"ORACLE_LINUX_6",
		"WINDOWS_SERVER_2016",
		"WINDOWS_SERVER_2019",
		"WINDOWS_SERVER_2022",
		"WINDOWS_SERVER_2025",
		"WINDOWS_11",
		"ALL",
		"UBUNTU_20_04",
		"UBUNTU_22_04",
		"UBUNTU_24_04",
	}
}

// GetMappingOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsFamilyEnum(val string) (OsFamilyEnum, bool) {
	enum, ok := mappingOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
