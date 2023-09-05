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

// VmwareSubscriptionStatesEnum Enum with underlying type: string
type VmwareSubscriptionStatesEnum string

// Set of constants representing the allowable values for VmwareSubscriptionStatesEnum
const (
	VmwareSubscriptionStatesActive    VmwareSubscriptionStatesEnum = "ACTIVE"
	VmwareSubscriptionStatesAvailable VmwareSubscriptionStatesEnum = "AVAILABLE"
	VmwareSubscriptionStatesExpired   VmwareSubscriptionStatesEnum = "EXPIRED"
)

var mappingVmwareSubscriptionStatesEnum = map[string]VmwareSubscriptionStatesEnum{
	"ACTIVE":    VmwareSubscriptionStatesActive,
	"AVAILABLE": VmwareSubscriptionStatesAvailable,
	"EXPIRED":   VmwareSubscriptionStatesExpired,
}

var mappingVmwareSubscriptionStatesEnumLowerCase = map[string]VmwareSubscriptionStatesEnum{
	"active":    VmwareSubscriptionStatesActive,
	"available": VmwareSubscriptionStatesAvailable,
	"expired":   VmwareSubscriptionStatesExpired,
}

// GetVmwareSubscriptionStatesEnumValues Enumerates the set of values for VmwareSubscriptionStatesEnum
func GetVmwareSubscriptionStatesEnumValues() []VmwareSubscriptionStatesEnum {
	values := make([]VmwareSubscriptionStatesEnum, 0)
	for _, v := range mappingVmwareSubscriptionStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetVmwareSubscriptionStatesEnumStringValues Enumerates the set of values in String for VmwareSubscriptionStatesEnum
func GetVmwareSubscriptionStatesEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"AVAILABLE",
		"EXPIRED",
	}
}

// GetMappingVmwareSubscriptionStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmwareSubscriptionStatesEnum(val string) (VmwareSubscriptionStatesEnum, bool) {
	enum, ok := mappingVmwareSubscriptionStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
