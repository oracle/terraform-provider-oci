// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ConditionValueTypeEnum Enum with underlying type: string
type ConditionValueTypeEnum string

// Set of constants representing the allowable values for ConditionValueTypeEnum
const (
	ConditionValueTypeManaged ConditionValueTypeEnum = "MANAGED"
	ConditionValueTypeCustom  ConditionValueTypeEnum = "CUSTOM"
)

var mappingConditionValueTypeEnum = map[string]ConditionValueTypeEnum{
	"MANAGED": ConditionValueTypeManaged,
	"CUSTOM":  ConditionValueTypeCustom,
}

var mappingConditionValueTypeEnumLowerCase = map[string]ConditionValueTypeEnum{
	"managed": ConditionValueTypeManaged,
	"custom":  ConditionValueTypeCustom,
}

// GetConditionValueTypeEnumValues Enumerates the set of values for ConditionValueTypeEnum
func GetConditionValueTypeEnumValues() []ConditionValueTypeEnum {
	values := make([]ConditionValueTypeEnum, 0)
	for _, v := range mappingConditionValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionValueTypeEnumStringValues Enumerates the set of values in String for ConditionValueTypeEnum
func GetConditionValueTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingConditionValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionValueTypeEnum(val string) (ConditionValueTypeEnum, bool) {
	enum, ok := mappingConditionValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
