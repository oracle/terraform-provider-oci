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

// ComputeTask compute task represents a single executable command together with its dependencies and resources.
type ComputeTask struct {

	// The UUID of batch task.
	Id *string `mandatory:"true" json:"id"`

	// The name of the batch task. It must be unique within its parent batch job.
	Name *string `mandatory:"true" json:"name"`

	// A list of resources (for example licences) this task needs for its execution.
	EntitlementClaims []string `mandatory:"true" json:"entitlementClaims"`

	// A list of tasks on which this tasks depends, referenced by name. Dependencies must be within the same parent (job or group task). For tasks within a group task, all dependencies must also be within that same group task.
	Dependencies []string `mandatory:"true" json:"dependencies"`

	// Environment variables to use for the task execution.
	EnvironmentVariables []EnvironmentVariable `mandatory:"true" json:"environmentVariables"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BatchTaskEnvironment.
	BatchTaskEnvironmentId *string `mandatory:"true" json:"batchTaskEnvironmentId"`

	// The hierarchical name of the task, which incorporates names of all parent group tasks, separated by "." (dot symbol). Maximum nesting depth is 4 levels. Example: groupTaskA.nestedGroupTaskB.thisTaskName
	HierarchicalName *string `mandatory:"false" json:"hierarchicalName"`

	// The hierarchical name of the group task. Null for top-level tasks.
	GroupTaskName *string `mandatory:"false" json:"groupTaskName"`

	// An optional description that provides additional context next to the displayName.
	Description *string `mandatory:"false" json:"description"`

	// A message that describes the current state of the batch task in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task profile used for this task.
	BatchTaskProfileId *string `mandatory:"false" json:"batchTaskProfileId"`

	// An executable command to start the processing of this task.
	Command []string `mandatory:"false" json:"command"`

	// Task arguments.
	Arguments []string `mandatory:"false" json:"arguments"`

	FleetAssignmentPolicy FleetAssignmentPolicy `mandatory:"false" json:"fleetAssignmentPolicy"`

	// List of up to 30 most recent execution history entries, from newest to oldest. Be aware that the maximum number of items returned may change in the future.
	MostRecentExecutionHistory []BatchTaskExecutionDetails `mandatory:"false" json:"mostRecentExecutionHistory"`

	// The current state of the batch task.
	LifecycleState BatchTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m ComputeTask) GetId() *string {
	return m.Id
}

// GetName returns Name
func (m ComputeTask) GetName() *string {
	return m.Name
}

// GetHierarchicalName returns HierarchicalName
func (m ComputeTask) GetHierarchicalName() *string {
	return m.HierarchicalName
}

// GetGroupTaskName returns GroupTaskName
func (m ComputeTask) GetGroupTaskName() *string {
	return m.GroupTaskName
}

// GetDescription returns Description
func (m ComputeTask) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m ComputeTask) GetLifecycleState() BatchTaskLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m ComputeTask) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetEntitlementClaims returns EntitlementClaims
func (m ComputeTask) GetEntitlementClaims() []string {
	return m.EntitlementClaims
}

// GetDependencies returns Dependencies
func (m ComputeTask) GetDependencies() []string {
	return m.Dependencies
}

// GetEnvironmentVariables returns EnvironmentVariables
func (m ComputeTask) GetEnvironmentVariables() []EnvironmentVariable {
	return m.EnvironmentVariables
}

func (m ComputeTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeTask) ValidateEnumValue() (bool, error) {
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
func (m ComputeTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeTask ComputeTask
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeComputeTask
	}{
		"COMPUTE",
		(MarshalTypeComputeTask)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ComputeTask) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		HierarchicalName           *string                     `json:"hierarchicalName"`
		GroupTaskName              *string                     `json:"groupTaskName"`
		Description                *string                     `json:"description"`
		LifecycleState             BatchTaskLifecycleStateEnum `json:"lifecycleState"`
		LifecycleDetails           *string                     `json:"lifecycleDetails"`
		BatchTaskProfileId         *string                     `json:"batchTaskProfileId"`
		Command                    []string                    `json:"command"`
		Arguments                  []string                    `json:"arguments"`
		FleetAssignmentPolicy      fleetassignmentpolicy       `json:"fleetAssignmentPolicy"`
		MostRecentExecutionHistory []batchtaskexecutiondetails `json:"mostRecentExecutionHistory"`
		Id                         *string                     `json:"id"`
		Name                       *string                     `json:"name"`
		EntitlementClaims          []string                    `json:"entitlementClaims"`
		Dependencies               []string                    `json:"dependencies"`
		EnvironmentVariables       []EnvironmentVariable       `json:"environmentVariables"`
		BatchTaskEnvironmentId     *string                     `json:"batchTaskEnvironmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.HierarchicalName = model.HierarchicalName

	m.GroupTaskName = model.GroupTaskName

	m.Description = model.Description

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

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

	m.MostRecentExecutionHistory = make([]BatchTaskExecutionDetails, len(model.MostRecentExecutionHistory))
	for i, n := range model.MostRecentExecutionHistory {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.MostRecentExecutionHistory[i] = nn.(BatchTaskExecutionDetails)
		} else {
			m.MostRecentExecutionHistory[i] = nil
		}
	}
	m.Id = model.Id

	m.Name = model.Name

	m.EntitlementClaims = make([]string, len(model.EntitlementClaims))
	copy(m.EntitlementClaims, model.EntitlementClaims)
	m.Dependencies = make([]string, len(model.Dependencies))
	copy(m.Dependencies, model.Dependencies)
	m.EnvironmentVariables = make([]EnvironmentVariable, len(model.EnvironmentVariables))
	copy(m.EnvironmentVariables, model.EnvironmentVariables)
	m.BatchTaskEnvironmentId = model.BatchTaskEnvironmentId

	return
}
