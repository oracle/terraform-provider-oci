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

// OracleDatabaseObjectTypesEnum Enum with underlying type: string
type OracleDatabaseObjectTypesEnum string

// Set of constants representing the allowable values for OracleDatabaseObjectTypesEnum
const (
	OracleDatabaseObjectTypesTable            OracleDatabaseObjectTypesEnum = "TABLE"
	OracleDatabaseObjectTypesIndex            OracleDatabaseObjectTypesEnum = "INDEX"
	OracleDatabaseObjectTypesView             OracleDatabaseObjectTypesEnum = "VIEW"
	OracleDatabaseObjectTypesSequence         OracleDatabaseObjectTypesEnum = "SEQUENCE"
	OracleDatabaseObjectTypesSynonym          OracleDatabaseObjectTypesEnum = "SYNONYM"
	OracleDatabaseObjectTypesCluster          OracleDatabaseObjectTypesEnum = "CLUSTER"
	OracleDatabaseObjectTypesProcedure        OracleDatabaseObjectTypesEnum = "PROCEDURE"
	OracleDatabaseObjectTypesFunction         OracleDatabaseObjectTypesEnum = "FUNCTION"
	OracleDatabaseObjectTypesPackage          OracleDatabaseObjectTypesEnum = "PACKAGE"
	OracleDatabaseObjectTypesPackageBody      OracleDatabaseObjectTypesEnum = "PACKAGE_BODY"
	OracleDatabaseObjectTypesTrigger          OracleDatabaseObjectTypesEnum = "TRIGGER"
	OracleDatabaseObjectTypesUser             OracleDatabaseObjectTypesEnum = "USER"
	OracleDatabaseObjectTypesRole             OracleDatabaseObjectTypesEnum = "ROLE"
	OracleDatabaseObjectTypesProfile          OracleDatabaseObjectTypesEnum = "PROFILE"
	OracleDatabaseObjectTypesMaterializedView OracleDatabaseObjectTypesEnum = "MATERIALIZED_VIEW"
	OracleDatabaseObjectTypesType             OracleDatabaseObjectTypesEnum = "TYPE"
	OracleDatabaseObjectTypesTypeBody         OracleDatabaseObjectTypesEnum = "TYPE_BODY"
	OracleDatabaseObjectTypesOperator         OracleDatabaseObjectTypesEnum = "OPERATOR"
	OracleDatabaseObjectTypesLibrary          OracleDatabaseObjectTypesEnum = "LIBRARY"
	OracleDatabaseObjectTypesDirectory        OracleDatabaseObjectTypesEnum = "DIRECTORY"
	OracleDatabaseObjectTypesTablespace       OracleDatabaseObjectTypesEnum = "TABLESPACE"
	OracleDatabaseObjectTypesDatabaseLink     OracleDatabaseObjectTypesEnum = "DATABASE_LINK"
	OracleDatabaseObjectTypesControlfile      OracleDatabaseObjectTypesEnum = "CONTROLFILE"
	OracleDatabaseObjectTypesDatafile         OracleDatabaseObjectTypesEnum = "DATAFILE"
	OracleDatabaseObjectTypesRedoLog          OracleDatabaseObjectTypesEnum = "REDO_LOG"
	OracleDatabaseObjectTypesJavaClass        OracleDatabaseObjectTypesEnum = "JAVA_CLASS"
	OracleDatabaseObjectTypesJavaSource       OracleDatabaseObjectTypesEnum = "JAVA_SOURCE"
	OracleDatabaseObjectTypesJavaResource     OracleDatabaseObjectTypesEnum = "JAVA_RESOURCE"
	OracleDatabaseObjectTypesXmlSchema        OracleDatabaseObjectTypesEnum = "XML_SCHEMA"
	OracleDatabaseObjectTypesQueue            OracleDatabaseObjectTypesEnum = "QUEUE"
	OracleDatabaseObjectTypesQueueTable       OracleDatabaseObjectTypesEnum = "QUEUE_TABLE"
	OracleDatabaseObjectTypesDimension        OracleDatabaseObjectTypesEnum = "DIMENSION"
	OracleDatabaseObjectTypesContext          OracleDatabaseObjectTypesEnum = "CONTEXT"
)

var mappingOracleDatabaseObjectTypesEnum = map[string]OracleDatabaseObjectTypesEnum{
	"TABLE":             OracleDatabaseObjectTypesTable,
	"INDEX":             OracleDatabaseObjectTypesIndex,
	"VIEW":              OracleDatabaseObjectTypesView,
	"SEQUENCE":          OracleDatabaseObjectTypesSequence,
	"SYNONYM":           OracleDatabaseObjectTypesSynonym,
	"CLUSTER":           OracleDatabaseObjectTypesCluster,
	"PROCEDURE":         OracleDatabaseObjectTypesProcedure,
	"FUNCTION":          OracleDatabaseObjectTypesFunction,
	"PACKAGE":           OracleDatabaseObjectTypesPackage,
	"PACKAGE_BODY":      OracleDatabaseObjectTypesPackageBody,
	"TRIGGER":           OracleDatabaseObjectTypesTrigger,
	"USER":              OracleDatabaseObjectTypesUser,
	"ROLE":              OracleDatabaseObjectTypesRole,
	"PROFILE":           OracleDatabaseObjectTypesProfile,
	"MATERIALIZED_VIEW": OracleDatabaseObjectTypesMaterializedView,
	"TYPE":              OracleDatabaseObjectTypesType,
	"TYPE_BODY":         OracleDatabaseObjectTypesTypeBody,
	"OPERATOR":          OracleDatabaseObjectTypesOperator,
	"LIBRARY":           OracleDatabaseObjectTypesLibrary,
	"DIRECTORY":         OracleDatabaseObjectTypesDirectory,
	"TABLESPACE":        OracleDatabaseObjectTypesTablespace,
	"DATABASE_LINK":     OracleDatabaseObjectTypesDatabaseLink,
	"CONTROLFILE":       OracleDatabaseObjectTypesControlfile,
	"DATAFILE":          OracleDatabaseObjectTypesDatafile,
	"REDO_LOG":          OracleDatabaseObjectTypesRedoLog,
	"JAVA_CLASS":        OracleDatabaseObjectTypesJavaClass,
	"JAVA_SOURCE":       OracleDatabaseObjectTypesJavaSource,
	"JAVA_RESOURCE":     OracleDatabaseObjectTypesJavaResource,
	"XML_SCHEMA":        OracleDatabaseObjectTypesXmlSchema,
	"QUEUE":             OracleDatabaseObjectTypesQueue,
	"QUEUE_TABLE":       OracleDatabaseObjectTypesQueueTable,
	"DIMENSION":         OracleDatabaseObjectTypesDimension,
	"CONTEXT":           OracleDatabaseObjectTypesContext,
}

var mappingOracleDatabaseObjectTypesEnumLowerCase = map[string]OracleDatabaseObjectTypesEnum{
	"table":             OracleDatabaseObjectTypesTable,
	"index":             OracleDatabaseObjectTypesIndex,
	"view":              OracleDatabaseObjectTypesView,
	"sequence":          OracleDatabaseObjectTypesSequence,
	"synonym":           OracleDatabaseObjectTypesSynonym,
	"cluster":           OracleDatabaseObjectTypesCluster,
	"procedure":         OracleDatabaseObjectTypesProcedure,
	"function":          OracleDatabaseObjectTypesFunction,
	"package":           OracleDatabaseObjectTypesPackage,
	"package_body":      OracleDatabaseObjectTypesPackageBody,
	"trigger":           OracleDatabaseObjectTypesTrigger,
	"user":              OracleDatabaseObjectTypesUser,
	"role":              OracleDatabaseObjectTypesRole,
	"profile":           OracleDatabaseObjectTypesProfile,
	"materialized_view": OracleDatabaseObjectTypesMaterializedView,
	"type":              OracleDatabaseObjectTypesType,
	"type_body":         OracleDatabaseObjectTypesTypeBody,
	"operator":          OracleDatabaseObjectTypesOperator,
	"library":           OracleDatabaseObjectTypesLibrary,
	"directory":         OracleDatabaseObjectTypesDirectory,
	"tablespace":        OracleDatabaseObjectTypesTablespace,
	"database_link":     OracleDatabaseObjectTypesDatabaseLink,
	"controlfile":       OracleDatabaseObjectTypesControlfile,
	"datafile":          OracleDatabaseObjectTypesDatafile,
	"redo_log":          OracleDatabaseObjectTypesRedoLog,
	"java_class":        OracleDatabaseObjectTypesJavaClass,
	"java_source":       OracleDatabaseObjectTypesJavaSource,
	"java_resource":     OracleDatabaseObjectTypesJavaResource,
	"xml_schema":        OracleDatabaseObjectTypesXmlSchema,
	"queue":             OracleDatabaseObjectTypesQueue,
	"queue_table":       OracleDatabaseObjectTypesQueueTable,
	"dimension":         OracleDatabaseObjectTypesDimension,
	"context":           OracleDatabaseObjectTypesContext,
}

// GetOracleDatabaseObjectTypesEnumValues Enumerates the set of values for OracleDatabaseObjectTypesEnum
func GetOracleDatabaseObjectTypesEnumValues() []OracleDatabaseObjectTypesEnum {
	values := make([]OracleDatabaseObjectTypesEnum, 0)
	for _, v := range mappingOracleDatabaseObjectTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDatabaseObjectTypesEnumStringValues Enumerates the set of values in String for OracleDatabaseObjectTypesEnum
func GetOracleDatabaseObjectTypesEnumStringValues() []string {
	return []string{
		"TABLE",
		"INDEX",
		"VIEW",
		"SEQUENCE",
		"SYNONYM",
		"CLUSTER",
		"PROCEDURE",
		"FUNCTION",
		"PACKAGE",
		"PACKAGE_BODY",
		"TRIGGER",
		"USER",
		"ROLE",
		"PROFILE",
		"MATERIALIZED_VIEW",
		"TYPE",
		"TYPE_BODY",
		"OPERATOR",
		"LIBRARY",
		"DIRECTORY",
		"TABLESPACE",
		"DATABASE_LINK",
		"CONTROLFILE",
		"DATAFILE",
		"REDO_LOG",
		"JAVA_CLASS",
		"JAVA_SOURCE",
		"JAVA_RESOURCE",
		"XML_SCHEMA",
		"QUEUE",
		"QUEUE_TABLE",
		"DIMENSION",
		"CONTEXT",
	}
}

// GetMappingOracleDatabaseObjectTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDatabaseObjectTypesEnum(val string) (OracleDatabaseObjectTypesEnum, bool) {
	enum, ok := mappingOracleDatabaseObjectTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
