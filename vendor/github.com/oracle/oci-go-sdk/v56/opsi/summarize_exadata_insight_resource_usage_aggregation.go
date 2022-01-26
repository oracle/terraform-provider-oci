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
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetric = map[string]SummarizeExadataInsightResourceUsageAggregationExadataResourceMetricEnum{
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
	for _, v := range mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceMetric {
		values = append(values, v)
	}
	return values
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

var mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceType = map[string]SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceUsageAggregationExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageAggregationExadataResourceType {
		values = append(values, v)
	}
	return values
}
