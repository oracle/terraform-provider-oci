// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeHostInsightResourceStatisticsAggregationCollection Returns list of hosts with resource statistics like usage, capacity, utilization, usage change percent and load.
type SummarizeHostInsightResourceStatisticsAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"true" json:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered lowly utilized.
	LowUtilizationThreshold *int `mandatory:"true" json:"lowUtilizationThreshold"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of Resource Statistics items
	Items []HostInsightResourceStatisticsAggregation `mandatory:"true" json:"items"`
}

func (m SummarizeHostInsightResourceStatisticsAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightResourceStatisticsAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum
const (
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricCpu            SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricMemory         SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricLogicalMemory  SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "LOGICAL_MEMORY"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricStorage        SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricNetwork        SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "NETWORK"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricGpuUtilization SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "GPU_UTILIZATION"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricGpuMemoryUsage SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "GPU_MEMORY_USAGE"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricIo             SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "IO"
)

var mappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = map[string]SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum{
	"CPU":              SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricCpu,
	"MEMORY":           SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricMemory,
	"LOGICAL_MEMORY":   SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricLogicalMemory,
	"STORAGE":          SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricStorage,
	"NETWORK":          SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricNetwork,
	"GPU_UTILIZATION":  SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricGpuUtilization,
	"GPU_MEMORY_USAGE": SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricGpuMemoryUsage,
	"IO":               SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricIo,
}

var mappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumLowerCase = map[string]SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum{
	"cpu":              SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricCpu,
	"memory":           SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricMemory,
	"logical_memory":   SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricLogicalMemory,
	"storage":          SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricStorage,
	"network":          SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricNetwork,
	"gpu_utilization":  SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricGpuUtilization,
	"gpu_memory_usage": SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricGpuMemoryUsage,
	"io":               SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricIo,
}

// GetSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum
func GetSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues() []SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum
func GetSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumStringValues() []string {
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

// GetMappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum(val string) (SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum
const (
	SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitCores   SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "CORES"
	SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitGb      SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "GB"
	SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitMbps    SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "MBPS"
	SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitIops    SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "IOPS"
	SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitPercent SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum = map[string]SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum{
	"CORES":   SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitCores,
	"GB":      SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitGb,
	"MBPS":    SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitMbps,
	"IOPS":    SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitIops,
	"PERCENT": SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitPercent,
}

var mappingSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnumLowerCase = map[string]SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum{
	"cores":   SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitCores,
	"gb":      SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitGb,
	"mbps":    SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitMbps,
	"iops":    SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitIops,
	"percent": SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitPercent,
}

// GetSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum
func GetSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnumValues() []SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum {
	values := make([]SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum
func GetSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum(val string) (SummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceStatisticsAggregationCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
