// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// LifecycleStateEnumEnum Enum with underlying type: string
type LifecycleStateEnumEnum string

// Set of constants representing the allowable values for LifecycleStateEnumEnum
const (
	LifecycleStateEnumCreating LifecycleStateEnumEnum = "CREATING"
	LifecycleStateEnumActive   LifecycleStateEnumEnum = "ACTIVE"
	LifecycleStateEnumUpdating LifecycleStateEnumEnum = "UPDATING"
	LifecycleStateEnumDeleting LifecycleStateEnumEnum = "DELETING"
	LifecycleStateEnumDeleted  LifecycleStateEnumEnum = "DELETED"
	LifecycleStateEnumFailed   LifecycleStateEnumEnum = "FAILED"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnumEnum{
	"CREATING": LifecycleStateEnumCreating,
	"ACTIVE":   LifecycleStateEnumActive,
	"UPDATING": LifecycleStateEnumUpdating,
	"DELETING": LifecycleStateEnumDeleting,
	"DELETED":  LifecycleStateEnumDeleted,
	"FAILED":   LifecycleStateEnumFailed,
}

// GetLifecycleStateEnumEnumValues Enumerates the set of values for LifecycleStateEnumEnum
func GetLifecycleStateEnumEnumValues() []LifecycleStateEnumEnum {
	values := make([]LifecycleStateEnumEnum, 0)
	for _, v := range mappingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}
