// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReportQuery The request of the generated Cost Analysis report.
type ReportQuery struct {

	// Tenant ID.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The usage granularity.
	// HOURLY - Hourly data aggregation.
	// DAILY - Daily data aggregation.
	// MONTHLY - Monthly data aggregation.
	// TOTAL - Not yet supported.
	Granularity ReportQueryGranularityEnum `mandatory:"true" json:"granularity"`

	// The usage start time.
	TimeUsageStarted *common.SDKTime `mandatory:"false" json:"timeUsageStarted"`

	// The usage end time.
	TimeUsageEnded *common.SDKTime `mandatory:"false" json:"timeUsageEnded"`

	// Whether aggregated by time. If isAggregateByTime is true, all usage/cost over the query time period will be added up.
	IsAggregateByTime *bool `mandatory:"false" json:"isAggregateByTime"`

	Forecast *Forecast `mandatory:"false" json:"forecast"`

	// The query usage type. COST by default if it is missing.
	// Usage - Query the usage data.
	// Cost - Query the cost/billing data.
	// Credit - Query the credit adjustments data.
	// ExpiredCredit - Query the expired credits data
	// AllCredit - Query the credit adjustments and expired credit
	QueryType ReportQueryQueryTypeEnum `mandatory:"false" json:"queryType,omitempty"`

	// Aggregate the result by.
	// example:
	//   `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit",
	//     "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd",
	//     "resourceId", "tenantId", "tenantName"]`
	GroupBy []string `mandatory:"false" json:"groupBy"`

	// GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list.
	// For example:
	//   `[{"namespace":"oracle", "key":"createdBy"]`
	GroupByTag []Tag `mandatory:"false" json:"groupByTag"`

	// The compartment depth level.
	CompartmentDepth *float32 `mandatory:"false" json:"compartmentDepth"`

	Filter *Filter `mandatory:"false" json:"filter"`

	// The UI date range, for example, LAST_THREE_MONTHS. Conflicts with timeUsageStarted and timeUsageEnded.
	DateRangeName ReportQueryDateRangeNameEnum `mandatory:"false" json:"dateRangeName,omitempty"`
}

func (m ReportQuery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReportQuery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReportQueryGranularityEnum(string(m.Granularity)); !ok && m.Granularity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Granularity: %s. Supported values are: %s.", m.Granularity, strings.Join(GetReportQueryGranularityEnumStringValues(), ",")))
	}

	if _, ok := GetMappingReportQueryQueryTypeEnum(string(m.QueryType)); !ok && m.QueryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QueryType: %s. Supported values are: %s.", m.QueryType, strings.Join(GetReportQueryQueryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReportQueryDateRangeNameEnum(string(m.DateRangeName)); !ok && m.DateRangeName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DateRangeName: %s. Supported values are: %s.", m.DateRangeName, strings.Join(GetReportQueryDateRangeNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReportQueryGranularityEnum Enum with underlying type: string
type ReportQueryGranularityEnum string

// Set of constants representing the allowable values for ReportQueryGranularityEnum
const (
	ReportQueryGranularityHourly  ReportQueryGranularityEnum = "HOURLY"
	ReportQueryGranularityDaily   ReportQueryGranularityEnum = "DAILY"
	ReportQueryGranularityMonthly ReportQueryGranularityEnum = "MONTHLY"
	ReportQueryGranularityTotal   ReportQueryGranularityEnum = "TOTAL"
)

var mappingReportQueryGranularityEnum = map[string]ReportQueryGranularityEnum{
	"HOURLY":  ReportQueryGranularityHourly,
	"DAILY":   ReportQueryGranularityDaily,
	"MONTHLY": ReportQueryGranularityMonthly,
	"TOTAL":   ReportQueryGranularityTotal,
}

var mappingReportQueryGranularityEnumLowerCase = map[string]ReportQueryGranularityEnum{
	"hourly":  ReportQueryGranularityHourly,
	"daily":   ReportQueryGranularityDaily,
	"monthly": ReportQueryGranularityMonthly,
	"total":   ReportQueryGranularityTotal,
}

// GetReportQueryGranularityEnumValues Enumerates the set of values for ReportQueryGranularityEnum
func GetReportQueryGranularityEnumValues() []ReportQueryGranularityEnum {
	values := make([]ReportQueryGranularityEnum, 0)
	for _, v := range mappingReportQueryGranularityEnum {
		values = append(values, v)
	}
	return values
}

// GetReportQueryGranularityEnumStringValues Enumerates the set of values in String for ReportQueryGranularityEnum
func GetReportQueryGranularityEnumStringValues() []string {
	return []string{
		"HOURLY",
		"DAILY",
		"MONTHLY",
		"TOTAL",
	}
}

// GetMappingReportQueryGranularityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportQueryGranularityEnum(val string) (ReportQueryGranularityEnum, bool) {
	enum, ok := mappingReportQueryGranularityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReportQueryQueryTypeEnum Enum with underlying type: string
type ReportQueryQueryTypeEnum string

// Set of constants representing the allowable values for ReportQueryQueryTypeEnum
const (
	ReportQueryQueryTypeUsage         ReportQueryQueryTypeEnum = "USAGE"
	ReportQueryQueryTypeCost          ReportQueryQueryTypeEnum = "COST"
	ReportQueryQueryTypeCredit        ReportQueryQueryTypeEnum = "CREDIT"
	ReportQueryQueryTypeExpiredcredit ReportQueryQueryTypeEnum = "EXPIREDCREDIT"
	ReportQueryQueryTypeAllcredit     ReportQueryQueryTypeEnum = "ALLCREDIT"
)

var mappingReportQueryQueryTypeEnum = map[string]ReportQueryQueryTypeEnum{
	"USAGE":         ReportQueryQueryTypeUsage,
	"COST":          ReportQueryQueryTypeCost,
	"CREDIT":        ReportQueryQueryTypeCredit,
	"EXPIREDCREDIT": ReportQueryQueryTypeExpiredcredit,
	"ALLCREDIT":     ReportQueryQueryTypeAllcredit,
}

var mappingReportQueryQueryTypeEnumLowerCase = map[string]ReportQueryQueryTypeEnum{
	"usage":         ReportQueryQueryTypeUsage,
	"cost":          ReportQueryQueryTypeCost,
	"credit":        ReportQueryQueryTypeCredit,
	"expiredcredit": ReportQueryQueryTypeExpiredcredit,
	"allcredit":     ReportQueryQueryTypeAllcredit,
}

// GetReportQueryQueryTypeEnumValues Enumerates the set of values for ReportQueryQueryTypeEnum
func GetReportQueryQueryTypeEnumValues() []ReportQueryQueryTypeEnum {
	values := make([]ReportQueryQueryTypeEnum, 0)
	for _, v := range mappingReportQueryQueryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReportQueryQueryTypeEnumStringValues Enumerates the set of values in String for ReportQueryQueryTypeEnum
func GetReportQueryQueryTypeEnumStringValues() []string {
	return []string{
		"USAGE",
		"COST",
		"CREDIT",
		"EXPIREDCREDIT",
		"ALLCREDIT",
	}
}

// GetMappingReportQueryQueryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportQueryQueryTypeEnum(val string) (ReportQueryQueryTypeEnum, bool) {
	enum, ok := mappingReportQueryQueryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReportQueryDateRangeNameEnum Enum with underlying type: string
type ReportQueryDateRangeNameEnum string

// Set of constants representing the allowable values for ReportQueryDateRangeNameEnum
const (
	ReportQueryDateRangeNameLastSevenDays   ReportQueryDateRangeNameEnum = "LAST_SEVEN_DAYS"
	ReportQueryDateRangeNameLastTenDays     ReportQueryDateRangeNameEnum = "LAST_TEN_DAYS"
	ReportQueryDateRangeNameMtd             ReportQueryDateRangeNameEnum = "MTD"
	ReportQueryDateRangeNameLastTwoMonths   ReportQueryDateRangeNameEnum = "LAST_TWO_MONTHS"
	ReportQueryDateRangeNameLastThreeMonths ReportQueryDateRangeNameEnum = "LAST_THREE_MONTHS"
	ReportQueryDateRangeNameAll             ReportQueryDateRangeNameEnum = "ALL"
	ReportQueryDateRangeNameLastSixMonths   ReportQueryDateRangeNameEnum = "LAST_SIX_MONTHS"
	ReportQueryDateRangeNameLastOneYear     ReportQueryDateRangeNameEnum = "LAST_ONE_YEAR"
	ReportQueryDateRangeNameYtd             ReportQueryDateRangeNameEnum = "YTD"
	ReportQueryDateRangeNameCustom          ReportQueryDateRangeNameEnum = "CUSTOM"
)

var mappingReportQueryDateRangeNameEnum = map[string]ReportQueryDateRangeNameEnum{
	"LAST_SEVEN_DAYS":   ReportQueryDateRangeNameLastSevenDays,
	"LAST_TEN_DAYS":     ReportQueryDateRangeNameLastTenDays,
	"MTD":               ReportQueryDateRangeNameMtd,
	"LAST_TWO_MONTHS":   ReportQueryDateRangeNameLastTwoMonths,
	"LAST_THREE_MONTHS": ReportQueryDateRangeNameLastThreeMonths,
	"ALL":               ReportQueryDateRangeNameAll,
	"LAST_SIX_MONTHS":   ReportQueryDateRangeNameLastSixMonths,
	"LAST_ONE_YEAR":     ReportQueryDateRangeNameLastOneYear,
	"YTD":               ReportQueryDateRangeNameYtd,
	"CUSTOM":            ReportQueryDateRangeNameCustom,
}

var mappingReportQueryDateRangeNameEnumLowerCase = map[string]ReportQueryDateRangeNameEnum{
	"last_seven_days":   ReportQueryDateRangeNameLastSevenDays,
	"last_ten_days":     ReportQueryDateRangeNameLastTenDays,
	"mtd":               ReportQueryDateRangeNameMtd,
	"last_two_months":   ReportQueryDateRangeNameLastTwoMonths,
	"last_three_months": ReportQueryDateRangeNameLastThreeMonths,
	"all":               ReportQueryDateRangeNameAll,
	"last_six_months":   ReportQueryDateRangeNameLastSixMonths,
	"last_one_year":     ReportQueryDateRangeNameLastOneYear,
	"ytd":               ReportQueryDateRangeNameYtd,
	"custom":            ReportQueryDateRangeNameCustom,
}

// GetReportQueryDateRangeNameEnumValues Enumerates the set of values for ReportQueryDateRangeNameEnum
func GetReportQueryDateRangeNameEnumValues() []ReportQueryDateRangeNameEnum {
	values := make([]ReportQueryDateRangeNameEnum, 0)
	for _, v := range mappingReportQueryDateRangeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetReportQueryDateRangeNameEnumStringValues Enumerates the set of values in String for ReportQueryDateRangeNameEnum
func GetReportQueryDateRangeNameEnumStringValues() []string {
	return []string{
		"LAST_SEVEN_DAYS",
		"LAST_TEN_DAYS",
		"MTD",
		"LAST_TWO_MONTHS",
		"LAST_THREE_MONTHS",
		"ALL",
		"LAST_SIX_MONTHS",
		"LAST_ONE_YEAR",
		"YTD",
		"CUSTOM",
	}
}

// GetMappingReportQueryDateRangeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportQueryDateRangeNameEnum(val string) (ReportQueryDateRangeNameEnum, bool) {
	enum, ok := mappingReportQueryDateRangeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
