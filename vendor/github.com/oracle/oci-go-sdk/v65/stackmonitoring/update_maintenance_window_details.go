// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMaintenanceWindowDetails Infomation to create a new Maintenance Window.
type UpdateMaintenanceWindowDetails struct {

	// Maintenance Window description.
	Description *string `mandatory:"false" json:"description"`

	// List of resource Ids which are part of the Maintenance Window
	Resources []CreateMaintenanceWindowResourceDetails `mandatory:"false" json:"resources"`

	Schedule MaintenanceWindowSchedule `mandatory:"false" json:"schedule"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMaintenanceWindowDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMaintenanceWindowDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMaintenanceWindowDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description  *string                                  `json:"description"`
		Resources    []CreateMaintenanceWindowResourceDetails `json:"resources"`
		Schedule     maintenancewindowschedule                `json:"schedule"`
		FreeformTags map[string]string                        `json:"freeformTags"`
		DefinedTags  map[string]map[string]interface{}        `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Resources = make([]CreateMaintenanceWindowResourceDetails, len(model.Resources))
	copy(m.Resources, model.Resources)
	nn, e = model.Schedule.UnmarshalPolymorphicJSON(model.Schedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schedule = nn.(MaintenanceWindowSchedule)
	} else {
		m.Schedule = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
