// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

// DatabaseTypeEnum Enum with underlying type: string
type DatabaseTypeEnum string

// Set of constants representing the allowable values for DatabaseTypeEnum
const (
	DatabaseTypeDatabaseCloudService DatabaseTypeEnum = "DATABASE_CLOUD_SERVICE"
	DatabaseTypeAutonomousDatabase   DatabaseTypeEnum = "AUTONOMOUS_DATABASE"
	DatabaseTypeInstalledDatabase    DatabaseTypeEnum = "INSTALLED_DATABASE"
)

var mappingDatabaseType = map[string]DatabaseTypeEnum{
	"DATABASE_CLOUD_SERVICE": DatabaseTypeDatabaseCloudService,
	"AUTONOMOUS_DATABASE":    DatabaseTypeAutonomousDatabase,
	"INSTALLED_DATABASE":     DatabaseTypeInstalledDatabase,
}

// GetDatabaseTypeEnumValues Enumerates the set of values for DatabaseTypeEnum
func GetDatabaseTypeEnumValues() []DatabaseTypeEnum {
	values := make([]DatabaseTypeEnum, 0)
	for _, v := range mappingDatabaseType {
		values = append(values, v)
	}
	return values
}
