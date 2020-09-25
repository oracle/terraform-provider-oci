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

// UpdateDeletionPolicyDetails Policy for how the DBSystem and related resources should be handled at the time of its deletion.
type UpdateDeletionPolicyDetails struct {

	// Specifies if automatic backups needs to be retained after dbsystem is deleted.
	AutomaticBackupRetention UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum `mandatory:"false" json:"automaticBackupRetention,omitempty"`

	// Specifies whether backup needs to be performed before deleting the DbSystem.
	FinalBackup UpdateDeletionPolicyDetailsFinalBackupEnum `mandatory:"false" json:"finalBackup,omitempty"`

	// This flag protects against deleting a DbSystem.
	IsDeleteProtected *bool `mandatory:"false" json:"isDeleteProtected"`
}

func (m UpdateDeletionPolicyDetails) String() string {
	return common.PointerString(m)
}

// UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum Enum with underlying type: string
type UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum string

// Set of constants representing the allowable values for UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum
const (
	UpdateDeletionPolicyDetailsAutomaticBackupRetentionDelete UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum = "DELETE"
	UpdateDeletionPolicyDetailsAutomaticBackupRetentionRetain UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum = "RETAIN"
)

var mappingUpdateDeletionPolicyDetailsAutomaticBackupRetention = map[string]UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"DELETE": UpdateDeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"RETAIN": UpdateDeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

// GetUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumValues Enumerates the set of values for UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumValues() []UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum {
	values := make([]UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum, 0)
	for _, v := range mappingUpdateDeletionPolicyDetailsAutomaticBackupRetention {
		values = append(values, v)
	}
	return values
}

// UpdateDeletionPolicyDetailsFinalBackupEnum Enum with underlying type: string
type UpdateDeletionPolicyDetailsFinalBackupEnum string

// Set of constants representing the allowable values for UpdateDeletionPolicyDetailsFinalBackupEnum
const (
	UpdateDeletionPolicyDetailsFinalBackupDoNotTake          UpdateDeletionPolicyDetailsFinalBackupEnum = "DO_NOT_TAKE"
	UpdateDeletionPolicyDetailsFinalBackupTakeBeforeDeletion UpdateDeletionPolicyDetailsFinalBackupEnum = "TAKE_BEFORE_DELETION"
)

var mappingUpdateDeletionPolicyDetailsFinalBackup = map[string]UpdateDeletionPolicyDetailsFinalBackupEnum{
	"DO_NOT_TAKE":          UpdateDeletionPolicyDetailsFinalBackupDoNotTake,
	"TAKE_BEFORE_DELETION": UpdateDeletionPolicyDetailsFinalBackupTakeBeforeDeletion,
}

// GetUpdateDeletionPolicyDetailsFinalBackupEnumValues Enumerates the set of values for UpdateDeletionPolicyDetailsFinalBackupEnum
func GetUpdateDeletionPolicyDetailsFinalBackupEnumValues() []UpdateDeletionPolicyDetailsFinalBackupEnum {
	values := make([]UpdateDeletionPolicyDetailsFinalBackupEnum, 0)
	for _, v := range mappingUpdateDeletionPolicyDetailsFinalBackup {
		values = append(values, v)
	}
	return values
}
