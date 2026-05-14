// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalLocationScheduleOperation The scheduled multicloud external location update details for an Autonomous AI Database.
type ExternalLocationScheduleOperation struct {

	// The requested target multicloud placement value for a scheduled update. Use a CSP region for regional placement or a CSP physical zone for explicit AZ placement.
	TargetExternalLocation *string `mandatory:"true" json:"targetExternalLocation"`

	// Schedule to earliest available time slot.
	IsScheduleUpdateToEarliest *bool `mandatory:"false" json:"isScheduleUpdateToEarliest"`

	// Scheduled time (RFC3339).
	TimeScheduledUpdate *common.SDKTime `mandatory:"false" json:"timeScheduledUpdate"`

	// Cancel existing scheduled update.
	IsDisableUpdateSchedule *bool `mandatory:"false" json:"isDisableUpdateSchedule"`
}

// GetIsScheduleUpdateToEarliest returns IsScheduleUpdateToEarliest
func (m ExternalLocationScheduleOperation) GetIsScheduleUpdateToEarliest() *bool {
	return m.IsScheduleUpdateToEarliest
}

// GetTimeScheduledUpdate returns TimeScheduledUpdate
func (m ExternalLocationScheduleOperation) GetTimeScheduledUpdate() *common.SDKTime {
	return m.TimeScheduledUpdate
}

// GetIsDisableUpdateSchedule returns IsDisableUpdateSchedule
func (m ExternalLocationScheduleOperation) GetIsDisableUpdateSchedule() *bool {
	return m.IsDisableUpdateSchedule
}

func (m ExternalLocationScheduleOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalLocationScheduleOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalLocationScheduleOperation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalLocationScheduleOperation ExternalLocationScheduleOperation
	s := struct {
		DiscriminatorParam string `json:"operationType"`
		MarshalTypeExternalLocationScheduleOperation
	}{
		"EXTERNAL_LOCATION_UPDATE",
		(MarshalTypeExternalLocationScheduleOperation)(m),
	}

	return json.Marshal(&s)
}
