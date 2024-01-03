// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"strings"
)

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateActive    LifecycleStateEnum = "ACTIVE"
	LifecycleStateFailed    LifecycleStateEnum = "FAILED"
	LifecycleStateInactive  LifecycleStateEnum = "INACTIVE"
	LifecycleStateAttaching LifecycleStateEnum = "ATTACHING"
	LifecycleStateDetaching LifecycleStateEnum = "DETACHING"
	LifecycleStateDeleting  LifecycleStateEnum = "DELETING"
	LifecycleStateDeleted   LifecycleStateEnum = "DELETED"
	LifecycleStateUpdating  LifecycleStateEnum = "UPDATING"
	LifecycleStateCreating  LifecycleStateEnum = "CREATING"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"ACTIVE":    LifecycleStateActive,
	"FAILED":    LifecycleStateFailed,
	"INACTIVE":  LifecycleStateInactive,
	"ATTACHING": LifecycleStateAttaching,
	"DETACHING": LifecycleStateDetaching,
	"DELETING":  LifecycleStateDeleting,
	"DELETED":   LifecycleStateDeleted,
	"UPDATING":  LifecycleStateUpdating,
	"CREATING":  LifecycleStateCreating,
}

var mappingLifecycleStateEnumLowerCase = map[string]LifecycleStateEnum{
	"active":    LifecycleStateActive,
	"failed":    LifecycleStateFailed,
	"inactive":  LifecycleStateInactive,
	"attaching": LifecycleStateAttaching,
	"detaching": LifecycleStateDetaching,
	"deleting":  LifecycleStateDeleting,
	"deleted":   LifecycleStateDeleted,
	"updating":  LifecycleStateUpdating,
	"creating":  LifecycleStateCreating,
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
		"ACTIVE",
		"FAILED",
		"INACTIVE",
		"ATTACHING",
		"DETACHING",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
	}
}

// GetMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStateEnum(val string) (LifecycleStateEnum, bool) {
	enum, ok := mappingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
