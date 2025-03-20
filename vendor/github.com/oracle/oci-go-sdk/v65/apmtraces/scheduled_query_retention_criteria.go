// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"strings"
)

// ScheduledQueryRetentionCriteriaEnum Enum with underlying type: string
type ScheduledQueryRetentionCriteriaEnum string

// Set of constants representing the allowable values for ScheduledQueryRetentionCriteriaEnum
const (
	ScheduledQueryRetentionCriteriaKeepDataUntilRetentionPeriod ScheduledQueryRetentionCriteriaEnum = "KEEP_DATA_UNTIL_RETENTION_PERIOD"
	ScheduledQueryRetentionCriteriaUpdate                       ScheduledQueryRetentionCriteriaEnum = "UPDATE"
)

var mappingScheduledQueryRetentionCriteriaEnum = map[string]ScheduledQueryRetentionCriteriaEnum{
	"KEEP_DATA_UNTIL_RETENTION_PERIOD": ScheduledQueryRetentionCriteriaKeepDataUntilRetentionPeriod,
	"UPDATE":                           ScheduledQueryRetentionCriteriaUpdate,
}

var mappingScheduledQueryRetentionCriteriaEnumLowerCase = map[string]ScheduledQueryRetentionCriteriaEnum{
	"keep_data_until_retention_period": ScheduledQueryRetentionCriteriaKeepDataUntilRetentionPeriod,
	"update":                           ScheduledQueryRetentionCriteriaUpdate,
}

// GetScheduledQueryRetentionCriteriaEnumValues Enumerates the set of values for ScheduledQueryRetentionCriteriaEnum
func GetScheduledQueryRetentionCriteriaEnumValues() []ScheduledQueryRetentionCriteriaEnum {
	values := make([]ScheduledQueryRetentionCriteriaEnum, 0)
	for _, v := range mappingScheduledQueryRetentionCriteriaEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledQueryRetentionCriteriaEnumStringValues Enumerates the set of values in String for ScheduledQueryRetentionCriteriaEnum
func GetScheduledQueryRetentionCriteriaEnumStringValues() []string {
	return []string{
		"KEEP_DATA_UNTIL_RETENTION_PERIOD",
		"UPDATE",
	}
}

// GetMappingScheduledQueryRetentionCriteriaEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledQueryRetentionCriteriaEnum(val string) (ScheduledQueryRetentionCriteriaEnum, bool) {
	enum, ok := mappingScheduledQueryRetentionCriteriaEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
