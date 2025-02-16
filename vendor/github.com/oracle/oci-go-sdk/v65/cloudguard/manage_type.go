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

// ManageTypeEnum Enum with underlying type: string
type ManageTypeEnum string

// Set of constants representing the allowable values for ManageTypeEnum
const (
	ManageTypeLocal      ManageTypeEnum = "LOCAL"
	ManageTypeGovernance ManageTypeEnum = "GOVERNANCE"
)

var mappingManageTypeEnum = map[string]ManageTypeEnum{
	"LOCAL":      ManageTypeLocal,
	"GOVERNANCE": ManageTypeGovernance,
}

var mappingManageTypeEnumLowerCase = map[string]ManageTypeEnum{
	"local":      ManageTypeLocal,
	"governance": ManageTypeGovernance,
}

// GetManageTypeEnumValues Enumerates the set of values for ManageTypeEnum
func GetManageTypeEnumValues() []ManageTypeEnum {
	values := make([]ManageTypeEnum, 0)
	for _, v := range mappingManageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManageTypeEnumStringValues Enumerates the set of values in String for ManageTypeEnum
func GetManageTypeEnumStringValues() []string {
	return []string{
		"LOCAL",
		"GOVERNANCE",
	}
}

// GetMappingManageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManageTypeEnum(val string) (ManageTypeEnum, bool) {
	enum, ok := mappingManageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
