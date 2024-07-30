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

// CreateRollbackActionDetails Rollback Exadata Fleet Update Action creation details. This action will attempt to rollback the
// specified Targets according to strategy to the source target version prior to patching in this
// Exadata Fleet Update Cycle and remove them from the Collection.
type CreateRollbackActionDetails struct {

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID identifier for the Exadata Fleet Update Cycle the Action will be part of.
	FsuCycleId *string `mandatory:"true" json:"fsuCycleId"`

	Details RollbackDetails `mandatory:"true" json:"details"`

	// Exadata Fleet Update Action display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	ScheduleDetails CreateScheduleDetails `mandatory:"false" json:"scheduleDetails"`
}

// GetDisplayName returns DisplayName
func (m CreateRollbackActionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateRollbackActionDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateRollbackActionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateRollbackActionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateRollbackActionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRollbackActionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateRollbackActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateRollbackActionDetails CreateRollbackActionDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateRollbackActionDetails
	}{
		"ROLLBACK_AND_REMOVE_TARGET",
		(MarshalTypeCreateRollbackActionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateRollbackActionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName     *string                           `json:"displayName"`
		FreeformTags    map[string]string                 `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{} `json:"definedTags"`
		ScheduleDetails createscheduledetails             `json:"scheduleDetails"`
		CompartmentId   *string                           `json:"compartmentId"`
		FsuCycleId      *string                           `json:"fsuCycleId"`
		Details         rollbackdetails                   `json:"details"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.ScheduleDetails.UnmarshalPolymorphicJSON(model.ScheduleDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ScheduleDetails = nn.(CreateScheduleDetails)
	} else {
		m.ScheduleDetails = nil
	}

	m.CompartmentId = model.CompartmentId

	m.FsuCycleId = model.FsuCycleId

	nn, e = model.Details.UnmarshalPolymorphicJSON(model.Details.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Details = nn.(RollbackDetails)
	} else {
		m.Details = nil
	}

	return
}
