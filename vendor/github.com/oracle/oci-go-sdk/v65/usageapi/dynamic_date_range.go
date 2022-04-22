// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DynamicDateRange The saved dynamic date range (Required when staticDateRange is missing).
type DynamicDateRange struct {
	DynamicDateRangeType DynamicDateRangeDynamicDateRangeTypeEnum `mandatory:"true" json:"dynamicDateRangeType"`
}

func (m DynamicDateRange) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicDateRange) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDynamicDateRangeDynamicDateRangeTypeEnum(string(m.DynamicDateRangeType)); !ok && m.DynamicDateRangeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DynamicDateRangeType: %s. Supported values are: %s.", m.DynamicDateRangeType, strings.Join(GetDynamicDateRangeDynamicDateRangeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DynamicDateRange) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDynamicDateRange DynamicDateRange
	s := struct {
		DiscriminatorParam string `json:"dateRangeType"`
		MarshalTypeDynamicDateRange
	}{
		"DYNAMIC",
		(MarshalTypeDynamicDateRange)(m),
	}

	return json.Marshal(&s)
}

// DynamicDateRangeDynamicDateRangeTypeEnum Enum with underlying type: string
type DynamicDateRangeDynamicDateRangeTypeEnum string

// Set of constants representing the allowable values for DynamicDateRangeDynamicDateRangeTypeEnum
const (
	DynamicDateRangeDynamicDateRangeTypeLast7Days         DynamicDateRangeDynamicDateRangeTypeEnum = "LAST_7_DAYS"
	DynamicDateRangeDynamicDateRangeTypeLastCalendarWeek  DynamicDateRangeDynamicDateRangeTypeEnum = "LAST_CALENDAR_WEEK"
	DynamicDateRangeDynamicDateRangeTypeLastCalendarMonth DynamicDateRangeDynamicDateRangeTypeEnum = "LAST_CALENDAR_MONTH"
	DynamicDateRangeDynamicDateRangeTypeLast30Days        DynamicDateRangeDynamicDateRangeTypeEnum = "LAST_30_DAYS"
	DynamicDateRangeDynamicDateRangeTypeMonthToDate       DynamicDateRangeDynamicDateRangeTypeEnum = "MONTH_TO_DATE"
	DynamicDateRangeDynamicDateRangeTypeLastYear          DynamicDateRangeDynamicDateRangeTypeEnum = "LAST_YEAR"
	DynamicDateRangeDynamicDateRangeTypeYearTodate        DynamicDateRangeDynamicDateRangeTypeEnum = "YEAR_TODATE"
)

var mappingDynamicDateRangeDynamicDateRangeTypeEnum = map[string]DynamicDateRangeDynamicDateRangeTypeEnum{
	"LAST_7_DAYS":         DynamicDateRangeDynamicDateRangeTypeLast7Days,
	"LAST_CALENDAR_WEEK":  DynamicDateRangeDynamicDateRangeTypeLastCalendarWeek,
	"LAST_CALENDAR_MONTH": DynamicDateRangeDynamicDateRangeTypeLastCalendarMonth,
	"LAST_30_DAYS":        DynamicDateRangeDynamicDateRangeTypeLast30Days,
	"MONTH_TO_DATE":       DynamicDateRangeDynamicDateRangeTypeMonthToDate,
	"LAST_YEAR":           DynamicDateRangeDynamicDateRangeTypeLastYear,
	"YEAR_TODATE":         DynamicDateRangeDynamicDateRangeTypeYearTodate,
}

var mappingDynamicDateRangeDynamicDateRangeTypeEnumLowerCase = map[string]DynamicDateRangeDynamicDateRangeTypeEnum{
	"last_7_days":         DynamicDateRangeDynamicDateRangeTypeLast7Days,
	"last_calendar_week":  DynamicDateRangeDynamicDateRangeTypeLastCalendarWeek,
	"last_calendar_month": DynamicDateRangeDynamicDateRangeTypeLastCalendarMonth,
	"last_30_days":        DynamicDateRangeDynamicDateRangeTypeLast30Days,
	"month_to_date":       DynamicDateRangeDynamicDateRangeTypeMonthToDate,
	"last_year":           DynamicDateRangeDynamicDateRangeTypeLastYear,
	"year_todate":         DynamicDateRangeDynamicDateRangeTypeYearTodate,
}

// GetDynamicDateRangeDynamicDateRangeTypeEnumValues Enumerates the set of values for DynamicDateRangeDynamicDateRangeTypeEnum
func GetDynamicDateRangeDynamicDateRangeTypeEnumValues() []DynamicDateRangeDynamicDateRangeTypeEnum {
	values := make([]DynamicDateRangeDynamicDateRangeTypeEnum, 0)
	for _, v := range mappingDynamicDateRangeDynamicDateRangeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDynamicDateRangeDynamicDateRangeTypeEnumStringValues Enumerates the set of values in String for DynamicDateRangeDynamicDateRangeTypeEnum
func GetDynamicDateRangeDynamicDateRangeTypeEnumStringValues() []string {
	return []string{
		"LAST_7_DAYS",
		"LAST_CALENDAR_WEEK",
		"LAST_CALENDAR_MONTH",
		"LAST_30_DAYS",
		"MONTH_TO_DATE",
		"LAST_YEAR",
		"YEAR_TODATE",
	}
}

// GetMappingDynamicDateRangeDynamicDateRangeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDynamicDateRangeDynamicDateRangeTypeEnum(val string) (DynamicDateRangeDynamicDateRangeTypeEnum, bool) {
	enum, ok := mappingDynamicDateRangeDynamicDateRangeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
