// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateUpgradeFsuCycle Upgrade Exadata Fleet Update Cycle resource creation details. Batching strategy should always be NON_ROLLING for DB collection upgrade.
// For GI upgrade case, Batching strategy may be NON_ROLLING or SEQUENTIAL(default).
type CreateUpgradeFsuCycle struct {

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID identifier for the Collection ID the Exadata Fleet Update Cycle will be assigned to.
	FsuCollectionId *string `mandatory:"true" json:"fsuCollectionId"`

	GoalVersionDetails FsuGoalVersionDetails `mandatory:"true" json:"goalVersionDetails"`

	// Exadata Fleet Update Cycle display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	BatchingStrategy CreateBatchingStrategyDetails `mandatory:"false" json:"batchingStrategy"`

	StageActionSchedule CreateScheduleDetails `mandatory:"false" json:"stageActionSchedule"`

	ApplyActionSchedule CreateScheduleDetails `mandatory:"false" json:"applyActionSchedule"`

	DiagnosticsCollection *DiagnosticsCollectionDetails `mandatory:"false" json:"diagnosticsCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	UpgradeDetails UpgradeDetails `mandatory:"false" json:"upgradeDetails"`
}

// GetDisplayName returns DisplayName
func (m CreateUpgradeFsuCycle) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateUpgradeFsuCycle) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFsuCollectionId returns FsuCollectionId
func (m CreateUpgradeFsuCycle) GetFsuCollectionId() *string {
	return m.FsuCollectionId
}

// GetGoalVersionDetails returns GoalVersionDetails
func (m CreateUpgradeFsuCycle) GetGoalVersionDetails() FsuGoalVersionDetails {
	return m.GoalVersionDetails
}

// GetBatchingStrategy returns BatchingStrategy
func (m CreateUpgradeFsuCycle) GetBatchingStrategy() CreateBatchingStrategyDetails {
	return m.BatchingStrategy
}

// GetStageActionSchedule returns StageActionSchedule
func (m CreateUpgradeFsuCycle) GetStageActionSchedule() CreateScheduleDetails {
	return m.StageActionSchedule
}

// GetApplyActionSchedule returns ApplyActionSchedule
func (m CreateUpgradeFsuCycle) GetApplyActionSchedule() CreateScheduleDetails {
	return m.ApplyActionSchedule
}

// GetDiagnosticsCollection returns DiagnosticsCollection
func (m CreateUpgradeFsuCycle) GetDiagnosticsCollection() *DiagnosticsCollectionDetails {
	return m.DiagnosticsCollection
}

// GetFreeformTags returns FreeformTags
func (m CreateUpgradeFsuCycle) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateUpgradeFsuCycle) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateUpgradeFsuCycle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateUpgradeFsuCycle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateUpgradeFsuCycle) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateUpgradeFsuCycle CreateUpgradeFsuCycle
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateUpgradeFsuCycle
	}{
		"UPGRADE",
		(MarshalTypeCreateUpgradeFsuCycle)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateUpgradeFsuCycle) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                           `json:"displayName"`
		BatchingStrategy      createbatchingstrategydetails     `json:"batchingStrategy"`
		StageActionSchedule   createscheduledetails             `json:"stageActionSchedule"`
		ApplyActionSchedule   createscheduledetails             `json:"applyActionSchedule"`
		DiagnosticsCollection *DiagnosticsCollectionDetails     `json:"diagnosticsCollection"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		UpgradeDetails        upgradedetails                    `json:"upgradeDetails"`
		CompartmentId         *string                           `json:"compartmentId"`
		FsuCollectionId       *string                           `json:"fsuCollectionId"`
		GoalVersionDetails    fsugoalversiondetails             `json:"goalVersionDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.BatchingStrategy.UnmarshalPolymorphicJSON(model.BatchingStrategy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BatchingStrategy = nn.(CreateBatchingStrategyDetails)
	} else {
		m.BatchingStrategy = nil
	}

	nn, e = model.StageActionSchedule.UnmarshalPolymorphicJSON(model.StageActionSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StageActionSchedule = nn.(CreateScheduleDetails)
	} else {
		m.StageActionSchedule = nil
	}

	nn, e = model.ApplyActionSchedule.UnmarshalPolymorphicJSON(model.ApplyActionSchedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ApplyActionSchedule = nn.(CreateScheduleDetails)
	} else {
		m.ApplyActionSchedule = nil
	}

	m.DiagnosticsCollection = model.DiagnosticsCollection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.UpgradeDetails.UnmarshalPolymorphicJSON(model.UpgradeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.UpgradeDetails = nn.(UpgradeDetails)
	} else {
		m.UpgradeDetails = nil
	}

	m.CompartmentId = model.CompartmentId

	m.FsuCollectionId = model.FsuCollectionId

	nn, e = model.GoalVersionDetails.UnmarshalPolymorphicJSON(model.GoalVersionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.GoalVersionDetails = nn.(FsuGoalVersionDetails)
	} else {
		m.GoalVersionDetails = nil
	}

	return
}
