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

// SummarizeExadataInsightResourceForecastTrendCollection Usage and Forecast results with breakdown by databases, hosts or storage servers.
type SummarizeExadataInsightResourceForecastTrendCollection struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
	ExadataInsightId *string `mandatory:"true" json:"exadataInsightId"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Displays usage unit ( CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of id, name , daysToReach Capacity, historical usage and projected usage forecast.
	Items []ExadataInsightResourceForecastTrendSummary `mandatory:"true" json:"items"`
}

func (m SummarizeExadataInsightResourceForecastTrendCollection) String() string {
	return common.PointerString(m)
}

// SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDatabase      SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeHost          SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeStorageServer SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceType = map[string]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceType {
		values = append(values, v)
	}
	return values
}

// SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricCpu        SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricStorage    SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIo         SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricMemory     SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIops       SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricThroughput SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetric = map[string]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetric {
		values = append(values, v)
	}
	return values
}
