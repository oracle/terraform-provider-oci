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

// CreateMaintenanceWindowDetails Infomation to create a new Maintenance Window.
type CreateMaintenanceWindowDetails struct {

	// Maintenance Window name.
	Name *string `mandatory:"true" json:"name"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of resource Ids which are part of the Maintenance Window
	Resources []CreateMaintenanceWindowResourceDetails `mandatory:"true" json:"resources"`

	Schedule MaintenanceWindowSchedule `mandatory:"true" json:"schedule"`

	// Maintenance Window description.
	Description *string `mandatory:"false" json:"description"`
}

func (m CreateMaintenanceWindowDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMaintenanceWindowDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateMaintenanceWindowDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description   *string                                  `json:"description"`
		Name          *string                                  `json:"name"`
		CompartmentId *string                                  `json:"compartmentId"`
		Resources     []CreateMaintenanceWindowResourceDetails `json:"resources"`
		Schedule      maintenancewindowschedule                `json:"schedule"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

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
