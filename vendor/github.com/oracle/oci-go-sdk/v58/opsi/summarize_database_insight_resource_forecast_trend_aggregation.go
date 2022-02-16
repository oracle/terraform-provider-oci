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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeDatabaseInsightResourceForecastTrendAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetUsageUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum(string(m.Pattern)); !ok && m.Pattern != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Pattern: %s. Supported values are: %s.", m.Pattern, strings.Join(GetSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum = map[string]SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum{
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
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum
func GetSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"MEMORY_PGA",
		"MEMORY_SGA",
	}
}

// GetMappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum(val string) (SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum, bool) {
	mappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnumIgnoreCase := make(map[string]SummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum)
	for k, v := range mappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnum {
		mappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeDatabaseInsightResourceForecastTrendAggregationResourceMetricEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum = map[string]SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum{
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
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum
func GetSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnumStringValues() []string {
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

// GetMappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum(val string) (SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum, bool) {
	mappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnumIgnoreCase := make(map[string]SummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum)
	for k, v := range mappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnum {
		mappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeDatabaseInsightResourceForecastTrendAggregationPatternEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
