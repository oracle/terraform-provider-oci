// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// OracleMigrationMethodsEnum Enum with underlying type: string
type OracleMigrationMethodsEnum string

// Set of constants representing the allowable values for OracleMigrationMethodsEnum
const (
	OracleMigrationMethodsOnlineLogical  OracleMigrationMethodsEnum = "ONLINE_LOGICAL"
	OracleMigrationMethodsOfflineLogical OracleMigrationMethodsEnum = "OFFLINE_LOGICAL"
	OracleMigrationMethodsOnlineStandby  OracleMigrationMethodsEnum = "ONLINE_STANDBY"
	OracleMigrationMethodsOnlinePhysical OracleMigrationMethodsEnum = "ONLINE_PHYSICAL"
)

var mappingOracleMigrationMethodsEnum = map[string]OracleMigrationMethodsEnum{
	"ONLINE_LOGICAL":  OracleMigrationMethodsOnlineLogical,
	"OFFLINE_LOGICAL": OracleMigrationMethodsOfflineLogical,
	"ONLINE_STANDBY":  OracleMigrationMethodsOnlineStandby,
	"ONLINE_PHYSICAL": OracleMigrationMethodsOnlinePhysical,
}

var mappingOracleMigrationMethodsEnumLowerCase = map[string]OracleMigrationMethodsEnum{
	"online_logical":  OracleMigrationMethodsOnlineLogical,
	"offline_logical": OracleMigrationMethodsOfflineLogical,
	"online_standby":  OracleMigrationMethodsOnlineStandby,
	"online_physical": OracleMigrationMethodsOnlinePhysical,
}

// GetOracleMigrationMethodsEnumValues Enumerates the set of values for OracleMigrationMethodsEnum
func GetOracleMigrationMethodsEnumValues() []OracleMigrationMethodsEnum {
	values := make([]OracleMigrationMethodsEnum, 0)
	for _, v := range mappingOracleMigrationMethodsEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleMigrationMethodsEnumStringValues Enumerates the set of values in String for OracleMigrationMethodsEnum
func GetOracleMigrationMethodsEnumStringValues() []string {
	return []string{
		"ONLINE_LOGICAL",
		"OFFLINE_LOGICAL",
		"ONLINE_STANDBY",
		"ONLINE_PHYSICAL",
	}
}

// GetMappingOracleMigrationMethodsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleMigrationMethodsEnum(val string) (OracleMigrationMethodsEnum, bool) {
	enum, ok := mappingOracleMigrationMethodsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
