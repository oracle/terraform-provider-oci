// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMaintenanceDetails The Maintenance Policy for the DB System or Read Replica that this model is included in.
type UpdateMaintenanceDetails struct {

	// The start of the 2 hour maintenance window.
	// This string is of the format: "{day-of-week} {time-of-day}".
	// "{day-of-week}" is a case-insensitive string like "mon", "tue", &c.
	// "{time-of-day}" is the "Time" portion of an RFC3339-formatted timestamp. Any second or sub-second time data will be truncated to zero.
	// If you set the read replica maintenance window to "", the read replica is set same as the DB system maintenance window. If not specific by the user,
	// there will be no changes to the maintenace window.
	WindowStartTime *string `mandatory:"false" json:"windowStartTime"`

	// The preferred version to target when performing an automatic MySQL upgrade. Defaults to OLDEST.
	// OLDEST: Choose the oldest available MySQL version based on the current version of the DB System.
	// SECOND_NEWEST: Choose the MySQL version before the newest for auto-upgrade.
	// NEWEST: Choose the latest and greatest MySQL version available for auto-upgrade.
	VersionPreference VersionPreferenceEnum `mandatory:"false" json:"versionPreference,omitempty"`

	// The preferred version track to target when performing an automatic MySQL upgrade. Defaults to FOLLOW.
	// LONG_TERM_SUPPORT: No MySQL database behavior changes.
	// INNOVATION:        Provides access to the latest features and all bug fixes.
	// FOLLOW:            Follows the track of the current MySQL version.
	VersionTrackPreference VersionTrackPreferenceEnum `mandatory:"false" json:"versionTrackPreference,omitempty"`

	// The maintenance schedule type of the DB system. Defaults to REGULAR.
	// EARLY:   Maintenance schedule follows a cycle where upgrades are performed when versions become deprecated.
	// REGULAR: Maintenance schedule follows the normal cycle where upgrades are performed when versions become unavailable.
	MaintenanceScheduleType MaintenanceScheduleTypeEnum `mandatory:"false" json:"maintenanceScheduleType,omitempty"`
}

func (m UpdateMaintenanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMaintenanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVersionPreferenceEnum(string(m.VersionPreference)); !ok && m.VersionPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VersionPreference: %s. Supported values are: %s.", m.VersionPreference, strings.Join(GetVersionPreferenceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVersionTrackPreferenceEnum(string(m.VersionTrackPreference)); !ok && m.VersionTrackPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VersionTrackPreference: %s. Supported values are: %s.", m.VersionTrackPreference, strings.Join(GetVersionTrackPreferenceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceScheduleTypeEnum(string(m.MaintenanceScheduleType)); !ok && m.MaintenanceScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceScheduleType: %s. Supported values are: %s.", m.MaintenanceScheduleType, strings.Join(GetMaintenanceScheduleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
