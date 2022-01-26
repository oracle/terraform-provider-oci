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

// ExadataInsightResourceForecastTrendSummary List of resource id, name , capacity insight value, pattern, historical usage and projected data.
type ExadataInsightResourceForecastTrendSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The name of the resource.
	Name *string `mandatory:"true" json:"name"`

	// Days to reach capacity for a storage server
	DaysToReachCapacity *int `mandatory:"true" json:"daysToReachCapacity"`

	// Time series patterns used in the forecasting.
	Pattern ExadataInsightResourceForecastTrendSummaryPatternEnum `mandatory:"true" json:"pattern"`

	// Time series data used for the forecast analysis.
	HistoricalData []HistoricalDataItem `mandatory:"true" json:"historicalData"`

	// Time series data result of the forecasting analysis.
	ProjectedData []ProjectedDataItem `mandatory:"true" json:"projectedData"`
}

func (m ExadataInsightResourceForecastTrendSummary) String() string {
	return common.PointerString(m)
}

// ExadataInsightResourceForecastTrendSummaryPatternEnum Enum with underlying type: string
type ExadataInsightResourceForecastTrendSummaryPatternEnum string

// Set of constants representing the allowable values for ExadataInsightResourceForecastTrendSummaryPatternEnum
const (
	ExadataInsightResourceForecastTrendSummaryPatternLinear                        ExadataInsightResourceForecastTrendSummaryPatternEnum = "LINEAR"
	ExadataInsightResourceForecastTrendSummaryPatternMonthlySeasons                ExadataInsightResourceForecastTrendSummaryPatternEnum = "MONTHLY_SEASONS"
	ExadataInsightResourceForecastTrendSummaryPatternMonthlyAndYearlySeasons       ExadataInsightResourceForecastTrendSummaryPatternEnum = "MONTHLY_AND_YEARLY_SEASONS"
	ExadataInsightResourceForecastTrendSummaryPatternWeeklySeasons                 ExadataInsightResourceForecastTrendSummaryPatternEnum = "WEEKLY_SEASONS"
	ExadataInsightResourceForecastTrendSummaryPatternWeeklyAndMonthlySeasons       ExadataInsightResourceForecastTrendSummaryPatternEnum = "WEEKLY_AND_MONTHLY_SEASONS"
	ExadataInsightResourceForecastTrendSummaryPatternWeeklyMonthlyAndYearlySeasons ExadataInsightResourceForecastTrendSummaryPatternEnum = "WEEKLY_MONTHLY_AND_YEARLY_SEASONS"
	ExadataInsightResourceForecastTrendSummaryPatternWeeklyAndYearlySeasons        ExadataInsightResourceForecastTrendSummaryPatternEnum = "WEEKLY_AND_YEARLY_SEASONS"
	ExadataInsightResourceForecastTrendSummaryPatternYearlySeasons                 ExadataInsightResourceForecastTrendSummaryPatternEnum = "YEARLY_SEASONS"
)

var mappingExadataInsightResourceForecastTrendSummaryPattern = map[string]ExadataInsightResourceForecastTrendSummaryPatternEnum{
	"LINEAR":                            ExadataInsightResourceForecastTrendSummaryPatternLinear,
	"MONTHLY_SEASONS":                   ExadataInsightResourceForecastTrendSummaryPatternMonthlySeasons,
	"MONTHLY_AND_YEARLY_SEASONS":        ExadataInsightResourceForecastTrendSummaryPatternMonthlyAndYearlySeasons,
	"WEEKLY_SEASONS":                    ExadataInsightResourceForecastTrendSummaryPatternWeeklySeasons,
	"WEEKLY_AND_MONTHLY_SEASONS":        ExadataInsightResourceForecastTrendSummaryPatternWeeklyAndMonthlySeasons,
	"WEEKLY_MONTHLY_AND_YEARLY_SEASONS": ExadataInsightResourceForecastTrendSummaryPatternWeeklyMonthlyAndYearlySeasons,
	"WEEKLY_AND_YEARLY_SEASONS":         ExadataInsightResourceForecastTrendSummaryPatternWeeklyAndYearlySeasons,
	"YEARLY_SEASONS":                    ExadataInsightResourceForecastTrendSummaryPatternYearlySeasons,
}

// GetExadataInsightResourceForecastTrendSummaryPatternEnumValues Enumerates the set of values for ExadataInsightResourceForecastTrendSummaryPatternEnum
func GetExadataInsightResourceForecastTrendSummaryPatternEnumValues() []ExadataInsightResourceForecastTrendSummaryPatternEnum {
	values := make([]ExadataInsightResourceForecastTrendSummaryPatternEnum, 0)
	for _, v := range mappingExadataInsightResourceForecastTrendSummaryPattern {
		values = append(values, v)
	}
	return values
}
