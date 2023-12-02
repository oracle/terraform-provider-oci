// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCccScheduleEvent A period where upgrades may be applied to Compute Cloud@Customer infrastructures
// associated with the schedule. All upgrade windows may not be used.
type CreateCccScheduleEvent struct {

	// A description of the Compute Cloud@Customer upgrade schedule time block.
	Description *string `mandatory:"true" json:"description"`

	// The date and time when the Compute Cloud@Customer upgrade schedule event starts,
	// inclusive. An RFC3339 formatted UTC datetime string. For an event with recurrences,
	// this is the date that a recurrence can start being applied.
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// The duration of this block of time. The duration must be specified and be of the
	// ISO-8601 format for durations.
	ScheduleEventDuration *string `mandatory:"true" json:"scheduleEventDuration"`

	// Frequency of recurrence of schedule block. When this field is not included, the event
	// is assumed to be a one time occurrence. The frequency field is strictly parsed and must
	// conform to RFC-5545 formatting for recurrences.
	ScheduleEventRecurrences *string `mandatory:"false" json:"scheduleEventRecurrences"`
}

func (m CreateCccScheduleEvent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCccScheduleEvent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
