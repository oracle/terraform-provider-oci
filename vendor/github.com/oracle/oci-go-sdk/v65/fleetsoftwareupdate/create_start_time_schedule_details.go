// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateStartTimeScheduleDetails Start time details for the Exadata Fleet Update Action.
// The specified time should not conflict with existing Exadata Infrastructure maintenance windows.
// If Stage and Apply Actions are created with a timeToStart specified during Exadata Fleet Update Cycle
// creation, Apply should be scheduled at least 24 hours after the start time of the Stage Action.
type CreateStartTimeScheduleDetails struct {

	// The date and time the Exadata Fleet Update Action is expected to start.
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeToStart *common.SDKTime `mandatory:"true" json:"timeToStart"`
}

func (m CreateStartTimeScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateStartTimeScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateStartTimeScheduleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateStartTimeScheduleDetails CreateStartTimeScheduleDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateStartTimeScheduleDetails
	}{
		"START_TIME",
		(MarshalTypeCreateStartTimeScheduleDetails)(m),
	}

	return json.Marshal(&s)
}
