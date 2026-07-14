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

// InfrastructureLifecycleStateEnum Enum with underlying type: string
type InfrastructureLifecycleStateEnum string

// Set of constants representing the allowable values for InfrastructureLifecycleStateEnum
const (
	InfrastructureLifecycleStateCreating              InfrastructureLifecycleStateEnum = "CREATING"
	InfrastructureLifecycleStateRequiresValidation    InfrastructureLifecycleStateEnum = "REQUIRES_VALIDATION"
	InfrastructureLifecycleStateValidating            InfrastructureLifecycleStateEnum = "VALIDATING"
	InfrastructureLifecycleStateValidationFailed      InfrastructureLifecycleStateEnum = "VALIDATION_FAILED"
	InfrastructureLifecycleStateRequiresActivation    InfrastructureLifecycleStateEnum = "REQUIRES_ACTIVATION"
	InfrastructureLifecycleStateActivating            InfrastructureLifecycleStateEnum = "ACTIVATING"
	InfrastructureLifecycleStateActive                InfrastructureLifecycleStateEnum = "ACTIVE"
	InfrastructureLifecycleStateActivationFailed      InfrastructureLifecycleStateEnum = "ACTIVATION_FAILED"
	InfrastructureLifecycleStateFailed                InfrastructureLifecycleStateEnum = "FAILED"
	InfrastructureLifecycleStateUpdating              InfrastructureLifecycleStateEnum = "UPDATING"
	InfrastructureLifecycleStateDeleting              InfrastructureLifecycleStateEnum = "DELETING"
	InfrastructureLifecycleStateDeleted               InfrastructureLifecycleStateEnum = "DELETED"
	InfrastructureLifecycleStateDisconnected          InfrastructureLifecycleStateEnum = "DISCONNECTED"
	InfrastructureLifecycleStateMaintenanceInProgress InfrastructureLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingInfrastructureLifecycleStateEnum = map[string]InfrastructureLifecycleStateEnum{
	"CREATING":                InfrastructureLifecycleStateCreating,
	"REQUIRES_VALIDATION":     InfrastructureLifecycleStateRequiresValidation,
	"VALIDATING":              InfrastructureLifecycleStateValidating,
	"VALIDATION_FAILED":       InfrastructureLifecycleStateValidationFailed,
	"REQUIRES_ACTIVATION":     InfrastructureLifecycleStateRequiresActivation,
	"ACTIVATING":              InfrastructureLifecycleStateActivating,
	"ACTIVE":                  InfrastructureLifecycleStateActive,
	"ACTIVATION_FAILED":       InfrastructureLifecycleStateActivationFailed,
	"FAILED":                  InfrastructureLifecycleStateFailed,
	"UPDATING":                InfrastructureLifecycleStateUpdating,
	"DELETING":                InfrastructureLifecycleStateDeleting,
	"DELETED":                 InfrastructureLifecycleStateDeleted,
	"DISCONNECTED":            InfrastructureLifecycleStateDisconnected,
	"MAINTENANCE_IN_PROGRESS": InfrastructureLifecycleStateMaintenanceInProgress,
}

var mappingInfrastructureLifecycleStateEnumLowerCase = map[string]InfrastructureLifecycleStateEnum{
	"creating":                InfrastructureLifecycleStateCreating,
	"requires_validation":     InfrastructureLifecycleStateRequiresValidation,
	"validating":              InfrastructureLifecycleStateValidating,
	"validation_failed":       InfrastructureLifecycleStateValidationFailed,
	"requires_activation":     InfrastructureLifecycleStateRequiresActivation,
	"activating":              InfrastructureLifecycleStateActivating,
	"active":                  InfrastructureLifecycleStateActive,
	"activation_failed":       InfrastructureLifecycleStateActivationFailed,
	"failed":                  InfrastructureLifecycleStateFailed,
	"updating":                InfrastructureLifecycleStateUpdating,
	"deleting":                InfrastructureLifecycleStateDeleting,
	"deleted":                 InfrastructureLifecycleStateDeleted,
	"disconnected":            InfrastructureLifecycleStateDisconnected,
	"maintenance_in_progress": InfrastructureLifecycleStateMaintenanceInProgress,
}

// GetInfrastructureLifecycleStateEnumValues Enumerates the set of values for InfrastructureLifecycleStateEnum
func GetInfrastructureLifecycleStateEnumValues() []InfrastructureLifecycleStateEnum {
	values := make([]InfrastructureLifecycleStateEnum, 0)
	for _, v := range mappingInfrastructureLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInfrastructureLifecycleStateEnumStringValues Enumerates the set of values in String for InfrastructureLifecycleStateEnum
func GetInfrastructureLifecycleStateEnumStringValues() []string {
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

// GetMappingInfrastructureLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInfrastructureLifecycleStateEnum(val string) (InfrastructureLifecycleStateEnum, bool) {
	enum, ok := mappingInfrastructureLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
