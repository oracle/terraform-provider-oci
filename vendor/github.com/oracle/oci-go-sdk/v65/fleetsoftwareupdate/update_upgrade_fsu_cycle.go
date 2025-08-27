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

// UpdateUpgradeFsuCycle Update Upgrade Exadata Fleet Update Cycle resource details.
type UpdateUpgradeFsuCycle struct {

	// Exadata Fleet Update Cycle display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	GoalVersionDetails FsuGoalVersionDetails `mandatory:"false" json:"goalVersionDetails"`

	BatchingStrategy UpdateBatchingStrategyDetails `mandatory:"false" json:"batchingStrategy"`

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
func (m UpdateUpgradeFsuCycle) GetDisplayName() *string {
	return m.DisplayName
}

// GetGoalVersionDetails returns GoalVersionDetails
func (m UpdateUpgradeFsuCycle) GetGoalVersionDetails() FsuGoalVersionDetails {
	return m.GoalVersionDetails
}

// GetBatchingStrategy returns BatchingStrategy
func (m UpdateUpgradeFsuCycle) GetBatchingStrategy() UpdateBatchingStrategyDetails {
	return m.BatchingStrategy
}

// GetDiagnosticsCollection returns DiagnosticsCollection
func (m UpdateUpgradeFsuCycle) GetDiagnosticsCollection() *DiagnosticsCollectionDetails {
	return m.DiagnosticsCollection
}

// GetFreeformTags returns FreeformTags
func (m UpdateUpgradeFsuCycle) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateUpgradeFsuCycle) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateUpgradeFsuCycle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateUpgradeFsuCycle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateUpgradeFsuCycle) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateUpgradeFsuCycle UpdateUpgradeFsuCycle
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateUpgradeFsuCycle
	}{
		"UPGRADE",
		(MarshalTypeUpdateUpgradeFsuCycle)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateUpgradeFsuCycle) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                           `json:"displayName"`
		GoalVersionDetails    fsugoalversiondetails             `json:"goalVersionDetails"`
		BatchingStrategy      updatebatchingstrategydetails     `json:"batchingStrategy"`
		DiagnosticsCollection *DiagnosticsCollectionDetails     `json:"diagnosticsCollection"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		UpgradeDetails        upgradedetails                    `json:"upgradeDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

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
		m.BatchingStrategy = nn.(UpdateBatchingStrategyDetails)
	} else {
		m.BatchingStrategy = nil
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

	return
}
