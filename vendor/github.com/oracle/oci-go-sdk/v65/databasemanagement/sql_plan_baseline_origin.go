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

// SqlPlanBaselineOriginEnum Enum with underlying type: string
type SqlPlanBaselineOriginEnum string

// Set of constants representing the allowable values for SqlPlanBaselineOriginEnum
const (
	SqlPlanBaselineOriginAddmSqltune               SqlPlanBaselineOriginEnum = "ADDM_SQLTUNE"
	SqlPlanBaselineOriginAutoCapture               SqlPlanBaselineOriginEnum = "AUTO_CAPTURE"
	SqlPlanBaselineOriginAutoSqltune               SqlPlanBaselineOriginEnum = "AUTO_SQLTUNE"
	SqlPlanBaselineOriginEvolveAutoIndexLoad       SqlPlanBaselineOriginEnum = "EVOLVE_AUTO_INDEX_LOAD"
	SqlPlanBaselineOriginEvolveCreateFromAdaptive  SqlPlanBaselineOriginEnum = "EVOLVE_CREATE_FROM_ADAPTIVE"
	SqlPlanBaselineOriginEvolveLoadFromSts         SqlPlanBaselineOriginEnum = "EVOLVE_LOAD_FROM_STS"
	SqlPlanBaselineOriginEvolveLoadFromAwr         SqlPlanBaselineOriginEnum = "EVOLVE_LOAD_FROM_AWR"
	SqlPlanBaselineOriginEvolveLoadFromCursorCache SqlPlanBaselineOriginEnum = "EVOLVE_LOAD_FROM_CURSOR_CACHE"
	SqlPlanBaselineOriginManualLoad                SqlPlanBaselineOriginEnum = "MANUAL_LOAD"
	SqlPlanBaselineOriginManualLoadFromAwr         SqlPlanBaselineOriginEnum = "MANUAL_LOAD_FROM_AWR"
	SqlPlanBaselineOriginManualLoadFromCursorCache SqlPlanBaselineOriginEnum = "MANUAL_LOAD_FROM_CURSOR_CACHE"
	SqlPlanBaselineOriginManualLoadFromSts         SqlPlanBaselineOriginEnum = "MANUAL_LOAD_FROM_STS"
	SqlPlanBaselineOriginManualSqltune             SqlPlanBaselineOriginEnum = "MANUAL_SQLTUNE"
	SqlPlanBaselineOriginStoredOutline             SqlPlanBaselineOriginEnum = "STORED_OUTLINE"
	SqlPlanBaselineOriginUnknown                   SqlPlanBaselineOriginEnum = "UNKNOWN"
)

var mappingSqlPlanBaselineOriginEnum = map[string]SqlPlanBaselineOriginEnum{
	"ADDM_SQLTUNE":                  SqlPlanBaselineOriginAddmSqltune,
	"AUTO_CAPTURE":                  SqlPlanBaselineOriginAutoCapture,
	"AUTO_SQLTUNE":                  SqlPlanBaselineOriginAutoSqltune,
	"EVOLVE_AUTO_INDEX_LOAD":        SqlPlanBaselineOriginEvolveAutoIndexLoad,
	"EVOLVE_CREATE_FROM_ADAPTIVE":   SqlPlanBaselineOriginEvolveCreateFromAdaptive,
	"EVOLVE_LOAD_FROM_STS":          SqlPlanBaselineOriginEvolveLoadFromSts,
	"EVOLVE_LOAD_FROM_AWR":          SqlPlanBaselineOriginEvolveLoadFromAwr,
	"EVOLVE_LOAD_FROM_CURSOR_CACHE": SqlPlanBaselineOriginEvolveLoadFromCursorCache,
	"MANUAL_LOAD":                   SqlPlanBaselineOriginManualLoad,
	"MANUAL_LOAD_FROM_AWR":          SqlPlanBaselineOriginManualLoadFromAwr,
	"MANUAL_LOAD_FROM_CURSOR_CACHE": SqlPlanBaselineOriginManualLoadFromCursorCache,
	"MANUAL_LOAD_FROM_STS":          SqlPlanBaselineOriginManualLoadFromSts,
	"MANUAL_SQLTUNE":                SqlPlanBaselineOriginManualSqltune,
	"STORED_OUTLINE":                SqlPlanBaselineOriginStoredOutline,
	"UNKNOWN":                       SqlPlanBaselineOriginUnknown,
}

var mappingSqlPlanBaselineOriginEnumLowerCase = map[string]SqlPlanBaselineOriginEnum{
	"addm_sqltune":                  SqlPlanBaselineOriginAddmSqltune,
	"auto_capture":                  SqlPlanBaselineOriginAutoCapture,
	"auto_sqltune":                  SqlPlanBaselineOriginAutoSqltune,
	"evolve_auto_index_load":        SqlPlanBaselineOriginEvolveAutoIndexLoad,
	"evolve_create_from_adaptive":   SqlPlanBaselineOriginEvolveCreateFromAdaptive,
	"evolve_load_from_sts":          SqlPlanBaselineOriginEvolveLoadFromSts,
	"evolve_load_from_awr":          SqlPlanBaselineOriginEvolveLoadFromAwr,
	"evolve_load_from_cursor_cache": SqlPlanBaselineOriginEvolveLoadFromCursorCache,
	"manual_load":                   SqlPlanBaselineOriginManualLoad,
	"manual_load_from_awr":          SqlPlanBaselineOriginManualLoadFromAwr,
	"manual_load_from_cursor_cache": SqlPlanBaselineOriginManualLoadFromCursorCache,
	"manual_load_from_sts":          SqlPlanBaselineOriginManualLoadFromSts,
	"manual_sqltune":                SqlPlanBaselineOriginManualSqltune,
	"stored_outline":                SqlPlanBaselineOriginStoredOutline,
	"unknown":                       SqlPlanBaselineOriginUnknown,
}

// GetSqlPlanBaselineOriginEnumValues Enumerates the set of values for SqlPlanBaselineOriginEnum
func GetSqlPlanBaselineOriginEnumValues() []SqlPlanBaselineOriginEnum {
	values := make([]SqlPlanBaselineOriginEnum, 0)
	for _, v := range mappingSqlPlanBaselineOriginEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineOriginEnumStringValues Enumerates the set of values in String for SqlPlanBaselineOriginEnum
func GetSqlPlanBaselineOriginEnumStringValues() []string {
	return []string{
		"ADDM_SQLTUNE",
		"AUTO_CAPTURE",
		"AUTO_SQLTUNE",
		"EVOLVE_AUTO_INDEX_LOAD",
		"EVOLVE_CREATE_FROM_ADAPTIVE",
		"EVOLVE_LOAD_FROM_STS",
		"EVOLVE_LOAD_FROM_AWR",
		"EVOLVE_LOAD_FROM_CURSOR_CACHE",
		"MANUAL_LOAD",
		"MANUAL_LOAD_FROM_AWR",
		"MANUAL_LOAD_FROM_CURSOR_CACHE",
		"MANUAL_LOAD_FROM_STS",
		"MANUAL_SQLTUNE",
		"STORED_OUTLINE",
		"UNKNOWN",
	}
}

// GetMappingSqlPlanBaselineOriginEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineOriginEnum(val string) (SqlPlanBaselineOriginEnum, bool) {
	enum, ok := mappingSqlPlanBaselineOriginEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
