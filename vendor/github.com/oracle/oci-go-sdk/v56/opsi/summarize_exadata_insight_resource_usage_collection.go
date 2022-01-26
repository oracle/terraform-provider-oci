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

// SummarizeExadataInsightResourceUsageCollection Resource usage , allocation, utilization and usage ChangePercent for the current time period
type SummarizeExadataInsightResourceUsageCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Displays usage unit (CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of Resource Usage Summary items
	Items []ResourceUsageSummary `mandatory:"true" json:"items"`
}

func (m SummarizeExadataInsightResourceUsageCollection) String() string {
	return common.PointerString(m)
}

// SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricCpu        SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricStorage    SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIo         SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricMemory     SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIops       SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricThroughput SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetric = map[string]SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetric {
		values = append(values, v)
	}
	return values
}

// SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDatabase      SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeHost          SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeStorageServer SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceType = map[string]SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceType {
		values = append(values, v)
	}
	return values
}
