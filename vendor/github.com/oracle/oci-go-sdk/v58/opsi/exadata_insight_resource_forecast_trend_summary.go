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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInsightResourceForecastTrendSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadataInsightResourceForecastTrendSummaryPatternEnum(string(m.Pattern)); !ok && m.Pattern != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Pattern: %s. Supported values are: %s.", m.Pattern, strings.Join(GetExadataInsightResourceForecastTrendSummaryPatternEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingExadataInsightResourceForecastTrendSummaryPatternEnum = map[string]ExadataInsightResourceForecastTrendSummaryPatternEnum{
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
	for _, v := range mappingExadataInsightResourceForecastTrendSummaryPatternEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInsightResourceForecastTrendSummaryPatternEnumStringValues Enumerates the set of values in String for ExadataInsightResourceForecastTrendSummaryPatternEnum
func GetExadataInsightResourceForecastTrendSummaryPatternEnumStringValues() []string {
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

// GetMappingExadataInsightResourceForecastTrendSummaryPatternEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInsightResourceForecastTrendSummaryPatternEnum(val string) (ExadataInsightResourceForecastTrendSummaryPatternEnum, bool) {
	mappingExadataInsightResourceForecastTrendSummaryPatternEnumIgnoreCase := make(map[string]ExadataInsightResourceForecastTrendSummaryPatternEnum)
	for k, v := range mappingExadataInsightResourceForecastTrendSummaryPatternEnum {
		mappingExadataInsightResourceForecastTrendSummaryPatternEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingExadataInsightResourceForecastTrendSummaryPatternEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
