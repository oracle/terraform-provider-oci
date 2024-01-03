// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AuditArchiveRetrievalLifecycleStateEnum Enum with underlying type: string
type AuditArchiveRetrievalLifecycleStateEnum string

// Set of constants representing the allowable values for AuditArchiveRetrievalLifecycleStateEnum
const (
	AuditArchiveRetrievalLifecycleStateCreating       AuditArchiveRetrievalLifecycleStateEnum = "CREATING"
	AuditArchiveRetrievalLifecycleStateActive         AuditArchiveRetrievalLifecycleStateEnum = "ACTIVE"
	AuditArchiveRetrievalLifecycleStateNeedsAttention AuditArchiveRetrievalLifecycleStateEnum = "NEEDS_ATTENTION"
	AuditArchiveRetrievalLifecycleStateFailed         AuditArchiveRetrievalLifecycleStateEnum = "FAILED"
	AuditArchiveRetrievalLifecycleStateDeleting       AuditArchiveRetrievalLifecycleStateEnum = "DELETING"
	AuditArchiveRetrievalLifecycleStateDeleted        AuditArchiveRetrievalLifecycleStateEnum = "DELETED"
	AuditArchiveRetrievalLifecycleStateUpdating       AuditArchiveRetrievalLifecycleStateEnum = "UPDATING"
)

var mappingAuditArchiveRetrievalLifecycleStateEnum = map[string]AuditArchiveRetrievalLifecycleStateEnum{
	"CREATING":        AuditArchiveRetrievalLifecycleStateCreating,
	"ACTIVE":          AuditArchiveRetrievalLifecycleStateActive,
	"NEEDS_ATTENTION": AuditArchiveRetrievalLifecycleStateNeedsAttention,
	"FAILED":          AuditArchiveRetrievalLifecycleStateFailed,
	"DELETING":        AuditArchiveRetrievalLifecycleStateDeleting,
	"DELETED":         AuditArchiveRetrievalLifecycleStateDeleted,
	"UPDATING":        AuditArchiveRetrievalLifecycleStateUpdating,
}

var mappingAuditArchiveRetrievalLifecycleStateEnumLowerCase = map[string]AuditArchiveRetrievalLifecycleStateEnum{
	"creating":        AuditArchiveRetrievalLifecycleStateCreating,
	"active":          AuditArchiveRetrievalLifecycleStateActive,
	"needs_attention": AuditArchiveRetrievalLifecycleStateNeedsAttention,
	"failed":          AuditArchiveRetrievalLifecycleStateFailed,
	"deleting":        AuditArchiveRetrievalLifecycleStateDeleting,
	"deleted":         AuditArchiveRetrievalLifecycleStateDeleted,
	"updating":        AuditArchiveRetrievalLifecycleStateUpdating,
}

// GetAuditArchiveRetrievalLifecycleStateEnumValues Enumerates the set of values for AuditArchiveRetrievalLifecycleStateEnum
func GetAuditArchiveRetrievalLifecycleStateEnumValues() []AuditArchiveRetrievalLifecycleStateEnum {
	values := make([]AuditArchiveRetrievalLifecycleStateEnum, 0)
	for _, v := range mappingAuditArchiveRetrievalLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditArchiveRetrievalLifecycleStateEnumStringValues Enumerates the set of values in String for AuditArchiveRetrievalLifecycleStateEnum
func GetAuditArchiveRetrievalLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
		"DELETED",
		"UPDATING",
	}
}

// GetMappingAuditArchiveRetrievalLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditArchiveRetrievalLifecycleStateEnum(val string) (AuditArchiveRetrievalLifecycleStateEnum, bool) {
	enum, ok := mappingAuditArchiveRetrievalLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
