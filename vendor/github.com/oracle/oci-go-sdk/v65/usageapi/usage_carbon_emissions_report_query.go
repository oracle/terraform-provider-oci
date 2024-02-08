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

// UsageCarbonEmissionsReportQuery The request of the generated usage carbon emissions report.
type UsageCarbonEmissionsReportQuery struct {

	// Tenant ID.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The usage start time.
	TimeUsageStarted *common.SDKTime `mandatory:"false" json:"timeUsageStarted"`

	// The usage end time.
	TimeUsageEnded *common.SDKTime `mandatory:"false" json:"timeUsageEnded"`

	// Specifies whether aggregated by time. If isAggregateByTime is true, all usage or cost over the query time period will be added up.
	IsAggregateByTime *bool `mandatory:"false" json:"isAggregateByTime"`

	// Specifies what to aggregate the result by.
	// For example:
	//   `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit",
	//     "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd",
	//     "resourceId", "tenantId", "tenantName"]`
	GroupBy []string `mandatory:"false" json:"groupBy"`

	// GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list.
	// For example:
	//   `[{"namespace":"oracle", "key":"createdBy"]`
	GroupByTag []Tag `mandatory:"false" json:"groupByTag"`

	// The compartment depth level.
	CompartmentDepth *int `mandatory:"false" json:"compartmentDepth"`

	Filter *Filter `mandatory:"false" json:"filter"`

	// The UI date range, for example, LAST_THREE_MONTHS. It will override timeUsageStarted and timeUsageEnded properties.
	DateRangeName UsageCarbonEmissionsReportQueryDateRangeNameEnum `mandatory:"false" json:"dateRangeName,omitempty"`
}

func (m UsageCarbonEmissionsReportQuery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageCarbonEmissionsReportQuery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUsageCarbonEmissionsReportQueryDateRangeNameEnum(string(m.DateRangeName)); !ok && m.DateRangeName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DateRangeName: %s. Supported values are: %s.", m.DateRangeName, strings.Join(GetUsageCarbonEmissionsReportQueryDateRangeNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UsageCarbonEmissionsReportQueryDateRangeNameEnum Enum with underlying type: string
type UsageCarbonEmissionsReportQueryDateRangeNameEnum string

// Set of constants representing the allowable values for UsageCarbonEmissionsReportQueryDateRangeNameEnum
const (
	UsageCarbonEmissionsReportQueryDateRangeNameLastTwoMonths   UsageCarbonEmissionsReportQueryDateRangeNameEnum = "LAST_TWO_MONTHS"
	UsageCarbonEmissionsReportQueryDateRangeNameLastThreeMonths UsageCarbonEmissionsReportQueryDateRangeNameEnum = "LAST_THREE_MONTHS"
	UsageCarbonEmissionsReportQueryDateRangeNameLastSixMonths   UsageCarbonEmissionsReportQueryDateRangeNameEnum = "LAST_SIX_MONTHS"
	UsageCarbonEmissionsReportQueryDateRangeNameLastOneYear     UsageCarbonEmissionsReportQueryDateRangeNameEnum = "LAST_ONE_YEAR"
	UsageCarbonEmissionsReportQueryDateRangeNameCustom          UsageCarbonEmissionsReportQueryDateRangeNameEnum = "CUSTOM"
)

var mappingUsageCarbonEmissionsReportQueryDateRangeNameEnum = map[string]UsageCarbonEmissionsReportQueryDateRangeNameEnum{
	"LAST_TWO_MONTHS":   UsageCarbonEmissionsReportQueryDateRangeNameLastTwoMonths,
	"LAST_THREE_MONTHS": UsageCarbonEmissionsReportQueryDateRangeNameLastThreeMonths,
	"LAST_SIX_MONTHS":   UsageCarbonEmissionsReportQueryDateRangeNameLastSixMonths,
	"LAST_ONE_YEAR":     UsageCarbonEmissionsReportQueryDateRangeNameLastOneYear,
	"CUSTOM":            UsageCarbonEmissionsReportQueryDateRangeNameCustom,
}

var mappingUsageCarbonEmissionsReportQueryDateRangeNameEnumLowerCase = map[string]UsageCarbonEmissionsReportQueryDateRangeNameEnum{
	"last_two_months":   UsageCarbonEmissionsReportQueryDateRangeNameLastTwoMonths,
	"last_three_months": UsageCarbonEmissionsReportQueryDateRangeNameLastThreeMonths,
	"last_six_months":   UsageCarbonEmissionsReportQueryDateRangeNameLastSixMonths,
	"last_one_year":     UsageCarbonEmissionsReportQueryDateRangeNameLastOneYear,
	"custom":            UsageCarbonEmissionsReportQueryDateRangeNameCustom,
}

// GetUsageCarbonEmissionsReportQueryDateRangeNameEnumValues Enumerates the set of values for UsageCarbonEmissionsReportQueryDateRangeNameEnum
func GetUsageCarbonEmissionsReportQueryDateRangeNameEnumValues() []UsageCarbonEmissionsReportQueryDateRangeNameEnum {
	values := make([]UsageCarbonEmissionsReportQueryDateRangeNameEnum, 0)
	for _, v := range mappingUsageCarbonEmissionsReportQueryDateRangeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageCarbonEmissionsReportQueryDateRangeNameEnumStringValues Enumerates the set of values in String for UsageCarbonEmissionsReportQueryDateRangeNameEnum
func GetUsageCarbonEmissionsReportQueryDateRangeNameEnumStringValues() []string {
	return []string{
		"LAST_TWO_MONTHS",
		"LAST_THREE_MONTHS",
		"LAST_SIX_MONTHS",
		"LAST_ONE_YEAR",
		"CUSTOM",
	}
}

// GetMappingUsageCarbonEmissionsReportQueryDateRangeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageCarbonEmissionsReportQueryDateRangeNameEnum(val string) (UsageCarbonEmissionsReportQueryDateRangeNameEnum, bool) {
	enum, ok := mappingUsageCarbonEmissionsReportQueryDateRangeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
