// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CronSchedule Cron schedule for a scheduled task.
type CronSchedule struct {

	// The date and time the scheduled task should execute first time after create or update;
	// thereafter the task will execute as specified in the schedule.
	TimeOfFirstExecution *common.SDKTime `mandatory:"false" json:"timeOfFirstExecution"`

	// Number of seconds to offset the query time window by to accommodate capture late arriving data. For example, a schedule run at 12:00 with a 10 minute interval and queryOffsetSecs=120 will use the query time window of 11:48-11:58 rather than 11:50-12:00 without queryOffsetSecs.
	QueryOffsetSecs *int `mandatory:"false" json:"queryOffsetSecs"`

	// End time for the schedule, even if the schedule would otherwise have remaining executions.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Value in cron format.
	Expression *string `mandatory:"false" json:"expression"`

	// Time zone, by default UTC.
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// Schedule misfire retry policy.
	MisfirePolicy ScheduleMisfirePolicyEnum `mandatory:"false" json:"misfirePolicy,omitempty"`
}

// GetMisfirePolicy returns MisfirePolicy
func (m CronSchedule) GetMisfirePolicy() ScheduleMisfirePolicyEnum {
	return m.MisfirePolicy
}

// GetTimeOfFirstExecution returns TimeOfFirstExecution
func (m CronSchedule) GetTimeOfFirstExecution() *common.SDKTime {
	return m.TimeOfFirstExecution
}

// GetQueryOffsetSecs returns QueryOffsetSecs
func (m CronSchedule) GetQueryOffsetSecs() *int {
	return m.QueryOffsetSecs
}

// GetTimeEnd returns TimeEnd
func (m CronSchedule) GetTimeEnd() *common.SDKTime {
	return m.TimeEnd
}

func (m CronSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CronSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingScheduleMisfirePolicyEnum(string(m.MisfirePolicy)); !ok && m.MisfirePolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MisfirePolicy: %s. Supported values are: %s.", m.MisfirePolicy, strings.Join(GetScheduleMisfirePolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CronSchedule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCronSchedule CronSchedule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCronSchedule
	}{
		"CRON",
		(MarshalTypeCronSchedule)(m),
	}

	return json.Marshal(&s)
}
