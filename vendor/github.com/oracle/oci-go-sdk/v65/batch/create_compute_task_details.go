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

// CreateComputeTaskDetails compute task represents a single executable command together with its dependencies and resources.
type CreateComputeTaskDetails struct {

	// The name of the batch task. It must be unique within its parent batch job.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BatchTaskEnvironment.
	BatchTaskEnvironmentId *string `mandatory:"true" json:"batchTaskEnvironmentId"`

	// An optional description that provides additional context next to the displayName.
	Description *string `mandatory:"false" json:"description"`

	// A list of resources (for example licences) this task needs for its execution.
	EntitlementClaims []string `mandatory:"false" json:"entitlementClaims"`

	// A list of tasks from the same job this task depends on referenced by name.
	Dependencies []string `mandatory:"false" json:"dependencies"`

	// Environment variables to use for the task execution.
	EnvironmentVariables []EnvironmentVariable `mandatory:"false" json:"environmentVariables"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task profile used for this task.
	BatchTaskProfileId *string `mandatory:"false" json:"batchTaskProfileId"`

	// An executable command to start the processing of this task.
	Command []string `mandatory:"false" json:"command"`

	// Task arguments.
	Arguments []string `mandatory:"false" json:"arguments"`

	FleetAssignmentPolicy FleetAssignmentPolicy `mandatory:"false" json:"fleetAssignmentPolicy"`
}

// GetName returns Name
func (m CreateComputeTaskDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateComputeTaskDetails) GetDescription() *string {
	return m.Description
}

// GetEntitlementClaims returns EntitlementClaims
func (m CreateComputeTaskDetails) GetEntitlementClaims() []string {
	return m.EntitlementClaims
}

// GetDependencies returns Dependencies
func (m CreateComputeTaskDetails) GetDependencies() []string {
	return m.Dependencies
}

// GetEnvironmentVariables returns EnvironmentVariables
func (m CreateComputeTaskDetails) GetEnvironmentVariables() []EnvironmentVariable {
	return m.EnvironmentVariables
}

func (m CreateComputeTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateComputeTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateComputeTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateComputeTaskDetails CreateComputeTaskDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateComputeTaskDetails
	}{
		"COMPUTE",
		(MarshalTypeCreateComputeTaskDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateComputeTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string               `json:"description"`
		EntitlementClaims      []string              `json:"entitlementClaims"`
		Dependencies           []string              `json:"dependencies"`
		EnvironmentVariables   []EnvironmentVariable `json:"environmentVariables"`
		BatchTaskProfileId     *string               `json:"batchTaskProfileId"`
		Command                []string              `json:"command"`
		Arguments              []string              `json:"arguments"`
		FleetAssignmentPolicy  fleetassignmentpolicy `json:"fleetAssignmentPolicy"`
		Name                   *string               `json:"name"`
		BatchTaskEnvironmentId *string               `json:"batchTaskEnvironmentId"`
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
	m.BatchTaskProfileId = model.BatchTaskProfileId

	m.Command = make([]string, len(model.Command))
	copy(m.Command, model.Command)
	m.Arguments = make([]string, len(model.Arguments))
	copy(m.Arguments, model.Arguments)
	nn, e = model.FleetAssignmentPolicy.UnmarshalPolymorphicJSON(model.FleetAssignmentPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FleetAssignmentPolicy = nn.(FleetAssignmentPolicy)
	} else {
		m.FleetAssignmentPolicy = nil
	}

	m.Name = model.Name

	m.BatchTaskEnvironmentId = model.BatchTaskEnvironmentId

	return
}
