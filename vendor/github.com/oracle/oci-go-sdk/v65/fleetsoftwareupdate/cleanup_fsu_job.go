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

// CleanupFsuJob Cleanup Exadata Fleet Update Job resource.
type CleanupFsuJob struct {

	// OCID identifier for the Exadata Fleet Update Job.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier, this will map to the owner Exadata Fleet Update Action resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the Exadata Fleet Update Action that this job is part of.
	FsuActionId *string `mandatory:"true" json:"fsuActionId"`

	// The time the Exadata Fleet Update Job was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// OCID of the Exadata Fleet Update Collection that the job is executing on.
	FsuCollectionId *string `mandatory:"true" json:"fsuCollectionId"`

	// OCID of the Exadata Fleet Update Cycle that this job is part of.
	FsuCycleId *string `mandatory:"true" json:"fsuCycleId"`

	// Exadata Fleet Update Job display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	Progress *JobProgressDetails `mandatory:"false" json:"progress"`

	// The time the Exadata Fleet Update Job started execution. An RFC3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Exadata Fleet Update Job was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time the Exadata Fleet Update Job completed execution. An RFC3339 formatted datetime string.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
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

	// OCID of Target resource on which the job is executing the action.
	TargetId *string `mandatory:"false" json:"targetId"`

	Schedule ScheduleDetails `mandatory:"false" json:"schedule"`

	// The current state of the Exadata Fleet Update Job.
	LifecycleState JobLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m CleanupFsuJob) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m CleanupFsuJob) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CleanupFsuJob) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFsuActionId returns FsuActionId
func (m CleanupFsuJob) GetFsuActionId() *string {
	return m.FsuActionId
}

// GetProgress returns Progress
func (m CleanupFsuJob) GetProgress() *JobProgressDetails {
	return m.Progress
}

// GetTimeCreated returns TimeCreated
func (m CleanupFsuJob) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeStarted returns TimeStarted
func (m CleanupFsuJob) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeUpdated returns TimeUpdated
func (m CleanupFsuJob) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeFinished returns TimeFinished
func (m CleanupFsuJob) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleState returns LifecycleState
func (m CleanupFsuJob) GetLifecycleState() JobLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m CleanupFsuJob) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m CleanupFsuJob) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CleanupFsuJob) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m CleanupFsuJob) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m CleanupFsuJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CleanupFsuJob) ValidateEnumValue() (bool, error) {
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
func (m CleanupFsuJob) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCleanupFsuJob CleanupFsuJob
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCleanupFsuJob
	}{
		"CLEANUP",
		(MarshalTypeCleanupFsuJob)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CleanupFsuJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Progress         *JobProgressDetails               `json:"progress"`
		TimeStarted      *common.SDKTime                   `json:"timeStarted"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		TimeFinished     *common.SDKTime                   `json:"timeFinished"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		TargetId         *string                           `json:"targetId"`
		Schedule         scheduledetails                   `json:"schedule"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		FsuActionId      *string                           `json:"fsuActionId"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		LifecycleState   JobLifecycleStatesEnum            `json:"lifecycleState"`
		FsuCollectionId  *string                           `json:"fsuCollectionId"`
		FsuCycleId       *string                           `json:"fsuCycleId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Progress = model.Progress

	m.TimeStarted = model.TimeStarted

	m.TimeUpdated = model.TimeUpdated

	m.TimeFinished = model.TimeFinished

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

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

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.FsuActionId = model.FsuActionId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.FsuCollectionId = model.FsuCollectionId

	m.FsuCycleId = model.FsuCycleId

	return
}
