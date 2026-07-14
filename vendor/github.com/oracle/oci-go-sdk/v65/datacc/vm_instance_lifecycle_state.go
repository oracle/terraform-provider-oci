// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// VmInstanceLifecycleStateEnum Enum with underlying type: string
type VmInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for VmInstanceLifecycleStateEnum
const (
	VmInstanceLifecycleStateCreating VmInstanceLifecycleStateEnum = "CREATING"
	VmInstanceLifecycleStateActive   VmInstanceLifecycleStateEnum = "ACTIVE"
	VmInstanceLifecycleStateInactive VmInstanceLifecycleStateEnum = "INACTIVE"
	VmInstanceLifecycleStateUpdating VmInstanceLifecycleStateEnum = "UPDATING"
	VmInstanceLifecycleStateDeleting VmInstanceLifecycleStateEnum = "DELETING"
	VmInstanceLifecycleStateDeleted  VmInstanceLifecycleStateEnum = "DELETED"
)

var mappingVmInstanceLifecycleStateEnum = map[string]VmInstanceLifecycleStateEnum{
	"CREATING": VmInstanceLifecycleStateCreating,
	"ACTIVE":   VmInstanceLifecycleStateActive,
	"INACTIVE": VmInstanceLifecycleStateInactive,
	"UPDATING": VmInstanceLifecycleStateUpdating,
	"DELETING": VmInstanceLifecycleStateDeleting,
	"DELETED":  VmInstanceLifecycleStateDeleted,
}

var mappingVmInstanceLifecycleStateEnumLowerCase = map[string]VmInstanceLifecycleStateEnum{
	"creating": VmInstanceLifecycleStateCreating,
	"active":   VmInstanceLifecycleStateActive,
	"inactive": VmInstanceLifecycleStateInactive,
	"updating": VmInstanceLifecycleStateUpdating,
	"deleting": VmInstanceLifecycleStateDeleting,
	"deleted":  VmInstanceLifecycleStateDeleted,
}

// GetVmInstanceLifecycleStateEnumValues Enumerates the set of values for VmInstanceLifecycleStateEnum
func GetVmInstanceLifecycleStateEnumValues() []VmInstanceLifecycleStateEnum {
	values := make([]VmInstanceLifecycleStateEnum, 0)
	for _, v := range mappingVmInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVmInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for VmInstanceLifecycleStateEnum
func GetVmInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
	}
}

// GetMappingVmInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmInstanceLifecycleStateEnum(val string) (VmInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingVmInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
