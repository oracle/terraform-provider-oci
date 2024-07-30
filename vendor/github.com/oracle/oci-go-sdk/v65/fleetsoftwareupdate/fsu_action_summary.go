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

// FsuActionSummary Exadata Fleet Update Action summary.
type FsuActionSummary interface {

	// OCID identifier for the Exadata Fleet Update Action.
	GetId() *string

	// Compartment Identifier.
	GetCompartmentId() *string

	// The date and time the Action was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeCreated() *common.SDKTime

	// The current state of the Exadata Fleet Update Action.
	GetLifecycleState() ActionLifecycleStatesEnum

	// Exadata Fleet Update Action display name.
	GetDisplayName() *string

	// The date and time the Action was started,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	GetTimeStarted() *common.SDKTime

	// The date and time the Action was finished,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	GetTimeFinished() *common.SDKTime

	// The date and time the Action was last updated, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeUpdated() *common.SDKTime

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

type fsuactionsummary struct {
	JsonData         []byte
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	TimeStarted      *common.SDKTime                   `mandatory:"false" json:"timeStarted"`
	TimeFinished     *common.SDKTime                   `mandatory:"false" json:"timeFinished"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id               *string                           `mandatory:"true" json:"id"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState   ActionLifecycleStatesEnum         `mandatory:"true" json:"lifecycleState"`
	Type             string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fsuactionsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfsuactionsummary fsuactionsummary
	s := struct {
		Model Unmarshalerfsuactionsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.TimeStarted = s.Model.TimeStarted
	m.TimeFinished = s.Model.TimeFinished
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fsuactionsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STAGE":
		mm := StageActionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLY":
		mm := ApplyActionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLEANUP":
		mm := CleanupActionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ROLLBACK_AND_REMOVE_TARGET":
		mm := RollbackActionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRECHECK":
		mm := PrecheckActionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FsuActionSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m fsuactionsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeStarted returns TimeStarted
func (m fsuactionsummary) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeFinished returns TimeFinished
func (m fsuactionsummary) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetTimeUpdated returns TimeUpdated
func (m fsuactionsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m fsuactionsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m fsuactionsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m fsuactionsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m fsuactionsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m fsuactionsummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m fsuactionsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m fsuactionsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m fsuactionsummary) GetLifecycleState() ActionLifecycleStatesEnum {
	return m.LifecycleState
}

func (m fsuactionsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fsuactionsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingActionLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetActionLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
