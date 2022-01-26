// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CronExecutionSchedule An autoscaling execution schedule that uses a cron expression.
type CronExecutionSchedule struct {

	// A cron expression that represents the time at which to execute the autoscaling policy.
	// Cron expressions have this format: `<second> <minute> <hour> <day of month> <month> <day of week> <year>`
	// You can use special characters that are supported with the Quartz cron implementation.
	// You must specify `0` as the value for seconds.
	// Example: `0 15 10 ? * *`
	Expression *string `mandatory:"true" json:"expression"`

	// The time zone for the execution schedule.
	Timezone ExecutionScheduleTimezoneEnum `mandatory:"true" json:"timezone"`
}

//GetTimezone returns Timezone
func (m CronExecutionSchedule) GetTimezone() ExecutionScheduleTimezoneEnum {
	return m.Timezone
}

func (m CronExecutionSchedule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CronExecutionSchedule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCronExecutionSchedule CronExecutionSchedule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCronExecutionSchedule
	}{
		"cron",
		(MarshalTypeCronExecutionSchedule)(m),
	}

	return json.Marshal(&s)
}
