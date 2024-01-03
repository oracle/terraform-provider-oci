// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// HeatWaveClusterMemoryEstimateStatusEnum Enum with underlying type: string
type HeatWaveClusterMemoryEstimateStatusEnum string

// Set of constants representing the allowable values for HeatWaveClusterMemoryEstimateStatusEnum
const (
	HeatWaveClusterMemoryEstimateStatusAccepted   HeatWaveClusterMemoryEstimateStatusEnum = "ACCEPTED"
	HeatWaveClusterMemoryEstimateStatusInProgress HeatWaveClusterMemoryEstimateStatusEnum = "IN_PROGRESS"
	HeatWaveClusterMemoryEstimateStatusFailed     HeatWaveClusterMemoryEstimateStatusEnum = "FAILED"
	HeatWaveClusterMemoryEstimateStatusSucceeded  HeatWaveClusterMemoryEstimateStatusEnum = "SUCCEEDED"
	HeatWaveClusterMemoryEstimateStatusCanceling  HeatWaveClusterMemoryEstimateStatusEnum = "CANCELING"
	HeatWaveClusterMemoryEstimateStatusCanceled   HeatWaveClusterMemoryEstimateStatusEnum = "CANCELED"
)

var mappingHeatWaveClusterMemoryEstimateStatusEnum = map[string]HeatWaveClusterMemoryEstimateStatusEnum{
	"ACCEPTED":    HeatWaveClusterMemoryEstimateStatusAccepted,
	"IN_PROGRESS": HeatWaveClusterMemoryEstimateStatusInProgress,
	"FAILED":      HeatWaveClusterMemoryEstimateStatusFailed,
	"SUCCEEDED":   HeatWaveClusterMemoryEstimateStatusSucceeded,
	"CANCELING":   HeatWaveClusterMemoryEstimateStatusCanceling,
	"CANCELED":    HeatWaveClusterMemoryEstimateStatusCanceled,
}

var mappingHeatWaveClusterMemoryEstimateStatusEnumLowerCase = map[string]HeatWaveClusterMemoryEstimateStatusEnum{
	"accepted":    HeatWaveClusterMemoryEstimateStatusAccepted,
	"in_progress": HeatWaveClusterMemoryEstimateStatusInProgress,
	"failed":      HeatWaveClusterMemoryEstimateStatusFailed,
	"succeeded":   HeatWaveClusterMemoryEstimateStatusSucceeded,
	"canceling":   HeatWaveClusterMemoryEstimateStatusCanceling,
	"canceled":    HeatWaveClusterMemoryEstimateStatusCanceled,
}

// GetHeatWaveClusterMemoryEstimateStatusEnumValues Enumerates the set of values for HeatWaveClusterMemoryEstimateStatusEnum
func GetHeatWaveClusterMemoryEstimateStatusEnumValues() []HeatWaveClusterMemoryEstimateStatusEnum {
	values := make([]HeatWaveClusterMemoryEstimateStatusEnum, 0)
	for _, v := range mappingHeatWaveClusterMemoryEstimateStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetHeatWaveClusterMemoryEstimateStatusEnumStringValues Enumerates the set of values in String for HeatWaveClusterMemoryEstimateStatusEnum
func GetHeatWaveClusterMemoryEstimateStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingHeatWaveClusterMemoryEstimateStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHeatWaveClusterMemoryEstimateStatusEnum(val string) (HeatWaveClusterMemoryEstimateStatusEnum, bool) {
	enum, ok := mappingHeatWaveClusterMemoryEstimateStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
