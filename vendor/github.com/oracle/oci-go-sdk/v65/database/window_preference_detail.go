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

// WindowPreferenceDetail The Single Scheduling Window details.
type WindowPreferenceDetail struct {

	// Weeks during the month when scheduled window should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week.
	// For example, to allow scheduling window during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Scheduling window cannot be scheduled for the fifth week of months that contain more than 28 days.
	// Note that this parameter works in conjunction with the  daysOfWeek and startTime parameters to allow you to specify specific days of the week and hours that scheduled window will be performed.
	WeeksOfMonth []int `mandatory:"true" json:"weeksOfMonth"`

	// Days during the week when scheduling window should be performed.
	DaysOfWeek []DayOfWeek `mandatory:"true" json:"daysOfWeek"`

	// The scheduling window start time. The value must use the ISO-8601 format "hh:mm".
	StartTime *string `mandatory:"true" json:"startTime"`

	// Duration window allows user to set a duration they plan to allocate for Scheduling window. The duration is in minutes.
	Duration *int `mandatory:"true" json:"duration"`

	// Indicates if duration the user plans to allocate for scheduling window is strictly enforced. The default value is `FALSE`.
	IsEnforcedDuration *bool `mandatory:"true" json:"isEnforcedDuration"`

	// Months during the year when scheduled window should be performed.
	Months []Month `mandatory:"false" json:"months"`
}

func (m WindowPreferenceDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WindowPreferenceDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
