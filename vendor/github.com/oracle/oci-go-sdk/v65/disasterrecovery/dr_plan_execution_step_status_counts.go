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

// DrPlanExecutionStepStatusCounts A categorized summary of step execution statuses and their corresponding counts.
type DrPlanExecutionStepStatusCounts struct {

	// The total number of steps in a DR plan execution.
	TotalSteps *int `mandatory:"true" json:"totalSteps"`

	RemainingSteps *DrPlanExecutionRemainingStepStatusCounts `mandatory:"true" json:"remainingSteps"`

	SkippedSteps *DrPlanExecutionSkippedStepStatusCounts `mandatory:"true" json:"skippedSteps"`

	SuccessfulSteps *DrPlanExecutionSuccessfulStepStatusCounts `mandatory:"true" json:"successfulSteps"`

	WarningSteps *DrPlanExecutionWarningStepStatusCounts `mandatory:"true" json:"warningSteps"`

	FailedSteps *DrPlanExecutionFailedStepStatusCounts `mandatory:"true" json:"failedSteps"`
}

func (m DrPlanExecutionStepStatusCounts) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanExecutionStepStatusCounts) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
