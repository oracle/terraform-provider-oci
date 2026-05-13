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

// DatabaseToolsMcpToolsetVersionTypeEnum Enum with underlying type: string
type DatabaseToolsMcpToolsetVersionTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpToolsetVersionTypeEnum
const (
	DatabaseToolsMcpToolsetVersionTypeCustomSqlTool              DatabaseToolsMcpToolsetVersionTypeEnum = "CUSTOM_SQL_TOOL"
	DatabaseToolsMcpToolsetVersionTypeBuiltInSqlTools            DatabaseToolsMcpToolsetVersionTypeEnum = "BUILT_IN_SQL_TOOLS"
	DatabaseToolsMcpToolsetVersionTypeCustomizableReportingTools DatabaseToolsMcpToolsetVersionTypeEnum = "CUSTOMIZABLE_REPORTING_TOOLS"
	DatabaseToolsMcpToolsetVersionTypeGenaiSqlAssistant          DatabaseToolsMcpToolsetVersionTypeEnum = "GENAI_SQL_ASSISTANT"
)

var mappingDatabaseToolsMcpToolsetVersionTypeEnum = map[string]DatabaseToolsMcpToolsetVersionTypeEnum{
	"CUSTOM_SQL_TOOL":              DatabaseToolsMcpToolsetVersionTypeCustomSqlTool,
	"BUILT_IN_SQL_TOOLS":           DatabaseToolsMcpToolsetVersionTypeBuiltInSqlTools,
	"CUSTOMIZABLE_REPORTING_TOOLS": DatabaseToolsMcpToolsetVersionTypeCustomizableReportingTools,
	"GENAI_SQL_ASSISTANT":          DatabaseToolsMcpToolsetVersionTypeGenaiSqlAssistant,
}

var mappingDatabaseToolsMcpToolsetVersionTypeEnumLowerCase = map[string]DatabaseToolsMcpToolsetVersionTypeEnum{
	"custom_sql_tool":              DatabaseToolsMcpToolsetVersionTypeCustomSqlTool,
	"built_in_sql_tools":           DatabaseToolsMcpToolsetVersionTypeBuiltInSqlTools,
	"customizable_reporting_tools": DatabaseToolsMcpToolsetVersionTypeCustomizableReportingTools,
	"genai_sql_assistant":          DatabaseToolsMcpToolsetVersionTypeGenaiSqlAssistant,
}

// GetDatabaseToolsMcpToolsetVersionTypeEnumValues Enumerates the set of values for DatabaseToolsMcpToolsetVersionTypeEnum
func GetDatabaseToolsMcpToolsetVersionTypeEnumValues() []DatabaseToolsMcpToolsetVersionTypeEnum {
	values := make([]DatabaseToolsMcpToolsetVersionTypeEnum, 0)
	for _, v := range mappingDatabaseToolsMcpToolsetVersionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpToolsetVersionTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpToolsetVersionTypeEnum
func GetDatabaseToolsMcpToolsetVersionTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_SQL_TOOL",
		"BUILT_IN_SQL_TOOLS",
		"CUSTOMIZABLE_REPORTING_TOOLS",
		"GENAI_SQL_ASSISTANT",
	}
}

// GetMappingDatabaseToolsMcpToolsetVersionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpToolsetVersionTypeEnum(val string) (DatabaseToolsMcpToolsetVersionTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpToolsetVersionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
