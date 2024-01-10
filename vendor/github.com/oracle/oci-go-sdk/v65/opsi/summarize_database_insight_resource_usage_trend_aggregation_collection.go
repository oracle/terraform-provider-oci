// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeDatabaseInsightResourceUsageTrendAggregationCollection Top level response object.
type SummarizeDatabaseInsightResourceUsageTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (example: CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Usage Data with time stamps
	UsageData []ResourceUsageTrendAggregation `mandatory:"true" json:"usageData"`
}

func (m SummarizeDatabaseInsightResourceUsageTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeDatabaseInsightResourceUsageTrendAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricCpu       SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricStorage   SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricIo        SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemory    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemoryPga SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "MEMORY_PGA"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemorySga SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "MEMORY_SGA"
)

var mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = map[string]SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum{
	"CPU":        SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricCpu,
	"STORAGE":    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricStorage,
	"IO":         SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricIo,
	"MEMORY":     SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemory,
	"MEMORY_PGA": SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemoryPga,
	"MEMORY_SGA": SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemorySga,
}

var mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumLowerCase = map[string]SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum{
	"cpu":        SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricCpu,
	"storage":    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricStorage,
	"io":         SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricIo,
	"memory":     SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemory,
	"memory_pga": SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemoryPga,
	"memory_sga": SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricMemorySga,
}

// GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumValues() []SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"MEMORY_PGA",
		"MEMORY_SGA",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum(val string) (SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum
const (
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitCores   SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum = "CORES"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitGb      SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum = "GB"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitMbps    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum = "MBPS"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitIops    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum = "IOPS"
	SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitPercent SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum = map[string]SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum{
	"CORES":   SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitCores,
	"GB":      SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitGb,
	"MBPS":    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitMbps,
	"IOPS":    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitIops,
	"PERCENT": SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitPercent,
}

var mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnumLowerCase = map[string]SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum{
	"cores":   SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitCores,
	"gb":      SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitGb,
	"mbps":    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitMbps,
	"iops":    SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitIops,
	"percent": SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitPercent,
}

// GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum
func GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnumValues() []SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum {
	values := make([]SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum
func GetSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum(val string) (SummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceUsageTrendAggregationCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
