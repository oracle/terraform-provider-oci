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

// VmClusterNetworkLifecycleStateEnum Enum with underlying type: string
type VmClusterNetworkLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterNetworkLifecycleStateEnum
const (
	VmClusterNetworkLifecycleStateCreating           VmClusterNetworkLifecycleStateEnum = "CREATING"
	VmClusterNetworkLifecycleStateRequiresValidation VmClusterNetworkLifecycleStateEnum = "REQUIRES_VALIDATION"
	VmClusterNetworkLifecycleStateValidating         VmClusterNetworkLifecycleStateEnum = "VALIDATING"
	VmClusterNetworkLifecycleStateValidated          VmClusterNetworkLifecycleStateEnum = "VALIDATED"
	VmClusterNetworkLifecycleStateValidationFailed   VmClusterNetworkLifecycleStateEnum = "VALIDATION_FAILED"
	VmClusterNetworkLifecycleStateUpdating           VmClusterNetworkLifecycleStateEnum = "UPDATING"
	VmClusterNetworkLifecycleStateAllocated          VmClusterNetworkLifecycleStateEnum = "ALLOCATED"
	VmClusterNetworkLifecycleStateDeleting           VmClusterNetworkLifecycleStateEnum = "DELETING"
	VmClusterNetworkLifecycleStateDeleted            VmClusterNetworkLifecycleStateEnum = "DELETED"
)

var mappingVmClusterNetworkLifecycleStateEnum = map[string]VmClusterNetworkLifecycleStateEnum{
	"CREATING":            VmClusterNetworkLifecycleStateCreating,
	"REQUIRES_VALIDATION": VmClusterNetworkLifecycleStateRequiresValidation,
	"VALIDATING":          VmClusterNetworkLifecycleStateValidating,
	"VALIDATED":           VmClusterNetworkLifecycleStateValidated,
	"VALIDATION_FAILED":   VmClusterNetworkLifecycleStateValidationFailed,
	"UPDATING":            VmClusterNetworkLifecycleStateUpdating,
	"ALLOCATED":           VmClusterNetworkLifecycleStateAllocated,
	"DELETING":            VmClusterNetworkLifecycleStateDeleting,
	"DELETED":             VmClusterNetworkLifecycleStateDeleted,
}

var mappingVmClusterNetworkLifecycleStateEnumLowerCase = map[string]VmClusterNetworkLifecycleStateEnum{
	"creating":            VmClusterNetworkLifecycleStateCreating,
	"requires_validation": VmClusterNetworkLifecycleStateRequiresValidation,
	"validating":          VmClusterNetworkLifecycleStateValidating,
	"validated":           VmClusterNetworkLifecycleStateValidated,
	"validation_failed":   VmClusterNetworkLifecycleStateValidationFailed,
	"updating":            VmClusterNetworkLifecycleStateUpdating,
	"allocated":           VmClusterNetworkLifecycleStateAllocated,
	"deleting":            VmClusterNetworkLifecycleStateDeleting,
	"deleted":             VmClusterNetworkLifecycleStateDeleted,
}

// GetVmClusterNetworkLifecycleStateEnumValues Enumerates the set of values for VmClusterNetworkLifecycleStateEnum
func GetVmClusterNetworkLifecycleStateEnumValues() []VmClusterNetworkLifecycleStateEnum {
	values := make([]VmClusterNetworkLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterNetworkLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterNetworkLifecycleStateEnumStringValues Enumerates the set of values in String for VmClusterNetworkLifecycleStateEnum
func GetVmClusterNetworkLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"REQUIRES_VALIDATION",
		"VALIDATING",
		"VALIDATED",
		"VALIDATION_FAILED",
		"UPDATING",
		"ALLOCATED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingVmClusterNetworkLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterNetworkLifecycleStateEnum(val string) (VmClusterNetworkLifecycleStateEnum, bool) {
	enum, ok := mappingVmClusterNetworkLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
