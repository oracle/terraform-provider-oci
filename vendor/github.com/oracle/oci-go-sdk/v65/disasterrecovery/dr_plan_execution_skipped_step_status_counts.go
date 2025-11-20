// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanExecutionSkippedStepStatusCounts A summary of steps that were skipped during a DR plan execution, including disabled, failed but ignored, timed out but ignored, and canceled steps.
type DrPlanExecutionSkippedStepStatusCounts struct {

	// The total number of steps that were skipped during a DR plan execution.
	TotalSkipped *int `mandatory:"true" json:"totalSkipped"`

	// The total number of disabled steps in a DR plan execution.
	Disabled *int `mandatory:"true" json:"disabled"`

	// The total number of steps that failed but were ignored during a DR plan execution.
	FailedIgnored *int `mandatory:"true" json:"failedIgnored"`

	// The total number of steps that timed out but were ignored during a DR plan execution.
	TimedOutIgnored *int `mandatory:"true" json:"timedOutIgnored"`

	// The total number of canceled steps in a DR plan execution.
	Canceled *int `mandatory:"true" json:"canceled"`
}

func (m DrPlanExecutionSkippedStepStatusCounts) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanExecutionSkippedStepStatusCounts) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
