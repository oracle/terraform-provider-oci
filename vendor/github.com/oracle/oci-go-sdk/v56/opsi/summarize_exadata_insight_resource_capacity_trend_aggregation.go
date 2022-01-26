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

	// Identifies the units of the current resource metric (CORES, GB).
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Capacity Data with time interval
	CapacityData []ExadataInsightResourceCapacityTrendAggregation `mandatory:"true" json:"capacityData"`
}

func (m SummarizeExadataInsightResourceCapacityTrendAggregation) String() string {
	return common.PointerString(m)
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

var mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetric = map[string]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceMetric {
		values = append(values, v)
	}
	return values
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

var mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceType = map[string]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendAggregationExadataResourceType {
		values = append(values, v)
	}
	return values
}
