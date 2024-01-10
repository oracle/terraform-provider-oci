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

// TargetTypeEnum Enum with underlying type: string
type TargetTypeEnum string

// Set of constants representing the allowable values for TargetTypeEnum
const (
	TargetTypeCompartment TargetTypeEnum = "COMPARTMENT"
	TargetTypeTag         TargetTypeEnum = "TAG"
)

var mappingTargetTypeEnum = map[string]TargetTypeEnum{
	"COMPARTMENT": TargetTypeCompartment,
	"TAG":         TargetTypeTag,
}

var mappingTargetTypeEnumLowerCase = map[string]TargetTypeEnum{
	"compartment": TargetTypeCompartment,
	"tag":         TargetTypeTag,
}

// GetTargetTypeEnumValues Enumerates the set of values for TargetTypeEnum
func GetTargetTypeEnumValues() []TargetTypeEnum {
	values := make([]TargetTypeEnum, 0)
	for _, v := range mappingTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetTypeEnumStringValues Enumerates the set of values in String for TargetTypeEnum
func GetTargetTypeEnumStringValues() []string {
	return []string{
		"COMPARTMENT",
		"TAG",
	}
}

// GetMappingTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetTypeEnum(val string) (TargetTypeEnum, bool) {
	enum, ok := mappingTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
