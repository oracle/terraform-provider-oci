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

// CreateSingleTargetApplyDetails Single Target Apply Action creation details.
type CreateSingleTargetApplyDetails struct {

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID identifier for the current Exadata Fleet Update Home of the Target.
	SourceFsuHomeId *string `mandatory:"true" json:"sourceFsuHomeId"`

	// OCID identifier for the goal Exadata Fleet Update Home of the Target for patching.
	GoalFsuHomeId *string `mandatory:"true" json:"goalFsuHomeId"`

	TargetDetails SingleTargetDetails `mandatory:"true" json:"targetDetails"`

	// Exadata Fleet Update Action display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Ignore all patches between the source and target homes during patching.
	IsIgnorePatches *bool `mandatory:"false" json:"isIgnorePatches"`

	// Ensure that services of administrator-managed Oracle RAC or Oracle RAC One databases are running on the same
	// instances before and after the move operation.
	IsKeepPlacement *bool `mandatory:"false" json:"isKeepPlacement"`

	// Service drain timeout specified in seconds.
	MaxDrainTimeoutInSeconds *int `mandatory:"false" json:"maxDrainTimeoutInSeconds"`

	// List of bug numbers to ignore.
	IgnoreMissingPatches []string `mandatory:"false" json:"ignoreMissingPatches"`
}

// GetDisplayName returns DisplayName
func (m CreateSingleTargetApplyDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateSingleTargetApplyDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateSingleTargetApplyDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateSingleTargetApplyDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateSingleTargetApplyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSingleTargetApplyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSingleTargetApplyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSingleTargetApplyDetails CreateSingleTargetApplyDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateSingleTargetApplyDetails
	}{
		"SINGLE_TARGET_APPLY",
		(MarshalTypeCreateSingleTargetApplyDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateSingleTargetApplyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                           `json:"displayName"`
		FreeformTags             map[string]string                 `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{} `json:"definedTags"`
		IsIgnorePatches          *bool                             `json:"isIgnorePatches"`
		IsKeepPlacement          *bool                             `json:"isKeepPlacement"`
		MaxDrainTimeoutInSeconds *int                              `json:"maxDrainTimeoutInSeconds"`
		IgnoreMissingPatches     []string                          `json:"ignoreMissingPatches"`
		CompartmentId            *string                           `json:"compartmentId"`
		SourceFsuHomeId          *string                           `json:"sourceFsuHomeId"`
		GoalFsuHomeId            *string                           `json:"goalFsuHomeId"`
		TargetDetails            singletargetdetails               `json:"targetDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.IsIgnorePatches = model.IsIgnorePatches

	m.IsKeepPlacement = model.IsKeepPlacement

	m.MaxDrainTimeoutInSeconds = model.MaxDrainTimeoutInSeconds

	m.IgnoreMissingPatches = make([]string, len(model.IgnoreMissingPatches))
	copy(m.IgnoreMissingPatches, model.IgnoreMissingPatches)
	m.CompartmentId = model.CompartmentId

	m.SourceFsuHomeId = model.SourceFsuHomeId

	m.GoalFsuHomeId = model.GoalFsuHomeId

	nn, e = model.TargetDetails.UnmarshalPolymorphicJSON(model.TargetDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetDetails = nn.(SingleTargetDetails)
	} else {
		m.TargetDetails = nil
	}

	return
}
