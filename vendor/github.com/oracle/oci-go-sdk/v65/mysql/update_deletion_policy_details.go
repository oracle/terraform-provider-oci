// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDeletionPolicyDetails Policy for how the DB System and related resources should be handled at the time of its deletion.
type UpdateDeletionPolicyDetails struct {

	// Specifies if any automatic backups created for a DB System should be retained or deleted when the DB System is deleted.
	AutomaticBackupRetention UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum `mandatory:"false" json:"automaticBackupRetention,omitempty"`

	// Specifies whether or not a backup is taken when the DB System is deleted.
	//   REQUIRE_FINAL_BACKUP: a backup is taken if the DB System is deleted.
	//   SKIP_FINAL_BACKUP: a backup is not taken if the DB System is deleted.
	FinalBackup UpdateDeletionPolicyDetailsFinalBackupEnum `mandatory:"false" json:"finalBackup,omitempty"`

	// Specifies whether the DB System can be deleted. Set to true to prevent deletion, false (default) to allow.
	IsDeleteProtected *bool `mandatory:"false" json:"isDeleteProtected"`
}

func (m UpdateDeletionPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDeletionPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum(string(m.AutomaticBackupRetention)); !ok && m.AutomaticBackupRetention != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutomaticBackupRetention: %s. Supported values are: %s.", m.AutomaticBackupRetention, strings.Join(GetUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateDeletionPolicyDetailsFinalBackupEnum(string(m.FinalBackup)); !ok && m.FinalBackup != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FinalBackup: %s. Supported values are: %s.", m.FinalBackup, strings.Join(GetUpdateDeletionPolicyDetailsFinalBackupEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum Enum with underlying type: string
type UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum string

// Set of constants representing the allowable values for UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum
const (
	UpdateDeletionPolicyDetailsAutomaticBackupRetentionDelete UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum = "DELETE"
	UpdateDeletionPolicyDetailsAutomaticBackupRetentionRetain UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum = "RETAIN"
)

var mappingUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum = map[string]UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"DELETE": UpdateDeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"RETAIN": UpdateDeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

var mappingUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumLowerCase = map[string]UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"delete": UpdateDeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"retain": UpdateDeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

// GetUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumValues Enumerates the set of values for UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumValues() []UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum {
	values := make([]UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum, 0)
	for _, v := range mappingUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues Enumerates the set of values in String for UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues() []string {
	return []string{
		"DELETE",
		"RETAIN",
	}
}

// GetMappingUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum(val string) (UpdateDeletionPolicyDetailsAutomaticBackupRetentionEnum, bool) {
	enum, ok := mappingUpdateDeletionPolicyDetailsAutomaticBackupRetentionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateDeletionPolicyDetailsFinalBackupEnum Enum with underlying type: string
type UpdateDeletionPolicyDetailsFinalBackupEnum string

// Set of constants representing the allowable values for UpdateDeletionPolicyDetailsFinalBackupEnum
const (
	UpdateDeletionPolicyDetailsFinalBackupSkipFinalBackup    UpdateDeletionPolicyDetailsFinalBackupEnum = "SKIP_FINAL_BACKUP"
	UpdateDeletionPolicyDetailsFinalBackupRequireFinalBackup UpdateDeletionPolicyDetailsFinalBackupEnum = "REQUIRE_FINAL_BACKUP"
)

var mappingUpdateDeletionPolicyDetailsFinalBackupEnum = map[string]UpdateDeletionPolicyDetailsFinalBackupEnum{
	"SKIP_FINAL_BACKUP":    UpdateDeletionPolicyDetailsFinalBackupSkipFinalBackup,
	"REQUIRE_FINAL_BACKUP": UpdateDeletionPolicyDetailsFinalBackupRequireFinalBackup,
}

var mappingUpdateDeletionPolicyDetailsFinalBackupEnumLowerCase = map[string]UpdateDeletionPolicyDetailsFinalBackupEnum{
	"skip_final_backup":    UpdateDeletionPolicyDetailsFinalBackupSkipFinalBackup,
	"require_final_backup": UpdateDeletionPolicyDetailsFinalBackupRequireFinalBackup,
}

// GetUpdateDeletionPolicyDetailsFinalBackupEnumValues Enumerates the set of values for UpdateDeletionPolicyDetailsFinalBackupEnum
func GetUpdateDeletionPolicyDetailsFinalBackupEnumValues() []UpdateDeletionPolicyDetailsFinalBackupEnum {
	values := make([]UpdateDeletionPolicyDetailsFinalBackupEnum, 0)
	for _, v := range mappingUpdateDeletionPolicyDetailsFinalBackupEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDeletionPolicyDetailsFinalBackupEnumStringValues Enumerates the set of values in String for UpdateDeletionPolicyDetailsFinalBackupEnum
func GetUpdateDeletionPolicyDetailsFinalBackupEnumStringValues() []string {
	return []string{
		"SKIP_FINAL_BACKUP",
		"REQUIRE_FINAL_BACKUP",
	}
}

// GetMappingUpdateDeletionPolicyDetailsFinalBackupEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDeletionPolicyDetailsFinalBackupEnum(val string) (UpdateDeletionPolicyDetailsFinalBackupEnum, bool) {
	enum, ok := mappingUpdateDeletionPolicyDetailsFinalBackupEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
