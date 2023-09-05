// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// VmwareAccountStatesEnum Enum with underlying type: string
type VmwareAccountStatesEnum string

// Set of constants representing the allowable values for VmwareAccountStatesEnum
const (
	VmwareAccountStatesActive    VmwareAccountStatesEnum = "ACTIVE"
	VmwareAccountStatesInactive  VmwareAccountStatesEnum = "INACTIVE"
	VmwareAccountStatesDeleted   VmwareAccountStatesEnum = "DELETED"
	VmwareAccountStatesSuspended VmwareAccountStatesEnum = "SUSPENDED"
)

var mappingVmwareAccountStatesEnum = map[string]VmwareAccountStatesEnum{
	"ACTIVE":    VmwareAccountStatesActive,
	"INACTIVE":  VmwareAccountStatesInactive,
	"DELETED":   VmwareAccountStatesDeleted,
	"SUSPENDED": VmwareAccountStatesSuspended,
}

var mappingVmwareAccountStatesEnumLowerCase = map[string]VmwareAccountStatesEnum{
	"active":    VmwareAccountStatesActive,
	"inactive":  VmwareAccountStatesInactive,
	"deleted":   VmwareAccountStatesDeleted,
	"suspended": VmwareAccountStatesSuspended,
}

// GetVmwareAccountStatesEnumValues Enumerates the set of values for VmwareAccountStatesEnum
func GetVmwareAccountStatesEnumValues() []VmwareAccountStatesEnum {
	values := make([]VmwareAccountStatesEnum, 0)
	for _, v := range mappingVmwareAccountStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetVmwareAccountStatesEnumStringValues Enumerates the set of values in String for VmwareAccountStatesEnum
func GetVmwareAccountStatesEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"SUSPENDED",
	}
}

// GetMappingVmwareAccountStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmwareAccountStatesEnum(val string) (VmwareAccountStatesEnum, bool) {
	enum, ok := mappingVmwareAccountStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
