// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeTaskExecutionDetails Execution details for a compute task.
type ComputeTaskExecutionDetails struct {

	// A unique identifier for the task execution. Created as "taskId:taskVersion:runNumber".
	ExecutionId *string `mandatory:"false" json:"executionId"`

	// The date and time when the lifecycleState was changed to Waiting, and it is waiting for its dependencies to run, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeTransitionedToWaitingDependencies *common.SDKTime `mandatory:"false" json:"timeTransitionedToWaitingDependencies"`

	// The date and time when the lifecycleState was changed to Waiting, and it is queued to be executed, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeTransitionedToWaitingQueued *common.SDKTime `mandatory:"false" json:"timeTransitionedToWaitingQueued"`

	// The date and time when the lifecycleState was changed to In_Progress, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time when the lifecycleState changed to Succeeded, or Failed, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	FleetShape FleetShapeExecutionDetails `mandatory:"false" json:"fleetShape"`

	// List of error messages related to this task execution. Be aware that the maximum number of items returned may change in the future.
	Errors []string `mandatory:"false" json:"errors"`

	// The terminal lifecycle state of the task for this execution. Valid values are: SUCCEEDED, NEEDS_ATTENTION, CANCELED, or FAILED.
	CompletionLifecycleState BatchTaskLifecycleStateEnum `mandatory:"false" json:"completionLifecycleState,omitempty"`
}

func (m ComputeTaskExecutionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeTaskExecutionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBatchTaskLifecycleStateEnum(string(m.CompletionLifecycleState)); !ok && m.CompletionLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompletionLifecycleState: %s. Supported values are: %s.", m.CompletionLifecycleState, strings.Join(GetBatchTaskLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ComputeTaskExecutionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeTaskExecutionDetails ComputeTaskExecutionDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeComputeTaskExecutionDetails
	}{
		"COMPUTE_TASK_EXECUTION_DETAILS",
		(MarshalTypeComputeTaskExecutionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ComputeTaskExecutionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ExecutionId                           *string                     `json:"executionId"`
		TimeTransitionedToWaitingDependencies *common.SDKTime             `json:"timeTransitionedToWaitingDependencies"`
		TimeTransitionedToWaitingQueued       *common.SDKTime             `json:"timeTransitionedToWaitingQueued"`
		TimeStarted                           *common.SDKTime             `json:"timeStarted"`
		TimeCompleted                         *common.SDKTime             `json:"timeCompleted"`
		FleetShape                            fleetshapeexecutiondetails  `json:"fleetShape"`
		CompletionLifecycleState              BatchTaskLifecycleStateEnum `json:"completionLifecycleState"`
		Errors                                []string                    `json:"errors"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ExecutionId = model.ExecutionId

	m.TimeTransitionedToWaitingDependencies = model.TimeTransitionedToWaitingDependencies

	m.TimeTransitionedToWaitingQueued = model.TimeTransitionedToWaitingQueued

	m.TimeStarted = model.TimeStarted

	m.TimeCompleted = model.TimeCompleted

	nn, e = model.FleetShape.UnmarshalPolymorphicJSON(model.FleetShape.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FleetShape = nn.(FleetShapeExecutionDetails)
	} else {
		m.FleetShape = nil
	}

	m.CompletionLifecycleState = model.CompletionLifecycleState

	m.Errors = make([]string, len(model.Errors))
	copy(m.Errors, model.Errors)
	return
}
