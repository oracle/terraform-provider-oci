// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

// ResetPeriodEnum Enum with underlying type: string
type ResetPeriodEnum string

// Set of constants representing the allowable values for ResetPeriodEnum
const (
	ResetPeriodMonthly ResetPeriodEnum = "MONTHLY"
)

var mappingResetPeriod = map[string]ResetPeriodEnum{
	"MONTHLY": ResetPeriodMonthly,
}

// GetResetPeriodEnumValues Enumerates the set of values for ResetPeriodEnum
func GetResetPeriodEnumValues() []ResetPeriodEnum {
	values := make([]ResetPeriodEnum, 0)
	for _, v := range mappingResetPeriod {
		values = append(values, v)
	}
	return values
}
