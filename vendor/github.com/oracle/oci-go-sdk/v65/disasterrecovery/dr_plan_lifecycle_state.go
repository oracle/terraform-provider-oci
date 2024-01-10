// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DrPlanLifecycleStateEnum Enum with underlying type: string
type DrPlanLifecycleStateEnum string

// Set of constants representing the allowable values for DrPlanLifecycleStateEnum
const (
	DrPlanLifecycleStateCreating       DrPlanLifecycleStateEnum = "CREATING"
	DrPlanLifecycleStateUpdating       DrPlanLifecycleStateEnum = "UPDATING"
	DrPlanLifecycleStateActive         DrPlanLifecycleStateEnum = "ACTIVE"
	DrPlanLifecycleStateInactive       DrPlanLifecycleStateEnum = "INACTIVE"
	DrPlanLifecycleStateDeleting       DrPlanLifecycleStateEnum = "DELETING"
	DrPlanLifecycleStateDeleted        DrPlanLifecycleStateEnum = "DELETED"
	DrPlanLifecycleStateFailed         DrPlanLifecycleStateEnum = "FAILED"
	DrPlanLifecycleStateNeedsAttention DrPlanLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDrPlanLifecycleStateEnum = map[string]DrPlanLifecycleStateEnum{
	"CREATING":        DrPlanLifecycleStateCreating,
	"UPDATING":        DrPlanLifecycleStateUpdating,
	"ACTIVE":          DrPlanLifecycleStateActive,
	"INACTIVE":        DrPlanLifecycleStateInactive,
	"DELETING":        DrPlanLifecycleStateDeleting,
	"DELETED":         DrPlanLifecycleStateDeleted,
	"FAILED":          DrPlanLifecycleStateFailed,
	"NEEDS_ATTENTION": DrPlanLifecycleStateNeedsAttention,
}

var mappingDrPlanLifecycleStateEnumLowerCase = map[string]DrPlanLifecycleStateEnum{
	"creating":        DrPlanLifecycleStateCreating,
	"updating":        DrPlanLifecycleStateUpdating,
	"active":          DrPlanLifecycleStateActive,
	"inactive":        DrPlanLifecycleStateInactive,
	"deleting":        DrPlanLifecycleStateDeleting,
	"deleted":         DrPlanLifecycleStateDeleted,
	"failed":          DrPlanLifecycleStateFailed,
	"needs_attention": DrPlanLifecycleStateNeedsAttention,
}

// GetDrPlanLifecycleStateEnumValues Enumerates the set of values for DrPlanLifecycleStateEnum
func GetDrPlanLifecycleStateEnumValues() []DrPlanLifecycleStateEnum {
	values := make([]DrPlanLifecycleStateEnum, 0)
	for _, v := range mappingDrPlanLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanLifecycleStateEnumStringValues Enumerates the set of values in String for DrPlanLifecycleStateEnum
func GetDrPlanLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDrPlanLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanLifecycleStateEnum(val string) (DrPlanLifecycleStateEnum, bool) {
	enum, ok := mappingDrPlanLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
