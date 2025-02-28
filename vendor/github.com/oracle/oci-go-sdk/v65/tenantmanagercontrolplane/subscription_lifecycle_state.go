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

// SubscriptionLifecycleStateEnum Enum with underlying type: string
type SubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionLifecycleStateEnum
const (
	SubscriptionLifecycleStateNeedsAttention SubscriptionLifecycleStateEnum = "NEEDS_ATTENTION"
	SubscriptionLifecycleStateActive         SubscriptionLifecycleStateEnum = "ACTIVE"
	SubscriptionLifecycleStateInactive       SubscriptionLifecycleStateEnum = "INACTIVE"
	SubscriptionLifecycleStateFailed         SubscriptionLifecycleStateEnum = "FAILED"
	SubscriptionLifecycleStateCreating       SubscriptionLifecycleStateEnum = "CREATING"
)

var mappingSubscriptionLifecycleStateEnum = map[string]SubscriptionLifecycleStateEnum{
	"NEEDS_ATTENTION": SubscriptionLifecycleStateNeedsAttention,
	"ACTIVE":          SubscriptionLifecycleStateActive,
	"INACTIVE":        SubscriptionLifecycleStateInactive,
	"FAILED":          SubscriptionLifecycleStateFailed,
	"CREATING":        SubscriptionLifecycleStateCreating,
}

var mappingSubscriptionLifecycleStateEnumLowerCase = map[string]SubscriptionLifecycleStateEnum{
	"needs_attention": SubscriptionLifecycleStateNeedsAttention,
	"active":          SubscriptionLifecycleStateActive,
	"inactive":        SubscriptionLifecycleStateInactive,
	"failed":          SubscriptionLifecycleStateFailed,
	"creating":        SubscriptionLifecycleStateCreating,
}

// GetSubscriptionLifecycleStateEnumValues Enumerates the set of values for SubscriptionLifecycleStateEnum
func GetSubscriptionLifecycleStateEnumValues() []SubscriptionLifecycleStateEnum {
	values := make([]SubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionLifecycleStateEnumStringValues Enumerates the set of values in String for SubscriptionLifecycleStateEnum
func GetSubscriptionLifecycleStateEnumStringValues() []string {
	return []string{
		"NEEDS_ATTENTION",
		"ACTIVE",
		"INACTIVE",
		"FAILED",
		"CREATING",
	}
}

// GetMappingSubscriptionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionLifecycleStateEnum(val string) (SubscriptionLifecycleStateEnum, bool) {
	enum, ok := mappingSubscriptionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
