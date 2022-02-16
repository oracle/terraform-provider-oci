// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CronSchedule Cron schedule for a scheduled task.
type CronSchedule struct {

	// Value in cron format.
	Expression *string `mandatory:"true" json:"expression"`

	// Time zone, by default UTC.
	TimeZone *string `mandatory:"true" json:"timeZone"`

	// The date and time the scheduled task should execute first time after create or update;
	// thereafter the task will execute as specified in the schedule.
	TimeOfFirstExecution *common.SDKTime `mandatory:"false" json:"timeOfFirstExecution"`

	// Schedule misfire retry policy.
	MisfirePolicy ScheduleMisfirePolicyEnum `mandatory:"false" json:"misfirePolicy,omitempty"`
}

//GetMisfirePolicy returns MisfirePolicy
func (m CronSchedule) GetMisfirePolicy() ScheduleMisfirePolicyEnum {
	return m.MisfirePolicy
}

//GetTimeOfFirstExecution returns TimeOfFirstExecution
func (m CronSchedule) GetTimeOfFirstExecution() *common.SDKTime {
	return m.TimeOfFirstExecution
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
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
