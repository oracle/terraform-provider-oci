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

// CreateGroupTaskDetails Group task is a construct that represents a container of tasks for execution.
type CreateGroupTaskDetails struct {

	// The name of the batch task. It must be unique within its parent batch job.
	Name *string `mandatory:"true" json:"name"`

	// A list of tasks to be executed within this group task.
	Tasks []CreateBatchTaskDetails `mandatory:"true" json:"tasks"`

	// An optional description that provides additional context next to the displayName.
	Description *string `mandatory:"false" json:"description"`

	// A list of resources (for example licences) this task needs for its execution.
	EntitlementClaims []string `mandatory:"false" json:"entitlementClaims"`

	// A list of tasks on which this tasks depends, referenced by name. Dependencies must be within the same parent (job or group task). For tasks within a group task, all dependencies must also be within that same group task.
	Dependencies []string `mandatory:"false" json:"dependencies"`

	// Environment variables to use for the task execution.
	EnvironmentVariables []EnvironmentVariable `mandatory:"false" json:"environmentVariables"`
}

// GetName returns Name
func (m CreateGroupTaskDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateGroupTaskDetails) GetDescription() *string {
	return m.Description
}

// GetEntitlementClaims returns EntitlementClaims
func (m CreateGroupTaskDetails) GetEntitlementClaims() []string {
	return m.EntitlementClaims
}

// GetDependencies returns Dependencies
func (m CreateGroupTaskDetails) GetDependencies() []string {
	return m.Dependencies
}

// GetEnvironmentVariables returns EnvironmentVariables
func (m CreateGroupTaskDetails) GetEnvironmentVariables() []EnvironmentVariable {
	return m.EnvironmentVariables
}

func (m CreateGroupTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGroupTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateGroupTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateGroupTaskDetails CreateGroupTaskDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateGroupTaskDetails
	}{
		"GROUP",
		(MarshalTypeCreateGroupTaskDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateGroupTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description          *string                  `json:"description"`
		EntitlementClaims    []string                 `json:"entitlementClaims"`
		Dependencies         []string                 `json:"dependencies"`
		EnvironmentVariables []EnvironmentVariable    `json:"environmentVariables"`
		Name                 *string                  `json:"name"`
		Tasks                []createbatchtaskdetails `json:"tasks"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.EntitlementClaims = make([]string, len(model.EntitlementClaims))
	copy(m.EntitlementClaims, model.EntitlementClaims)
	m.Dependencies = make([]string, len(model.Dependencies))
	copy(m.Dependencies, model.Dependencies)
	m.EnvironmentVariables = make([]EnvironmentVariable, len(model.EnvironmentVariables))
	copy(m.EnvironmentVariables, model.EnvironmentVariables)
	m.Name = model.Name

	m.Tasks = make([]CreateBatchTaskDetails, len(model.Tasks))
	for i, n := range model.Tasks {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Tasks[i] = nn.(CreateBatchTaskDetails)
		} else {
			m.Tasks[i] = nil
		}
	}
	return
}
