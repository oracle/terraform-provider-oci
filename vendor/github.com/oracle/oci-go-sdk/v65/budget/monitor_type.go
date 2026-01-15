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

// MonitorTypeEnum Enum with underlying type: string
type MonitorTypeEnum string

// Set of constants representing the allowable values for MonitorTypeEnum
const (
	MonitorTypeDefault MonitorTypeEnum = "DEFAULT"
	MonitorTypeCustom  MonitorTypeEnum = "CUSTOM"
)

var mappingMonitorTypeEnum = map[string]MonitorTypeEnum{
	"DEFAULT": MonitorTypeDefault,
	"CUSTOM":  MonitorTypeCustom,
}

var mappingMonitorTypeEnumLowerCase = map[string]MonitorTypeEnum{
	"default": MonitorTypeDefault,
	"custom":  MonitorTypeCustom,
}

// GetMonitorTypeEnumValues Enumerates the set of values for MonitorTypeEnum
func GetMonitorTypeEnumValues() []MonitorTypeEnum {
	values := make([]MonitorTypeEnum, 0)
	for _, v := range mappingMonitorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitorTypeEnumStringValues Enumerates the set of values in String for MonitorTypeEnum
func GetMonitorTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM",
	}
}

// GetMappingMonitorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitorTypeEnum(val string) (MonitorTypeEnum, bool) {
	enum, ok := mappingMonitorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
