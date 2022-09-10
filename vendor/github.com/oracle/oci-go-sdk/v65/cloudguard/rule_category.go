// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// RuleCategoryEnum Enum with underlying type: string
type RuleCategoryEnum string

// Set of constants representing the allowable values for RuleCategoryEnum
const (
	RuleCategoryRule     RuleCategoryEnum = "RULE"
	RuleCategoryTemplate RuleCategoryEnum = "TEMPLATE"
)

var mappingRuleCategoryEnum = map[string]RuleCategoryEnum{
	"RULE":     RuleCategoryRule,
	"TEMPLATE": RuleCategoryTemplate,
}

var mappingRuleCategoryEnumLowerCase = map[string]RuleCategoryEnum{
	"rule":     RuleCategoryRule,
	"template": RuleCategoryTemplate,
}

// GetRuleCategoryEnumValues Enumerates the set of values for RuleCategoryEnum
func GetRuleCategoryEnumValues() []RuleCategoryEnum {
	values := make([]RuleCategoryEnum, 0)
	for _, v := range mappingRuleCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleCategoryEnumStringValues Enumerates the set of values in String for RuleCategoryEnum
func GetRuleCategoryEnumStringValues() []string {
	return []string{
		"RULE",
		"TEMPLATE",
	}
}

// GetMappingRuleCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleCategoryEnum(val string) (RuleCategoryEnum, bool) {
	enum, ok := mappingRuleCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
