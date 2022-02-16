// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateStandardTaskDetails The details for updating a schedule task.
type UpdateStandardTaskDetails struct {

	// A user-friendly name that is changeable and that does not have to be unique.
	// Format: a leading alphanumeric, followed by zero or more
	// alphanumerics, underscores, spaces, backslashes, or hyphens in any order).
	// No trailing spaces allowed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Schedules may be updated for task types SAVED_SEARCH and PURGE.
	// Note there may only be a single schedule for SAVED_SEARCH and PURGE scheduled tasks.
	Schedules []Schedule `mandatory:"false" json:"schedules"`

	Action Action `mandatory:"false" json:"action"`
}

//GetDisplayName returns DisplayName
func (m UpdateStandardTaskDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m UpdateStandardTaskDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateStandardTaskDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSchedules returns Schedules
func (m UpdateStandardTaskDetails) GetSchedules() []Schedule {
	return m.Schedules
}

func (m UpdateStandardTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateStandardTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateStandardTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateStandardTaskDetails UpdateStandardTaskDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeUpdateStandardTaskDetails
	}{
		"STANDARD",
		(MarshalTypeUpdateStandardTaskDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateStandardTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName  *string                           `json:"displayName"`
		FreeformTags map[string]string                 `json:"freeformTags"`
		DefinedTags  map[string]map[string]interface{} `json:"definedTags"`
		Schedules    []schedule                        `json:"schedules"`
		Action       action                            `json:"action"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Schedules = make([]Schedule, len(model.Schedules))
	for i, n := range model.Schedules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Schedules[i] = nn.(Schedule)
		} else {
			m.Schedules[i] = nil
		}
	}

	nn, e = model.Action.UnmarshalPolymorphicJSON(model.Action.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Action = nn.(Action)
	} else {
		m.Action = nil
	}

	return
}
