// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

// ThresholdTypeEnum Enum with underlying type: string
type ThresholdTypeEnum string

// Set of constants representing the allowable values for ThresholdTypeEnum
const (
	ThresholdTypePercentage ThresholdTypeEnum = "PERCENTAGE"
	ThresholdTypeAbsolute   ThresholdTypeEnum = "ABSOLUTE"
)

var mappingThresholdType = map[string]ThresholdTypeEnum{
	"PERCENTAGE": ThresholdTypePercentage,
	"ABSOLUTE":   ThresholdTypeAbsolute,
}

// GetThresholdTypeEnumValues Enumerates the set of values for ThresholdTypeEnum
func GetThresholdTypeEnumValues() []ThresholdTypeEnum {
	values := make([]ThresholdTypeEnum, 0)
	for _, v := range mappingThresholdType {
		values = append(values, v)
	}
	return values
}
