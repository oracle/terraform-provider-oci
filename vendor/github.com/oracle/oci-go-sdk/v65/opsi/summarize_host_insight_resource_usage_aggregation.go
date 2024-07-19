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

// SummarizeHostInsightResourceUsageAggregation Resource usage summation for the current time period.
type SummarizeHostInsightResourceUsageAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceUsageAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeHostInsightResourceUsageAggregationUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE) for a set of databases.
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// Percentage change in resource usage during the current period calculated using linear regression functions
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`
}

func (m SummarizeHostInsightResourceUsageAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightResourceUsageAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeHostInsightResourceUsageAggregationUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeHostInsightResourceUsageAggregationUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceUsageAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceUsageAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUsageAggregationResourceMetricEnum
const (
	SummarizeHostInsightResourceUsageAggregationResourceMetricCpu            SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceUsageAggregationResourceMetricMemory         SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceUsageAggregationResourceMetricLogicalMemory  SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "LOGICAL_MEMORY"
	SummarizeHostInsightResourceUsageAggregationResourceMetricStorage        SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "STORAGE"
	SummarizeHostInsightResourceUsageAggregationResourceMetricNetwork        SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "NETWORK"
	SummarizeHostInsightResourceUsageAggregationResourceMetricGpuUtilization SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "GPU_UTILIZATION"
	SummarizeHostInsightResourceUsageAggregationResourceMetricGpuMemoryUsage SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "GPU_MEMORY_USAGE"
)

var mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum = map[string]SummarizeHostInsightResourceUsageAggregationResourceMetricEnum{
	"CPU":              SummarizeHostInsightResourceUsageAggregationResourceMetricCpu,
	"MEMORY":           SummarizeHostInsightResourceUsageAggregationResourceMetricMemory,
	"LOGICAL_MEMORY":   SummarizeHostInsightResourceUsageAggregationResourceMetricLogicalMemory,
	"STORAGE":          SummarizeHostInsightResourceUsageAggregationResourceMetricStorage,
	"NETWORK":          SummarizeHostInsightResourceUsageAggregationResourceMetricNetwork,
	"GPU_UTILIZATION":  SummarizeHostInsightResourceUsageAggregationResourceMetricGpuUtilization,
	"GPU_MEMORY_USAGE": SummarizeHostInsightResourceUsageAggregationResourceMetricGpuMemoryUsage,
}

var mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnumLowerCase = map[string]SummarizeHostInsightResourceUsageAggregationResourceMetricEnum{
	"cpu":              SummarizeHostInsightResourceUsageAggregationResourceMetricCpu,
	"memory":           SummarizeHostInsightResourceUsageAggregationResourceMetricMemory,
	"logical_memory":   SummarizeHostInsightResourceUsageAggregationResourceMetricLogicalMemory,
	"storage":          SummarizeHostInsightResourceUsageAggregationResourceMetricStorage,
	"network":          SummarizeHostInsightResourceUsageAggregationResourceMetricNetwork,
	"gpu_utilization":  SummarizeHostInsightResourceUsageAggregationResourceMetricGpuUtilization,
	"gpu_memory_usage": SummarizeHostInsightResourceUsageAggregationResourceMetricGpuMemoryUsage,
}

// GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceUsageAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumValues() []SummarizeHostInsightResourceUsageAggregationResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceUsageAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceUsageAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"LOGICAL_MEMORY",
		"STORAGE",
		"NETWORK",
		"GPU_UTILIZATION",
		"GPU_MEMORY_USAGE",
	}
}

// GetMappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum(val string) (SummarizeHostInsightResourceUsageAggregationResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceUsageAggregationUsageUnitEnum Enum with underlying type: string
type SummarizeHostInsightResourceUsageAggregationUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUsageAggregationUsageUnitEnum
const (
	SummarizeHostInsightResourceUsageAggregationUsageUnitCores   SummarizeHostInsightResourceUsageAggregationUsageUnitEnum = "CORES"
	SummarizeHostInsightResourceUsageAggregationUsageUnitGb      SummarizeHostInsightResourceUsageAggregationUsageUnitEnum = "GB"
	SummarizeHostInsightResourceUsageAggregationUsageUnitMbps    SummarizeHostInsightResourceUsageAggregationUsageUnitEnum = "MBPS"
	SummarizeHostInsightResourceUsageAggregationUsageUnitIops    SummarizeHostInsightResourceUsageAggregationUsageUnitEnum = "IOPS"
	SummarizeHostInsightResourceUsageAggregationUsageUnitPercent SummarizeHostInsightResourceUsageAggregationUsageUnitEnum = "PERCENT"
)

var mappingSummarizeHostInsightResourceUsageAggregationUsageUnitEnum = map[string]SummarizeHostInsightResourceUsageAggregationUsageUnitEnum{
	"CORES":   SummarizeHostInsightResourceUsageAggregationUsageUnitCores,
	"GB":      SummarizeHostInsightResourceUsageAggregationUsageUnitGb,
	"MBPS":    SummarizeHostInsightResourceUsageAggregationUsageUnitMbps,
	"IOPS":    SummarizeHostInsightResourceUsageAggregationUsageUnitIops,
	"PERCENT": SummarizeHostInsightResourceUsageAggregationUsageUnitPercent,
}

var mappingSummarizeHostInsightResourceUsageAggregationUsageUnitEnumLowerCase = map[string]SummarizeHostInsightResourceUsageAggregationUsageUnitEnum{
	"cores":   SummarizeHostInsightResourceUsageAggregationUsageUnitCores,
	"gb":      SummarizeHostInsightResourceUsageAggregationUsageUnitGb,
	"mbps":    SummarizeHostInsightResourceUsageAggregationUsageUnitMbps,
	"iops":    SummarizeHostInsightResourceUsageAggregationUsageUnitIops,
	"percent": SummarizeHostInsightResourceUsageAggregationUsageUnitPercent,
}

// GetSummarizeHostInsightResourceUsageAggregationUsageUnitEnumValues Enumerates the set of values for SummarizeHostInsightResourceUsageAggregationUsageUnitEnum
func GetSummarizeHostInsightResourceUsageAggregationUsageUnitEnumValues() []SummarizeHostInsightResourceUsageAggregationUsageUnitEnum {
	values := make([]SummarizeHostInsightResourceUsageAggregationUsageUnitEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUsageAggregationUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceUsageAggregationUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceUsageAggregationUsageUnitEnum
func GetSummarizeHostInsightResourceUsageAggregationUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeHostInsightResourceUsageAggregationUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceUsageAggregationUsageUnitEnum(val string) (SummarizeHostInsightResourceUsageAggregationUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceUsageAggregationUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
