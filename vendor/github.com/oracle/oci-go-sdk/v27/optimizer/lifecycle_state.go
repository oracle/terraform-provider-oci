// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

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

var mappingLifecycleState = map[string]LifecycleStateEnum{
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

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleState {
		values = append(values, v)
	}
	return values
}
