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

// FsuCycleSummary Exadata Fleet Update Cycle Summary.
type FsuCycleSummary struct {

	// OCID identifier for the Exadata Fleet Update Cycle.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of Exadata Fleet Update Cycle.
	Type CycleTypesEnum `mandatory:"true" json:"type"`

	// OCID identifier for the Collection ID the Exadata Fleet Update Cycle is assigned to.
	FsuCollectionId *string `mandatory:"true" json:"fsuCollectionId"`

	// Type of Collection this Exadata Fleet Update Cycle belongs to.
	CollectionType CollectionTypesEnum `mandatory:"true" json:"collectionType"`

	GoalVersionDetails FsuGoalVersionDetails `mandatory:"true" json:"goalVersionDetails"`

	// The date and time the Exadata Fleet Update Cycle was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Exadata Fleet Update Cycle.
	LifecycleState CycleLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Exadata Fleet Update Cycle display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// OCID identifier for the Action that is currently in execution, if applicable.
	ExecutingFsuActionId *string `mandatory:"false" json:"executingFsuActionId"`

	// In this array all the possible actions will be listed. The first element is the suggested Action.
	NextActionToExecute []NextActionToExecuteDetails `mandatory:"false" json:"nextActionToExecute"`

	// The latest Action type that was completed in the Exadata Fleet Update Cycle.
	// No value would indicate that the Cycle has not completed any Action yet.
	LastCompletedAction DetailedActionTypesEnum `mandatory:"false" json:"lastCompletedAction,omitempty"`

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
}

func (m FsuCycleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FsuCycleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCycleTypesEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCycleTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCollectionTypesEnum(string(m.CollectionType)); !ok && m.CollectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CollectionType: %s. Supported values are: %s.", m.CollectionType, strings.Join(GetCollectionTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCycleLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCycleLifecycleStatesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDetailedActionTypesEnum(string(m.LastCompletedAction)); !ok && m.LastCompletedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastCompletedAction: %s. Supported values are: %s.", m.LastCompletedAction, strings.Join(GetDetailedActionTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FsuCycleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                           `json:"displayName"`
		ExecutingFsuActionId  *string                           `json:"executingFsuActionId"`
		NextActionToExecute   []NextActionToExecuteDetails      `json:"nextActionToExecute"`
		LastCompletedAction   DetailedActionTypesEnum           `json:"lastCompletedAction"`
		DiagnosticsCollection *DiagnosticsCollectionDetails     `json:"diagnosticsCollection"`
		TimeUpdated           *common.SDKTime                   `json:"timeUpdated"`
		TimeFinished          *common.SDKTime                   `json:"timeFinished"`
		LifecycleDetails      *string                           `json:"lifecycleDetails"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		SystemTags            map[string]map[string]interface{} `json:"systemTags"`
		Id                    *string                           `json:"id"`
		CompartmentId         *string                           `json:"compartmentId"`
		Type                  CycleTypesEnum                    `json:"type"`
		FsuCollectionId       *string                           `json:"fsuCollectionId"`
		CollectionType        CollectionTypesEnum               `json:"collectionType"`
		GoalVersionDetails    fsugoalversiondetails             `json:"goalVersionDetails"`
		TimeCreated           *common.SDKTime                   `json:"timeCreated"`
		LifecycleState        CycleLifecycleStatesEnum          `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.ExecutingFsuActionId = model.ExecutingFsuActionId

	m.NextActionToExecute = make([]NextActionToExecuteDetails, len(model.NextActionToExecute))
	copy(m.NextActionToExecute, model.NextActionToExecute)
	m.LastCompletedAction = model.LastCompletedAction

	m.DiagnosticsCollection = model.DiagnosticsCollection

	m.TimeUpdated = model.TimeUpdated

	m.TimeFinished = model.TimeFinished

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.Type = model.Type

	m.FsuCollectionId = model.FsuCollectionId

	m.CollectionType = model.CollectionType

	nn, e = model.GoalVersionDetails.UnmarshalPolymorphicJSON(model.GoalVersionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.GoalVersionDetails = nn.(FsuGoalVersionDetails)
	} else {
		m.GoalVersionDetails = nil
	}

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
