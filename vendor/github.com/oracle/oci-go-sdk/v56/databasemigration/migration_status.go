// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// MigrationStatusEnum Enum with underlying type: string
type MigrationStatusEnum string

// Set of constants representing the allowable values for MigrationStatusEnum
const (
	MigrationStatusReady      MigrationStatusEnum = "READY"
	MigrationStatusAborting   MigrationStatusEnum = "ABORTING"
	MigrationStatusValidating MigrationStatusEnum = "VALIDATING"
	MigrationStatusValidated  MigrationStatusEnum = "VALIDATED"
	MigrationStatusWaiting    MigrationStatusEnum = "WAITING"
	MigrationStatusMigrating  MigrationStatusEnum = "MIGRATING"
	MigrationStatusDone       MigrationStatusEnum = "DONE"
)

var mappingMigrationStatus = map[string]MigrationStatusEnum{
	"READY":      MigrationStatusReady,
	"ABORTING":   MigrationStatusAborting,
	"VALIDATING": MigrationStatusValidating,
	"VALIDATED":  MigrationStatusValidated,
	"WAITING":    MigrationStatusWaiting,
	"MIGRATING":  MigrationStatusMigrating,
	"DONE":       MigrationStatusDone,
}

// GetMigrationStatusEnumValues Enumerates the set of values for MigrationStatusEnum
func GetMigrationStatusEnumValues() []MigrationStatusEnum {
	values := make([]MigrationStatusEnum, 0)
	for _, v := range mappingMigrationStatus {
		values = append(values, v)
	}
	return values
}
