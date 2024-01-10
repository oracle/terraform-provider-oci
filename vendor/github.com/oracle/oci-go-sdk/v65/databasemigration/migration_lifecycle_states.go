// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// MigrationLifecycleStatesEnum Enum with underlying type: string
type MigrationLifecycleStatesEnum string

// Set of constants representing the allowable values for MigrationLifecycleStatesEnum
const (
	MigrationLifecycleStatesCreating       MigrationLifecycleStatesEnum = "CREATING"
	MigrationLifecycleStatesUpdating       MigrationLifecycleStatesEnum = "UPDATING"
	MigrationLifecycleStatesActive         MigrationLifecycleStatesEnum = "ACTIVE"
	MigrationLifecycleStatesInProgress     MigrationLifecycleStatesEnum = "IN_PROGRESS"
	MigrationLifecycleStatesAccepted       MigrationLifecycleStatesEnum = "ACCEPTED"
	MigrationLifecycleStatesSucceeded      MigrationLifecycleStatesEnum = "SUCCEEDED"
	MigrationLifecycleStatesCanceled       MigrationLifecycleStatesEnum = "CANCELED"
	MigrationLifecycleStatesWaiting        MigrationLifecycleStatesEnum = "WAITING"
	MigrationLifecycleStatesNeedsAttention MigrationLifecycleStatesEnum = "NEEDS_ATTENTION"
	MigrationLifecycleStatesInactive       MigrationLifecycleStatesEnum = "INACTIVE"
	MigrationLifecycleStatesDeleting       MigrationLifecycleStatesEnum = "DELETING"
	MigrationLifecycleStatesDeleted        MigrationLifecycleStatesEnum = "DELETED"
	MigrationLifecycleStatesFailed         MigrationLifecycleStatesEnum = "FAILED"
)

var mappingMigrationLifecycleStatesEnum = map[string]MigrationLifecycleStatesEnum{
	"CREATING":        MigrationLifecycleStatesCreating,
	"UPDATING":        MigrationLifecycleStatesUpdating,
	"ACTIVE":          MigrationLifecycleStatesActive,
	"IN_PROGRESS":     MigrationLifecycleStatesInProgress,
	"ACCEPTED":        MigrationLifecycleStatesAccepted,
	"SUCCEEDED":       MigrationLifecycleStatesSucceeded,
	"CANCELED":        MigrationLifecycleStatesCanceled,
	"WAITING":         MigrationLifecycleStatesWaiting,
	"NEEDS_ATTENTION": MigrationLifecycleStatesNeedsAttention,
	"INACTIVE":        MigrationLifecycleStatesInactive,
	"DELETING":        MigrationLifecycleStatesDeleting,
	"DELETED":         MigrationLifecycleStatesDeleted,
	"FAILED":          MigrationLifecycleStatesFailed,
}

var mappingMigrationLifecycleStatesEnumLowerCase = map[string]MigrationLifecycleStatesEnum{
	"creating":        MigrationLifecycleStatesCreating,
	"updating":        MigrationLifecycleStatesUpdating,
	"active":          MigrationLifecycleStatesActive,
	"in_progress":     MigrationLifecycleStatesInProgress,
	"accepted":        MigrationLifecycleStatesAccepted,
	"succeeded":       MigrationLifecycleStatesSucceeded,
	"canceled":        MigrationLifecycleStatesCanceled,
	"waiting":         MigrationLifecycleStatesWaiting,
	"needs_attention": MigrationLifecycleStatesNeedsAttention,
	"inactive":        MigrationLifecycleStatesInactive,
	"deleting":        MigrationLifecycleStatesDeleting,
	"deleted":         MigrationLifecycleStatesDeleted,
	"failed":          MigrationLifecycleStatesFailed,
}

// GetMigrationLifecycleStatesEnumValues Enumerates the set of values for MigrationLifecycleStatesEnum
func GetMigrationLifecycleStatesEnumValues() []MigrationLifecycleStatesEnum {
	values := make([]MigrationLifecycleStatesEnum, 0)
	for _, v := range mappingMigrationLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationLifecycleStatesEnumStringValues Enumerates the set of values in String for MigrationLifecycleStatesEnum
func GetMigrationLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"IN_PROGRESS",
		"ACCEPTED",
		"SUCCEEDED",
		"CANCELED",
		"WAITING",
		"NEEDS_ATTENTION",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMigrationLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationLifecycleStatesEnum(val string) (MigrationLifecycleStatesEnum, bool) {
	enum, ok := mappingMigrationLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
