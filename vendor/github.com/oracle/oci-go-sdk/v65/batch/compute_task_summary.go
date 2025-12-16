// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ComputeTaskSummary Summary information about a compute task.
type ComputeTaskSummary struct {

	// The UUID of batch task.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job to which this task belongs.
	JobId *string `mandatory:"true" json:"jobId"`

	// The name of the batch task. It must be unique within its parent batch job.
	Name *string `mandatory:"true" json:"name"`

	// An optional description that provides additional context next to the displayName.
	Description *string `mandatory:"false" json:"description"`

	// A message that describes the current state of the batch task in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the batch task.
	LifecycleState BatchTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m ComputeTaskSummary) GetId() *string {
	return m.Id
}

// GetJobId returns JobId
func (m ComputeTaskSummary) GetJobId() *string {
	return m.JobId
}

// GetName returns Name
func (m ComputeTaskSummary) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m ComputeTaskSummary) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m ComputeTaskSummary) GetLifecycleState() BatchTaskLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m ComputeTaskSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m ComputeTaskSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeTaskSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBatchTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBatchTaskLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ComputeTaskSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeTaskSummary ComputeTaskSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeComputeTaskSummary
	}{
		"COMPUTE",
		(MarshalTypeComputeTaskSummary)(m),
	}

	return json.Marshal(&s)
}
