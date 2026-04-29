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

// GroupTaskSummary Summary information about a group task.
type GroupTaskSummary struct {

	// The UUID of batch task.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch job to which this task belongs.
	JobId *string `mandatory:"true" json:"jobId"`

	// The name of the batch task. It must be unique within its parent batch job.
	Name *string `mandatory:"true" json:"name"`

	// The hierarchical name of the task, which incorporates names of all parent group tasks, separated by "." (dot symbol). Maximum nesting depth is 4 levels. Example: groupTaskA.nestedGroupTaskB.thisTaskName
	HierarchicalName *string `mandatory:"false" json:"hierarchicalName"`

	// The hierarchical name of the group task. Null for top-level tasks.
	GroupTaskName *string `mandatory:"false" json:"groupTaskName"`

	// An optional description that provides additional context next to the displayName.
	Description *string `mandatory:"false" json:"description"`

	// A message that describes the current state of the batch task in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The count of tasks contained directly (non-recursively) within this group task.
	TaskCount *int `mandatory:"false" json:"taskCount"`

	// The current state of the batch task.
	LifecycleState BatchTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m GroupTaskSummary) GetId() *string {
	return m.Id
}

// GetJobId returns JobId
func (m GroupTaskSummary) GetJobId() *string {
	return m.JobId
}

// GetName returns Name
func (m GroupTaskSummary) GetName() *string {
	return m.Name
}

// GetHierarchicalName returns HierarchicalName
func (m GroupTaskSummary) GetHierarchicalName() *string {
	return m.HierarchicalName
}

// GetGroupTaskName returns GroupTaskName
func (m GroupTaskSummary) GetGroupTaskName() *string {
	return m.GroupTaskName
}

// GetDescription returns Description
func (m GroupTaskSummary) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m GroupTaskSummary) GetLifecycleState() BatchTaskLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m GroupTaskSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m GroupTaskSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupTaskSummary) ValidateEnumValue() (bool, error) {
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
func (m GroupTaskSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGroupTaskSummary GroupTaskSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGroupTaskSummary
	}{
		"GROUP",
		(MarshalTypeGroupTaskSummary)(m),
	}

	return json.Marshal(&s)
}
