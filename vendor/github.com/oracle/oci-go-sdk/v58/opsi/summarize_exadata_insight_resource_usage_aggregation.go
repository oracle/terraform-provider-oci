// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
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

	// Displays usage unit (CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE).
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// Percentage change in resource usage during the current period calculated using linear regression functions
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`
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
	if _, ok := GetMappingUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetUsageUnitEnumStringValues(), ",")))
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
	mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum)
	for k, v := range mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum {
		mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnumIgnoreCase[strings.ToLower(val)]
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
	mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum)
	for k, v := range mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum {
		mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
