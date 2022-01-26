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

// SummarizeHostInsightResourceForecastTrendAggregation Forecast results from the selected time period.
type SummarizeHostInsightResourceForecastTrendAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit (CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time series patterns used in the forecasting.
	Pattern SummarizeHostInsightResourceForecastTrendAggregationPatternEnum `mandatory:"true" json:"pattern"`

	// Time series data used for the forecast analysis.
	HistoricalData []HistoricalDataItem `mandatory:"true" json:"historicalData"`

	// Time series data result of the forecasting analysis.
	ProjectedData []ProjectedDataItem `mandatory:"true" json:"projectedData"`
}

func (m SummarizeHostInsightResourceForecastTrendAggregation) String() string {
	return common.PointerString(m)
}

// SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum
const (
	SummarizeHostInsightResourceForecastTrendAggregationResourceMetricCpu           SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceForecastTrendAggregationResourceMetricMemory        SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceForecastTrendAggregationResourceMetricLogicalMemory SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = "LOGICAL_MEMORY"
)

var mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetric = map[string]SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceForecastTrendAggregationResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceForecastTrendAggregationResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceForecastTrendAggregationResourceMetricLogicalMemory,
}

// GetSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumValues() []SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetric {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceForecastTrendAggregationPatternEnum Enum with underlying type: string
type SummarizeHostInsightResourceForecastTrendAggregationPatternEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceForecastTrendAggregationPatternEnum
const (
	SummarizeHostInsightResourceForecastTrendAggregationPatternLinear                        SummarizeHostInsightResourceForecastTrendAggregationPatternEnum = "LINEAR"
	SummarizeHostInsightResourceForecastTrendAggregationPatternMonthlySeasons                SummarizeHostInsightResourceForecastTrendAggregationPatternEnum = "MONTHLY_SEASONS"
	SummarizeHostInsightResourceForecastTrendAggregationPatternMonthlyAndYearlySeasons       SummarizeHostInsightResourceForecastTrendAggregationPatternEnum = "MONTHLY_AND_YEARLY_SEASONS"
	SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklySeasons                 SummarizeHostInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_SEASONS"
	SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyAndMonthlySeasons       SummarizeHostInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_AND_MONTHLY_SEASONS"
	SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyMonthlyAndYearlySeasons SummarizeHostInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_MONTHLY_AND_YEARLY_SEASONS"
	SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyAndYearlySeasons        SummarizeHostInsightResourceForecastTrendAggregationPatternEnum = "WEEKLY_AND_YEARLY_SEASONS"
	SummarizeHostInsightResourceForecastTrendAggregationPatternYearlySeasons                 SummarizeHostInsightResourceForecastTrendAggregationPatternEnum = "YEARLY_SEASONS"
)

var mappingSummarizeHostInsightResourceForecastTrendAggregationPattern = map[string]SummarizeHostInsightResourceForecastTrendAggregationPatternEnum{
	"LINEAR":                            SummarizeHostInsightResourceForecastTrendAggregationPatternLinear,
	"MONTHLY_SEASONS":                   SummarizeHostInsightResourceForecastTrendAggregationPatternMonthlySeasons,
	"MONTHLY_AND_YEARLY_SEASONS":        SummarizeHostInsightResourceForecastTrendAggregationPatternMonthlyAndYearlySeasons,
	"WEEKLY_SEASONS":                    SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklySeasons,
	"WEEKLY_AND_MONTHLY_SEASONS":        SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyAndMonthlySeasons,
	"WEEKLY_MONTHLY_AND_YEARLY_SEASONS": SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyMonthlyAndYearlySeasons,
	"WEEKLY_AND_YEARLY_SEASONS":         SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyAndYearlySeasons,
	"YEARLY_SEASONS":                    SummarizeHostInsightResourceForecastTrendAggregationPatternYearlySeasons,
}

// GetSummarizeHostInsightResourceForecastTrendAggregationPatternEnumValues Enumerates the set of values for SummarizeHostInsightResourceForecastTrendAggregationPatternEnum
func GetSummarizeHostInsightResourceForecastTrendAggregationPatternEnumValues() []SummarizeHostInsightResourceForecastTrendAggregationPatternEnum {
	values := make([]SummarizeHostInsightResourceForecastTrendAggregationPatternEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendAggregationPattern {
		values = append(values, v)
	}
	return values
}
