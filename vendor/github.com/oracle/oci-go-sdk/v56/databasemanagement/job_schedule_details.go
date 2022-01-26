// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingJobScheduleDetailsIntervalType = map[string]JobScheduleDetailsIntervalTypeEnum{
	"DAILY":   JobScheduleDetailsIntervalTypeDaily,
	"HOURLY":  JobScheduleDetailsIntervalTypeHourly,
	"WEEKLY":  JobScheduleDetailsIntervalTypeWeekly,
	"MONTHLY": JobScheduleDetailsIntervalTypeMonthly,
	"NEVER":   JobScheduleDetailsIntervalTypeNever,
}

// GetJobScheduleDetailsIntervalTypeEnumValues Enumerates the set of values for JobScheduleDetailsIntervalTypeEnum
func GetJobScheduleDetailsIntervalTypeEnumValues() []JobScheduleDetailsIntervalTypeEnum {
	values := make([]JobScheduleDetailsIntervalTypeEnum, 0)
	for _, v := range mappingJobScheduleDetailsIntervalType {
		values = append(values, v)
	}
	return values
}
