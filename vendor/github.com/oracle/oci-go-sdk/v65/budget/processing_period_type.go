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

// ProcessingPeriodTypeEnum Enum with underlying type: string
type ProcessingPeriodTypeEnum string

// Set of constants representing the allowable values for ProcessingPeriodTypeEnum
const (
	ProcessingPeriodTypeInvoice   ProcessingPeriodTypeEnum = "INVOICE"
	ProcessingPeriodTypeMonth     ProcessingPeriodTypeEnum = "MONTH"
	ProcessingPeriodTypeSingleUse ProcessingPeriodTypeEnum = "SINGLE_USE"
)

var mappingProcessingPeriodTypeEnum = map[string]ProcessingPeriodTypeEnum{
	"INVOICE":    ProcessingPeriodTypeInvoice,
	"MONTH":      ProcessingPeriodTypeMonth,
	"SINGLE_USE": ProcessingPeriodTypeSingleUse,
}

var mappingProcessingPeriodTypeEnumLowerCase = map[string]ProcessingPeriodTypeEnum{
	"invoice":    ProcessingPeriodTypeInvoice,
	"month":      ProcessingPeriodTypeMonth,
	"single_use": ProcessingPeriodTypeSingleUse,
}

// GetProcessingPeriodTypeEnumValues Enumerates the set of values for ProcessingPeriodTypeEnum
func GetProcessingPeriodTypeEnumValues() []ProcessingPeriodTypeEnum {
	values := make([]ProcessingPeriodTypeEnum, 0)
	for _, v := range mappingProcessingPeriodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProcessingPeriodTypeEnumStringValues Enumerates the set of values in String for ProcessingPeriodTypeEnum
func GetProcessingPeriodTypeEnumStringValues() []string {
	return []string{
		"INVOICE",
		"MONTH",
		"SINGLE_USE",
	}
}

// GetMappingProcessingPeriodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProcessingPeriodTypeEnum(val string) (ProcessingPeriodTypeEnum, bool) {
	enum, ok := mappingProcessingPeriodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
