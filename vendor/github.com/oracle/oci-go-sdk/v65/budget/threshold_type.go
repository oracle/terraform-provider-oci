// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ThresholdTypeEnum Enum with underlying type: string
type ThresholdTypeEnum string

// Set of constants representing the allowable values for ThresholdTypeEnum
const (
	ThresholdTypePercentage ThresholdTypeEnum = "PERCENTAGE"
	ThresholdTypeAbsolute   ThresholdTypeEnum = "ABSOLUTE"
)

var mappingThresholdTypeEnum = map[string]ThresholdTypeEnum{
	"PERCENTAGE": ThresholdTypePercentage,
	"ABSOLUTE":   ThresholdTypeAbsolute,
}

var mappingThresholdTypeEnumLowerCase = map[string]ThresholdTypeEnum{
	"percentage": ThresholdTypePercentage,
	"absolute":   ThresholdTypeAbsolute,
}

// GetThresholdTypeEnumValues Enumerates the set of values for ThresholdTypeEnum
func GetThresholdTypeEnumValues() []ThresholdTypeEnum {
	values := make([]ThresholdTypeEnum, 0)
	for _, v := range mappingThresholdTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetThresholdTypeEnumStringValues Enumerates the set of values in String for ThresholdTypeEnum
func GetThresholdTypeEnumStringValues() []string {
	return []string{
		"PERCENTAGE",
		"ABSOLUTE",
	}
}

// GetMappingThresholdTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingThresholdTypeEnum(val string) (ThresholdTypeEnum, bool) {
	enum, ok := mappingThresholdTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
