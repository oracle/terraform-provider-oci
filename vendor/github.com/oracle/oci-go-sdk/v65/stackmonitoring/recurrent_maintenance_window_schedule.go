// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecurrentMaintenanceWindowSchedule Schedule information for the Maintenance Window that is executed multiple times.
type RecurrentMaintenanceWindowSchedule struct {

	// A RFC5545 formatted recurrence string which represents the Maintenance Window Recurrence.
	// Please refer this for details:https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10
	// FREQ: Frequency of the Maintenance Window. The supported values are: DAILY and WEEKLY.
	// BYDAY: Comma separated days for Weekly Maintenance Window.
	// BYHOUR: Specifies the start hour of each recurrence after `timeMaintenanceWindowStart` value.
	// BYMINUTE: Specifies the start minute of each reccurrence after `timeMaintenanceWindowStart` value. The default value is 00
	// BYSECOND: Specifies the start second of each reccurrence after `timeMaintenanceWindowStart` value. The default value is 00
	// Other Rules are not supported.
	MaintenanceWindowRecurrences *string `mandatory:"true" json:"maintenanceWindowRecurrences"`

	// Start time of Maintenance window. A RFC3339 formatted datetime string
	TimeMaintenanceWindowStart *common.SDKTime `mandatory:"false" json:"timeMaintenanceWindowStart"`

	// Start time of Maintenance window. A RFC3339 formatted datetime string
	TimeMaintenanceWindowEnd *common.SDKTime `mandatory:"false" json:"timeMaintenanceWindowEnd"`

	// Duration time of each recurrence of each Maintenance Window.
	// It must be specified as a string in ISO 8601 extended format.
	MaintenanceWindowDuration *string `mandatory:"false" json:"maintenanceWindowDuration"`
}

func (m RecurrentMaintenanceWindowSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecurrentMaintenanceWindowSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RecurrentMaintenanceWindowSchedule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRecurrentMaintenanceWindowSchedule RecurrentMaintenanceWindowSchedule
	s := struct {
		DiscriminatorParam string `json:"scheduleType"`
		MarshalTypeRecurrentMaintenanceWindowSchedule
	}{
		"RECURRENT",
		(MarshalTypeRecurrentMaintenanceWindowSchedule)(m),
	}

	return json.Marshal(&s)
}
