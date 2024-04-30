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

// ExternalDbSystemComponentTypeEnum Enum with underlying type: string
type ExternalDbSystemComponentTypeEnum string

// Set of constants representing the allowable values for ExternalDbSystemComponentTypeEnum
const (
	ExternalDbSystemComponentTypeAsm               ExternalDbSystemComponentTypeEnum = "ASM"
	ExternalDbSystemComponentTypeAsmInstance       ExternalDbSystemComponentTypeEnum = "ASM_INSTANCE"
	ExternalDbSystemComponentTypeCluster           ExternalDbSystemComponentTypeEnum = "CLUSTER"
	ExternalDbSystemComponentTypeClusterInstance   ExternalDbSystemComponentTypeEnum = "CLUSTER_INSTANCE"
	ExternalDbSystemComponentTypeDatabase          ExternalDbSystemComponentTypeEnum = "DATABASE"
	ExternalDbSystemComponentTypeDatabaseInstance  ExternalDbSystemComponentTypeEnum = "DATABASE_INSTANCE"
	ExternalDbSystemComponentTypeDatabaseHome      ExternalDbSystemComponentTypeEnum = "DATABASE_HOME"
	ExternalDbSystemComponentTypeDatabaseNode      ExternalDbSystemComponentTypeEnum = "DATABASE_NODE"
	ExternalDbSystemComponentTypeDbsystem          ExternalDbSystemComponentTypeEnum = "DBSYSTEM"
	ExternalDbSystemComponentTypeListener          ExternalDbSystemComponentTypeEnum = "LISTENER"
	ExternalDbSystemComponentTypePluggableDatabase ExternalDbSystemComponentTypeEnum = "PLUGGABLE_DATABASE"
)

var mappingExternalDbSystemComponentTypeEnum = map[string]ExternalDbSystemComponentTypeEnum{
	"ASM":                ExternalDbSystemComponentTypeAsm,
	"ASM_INSTANCE":       ExternalDbSystemComponentTypeAsmInstance,
	"CLUSTER":            ExternalDbSystemComponentTypeCluster,
	"CLUSTER_INSTANCE":   ExternalDbSystemComponentTypeClusterInstance,
	"DATABASE":           ExternalDbSystemComponentTypeDatabase,
	"DATABASE_INSTANCE":  ExternalDbSystemComponentTypeDatabaseInstance,
	"DATABASE_HOME":      ExternalDbSystemComponentTypeDatabaseHome,
	"DATABASE_NODE":      ExternalDbSystemComponentTypeDatabaseNode,
	"DBSYSTEM":           ExternalDbSystemComponentTypeDbsystem,
	"LISTENER":           ExternalDbSystemComponentTypeListener,
	"PLUGGABLE_DATABASE": ExternalDbSystemComponentTypePluggableDatabase,
}

var mappingExternalDbSystemComponentTypeEnumLowerCase = map[string]ExternalDbSystemComponentTypeEnum{
	"asm":                ExternalDbSystemComponentTypeAsm,
	"asm_instance":       ExternalDbSystemComponentTypeAsmInstance,
	"cluster":            ExternalDbSystemComponentTypeCluster,
	"cluster_instance":   ExternalDbSystemComponentTypeClusterInstance,
	"database":           ExternalDbSystemComponentTypeDatabase,
	"database_instance":  ExternalDbSystemComponentTypeDatabaseInstance,
	"database_home":      ExternalDbSystemComponentTypeDatabaseHome,
	"database_node":      ExternalDbSystemComponentTypeDatabaseNode,
	"dbsystem":           ExternalDbSystemComponentTypeDbsystem,
	"listener":           ExternalDbSystemComponentTypeListener,
	"pluggable_database": ExternalDbSystemComponentTypePluggableDatabase,
}

// GetExternalDbSystemComponentTypeEnumValues Enumerates the set of values for ExternalDbSystemComponentTypeEnum
func GetExternalDbSystemComponentTypeEnumValues() []ExternalDbSystemComponentTypeEnum {
	values := make([]ExternalDbSystemComponentTypeEnum, 0)
	for _, v := range mappingExternalDbSystemComponentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemComponentTypeEnumStringValues Enumerates the set of values in String for ExternalDbSystemComponentTypeEnum
func GetExternalDbSystemComponentTypeEnumStringValues() []string {
	return []string{
		"ASM",
		"ASM_INSTANCE",
		"CLUSTER",
		"CLUSTER_INSTANCE",
		"DATABASE",
		"DATABASE_INSTANCE",
		"DATABASE_HOME",
		"DATABASE_NODE",
		"DBSYSTEM",
		"LISTENER",
		"PLUGGABLE_DATABASE",
	}
}

// GetMappingExternalDbSystemComponentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemComponentTypeEnum(val string) (ExternalDbSystemComponentTypeEnum, bool) {
	enum, ok := mappingExternalDbSystemComponentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
