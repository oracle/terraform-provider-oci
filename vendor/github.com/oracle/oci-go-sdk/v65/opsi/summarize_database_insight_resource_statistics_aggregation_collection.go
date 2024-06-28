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

// SummarizeDatabaseInsightResourceStatisticsAggregationCollection Returns list of the Databases with resource statistics like usage, capacity, utilization and usage change percent.
type SummarizeDatabaseInsightResourceStatisticsAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"true" json:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered lowly utilized.
	LowUtilizationThreshold *int `mandatory:"true" json:"lowUtilizationThreshold"`

	// Defines the type of resource metric (example: CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of Resource Statistics items
	Items []ResourceStatisticsAggregation `mandatory:"true" json:"items"`
}

func (m SummarizeDatabaseInsightResourceStatisticsAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeDatabaseInsightResourceStatisticsAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricCpu       SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricStorage   SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricIo        SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemory    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemoryPga SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY_PGA"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemorySga SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY_SGA"
)

var mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = map[string]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum{
	"CPU":        SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricCpu,
	"STORAGE":    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricStorage,
	"IO":         SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricIo,
	"MEMORY":     SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemory,
	"MEMORY_PGA": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemoryPga,
	"MEMORY_SGA": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemorySga,
}

var mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumLowerCase = map[string]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum{
	"cpu":        SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricCpu,
	"storage":    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricStorage,
	"io":         SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricIo,
	"memory":     SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemory,
	"memory_pga": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemoryPga,
	"memory_sga": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemorySga,
}

// GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues() []SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"MEMORY_PGA",
		"MEMORY_SGA",
	}
}

// GetMappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum(val string) (SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum
const (
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitCores   SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "CORES"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitGb      SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "GB"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitMbps    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "MBPS"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitIops    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "IOPS"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitPercent SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum = map[string]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum{
	"CORES":   SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitCores,
	"GB":      SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitGb,
	"MBPS":    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitMbps,
	"IOPS":    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitIops,
	"PERCENT": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitPercent,
}

var mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnumLowerCase = map[string]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum{
	"cores":   SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitCores,
	"gb":      SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitGb,
	"mbps":    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitMbps,
	"iops":    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitIops,
	"percent": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitPercent,
}

// GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum
func GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnumValues() []SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum {
	values := make([]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum
func GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum(val string) (SummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
