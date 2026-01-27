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

// ManagementApplianceLifecycleDetailsEnum Enum with underlying type: string
type ManagementApplianceLifecycleDetailsEnum string

// Set of constants representing the allowable values for ManagementApplianceLifecycleDetailsEnum
const (
	ManagementApplianceLifecycleDetailsHealthy                 ManagementApplianceLifecycleDetailsEnum = "HEALTHY"
	ManagementApplianceLifecycleDetailsUpdating                ManagementApplianceLifecycleDetailsEnum = "UPDATING"
	ManagementApplianceLifecycleDetailsDeleting                ManagementApplianceLifecycleDetailsEnum = "DELETING"
	ManagementApplianceLifecycleDetailsDeleted                 ManagementApplianceLifecycleDetailsEnum = "DELETED"
	ManagementApplianceLifecycleDetailsFailed                  ManagementApplianceLifecycleDetailsEnum = "FAILED"
	ManagementApplianceLifecycleDetailsCreating                ManagementApplianceLifecycleDetailsEnum = "CREATING"
	ManagementApplianceLifecycleDetailsWaitingForHeartbeat     ManagementApplianceLifecycleDetailsEnum = "WAITING_FOR_HEARTBEAT"
	ManagementApplianceLifecycleDetailsHeartbeatTimeout        ManagementApplianceLifecycleDetailsEnum = "HEARTBEAT_TIMEOUT"
	ManagementApplianceLifecycleDetailsCanNotConnectToVcenter  ManagementApplianceLifecycleDetailsEnum = "CAN_NOT_CONNECT_TO_VCENTER"
	ManagementApplianceLifecycleDetailsUiPluginIsNotRegistered ManagementApplianceLifecycleDetailsEnum = "UI_PLUGIN_IS_NOT_REGISTERED"
	ManagementApplianceLifecycleDetailsUnknownDetails          ManagementApplianceLifecycleDetailsEnum = "UNKNOWN_DETAILS"
)

var mappingManagementApplianceLifecycleDetailsEnum = map[string]ManagementApplianceLifecycleDetailsEnum{
	"HEALTHY":                     ManagementApplianceLifecycleDetailsHealthy,
	"UPDATING":                    ManagementApplianceLifecycleDetailsUpdating,
	"DELETING":                    ManagementApplianceLifecycleDetailsDeleting,
	"DELETED":                     ManagementApplianceLifecycleDetailsDeleted,
	"FAILED":                      ManagementApplianceLifecycleDetailsFailed,
	"CREATING":                    ManagementApplianceLifecycleDetailsCreating,
	"WAITING_FOR_HEARTBEAT":       ManagementApplianceLifecycleDetailsWaitingForHeartbeat,
	"HEARTBEAT_TIMEOUT":           ManagementApplianceLifecycleDetailsHeartbeatTimeout,
	"CAN_NOT_CONNECT_TO_VCENTER":  ManagementApplianceLifecycleDetailsCanNotConnectToVcenter,
	"UI_PLUGIN_IS_NOT_REGISTERED": ManagementApplianceLifecycleDetailsUiPluginIsNotRegistered,
	"UNKNOWN_DETAILS":             ManagementApplianceLifecycleDetailsUnknownDetails,
}

var mappingManagementApplianceLifecycleDetailsEnumLowerCase = map[string]ManagementApplianceLifecycleDetailsEnum{
	"healthy":                     ManagementApplianceLifecycleDetailsHealthy,
	"updating":                    ManagementApplianceLifecycleDetailsUpdating,
	"deleting":                    ManagementApplianceLifecycleDetailsDeleting,
	"deleted":                     ManagementApplianceLifecycleDetailsDeleted,
	"failed":                      ManagementApplianceLifecycleDetailsFailed,
	"creating":                    ManagementApplianceLifecycleDetailsCreating,
	"waiting_for_heartbeat":       ManagementApplianceLifecycleDetailsWaitingForHeartbeat,
	"heartbeat_timeout":           ManagementApplianceLifecycleDetailsHeartbeatTimeout,
	"can_not_connect_to_vcenter":  ManagementApplianceLifecycleDetailsCanNotConnectToVcenter,
	"ui_plugin_is_not_registered": ManagementApplianceLifecycleDetailsUiPluginIsNotRegistered,
	"unknown_details":             ManagementApplianceLifecycleDetailsUnknownDetails,
}

// GetManagementApplianceLifecycleDetailsEnumValues Enumerates the set of values for ManagementApplianceLifecycleDetailsEnum
func GetManagementApplianceLifecycleDetailsEnumValues() []ManagementApplianceLifecycleDetailsEnum {
	values := make([]ManagementApplianceLifecycleDetailsEnum, 0)
	for _, v := range mappingManagementApplianceLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementApplianceLifecycleDetailsEnumStringValues Enumerates the set of values in String for ManagementApplianceLifecycleDetailsEnum
func GetManagementApplianceLifecycleDetailsEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"CREATING",
		"WAITING_FOR_HEARTBEAT",
		"HEARTBEAT_TIMEOUT",
		"CAN_NOT_CONNECT_TO_VCENTER",
		"UI_PLUGIN_IS_NOT_REGISTERED",
		"UNKNOWN_DETAILS",
	}
}

// GetMappingManagementApplianceLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementApplianceLifecycleDetailsEnum(val string) (ManagementApplianceLifecycleDetailsEnum, bool) {
	enum, ok := mappingManagementApplianceLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
