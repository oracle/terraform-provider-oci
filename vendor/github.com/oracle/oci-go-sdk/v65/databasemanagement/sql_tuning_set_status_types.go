// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// SqlTuningSetStatusTypesEnum Enum with underlying type: string
type SqlTuningSetStatusTypesEnum string

// Set of constants representing the allowable values for SqlTuningSetStatusTypesEnum
const (
	SqlTuningSetStatusTypesDisabled            SqlTuningSetStatusTypesEnum = "DISABLED"
	SqlTuningSetStatusTypesRetryScheduled      SqlTuningSetStatusTypesEnum = "RETRY_SCHEDULED"
	SqlTuningSetStatusTypesScheduled           SqlTuningSetStatusTypesEnum = "SCHEDULED"
	SqlTuningSetStatusTypesBlocked             SqlTuningSetStatusTypesEnum = "BLOCKED"
	SqlTuningSetStatusTypesRunning             SqlTuningSetStatusTypesEnum = "RUNNING"
	SqlTuningSetStatusTypesCompleted           SqlTuningSetStatusTypesEnum = "COMPLETED"
	SqlTuningSetStatusTypesBroken              SqlTuningSetStatusTypesEnum = "BROKEN"
	SqlTuningSetStatusTypesFailed              SqlTuningSetStatusTypesEnum = "FAILED"
	SqlTuningSetStatusTypesRemote              SqlTuningSetStatusTypesEnum = "REMOTE"
	SqlTuningSetStatusTypesResourceUnavailable SqlTuningSetStatusTypesEnum = "RESOURCE_UNAVAILABLE"
	SqlTuningSetStatusTypesSucceeded           SqlTuningSetStatusTypesEnum = "SUCCEEDED"
	SqlTuningSetStatusTypesChainStalled        SqlTuningSetStatusTypesEnum = "CHAIN_STALLED"
)

var mappingSqlTuningSetStatusTypesEnum = map[string]SqlTuningSetStatusTypesEnum{
	"DISABLED":             SqlTuningSetStatusTypesDisabled,
	"RETRY_SCHEDULED":      SqlTuningSetStatusTypesRetryScheduled,
	"SCHEDULED":            SqlTuningSetStatusTypesScheduled,
	"BLOCKED":              SqlTuningSetStatusTypesBlocked,
	"RUNNING":              SqlTuningSetStatusTypesRunning,
	"COMPLETED":            SqlTuningSetStatusTypesCompleted,
	"BROKEN":               SqlTuningSetStatusTypesBroken,
	"FAILED":               SqlTuningSetStatusTypesFailed,
	"REMOTE":               SqlTuningSetStatusTypesRemote,
	"RESOURCE_UNAVAILABLE": SqlTuningSetStatusTypesResourceUnavailable,
	"SUCCEEDED":            SqlTuningSetStatusTypesSucceeded,
	"CHAIN_STALLED":        SqlTuningSetStatusTypesChainStalled,
}

var mappingSqlTuningSetStatusTypesEnumLowerCase = map[string]SqlTuningSetStatusTypesEnum{
	"disabled":             SqlTuningSetStatusTypesDisabled,
	"retry_scheduled":      SqlTuningSetStatusTypesRetryScheduled,
	"scheduled":            SqlTuningSetStatusTypesScheduled,
	"blocked":              SqlTuningSetStatusTypesBlocked,
	"running":              SqlTuningSetStatusTypesRunning,
	"completed":            SqlTuningSetStatusTypesCompleted,
	"broken":               SqlTuningSetStatusTypesBroken,
	"failed":               SqlTuningSetStatusTypesFailed,
	"remote":               SqlTuningSetStatusTypesRemote,
	"resource_unavailable": SqlTuningSetStatusTypesResourceUnavailable,
	"succeeded":            SqlTuningSetStatusTypesSucceeded,
	"chain_stalled":        SqlTuningSetStatusTypesChainStalled,
}

// GetSqlTuningSetStatusTypesEnumValues Enumerates the set of values for SqlTuningSetStatusTypesEnum
func GetSqlTuningSetStatusTypesEnumValues() []SqlTuningSetStatusTypesEnum {
	values := make([]SqlTuningSetStatusTypesEnum, 0)
	for _, v := range mappingSqlTuningSetStatusTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningSetStatusTypesEnumStringValues Enumerates the set of values in String for SqlTuningSetStatusTypesEnum
func GetSqlTuningSetStatusTypesEnumStringValues() []string {
	return []string{
		"DISABLED",
		"RETRY_SCHEDULED",
		"SCHEDULED",
		"BLOCKED",
		"RUNNING",
		"COMPLETED",
		"BROKEN",
		"FAILED",
		"REMOTE",
		"RESOURCE_UNAVAILABLE",
		"SUCCEEDED",
		"CHAIN_STALLED",
	}
}

// GetMappingSqlTuningSetStatusTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningSetStatusTypesEnum(val string) (SqlTuningSetStatusTypesEnum, bool) {
	enum, ok := mappingSqlTuningSetStatusTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
