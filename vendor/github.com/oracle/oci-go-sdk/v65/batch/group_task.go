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

// GroupTask Group task is a construct that represents a container of tasks for execution.
type GroupTask struct {

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

	// The hierarchical name of the task, which incorporates names of all parent group tasks, separated by "." (dot symbol). Maximum nesting depth is 4 levels. Example: groupTaskA.nestedGroupTaskB.thisTaskName
	HierarchicalName *string `mandatory:"false" json:"hierarchicalName"`

	// The hierarchical name of the group task. Null for top-level tasks.
	GroupTaskName *string `mandatory:"false" json:"groupTaskName"`

	// An optional description that provides additional context next to the displayName.
	Description *string `mandatory:"false" json:"description"`

	// A message that describes the current state of the batch task in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The names of tasks contained directly (non-recursively) within this group task.
	Tasks []string `mandatory:"false" json:"tasks"`

	// The current state of the batch task.
	LifecycleState BatchTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m GroupTask) GetId() *string {
	return m.Id
}

// GetName returns Name
func (m GroupTask) GetName() *string {
	return m.Name
}

// GetHierarchicalName returns HierarchicalName
func (m GroupTask) GetHierarchicalName() *string {
	return m.HierarchicalName
}

// GetGroupTaskName returns GroupTaskName
func (m GroupTask) GetGroupTaskName() *string {
	return m.GroupTaskName
}

// GetDescription returns Description
func (m GroupTask) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m GroupTask) GetLifecycleState() BatchTaskLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m GroupTask) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetEntitlementClaims returns EntitlementClaims
func (m GroupTask) GetEntitlementClaims() []string {
	return m.EntitlementClaims
}

// GetDependencies returns Dependencies
func (m GroupTask) GetDependencies() []string {
	return m.Dependencies
}

// GetEnvironmentVariables returns EnvironmentVariables
func (m GroupTask) GetEnvironmentVariables() []EnvironmentVariable {
	return m.EnvironmentVariables
}

func (m GroupTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupTask) ValidateEnumValue() (bool, error) {
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
func (m GroupTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGroupTask GroupTask
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGroupTask
	}{
		"GROUP",
		(MarshalTypeGroupTask)(m),
	}

	return json.Marshal(&s)
}
