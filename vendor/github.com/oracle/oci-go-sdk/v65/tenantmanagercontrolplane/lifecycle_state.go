// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"strings"
)

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateCreating   LifecycleStateEnum = "CREATING"
	LifecycleStateActive     LifecycleStateEnum = "ACTIVE"
	LifecycleStateInactive   LifecycleStateEnum = "INACTIVE"
	LifecycleStateUpdating   LifecycleStateEnum = "UPDATING"
	LifecycleStateFailed     LifecycleStateEnum = "FAILED"
	LifecycleStateTerminated LifecycleStateEnum = "TERMINATED"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"CREATING":   LifecycleStateCreating,
	"ACTIVE":     LifecycleStateActive,
	"INACTIVE":   LifecycleStateInactive,
	"UPDATING":   LifecycleStateUpdating,
	"FAILED":     LifecycleStateFailed,
	"TERMINATED": LifecycleStateTerminated,
}

var mappingLifecycleStateEnumLowerCase = map[string]LifecycleStateEnum{
	"creating":   LifecycleStateCreating,
	"active":     LifecycleStateActive,
	"inactive":   LifecycleStateInactive,
	"updating":   LifecycleStateUpdating,
	"failed":     LifecycleStateFailed,
	"terminated": LifecycleStateTerminated,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleStateEnumStringValues Enumerates the set of values in String for LifecycleStateEnum
func GetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"FAILED",
		"TERMINATED",
	}
}

// GetMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStateEnum(val string) (LifecycleStateEnum, bool) {
	enum, ok := mappingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
