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

// SummarizeExadataInsightResourceUsageAggregation Resource usage summation for the current time period
type SummarizeExadataInsightResourceUsageAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE) for a set of databases.
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// Percentage change in resource usage during the current period calculated using linear regression functions
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`

	// The maximum host CPUs (cores x threads/core) on the underlying infrastructure. This only applies to CPU and does not not apply for Autonomous Databases.
	TotalHostCapacity *float64 `mandatory:"false" json:"totalHostCapacity"`
}

func (m SummarizeExadataInsightResourceUsageAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeExadataInsightResourceUsageAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum(string(m.ExadataResourceMetric)); !ok && m.ExadataResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceMetric: %s. Supported values are: %s.", m.ExadataResourceMetric, strings.Join(GetSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum(string(m.ExadataResourceType)); !ok && m.ExadataResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceType: %s. Supported values are: %s.", m.ExadataResourceType, strings.Join(GetSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceUsageAggregationUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeExadataInsightResourceUsageAggregationUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricCpu        SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricStorage    SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricIo         SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricMemory     SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricIops       SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricThroughput SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum = map[string]SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricThroughput,
}

var mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumLowerCase = map[string]SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum{
	"cpu":        SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricCpu,
	"storage":    SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricStorage,
	"io":         SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricIo,
	"memory":     SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricMemory,
	"iops":       SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricIops,
	"throughput": SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum(val string) (SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeDatabase      SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeHost          SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeStorageServer SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum = map[string]SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeDiskgroup,
}

var mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumLowerCase = map[string]SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum{
	"database":       SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeDatabase,
	"host":           SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeHost,
	"storage_server": SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeStorageServer,
	"diskgroup":      SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"HOST",
		"STORAGE_SERVER",
		"DISKGROUP",
	}
}

// GetMappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum(val string) (SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum
const (
	SummarizeExadataInsightResourceUsageAggregationUsageUnitCores   SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum = "CORES"
	SummarizeExadataInsightResourceUsageAggregationUsageUnitGb      SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum = "GB"
	SummarizeExadataInsightResourceUsageAggregationUsageUnitMbps    SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum = "MBPS"
	SummarizeExadataInsightResourceUsageAggregationUsageUnitIops    SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum = "IOPS"
	SummarizeExadataInsightResourceUsageAggregationUsageUnitPercent SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum = "PERCENT"
)

var mappingSummarizeExadataInsightResourceUsageAggregationUsageUnitEnum = map[string]SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum{
	"CORES":   SummarizeExadataInsightResourceUsageAggregationUsageUnitCores,
	"GB":      SummarizeExadataInsightResourceUsageAggregationUsageUnitGb,
	"MBPS":    SummarizeExadataInsightResourceUsageAggregationUsageUnitMbps,
	"IOPS":    SummarizeExadataInsightResourceUsageAggregationUsageUnitIops,
	"PERCENT": SummarizeExadataInsightResourceUsageAggregationUsageUnitPercent,
}

var mappingSummarizeExadataInsightResourceUsageAggregationUsageUnitEnumLowerCase = map[string]SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum{
	"cores":   SummarizeExadataInsightResourceUsageAggregationUsageUnitCores,
	"gb":      SummarizeExadataInsightResourceUsageAggregationUsageUnitGb,
	"mbps":    SummarizeExadataInsightResourceUsageAggregationUsageUnitMbps,
	"iops":    SummarizeExadataInsightResourceUsageAggregationUsageUnitIops,
	"percent": SummarizeExadataInsightResourceUsageAggregationUsageUnitPercent,
}

// GetSummarizeExadataInsightResourceUsageAggregationUsageUnitEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum
func GetSummarizeExadataInsightResourceUsageAggregationUsageUnitEnumValues() []SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum {
	values := make([]SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageAggregationUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUsageAggregationUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum
func GetSummarizeExadataInsightResourceUsageAggregationUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeExadataInsightResourceUsageAggregationUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUsageAggregationUsageUnitEnum(val string) (SummarizeExadataInsightResourceUsageAggregationUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceUsageAggregationUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
