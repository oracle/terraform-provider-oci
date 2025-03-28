// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeHostInsightResourceCapacityTrendAggregationCollection Top level response object.
type SummarizeHostInsightResourceCapacityTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"true" json:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered lowly utilized.
	LowUtilizationThreshold *int `mandatory:"true" json:"lowUtilizationThreshold"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Capacity Data with timestamp.
	CapacityData []HostResourceCapacityTrendAggregation `mandatory:"true" json:"capacityData"`
}

func (m SummarizeHostInsightResourceCapacityTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightResourceCapacityTrendAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
const (
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu            SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory         SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricLogicalMemory  SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "LOGICAL_MEMORY"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricStorage        SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricNetwork        SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "NETWORK"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricGpuUtilization SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "GPU_UTILIZATION"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricGpuMemoryUsage SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "GPU_MEMORY_USAGE"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricIo             SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "IO"
)

var mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = map[string]SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum{
	"CPU":              SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu,
	"MEMORY":           SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory,
	"LOGICAL_MEMORY":   SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricLogicalMemory,
	"STORAGE":          SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricStorage,
	"NETWORK":          SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricNetwork,
	"GPU_UTILIZATION":  SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricGpuUtilization,
	"GPU_MEMORY_USAGE": SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricGpuMemoryUsage,
	"IO":               SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricIo,
}

var mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumLowerCase = map[string]SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum{
	"cpu":              SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu,
	"memory":           SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory,
	"logical_memory":   SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricLogicalMemory,
	"storage":          SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricStorage,
	"network":          SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricNetwork,
	"gpu_utilization":  SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricGpuUtilization,
	"gpu_memory_usage": SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricGpuMemoryUsage,
	"io":               SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricIo,
}

// GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
func GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumValues() []SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
func GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"LOGICAL_MEMORY",
		"STORAGE",
		"NETWORK",
		"GPU_UTILIZATION",
		"GPU_MEMORY_USAGE",
		"IO",
	}
}

// GetMappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum(val string) (SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum
const (
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitCores   SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "CORES"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitGb      SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "GB"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitMbps    SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "MBPS"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitIops    SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "IOPS"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitPercent SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = map[string]SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum{
	"CORES":   SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitCores,
	"GB":      SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitGb,
	"MBPS":    SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitMbps,
	"IOPS":    SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitIops,
	"PERCENT": SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitPercent,
}

var mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumLowerCase = map[string]SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum{
	"cores":   SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitCores,
	"gb":      SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitGb,
	"mbps":    SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitMbps,
	"iops":    SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitIops,
	"percent": SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitPercent,
}

// GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum
func GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumValues() []SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum
func GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum(val string) (SummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
