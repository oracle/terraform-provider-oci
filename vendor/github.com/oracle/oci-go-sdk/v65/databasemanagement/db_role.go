// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// DbRoleEnum Enum with underlying type: string
type DbRoleEnum string

// Set of constants representing the allowable values for DbRoleEnum
const (
	DbRoleSnapshotStandby DbRoleEnum = "SNAPSHOT_STANDBY"
	DbRoleLogicalStandby  DbRoleEnum = "LOGICAL_STANDBY"
	DbRolePhysicalStandby DbRoleEnum = "PHYSICAL_STANDBY"
	DbRolePrimary         DbRoleEnum = "PRIMARY"
	DbRoleFarSync         DbRoleEnum = "FAR_SYNC"
)

var mappingDbRoleEnum = map[string]DbRoleEnum{
	"SNAPSHOT_STANDBY": DbRoleSnapshotStandby,
	"LOGICAL_STANDBY":  DbRoleLogicalStandby,
	"PHYSICAL_STANDBY": DbRolePhysicalStandby,
	"PRIMARY":          DbRolePrimary,
	"FAR_SYNC":         DbRoleFarSync,
}

var mappingDbRoleEnumLowerCase = map[string]DbRoleEnum{
	"snapshot_standby": DbRoleSnapshotStandby,
	"logical_standby":  DbRoleLogicalStandby,
	"physical_standby": DbRolePhysicalStandby,
	"primary":          DbRolePrimary,
	"far_sync":         DbRoleFarSync,
}

// GetDbRoleEnumValues Enumerates the set of values for DbRoleEnum
func GetDbRoleEnumValues() []DbRoleEnum {
	values := make([]DbRoleEnum, 0)
	for _, v := range mappingDbRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetDbRoleEnumStringValues Enumerates the set of values in String for DbRoleEnum
func GetDbRoleEnumStringValues() []string {
	return []string{
		"SNAPSHOT_STANDBY",
		"LOGICAL_STANDBY",
		"PHYSICAL_STANDBY",
		"PRIMARY",
		"FAR_SYNC",
	}
}

// GetMappingDbRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbRoleEnum(val string) (DbRoleEnum, bool) {
	enum, ok := mappingDbRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
