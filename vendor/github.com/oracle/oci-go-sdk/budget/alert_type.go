// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

// AlertTypeEnum Enum with underlying type: string
type AlertTypeEnum string

// Set of constants representing the allowable values for AlertTypeEnum
const (
	AlertTypeActual   AlertTypeEnum = "ACTUAL"
	AlertTypeForecast AlertTypeEnum = "FORECAST"
)

var mappingAlertType = map[string]AlertTypeEnum{
	"ACTUAL":   AlertTypeActual,
	"FORECAST": AlertTypeForecast,
}

// GetAlertTypeEnumValues Enumerates the set of values for AlertTypeEnum
func GetAlertTypeEnumValues() []AlertTypeEnum {
	values := make([]AlertTypeEnum, 0)
	for _, v := range mappingAlertType {
		values = append(values, v)
	}
	return values
}
