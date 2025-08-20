// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchOperationsSummary Patch operation list summary by status record.
type PatchOperationsSummary struct {

	// Total number of scheduled Patch operations.
	ScheduledPatchOpsCount *int `mandatory:"false" json:"scheduledPatchOpsCount"`

	// Total number of in progress Patch operations.
	RunningPatchOpsCount *int `mandatory:"false" json:"runningPatchOpsCount"`

	// Total number of successful Patch operations.
	SuccessfulPatchOpsCount *int `mandatory:"false" json:"successfulPatchOpsCount"`

	// Total number of Patch operations that have warnings.
	WarningsPatchOpsCount *int `mandatory:"false" json:"warningsPatchOpsCount"`

	// Total number of failed Patch operations.
	FailedPatchOpsCount *int `mandatory:"false" json:"failedPatchOpsCount"`
}

func (m PatchOperationsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchOperationsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
