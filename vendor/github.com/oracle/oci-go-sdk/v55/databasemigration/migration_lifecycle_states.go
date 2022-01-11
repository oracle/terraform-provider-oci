// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

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

var mappingMigrationLifecycleStates = map[string]MigrationLifecycleStatesEnum{
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

// GetMigrationLifecycleStatesEnumValues Enumerates the set of values for MigrationLifecycleStatesEnum
func GetMigrationLifecycleStatesEnumValues() []MigrationLifecycleStatesEnum {
	values := make([]MigrationLifecycleStatesEnum, 0)
	for _, v := range mappingMigrationLifecycleStates {
		values = append(values, v)
	}
	return values
}
