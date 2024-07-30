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

// FsuJob Exadata Fleet Update Job resource.
type FsuJob interface {

	// OCID identifier for the Exadata Fleet Update Job.
	GetId() *string

	// Compartment Identifier, this will map to the owner Exadata Fleet Update Action resource.
	GetCompartmentId() *string

	// OCID of the Exadata Fleet Update Action that this job is part of.
	GetFsuActionId() *string

	// The time the Exadata Fleet Update Job was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The current state of the Exadata Fleet Update Job.
	GetLifecycleState() JobLifecycleStatesEnum

	// Exadata Fleet Update Job display name.
	GetDisplayName() *string

	GetProgress() *JobProgressDetails

	// The time the Exadata Fleet Update Job started execution. An RFC3339 formatted datetime string.
	GetTimeStarted() *common.SDKTime

	// The time the Exadata Fleet Update Job was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// The time the Exadata Fleet Update Job completed execution. An RFC3339 formatted datetime string.
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

type fsujob struct {
	JsonData         []byte
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	Progress         *JobProgressDetails               `mandatory:"false" json:"progress"`
	TimeStarted      *common.SDKTime                   `mandatory:"false" json:"timeStarted"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	TimeFinished     *common.SDKTime                   `mandatory:"false" json:"timeFinished"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id               *string                           `mandatory:"true" json:"id"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	FsuActionId      *string                           `mandatory:"true" json:"fsuActionId"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState   JobLifecycleStatesEnum            `mandatory:"true" json:"lifecycleState"`
	Type             string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fsujob) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfsujob fsujob
	s := struct {
		Model Unmarshalerfsujob
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.FsuActionId = s.Model.FsuActionId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.Progress = s.Model.Progress
	m.TimeStarted = s.Model.TimeStarted
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
func (m *fsujob) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "APPLY":
		mm := ApplyFsuJob{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STAGE":
		mm := StageFsuJob{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRECHECK":
		mm := PrecheckFsuJob{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ROLLBACK_AND_REMOVE_TARGET":
		mm := RollbackFsuJob{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLEANUP":
		mm := CleanupFsuJob{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FsuJob: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m fsujob) GetDisplayName() *string {
	return m.DisplayName
}

// GetProgress returns Progress
func (m fsujob) GetProgress() *JobProgressDetails {
	return m.Progress
}

// GetTimeStarted returns TimeStarted
func (m fsujob) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeUpdated returns TimeUpdated
func (m fsujob) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeFinished returns TimeFinished
func (m fsujob) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleDetails returns LifecycleDetails
func (m fsujob) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m fsujob) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m fsujob) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m fsujob) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m fsujob) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m fsujob) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFsuActionId returns FsuActionId
func (m fsujob) GetFsuActionId() *string {
	return m.FsuActionId
}

// GetTimeCreated returns TimeCreated
func (m fsujob) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m fsujob) GetLifecycleState() JobLifecycleStatesEnum {
	return m.LifecycleState
}

func (m fsujob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fsujob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
