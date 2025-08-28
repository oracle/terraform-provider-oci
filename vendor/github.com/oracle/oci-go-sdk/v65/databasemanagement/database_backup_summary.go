// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseBackupSummary The summary of the High Availability (HA) and backup for a database.
type DatabaseBackupSummary struct {

	// The backup status of the database.
	BackupStatus *string `mandatory:"true" json:"backupStatus"`

	// The database backup completion date and time in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	TimeBackupCompleted *common.SDKTime `mandatory:"true" json:"timeBackupCompleted"`

	// The backup duration of the database in seconds.
	BackupDurationInSeconds *int `mandatory:"true" json:"backupDurationInSeconds"`

	// The backup type of the database (FULL/INCREMENTAL).
	BackupType *string `mandatory:"true" json:"backupType"`

	// The backup destination of the database.
	BackupDestination DatabaseBackupSummaryBackupDestinationEnum `mandatory:"true" json:"backupDestination"`

	// The backup size of the database.
	BackupSizeInGBs *float32 `mandatory:"true" json:"backupSizeInGBs"`
}

func (m DatabaseBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseBackupSummaryBackupDestinationEnum(string(m.BackupDestination)); !ok && m.BackupDestination != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupDestination: %s. Supported values are: %s.", m.BackupDestination, strings.Join(GetDatabaseBackupSummaryBackupDestinationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseBackupSummaryBackupDestinationEnum Enum with underlying type: string
type DatabaseBackupSummaryBackupDestinationEnum string

// Set of constants representing the allowable values for DatabaseBackupSummaryBackupDestinationEnum
const (
	DatabaseBackupSummaryBackupDestinationDisk              DatabaseBackupSummaryBackupDestinationEnum = "DISK"
	DatabaseBackupSummaryBackupDestinationTape              DatabaseBackupSummaryBackupDestinationEnum = "TAPE"
	DatabaseBackupSummaryBackupDestinationNfs               DatabaseBackupSummaryBackupDestinationEnum = "NFS"
	DatabaseBackupSummaryBackupDestinationLocal             DatabaseBackupSummaryBackupDestinationEnum = "LOCAL"
	DatabaseBackupSummaryBackupDestinationDbrs              DatabaseBackupSummaryBackupDestinationEnum = "DBRS"
	DatabaseBackupSummaryBackupDestinationObjectStore       DatabaseBackupSummaryBackupDestinationEnum = "OBJECT_STORE"
	DatabaseBackupSummaryBackupDestinationRecoveryAppliance DatabaseBackupSummaryBackupDestinationEnum = "RECOVERY_APPLIANCE"
)

var mappingDatabaseBackupSummaryBackupDestinationEnum = map[string]DatabaseBackupSummaryBackupDestinationEnum{
	"DISK":               DatabaseBackupSummaryBackupDestinationDisk,
	"TAPE":               DatabaseBackupSummaryBackupDestinationTape,
	"NFS":                DatabaseBackupSummaryBackupDestinationNfs,
	"LOCAL":              DatabaseBackupSummaryBackupDestinationLocal,
	"DBRS":               DatabaseBackupSummaryBackupDestinationDbrs,
	"OBJECT_STORE":       DatabaseBackupSummaryBackupDestinationObjectStore,
	"RECOVERY_APPLIANCE": DatabaseBackupSummaryBackupDestinationRecoveryAppliance,
}

var mappingDatabaseBackupSummaryBackupDestinationEnumLowerCase = map[string]DatabaseBackupSummaryBackupDestinationEnum{
	"disk":               DatabaseBackupSummaryBackupDestinationDisk,
	"tape":               DatabaseBackupSummaryBackupDestinationTape,
	"nfs":                DatabaseBackupSummaryBackupDestinationNfs,
	"local":              DatabaseBackupSummaryBackupDestinationLocal,
	"dbrs":               DatabaseBackupSummaryBackupDestinationDbrs,
	"object_store":       DatabaseBackupSummaryBackupDestinationObjectStore,
	"recovery_appliance": DatabaseBackupSummaryBackupDestinationRecoveryAppliance,
}

// GetDatabaseBackupSummaryBackupDestinationEnumValues Enumerates the set of values for DatabaseBackupSummaryBackupDestinationEnum
func GetDatabaseBackupSummaryBackupDestinationEnumValues() []DatabaseBackupSummaryBackupDestinationEnum {
	values := make([]DatabaseBackupSummaryBackupDestinationEnum, 0)
	for _, v := range mappingDatabaseBackupSummaryBackupDestinationEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseBackupSummaryBackupDestinationEnumStringValues Enumerates the set of values in String for DatabaseBackupSummaryBackupDestinationEnum
func GetDatabaseBackupSummaryBackupDestinationEnumStringValues() []string {
	return []string{
		"DISK",
		"TAPE",
		"NFS",
		"LOCAL",
		"DBRS",
		"OBJECT_STORE",
		"RECOVERY_APPLIANCE",
	}
}

// GetMappingDatabaseBackupSummaryBackupDestinationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseBackupSummaryBackupDestinationEnum(val string) (DatabaseBackupSummaryBackupDestinationEnum, bool) {
	enum, ok := mappingDatabaseBackupSummaryBackupDestinationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
