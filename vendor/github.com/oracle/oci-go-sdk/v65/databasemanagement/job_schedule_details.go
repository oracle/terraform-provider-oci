// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JobScheduleDetails The details of the job schedule.
type JobScheduleDetails struct {

	// The start time of the scheduled job in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	StartTime *string `mandatory:"false" json:"startTime"`

	// The end time of the scheduled job in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	EndTime *string `mandatory:"false" json:"endTime"`

	// The interval type for a recurring scheduled job. For a non-recurring (one time) job, NEVER must be specified as the interval type.
	IntervalType JobScheduleDetailsIntervalTypeEnum `mandatory:"false" json:"intervalType,omitempty"`

	// The value for the interval period for a recurring scheduled job.
	IntervalValue *string `mandatory:"false" json:"intervalValue"`
}

func (m JobScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobScheduleDetailsIntervalTypeEnum(string(m.IntervalType)); !ok && m.IntervalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IntervalType: %s. Supported values are: %s.", m.IntervalType, strings.Join(GetJobScheduleDetailsIntervalTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobScheduleDetailsIntervalTypeEnum Enum with underlying type: string
type JobScheduleDetailsIntervalTypeEnum string

// Set of constants representing the allowable values for JobScheduleDetailsIntervalTypeEnum
const (
	JobScheduleDetailsIntervalTypeDaily   JobScheduleDetailsIntervalTypeEnum = "DAILY"
	JobScheduleDetailsIntervalTypeHourly  JobScheduleDetailsIntervalTypeEnum = "HOURLY"
	JobScheduleDetailsIntervalTypeWeekly  JobScheduleDetailsIntervalTypeEnum = "WEEKLY"
	JobScheduleDetailsIntervalTypeMonthly JobScheduleDetailsIntervalTypeEnum = "MONTHLY"
	JobScheduleDetailsIntervalTypeNever   JobScheduleDetailsIntervalTypeEnum = "NEVER"
)

var mappingJobScheduleDetailsIntervalTypeEnum = map[string]JobScheduleDetailsIntervalTypeEnum{
	"DAILY":   JobScheduleDetailsIntervalTypeDaily,
	"HOURLY":  JobScheduleDetailsIntervalTypeHourly,
	"WEEKLY":  JobScheduleDetailsIntervalTypeWeekly,
	"MONTHLY": JobScheduleDetailsIntervalTypeMonthly,
	"NEVER":   JobScheduleDetailsIntervalTypeNever,
}

var mappingJobScheduleDetailsIntervalTypeEnumLowerCase = map[string]JobScheduleDetailsIntervalTypeEnum{
	"daily":   JobScheduleDetailsIntervalTypeDaily,
	"hourly":  JobScheduleDetailsIntervalTypeHourly,
	"weekly":  JobScheduleDetailsIntervalTypeWeekly,
	"monthly": JobScheduleDetailsIntervalTypeMonthly,
	"never":   JobScheduleDetailsIntervalTypeNever,
}

// GetJobScheduleDetailsIntervalTypeEnumValues Enumerates the set of values for JobScheduleDetailsIntervalTypeEnum
func GetJobScheduleDetailsIntervalTypeEnumValues() []JobScheduleDetailsIntervalTypeEnum {
	values := make([]JobScheduleDetailsIntervalTypeEnum, 0)
	for _, v := range mappingJobScheduleDetailsIntervalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobScheduleDetailsIntervalTypeEnumStringValues Enumerates the set of values in String for JobScheduleDetailsIntervalTypeEnum
func GetJobScheduleDetailsIntervalTypeEnumStringValues() []string {
	return []string{
		"DAILY",
		"HOURLY",
		"WEEKLY",
		"MONTHLY",
		"NEVER",
	}
}

// GetMappingJobScheduleDetailsIntervalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobScheduleDetailsIntervalTypeEnum(val string) (JobScheduleDetailsIntervalTypeEnum, bool) {
	enum, ok := mappingJobScheduleDetailsIntervalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
