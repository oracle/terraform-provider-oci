// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeHostInsightResourceForecastTrendAggregation Forecast results from the selected time period.
type SummarizeHostInsightResourceForecastTrendAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"true" json:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered lowly utilized.
	LowUtilizationThreshold *int `mandatory:"true" json:"lowUtilizationThreshold"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time series patterns used in the forecasting.
	Pattern SummarizeHostInsightResourceForecastTrendAggregationPatternEnum `mandatory:"true" json:"pattern"`

	// Time series data used for the forecast analysis.
	HistoricalData []HistoricalDataItem `mandatory:"true" json:"historicalData"`

	// Time series data result of the forecasting analysis.
	ProjectedData []ProjectedDataItem `mandatory:"true" json:"projectedData"`

	// Auto-ML algorithm leveraged for the forecast. Only applicable for Auto-ML forecast.
	SelectedForecastAlgorithm *string `mandatory:"false" json:"selectedForecastAlgorithm"`
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
	if _, ok := GetMappingSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnumStringValues(), ",")))
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
	SummarizeHostInsightResourceForecastTrendAggregationResourceMetricStorage       SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = "STORAGE"
	SummarizeHostInsightResourceForecastTrendAggregationResourceMetricNetwork       SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = "NETWORK"
)

var mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum = map[string]SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceForecastTrendAggregationResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceForecastTrendAggregationResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceForecastTrendAggregationResourceMetricLogicalMemory,
	"STORAGE":        SummarizeHostInsightResourceForecastTrendAggregationResourceMetricStorage,
	"NETWORK":        SummarizeHostInsightResourceForecastTrendAggregationResourceMetricNetwork,
}

var mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumLowerCase = map[string]SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum{
	"cpu":            SummarizeHostInsightResourceForecastTrendAggregationResourceMetricCpu,
	"memory":         SummarizeHostInsightResourceForecastTrendAggregationResourceMetricMemory,
	"logical_memory": SummarizeHostInsightResourceForecastTrendAggregationResourceMetricLogicalMemory,
	"storage":        SummarizeHostInsightResourceForecastTrendAggregationResourceMetricStorage,
	"network":        SummarizeHostInsightResourceForecastTrendAggregationResourceMetricNetwork,
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
		"STORAGE",
		"NETWORK",
	}
}

// GetMappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum(val string) (SummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceForecastTrendAggregationResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum Enum with underlying type: string
type SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum
const (
	SummarizeHostInsightResourceForecastTrendAggregationUsageUnitCores   SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum = "CORES"
	SummarizeHostInsightResourceForecastTrendAggregationUsageUnitGb      SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum = "GB"
	SummarizeHostInsightResourceForecastTrendAggregationUsageUnitMbps    SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum = "MBPS"
	SummarizeHostInsightResourceForecastTrendAggregationUsageUnitIops    SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum = "IOPS"
	SummarizeHostInsightResourceForecastTrendAggregationUsageUnitPercent SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum = "PERCENT"
)

var mappingSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum = map[string]SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum{
	"CORES":   SummarizeHostInsightResourceForecastTrendAggregationUsageUnitCores,
	"GB":      SummarizeHostInsightResourceForecastTrendAggregationUsageUnitGb,
	"MBPS":    SummarizeHostInsightResourceForecastTrendAggregationUsageUnitMbps,
	"IOPS":    SummarizeHostInsightResourceForecastTrendAggregationUsageUnitIops,
	"PERCENT": SummarizeHostInsightResourceForecastTrendAggregationUsageUnitPercent,
}

var mappingSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnumLowerCase = map[string]SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum{
	"cores":   SummarizeHostInsightResourceForecastTrendAggregationUsageUnitCores,
	"gb":      SummarizeHostInsightResourceForecastTrendAggregationUsageUnitGb,
	"mbps":    SummarizeHostInsightResourceForecastTrendAggregationUsageUnitMbps,
	"iops":    SummarizeHostInsightResourceForecastTrendAggregationUsageUnitIops,
	"percent": SummarizeHostInsightResourceForecastTrendAggregationUsageUnitPercent,
}

// GetSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnumValues Enumerates the set of values for SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum
func GetSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnumValues() []SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum {
	values := make([]SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum
func GetSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum(val string) (SummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceForecastTrendAggregationUsageUnitEnumLowerCase[strings.ToLower(val)]
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

var mappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnumLowerCase = map[string]SummarizeHostInsightResourceForecastTrendAggregationPatternEnum{
	"linear":                            SummarizeHostInsightResourceForecastTrendAggregationPatternLinear,
	"monthly_seasons":                   SummarizeHostInsightResourceForecastTrendAggregationPatternMonthlySeasons,
	"monthly_and_yearly_seasons":        SummarizeHostInsightResourceForecastTrendAggregationPatternMonthlyAndYearlySeasons,
	"weekly_seasons":                    SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklySeasons,
	"weekly_and_monthly_seasons":        SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyAndMonthlySeasons,
	"weekly_monthly_and_yearly_seasons": SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyMonthlyAndYearlySeasons,
	"weekly_and_yearly_seasons":         SummarizeHostInsightResourceForecastTrendAggregationPatternWeeklyAndYearlySeasons,
	"yearly_seasons":                    SummarizeHostInsightResourceForecastTrendAggregationPatternYearlySeasons,
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
	enum, ok := mappingSummarizeHostInsightResourceForecastTrendAggregationPatternEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
