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

// ScopeTypeEnum Enum with underlying type: string
type ScopeTypeEnum string

// Set of constants representing the allowable values for ScopeTypeEnum
const (
	ScopeTypeAll    ScopeTypeEnum = "ALL"
	ScopeTypeCustom ScopeTypeEnum = "CUSTOM"
)

var mappingScopeTypeEnum = map[string]ScopeTypeEnum{
	"ALL":    ScopeTypeAll,
	"CUSTOM": ScopeTypeCustom,
}

var mappingScopeTypeEnumLowerCase = map[string]ScopeTypeEnum{
	"all":    ScopeTypeAll,
	"custom": ScopeTypeCustom,
}

// GetScopeTypeEnumValues Enumerates the set of values for ScopeTypeEnum
func GetScopeTypeEnumValues() []ScopeTypeEnum {
	values := make([]ScopeTypeEnum, 0)
	for _, v := range mappingScopeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScopeTypeEnumStringValues Enumerates the set of values in String for ScopeTypeEnum
func GetScopeTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"CUSTOM",
	}
}

// GetMappingScopeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScopeTypeEnum(val string) (ScopeTypeEnum, bool) {
	enum, ok := mappingScopeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
