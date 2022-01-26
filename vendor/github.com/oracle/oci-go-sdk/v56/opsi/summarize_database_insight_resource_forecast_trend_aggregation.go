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

// SummarizeDatabaseInsightResourceForecastTrendAggregation Forecast results from the selected time period.
type SummarizeDatabaseInsightResourceForecastTrendAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (example: CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time series patterns used in the forecasting.
	Pattern SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum `mandatory:"true" json:"pattern"`

	// The name of tablespace.
	TablespaceName *string `mandatory:"true" json:"tablespaceName"`

	// Time series data used for the forecast analysis.
	HistoricalData []HistoricalDataItem `mandatory:"true" json:"historicalData"`

	// Time series data result of the forecasting analysis.
	ProjectedData []ProjectedDataItem `mandatory:"true" json:"projectedData"`
}

func (m SummarizeDatabaseInsightResourceForecastTrendAggregation) String() string {
	return common.PointerString(m)
}

// SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricCpu       SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricStorage   SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricIo        SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricMemory    SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum = "MEMORY"
	SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricMemoryPga SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum = "MEMORY_PGA"
	SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricMemorySga SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum = "MEMORY_SGA"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetric = map[string]SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum{
	"CPU":        SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricCpu,
	"STORAGE":    SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricStorage,
	"IO":         SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricIo,
	"MEMORY":     SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricMemory,
	"MEMORY_PGA": SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricMemoryPga,
	"MEMORY_SGA": SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricMemorySga,
}

// GetSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum
func GetSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnumValues() []SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetric {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum
const (
	SummarizeDatabaseInsightResourceForecastTrendAggregationPatternLinear                        SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = "LINEAR"
	SummarizeDatabaseInsightResourceForecastTrendAggregationPatternMonthlySeasons                SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = "MONTHLY_SEASONS"
	SummarizeDatabaseInsightResourceForecastTrendAggregationPatternMonthlyAndYearlySeasons       SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = "MONTHLY_AND_YEARLY_SEASONS"
	SummarizeDatabaseInsightResourceForecastTrendAggregationPatternWeeklySeasons                 SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_SEASONS"
	SummarizeDatabaseInsightResourceForecastTrendAggregationPatternWeeklyAndMonthlySeasons       SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_AND_MONTHLY_SEASONS"
	SummarizeDatabaseInsightResourceForecastTrendAggregationPatternWeeklyMonthlyAndYearlySeasons SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_MONTHLY_AND_YEARLY_SEASONS"
	SummarizeDatabaseInsightResourceForecastTrendAggregationPatternWeeklyAndYearlySeasons        SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_AND_YEARLY_SEASONS"
	SummarizeDatabaseInsightResourceForecastTrendAggregationPatternYearlySeasons                 SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = "YEARLY_SEASONS"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendAggregationPattern = map[string]SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum{
	"LINEAR":                            SummarizeDatabaseInsightResourceForecastTrendAggregationPatternLinear,
	"MONTHLY_SEASONS":                   SummarizeDatabaseInsightResourceForecastTrendAggregationPatternMonthlySeasons,
	"MONTHLY_AND_YEARLY_SEASONS":        SummarizeDatabaseInsightResourceForecastTrendAggregationPatternMonthlyAndYearlySeasons,
	"WEEKLY_SEASONS":                    SummarizeDatabaseInsightResourceForecastTrendAggregationPatternWeeklySeasons,
	"WEEKLY_AND_MONTHLY_SEASONS":        SummarizeDatabaseInsightResourceForecastTrendAggregationPatternWeeklyAndMonthlySeasons,
	"WEEKLY_MONTHLY_AND_YEARLY_SEASONS": SummarizeDatabaseInsightResourceForecastTrendAggregationPatternWeeklyMonthlyAndYearlySeasons,
	"WEEKLY_AND_YEARLY_SEASONS":         SummarizeDatabaseInsightResourceForecastTrendAggregationPatternWeeklyAndYearlySeasons,
	"YEARLY_SEASONS":                    SummarizeDatabaseInsightResourceForecastTrendAggregationPatternYearlySeasons,
}

// GetSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum
func GetSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnumValues() []SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendAggregationPattern {
		values = append(values, v)
	}
	return values
}
