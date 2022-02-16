// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// AuditTrailLifecycleStateEnum Enum with underlying type: string
type AuditTrailLifecycleStateEnum string

// Set of constants representing the allowable values for AuditTrailLifecycleStateEnum
const (
	AuditTrailLifecycleStateInactive       AuditTrailLifecycleStateEnum = "INACTIVE"
	AuditTrailLifecycleStateUpdating       AuditTrailLifecycleStateEnum = "UPDATING"
	AuditTrailLifecycleStateActive         AuditTrailLifecycleStateEnum = "ACTIVE"
	AuditTrailLifecycleStateDeleting       AuditTrailLifecycleStateEnum = "DELETING"
	AuditTrailLifecycleStateFailed         AuditTrailLifecycleStateEnum = "FAILED"
	AuditTrailLifecycleStateNeedsAttention AuditTrailLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingAuditTrailLifecycleStateEnum = map[string]AuditTrailLifecycleStateEnum{
	"INACTIVE":        AuditTrailLifecycleStateInactive,
	"UPDATING":        AuditTrailLifecycleStateUpdating,
	"ACTIVE":          AuditTrailLifecycleStateActive,
	"DELETING":        AuditTrailLifecycleStateDeleting,
	"FAILED":          AuditTrailLifecycleStateFailed,
	"NEEDS_ATTENTION": AuditTrailLifecycleStateNeedsAttention,
}

// GetAuditTrailLifecycleStateEnumValues Enumerates the set of values for AuditTrailLifecycleStateEnum
func GetAuditTrailLifecycleStateEnumValues() []AuditTrailLifecycleStateEnum {
	values := make([]AuditTrailLifecycleStateEnum, 0)
	for _, v := range mappingAuditTrailLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditTrailLifecycleStateEnumStringValues Enumerates the set of values in String for AuditTrailLifecycleStateEnum
func GetAuditTrailLifecycleStateEnumStringValues() []string {
	return []string{
		"INACTIVE",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingAuditTrailLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditTrailLifecycleStateEnum(val string) (AuditTrailLifecycleStateEnum, bool) {
	mappingAuditTrailLifecycleStateEnumIgnoreCase := make(map[string]AuditTrailLifecycleStateEnum)
	for k, v := range mappingAuditTrailLifecycleStateEnum {
		mappingAuditTrailLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAuditTrailLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
