// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ConditionOperatorNameEnum Enum with underlying type: string
type ConditionOperatorNameEnum string

// Set of constants representing the allowable values for ConditionOperatorNameEnum
const (
	ConditionOperatorNameAnd         ConditionOperatorNameEnum = "AND"
	ConditionOperatorNameOr          ConditionOperatorNameEnum = "OR"
	ConditionOperatorNameIn          ConditionOperatorNameEnum = "IN"
	ConditionOperatorNameNotIn       ConditionOperatorNameEnum = "NOT_IN"
	ConditionOperatorNameEquals      ConditionOperatorNameEnum = "EQUALS"
	ConditionOperatorNameNotEquals   ConditionOperatorNameEnum = "NOT_EQUALS"
	ConditionOperatorNameLessThan    ConditionOperatorNameEnum = "LESS_THAN"
	ConditionOperatorNameGreaterThan ConditionOperatorNameEnum = "GREATER_THAN"
	ConditionOperatorNameRange       ConditionOperatorNameEnum = "RANGE"
)

var mappingConditionOperatorNameEnum = map[string]ConditionOperatorNameEnum{
	"AND":          ConditionOperatorNameAnd,
	"OR":           ConditionOperatorNameOr,
	"IN":           ConditionOperatorNameIn,
	"NOT_IN":       ConditionOperatorNameNotIn,
	"EQUALS":       ConditionOperatorNameEquals,
	"NOT_EQUALS":   ConditionOperatorNameNotEquals,
	"LESS_THAN":    ConditionOperatorNameLessThan,
	"GREATER_THAN": ConditionOperatorNameGreaterThan,
	"RANGE":        ConditionOperatorNameRange,
}

var mappingConditionOperatorNameEnumLowerCase = map[string]ConditionOperatorNameEnum{
	"and":          ConditionOperatorNameAnd,
	"or":           ConditionOperatorNameOr,
	"in":           ConditionOperatorNameIn,
	"not_in":       ConditionOperatorNameNotIn,
	"equals":       ConditionOperatorNameEquals,
	"not_equals":   ConditionOperatorNameNotEquals,
	"less_than":    ConditionOperatorNameLessThan,
	"greater_than": ConditionOperatorNameGreaterThan,
	"range":        ConditionOperatorNameRange,
}

// GetConditionOperatorNameEnumValues Enumerates the set of values for ConditionOperatorNameEnum
func GetConditionOperatorNameEnumValues() []ConditionOperatorNameEnum {
	values := make([]ConditionOperatorNameEnum, 0)
	for _, v := range mappingConditionOperatorNameEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionOperatorNameEnumStringValues Enumerates the set of values in String for ConditionOperatorNameEnum
func GetConditionOperatorNameEnumStringValues() []string {
	return []string{
		"AND",
		"OR",
		"IN",
		"NOT_IN",
		"EQUALS",
		"NOT_EQUALS",
		"LESS_THAN",
		"GREATER_THAN",
		"RANGE",
	}
}

// GetMappingConditionOperatorNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionOperatorNameEnum(val string) (ConditionOperatorNameEnum, bool) {
	enum, ok := mappingConditionOperatorNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
