// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DbBackupConfig Backup Options
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type DbBackupConfig struct {

	// If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	AutoBackupEnabled *bool `mandatory:"false" json:"autoBackupEnabled"`

	// Number of days between the current and the earliest point of recoverability covered by automatic backups.
	// This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window.
	// When the value is updated, it is applied to all existing automatic backups.
	RecoveryWindowInDays *int `mandatory:"false" json:"recoveryWindowInDays"`

	// Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).
	// Example: `SLOT_TWO`
	AutoBackupWindow DbBackupConfigAutoBackupWindowEnum `mandatory:"false" json:"autoBackupWindow,omitempty"`

	// Backup destination details.
	BackupDestinationDetails []BackupDestinationDetails `mandatory:"false" json:"backupDestinationDetails"`
}

func (m DbBackupConfig) String() string {
	return common.PointerString(m)
}

// DbBackupConfigAutoBackupWindowEnum Enum with underlying type: string
type DbBackupConfigAutoBackupWindowEnum string

// Set of constants representing the allowable values for DbBackupConfigAutoBackupWindowEnum
const (
	DbBackupConfigAutoBackupWindowOne    DbBackupConfigAutoBackupWindowEnum = "SLOT_ONE"
	DbBackupConfigAutoBackupWindowTwo    DbBackupConfigAutoBackupWindowEnum = "SLOT_TWO"
	DbBackupConfigAutoBackupWindowThree  DbBackupConfigAutoBackupWindowEnum = "SLOT_THREE"
	DbBackupConfigAutoBackupWindowFour   DbBackupConfigAutoBackupWindowEnum = "SLOT_FOUR"
	DbBackupConfigAutoBackupWindowFive   DbBackupConfigAutoBackupWindowEnum = "SLOT_FIVE"
	DbBackupConfigAutoBackupWindowSix    DbBackupConfigAutoBackupWindowEnum = "SLOT_SIX"
	DbBackupConfigAutoBackupWindowSeven  DbBackupConfigAutoBackupWindowEnum = "SLOT_SEVEN"
	DbBackupConfigAutoBackupWindowEight  DbBackupConfigAutoBackupWindowEnum = "SLOT_EIGHT"
	DbBackupConfigAutoBackupWindowNine   DbBackupConfigAutoBackupWindowEnum = "SLOT_NINE"
	DbBackupConfigAutoBackupWindowTen    DbBackupConfigAutoBackupWindowEnum = "SLOT_TEN"
	DbBackupConfigAutoBackupWindowEleven DbBackupConfigAutoBackupWindowEnum = "SLOT_ELEVEN"
	DbBackupConfigAutoBackupWindowTwelve DbBackupConfigAutoBackupWindowEnum = "SLOT_TWELVE"
)

var mappingDbBackupConfigAutoBackupWindow = map[string]DbBackupConfigAutoBackupWindowEnum{
	"SLOT_ONE":    DbBackupConfigAutoBackupWindowOne,
	"SLOT_TWO":    DbBackupConfigAutoBackupWindowTwo,
	"SLOT_THREE":  DbBackupConfigAutoBackupWindowThree,
	"SLOT_FOUR":   DbBackupConfigAutoBackupWindowFour,
	"SLOT_FIVE":   DbBackupConfigAutoBackupWindowFive,
	"SLOT_SIX":    DbBackupConfigAutoBackupWindowSix,
	"SLOT_SEVEN":  DbBackupConfigAutoBackupWindowSeven,
	"SLOT_EIGHT":  DbBackupConfigAutoBackupWindowEight,
	"SLOT_NINE":   DbBackupConfigAutoBackupWindowNine,
	"SLOT_TEN":    DbBackupConfigAutoBackupWindowTen,
	"SLOT_ELEVEN": DbBackupConfigAutoBackupWindowEleven,
	"SLOT_TWELVE": DbBackupConfigAutoBackupWindowTwelve,
}

// GetDbBackupConfigAutoBackupWindowEnumValues Enumerates the set of values for DbBackupConfigAutoBackupWindowEnum
func GetDbBackupConfigAutoBackupWindowEnumValues() []DbBackupConfigAutoBackupWindowEnum {
	values := make([]DbBackupConfigAutoBackupWindowEnum, 0)
	for _, v := range mappingDbBackupConfigAutoBackupWindow {
		values = append(values, v)
	}
	return values
}
