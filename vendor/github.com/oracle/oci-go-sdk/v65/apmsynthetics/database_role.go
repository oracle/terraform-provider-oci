// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// DatabaseRoleEnum Enum with underlying type: string
type DatabaseRoleEnum string

// Set of constants representing the allowable values for DatabaseRoleEnum
const (
	DatabaseRoleDefault   DatabaseRoleEnum = "DEFAULT"
	DatabaseRoleSysdba    DatabaseRoleEnum = "SYSDBA"
	DatabaseRoleSysoper   DatabaseRoleEnum = "SYSOPER"
	DatabaseRoleSysbackup DatabaseRoleEnum = "SYSBACKUP"
	DatabaseRoleSysdg     DatabaseRoleEnum = "SYSDG"
	DatabaseRoleSyskm     DatabaseRoleEnum = "SYSKM"
	DatabaseRoleSysasm    DatabaseRoleEnum = "SYSASM"
)

var mappingDatabaseRoleEnum = map[string]DatabaseRoleEnum{
	"DEFAULT":   DatabaseRoleDefault,
	"SYSDBA":    DatabaseRoleSysdba,
	"SYSOPER":   DatabaseRoleSysoper,
	"SYSBACKUP": DatabaseRoleSysbackup,
	"SYSDG":     DatabaseRoleSysdg,
	"SYSKM":     DatabaseRoleSyskm,
	"SYSASM":    DatabaseRoleSysasm,
}

var mappingDatabaseRoleEnumLowerCase = map[string]DatabaseRoleEnum{
	"default":   DatabaseRoleDefault,
	"sysdba":    DatabaseRoleSysdba,
	"sysoper":   DatabaseRoleSysoper,
	"sysbackup": DatabaseRoleSysbackup,
	"sysdg":     DatabaseRoleSysdg,
	"syskm":     DatabaseRoleSyskm,
	"sysasm":    DatabaseRoleSysasm,
}

// GetDatabaseRoleEnumValues Enumerates the set of values for DatabaseRoleEnum
func GetDatabaseRoleEnumValues() []DatabaseRoleEnum {
	values := make([]DatabaseRoleEnum, 0)
	for _, v := range mappingDatabaseRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseRoleEnumStringValues Enumerates the set of values in String for DatabaseRoleEnum
func GetDatabaseRoleEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"SYSDBA",
		"SYSOPER",
		"SYSBACKUP",
		"SYSDG",
		"SYSKM",
		"SYSASM",
	}
}

// GetMappingDatabaseRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseRoleEnum(val string) (DatabaseRoleEnum, bool) {
	enum, ok := mappingDatabaseRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
