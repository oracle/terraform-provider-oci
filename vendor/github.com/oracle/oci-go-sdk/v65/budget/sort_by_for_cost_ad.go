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

// SortByForCostAdEnum Enum with underlying type: string
type SortByForCostAdEnum string

// Set of constants representing the allowable values for SortByForCostAdEnum
const (
	SortByForCostAdTimeCreated SortByForCostAdEnum = "timeCreated"
	SortByForCostAdName        SortByForCostAdEnum = "name"
	SortByForCostAdId          SortByForCostAdEnum = "id"
)

var mappingSortByForCostAdEnum = map[string]SortByForCostAdEnum{
	"timeCreated": SortByForCostAdTimeCreated,
	"name":        SortByForCostAdName,
	"id":          SortByForCostAdId,
}

var mappingSortByForCostAdEnumLowerCase = map[string]SortByForCostAdEnum{
	"timecreated": SortByForCostAdTimeCreated,
	"name":        SortByForCostAdName,
	"id":          SortByForCostAdId,
}

// GetSortByForCostAdEnumValues Enumerates the set of values for SortByForCostAdEnum
func GetSortByForCostAdEnumValues() []SortByForCostAdEnum {
	values := make([]SortByForCostAdEnum, 0)
	for _, v := range mappingSortByForCostAdEnum {
		values = append(values, v)
	}
	return values
}

// GetSortByForCostAdEnumStringValues Enumerates the set of values in String for SortByForCostAdEnum
func GetSortByForCostAdEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"id",
	}
}

// GetMappingSortByForCostAdEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortByForCostAdEnum(val string) (SortByForCostAdEnum, bool) {
	enum, ok := mappingSortByForCostAdEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
