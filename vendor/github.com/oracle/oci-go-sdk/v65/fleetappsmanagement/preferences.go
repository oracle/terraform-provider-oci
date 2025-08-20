// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// Preferences Preferences to send notifications on the fleet activities.
type Preferences struct {
	UpcomingSchedule *UpcomingSchedule `mandatory:"false" json:"upcomingSchedule"`

	// Enables or disables notification on Job Failures.
	OnJobFailure *bool `mandatory:"false" json:"onJobFailure"`

	// Enables or disables notification on Environment Fleet Topology Modification.
	OnTopologyModification *bool `mandatory:"false" json:"onTopologyModification"`

	// Enables or disables notification when a task is paused.
	OnTaskPause *bool `mandatory:"false" json:"onTaskPause"`

	// Enables or disables notification on task failure.
	OnTaskFailure *bool `mandatory:"false" json:"onTaskFailure"`

	// Enables or disables notification on task success.
	OnTaskSuccess *bool `mandatory:"false" json:"onTaskSuccess"`

	// Enables or disables notification when fleet resource becomes non compliant.
	OnResourceNonCompliance *bool `mandatory:"false" json:"onResourceNonCompliance"`

	// Enables or disables notification when a newer version of runbook associated with a fleet is available
	OnRunbookNewerVersion *bool `mandatory:"false" json:"onRunbookNewerVersion"`
}

func (m Preferences) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Preferences) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
