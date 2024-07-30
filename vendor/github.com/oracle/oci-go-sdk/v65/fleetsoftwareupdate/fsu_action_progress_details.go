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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FsuActionProgressDetails Progress of the Action in execution. If the Exadata Fleet Update Action has not started yet, this will be omitted.
type FsuActionProgressDetails struct {

	// Number of targets with jobs in progress.
	InProgressTargets *int `mandatory:"false" json:"inProgressTargets"`

	// Number of targets with completed jobs.
	CompletedTargets *int `mandatory:"false" json:"completedTargets"`

	// Number of targets with failed jobs.
	FailedTargets *int `mandatory:"false" json:"failedTargets"`

	// Number of targets with jobs waiting for batch to execute or for user to resume.
	WaitingTargets *int `mandatory:"false" json:"waitingTargets"`

	// Total number of targets impacted by Exadata Fleet Update Action.
	TotalTargets *int `mandatory:"false" json:"totalTargets"`
}

func (m FsuActionProgressDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FsuActionProgressDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
