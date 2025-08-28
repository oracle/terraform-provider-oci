// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSchedulerDefinitionDetails The information to be updated.
type UpdateSchedulerDefinitionDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	Schedule Schedule `mandatory:"false" json:"schedule"`

	// Action Groups associated with the Schedule.
	ActionGroups []ActionGroup `mandatory:"false" json:"actionGroups"`

	// Runbooks.
	RunBooks []OperationRunbook `mandatory:"false" json:"runBooks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateSchedulerDefinitionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSchedulerDefinitionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSchedulerDefinitionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName  *string                           `json:"displayName"`
		Description  *string                           `json:"description"`
		Schedule     schedule                          `json:"schedule"`
		ActionGroups []actiongroup                     `json:"actionGroups"`
		RunBooks     []OperationRunbook                `json:"runBooks"`
		FreeformTags map[string]string                 `json:"freeformTags"`
		DefinedTags  map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.Schedule.UnmarshalPolymorphicJSON(model.Schedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schedule = nn.(Schedule)
	} else {
		m.Schedule = nil
	}

	m.ActionGroups = make([]ActionGroup, len(model.ActionGroups))
	for i, n := range model.ActionGroups {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ActionGroups[i] = nn.(ActionGroup)
		} else {
			m.ActionGroups[i] = nil
		}
	}
	m.RunBooks = make([]OperationRunbook, len(model.RunBooks))
	copy(m.RunBooks, model.RunBooks)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
