// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts. For more information, see Budgets Overview (https://docs.oracle.com/iaas/Content/Billing/Concepts/budgetsoverview.htm).
//

package budget

import (
	"strings"
)

// SortByForCostAnomalyEventEnum Enum with underlying type: string
type SortByForCostAnomalyEventEnum string

// Set of constants representing the allowable values for SortByForCostAnomalyEventEnum
const (
	SortByForCostAnomalyEventTimeAnomalyEventDate SortByForCostAnomalyEventEnum = "timeAnomalyEventDate"
	SortByForCostAnomalyEventCostAnomalyName      SortByForCostAnomalyEventEnum = "costAnomalyName"
	SortByForCostAnomalyEventId                   SortByForCostAnomalyEventEnum = "id"
)

var mappingSortByForCostAnomalyEventEnum = map[string]SortByForCostAnomalyEventEnum{
	"timeAnomalyEventDate": SortByForCostAnomalyEventTimeAnomalyEventDate,
	"costAnomalyName":      SortByForCostAnomalyEventCostAnomalyName,
	"id":                   SortByForCostAnomalyEventId,
}

var mappingSortByForCostAnomalyEventEnumLowerCase = map[string]SortByForCostAnomalyEventEnum{
	"timeanomalyeventdate": SortByForCostAnomalyEventTimeAnomalyEventDate,
	"costanomalyname":      SortByForCostAnomalyEventCostAnomalyName,
	"id":                   SortByForCostAnomalyEventId,
}

// GetSortByForCostAnomalyEventEnumValues Enumerates the set of values for SortByForCostAnomalyEventEnum
func GetSortByForCostAnomalyEventEnumValues() []SortByForCostAnomalyEventEnum {
	values := make([]SortByForCostAnomalyEventEnum, 0)
	for _, v := range mappingSortByForCostAnomalyEventEnum {
		values = append(values, v)
	}
	return values
}

// GetSortByForCostAnomalyEventEnumStringValues Enumerates the set of values in String for SortByForCostAnomalyEventEnum
func GetSortByForCostAnomalyEventEnumStringValues() []string {
	return []string{
		"timeAnomalyEventDate",
		"costAnomalyName",
		"id",
	}
}

// GetMappingSortByForCostAnomalyEventEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortByForCostAnomalyEventEnum(val string) (SortByForCostAnomalyEventEnum, bool) {
	enum, ok := mappingSortByForCostAnomalyEventEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
