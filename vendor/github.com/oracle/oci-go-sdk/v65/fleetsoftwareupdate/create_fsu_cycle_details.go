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

// CreateFsuCycleDetails Exadata Fleet Update Cycle resource creation details.
type CreateFsuCycleDetails interface {

	// Compartment Identifier.
	GetCompartmentId() *string

	// OCID identifier for the Collection ID the Exadata Fleet Update Cycle will be assigned to.
	GetFsuCollectionId() *string

	GetGoalVersionDetails() FsuGoalVersionDetails

	// Exadata Fleet Update Cycle display name.
	GetDisplayName() *string

	GetBatchingStrategy() CreateBatchingStrategyDetails

	GetStageActionSchedule() CreateScheduleDetails

	GetApplyActionSchedule() CreateScheduleDetails

	GetDiagnosticsCollection() *DiagnosticsCollectionDetails

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createfsucycledetails struct {
	JsonData              []byte
	DisplayName           *string                           `mandatory:"false" json:"displayName"`
	BatchingStrategy      createbatchingstrategydetails     `mandatory:"false" json:"batchingStrategy"`
	StageActionSchedule   createscheduledetails             `mandatory:"false" json:"stageActionSchedule"`
	ApplyActionSchedule   createscheduledetails             `mandatory:"false" json:"applyActionSchedule"`
	DiagnosticsCollection *DiagnosticsCollectionDetails     `mandatory:"false" json:"diagnosticsCollection"`
	FreeformTags          map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags           map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId         *string                           `mandatory:"true" json:"compartmentId"`
	FsuCollectionId       *string                           `mandatory:"true" json:"fsuCollectionId"`
	GoalVersionDetails    fsugoalversiondetails             `mandatory:"true" json:"goalVersionDetails"`
	Type                  string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createfsucycledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatefsucycledetails createfsucycledetails
	s := struct {
		Model Unmarshalercreatefsucycledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.FsuCollectionId = s.Model.FsuCollectionId
	m.GoalVersionDetails = s.Model.GoalVersionDetails
	m.DisplayName = s.Model.DisplayName
	m.BatchingStrategy = s.Model.BatchingStrategy
	m.StageActionSchedule = s.Model.StageActionSchedule
	m.ApplyActionSchedule = s.Model.ApplyActionSchedule
	m.DiagnosticsCollection = s.Model.DiagnosticsCollection
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createfsucycledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "PATCH":
		mm := CreatePatchFsuCycle{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateFsuCycleDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createfsucycledetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetBatchingStrategy returns BatchingStrategy
func (m createfsucycledetails) GetBatchingStrategy() createbatchingstrategydetails {
	return m.BatchingStrategy
}

// GetStageActionSchedule returns StageActionSchedule
func (m createfsucycledetails) GetStageActionSchedule() createscheduledetails {
	return m.StageActionSchedule
}

// GetApplyActionSchedule returns ApplyActionSchedule
func (m createfsucycledetails) GetApplyActionSchedule() createscheduledetails {
	return m.ApplyActionSchedule
}

// GetDiagnosticsCollection returns DiagnosticsCollection
func (m createfsucycledetails) GetDiagnosticsCollection() *DiagnosticsCollectionDetails {
	return m.DiagnosticsCollection
}

// GetFreeformTags returns FreeformTags
func (m createfsucycledetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createfsucycledetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m createfsucycledetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFsuCollectionId returns FsuCollectionId
func (m createfsucycledetails) GetFsuCollectionId() *string {
	return m.FsuCollectionId
}

// GetGoalVersionDetails returns GoalVersionDetails
func (m createfsucycledetails) GetGoalVersionDetails() fsugoalversiondetails {
	return m.GoalVersionDetails
}

func (m createfsucycledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createfsucycledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
