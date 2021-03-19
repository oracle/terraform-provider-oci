// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

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

var mappingHeatWaveClusterMemoryEstimateStatus = map[string]HeatWaveClusterMemoryEstimateStatusEnum{
	"ACCEPTED":    HeatWaveClusterMemoryEstimateStatusAccepted,
	"IN_PROGRESS": HeatWaveClusterMemoryEstimateStatusInProgress,
	"FAILED":      HeatWaveClusterMemoryEstimateStatusFailed,
	"SUCCEEDED":   HeatWaveClusterMemoryEstimateStatusSucceeded,
	"CANCELING":   HeatWaveClusterMemoryEstimateStatusCanceling,
	"CANCELED":    HeatWaveClusterMemoryEstimateStatusCanceled,
}

// GetHeatWaveClusterMemoryEstimateStatusEnumValues Enumerates the set of values for HeatWaveClusterMemoryEstimateStatusEnum
func GetHeatWaveClusterMemoryEstimateStatusEnumValues() []HeatWaveClusterMemoryEstimateStatusEnum {
	values := make([]HeatWaveClusterMemoryEstimateStatusEnum, 0)
	for _, v := range mappingHeatWaveClusterMemoryEstimateStatus {
		values = append(values, v)
	}
	return values
}
