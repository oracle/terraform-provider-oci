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

// SummarizeExadataInsightResourceCapacityTrendAggregation Collection of resource capacity trend.
type SummarizeExadataInsightResourceCapacityTrendAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Capacity Data with time interval
	CapacityData []ExadataInsightResourceCapacityTrendAggregation `mandatory:"true" json:"capacityData"`
}

func (m SummarizeExadataInsightResourceCapacityTrendAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeExadataInsightResourceCapacityTrendAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum(string(m.ExadataResourceMetric)); !ok && m.ExadataResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceMetric: %s. Supported values are: %s.", m.ExadataResourceMetric, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum(string(m.ExadataResourceType)); !ok && m.ExadataResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceType: %s. Supported values are: %s.", m.ExadataResourceType, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricCpu        SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricStorage    SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricIo         SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricMemory     SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricIops       SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricThroughput SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum = map[string]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricThroughput,
}

var mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumLowerCase = map[string]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum{
	"cpu":        SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricCpu,
	"storage":    SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricStorage,
	"io":         SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricIo,
	"memory":     SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricMemory,
	"iops":       SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricIops,
	"throughput": SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum(val string) (SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeDatabase      SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeHost          SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeStorageServer SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum = map[string]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeDiskgroup,
}

var mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumLowerCase = map[string]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum{
	"database":       SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeDatabase,
	"host":           SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeHost,
	"storage_server": SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeStorageServer,
	"diskgroup":      SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"HOST",
		"STORAGE_SERVER",
		"DISKGROUP",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum(val string) (SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum
const (
	SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitCores   SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum = "CORES"
	SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitGb      SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum = "GB"
	SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitMbps    SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum = "MBPS"
	SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitIops    SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum = "IOPS"
	SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitPercent SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum = "PERCENT"
)

var mappingSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum = map[string]SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum{
	"CORES":   SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitCores,
	"GB":      SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitGb,
	"MBPS":    SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitMbps,
	"IOPS":    SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitIops,
	"PERCENT": SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitPercent,
}

var mappingSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnumLowerCase = map[string]SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum{
	"cores":   SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitCores,
	"gb":      SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitGb,
	"mbps":    SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitMbps,
	"iops":    SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitIops,
	"percent": SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitPercent,
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnumValues() []SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum(val string) (SummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendAggregationUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
