// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts. For more information, see Budgets Overview (https://docs.cloud.oracle.com/iaas/Content/Billing/Concepts/budgetsoverview.htm).
//

package budget

import (
	"strings"
)

// AlertTypeEnum Enum with underlying type: string
type AlertTypeEnum string

// Set of constants representing the allowable values for AlertTypeEnum
const (
	AlertTypeActual   AlertTypeEnum = "ACTUAL"
	AlertTypeForecast AlertTypeEnum = "FORECAST"
)

var mappingAlertTypeEnum = map[string]AlertTypeEnum{
	"ACTUAL":   AlertTypeActual,
	"FORECAST": AlertTypeForecast,
}

var mappingAlertTypeEnumLowerCase = map[string]AlertTypeEnum{
	"actual":   AlertTypeActual,
	"forecast": AlertTypeForecast,
}

// GetAlertTypeEnumValues Enumerates the set of values for AlertTypeEnum
func GetAlertTypeEnumValues() []AlertTypeEnum {
	values := make([]AlertTypeEnum, 0)
	for _, v := range mappingAlertTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertTypeEnumStringValues Enumerates the set of values in String for AlertTypeEnum
func GetAlertTypeEnumStringValues() []string {
	return []string{
		"ACTUAL",
		"FORECAST",
	}
}

// GetMappingAlertTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertTypeEnum(val string) (AlertTypeEnum, bool) {
	enum, ok := mappingAlertTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
