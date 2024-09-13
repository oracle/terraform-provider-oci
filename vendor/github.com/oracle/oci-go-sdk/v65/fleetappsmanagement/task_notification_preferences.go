// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TaskNotificationPreferences Preferences to send notifications on the task activities.
type TaskNotificationPreferences struct {

	// Enables notification on pause.
	ShouldNotifyOnPause *bool `mandatory:"false" json:"shouldNotifyOnPause"`

	// Enables or disables notification on Task Failures.
	ShouldNotifyOnTaskFailure *bool `mandatory:"false" json:"shouldNotifyOnTaskFailure"`

	// Enables or disables notification on Task Success.
	ShouldNotifyOnTaskSuccess *bool `mandatory:"false" json:"shouldNotifyOnTaskSuccess"`
}

func (m TaskNotificationPreferences) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskNotificationPreferences) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
