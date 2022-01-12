// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

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

var mappingDatabaseType = map[string]DatabaseTypeEnum{
	"EXTERNAL_SIDB": DatabaseTypeExternalSidb,
	"EXTERNAL_RAC":  DatabaseTypeExternalRac,
	"CLOUD_SIDB":    DatabaseTypeCloudSidb,
	"CLOUD_RAC":     DatabaseTypeCloudRac,
	"SHARED":        DatabaseTypeShared,
	"DEDICATED":     DatabaseTypeDedicated,
}

// GetDatabaseTypeEnumValues Enumerates the set of values for DatabaseTypeEnum
func GetDatabaseTypeEnumValues() []DatabaseTypeEnum {
	values := make([]DatabaseTypeEnum, 0)
	for _, v := range mappingDatabaseType {
		values = append(values, v)
	}
	return values
}
