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

// CreateDeletionPolicyDetails Policy for how the DB System and related resources should be handled at the time of its deletion.
type CreateDeletionPolicyDetails struct {

	// Specifies if any automatic backups created for a DB System should be retained or deleted when the DB System is deleted.
	AutomaticBackupRetention CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum `mandatory:"false" json:"automaticBackupRetention,omitempty"`

	// Specifies whether or not a backup is taken when the DB System is deleted.
	//   REQUIRE_FINAL_BACKUP: a backup is taken if the DB System is deleted.
	//   SKIP_FINAL_BACKUP: a backup is not taken if the DB System is deleted.
	FinalBackup CreateDeletionPolicyDetailsFinalBackupEnum `mandatory:"false" json:"finalBackup,omitempty"`

	// Specifies whether the DB System can be deleted. Set to true to prevent deletion, false (default) to allow.
	IsDeleteProtected *bool `mandatory:"false" json:"isDeleteProtected"`
}

func (m CreateDeletionPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDeletionPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateDeletionPolicyDetailsAutomaticBackupRetentionEnum(string(m.AutomaticBackupRetention)); !ok && m.AutomaticBackupRetention != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutomaticBackupRetention: %s. Supported values are: %s.", m.AutomaticBackupRetention, strings.Join(GetCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateDeletionPolicyDetailsFinalBackupEnum(string(m.FinalBackup)); !ok && m.FinalBackup != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FinalBackup: %s. Supported values are: %s.", m.FinalBackup, strings.Join(GetCreateDeletionPolicyDetailsFinalBackupEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum Enum with underlying type: string
type CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum string

// Set of constants representing the allowable values for CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum
const (
	CreateDeletionPolicyDetailsAutomaticBackupRetentionDelete CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum = "DELETE"
	CreateDeletionPolicyDetailsAutomaticBackupRetentionRetain CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum = "RETAIN"
)

var mappingCreateDeletionPolicyDetailsAutomaticBackupRetentionEnum = map[string]CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"DELETE": CreateDeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"RETAIN": CreateDeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

var mappingCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumLowerCase = map[string]CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"delete": CreateDeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"retain": CreateDeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

// GetCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumValues Enumerates the set of values for CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumValues() []CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum {
	values := make([]CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum, 0)
	for _, v := range mappingCreateDeletionPolicyDetailsAutomaticBackupRetentionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues Enumerates the set of values in String for CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues() []string {
	return []string{
		"DELETE",
		"RETAIN",
	}
}

// GetMappingCreateDeletionPolicyDetailsAutomaticBackupRetentionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDeletionPolicyDetailsAutomaticBackupRetentionEnum(val string) (CreateDeletionPolicyDetailsAutomaticBackupRetentionEnum, bool) {
	enum, ok := mappingCreateDeletionPolicyDetailsAutomaticBackupRetentionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateDeletionPolicyDetailsFinalBackupEnum Enum with underlying type: string
type CreateDeletionPolicyDetailsFinalBackupEnum string

// Set of constants representing the allowable values for CreateDeletionPolicyDetailsFinalBackupEnum
const (
	CreateDeletionPolicyDetailsFinalBackupSkipFinalBackup    CreateDeletionPolicyDetailsFinalBackupEnum = "SKIP_FINAL_BACKUP"
	CreateDeletionPolicyDetailsFinalBackupRequireFinalBackup CreateDeletionPolicyDetailsFinalBackupEnum = "REQUIRE_FINAL_BACKUP"
)

var mappingCreateDeletionPolicyDetailsFinalBackupEnum = map[string]CreateDeletionPolicyDetailsFinalBackupEnum{
	"SKIP_FINAL_BACKUP":    CreateDeletionPolicyDetailsFinalBackupSkipFinalBackup,
	"REQUIRE_FINAL_BACKUP": CreateDeletionPolicyDetailsFinalBackupRequireFinalBackup,
}

var mappingCreateDeletionPolicyDetailsFinalBackupEnumLowerCase = map[string]CreateDeletionPolicyDetailsFinalBackupEnum{
	"skip_final_backup":    CreateDeletionPolicyDetailsFinalBackupSkipFinalBackup,
	"require_final_backup": CreateDeletionPolicyDetailsFinalBackupRequireFinalBackup,
}

// GetCreateDeletionPolicyDetailsFinalBackupEnumValues Enumerates the set of values for CreateDeletionPolicyDetailsFinalBackupEnum
func GetCreateDeletionPolicyDetailsFinalBackupEnumValues() []CreateDeletionPolicyDetailsFinalBackupEnum {
	values := make([]CreateDeletionPolicyDetailsFinalBackupEnum, 0)
	for _, v := range mappingCreateDeletionPolicyDetailsFinalBackupEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDeletionPolicyDetailsFinalBackupEnumStringValues Enumerates the set of values in String for CreateDeletionPolicyDetailsFinalBackupEnum
func GetCreateDeletionPolicyDetailsFinalBackupEnumStringValues() []string {
	return []string{
		"SKIP_FINAL_BACKUP",
		"REQUIRE_FINAL_BACKUP",
	}
}

// GetMappingCreateDeletionPolicyDetailsFinalBackupEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDeletionPolicyDetailsFinalBackupEnum(val string) (CreateDeletionPolicyDetailsFinalBackupEnum, bool) {
	enum, ok := mappingCreateDeletionPolicyDetailsFinalBackupEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
