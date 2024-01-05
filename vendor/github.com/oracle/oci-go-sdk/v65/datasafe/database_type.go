// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// DatabaseTypeEnum Enum with underlying type: string
type DatabaseTypeEnum string

// Set of constants representing the allowable values for DatabaseTypeEnum
const (
	DatabaseTypeDatabaseCloudService DatabaseTypeEnum = "DATABASE_CLOUD_SERVICE"
	DatabaseTypeAutonomousDatabase   DatabaseTypeEnum = "AUTONOMOUS_DATABASE"
	DatabaseTypeInstalledDatabase    DatabaseTypeEnum = "INSTALLED_DATABASE"
)

var mappingDatabaseTypeEnum = map[string]DatabaseTypeEnum{
	"DATABASE_CLOUD_SERVICE": DatabaseTypeDatabaseCloudService,
	"AUTONOMOUS_DATABASE":    DatabaseTypeAutonomousDatabase,
	"INSTALLED_DATABASE":     DatabaseTypeInstalledDatabase,
}

var mappingDatabaseTypeEnumLowerCase = map[string]DatabaseTypeEnum{
	"database_cloud_service": DatabaseTypeDatabaseCloudService,
	"autonomous_database":    DatabaseTypeAutonomousDatabase,
	"installed_database":     DatabaseTypeInstalledDatabase,
}

// GetDatabaseTypeEnumValues Enumerates the set of values for DatabaseTypeEnum
func GetDatabaseTypeEnumValues() []DatabaseTypeEnum {
	values := make([]DatabaseTypeEnum, 0)
	for _, v := range mappingDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseTypeEnumStringValues Enumerates the set of values in String for DatabaseTypeEnum
func GetDatabaseTypeEnumStringValues() []string {
	return []string{
		"DATABASE_CLOUD_SERVICE",
		"AUTONOMOUS_DATABASE",
		"INSTALLED_DATABASE",
	}
}

// GetMappingDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseTypeEnum(val string) (DatabaseTypeEnum, bool) {
	enum, ok := mappingDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
