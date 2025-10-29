// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// AutomaticDrConfigurationLifecycleStateEnum Enum with underlying type: string
type AutomaticDrConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for AutomaticDrConfigurationLifecycleStateEnum
const (
	AutomaticDrConfigurationLifecycleStateCreating       AutomaticDrConfigurationLifecycleStateEnum = "CREATING"
	AutomaticDrConfigurationLifecycleStateUpdating       AutomaticDrConfigurationLifecycleStateEnum = "UPDATING"
	AutomaticDrConfigurationLifecycleStateActive         AutomaticDrConfigurationLifecycleStateEnum = "ACTIVE"
	AutomaticDrConfigurationLifecycleStateInactive       AutomaticDrConfigurationLifecycleStateEnum = "INACTIVE"
	AutomaticDrConfigurationLifecycleStateNeedsAttention AutomaticDrConfigurationLifecycleStateEnum = "NEEDS_ATTENTION"
	AutomaticDrConfigurationLifecycleStateDeleting       AutomaticDrConfigurationLifecycleStateEnum = "DELETING"
	AutomaticDrConfigurationLifecycleStateDeleted        AutomaticDrConfigurationLifecycleStateEnum = "DELETED"
	AutomaticDrConfigurationLifecycleStateFailed         AutomaticDrConfigurationLifecycleStateEnum = "FAILED"
)

var mappingAutomaticDrConfigurationLifecycleStateEnum = map[string]AutomaticDrConfigurationLifecycleStateEnum{
	"CREATING":        AutomaticDrConfigurationLifecycleStateCreating,
	"UPDATING":        AutomaticDrConfigurationLifecycleStateUpdating,
	"ACTIVE":          AutomaticDrConfigurationLifecycleStateActive,
	"INACTIVE":        AutomaticDrConfigurationLifecycleStateInactive,
	"NEEDS_ATTENTION": AutomaticDrConfigurationLifecycleStateNeedsAttention,
	"DELETING":        AutomaticDrConfigurationLifecycleStateDeleting,
	"DELETED":         AutomaticDrConfigurationLifecycleStateDeleted,
	"FAILED":          AutomaticDrConfigurationLifecycleStateFailed,
}

var mappingAutomaticDrConfigurationLifecycleStateEnumLowerCase = map[string]AutomaticDrConfigurationLifecycleStateEnum{
	"creating":        AutomaticDrConfigurationLifecycleStateCreating,
	"updating":        AutomaticDrConfigurationLifecycleStateUpdating,
	"active":          AutomaticDrConfigurationLifecycleStateActive,
	"inactive":        AutomaticDrConfigurationLifecycleStateInactive,
	"needs_attention": AutomaticDrConfigurationLifecycleStateNeedsAttention,
	"deleting":        AutomaticDrConfigurationLifecycleStateDeleting,
	"deleted":         AutomaticDrConfigurationLifecycleStateDeleted,
	"failed":          AutomaticDrConfigurationLifecycleStateFailed,
}

// GetAutomaticDrConfigurationLifecycleStateEnumValues Enumerates the set of values for AutomaticDrConfigurationLifecycleStateEnum
func GetAutomaticDrConfigurationLifecycleStateEnumValues() []AutomaticDrConfigurationLifecycleStateEnum {
	values := make([]AutomaticDrConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingAutomaticDrConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutomaticDrConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for AutomaticDrConfigurationLifecycleStateEnum
func GetAutomaticDrConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAutomaticDrConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutomaticDrConfigurationLifecycleStateEnum(val string) (AutomaticDrConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingAutomaticDrConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
