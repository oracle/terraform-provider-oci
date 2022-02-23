// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// AnalyticsClusterMemoryEstimateStatusEnum Enum with underlying type: string
type AnalyticsClusterMemoryEstimateStatusEnum string

// Set of constants representing the allowable values for AnalyticsClusterMemoryEstimateStatusEnum
const (
	AnalyticsClusterMemoryEstimateStatusAccepted   AnalyticsClusterMemoryEstimateStatusEnum = "ACCEPTED"
	AnalyticsClusterMemoryEstimateStatusInProgress AnalyticsClusterMemoryEstimateStatusEnum = "IN_PROGRESS"
	AnalyticsClusterMemoryEstimateStatusFailed     AnalyticsClusterMemoryEstimateStatusEnum = "FAILED"
	AnalyticsClusterMemoryEstimateStatusSucceeded  AnalyticsClusterMemoryEstimateStatusEnum = "SUCCEEDED"
	AnalyticsClusterMemoryEstimateStatusCanceling  AnalyticsClusterMemoryEstimateStatusEnum = "CANCELING"
	AnalyticsClusterMemoryEstimateStatusCanceled   AnalyticsClusterMemoryEstimateStatusEnum = "CANCELED"
)

var mappingAnalyticsClusterMemoryEstimateStatusEnum = map[string]AnalyticsClusterMemoryEstimateStatusEnum{
	"ACCEPTED":    AnalyticsClusterMemoryEstimateStatusAccepted,
	"IN_PROGRESS": AnalyticsClusterMemoryEstimateStatusInProgress,
	"FAILED":      AnalyticsClusterMemoryEstimateStatusFailed,
	"SUCCEEDED":   AnalyticsClusterMemoryEstimateStatusSucceeded,
	"CANCELING":   AnalyticsClusterMemoryEstimateStatusCanceling,
	"CANCELED":    AnalyticsClusterMemoryEstimateStatusCanceled,
}

var mappingAnalyticsClusterMemoryEstimateStatusEnumLowerCase = map[string]AnalyticsClusterMemoryEstimateStatusEnum{
	"accepted":    AnalyticsClusterMemoryEstimateStatusAccepted,
	"in_progress": AnalyticsClusterMemoryEstimateStatusInProgress,
	"failed":      AnalyticsClusterMemoryEstimateStatusFailed,
	"succeeded":   AnalyticsClusterMemoryEstimateStatusSucceeded,
	"canceling":   AnalyticsClusterMemoryEstimateStatusCanceling,
	"canceled":    AnalyticsClusterMemoryEstimateStatusCanceled,
}

// GetAnalyticsClusterMemoryEstimateStatusEnumValues Enumerates the set of values for AnalyticsClusterMemoryEstimateStatusEnum
func GetAnalyticsClusterMemoryEstimateStatusEnumValues() []AnalyticsClusterMemoryEstimateStatusEnum {
	values := make([]AnalyticsClusterMemoryEstimateStatusEnum, 0)
	for _, v := range mappingAnalyticsClusterMemoryEstimateStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAnalyticsClusterMemoryEstimateStatusEnumStringValues Enumerates the set of values in String for AnalyticsClusterMemoryEstimateStatusEnum
func GetAnalyticsClusterMemoryEstimateStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingAnalyticsClusterMemoryEstimateStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnalyticsClusterMemoryEstimateStatusEnum(val string) (AnalyticsClusterMemoryEstimateStatusEnum, bool) {
	enum, ok := mappingAnalyticsClusterMemoryEstimateStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
