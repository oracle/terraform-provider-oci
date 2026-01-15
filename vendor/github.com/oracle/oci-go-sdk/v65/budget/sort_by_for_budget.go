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

// SortByForBudgetEnum Enum with underlying type: string
type SortByForBudgetEnum string

// Set of constants representing the allowable values for SortByForBudgetEnum
const (
	SortByForBudgetTimeCreated SortByForBudgetEnum = "timeCreated"
	SortByForBudgetDisplayName SortByForBudgetEnum = "displayName"
)

var mappingSortByForBudgetEnum = map[string]SortByForBudgetEnum{
	"timeCreated": SortByForBudgetTimeCreated,
	"displayName": SortByForBudgetDisplayName,
}

var mappingSortByForBudgetEnumLowerCase = map[string]SortByForBudgetEnum{
	"timecreated": SortByForBudgetTimeCreated,
	"displayname": SortByForBudgetDisplayName,
}

// GetSortByForBudgetEnumValues Enumerates the set of values for SortByForBudgetEnum
func GetSortByForBudgetEnumValues() []SortByForBudgetEnum {
	values := make([]SortByForBudgetEnum, 0)
	for _, v := range mappingSortByForBudgetEnum {
		values = append(values, v)
	}
	return values
}

// GetSortByForBudgetEnumStringValues Enumerates the set of values in String for SortByForBudgetEnum
func GetSortByForBudgetEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingSortByForBudgetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortByForBudgetEnum(val string) (SortByForBudgetEnum, bool) {
	enum, ok := mappingSortByForBudgetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
