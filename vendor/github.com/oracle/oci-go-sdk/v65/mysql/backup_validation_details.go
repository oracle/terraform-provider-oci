// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// BackupValidationDetails Backup validation details.
type BackupValidationDetails struct {

	// The status of backup validation:
	// NOT_VALIDATED (Default): The backup has not been validated.
	// VALIDATED: The backup has been validated successfully.
	// NEEDS_ATTENTION: The backup validation failed due to a transient issue. Validation should be retried.
	// FAILED: The backup cannot be restored.
	ValidationStatus BackupValidationDetailsValidationStatusEnum `mandatory:"false" json:"validationStatus,omitempty"`

	// The date and time of the most recent validation performed on the backup.
	TimeLastValidated *common.SDKTime `mandatory:"false" json:"timeLastValidated"`

	// The estimated restore duration of the backup.
	EstimatedRestoreDuration *string `mandatory:"false" json:"estimatedRestoreDuration"`

	// Error message if the backup validation has failed.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// Indicates whether the backup has been prepared successfully.
	// PREPARED: The backup is prepared one.
	// NOT_PREPARED: The backup is not prepared.
	BackupPreparationStatus BackupValidationDetailsBackupPreparationStatusEnum `mandatory:"false" json:"backupPreparationStatus,omitempty"`

	PreparedBackupDetails *PreparedBackupDetails `mandatory:"false" json:"preparedBackupDetails"`
}

func (m BackupValidationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupValidationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBackupValidationDetailsValidationStatusEnum(string(m.ValidationStatus)); !ok && m.ValidationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValidationStatus: %s. Supported values are: %s.", m.ValidationStatus, strings.Join(GetBackupValidationDetailsValidationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBackupValidationDetailsBackupPreparationStatusEnum(string(m.BackupPreparationStatus)); !ok && m.BackupPreparationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupPreparationStatus: %s. Supported values are: %s.", m.BackupPreparationStatus, strings.Join(GetBackupValidationDetailsBackupPreparationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupValidationDetailsValidationStatusEnum Enum with underlying type: string
type BackupValidationDetailsValidationStatusEnum string

// Set of constants representing the allowable values for BackupValidationDetailsValidationStatusEnum
const (
	BackupValidationDetailsValidationStatusNotValidated   BackupValidationDetailsValidationStatusEnum = "NOT_VALIDATED"
	BackupValidationDetailsValidationStatusValidated      BackupValidationDetailsValidationStatusEnum = "VALIDATED"
	BackupValidationDetailsValidationStatusNeedsAttention BackupValidationDetailsValidationStatusEnum = "NEEDS_ATTENTION"
	BackupValidationDetailsValidationStatusFailed         BackupValidationDetailsValidationStatusEnum = "FAILED"
)

var mappingBackupValidationDetailsValidationStatusEnum = map[string]BackupValidationDetailsValidationStatusEnum{
	"NOT_VALIDATED":   BackupValidationDetailsValidationStatusNotValidated,
	"VALIDATED":       BackupValidationDetailsValidationStatusValidated,
	"NEEDS_ATTENTION": BackupValidationDetailsValidationStatusNeedsAttention,
	"FAILED":          BackupValidationDetailsValidationStatusFailed,
}

var mappingBackupValidationDetailsValidationStatusEnumLowerCase = map[string]BackupValidationDetailsValidationStatusEnum{
	"not_validated":   BackupValidationDetailsValidationStatusNotValidated,
	"validated":       BackupValidationDetailsValidationStatusValidated,
	"needs_attention": BackupValidationDetailsValidationStatusNeedsAttention,
	"failed":          BackupValidationDetailsValidationStatusFailed,
}

// GetBackupValidationDetailsValidationStatusEnumValues Enumerates the set of values for BackupValidationDetailsValidationStatusEnum
func GetBackupValidationDetailsValidationStatusEnumValues() []BackupValidationDetailsValidationStatusEnum {
	values := make([]BackupValidationDetailsValidationStatusEnum, 0)
	for _, v := range mappingBackupValidationDetailsValidationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupValidationDetailsValidationStatusEnumStringValues Enumerates the set of values in String for BackupValidationDetailsValidationStatusEnum
func GetBackupValidationDetailsValidationStatusEnumStringValues() []string {
	return []string{
		"NOT_VALIDATED",
		"VALIDATED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingBackupValidationDetailsValidationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupValidationDetailsValidationStatusEnum(val string) (BackupValidationDetailsValidationStatusEnum, bool) {
	enum, ok := mappingBackupValidationDetailsValidationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BackupValidationDetailsBackupPreparationStatusEnum Enum with underlying type: string
type BackupValidationDetailsBackupPreparationStatusEnum string

// Set of constants representing the allowable values for BackupValidationDetailsBackupPreparationStatusEnum
const (
	BackupValidationDetailsBackupPreparationStatusPrepared    BackupValidationDetailsBackupPreparationStatusEnum = "PREPARED"
	BackupValidationDetailsBackupPreparationStatusNotPrepared BackupValidationDetailsBackupPreparationStatusEnum = "NOT_PREPARED"
)

var mappingBackupValidationDetailsBackupPreparationStatusEnum = map[string]BackupValidationDetailsBackupPreparationStatusEnum{
	"PREPARED":     BackupValidationDetailsBackupPreparationStatusPrepared,
	"NOT_PREPARED": BackupValidationDetailsBackupPreparationStatusNotPrepared,
}

var mappingBackupValidationDetailsBackupPreparationStatusEnumLowerCase = map[string]BackupValidationDetailsBackupPreparationStatusEnum{
	"prepared":     BackupValidationDetailsBackupPreparationStatusPrepared,
	"not_prepared": BackupValidationDetailsBackupPreparationStatusNotPrepared,
}

// GetBackupValidationDetailsBackupPreparationStatusEnumValues Enumerates the set of values for BackupValidationDetailsBackupPreparationStatusEnum
func GetBackupValidationDetailsBackupPreparationStatusEnumValues() []BackupValidationDetailsBackupPreparationStatusEnum {
	values := make([]BackupValidationDetailsBackupPreparationStatusEnum, 0)
	for _, v := range mappingBackupValidationDetailsBackupPreparationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupValidationDetailsBackupPreparationStatusEnumStringValues Enumerates the set of values in String for BackupValidationDetailsBackupPreparationStatusEnum
func GetBackupValidationDetailsBackupPreparationStatusEnumStringValues() []string {
	return []string{
		"PREPARED",
		"NOT_PREPARED",
	}
}

// GetMappingBackupValidationDetailsBackupPreparationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupValidationDetailsBackupPreparationStatusEnum(val string) (BackupValidationDetailsBackupPreparationStatusEnum, bool) {
	enum, ok := mappingBackupValidationDetailsBackupPreparationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
