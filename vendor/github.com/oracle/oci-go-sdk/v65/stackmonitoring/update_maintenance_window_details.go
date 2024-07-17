// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMaintenanceWindowDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description *string                                  `json:"description"`
		Resources   []CreateMaintenanceWindowResourceDetails `json:"resources"`
		Schedule    maintenancewindowschedule                `json:"schedule"`
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

	return
}
