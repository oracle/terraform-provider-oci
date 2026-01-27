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

// ManagementApplianceHeartbeatConnectionTypeEnum Enum with underlying type: string
type ManagementApplianceHeartbeatConnectionTypeEnum string

// Set of constants representing the allowable values for ManagementApplianceHeartbeatConnectionTypeEnum
const (
	ManagementApplianceHeartbeatConnectionTypeVcenter              ManagementApplianceHeartbeatConnectionTypeEnum = "VCENTER"
	ManagementApplianceHeartbeatConnectionTypeAdminVcenter         ManagementApplianceHeartbeatConnectionTypeEnum = "ADMIN_VCENTER"
	ManagementApplianceHeartbeatConnectionTypeUiPluginRegistration ManagementApplianceHeartbeatConnectionTypeEnum = "UI_PLUGIN_REGISTRATION"
	ManagementApplianceHeartbeatConnectionTypeNsx                  ManagementApplianceHeartbeatConnectionTypeEnum = "NSX"
)

var mappingManagementApplianceHeartbeatConnectionTypeEnum = map[string]ManagementApplianceHeartbeatConnectionTypeEnum{
	"VCENTER":                ManagementApplianceHeartbeatConnectionTypeVcenter,
	"ADMIN_VCENTER":          ManagementApplianceHeartbeatConnectionTypeAdminVcenter,
	"UI_PLUGIN_REGISTRATION": ManagementApplianceHeartbeatConnectionTypeUiPluginRegistration,
	"NSX":                    ManagementApplianceHeartbeatConnectionTypeNsx,
}

var mappingManagementApplianceHeartbeatConnectionTypeEnumLowerCase = map[string]ManagementApplianceHeartbeatConnectionTypeEnum{
	"vcenter":                ManagementApplianceHeartbeatConnectionTypeVcenter,
	"admin_vcenter":          ManagementApplianceHeartbeatConnectionTypeAdminVcenter,
	"ui_plugin_registration": ManagementApplianceHeartbeatConnectionTypeUiPluginRegistration,
	"nsx":                    ManagementApplianceHeartbeatConnectionTypeNsx,
}

// GetManagementApplianceHeartbeatConnectionTypeEnumValues Enumerates the set of values for ManagementApplianceHeartbeatConnectionTypeEnum
func GetManagementApplianceHeartbeatConnectionTypeEnumValues() []ManagementApplianceHeartbeatConnectionTypeEnum {
	values := make([]ManagementApplianceHeartbeatConnectionTypeEnum, 0)
	for _, v := range mappingManagementApplianceHeartbeatConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementApplianceHeartbeatConnectionTypeEnumStringValues Enumerates the set of values in String for ManagementApplianceHeartbeatConnectionTypeEnum
func GetManagementApplianceHeartbeatConnectionTypeEnumStringValues() []string {
	return []string{
		"VCENTER",
		"ADMIN_VCENTER",
		"UI_PLUGIN_REGISTRATION",
		"NSX",
	}
}

// GetMappingManagementApplianceHeartbeatConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementApplianceHeartbeatConnectionTypeEnum(val string) (ManagementApplianceHeartbeatConnectionTypeEnum, bool) {
	enum, ok := mappingManagementApplianceHeartbeatConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
