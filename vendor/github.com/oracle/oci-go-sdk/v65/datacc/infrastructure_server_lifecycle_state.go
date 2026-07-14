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

// InfrastructureServerLifecycleStateEnum Enum with underlying type: string
type InfrastructureServerLifecycleStateEnum string

// Set of constants representing the allowable values for InfrastructureServerLifecycleStateEnum
const (
	InfrastructureServerLifecycleStateCreating              InfrastructureServerLifecycleStateEnum = "CREATING"
	InfrastructureServerLifecycleStateRequiresValidation    InfrastructureServerLifecycleStateEnum = "REQUIRES_VALIDATION"
	InfrastructureServerLifecycleStateValidating            InfrastructureServerLifecycleStateEnum = "VALIDATING"
	InfrastructureServerLifecycleStateValidationFailed      InfrastructureServerLifecycleStateEnum = "VALIDATION_FAILED"
	InfrastructureServerLifecycleStateRequiresActivation    InfrastructureServerLifecycleStateEnum = "REQUIRES_ACTIVATION"
	InfrastructureServerLifecycleStateActivating            InfrastructureServerLifecycleStateEnum = "ACTIVATING"
	InfrastructureServerLifecycleStateActive                InfrastructureServerLifecycleStateEnum = "ACTIVE"
	InfrastructureServerLifecycleStateActivationFailed      InfrastructureServerLifecycleStateEnum = "ACTIVATION_FAILED"
	InfrastructureServerLifecycleStateFailed                InfrastructureServerLifecycleStateEnum = "FAILED"
	InfrastructureServerLifecycleStateUpdating              InfrastructureServerLifecycleStateEnum = "UPDATING"
	InfrastructureServerLifecycleStateDeleting              InfrastructureServerLifecycleStateEnum = "DELETING"
	InfrastructureServerLifecycleStateDeleted               InfrastructureServerLifecycleStateEnum = "DELETED"
	InfrastructureServerLifecycleStateDisconnected          InfrastructureServerLifecycleStateEnum = "DISCONNECTED"
	InfrastructureServerLifecycleStateMaintenanceInProgress InfrastructureServerLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingInfrastructureServerLifecycleStateEnum = map[string]InfrastructureServerLifecycleStateEnum{
	"CREATING":                InfrastructureServerLifecycleStateCreating,
	"REQUIRES_VALIDATION":     InfrastructureServerLifecycleStateRequiresValidation,
	"VALIDATING":              InfrastructureServerLifecycleStateValidating,
	"VALIDATION_FAILED":       InfrastructureServerLifecycleStateValidationFailed,
	"REQUIRES_ACTIVATION":     InfrastructureServerLifecycleStateRequiresActivation,
	"ACTIVATING":              InfrastructureServerLifecycleStateActivating,
	"ACTIVE":                  InfrastructureServerLifecycleStateActive,
	"ACTIVATION_FAILED":       InfrastructureServerLifecycleStateActivationFailed,
	"FAILED":                  InfrastructureServerLifecycleStateFailed,
	"UPDATING":                InfrastructureServerLifecycleStateUpdating,
	"DELETING":                InfrastructureServerLifecycleStateDeleting,
	"DELETED":                 InfrastructureServerLifecycleStateDeleted,
	"DISCONNECTED":            InfrastructureServerLifecycleStateDisconnected,
	"MAINTENANCE_IN_PROGRESS": InfrastructureServerLifecycleStateMaintenanceInProgress,
}

var mappingInfrastructureServerLifecycleStateEnumLowerCase = map[string]InfrastructureServerLifecycleStateEnum{
	"creating":                InfrastructureServerLifecycleStateCreating,
	"requires_validation":     InfrastructureServerLifecycleStateRequiresValidation,
	"validating":              InfrastructureServerLifecycleStateValidating,
	"validation_failed":       InfrastructureServerLifecycleStateValidationFailed,
	"requires_activation":     InfrastructureServerLifecycleStateRequiresActivation,
	"activating":              InfrastructureServerLifecycleStateActivating,
	"active":                  InfrastructureServerLifecycleStateActive,
	"activation_failed":       InfrastructureServerLifecycleStateActivationFailed,
	"failed":                  InfrastructureServerLifecycleStateFailed,
	"updating":                InfrastructureServerLifecycleStateUpdating,
	"deleting":                InfrastructureServerLifecycleStateDeleting,
	"deleted":                 InfrastructureServerLifecycleStateDeleted,
	"disconnected":            InfrastructureServerLifecycleStateDisconnected,
	"maintenance_in_progress": InfrastructureServerLifecycleStateMaintenanceInProgress,
}

// GetInfrastructureServerLifecycleStateEnumValues Enumerates the set of values for InfrastructureServerLifecycleStateEnum
func GetInfrastructureServerLifecycleStateEnumValues() []InfrastructureServerLifecycleStateEnum {
	values := make([]InfrastructureServerLifecycleStateEnum, 0)
	for _, v := range mappingInfrastructureServerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInfrastructureServerLifecycleStateEnumStringValues Enumerates the set of values in String for InfrastructureServerLifecycleStateEnum
func GetInfrastructureServerLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"REQUIRES_VALIDATION",
		"VALIDATING",
		"VALIDATION_FAILED",
		"REQUIRES_ACTIVATION",
		"ACTIVATING",
		"ACTIVE",
		"ACTIVATION_FAILED",
		"FAILED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"DISCONNECTED",
		"MAINTENANCE_IN_PROGRESS",
	}
}

// GetMappingInfrastructureServerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfrastructureServerLifecycleStateEnum(val string) (InfrastructureServerLifecycleStateEnum, bool) {
	enum, ok := mappingInfrastructureServerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
