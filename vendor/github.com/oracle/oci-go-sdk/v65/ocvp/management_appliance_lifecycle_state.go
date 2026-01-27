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

// ManagementApplianceLifecycleStateEnum Enum with underlying type: string
type ManagementApplianceLifecycleStateEnum string

// Set of constants representing the allowable values for ManagementApplianceLifecycleStateEnum
const (
	ManagementApplianceLifecycleStateCreating       ManagementApplianceLifecycleStateEnum = "CREATING"
	ManagementApplianceLifecycleStateUpdating       ManagementApplianceLifecycleStateEnum = "UPDATING"
	ManagementApplianceLifecycleStateActive         ManagementApplianceLifecycleStateEnum = "ACTIVE"
	ManagementApplianceLifecycleStateNeedsAttention ManagementApplianceLifecycleStateEnum = "NEEDS_ATTENTION"
	ManagementApplianceLifecycleStateDeleting       ManagementApplianceLifecycleStateEnum = "DELETING"
	ManagementApplianceLifecycleStateDeleted        ManagementApplianceLifecycleStateEnum = "DELETED"
	ManagementApplianceLifecycleStateFailed         ManagementApplianceLifecycleStateEnum = "FAILED"
)

var mappingManagementApplianceLifecycleStateEnum = map[string]ManagementApplianceLifecycleStateEnum{
	"CREATING":        ManagementApplianceLifecycleStateCreating,
	"UPDATING":        ManagementApplianceLifecycleStateUpdating,
	"ACTIVE":          ManagementApplianceLifecycleStateActive,
	"NEEDS_ATTENTION": ManagementApplianceLifecycleStateNeedsAttention,
	"DELETING":        ManagementApplianceLifecycleStateDeleting,
	"DELETED":         ManagementApplianceLifecycleStateDeleted,
	"FAILED":          ManagementApplianceLifecycleStateFailed,
}

var mappingManagementApplianceLifecycleStateEnumLowerCase = map[string]ManagementApplianceLifecycleStateEnum{
	"creating":        ManagementApplianceLifecycleStateCreating,
	"updating":        ManagementApplianceLifecycleStateUpdating,
	"active":          ManagementApplianceLifecycleStateActive,
	"needs_attention": ManagementApplianceLifecycleStateNeedsAttention,
	"deleting":        ManagementApplianceLifecycleStateDeleting,
	"deleted":         ManagementApplianceLifecycleStateDeleted,
	"failed":          ManagementApplianceLifecycleStateFailed,
}

// GetManagementApplianceLifecycleStateEnumValues Enumerates the set of values for ManagementApplianceLifecycleStateEnum
func GetManagementApplianceLifecycleStateEnumValues() []ManagementApplianceLifecycleStateEnum {
	values := make([]ManagementApplianceLifecycleStateEnum, 0)
	for _, v := range mappingManagementApplianceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementApplianceLifecycleStateEnumStringValues Enumerates the set of values in String for ManagementApplianceLifecycleStateEnum
func GetManagementApplianceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingManagementApplianceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementApplianceLifecycleStateEnum(val string) (ManagementApplianceLifecycleStateEnum, bool) {
	enum, ok := mappingManagementApplianceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
