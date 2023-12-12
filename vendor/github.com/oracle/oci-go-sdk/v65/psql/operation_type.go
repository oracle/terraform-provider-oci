// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreatePostgresqlDbSystem       OperationTypeEnum = "CREATE_POSTGRESQL_DB_SYSTEM"
	OperationTypeUpdatePostgresqlDbSystem       OperationTypeEnum = "UPDATE_POSTGRESQL_DB_SYSTEM"
	OperationTypeDeletePostgresqlDbSystem       OperationTypeEnum = "DELETE_POSTGRESQL_DB_SYSTEM"
	OperationTypeMovePostgresqlDbSystem         OperationTypeEnum = "MOVE_POSTGRESQL_DB_SYSTEM"
	OperationTypeCreatePostgresqlDbSystemBackup OperationTypeEnum = "CREATE_POSTGRESQL_DB_SYSTEM_BACKUP"
	OperationTypeUpdatePostgresqlDbSystemBackup OperationTypeEnum = "UPDATE_POSTGRESQL_DB_SYSTEM_BACKUP"
	OperationTypeDeletePostgresqlDbSystemBackup OperationTypeEnum = "DELETE_POSTGRESQL_DB_SYSTEM_BACKUP"
	OperationTypeMovePostgresqlDbSystemBackup   OperationTypeEnum = "MOVE_POSTGRESQL_DB_SYSTEM_BACKUP"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_POSTGRESQL_DB_SYSTEM":        OperationTypeCreatePostgresqlDbSystem,
	"UPDATE_POSTGRESQL_DB_SYSTEM":        OperationTypeUpdatePostgresqlDbSystem,
	"DELETE_POSTGRESQL_DB_SYSTEM":        OperationTypeDeletePostgresqlDbSystem,
	"MOVE_POSTGRESQL_DB_SYSTEM":          OperationTypeMovePostgresqlDbSystem,
	"CREATE_POSTGRESQL_DB_SYSTEM_BACKUP": OperationTypeCreatePostgresqlDbSystemBackup,
	"UPDATE_POSTGRESQL_DB_SYSTEM_BACKUP": OperationTypeUpdatePostgresqlDbSystemBackup,
	"DELETE_POSTGRESQL_DB_SYSTEM_BACKUP": OperationTypeDeletePostgresqlDbSystemBackup,
	"MOVE_POSTGRESQL_DB_SYSTEM_BACKUP":   OperationTypeMovePostgresqlDbSystemBackup,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_postgresql_db_system":        OperationTypeCreatePostgresqlDbSystem,
	"update_postgresql_db_system":        OperationTypeUpdatePostgresqlDbSystem,
	"delete_postgresql_db_system":        OperationTypeDeletePostgresqlDbSystem,
	"move_postgresql_db_system":          OperationTypeMovePostgresqlDbSystem,
	"create_postgresql_db_system_backup": OperationTypeCreatePostgresqlDbSystemBackup,
	"update_postgresql_db_system_backup": OperationTypeUpdatePostgresqlDbSystemBackup,
	"delete_postgresql_db_system_backup": OperationTypeDeletePostgresqlDbSystemBackup,
	"move_postgresql_db_system_backup":   OperationTypeMovePostgresqlDbSystemBackup,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_POSTGRESQL_DB_SYSTEM",
		"UPDATE_POSTGRESQL_DB_SYSTEM",
		"DELETE_POSTGRESQL_DB_SYSTEM",
		"MOVE_POSTGRESQL_DB_SYSTEM",
		"CREATE_POSTGRESQL_DB_SYSTEM_BACKUP",
		"UPDATE_POSTGRESQL_DB_SYSTEM_BACKUP",
		"DELETE_POSTGRESQL_DB_SYSTEM_BACKUP",
		"MOVE_POSTGRESQL_DB_SYSTEM_BACKUP",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
