// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RollbackFsuJobSummary Summary of Rolback & Remove Targets Exadata Fleet Update Job resource.
type RollbackFsuJobSummary struct {

	// OCID identifier for the Exadata Fleet Update Job.
	Id *string `mandatory:"false" json:"id"`

	// Exadata Fleet Update Job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Compartment Identifier, this will map to the owner Exadata Fleet Update Action resource.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// OCID of the Exadata Fleet Update Action that this job is part of.
	FsuActionId *string `mandatory:"false" json:"fsuActionId"`

	Progress *JobProgress `mandatory:"false" json:"progress"`

	// The time the Exadata Fleet Update Job was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Exadata Fleet Update Job started execution. An RFC3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Exadata Fleet Update Job was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time the Exadata Fleet Update Job completed execution. An RFC3339 formatted datetime string.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// OCID of the Exadata Fleet Update Collection that the job is executing on.
	FsuCollectionId *string `mandatory:"false" json:"fsuCollectionId"`

	// OCID of the Exadata Fleet Update Cycle that this job is part of.
	FsuCycleId *string `mandatory:"false" json:"fsuCycleId"`

	// OCID of Target resource on which the job is executing the action.
	TargetId *string `mandatory:"false" json:"targetId"`

	Schedule ScheduleDetails `mandatory:"false" json:"schedule"`

	// The current state of the Job.
	LifecycleState JobLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m RollbackFsuJobSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m RollbackFsuJobSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m RollbackFsuJobSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFsuActionId returns FsuActionId
func (m RollbackFsuJobSummary) GetFsuActionId() *string {
	return m.FsuActionId
}

// GetProgress returns Progress
func (m RollbackFsuJobSummary) GetProgress() *JobProgress {
	return m.Progress
}

// GetTimeCreated returns TimeCreated
func (m RollbackFsuJobSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeStarted returns TimeStarted
func (m RollbackFsuJobSummary) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeUpdated returns TimeUpdated
func (m RollbackFsuJobSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeFinished returns TimeFinished
func (m RollbackFsuJobSummary) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleState returns LifecycleState
func (m RollbackFsuJobSummary) GetLifecycleState() JobLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m RollbackFsuJobSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m RollbackFsuJobSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m RollbackFsuJobSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m RollbackFsuJobSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m RollbackFsuJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RollbackFsuJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RollbackFsuJobSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRollbackFsuJobSummary RollbackFsuJobSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRollbackFsuJobSummary
	}{
		"ROLLBACK_AND_REMOVE_TARGET",
		(MarshalTypeRollbackFsuJobSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *RollbackFsuJobSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id               *string                           `json:"id"`
		DisplayName      *string                           `json:"displayName"`
		CompartmentId    *string                           `json:"compartmentId"`
		FsuActionId      *string                           `json:"fsuActionId"`
		Progress         *JobProgress                      `json:"progress"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeStarted      *common.SDKTime                   `json:"timeStarted"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		TimeFinished     *common.SDKTime                   `json:"timeFinished"`
		LifecycleState   JobLifecycleStatesEnum            `json:"lifecycleState"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		FsuCollectionId  *string                           `json:"fsuCollectionId"`
		FsuCycleId       *string                           `json:"fsuCycleId"`
		TargetId         *string                           `json:"targetId"`
		Schedule         scheduledetails                   `json:"schedule"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.FsuActionId = model.FsuActionId

	m.Progress = model.Progress

	m.TimeCreated = model.TimeCreated

	m.TimeStarted = model.TimeStarted

	m.TimeUpdated = model.TimeUpdated

	m.TimeFinished = model.TimeFinished

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.FsuCollectionId = model.FsuCollectionId

	m.FsuCycleId = model.FsuCycleId

	m.TargetId = model.TargetId

	nn, e = model.Schedule.UnmarshalPolymorphicJSON(model.Schedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schedule = nn.(ScheduleDetails)
	} else {
		m.Schedule = nil
	}

	return
}
