// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// CloudDbSystemComponentTypeEnum Enum with underlying type: string
type CloudDbSystemComponentTypeEnum string

// Set of constants representing the allowable values for CloudDbSystemComponentTypeEnum
const (
	CloudDbSystemComponentTypeAsm               CloudDbSystemComponentTypeEnum = "ASM"
	CloudDbSystemComponentTypeAsmInstance       CloudDbSystemComponentTypeEnum = "ASM_INSTANCE"
	CloudDbSystemComponentTypeCluster           CloudDbSystemComponentTypeEnum = "CLUSTER"
	CloudDbSystemComponentTypeClusterInstance   CloudDbSystemComponentTypeEnum = "CLUSTER_INSTANCE"
	CloudDbSystemComponentTypeDatabase          CloudDbSystemComponentTypeEnum = "DATABASE"
	CloudDbSystemComponentTypeDatabaseInstance  CloudDbSystemComponentTypeEnum = "DATABASE_INSTANCE"
	CloudDbSystemComponentTypeDatabaseHome      CloudDbSystemComponentTypeEnum = "DATABASE_HOME"
	CloudDbSystemComponentTypeDatabaseNode      CloudDbSystemComponentTypeEnum = "DATABASE_NODE"
	CloudDbSystemComponentTypeDbsystem          CloudDbSystemComponentTypeEnum = "DBSYSTEM"
	CloudDbSystemComponentTypeListener          CloudDbSystemComponentTypeEnum = "LISTENER"
	CloudDbSystemComponentTypePluggableDatabase CloudDbSystemComponentTypeEnum = "PLUGGABLE_DATABASE"
)

var mappingCloudDbSystemComponentTypeEnum = map[string]CloudDbSystemComponentTypeEnum{
	"ASM":                CloudDbSystemComponentTypeAsm,
	"ASM_INSTANCE":       CloudDbSystemComponentTypeAsmInstance,
	"CLUSTER":            CloudDbSystemComponentTypeCluster,
	"CLUSTER_INSTANCE":   CloudDbSystemComponentTypeClusterInstance,
	"DATABASE":           CloudDbSystemComponentTypeDatabase,
	"DATABASE_INSTANCE":  CloudDbSystemComponentTypeDatabaseInstance,
	"DATABASE_HOME":      CloudDbSystemComponentTypeDatabaseHome,
	"DATABASE_NODE":      CloudDbSystemComponentTypeDatabaseNode,
	"DBSYSTEM":           CloudDbSystemComponentTypeDbsystem,
	"LISTENER":           CloudDbSystemComponentTypeListener,
	"PLUGGABLE_DATABASE": CloudDbSystemComponentTypePluggableDatabase,
}

var mappingCloudDbSystemComponentTypeEnumLowerCase = map[string]CloudDbSystemComponentTypeEnum{
	"asm":                CloudDbSystemComponentTypeAsm,
	"asm_instance":       CloudDbSystemComponentTypeAsmInstance,
	"cluster":            CloudDbSystemComponentTypeCluster,
	"cluster_instance":   CloudDbSystemComponentTypeClusterInstance,
	"database":           CloudDbSystemComponentTypeDatabase,
	"database_instance":  CloudDbSystemComponentTypeDatabaseInstance,
	"database_home":      CloudDbSystemComponentTypeDatabaseHome,
	"database_node":      CloudDbSystemComponentTypeDatabaseNode,
	"dbsystem":           CloudDbSystemComponentTypeDbsystem,
	"listener":           CloudDbSystemComponentTypeListener,
	"pluggable_database": CloudDbSystemComponentTypePluggableDatabase,
}

// GetCloudDbSystemComponentTypeEnumValues Enumerates the set of values for CloudDbSystemComponentTypeEnum
func GetCloudDbSystemComponentTypeEnumValues() []CloudDbSystemComponentTypeEnum {
	values := make([]CloudDbSystemComponentTypeEnum, 0)
	for _, v := range mappingCloudDbSystemComponentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudDbSystemComponentTypeEnumStringValues Enumerates the set of values in String for CloudDbSystemComponentTypeEnum
func GetCloudDbSystemComponentTypeEnumStringValues() []string {
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

// GetMappingCloudDbSystemComponentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudDbSystemComponentTypeEnum(val string) (CloudDbSystemComponentTypeEnum, bool) {
	enum, ok := mappingCloudDbSystemComponentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
