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

// SummarizeDatabaseInsightResourceUsageAggregation Resource usage summation for the current time period
type SummarizeDatabaseInsightResourceUsageAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (example: CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE) for a set of databases.
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// Percentage change in resource usage during the current period calculated using linear regression functions
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`

	// The maximum host CPUs (cores x threads/core) on the underlying infrastructure. This only applies to CPU and does not not apply for Autonomous Databases.
	TotalHostCapacity *float64 `mandatory:"false" json:"totalHostCapacity"`
}

func (m SummarizeDatabaseInsightResourceUsageAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeDatabaseInsightResourceUsageAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceUsageAggregationResourceMetricCpu       SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceUsageAggregationResourceMetricStorage   SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceUsageAggregationResourceMetricIo        SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemory    SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = "MEMORY"
	SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemoryPga SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = "MEMORY_PGA"
	SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemorySga SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = "MEMORY_SGA"
)

var mappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum = map[string]SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum{
	"CPU":        SummarizeDatabaseInsightResourceUsageAggregationResourceMetricCpu,
	"STORAGE":    SummarizeDatabaseInsightResourceUsageAggregationResourceMetricStorage,
	"IO":         SummarizeDatabaseInsightResourceUsageAggregationResourceMetricIo,
	"MEMORY":     SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemory,
	"MEMORY_PGA": SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemoryPga,
	"MEMORY_SGA": SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemorySga,
}

var mappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumLowerCase = map[string]SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum{
	"cpu":        SummarizeDatabaseInsightResourceUsageAggregationResourceMetricCpu,
	"storage":    SummarizeDatabaseInsightResourceUsageAggregationResourceMetricStorage,
	"io":         SummarizeDatabaseInsightResourceUsageAggregationResourceMetricIo,
	"memory":     SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemory,
	"memory_pga": SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemoryPga,
	"memory_sga": SummarizeDatabaseInsightResourceUsageAggregationResourceMetricMemorySga,
}

// GetSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumValues() []SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"MEMORY_PGA",
		"MEMORY_SGA",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum(val string) (SummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceUsageAggregationResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum
const (
	SummarizeDatabaseInsightResourceUsageAggregationUsageUnitCores   SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum = "CORES"
	SummarizeDatabaseInsightResourceUsageAggregationUsageUnitGb      SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum = "GB"
	SummarizeDatabaseInsightResourceUsageAggregationUsageUnitMbps    SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum = "MBPS"
	SummarizeDatabaseInsightResourceUsageAggregationUsageUnitIops    SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum = "IOPS"
	SummarizeDatabaseInsightResourceUsageAggregationUsageUnitPercent SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum = "PERCENT"
)

var mappingSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum = map[string]SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum{
	"CORES":   SummarizeDatabaseInsightResourceUsageAggregationUsageUnitCores,
	"GB":      SummarizeDatabaseInsightResourceUsageAggregationUsageUnitGb,
	"MBPS":    SummarizeDatabaseInsightResourceUsageAggregationUsageUnitMbps,
	"IOPS":    SummarizeDatabaseInsightResourceUsageAggregationUsageUnitIops,
	"PERCENT": SummarizeDatabaseInsightResourceUsageAggregationUsageUnitPercent,
}

var mappingSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnumLowerCase = map[string]SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum{
	"cores":   SummarizeDatabaseInsightResourceUsageAggregationUsageUnitCores,
	"gb":      SummarizeDatabaseInsightResourceUsageAggregationUsageUnitGb,
	"mbps":    SummarizeDatabaseInsightResourceUsageAggregationUsageUnitMbps,
	"iops":    SummarizeDatabaseInsightResourceUsageAggregationUsageUnitIops,
	"percent": SummarizeDatabaseInsightResourceUsageAggregationUsageUnitPercent,
}

// GetSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum
func GetSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnumValues() []SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum
func GetSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum(val string) (SummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceUsageAggregationUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
