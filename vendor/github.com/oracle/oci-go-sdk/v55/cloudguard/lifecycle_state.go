// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateCreating LifecycleStateEnum = "CREATING"
	LifecycleStateUpdating LifecycleStateEnum = "UPDATING"
	LifecycleStateActive   LifecycleStateEnum = "ACTIVE"
	LifecycleStateInactive LifecycleStateEnum = "INACTIVE"
	LifecycleStateDeleting LifecycleStateEnum = "DELETING"
	LifecycleStateDeleted  LifecycleStateEnum = "DELETED"
	LifecycleStateFailed   LifecycleStateEnum = "FAILED"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"CREATING": LifecycleStateCreating,
	"UPDATING": LifecycleStateUpdating,
	"ACTIVE":   LifecycleStateActive,
	"INACTIVE": LifecycleStateInactive,
	"DELETING": LifecycleStateDeleting,
	"DELETED":  LifecycleStateDeleted,
	"FAILED":   LifecycleStateFailed,
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
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}
