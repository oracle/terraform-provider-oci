// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedDbBackupConfig Backup Options
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type DistributedDbBackupConfig struct {

	// If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	IsAutoBackupEnabled *bool `mandatory:"false" json:"isAutoBackupEnabled"`

	// Number of days between the current and the earliest point of recoverability covered by automatic backups.
	// This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window.
	// When the value is updated, it is applied to all existing automatic backups.
	RecoveryWindowInDays *int `mandatory:"false" json:"recoveryWindowInDays"`

	// Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).
	// Example: `SLOT_TWO`
	AutoBackupWindow DistributedDbBackupConfigAutoBackupWindowEnum `mandatory:"false" json:"autoBackupWindow,omitempty"`

	// Time window selected for initiating full backup for the database system. There are twelve available two-hour time windows. If no option is selected, the value is null and a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).
	// Example: `SLOT_TWO`
	AutoFullBackupWindow DistributedDbBackupConfigAutoFullBackupWindowEnum `mandatory:"false" json:"autoFullBackupWindow,omitempty"`

	// Day of the week the full backup should be applied on the database system. If no option is selected, the value is null and we will default to Sunday.
	AutoFullBackupDay DistributedDbBackupConfigAutoFullBackupDayEnum `mandatory:"false" json:"autoFullBackupDay,omitempty"`

	// If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
	CanRunImmediateFullBackup *bool `mandatory:"false" json:"canRunImmediateFullBackup"`

	// If set to true, configures automatic incremental backups in the local region (the region of the DB system) and the remote region with a default frequency of 1 hour.
	// If you previously used RMAN or dbcli to configure backups, using the Console or the API for manged backups creates a new backup configuration for your database. The new configuration replaces the configuration created with RMAN or dbcli.
	// This means that you can no longer rely on your previously configured unmanaged backups to work.
	IsRemoteBackupEnabled *bool `mandatory:"false" json:"isRemoteBackupEnabled"`

	// The name of the remote region where the remote automatic incremental backups will be stored.
	// For information about valid region names, see
	// Regions and Availability Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm).
	RemoteRegion *string `mandatory:"false" json:"remoteRegion"`

	// Backup destination details.
	BackupDestinationDetails []DistributedDbBackupDestination `mandatory:"false" json:"backupDestinationDetails"`

	// This defines when the backups will be deleted. - IMMEDIATE option keep the backup for predefined time i.e 72 hours and then delete permanently... - RETAIN will keep the backups as per the policy defined for database backups.
	BackupDeletionPolicy DistributedDbBackupConfigBackupDeletionPolicyEnum `mandatory:"false" json:"backupDeletionPolicy,omitempty"`
}

func (m DistributedDbBackupConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDbBackupConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDistributedDbBackupConfigAutoBackupWindowEnum(string(m.AutoBackupWindow)); !ok && m.AutoBackupWindow != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoBackupWindow: %s. Supported values are: %s.", m.AutoBackupWindow, strings.Join(GetDistributedDbBackupConfigAutoBackupWindowEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDbBackupConfigAutoFullBackupWindowEnum(string(m.AutoFullBackupWindow)); !ok && m.AutoFullBackupWindow != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoFullBackupWindow: %s. Supported values are: %s.", m.AutoFullBackupWindow, strings.Join(GetDistributedDbBackupConfigAutoFullBackupWindowEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDbBackupConfigAutoFullBackupDayEnum(string(m.AutoFullBackupDay)); !ok && m.AutoFullBackupDay != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoFullBackupDay: %s. Supported values are: %s.", m.AutoFullBackupDay, strings.Join(GetDistributedDbBackupConfigAutoFullBackupDayEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDbBackupConfigBackupDeletionPolicyEnum(string(m.BackupDeletionPolicy)); !ok && m.BackupDeletionPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupDeletionPolicy: %s. Supported values are: %s.", m.BackupDeletionPolicy, strings.Join(GetDistributedDbBackupConfigBackupDeletionPolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDbBackupConfigAutoBackupWindowEnum Enum with underlying type: string
type DistributedDbBackupConfigAutoBackupWindowEnum string

// Set of constants representing the allowable values for DistributedDbBackupConfigAutoBackupWindowEnum
const (
	DistributedDbBackupConfigAutoBackupWindowOne    DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_ONE"
	DistributedDbBackupConfigAutoBackupWindowTwo    DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_TWO"
	DistributedDbBackupConfigAutoBackupWindowThree  DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_THREE"
	DistributedDbBackupConfigAutoBackupWindowFour   DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_FOUR"
	DistributedDbBackupConfigAutoBackupWindowFive   DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_FIVE"
	DistributedDbBackupConfigAutoBackupWindowSix    DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_SIX"
	DistributedDbBackupConfigAutoBackupWindowSeven  DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_SEVEN"
	DistributedDbBackupConfigAutoBackupWindowEight  DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_EIGHT"
	DistributedDbBackupConfigAutoBackupWindowNine   DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_NINE"
	DistributedDbBackupConfigAutoBackupWindowTen    DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_TEN"
	DistributedDbBackupConfigAutoBackupWindowEleven DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_ELEVEN"
	DistributedDbBackupConfigAutoBackupWindowTwelve DistributedDbBackupConfigAutoBackupWindowEnum = "SLOT_TWELVE"
)

var mappingDistributedDbBackupConfigAutoBackupWindowEnum = map[string]DistributedDbBackupConfigAutoBackupWindowEnum{
	"SLOT_ONE":    DistributedDbBackupConfigAutoBackupWindowOne,
	"SLOT_TWO":    DistributedDbBackupConfigAutoBackupWindowTwo,
	"SLOT_THREE":  DistributedDbBackupConfigAutoBackupWindowThree,
	"SLOT_FOUR":   DistributedDbBackupConfigAutoBackupWindowFour,
	"SLOT_FIVE":   DistributedDbBackupConfigAutoBackupWindowFive,
	"SLOT_SIX":    DistributedDbBackupConfigAutoBackupWindowSix,
	"SLOT_SEVEN":  DistributedDbBackupConfigAutoBackupWindowSeven,
	"SLOT_EIGHT":  DistributedDbBackupConfigAutoBackupWindowEight,
	"SLOT_NINE":   DistributedDbBackupConfigAutoBackupWindowNine,
	"SLOT_TEN":    DistributedDbBackupConfigAutoBackupWindowTen,
	"SLOT_ELEVEN": DistributedDbBackupConfigAutoBackupWindowEleven,
	"SLOT_TWELVE": DistributedDbBackupConfigAutoBackupWindowTwelve,
}

var mappingDistributedDbBackupConfigAutoBackupWindowEnumLowerCase = map[string]DistributedDbBackupConfigAutoBackupWindowEnum{
	"slot_one":    DistributedDbBackupConfigAutoBackupWindowOne,
	"slot_two":    DistributedDbBackupConfigAutoBackupWindowTwo,
	"slot_three":  DistributedDbBackupConfigAutoBackupWindowThree,
	"slot_four":   DistributedDbBackupConfigAutoBackupWindowFour,
	"slot_five":   DistributedDbBackupConfigAutoBackupWindowFive,
	"slot_six":    DistributedDbBackupConfigAutoBackupWindowSix,
	"slot_seven":  DistributedDbBackupConfigAutoBackupWindowSeven,
	"slot_eight":  DistributedDbBackupConfigAutoBackupWindowEight,
	"slot_nine":   DistributedDbBackupConfigAutoBackupWindowNine,
	"slot_ten":    DistributedDbBackupConfigAutoBackupWindowTen,
	"slot_eleven": DistributedDbBackupConfigAutoBackupWindowEleven,
	"slot_twelve": DistributedDbBackupConfigAutoBackupWindowTwelve,
}

// GetDistributedDbBackupConfigAutoBackupWindowEnumValues Enumerates the set of values for DistributedDbBackupConfigAutoBackupWindowEnum
func GetDistributedDbBackupConfigAutoBackupWindowEnumValues() []DistributedDbBackupConfigAutoBackupWindowEnum {
	values := make([]DistributedDbBackupConfigAutoBackupWindowEnum, 0)
	for _, v := range mappingDistributedDbBackupConfigAutoBackupWindowEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDbBackupConfigAutoBackupWindowEnumStringValues Enumerates the set of values in String for DistributedDbBackupConfigAutoBackupWindowEnum
func GetDistributedDbBackupConfigAutoBackupWindowEnumStringValues() []string {
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

// GetMappingDistributedDbBackupConfigAutoBackupWindowEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDbBackupConfigAutoBackupWindowEnum(val string) (DistributedDbBackupConfigAutoBackupWindowEnum, bool) {
	enum, ok := mappingDistributedDbBackupConfigAutoBackupWindowEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDbBackupConfigAutoFullBackupWindowEnum Enum with underlying type: string
type DistributedDbBackupConfigAutoFullBackupWindowEnum string

// Set of constants representing the allowable values for DistributedDbBackupConfigAutoFullBackupWindowEnum
const (
	DistributedDbBackupConfigAutoFullBackupWindowOne    DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_ONE"
	DistributedDbBackupConfigAutoFullBackupWindowTwo    DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_TWO"
	DistributedDbBackupConfigAutoFullBackupWindowThree  DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_THREE"
	DistributedDbBackupConfigAutoFullBackupWindowFour   DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_FOUR"
	DistributedDbBackupConfigAutoFullBackupWindowFive   DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_FIVE"
	DistributedDbBackupConfigAutoFullBackupWindowSix    DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_SIX"
	DistributedDbBackupConfigAutoFullBackupWindowSeven  DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_SEVEN"
	DistributedDbBackupConfigAutoFullBackupWindowEight  DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_EIGHT"
	DistributedDbBackupConfigAutoFullBackupWindowNine   DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_NINE"
	DistributedDbBackupConfigAutoFullBackupWindowTen    DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_TEN"
	DistributedDbBackupConfigAutoFullBackupWindowEleven DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_ELEVEN"
	DistributedDbBackupConfigAutoFullBackupWindowTwelve DistributedDbBackupConfigAutoFullBackupWindowEnum = "SLOT_TWELVE"
)

var mappingDistributedDbBackupConfigAutoFullBackupWindowEnum = map[string]DistributedDbBackupConfigAutoFullBackupWindowEnum{
	"SLOT_ONE":    DistributedDbBackupConfigAutoFullBackupWindowOne,
	"SLOT_TWO":    DistributedDbBackupConfigAutoFullBackupWindowTwo,
	"SLOT_THREE":  DistributedDbBackupConfigAutoFullBackupWindowThree,
	"SLOT_FOUR":   DistributedDbBackupConfigAutoFullBackupWindowFour,
	"SLOT_FIVE":   DistributedDbBackupConfigAutoFullBackupWindowFive,
	"SLOT_SIX":    DistributedDbBackupConfigAutoFullBackupWindowSix,
	"SLOT_SEVEN":  DistributedDbBackupConfigAutoFullBackupWindowSeven,
	"SLOT_EIGHT":  DistributedDbBackupConfigAutoFullBackupWindowEight,
	"SLOT_NINE":   DistributedDbBackupConfigAutoFullBackupWindowNine,
	"SLOT_TEN":    DistributedDbBackupConfigAutoFullBackupWindowTen,
	"SLOT_ELEVEN": DistributedDbBackupConfigAutoFullBackupWindowEleven,
	"SLOT_TWELVE": DistributedDbBackupConfigAutoFullBackupWindowTwelve,
}

var mappingDistributedDbBackupConfigAutoFullBackupWindowEnumLowerCase = map[string]DistributedDbBackupConfigAutoFullBackupWindowEnum{
	"slot_one":    DistributedDbBackupConfigAutoFullBackupWindowOne,
	"slot_two":    DistributedDbBackupConfigAutoFullBackupWindowTwo,
	"slot_three":  DistributedDbBackupConfigAutoFullBackupWindowThree,
	"slot_four":   DistributedDbBackupConfigAutoFullBackupWindowFour,
	"slot_five":   DistributedDbBackupConfigAutoFullBackupWindowFive,
	"slot_six":    DistributedDbBackupConfigAutoFullBackupWindowSix,
	"slot_seven":  DistributedDbBackupConfigAutoFullBackupWindowSeven,
	"slot_eight":  DistributedDbBackupConfigAutoFullBackupWindowEight,
	"slot_nine":   DistributedDbBackupConfigAutoFullBackupWindowNine,
	"slot_ten":    DistributedDbBackupConfigAutoFullBackupWindowTen,
	"slot_eleven": DistributedDbBackupConfigAutoFullBackupWindowEleven,
	"slot_twelve": DistributedDbBackupConfigAutoFullBackupWindowTwelve,
}

// GetDistributedDbBackupConfigAutoFullBackupWindowEnumValues Enumerates the set of values for DistributedDbBackupConfigAutoFullBackupWindowEnum
func GetDistributedDbBackupConfigAutoFullBackupWindowEnumValues() []DistributedDbBackupConfigAutoFullBackupWindowEnum {
	values := make([]DistributedDbBackupConfigAutoFullBackupWindowEnum, 0)
	for _, v := range mappingDistributedDbBackupConfigAutoFullBackupWindowEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDbBackupConfigAutoFullBackupWindowEnumStringValues Enumerates the set of values in String for DistributedDbBackupConfigAutoFullBackupWindowEnum
func GetDistributedDbBackupConfigAutoFullBackupWindowEnumStringValues() []string {
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

// GetMappingDistributedDbBackupConfigAutoFullBackupWindowEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDbBackupConfigAutoFullBackupWindowEnum(val string) (DistributedDbBackupConfigAutoFullBackupWindowEnum, bool) {
	enum, ok := mappingDistributedDbBackupConfigAutoFullBackupWindowEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDbBackupConfigAutoFullBackupDayEnum Enum with underlying type: string
type DistributedDbBackupConfigAutoFullBackupDayEnum string

// Set of constants representing the allowable values for DistributedDbBackupConfigAutoFullBackupDayEnum
const (
	DistributedDbBackupConfigAutoFullBackupDaySunday    DistributedDbBackupConfigAutoFullBackupDayEnum = "SUNDAY"
	DistributedDbBackupConfigAutoFullBackupDayMonday    DistributedDbBackupConfigAutoFullBackupDayEnum = "MONDAY"
	DistributedDbBackupConfigAutoFullBackupDayTuesday   DistributedDbBackupConfigAutoFullBackupDayEnum = "TUESDAY"
	DistributedDbBackupConfigAutoFullBackupDayWednesday DistributedDbBackupConfigAutoFullBackupDayEnum = "WEDNESDAY"
	DistributedDbBackupConfigAutoFullBackupDayThursday  DistributedDbBackupConfigAutoFullBackupDayEnum = "THURSDAY"
	DistributedDbBackupConfigAutoFullBackupDayFriday    DistributedDbBackupConfigAutoFullBackupDayEnum = "FRIDAY"
	DistributedDbBackupConfigAutoFullBackupDaySaturday  DistributedDbBackupConfigAutoFullBackupDayEnum = "SATURDAY"
)

var mappingDistributedDbBackupConfigAutoFullBackupDayEnum = map[string]DistributedDbBackupConfigAutoFullBackupDayEnum{
	"SUNDAY":    DistributedDbBackupConfigAutoFullBackupDaySunday,
	"MONDAY":    DistributedDbBackupConfigAutoFullBackupDayMonday,
	"TUESDAY":   DistributedDbBackupConfigAutoFullBackupDayTuesday,
	"WEDNESDAY": DistributedDbBackupConfigAutoFullBackupDayWednesday,
	"THURSDAY":  DistributedDbBackupConfigAutoFullBackupDayThursday,
	"FRIDAY":    DistributedDbBackupConfigAutoFullBackupDayFriday,
	"SATURDAY":  DistributedDbBackupConfigAutoFullBackupDaySaturday,
}

var mappingDistributedDbBackupConfigAutoFullBackupDayEnumLowerCase = map[string]DistributedDbBackupConfigAutoFullBackupDayEnum{
	"sunday":    DistributedDbBackupConfigAutoFullBackupDaySunday,
	"monday":    DistributedDbBackupConfigAutoFullBackupDayMonday,
	"tuesday":   DistributedDbBackupConfigAutoFullBackupDayTuesday,
	"wednesday": DistributedDbBackupConfigAutoFullBackupDayWednesday,
	"thursday":  DistributedDbBackupConfigAutoFullBackupDayThursday,
	"friday":    DistributedDbBackupConfigAutoFullBackupDayFriday,
	"saturday":  DistributedDbBackupConfigAutoFullBackupDaySaturday,
}

// GetDistributedDbBackupConfigAutoFullBackupDayEnumValues Enumerates the set of values for DistributedDbBackupConfigAutoFullBackupDayEnum
func GetDistributedDbBackupConfigAutoFullBackupDayEnumValues() []DistributedDbBackupConfigAutoFullBackupDayEnum {
	values := make([]DistributedDbBackupConfigAutoFullBackupDayEnum, 0)
	for _, v := range mappingDistributedDbBackupConfigAutoFullBackupDayEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDbBackupConfigAutoFullBackupDayEnumStringValues Enumerates the set of values in String for DistributedDbBackupConfigAutoFullBackupDayEnum
func GetDistributedDbBackupConfigAutoFullBackupDayEnumStringValues() []string {
	return []string{
		"SUNDAY",
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
	}
}

// GetMappingDistributedDbBackupConfigAutoFullBackupDayEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDbBackupConfigAutoFullBackupDayEnum(val string) (DistributedDbBackupConfigAutoFullBackupDayEnum, bool) {
	enum, ok := mappingDistributedDbBackupConfigAutoFullBackupDayEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DistributedDbBackupConfigBackupDeletionPolicyEnum Enum with underlying type: string
type DistributedDbBackupConfigBackupDeletionPolicyEnum string

// Set of constants representing the allowable values for DistributedDbBackupConfigBackupDeletionPolicyEnum
const (
	DistributedDbBackupConfigBackupDeletionPolicyImmediately          DistributedDbBackupConfigBackupDeletionPolicyEnum = "DELETE_IMMEDIATELY"
	DistributedDbBackupConfigBackupDeletionPolicyAfterRetentionPeriod DistributedDbBackupConfigBackupDeletionPolicyEnum = "DELETE_AFTER_RETENTION_PERIOD"
)

var mappingDistributedDbBackupConfigBackupDeletionPolicyEnum = map[string]DistributedDbBackupConfigBackupDeletionPolicyEnum{
	"DELETE_IMMEDIATELY":            DistributedDbBackupConfigBackupDeletionPolicyImmediately,
	"DELETE_AFTER_RETENTION_PERIOD": DistributedDbBackupConfigBackupDeletionPolicyAfterRetentionPeriod,
}

var mappingDistributedDbBackupConfigBackupDeletionPolicyEnumLowerCase = map[string]DistributedDbBackupConfigBackupDeletionPolicyEnum{
	"delete_immediately":            DistributedDbBackupConfigBackupDeletionPolicyImmediately,
	"delete_after_retention_period": DistributedDbBackupConfigBackupDeletionPolicyAfterRetentionPeriod,
}

// GetDistributedDbBackupConfigBackupDeletionPolicyEnumValues Enumerates the set of values for DistributedDbBackupConfigBackupDeletionPolicyEnum
func GetDistributedDbBackupConfigBackupDeletionPolicyEnumValues() []DistributedDbBackupConfigBackupDeletionPolicyEnum {
	values := make([]DistributedDbBackupConfigBackupDeletionPolicyEnum, 0)
	for _, v := range mappingDistributedDbBackupConfigBackupDeletionPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDbBackupConfigBackupDeletionPolicyEnumStringValues Enumerates the set of values in String for DistributedDbBackupConfigBackupDeletionPolicyEnum
func GetDistributedDbBackupConfigBackupDeletionPolicyEnumStringValues() []string {
	return []string{
		"DELETE_IMMEDIATELY",
		"DELETE_AFTER_RETENTION_PERIOD",
	}
}

// GetMappingDistributedDbBackupConfigBackupDeletionPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDbBackupConfigBackupDeletionPolicyEnum(val string) (DistributedDbBackupConfigBackupDeletionPolicyEnum, bool) {
	enum, ok := mappingDistributedDbBackupConfigBackupDeletionPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
