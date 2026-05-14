// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsMcpToolsetTypeEnum Enum with underlying type: string
type DatabaseToolsMcpToolsetTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpToolsetTypeEnum
const (
	DatabaseToolsMcpToolsetTypeCustomSqlTool              DatabaseToolsMcpToolsetTypeEnum = "CUSTOM_SQL_TOOL"
	DatabaseToolsMcpToolsetTypeBuiltInSqlTools            DatabaseToolsMcpToolsetTypeEnum = "BUILT_IN_SQL_TOOLS"
	DatabaseToolsMcpToolsetTypeCustomizableReportingTools DatabaseToolsMcpToolsetTypeEnum = "CUSTOMIZABLE_REPORTING_TOOLS"
	DatabaseToolsMcpToolsetTypeGenaiSqlAssistant          DatabaseToolsMcpToolsetTypeEnum = "GENAI_SQL_ASSISTANT"
)

var mappingDatabaseToolsMcpToolsetTypeEnum = map[string]DatabaseToolsMcpToolsetTypeEnum{
	"CUSTOM_SQL_TOOL":              DatabaseToolsMcpToolsetTypeCustomSqlTool,
	"BUILT_IN_SQL_TOOLS":           DatabaseToolsMcpToolsetTypeBuiltInSqlTools,
	"CUSTOMIZABLE_REPORTING_TOOLS": DatabaseToolsMcpToolsetTypeCustomizableReportingTools,
	"GENAI_SQL_ASSISTANT":          DatabaseToolsMcpToolsetTypeGenaiSqlAssistant,
}

var mappingDatabaseToolsMcpToolsetTypeEnumLowerCase = map[string]DatabaseToolsMcpToolsetTypeEnum{
	"custom_sql_tool":              DatabaseToolsMcpToolsetTypeCustomSqlTool,
	"built_in_sql_tools":           DatabaseToolsMcpToolsetTypeBuiltInSqlTools,
	"customizable_reporting_tools": DatabaseToolsMcpToolsetTypeCustomizableReportingTools,
	"genai_sql_assistant":          DatabaseToolsMcpToolsetTypeGenaiSqlAssistant,
}

// GetDatabaseToolsMcpToolsetTypeEnumValues Enumerates the set of values for DatabaseToolsMcpToolsetTypeEnum
func GetDatabaseToolsMcpToolsetTypeEnumValues() []DatabaseToolsMcpToolsetTypeEnum {
	values := make([]DatabaseToolsMcpToolsetTypeEnum, 0)
	for _, v := range mappingDatabaseToolsMcpToolsetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpToolsetTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpToolsetTypeEnum
func GetDatabaseToolsMcpToolsetTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_SQL_TOOL",
		"BUILT_IN_SQL_TOOLS",
		"CUSTOMIZABLE_REPORTING_TOOLS",
		"GENAI_SQL_ASSISTANT",
	}
}

// GetMappingDatabaseToolsMcpToolsetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpToolsetTypeEnum(val string) (DatabaseToolsMcpToolsetTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpToolsetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
