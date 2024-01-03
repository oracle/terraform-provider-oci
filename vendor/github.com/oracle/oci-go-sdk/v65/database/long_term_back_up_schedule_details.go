// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// LongTermBackUpScheduleDetails Details for the long-term backup schedule.
type LongTermBackUpScheduleDetails struct {

	// The frequency of the long-term backup schedule
	RepeatCadence LongTermBackUpScheduleDetailsRepeatCadenceEnum `mandatory:"false" json:"repeatCadence,omitempty"`

	// The timestamp for the long-term backup schedule. For a MONTHLY cadence, months having fewer days than the provided date will have the backup taken on the last day of that month.
	TimeOfBackup *common.SDKTime `mandatory:"false" json:"timeOfBackup"`

	// Retention period, in days, for long-term backups
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`

	// Indicates if the long-term backup schedule should be deleted. The default value is `FALSE`.
	IsDisabled *bool `mandatory:"false" json:"isDisabled"`
}

func (m LongTermBackUpScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LongTermBackUpScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLongTermBackUpScheduleDetailsRepeatCadenceEnum(string(m.RepeatCadence)); !ok && m.RepeatCadence != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RepeatCadence: %s. Supported values are: %s.", m.RepeatCadence, strings.Join(GetLongTermBackUpScheduleDetailsRepeatCadenceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LongTermBackUpScheduleDetailsRepeatCadenceEnum Enum with underlying type: string
type LongTermBackUpScheduleDetailsRepeatCadenceEnum string

// Set of constants representing the allowable values for LongTermBackUpScheduleDetailsRepeatCadenceEnum
const (
	LongTermBackUpScheduleDetailsRepeatCadenceOneTime LongTermBackUpScheduleDetailsRepeatCadenceEnum = "ONE_TIME"
	LongTermBackUpScheduleDetailsRepeatCadenceWeekly  LongTermBackUpScheduleDetailsRepeatCadenceEnum = "WEEKLY"
	LongTermBackUpScheduleDetailsRepeatCadenceMonthly LongTermBackUpScheduleDetailsRepeatCadenceEnum = "MONTHLY"
	LongTermBackUpScheduleDetailsRepeatCadenceYearly  LongTermBackUpScheduleDetailsRepeatCadenceEnum = "YEARLY"
)

var mappingLongTermBackUpScheduleDetailsRepeatCadenceEnum = map[string]LongTermBackUpScheduleDetailsRepeatCadenceEnum{
	"ONE_TIME": LongTermBackUpScheduleDetailsRepeatCadenceOneTime,
	"WEEKLY":   LongTermBackUpScheduleDetailsRepeatCadenceWeekly,
	"MONTHLY":  LongTermBackUpScheduleDetailsRepeatCadenceMonthly,
	"YEARLY":   LongTermBackUpScheduleDetailsRepeatCadenceYearly,
}

var mappingLongTermBackUpScheduleDetailsRepeatCadenceEnumLowerCase = map[string]LongTermBackUpScheduleDetailsRepeatCadenceEnum{
	"one_time": LongTermBackUpScheduleDetailsRepeatCadenceOneTime,
	"weekly":   LongTermBackUpScheduleDetailsRepeatCadenceWeekly,
	"monthly":  LongTermBackUpScheduleDetailsRepeatCadenceMonthly,
	"yearly":   LongTermBackUpScheduleDetailsRepeatCadenceYearly,
}

// GetLongTermBackUpScheduleDetailsRepeatCadenceEnumValues Enumerates the set of values for LongTermBackUpScheduleDetailsRepeatCadenceEnum
func GetLongTermBackUpScheduleDetailsRepeatCadenceEnumValues() []LongTermBackUpScheduleDetailsRepeatCadenceEnum {
	values := make([]LongTermBackUpScheduleDetailsRepeatCadenceEnum, 0)
	for _, v := range mappingLongTermBackUpScheduleDetailsRepeatCadenceEnum {
		values = append(values, v)
	}
	return values
}

// GetLongTermBackUpScheduleDetailsRepeatCadenceEnumStringValues Enumerates the set of values in String for LongTermBackUpScheduleDetailsRepeatCadenceEnum
func GetLongTermBackUpScheduleDetailsRepeatCadenceEnumStringValues() []string {
	return []string{
		"ONE_TIME",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
}

// GetMappingLongTermBackUpScheduleDetailsRepeatCadenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLongTermBackUpScheduleDetailsRepeatCadenceEnum(val string) (LongTermBackUpScheduleDetailsRepeatCadenceEnum, bool) {
	enum, ok := mappingLongTermBackUpScheduleDetailsRepeatCadenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
