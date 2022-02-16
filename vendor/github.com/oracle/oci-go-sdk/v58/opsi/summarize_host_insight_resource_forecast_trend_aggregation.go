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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightResourceForecastTrendAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetUsageUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnum(string(m.Pattern)); !ok && m.Pattern != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Pattern: %s. Supported values are: %s.", m.Pattern, strings.Join(GetSummarizeHostInsightResourceForecastTrendAggregationPatternEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum
const (
	SummarizeHostInsightResourceForecastTrendAggregationResourceMetricCpu           SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceForecastTrendAggregationResourceMetricMemory        SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceForecastTrendAggregationResourceMetricLogicalMemory SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = "LOGICAL_MEMORY"
)

var mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = map[string]SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceForecastTrendAggregationResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceForecastTrendAggregationResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceForecastTrendAggregationResourceMetricLogicalMemory,
}

// GetSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumValues() []SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"LOGICAL_MEMORY",
	}
}

// GetMappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum(val string) (SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum, bool) {
	mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumIgnoreCase := make(map[string]SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum)
	for k, v := range mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum {
		mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnum = map[string]SummarizeHostInsightResourceForecastTrendAggregationPatternEnum{
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
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceForecastTrendAggregationPatternEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceForecastTrendAggregationPatternEnum
func GetSummarizeHostInsightResourceForecastTrendAggregationPatternEnumStringValues() []string {
	return []string{
		"LINEAR",
		"MONTHLY_SEASONS",
		"MONTHLY_AND_YEARLY_SEASONS",
		"WEEKLY_SEASONS",
		"WEEKLY_AND_MONTHLY_SEASONS",
		"WEEKLY_MONTHLY_AND_YEARLY_SEASONS",
		"WEEKLY_AND_YEARLY_SEASONS",
		"YEARLY_SEASONS",
	}
}

// GetMappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnum(val string) (SummarizeHostInsightResourceForecastTrendAggregationPatternEnum, bool) {
	mappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnumIgnoreCase := make(map[string]SummarizeHostInsightResourceForecastTrendAggregationPatternEnum)
	for k, v := range mappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnum {
		mappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
