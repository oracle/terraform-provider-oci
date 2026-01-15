// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ClassicSubscriptionLifecycleStateEnum Enum with underlying type: string
type ClassicSubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for ClassicSubscriptionLifecycleStateEnum
const (
	ClassicSubscriptionLifecycleStateCreating ClassicSubscriptionLifecycleStateEnum = "CREATING"
	ClassicSubscriptionLifecycleStateActive   ClassicSubscriptionLifecycleStateEnum = "ACTIVE"
	ClassicSubscriptionLifecycleStateInactive ClassicSubscriptionLifecycleStateEnum = "INACTIVE"
	ClassicSubscriptionLifecycleStateUpdating ClassicSubscriptionLifecycleStateEnum = "UPDATING"
	ClassicSubscriptionLifecycleStateDeleting ClassicSubscriptionLifecycleStateEnum = "DELETING"
	ClassicSubscriptionLifecycleStateDeleted  ClassicSubscriptionLifecycleStateEnum = "DELETED"
	ClassicSubscriptionLifecycleStateFailed   ClassicSubscriptionLifecycleStateEnum = "FAILED"
)

var mappingClassicSubscriptionLifecycleStateEnum = map[string]ClassicSubscriptionLifecycleStateEnum{
	"CREATING": ClassicSubscriptionLifecycleStateCreating,
	"ACTIVE":   ClassicSubscriptionLifecycleStateActive,
	"INACTIVE": ClassicSubscriptionLifecycleStateInactive,
	"UPDATING": ClassicSubscriptionLifecycleStateUpdating,
	"DELETING": ClassicSubscriptionLifecycleStateDeleting,
	"DELETED":  ClassicSubscriptionLifecycleStateDeleted,
	"FAILED":   ClassicSubscriptionLifecycleStateFailed,
}

var mappingClassicSubscriptionLifecycleStateEnumLowerCase = map[string]ClassicSubscriptionLifecycleStateEnum{
	"creating": ClassicSubscriptionLifecycleStateCreating,
	"active":   ClassicSubscriptionLifecycleStateActive,
	"inactive": ClassicSubscriptionLifecycleStateInactive,
	"updating": ClassicSubscriptionLifecycleStateUpdating,
	"deleting": ClassicSubscriptionLifecycleStateDeleting,
	"deleted":  ClassicSubscriptionLifecycleStateDeleted,
	"failed":   ClassicSubscriptionLifecycleStateFailed,
}

// GetClassicSubscriptionLifecycleStateEnumValues Enumerates the set of values for ClassicSubscriptionLifecycleStateEnum
func GetClassicSubscriptionLifecycleStateEnumValues() []ClassicSubscriptionLifecycleStateEnum {
	values := make([]ClassicSubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingClassicSubscriptionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetClassicSubscriptionLifecycleStateEnumStringValues Enumerates the set of values in String for ClassicSubscriptionLifecycleStateEnum
func GetClassicSubscriptionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingClassicSubscriptionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClassicSubscriptionLifecycleStateEnum(val string) (ClassicSubscriptionLifecycleStateEnum, bool) {
	enum, ok := mappingClassicSubscriptionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
