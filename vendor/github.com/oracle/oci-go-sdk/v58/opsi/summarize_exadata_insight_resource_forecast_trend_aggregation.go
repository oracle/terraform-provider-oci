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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeExadataInsightResourceForecastTrendAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum(string(m.ExadataResourceMetric)); !ok && m.ExadataResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceMetric: %s. Supported values are: %s.", m.ExadataResourceMetric, strings.Join(GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum(string(m.ExadataResourceType)); !ok && m.ExadataResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceType: %s. Supported values are: %s.", m.ExadataResourceType, strings.Join(GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetUsageUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnum(string(m.Pattern)); !ok && m.Pattern != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Pattern: %s. Supported values are: %s.", m.Pattern, strings.Join(GetSummarizeExadataInsightResourceForecastTrendAggregationPatternEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum = map[string]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum{
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
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum(val string) (SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum, bool) {
	mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum)
	for k, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnum {
		mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceMetricEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum = map[string]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"HOST",
		"STORAGE_SERVER",
		"DISKGROUP",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum(val string) (SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum, bool) {
	mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum)
	for k, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnum {
		mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendAggregationExadataResourceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnum = map[string]SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum{
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
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendAggregationPatternEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregationPatternEnumStringValues() []string {
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

// GetMappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnum(val string) (SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum, bool) {
	mappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceForecastTrendAggregationPatternEnum)
	for k, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnum {
		mappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendAggregationPatternEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
