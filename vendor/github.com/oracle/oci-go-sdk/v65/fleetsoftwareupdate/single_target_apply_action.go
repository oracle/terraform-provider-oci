// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SingleTargetApplyAction Single Target Apply Exadata Fleet Update Action details.
type SingleTargetApplyAction struct {

	// OCID identifier for the Exadata Fleet Update Action.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the Action was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Exadata Fleet Update Action display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the Action was started,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the Action was finished,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The date and time the Action was last updated, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

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

	// OCID identifier for the current Exadata Fleet Update Home of the Target.
	SourceFsuHomeId *string `mandatory:"false" json:"sourceFsuHomeId"`

	// OCID identifier for the goal Exadata Fleet Update Home of the Target for patching.
	GoalFsuHomeId *string `mandatory:"false" json:"goalFsuHomeId"`

	// Ignore all patches between the source and target homes during patching.
	IsIgnorePatches *bool `mandatory:"false" json:"isIgnorePatches"`

	// Ensure that services of administrator-managed Oracle RAC or Oracle RAC One databases are running on the same
	// instances before and after the move operation.
	IsKeepPlacement *bool `mandatory:"false" json:"isKeepPlacement"`

	// Service drain timeout specified in seconds.
	MaxDrainTimeoutInSeconds *int `mandatory:"false" json:"maxDrainTimeoutInSeconds"`

	// List of bug numbers to ignore.
	IgnoreMissingPatches []string `mandatory:"false" json:"ignoreMissingPatches"`

	Progress *SingleTargetFsuActionProgressDetails `mandatory:"false" json:"progress"`

	TargetDetails SingleTargetDetails `mandatory:"false" json:"targetDetails"`

	// The current state of the Exadata Fleet Update Action.
	LifecycleState ActionLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m SingleTargetApplyAction) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m SingleTargetApplyAction) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m SingleTargetApplyAction) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m SingleTargetApplyAction) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeStarted returns TimeStarted
func (m SingleTargetApplyAction) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeFinished returns TimeFinished
func (m SingleTargetApplyAction) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetTimeUpdated returns TimeUpdated
func (m SingleTargetApplyAction) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m SingleTargetApplyAction) GetLifecycleState() ActionLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m SingleTargetApplyAction) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m SingleTargetApplyAction) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m SingleTargetApplyAction) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m SingleTargetApplyAction) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m SingleTargetApplyAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SingleTargetApplyAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingActionLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetActionLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SingleTargetApplyAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSingleTargetApplyAction SingleTargetApplyAction
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSingleTargetApplyAction
	}{
		"SINGLE_TARGET_APPLY",
		(MarshalTypeSingleTargetApplyAction)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *SingleTargetApplyAction) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                               `json:"displayName"`
		TimeStarted              *common.SDKTime                       `json:"timeStarted"`
		TimeFinished             *common.SDKTime                       `json:"timeFinished"`
		TimeUpdated              *common.SDKTime                       `json:"timeUpdated"`
		LifecycleDetails         *string                               `json:"lifecycleDetails"`
		FreeformTags             map[string]string                     `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags               map[string]map[string]interface{}     `json:"systemTags"`
		SourceFsuHomeId          *string                               `json:"sourceFsuHomeId"`
		GoalFsuHomeId            *string                               `json:"goalFsuHomeId"`
		IsIgnorePatches          *bool                                 `json:"isIgnorePatches"`
		IsKeepPlacement          *bool                                 `json:"isKeepPlacement"`
		MaxDrainTimeoutInSeconds *int                                  `json:"maxDrainTimeoutInSeconds"`
		IgnoreMissingPatches     []string                              `json:"ignoreMissingPatches"`
		Progress                 *SingleTargetFsuActionProgressDetails `json:"progress"`
		TargetDetails            singletargetdetails                   `json:"targetDetails"`
		Id                       *string                               `json:"id"`
		CompartmentId            *string                               `json:"compartmentId"`
		TimeCreated              *common.SDKTime                       `json:"timeCreated"`
		LifecycleState           ActionLifecycleStatesEnum             `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.SourceFsuHomeId = model.SourceFsuHomeId

	m.GoalFsuHomeId = model.GoalFsuHomeId

	m.IsIgnorePatches = model.IsIgnorePatches

	m.IsKeepPlacement = model.IsKeepPlacement

	m.MaxDrainTimeoutInSeconds = model.MaxDrainTimeoutInSeconds

	m.IgnoreMissingPatches = make([]string, len(model.IgnoreMissingPatches))
	copy(m.IgnoreMissingPatches, model.IgnoreMissingPatches)
	m.Progress = model.Progress

	nn, e = model.TargetDetails.UnmarshalPolymorphicJSON(model.TargetDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetDetails = nn.(SingleTargetDetails)
	} else {
		m.TargetDetails = nil
	}

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}
