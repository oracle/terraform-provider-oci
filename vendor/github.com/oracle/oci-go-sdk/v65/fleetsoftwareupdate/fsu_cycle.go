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

// FsuCycle Exadata Fleet Update Cycle resource details.
type FsuCycle interface {

	// OCID identifier for the Exadata Fleet Update Cycle.
	GetId() *string

	// Compartment Identifier.
	GetCompartmentId() *string

	// OCID identifier for the Collection ID the Exadata Fleet Update Cycle is assigned to.
	GetFsuCollectionId() *string

	// The date and time the Exadata Fleet Update Cycle was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeCreated() *common.SDKTime

	// The current state of the Exadata Fleet Update Cycle.
	GetLifecycleState() CycleLifecycleStatesEnum

	// Exadata Fleet Update Cycle display name.
	GetDisplayName() *string

	// Type of Collection this Exadata Fleet Update Cycle belongs to.
	GetCollectionType() CollectionTypesEnum

	// OCID identifier for the Action that is currently in execution, if applicable.
	GetExecutingFsuActionId() *string

	// In this array all the possible actions will be listed. The first element is the suggested Action.
	GetNextActionToExecute() []NextActionToExecuteDetails

	// The latest Action type that was completed in the Exadata Fleet Update Cycle.
	// No value would indicate that the Cycle has not completed any Action yet.
	GetLastCompletedAction() DetailedActionTypesEnum

	GetGoalVersionDetails() FsuGoalVersionDetails

	GetBatchingStrategy() BatchingStrategyDetails

	GetStageActionSchedule() ScheduleDetails

	GetApplyActionSchedule() ScheduleDetails

	GetDiagnosticsCollection() *DiagnosticsCollectionDetails

	// The date and time the Exadata Fleet Update Cycle was updated,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	GetTimeUpdated() *common.SDKTime

	// The date and time the Exadata Fleet Update Cycle was finished,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	GetTimeFinished() *common.SDKTime

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type fsucycle struct {
	JsonData              []byte
	DisplayName           *string                           `mandatory:"false" json:"displayName"`
	CollectionType        CollectionTypesEnum               `mandatory:"false" json:"collectionType,omitempty"`
	ExecutingFsuActionId  *string                           `mandatory:"false" json:"executingFsuActionId"`
	NextActionToExecute   []NextActionToExecuteDetails      `mandatory:"false" json:"nextActionToExecute"`
	LastCompletedAction   DetailedActionTypesEnum           `mandatory:"false" json:"lastCompletedAction,omitempty"`
	GoalVersionDetails    fsugoalversiondetails             `mandatory:"false" json:"goalVersionDetails"`
	BatchingStrategy      batchingstrategydetails           `mandatory:"false" json:"batchingStrategy"`
	StageActionSchedule   scheduledetails                   `mandatory:"false" json:"stageActionSchedule"`
	ApplyActionSchedule   scheduledetails                   `mandatory:"false" json:"applyActionSchedule"`
	DiagnosticsCollection *DiagnosticsCollectionDetails     `mandatory:"false" json:"diagnosticsCollection"`
	TimeUpdated           *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	TimeFinished          *common.SDKTime                   `mandatory:"false" json:"timeFinished"`
	LifecycleDetails      *string                           `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags          map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags           map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags            map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                    *string                           `mandatory:"true" json:"id"`
	CompartmentId         *string                           `mandatory:"true" json:"compartmentId"`
	FsuCollectionId       *string                           `mandatory:"true" json:"fsuCollectionId"`
	TimeCreated           *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState        CycleLifecycleStatesEnum          `mandatory:"true" json:"lifecycleState"`
	Type                  string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fsucycle) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfsucycle fsucycle
	s := struct {
		Model Unmarshalerfsucycle
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.FsuCollectionId = s.Model.FsuCollectionId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.CollectionType = s.Model.CollectionType
	m.ExecutingFsuActionId = s.Model.ExecutingFsuActionId
	m.NextActionToExecute = s.Model.NextActionToExecute
	m.LastCompletedAction = s.Model.LastCompletedAction
	m.GoalVersionDetails = s.Model.GoalVersionDetails
	m.BatchingStrategy = s.Model.BatchingStrategy
	m.StageActionSchedule = s.Model.StageActionSchedule
	m.ApplyActionSchedule = s.Model.ApplyActionSchedule
	m.DiagnosticsCollection = s.Model.DiagnosticsCollection
	m.TimeUpdated = s.Model.TimeUpdated
	m.TimeFinished = s.Model.TimeFinished
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fsucycle) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "PATCH":
		mm := PatchFsuCycle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FsuCycle: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m fsucycle) GetDisplayName() *string {
	return m.DisplayName
}

// GetCollectionType returns CollectionType
func (m fsucycle) GetCollectionType() CollectionTypesEnum {
	return m.CollectionType
}

// GetExecutingFsuActionId returns ExecutingFsuActionId
func (m fsucycle) GetExecutingFsuActionId() *string {
	return m.ExecutingFsuActionId
}

// GetNextActionToExecute returns NextActionToExecute
func (m fsucycle) GetNextActionToExecute() []NextActionToExecuteDetails {
	return m.NextActionToExecute
}

// GetLastCompletedAction returns LastCompletedAction
func (m fsucycle) GetLastCompletedAction() DetailedActionTypesEnum {
	return m.LastCompletedAction
}

// GetGoalVersionDetails returns GoalVersionDetails
func (m fsucycle) GetGoalVersionDetails() fsugoalversiondetails {
	return m.GoalVersionDetails
}

// GetBatchingStrategy returns BatchingStrategy
func (m fsucycle) GetBatchingStrategy() batchingstrategydetails {
	return m.BatchingStrategy
}

// GetStageActionSchedule returns StageActionSchedule
func (m fsucycle) GetStageActionSchedule() scheduledetails {
	return m.StageActionSchedule
}

// GetApplyActionSchedule returns ApplyActionSchedule
func (m fsucycle) GetApplyActionSchedule() scheduledetails {
	return m.ApplyActionSchedule
}

// GetDiagnosticsCollection returns DiagnosticsCollection
func (m fsucycle) GetDiagnosticsCollection() *DiagnosticsCollectionDetails {
	return m.DiagnosticsCollection
}

// GetTimeUpdated returns TimeUpdated
func (m fsucycle) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeFinished returns TimeFinished
func (m fsucycle) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleDetails returns LifecycleDetails
func (m fsucycle) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m fsucycle) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m fsucycle) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m fsucycle) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m fsucycle) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m fsucycle) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFsuCollectionId returns FsuCollectionId
func (m fsucycle) GetFsuCollectionId() *string {
	return m.FsuCollectionId
}

// GetTimeCreated returns TimeCreated
func (m fsucycle) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m fsucycle) GetLifecycleState() CycleLifecycleStatesEnum {
	return m.LifecycleState
}

func (m fsucycle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fsucycle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCycleLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCycleLifecycleStatesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCollectionTypesEnum(string(m.CollectionType)); !ok && m.CollectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CollectionType: %s. Supported values are: %s.", m.CollectionType, strings.Join(GetCollectionTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDetailedActionTypesEnum(string(m.LastCompletedAction)); !ok && m.LastCompletedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastCompletedAction: %s. Supported values are: %s.", m.LastCompletedAction, strings.Join(GetDetailedActionTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
