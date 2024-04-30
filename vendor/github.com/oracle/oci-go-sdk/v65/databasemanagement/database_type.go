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

// DatabaseTypeEnum Enum with underlying type: string
type DatabaseTypeEnum string

// Set of constants representing the allowable values for DatabaseTypeEnum
const (
	DatabaseTypeExternalSidb DatabaseTypeEnum = "EXTERNAL_SIDB"
	DatabaseTypeExternalRac  DatabaseTypeEnum = "EXTERNAL_RAC"
	DatabaseTypeCloudSidb    DatabaseTypeEnum = "CLOUD_SIDB"
	DatabaseTypeCloudRac     DatabaseTypeEnum = "CLOUD_RAC"
	DatabaseTypeShared       DatabaseTypeEnum = "SHARED"
	DatabaseTypeDedicated    DatabaseTypeEnum = "DEDICATED"
)

var mappingDatabaseTypeEnum = map[string]DatabaseTypeEnum{
	"EXTERNAL_SIDB": DatabaseTypeExternalSidb,
	"EXTERNAL_RAC":  DatabaseTypeExternalRac,
	"CLOUD_SIDB":    DatabaseTypeCloudSidb,
	"CLOUD_RAC":     DatabaseTypeCloudRac,
	"SHARED":        DatabaseTypeShared,
	"DEDICATED":     DatabaseTypeDedicated,
}

var mappingDatabaseTypeEnumLowerCase = map[string]DatabaseTypeEnum{
	"external_sidb": DatabaseTypeExternalSidb,
	"external_rac":  DatabaseTypeExternalRac,
	"cloud_sidb":    DatabaseTypeCloudSidb,
	"cloud_rac":     DatabaseTypeCloudRac,
	"shared":        DatabaseTypeShared,
	"dedicated":     DatabaseTypeDedicated,
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
		"EXTERNAL_SIDB",
		"EXTERNAL_RAC",
		"CLOUD_SIDB",
		"CLOUD_RAC",
		"SHARED",
		"DEDICATED",
	}
}

// GetMappingDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseTypeEnum(val string) (DatabaseTypeEnum, bool) {
	enum, ok := mappingDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
