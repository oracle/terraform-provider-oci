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

// BatchTaskSummary Summary information about a batch task.
type BatchTaskSummary interface {

	// The UUID of batch task.
	GetId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job to which this task belongs.
	GetJobId() *string

	// The name of the batch task. It must be unique within its parent batch job.
	GetName() *string

	// An optional description that provides additional context next to the displayName.
	GetDescription() *string

	// The current state of the batch task.
	GetLifecycleState() BatchTaskLifecycleStateEnum

	// A message that describes the current state of the batch task in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
	GetLifecycleDetails() *string
}

type batchtasksummary struct {
	JsonData         []byte
	Description      *string                     `mandatory:"false" json:"description"`
	LifecycleState   BatchTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails *string                     `mandatory:"false" json:"lifecycleDetails"`
	Id               *string                     `mandatory:"true" json:"id"`
	JobId            *string                     `mandatory:"true" json:"jobId"`
	Name             *string                     `mandatory:"true" json:"name"`
	Type             string                      `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *batchtasksummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbatchtasksummary batchtasksummary
	s := struct {
		Model Unmarshalerbatchtasksummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.JobId = s.Model.JobId
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *batchtasksummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "COMPUTE":
		mm := ComputeTaskSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BatchTaskSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m batchtasksummary) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m batchtasksummary) GetLifecycleState() BatchTaskLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m batchtasksummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m batchtasksummary) GetId() *string {
	return m.Id
}

// GetJobId returns JobId
func (m batchtasksummary) GetJobId() *string {
	return m.JobId
}

// GetName returns Name
func (m batchtasksummary) GetName() *string {
	return m.Name
}

func (m batchtasksummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m batchtasksummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBatchTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBatchTaskLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
