// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

import (
	"strings"
)

// ResetPeriodEnum Enum with underlying type: string
type ResetPeriodEnum string

// Set of constants representing the allowable values for ResetPeriodEnum
const (
	ResetPeriodMonthly ResetPeriodEnum = "MONTHLY"
)

var mappingResetPeriodEnum = map[string]ResetPeriodEnum{
	"MONTHLY": ResetPeriodMonthly,
}

// GetResetPeriodEnumValues Enumerates the set of values for ResetPeriodEnum
func GetResetPeriodEnumValues() []ResetPeriodEnum {
	values := make([]ResetPeriodEnum, 0)
	for _, v := range mappingResetPeriodEnum {
		values = append(values, v)
	}
	return values
}

// GetResetPeriodEnumStringValues Enumerates the set of values in String for ResetPeriodEnum
func GetResetPeriodEnumStringValues() []string {
	return []string{
		"MONTHLY",
	}
}

// GetMappingResetPeriodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResetPeriodEnum(val string) (ResetPeriodEnum, bool) {
	mappingResetPeriodEnumIgnoreCase := make(map[string]ResetPeriodEnum)
	for k, v := range mappingResetPeriodEnum {
		mappingResetPeriodEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingResetPeriodEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
