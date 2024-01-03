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

// DrProtectionGroupLifecycleStateEnum Enum with underlying type: string
type DrProtectionGroupLifecycleStateEnum string

// Set of constants representing the allowable values for DrProtectionGroupLifecycleStateEnum
const (
	DrProtectionGroupLifecycleStateCreating       DrProtectionGroupLifecycleStateEnum = "CREATING"
	DrProtectionGroupLifecycleStateActive         DrProtectionGroupLifecycleStateEnum = "ACTIVE"
	DrProtectionGroupLifecycleStateUpdating       DrProtectionGroupLifecycleStateEnum = "UPDATING"
	DrProtectionGroupLifecycleStateInactive       DrProtectionGroupLifecycleStateEnum = "INACTIVE"
	DrProtectionGroupLifecycleStateNeedsAttention DrProtectionGroupLifecycleStateEnum = "NEEDS_ATTENTION"
	DrProtectionGroupLifecycleStateDeleting       DrProtectionGroupLifecycleStateEnum = "DELETING"
	DrProtectionGroupLifecycleStateDeleted        DrProtectionGroupLifecycleStateEnum = "DELETED"
	DrProtectionGroupLifecycleStateFailed         DrProtectionGroupLifecycleStateEnum = "FAILED"
)

var mappingDrProtectionGroupLifecycleStateEnum = map[string]DrProtectionGroupLifecycleStateEnum{
	"CREATING":        DrProtectionGroupLifecycleStateCreating,
	"ACTIVE":          DrProtectionGroupLifecycleStateActive,
	"UPDATING":        DrProtectionGroupLifecycleStateUpdating,
	"INACTIVE":        DrProtectionGroupLifecycleStateInactive,
	"NEEDS_ATTENTION": DrProtectionGroupLifecycleStateNeedsAttention,
	"DELETING":        DrProtectionGroupLifecycleStateDeleting,
	"DELETED":         DrProtectionGroupLifecycleStateDeleted,
	"FAILED":          DrProtectionGroupLifecycleStateFailed,
}

var mappingDrProtectionGroupLifecycleStateEnumLowerCase = map[string]DrProtectionGroupLifecycleStateEnum{
	"creating":        DrProtectionGroupLifecycleStateCreating,
	"active":          DrProtectionGroupLifecycleStateActive,
	"updating":        DrProtectionGroupLifecycleStateUpdating,
	"inactive":        DrProtectionGroupLifecycleStateInactive,
	"needs_attention": DrProtectionGroupLifecycleStateNeedsAttention,
	"deleting":        DrProtectionGroupLifecycleStateDeleting,
	"deleted":         DrProtectionGroupLifecycleStateDeleted,
	"failed":          DrProtectionGroupLifecycleStateFailed,
}

// GetDrProtectionGroupLifecycleStateEnumValues Enumerates the set of values for DrProtectionGroupLifecycleStateEnum
func GetDrProtectionGroupLifecycleStateEnumValues() []DrProtectionGroupLifecycleStateEnum {
	values := make([]DrProtectionGroupLifecycleStateEnum, 0)
	for _, v := range mappingDrProtectionGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDrProtectionGroupLifecycleStateEnumStringValues Enumerates the set of values in String for DrProtectionGroupLifecycleStateEnum
func GetDrProtectionGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDrProtectionGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrProtectionGroupLifecycleStateEnum(val string) (DrProtectionGroupLifecycleStateEnum, bool) {
	enum, ok := mappingDrProtectionGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
