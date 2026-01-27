// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// ManagementApplianceConnectionTypeEnum Enum with underlying type: string
type ManagementApplianceConnectionTypeEnum string

// Set of constants representing the allowable values for ManagementApplianceConnectionTypeEnum
const (
	ManagementApplianceConnectionTypeVcenter      ManagementApplianceConnectionTypeEnum = "VCENTER"
	ManagementApplianceConnectionTypeAdminVcenter ManagementApplianceConnectionTypeEnum = "ADMIN_VCENTER"
	ManagementApplianceConnectionTypeNsx          ManagementApplianceConnectionTypeEnum = "NSX"
)

var mappingManagementApplianceConnectionTypeEnum = map[string]ManagementApplianceConnectionTypeEnum{
	"VCENTER":       ManagementApplianceConnectionTypeVcenter,
	"ADMIN_VCENTER": ManagementApplianceConnectionTypeAdminVcenter,
	"NSX":           ManagementApplianceConnectionTypeNsx,
}

var mappingManagementApplianceConnectionTypeEnumLowerCase = map[string]ManagementApplianceConnectionTypeEnum{
	"vcenter":       ManagementApplianceConnectionTypeVcenter,
	"admin_vcenter": ManagementApplianceConnectionTypeAdminVcenter,
	"nsx":           ManagementApplianceConnectionTypeNsx,
}

// GetManagementApplianceConnectionTypeEnumValues Enumerates the set of values for ManagementApplianceConnectionTypeEnum
func GetManagementApplianceConnectionTypeEnumValues() []ManagementApplianceConnectionTypeEnum {
	values := make([]ManagementApplianceConnectionTypeEnum, 0)
	for _, v := range mappingManagementApplianceConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementApplianceConnectionTypeEnumStringValues Enumerates the set of values in String for ManagementApplianceConnectionTypeEnum
func GetManagementApplianceConnectionTypeEnumStringValues() []string {
	return []string{
		"VCENTER",
		"ADMIN_VCENTER",
		"NSX",
	}
}

// GetMappingManagementApplianceConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementApplianceConnectionTypeEnum(val string) (ManagementApplianceConnectionTypeEnum, bool) {
	enum, ok := mappingManagementApplianceConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
