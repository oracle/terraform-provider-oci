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

// PatchFsuCycle Patch Exadata Fleet Update Cycle resource details.
type PatchFsuCycle struct {

	// OCID identifier for the Exadata Fleet Update Cycle.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID identifier for the Collection ID the Exadata Fleet Update Cycle is assigned to.
	FsuCollectionId *string `mandatory:"true" json:"fsuCollectionId"`

	// The date and time the Exadata Fleet Update Cycle was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Exadata Fleet Update Cycle display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// OCID identifier for the Action that is currently in execution, if applicable.
	ExecutingFsuActionId *string `mandatory:"false" json:"executingFsuActionId"`

	// In this array all the possible actions will be listed. The first element is the suggested Action.
	NextActionToExecute []NextActionToExecuteDetails `mandatory:"false" json:"nextActionToExecute"`

	GoalVersionDetails FsuGoalVersionDetails `mandatory:"false" json:"goalVersionDetails"`

	BatchingStrategy BatchingStrategyDetails `mandatory:"false" json:"batchingStrategy"`

	StageActionSchedule ScheduleDetails `mandatory:"false" json:"stageActionSchedule"`

	ApplyActionSchedule ScheduleDetails `mandatory:"false" json:"applyActionSchedule"`

	DiagnosticsCollection *DiagnosticsCollectionDetails `mandatory:"false" json:"diagnosticsCollection"`

	// The date and time the Exadata Fleet Update Cycle was updated,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the Exadata Fleet Update Cycle was finished,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
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

	// Ignore all patches between the source and target homes during patching.
	IsIgnorePatches *bool `mandatory:"false" json:"isIgnorePatches"`

	// List of bug numbers to ignore.
	IsIgnoreMissingPatches []string `mandatory:"false" json:"isIgnoreMissingPatches"`

	// Service drain timeout specified in seconds.
	MaxDrainTimeoutInSeconds *int `mandatory:"false" json:"maxDrainTimeoutInSeconds"`

	// Ensure that services of administrator-managed Oracle RAC or Oracle RAC One databases are running on the same
	// instances before and after the move operation.
	IsKeepPlacement *bool `mandatory:"false" json:"isKeepPlacement"`

	// Type of Collection this Exadata Fleet Update Cycle belongs to.
	CollectionType CollectionTypesEnum `mandatory:"false" json:"collectionType,omitempty"`

	// The latest Action type that was completed in the Exadata Fleet Update Cycle.
	// No value would indicate that the Cycle has not completed any Action yet.
	LastCompletedAction DetailedActionTypesEnum `mandatory:"false" json:"lastCompletedAction,omitempty"`

	// The current state of the Exadata Fleet Update Cycle.
	LifecycleState CycleLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m PatchFsuCycle) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m PatchFsuCycle) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m PatchFsuCycle) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFsuCollectionId returns FsuCollectionId
func (m PatchFsuCycle) GetFsuCollectionId() *string {
	return m.FsuCollectionId
}

// GetCollectionType returns CollectionType
func (m PatchFsuCycle) GetCollectionType() CollectionTypesEnum {
	return m.CollectionType
}

// GetExecutingFsuActionId returns ExecutingFsuActionId
func (m PatchFsuCycle) GetExecutingFsuActionId() *string {
	return m.ExecutingFsuActionId
}

// GetNextActionToExecute returns NextActionToExecute
func (m PatchFsuCycle) GetNextActionToExecute() []NextActionToExecuteDetails {
	return m.NextActionToExecute
}

// GetLastCompletedAction returns LastCompletedAction
func (m PatchFsuCycle) GetLastCompletedAction() DetailedActionTypesEnum {
	return m.LastCompletedAction
}

// GetGoalVersionDetails returns GoalVersionDetails
func (m PatchFsuCycle) GetGoalVersionDetails() FsuGoalVersionDetails {
	return m.GoalVersionDetails
}

// GetBatchingStrategy returns BatchingStrategy
func (m PatchFsuCycle) GetBatchingStrategy() BatchingStrategyDetails {
	return m.BatchingStrategy
}

// GetStageActionSchedule returns StageActionSchedule
func (m PatchFsuCycle) GetStageActionSchedule() ScheduleDetails {
	return m.StageActionSchedule
}

// GetApplyActionSchedule returns ApplyActionSchedule
func (m PatchFsuCycle) GetApplyActionSchedule() ScheduleDetails {
	return m.ApplyActionSchedule
}

// GetDiagnosticsCollection returns DiagnosticsCollection
func (m PatchFsuCycle) GetDiagnosticsCollection() *DiagnosticsCollectionDetails {
	return m.DiagnosticsCollection
}

// GetTimeCreated returns TimeCreated
func (m PatchFsuCycle) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m PatchFsuCycle) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeFinished returns TimeFinished
func (m PatchFsuCycle) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleState returns LifecycleState
func (m PatchFsuCycle) GetLifecycleState() CycleLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m PatchFsuCycle) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m PatchFsuCycle) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m PatchFsuCycle) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m PatchFsuCycle) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m PatchFsuCycle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchFsuCycle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCollectionTypesEnum(string(m.CollectionType)); !ok && m.CollectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CollectionType: %s. Supported values are: %s.", m.CollectionType, strings.Join(GetCollectionTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDetailedActionTypesEnum(string(m.LastCompletedAction)); !ok && m.LastCompletedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastCompletedAction: %s. Supported values are: %s.", m.LastCompletedAction, strings.Join(GetDetailedActionTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCycleLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCycleLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatchFsuCycle) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatchFsuCycle PatchFsuCycle
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypePatchFsuCycle
	}{
		"PATCH",
		(MarshalTypePatchFsuCycle)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *PatchFsuCycle) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                           `json:"displayName"`
		CollectionType           CollectionTypesEnum               `json:"collectionType"`
		ExecutingFsuActionId     *string                           `json:"executingFsuActionId"`
		NextActionToExecute      []NextActionToExecuteDetails      `json:"nextActionToExecute"`
		LastCompletedAction      DetailedActionTypesEnum           `json:"lastCompletedAction"`
		GoalVersionDetails       fsugoalversiondetails             `json:"goalVersionDetails"`
		BatchingStrategy         batchingstrategydetails           `json:"batchingStrategy"`
		StageActionSchedule      scheduledetails                   `json:"stageActionSchedule"`
		ApplyActionSchedule      scheduledetails                   `json:"applyActionSchedule"`
		DiagnosticsCollection    *DiagnosticsCollectionDetails     `json:"diagnosticsCollection"`
		TimeUpdated              *common.SDKTime                   `json:"timeUpdated"`
		TimeFinished             *common.SDKTime                   `json:"timeFinished"`
		LifecycleDetails         *string                           `json:"lifecycleDetails"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
		SystemTags               map[string]map[string]interface{} `json:"systemTags"`
		IsIgnorePatches          *bool                             `json:"isIgnorePatches"`
		IsIgnoreMissingPatches   []string                          `json:"isIgnoreMissingPatches"`
		MaxDrainTimeoutInSeconds *int                              `json:"maxDrainTimeoutInSeconds"`
		IsKeepPlacement          *bool                             `json:"isKeepPlacement"`
		Id                       *string                           `json:"id"`
		CompartmentId            *string                           `json:"compartmentId"`
		FsuCollectionId          *string                           `json:"fsuCollectionId"`
		TimeCreated              *common.SDKTime                   `json:"timeCreated"`
		LifecycleState           CycleLifecycleStatesEnum          `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.CollectionType = model.CollectionType

	m.ExecutingFsuActionId = model.ExecutingFsuActionId

	m.NextActionToExecute = make([]NextActionToExecuteDetails, len(model.NextActionToExecute))
	copy(m.NextActionToExecute, model.NextActionToExecute)
	m.LastCompletedAction = model.LastCompletedAction

	nn, e = model.GoalVersionDetails.UnmarshalPolymorphicJSON(model.GoalVersionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.GoalVersionDetails = nn.(FsuGoalVersionDetails)
	} else {
		m.GoalVersionDetails = nil
	}

	nn, e = model.BatchingStrategy.UnmarshalPolymorphicJSON(model.BatchingStrategy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BatchingStrategy = nn.(BatchingStrategyDetails)
	} else {
		m.BatchingStrategy = nil
	}

	nn, e = model.StageActionSchedule.UnmarshalPolymorphicJSON(model.StageActionSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StageActionSchedule = nn.(ScheduleDetails)
	} else {
		m.StageActionSchedule = nil
	}

	nn, e = model.ApplyActionSchedule.UnmarshalPolymorphicJSON(model.ApplyActionSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ApplyActionSchedule = nn.(ScheduleDetails)
	} else {
		m.ApplyActionSchedule = nil
	}

	m.DiagnosticsCollection = model.DiagnosticsCollection

	m.TimeUpdated = model.TimeUpdated

	m.TimeFinished = model.TimeFinished

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.IsIgnorePatches = model.IsIgnorePatches

	m.IsIgnoreMissingPatches = make([]string, len(model.IsIgnoreMissingPatches))
	copy(m.IsIgnoreMissingPatches, model.IsIgnoreMissingPatches)
	m.MaxDrainTimeoutInSeconds = model.MaxDrainTimeoutInSeconds

	m.IsKeepPlacement = model.IsKeepPlacement

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.FsuCollectionId = model.FsuCollectionId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
