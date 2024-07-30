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

// FsuJobSummary Exadata Fleet Update Job resource.
type FsuJobSummary interface {

	// OCID identifier for the Exadata Fleet Update Job.
	GetId() *string

	// Exadata Fleet Update Job display name.
	GetDisplayName() *string

	// Compartment Identifier, this will map to the owner Exadata Fleet Update Action resource.
	GetCompartmentId() *string

	// OCID of the Exadata Fleet Update Action that this job is part of.
	GetFsuActionId() *string

	GetProgress() *JobProgress

	// The time the Exadata Fleet Update Job was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the Exadata Fleet Update Job started execution. An RFC3339 formatted datetime string.
	GetTimeStarted() *common.SDKTime

	// The time the Exadata Fleet Update Job was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// The time the Exadata Fleet Update Job completed execution. An RFC3339 formatted datetime string.
	GetTimeFinished() *common.SDKTime

	// The current state of the Job.
	GetLifecycleState() JobLifecycleStatesEnum

	// A message describing the current state in more detail.
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

type fsujobsummary struct {
	JsonData         []byte
	Id               *string                           `mandatory:"false" json:"id"`
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	CompartmentId    *string                           `mandatory:"false" json:"compartmentId"`
	FsuActionId      *string                           `mandatory:"false" json:"fsuActionId"`
	Progress         *JobProgress                      `mandatory:"false" json:"progress"`
	TimeCreated      *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeStarted      *common.SDKTime                   `mandatory:"false" json:"timeStarted"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	TimeFinished     *common.SDKTime                   `mandatory:"false" json:"timeFinished"`
	LifecycleState   JobLifecycleStatesEnum            `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Type             string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fsujobsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfsujobsummary fsujobsummary
	s := struct {
		Model Unmarshalerfsujobsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.FsuActionId = s.Model.FsuActionId
	m.Progress = s.Model.Progress
	m.TimeCreated = s.Model.TimeCreated
	m.TimeStarted = s.Model.TimeStarted
	m.TimeUpdated = s.Model.TimeUpdated
	m.TimeFinished = s.Model.TimeFinished
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fsujobsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "STAGE":
		mm := StageFsuJobSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRECHECK":
		mm := PrecheckFsuJobSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLY":
		mm := ApplyFsuJobSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLEANUP":
		mm := CleanupFsuJobSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ROLLBACK_AND_REMOVE_TARGET":
		mm := RollbackFsuJobSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FsuJobSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetId returns Id
func (m fsujobsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m fsujobsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m fsujobsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFsuActionId returns FsuActionId
func (m fsujobsummary) GetFsuActionId() *string {
	return m.FsuActionId
}

// GetProgress returns Progress
func (m fsujobsummary) GetProgress() *JobProgress {
	return m.Progress
}

// GetTimeCreated returns TimeCreated
func (m fsujobsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeStarted returns TimeStarted
func (m fsujobsummary) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeUpdated returns TimeUpdated
func (m fsujobsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeFinished returns TimeFinished
func (m fsujobsummary) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleState returns LifecycleState
func (m fsujobsummary) GetLifecycleState() JobLifecycleStatesEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m fsujobsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m fsujobsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m fsujobsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m fsujobsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m fsujobsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fsujobsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
