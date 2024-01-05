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

// DeletionPolicyDetails The Deletion policy for the DB System.
type DeletionPolicyDetails struct {

	// Specifies if any automatic backups created for a DB System should be retained or deleted when the DB System is deleted.
	AutomaticBackupRetention DeletionPolicyDetailsAutomaticBackupRetentionEnum `mandatory:"true" json:"automaticBackupRetention"`

	// Specifies whether or not a backup is taken when the DB System is deleted.
	//   REQUIRE_FINAL_BACKUP: a backup is taken if the DB System is deleted.
	//   SKIP_FINAL_BACKUP: a backup is not taken if the DB System is deleted.
	FinalBackup DeletionPolicyDetailsFinalBackupEnum `mandatory:"true" json:"finalBackup"`

	// Specifies whether the DB System can be deleted. Set to true to prevent deletion, false (default) to allow.
	IsDeleteProtected *bool `mandatory:"true" json:"isDeleteProtected"`
}

func (m DeletionPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeletionPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeletionPolicyDetailsAutomaticBackupRetentionEnum(string(m.AutomaticBackupRetention)); !ok && m.AutomaticBackupRetention != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutomaticBackupRetention: %s. Supported values are: %s.", m.AutomaticBackupRetention, strings.Join(GetDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeletionPolicyDetailsFinalBackupEnum(string(m.FinalBackup)); !ok && m.FinalBackup != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FinalBackup: %s. Supported values are: %s.", m.FinalBackup, strings.Join(GetDeletionPolicyDetailsFinalBackupEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeletionPolicyDetailsAutomaticBackupRetentionEnum Enum with underlying type: string
type DeletionPolicyDetailsAutomaticBackupRetentionEnum string

// Set of constants representing the allowable values for DeletionPolicyDetailsAutomaticBackupRetentionEnum
const (
	DeletionPolicyDetailsAutomaticBackupRetentionDelete DeletionPolicyDetailsAutomaticBackupRetentionEnum = "DELETE"
	DeletionPolicyDetailsAutomaticBackupRetentionRetain DeletionPolicyDetailsAutomaticBackupRetentionEnum = "RETAIN"
)

var mappingDeletionPolicyDetailsAutomaticBackupRetentionEnum = map[string]DeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"DELETE": DeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"RETAIN": DeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

var mappingDeletionPolicyDetailsAutomaticBackupRetentionEnumLowerCase = map[string]DeletionPolicyDetailsAutomaticBackupRetentionEnum{
	"delete": DeletionPolicyDetailsAutomaticBackupRetentionDelete,
	"retain": DeletionPolicyDetailsAutomaticBackupRetentionRetain,
}

// GetDeletionPolicyDetailsAutomaticBackupRetentionEnumValues Enumerates the set of values for DeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetDeletionPolicyDetailsAutomaticBackupRetentionEnumValues() []DeletionPolicyDetailsAutomaticBackupRetentionEnum {
	values := make([]DeletionPolicyDetailsAutomaticBackupRetentionEnum, 0)
	for _, v := range mappingDeletionPolicyDetailsAutomaticBackupRetentionEnum {
		values = append(values, v)
	}
	return values
}

// GetDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues Enumerates the set of values in String for DeletionPolicyDetailsAutomaticBackupRetentionEnum
func GetDeletionPolicyDetailsAutomaticBackupRetentionEnumStringValues() []string {
	return []string{
		"DELETE",
		"RETAIN",
	}
}

// GetMappingDeletionPolicyDetailsAutomaticBackupRetentionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeletionPolicyDetailsAutomaticBackupRetentionEnum(val string) (DeletionPolicyDetailsAutomaticBackupRetentionEnum, bool) {
	enum, ok := mappingDeletionPolicyDetailsAutomaticBackupRetentionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DeletionPolicyDetailsFinalBackupEnum Enum with underlying type: string
type DeletionPolicyDetailsFinalBackupEnum string

// Set of constants representing the allowable values for DeletionPolicyDetailsFinalBackupEnum
const (
	DeletionPolicyDetailsFinalBackupSkipFinalBackup    DeletionPolicyDetailsFinalBackupEnum = "SKIP_FINAL_BACKUP"
	DeletionPolicyDetailsFinalBackupRequireFinalBackup DeletionPolicyDetailsFinalBackupEnum = "REQUIRE_FINAL_BACKUP"
)

var mappingDeletionPolicyDetailsFinalBackupEnum = map[string]DeletionPolicyDetailsFinalBackupEnum{
	"SKIP_FINAL_BACKUP":    DeletionPolicyDetailsFinalBackupSkipFinalBackup,
	"REQUIRE_FINAL_BACKUP": DeletionPolicyDetailsFinalBackupRequireFinalBackup,
}

var mappingDeletionPolicyDetailsFinalBackupEnumLowerCase = map[string]DeletionPolicyDetailsFinalBackupEnum{
	"skip_final_backup":    DeletionPolicyDetailsFinalBackupSkipFinalBackup,
	"require_final_backup": DeletionPolicyDetailsFinalBackupRequireFinalBackup,
}

// GetDeletionPolicyDetailsFinalBackupEnumValues Enumerates the set of values for DeletionPolicyDetailsFinalBackupEnum
func GetDeletionPolicyDetailsFinalBackupEnumValues() []DeletionPolicyDetailsFinalBackupEnum {
	values := make([]DeletionPolicyDetailsFinalBackupEnum, 0)
	for _, v := range mappingDeletionPolicyDetailsFinalBackupEnum {
		values = append(values, v)
	}
	return values
}

// GetDeletionPolicyDetailsFinalBackupEnumStringValues Enumerates the set of values in String for DeletionPolicyDetailsFinalBackupEnum
func GetDeletionPolicyDetailsFinalBackupEnumStringValues() []string {
	return []string{
		"SKIP_FINAL_BACKUP",
		"REQUIRE_FINAL_BACKUP",
	}
}

// GetMappingDeletionPolicyDetailsFinalBackupEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeletionPolicyDetailsFinalBackupEnum(val string) (DeletionPolicyDetailsFinalBackupEnum, bool) {
	enum, ok := mappingDeletionPolicyDetailsFinalBackupEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
