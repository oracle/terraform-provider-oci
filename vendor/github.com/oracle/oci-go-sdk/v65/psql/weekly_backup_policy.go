// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WeeklyBackupPolicy Weekly backup policy.
type WeeklyBackupPolicy struct {

	// Hour of the day when the backup starts.
	BackupStart *string `mandatory:"true" json:"backupStart"`

	// How many days the data should be stored after the database system deletion.
	RetentionDays *int `mandatory:"false" json:"retentionDays"`

	// The day of the week that the backup starts.
	DaysOfTheWeek []WeeklyBackupPolicyDaysOfTheWeekEnum `mandatory:"true" json:"daysOfTheWeek"`
}

// GetRetentionDays returns RetentionDays
func (m WeeklyBackupPolicy) GetRetentionDays() *int {
	return m.RetentionDays
}

func (m WeeklyBackupPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WeeklyBackupPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.DaysOfTheWeek {
		if _, ok := GetMappingWeeklyBackupPolicyDaysOfTheWeekEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DaysOfTheWeek: %s. Supported values are: %s.", val, strings.Join(GetWeeklyBackupPolicyDaysOfTheWeekEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m WeeklyBackupPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeWeeklyBackupPolicy WeeklyBackupPolicy
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeWeeklyBackupPolicy
	}{
		"WEEKLY",
		(MarshalTypeWeeklyBackupPolicy)(m),
	}

	return json.Marshal(&s)
}

// WeeklyBackupPolicyDaysOfTheWeekEnum Enum with underlying type: string
type WeeklyBackupPolicyDaysOfTheWeekEnum string

// Set of constants representing the allowable values for WeeklyBackupPolicyDaysOfTheWeekEnum
const (
	WeeklyBackupPolicyDaysOfTheWeekSunday    WeeklyBackupPolicyDaysOfTheWeekEnum = "SUNDAY"
	WeeklyBackupPolicyDaysOfTheWeekMonday    WeeklyBackupPolicyDaysOfTheWeekEnum = "MONDAY"
	WeeklyBackupPolicyDaysOfTheWeekTuesday   WeeklyBackupPolicyDaysOfTheWeekEnum = "TUESDAY"
	WeeklyBackupPolicyDaysOfTheWeekWednesday WeeklyBackupPolicyDaysOfTheWeekEnum = "WEDNESDAY"
	WeeklyBackupPolicyDaysOfTheWeekThursday  WeeklyBackupPolicyDaysOfTheWeekEnum = "THURSDAY"
	WeeklyBackupPolicyDaysOfTheWeekFriday    WeeklyBackupPolicyDaysOfTheWeekEnum = "FRIDAY"
	WeeklyBackupPolicyDaysOfTheWeekSaturday  WeeklyBackupPolicyDaysOfTheWeekEnum = "SATURDAY"
)

var mappingWeeklyBackupPolicyDaysOfTheWeekEnum = map[string]WeeklyBackupPolicyDaysOfTheWeekEnum{
	"SUNDAY":    WeeklyBackupPolicyDaysOfTheWeekSunday,
	"MONDAY":    WeeklyBackupPolicyDaysOfTheWeekMonday,
	"TUESDAY":   WeeklyBackupPolicyDaysOfTheWeekTuesday,
	"WEDNESDAY": WeeklyBackupPolicyDaysOfTheWeekWednesday,
	"THURSDAY":  WeeklyBackupPolicyDaysOfTheWeekThursday,
	"FRIDAY":    WeeklyBackupPolicyDaysOfTheWeekFriday,
	"SATURDAY":  WeeklyBackupPolicyDaysOfTheWeekSaturday,
}

var mappingWeeklyBackupPolicyDaysOfTheWeekEnumLowerCase = map[string]WeeklyBackupPolicyDaysOfTheWeekEnum{
	"sunday":    WeeklyBackupPolicyDaysOfTheWeekSunday,
	"monday":    WeeklyBackupPolicyDaysOfTheWeekMonday,
	"tuesday":   WeeklyBackupPolicyDaysOfTheWeekTuesday,
	"wednesday": WeeklyBackupPolicyDaysOfTheWeekWednesday,
	"thursday":  WeeklyBackupPolicyDaysOfTheWeekThursday,
	"friday":    WeeklyBackupPolicyDaysOfTheWeekFriday,
	"saturday":  WeeklyBackupPolicyDaysOfTheWeekSaturday,
}

// GetWeeklyBackupPolicyDaysOfTheWeekEnumValues Enumerates the set of values for WeeklyBackupPolicyDaysOfTheWeekEnum
func GetWeeklyBackupPolicyDaysOfTheWeekEnumValues() []WeeklyBackupPolicyDaysOfTheWeekEnum {
	values := make([]WeeklyBackupPolicyDaysOfTheWeekEnum, 0)
	for _, v := range mappingWeeklyBackupPolicyDaysOfTheWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetWeeklyBackupPolicyDaysOfTheWeekEnumStringValues Enumerates the set of values in String for WeeklyBackupPolicyDaysOfTheWeekEnum
func GetWeeklyBackupPolicyDaysOfTheWeekEnumStringValues() []string {
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

// GetMappingWeeklyBackupPolicyDaysOfTheWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWeeklyBackupPolicyDaysOfTheWeekEnum(val string) (WeeklyBackupPolicyDaysOfTheWeekEnum, bool) {
	enum, ok := mappingWeeklyBackupPolicyDaysOfTheWeekEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
