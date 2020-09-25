// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// CreateDeletionPolicyDetails Policy for how the DBSystem and related resources should be handled at the time of its deletion.
type CreateDeletionPolicyDetails struct {

	// Specifies if automatic backups needs to be retained after dbsystem is deleted.
	AutomaticBackupRetention CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum `mandatory:"false" json:"automaticBackupRetention,omitempty"`

	// Specifies whether backup needs to be performed before deleting the DbSystem.
	FinalBackup CreateDeletionPolicyDetailsFinalBackupEnum `mandatory:"false" json:"finalBackup,omitempty"`

	// This flag protects against deleting a DbSystem.
	IsDeleteProtected *bool `mandatory:"false" json:"isDeleteProtected"`
}

func (m CreateDeletionPolicyDetails) String() string {
	return common.PointerString(m)
}

// CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum Enum with underlying type: string
type CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum string

// Set of constants representing the allowable values for CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum
const (
	CreateDeletionPolicyDetailsAutomaticBackupRetentionDelete CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum = "DELETE"
	CreateDeletionPolicyDetailsAutomaticBackupRetentionRetain CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum = "RETAIN"
)

var mappingCreateDeletionPolicyDetailsAutomaticBackupRetention = map[string]CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"DELETE": CreateDeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"RETAIN": CreateDeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

// GetCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumValues Enumerates the set of values for CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumValues() []CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum {
	values := make([]CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum, 0)
	for _, v := range mappingCreateDeletionPolicyDetailsAutomaticBackupRetention {
		values = append(values, v)
	}
	return values
}

// CreateDeletionPolicyDetailsFinalBackupEnum Enum with underlying type: string
type CreateDeletionPolicyDetailsFinalBackupEnum string

// Set of constants representing the allowable values for CreateDeletionPolicyDetailsFinalBackupEnum
const (
	CreateDeletionPolicyDetailsFinalBackupDoNotTake          CreateDeletionPolicyDetailsFinalBackupEnum = "DO_NOT_TAKE"
	CreateDeletionPolicyDetailsFinalBackupTakeBeforeDeletion CreateDeletionPolicyDetailsFinalBackupEnum = "TAKE_BEFORE_DELETION"
)

var mappingCreateDeletionPolicyDetailsFinalBackup = map[string]CreateDeletionPolicyDetailsFinalBackupEnum{
	"DO_NOT_TAKE":          CreateDeletionPolicyDetailsFinalBackupDoNotTake,
	"TAKE_BEFORE_DELETION": CreateDeletionPolicyDetailsFinalBackupTakeBeforeDeletion,
}

// GetCreateDeletionPolicyDetailsFinalBackupEnumValues Enumerates the set of values for CreateDeletionPolicyDetailsFinalBackupEnum
func GetCreateDeletionPolicyDetailsFinalBackupEnumValues() []CreateDeletionPolicyDetailsFinalBackupEnum {
	values := make([]CreateDeletionPolicyDetailsFinalBackupEnum, 0)
	for _, v := range mappingCreateDeletionPolicyDetailsFinalBackup {
		values = append(values, v)
	}
	return values
}
