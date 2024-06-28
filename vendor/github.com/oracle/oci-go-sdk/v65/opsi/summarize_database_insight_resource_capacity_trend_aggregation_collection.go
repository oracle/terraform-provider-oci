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

// SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection Collection of resource capacity trend.
type SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"true" json:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered lowly utilized.
	LowUtilizationThreshold *int `mandatory:"true" json:"lowUtilizationThreshold"`

	// Defines the type of resource metric (example: CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Capacity Data with time interval
	CapacityData []ResourceCapacityTrendAggregation `mandatory:"true" json:"capacityData"`
}

func (m SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu       SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricStorage   SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricIo        SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemoryPga SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "MEMORY_PGA"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemorySga SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "MEMORY_SGA"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = map[string]SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum{
	"CPU":        SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu,
	"STORAGE":    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricStorage,
	"IO":         SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricIo,
	"MEMORY":     SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory,
	"MEMORY_PGA": SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemoryPga,
	"MEMORY_SGA": SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemorySga,
}

var mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumLowerCase = map[string]SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum{
	"cpu":        SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu,
	"storage":    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricStorage,
	"io":         SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricIo,
	"memory":     SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory,
	"memory_pga": SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemoryPga,
	"memory_sga": SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemorySga,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"MEMORY_PGA",
		"MEMORY_SGA",
	}
}

// GetMappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum(val string) (SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitCores   SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "CORES"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitGb      SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "GB"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitMbps    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "MBPS"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitIops    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "IOPS"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitPercent SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum = map[string]SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum{
	"CORES":   SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitCores,
	"GB":      SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitGb,
	"MBPS":    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitMbps,
	"IOPS":    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitIops,
	"PERCENT": SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitPercent,
}

var mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumLowerCase = map[string]SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum{
	"cores":   SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitCores,
	"gb":      SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitGb,
	"mbps":    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitMbps,
	"iops":    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitIops,
	"percent": SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitPercent,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum(val string) (SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
