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

// DeletionPolicyDetails The Deletion policy for the DB System.
type DeletionPolicyDetails struct {

	// Specifies if automatic backups needs to be retained after dbsystem is deleted.
	AutomaticBackupRetention DeletionPolicyDetailsAutomaticBackupRetentionEnum `mandatory:"true" json:"automaticBackupRetention"`

	// Specifies whether backup needs to be performed before deleting the DbSystem.
	FinalBackup DeletionPolicyDetailsFinalBackupEnum `mandatory:"true" json:"finalBackup"`

	// This flag protects against deleting a DbSystem.
	IsDeleteProtected *bool `mandatory:"true" json:"isDeleteProtected"`
}

func (m DeletionPolicyDetails) String() string {
	return common.PointerString(m)
}

// DeletionPolicyDetailsAutomaticBackupRetentionEnum Enum with underlying type: string
type DeletionPolicyDetailsAutomaticBackupRetentionEnum string

// Set of constants representing the allowable values for DeletionPolicyDetailsAutomaticBackupRetentionEnum
const (
	DeletionPolicyDetailsAutomaticBackupRetentionDelete DeletionPolicyDetailsAutomaticBackupRetentionEnum = "DELETE"
	DeletionPolicyDetailsAutomaticBackupRetentionRetain DeletionPolicyDetailsAutomaticBackupRetentionEnum = "RETAIN"
)

var mappingDeletionPolicyDetailsAutomaticBackupRetention = map[string]DeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"DELETE": DeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"RETAIN": DeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

// GetDeletionPolicyDetailsAutomaticBackupRetentionEnumValues Enumerates the set of values for DeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetDeletionPolicyDetailsAutomaticBackupRetentionEnumValues() []DeletionPolicyDetailsAutomaticBackupRetentionEnum {
	values := make([]DeletionPolicyDetailsAutomaticBackupRetentionEnum, 0)
	for _, v := range mappingDeletionPolicyDetailsAutomaticBackupRetention {
		values = append(values, v)
	}
	return values
}

// DeletionPolicyDetailsFinalBackupEnum Enum with underlying type: string
type DeletionPolicyDetailsFinalBackupEnum string

// Set of constants representing the allowable values for DeletionPolicyDetailsFinalBackupEnum
const (
	DeletionPolicyDetailsFinalBackupDoNotTake          DeletionPolicyDetailsFinalBackupEnum = "DO_NOT_TAKE"
	DeletionPolicyDetailsFinalBackupTakeBeforeDeletion DeletionPolicyDetailsFinalBackupEnum = "TAKE_BEFORE_DELETION"
)

var mappingDeletionPolicyDetailsFinalBackup = map[string]DeletionPolicyDetailsFinalBackupEnum{
	"DO_NOT_TAKE":          DeletionPolicyDetailsFinalBackupDoNotTake,
	"TAKE_BEFORE_DELETION": DeletionPolicyDetailsFinalBackupTakeBeforeDeletion,
}

// GetDeletionPolicyDetailsFinalBackupEnumValues Enumerates the set of values for DeletionPolicyDetailsFinalBackupEnum
func GetDeletionPolicyDetailsFinalBackupEnumValues() []DeletionPolicyDetailsFinalBackupEnum {
	values := make([]DeletionPolicyDetailsFinalBackupEnum, 0)
	for _, v := range mappingDeletionPolicyDetailsFinalBackup {
		values = append(values, v)
	}
	return values
}
