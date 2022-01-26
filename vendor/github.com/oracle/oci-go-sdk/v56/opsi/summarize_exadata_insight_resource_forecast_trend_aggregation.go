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

// SummarizeExadataInsightResourceForecastTrendAggregation Usage and Forecast results from the selected time period.
type SummarizeExadataInsightResourceForecastTrendAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Displays usage unit ( CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time series patterns used in the forecasting.
	Pattern SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum `mandatory:"true" json:"pattern"`

	// Days to reach capacity for a storage server
	DaysToReachCapacity *int `mandatory:"true" json:"daysToReachCapacity"`

	// Time series data used for the forecast analysis.
	HistoricalData []HistoricalDataItem `mandatory:"true" json:"historicalData"`

	// Time series data result of the forecasting analysis.
	ProjectedData []ProjectedDataItem `mandatory:"true" json:"projectedData"`
}

func (m SummarizeExadataInsightResourceForecastTrendAggregation) String() string {
	return common.PointerString(m)
}

// SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricCpu        SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricStorage    SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricIo         SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricMemory     SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricIops       SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricThroughput SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetric = map[string]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetric {
		values = append(values, v)
	}
	return values
}

// SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeDatabase      SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeHost          SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeStorageServer SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceType = map[string]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceType {
		values = append(values, v)
	}
	return values
}

// SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum
const (
	SummarizeExadataInsightResourceForecastTrendAggregationPatternLinear                        SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = "LINEAR"
	SummarizeExadataInsightResourceForecastTrendAggregationPatternMonthlySeasons                SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = "MONTHLY_SEASONS"
	SummarizeExadataInsightResourceForecastTrendAggregationPatternMonthlyAndYearlySeasons       SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = "MONTHLY_AND_YEARLY_SEASONS"
	SummarizeExadataInsightResourceForecastTrendAggregationPatternWeeklySeasons                 SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_SEASONS"
	SummarizeExadataInsightResourceForecastTrendAggregationPatternWeeklyAndMonthlySeasons       SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_AND_MONTHLY_SEASONS"
	SummarizeExadataInsightResourceForecastTrendAggregationPatternWeeklyMonthlyAndYearlySeasons SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_MONTHLY_AND_YEARLY_SEASONS"
	SummarizeExadataInsightResourceForecastTrendAggregationPatternWeeklyAndYearlySeasons        SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_AND_YEARLY_SEASONS"
	SummarizeExadataInsightResourceForecastTrendAggregationPatternYearlySeasons                 SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = "YEARLY_SEASONS"
)

var mappingSummarizeExadataInsightResourceForecastTrendAggregationPattern = map[string]SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum{
	"LINEAR":                            SummarizeExadataInsightResourceForecastTrendAggregationPatternLinear,
	"MONTHLY_SEASONS":                   SummarizeExadataInsightResourceForecastTrendAggregationPatternMonthlySeasons,
	"MONTHLY_AND_YEARLY_SEASONS":        SummarizeExadataInsightResourceForecastTrendAggregationPatternMonthlyAndYearlySeasons,
	"WEEKLY_SEASONS":                    SummarizeExadataInsightResourceForecastTrendAggregationPatternWeeklySeasons,
	"WEEKLY_AND_MONTHLY_SEASONS":        SummarizeExadataInsightResourceForecastTrendAggregationPatternWeeklyAndMonthlySeasons,
	"WEEKLY_MONTHLY_AND_YEARLY_SEASONS": SummarizeExadataInsightResourceForecastTrendAggregationPatternWeeklyMonthlyAndYearlySeasons,
	"WEEKLY_AND_YEARLY_SEASONS":         SummarizeExadataInsightResourceForecastTrendAggregationPatternWeeklyAndYearlySeasons,
	"YEARLY_SEASONS":                    SummarizeExadataInsightResourceForecastTrendAggregationPatternYearlySeasons,
}

// GetSummarizeExadataInsightResourceForecastTrendAggregationPatternEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregationPatternEnumValues() []SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationPattern {
		values = append(values, v)
	}
	return values
}
