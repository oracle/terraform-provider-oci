// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleICalTrigger The iCal schedule trigger.
type ScheduleICalTrigger struct {

	// This recurrence field conforms to RFC-5545 formatting
	Recurrence *string `mandatory:"true" json:"recurrence"`

	// The schedule starting date time, if null, System set the time when schedule is created.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The schedule end date time, if null, the schedule will never expire.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`
}

// GetTimeStart returns TimeStart
func (m ScheduleICalTrigger) GetTimeStart() *common.SDKTime {
	return m.TimeStart
}

// GetTimeEnd returns TimeEnd
func (m ScheduleICalTrigger) GetTimeEnd() *common.SDKTime {
	return m.TimeEnd
}

func (m ScheduleICalTrigger) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleICalTrigger) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScheduleICalTrigger) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScheduleICalTrigger ScheduleICalTrigger
	s := struct {
		DiscriminatorParam string `json:"triggerType"`
		MarshalTypeScheduleICalTrigger
	}{
		"ICAL",
		(MarshalTypeScheduleICalTrigger)(m),
	}

	return json.Marshal(&s)
}
