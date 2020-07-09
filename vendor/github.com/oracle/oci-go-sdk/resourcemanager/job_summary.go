// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// JobSummary Returns a listing of all of the specified job's properties and their values.
type JobSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack that is associated with the specified job.
	StackId *string `mandatory:"false" json:"stackId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the stack of the associated job resides.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The job's display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of job executing
	Operation JobOperationEnum `mandatory:"false" json:"operation,omitempty"`

	JobOperationDetails JobOperationDetailsSummary `mandatory:"false" json:"jobOperationDetails"`

	ApplyJobPlanResolution *ApplyJobPlanResolution `mandatory:"false" json:"applyJobPlanResolution"`

	// Deprecated. Use the property `executionPlanJobId` in `jobOperationDetails` instead.
	// The plan job OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that was used (if this was an apply job and was not auto-approved).
	ResolvedPlanJobId *string `mandatory:"false" json:"resolvedPlanJobId"`

	// The date and time the job was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the job succeeded or failed.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Current state of the specified job.
	// For more information about job lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#JobStates).
	// Allowable values:
	// - ACCEPTED
	// - IN_PROGRESS
	// - FAILED
	// - SUCCEEDED
	// - CANCELING
	// - CANCELED
	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m JobSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *JobSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                     *string                           `json:"id"`
		StackId                *string                           `json:"stackId"`
		CompartmentId          *string                           `json:"compartmentId"`
		DisplayName            *string                           `json:"displayName"`
		Operation              JobOperationEnum                  `json:"operation"`
		JobOperationDetails    joboperationdetailssummary        `json:"jobOperationDetails"`
		ApplyJobPlanResolution *ApplyJobPlanResolution           `json:"applyJobPlanResolution"`
		ResolvedPlanJobId      *string                           `json:"resolvedPlanJobId"`
		TimeCreated            *common.SDKTime                   `json:"timeCreated"`
		TimeFinished           *common.SDKTime                   `json:"timeFinished"`
		LifecycleState         JobLifecycleStateEnum             `json:"lifecycleState"`
		FreeformTags           map[string]string                 `json:"freeformTags"`
		DefinedTags            map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.StackId = model.StackId

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Operation = model.Operation

	nn, e = model.JobOperationDetails.UnmarshalPolymorphicJSON(model.JobOperationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobOperationDetails = nn.(JobOperationDetailsSummary)
	} else {
		m.JobOperationDetails = nil
	}

	m.ApplyJobPlanResolution = model.ApplyJobPlanResolution

	m.ResolvedPlanJobId = model.ResolvedPlanJobId

	m.TimeCreated = model.TimeCreated

	m.TimeFinished = model.TimeFinished

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
