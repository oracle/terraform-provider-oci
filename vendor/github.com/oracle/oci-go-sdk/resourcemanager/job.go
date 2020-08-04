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

// Job The properties that define a job. Jobs perform the actions that are defined in your configuration.
// - **Plan job**. A plan job takes your Terraform configuration, parses it, and creates an execution plan.
// - **Apply job**. The apply job takes your execution plan, applies it to the associated stack, then executes
// the configuration's instructions.
// - **Destroy job**. To clean up the infrastructure controlled by the stack, you run a destroy job.
// A destroy job does not delete the stack or associated job resources,
// but instead releases the resources managed by the stack.
// - **Import_TF_State job**. An import Terraform state job takes a Terraform state file and sets it as the current
// state of the stack. This is used to migrate local Terraform environments to Resource Manager.
type Job struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack that is associated with the job.
	StackId *string `mandatory:"false" json:"stackId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the job's associated stack resides.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The job's display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of job executing.
	Operation JobOperationEnum `mandatory:"false" json:"operation,omitempty"`

	JobOperationDetails JobOperationDetails `mandatory:"false" json:"jobOperationDetails"`

	ApplyJobPlanResolution *ApplyJobPlanResolution `mandatory:"false" json:"applyJobPlanResolution"`

	// Deprecated. Use the property `executionPlanJobId` in `jobOperationDetails` instead.
	// The plan job OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that was used (if this was an apply job and was not auto-approved).
	ResolvedPlanJobId *string `mandatory:"false" json:"resolvedPlanJobId"`

	// The date and time when the job was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the job stopped running, irrespective of whether the job ran successfully.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Current state of the specified job.
	// For more information about job lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#JobStates).
	LifecycleState JobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	FailureDetails *FailureDetails `mandatory:"false" json:"failureDetails"`

	// File path to the directory from which Terraform runs.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// Terraform variables associated with this resource.
	// Maximum number of variables supported is 250.
	// The maximum size of each variable, including both name and value, is 4096 bytes.
	// Example: `{"CompartmentId": "compartment-id-value"}`
	Variables map[string]string `mandatory:"false" json:"variables"`

	ConfigSource ConfigSourceRecord `mandatory:"false" json:"configSource"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Job) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Job) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id                     *string                           `json:"id"`
		StackId                *string                           `json:"stackId"`
		CompartmentId          *string                           `json:"compartmentId"`
		DisplayName            *string                           `json:"displayName"`
		Operation              JobOperationEnum                  `json:"operation"`
		JobOperationDetails    joboperationdetails               `json:"jobOperationDetails"`
		ApplyJobPlanResolution *ApplyJobPlanResolution           `json:"applyJobPlanResolution"`
		ResolvedPlanJobId      *string                           `json:"resolvedPlanJobId"`
		TimeCreated            *common.SDKTime                   `json:"timeCreated"`
		TimeFinished           *common.SDKTime                   `json:"timeFinished"`
		LifecycleState         JobLifecycleStateEnum             `json:"lifecycleState"`
		FailureDetails         *FailureDetails                   `json:"failureDetails"`
		WorkingDirectory       *string                           `json:"workingDirectory"`
		Variables              map[string]string                 `json:"variables"`
		ConfigSource           configsourcerecord                `json:"configSource"`
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
		m.JobOperationDetails = nn.(JobOperationDetails)
	} else {
		m.JobOperationDetails = nil
	}

	m.ApplyJobPlanResolution = model.ApplyJobPlanResolution

	m.ResolvedPlanJobId = model.ResolvedPlanJobId

	m.TimeCreated = model.TimeCreated

	m.TimeFinished = model.TimeFinished

	m.LifecycleState = model.LifecycleState

	m.FailureDetails = model.FailureDetails

	m.WorkingDirectory = model.WorkingDirectory

	m.Variables = model.Variables

	nn, e = model.ConfigSource.UnmarshalPolymorphicJSON(model.ConfigSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConfigSource = nn.(ConfigSourceRecord)
	} else {
		m.ConfigSource = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// JobOperationEnum Enum with underlying type: string
type JobOperationEnum string

// Set of constants representing the allowable values for JobOperationEnum
const (
	JobOperationPlan          JobOperationEnum = "PLAN"
	JobOperationApply         JobOperationEnum = "APPLY"
	JobOperationDestroy       JobOperationEnum = "DESTROY"
	JobOperationImportTfState JobOperationEnum = "IMPORT_TF_STATE"
)

var mappingJobOperation = map[string]JobOperationEnum{
	"PLAN":            JobOperationPlan,
	"APPLY":           JobOperationApply,
	"DESTROY":         JobOperationDestroy,
	"IMPORT_TF_STATE": JobOperationImportTfState,
}

// GetJobOperationEnumValues Enumerates the set of values for JobOperationEnum
func GetJobOperationEnumValues() []JobOperationEnum {
	values := make([]JobOperationEnum, 0)
	for _, v := range mappingJobOperation {
		values = append(values, v)
	}
	return values
}

// JobLifecycleStateEnum Enum with underlying type: string
type JobLifecycleStateEnum string

// Set of constants representing the allowable values for JobLifecycleStateEnum
const (
	JobLifecycleStateAccepted   JobLifecycleStateEnum = "ACCEPTED"
	JobLifecycleStateInProgress JobLifecycleStateEnum = "IN_PROGRESS"
	JobLifecycleStateFailed     JobLifecycleStateEnum = "FAILED"
	JobLifecycleStateSucceeded  JobLifecycleStateEnum = "SUCCEEDED"
	JobLifecycleStateCanceling  JobLifecycleStateEnum = "CANCELING"
	JobLifecycleStateCanceled   JobLifecycleStateEnum = "CANCELED"
)

var mappingJobLifecycleState = map[string]JobLifecycleStateEnum{
	"ACCEPTED":    JobLifecycleStateAccepted,
	"IN_PROGRESS": JobLifecycleStateInProgress,
	"FAILED":      JobLifecycleStateFailed,
	"SUCCEEDED":   JobLifecycleStateSucceeded,
	"CANCELING":   JobLifecycleStateCanceling,
	"CANCELED":    JobLifecycleStateCanceled,
}

// GetJobLifecycleStateEnumValues Enumerates the set of values for JobLifecycleStateEnum
func GetJobLifecycleStateEnumValues() []JobLifecycleStateEnum {
	values := make([]JobLifecycleStateEnum, 0)
	for _, v := range mappingJobLifecycleState {
		values = append(values, v)
	}
	return values
}
