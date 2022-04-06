// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbBackupConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbBackupConfigAutoBackupWindowEnum(string(m.AutoBackupWindow)); !ok && m.AutoBackupWindow != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoBackupWindow: %s. Supported values are: %s.", m.AutoBackupWindow, strings.Join(GetDbBackupConfigAutoBackupWindowEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingDbBackupConfigAutoBackupWindowEnum = map[string]DbBackupConfigAutoBackupWindowEnum{
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

var mappingDbBackupConfigAutoBackupWindowEnumLowerCase = map[string]DbBackupConfigAutoBackupWindowEnum{
	"slot_one":    DbBackupConfigAutoBackupWindowOne,
	"slot_two":    DbBackupConfigAutoBackupWindowTwo,
	"slot_three":  DbBackupConfigAutoBackupWindowThree,
	"slot_four":   DbBackupConfigAutoBackupWindowFour,
	"slot_five":   DbBackupConfigAutoBackupWindowFive,
	"slot_six":    DbBackupConfigAutoBackupWindowSix,
	"slot_seven":  DbBackupConfigAutoBackupWindowSeven,
	"slot_eight":  DbBackupConfigAutoBackupWindowEight,
	"slot_nine":   DbBackupConfigAutoBackupWindowNine,
	"slot_ten":    DbBackupConfigAutoBackupWindowTen,
	"slot_eleven": DbBackupConfigAutoBackupWindowEleven,
	"slot_twelve": DbBackupConfigAutoBackupWindowTwelve,
}

// GetDbBackupConfigAutoBackupWindowEnumValues Enumerates the set of values for DbBackupConfigAutoBackupWindowEnum
func GetDbBackupConfigAutoBackupWindowEnumValues() []DbBackupConfigAutoBackupWindowEnum {
	values := make([]DbBackupConfigAutoBackupWindowEnum, 0)
	for _, v := range mappingDbBackupConfigAutoBackupWindowEnum {
		values = append(values, v)
	}
	return values
}

// GetDbBackupConfigAutoBackupWindowEnumStringValues Enumerates the set of values in String for DbBackupConfigAutoBackupWindowEnum
func GetDbBackupConfigAutoBackupWindowEnumStringValues() []string {
	return []string{
		"SLOT_ONE",
		"SLOT_TWO",
		"SLOT_THREE",
		"SLOT_FOUR",
		"SLOT_FIVE",
		"SLOT_SIX",
		"SLOT_SEVEN",
		"SLOT_EIGHT",
		"SLOT_NINE",
		"SLOT_TEN",
		"SLOT_ELEVEN",
		"SLOT_TWELVE",
	}
}

// GetMappingDbBackupConfigAutoBackupWindowEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbBackupConfigAutoBackupWindowEnum(val string) (DbBackupConfigAutoBackupWindowEnum, bool) {
	enum, ok := mappingDbBackupConfigAutoBackupWindowEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
