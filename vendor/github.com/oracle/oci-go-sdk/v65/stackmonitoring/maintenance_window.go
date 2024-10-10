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

// MaintenanceWindow Maintenance Window object. It contains all the information of the Maintenance window.
// Used in the Create and Get operations.
type MaintenanceWindow struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of maintenance window.
	Id *string `mandatory:"true" json:"id"`

	// Maintenance Window name.
	Name *string `mandatory:"true" json:"name"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Maintenance Window description.
	Description *string `mandatory:"false" json:"description"`

	// List of resource Ids which are part of the Maintenance Window
	Resources []CreateMaintenanceWindowResourceDetails `mandatory:"false" json:"resources"`

	// List of resource details that are part of the Maintenance Window.
	ResourcesDetails []MonitoredResourceDetails `mandatory:"false" json:"resourcesDetails"`

	// Lifecycle state of the monitored resource.
	LifecycleState MaintenanceWindowLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Lifecycle Details of the Maintenance Window.
	LifecycleDetails MaintenanceWindowLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	Schedule MaintenanceWindowSchedule `mandatory:"false" json:"schedule"`

	// The time the the maintenance window was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the the mainteance window was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m MaintenanceWindow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceWindow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMaintenanceWindowLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaintenanceWindowLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceWindowLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetMaintenanceWindowLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MaintenanceWindow) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                                  `json:"description"`
		Resources        []CreateMaintenanceWindowResourceDetails `json:"resources"`
		ResourcesDetails []MonitoredResourceDetails               `json:"resourcesDetails"`
		LifecycleState   MaintenanceWindowLifecycleStateEnum      `json:"lifecycleState"`
		LifecycleDetails MaintenanceWindowLifecycleDetailsEnum    `json:"lifecycleDetails"`
		Schedule         maintenancewindowschedule                `json:"schedule"`
		TimeCreated      *common.SDKTime                          `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                          `json:"timeUpdated"`
		Id               *string                                  `json:"id"`
		Name             *string                                  `json:"name"`
		CompartmentId    *string                                  `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Resources = make([]CreateMaintenanceWindowResourceDetails, len(model.Resources))
	copy(m.Resources, model.Resources)
	m.ResourcesDetails = make([]MonitoredResourceDetails, len(model.ResourcesDetails))
	copy(m.ResourcesDetails, model.ResourcesDetails)
	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	nn, e = model.Schedule.UnmarshalPolymorphicJSON(model.Schedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schedule = nn.(MaintenanceWindowSchedule)
	} else {
		m.Schedule = nil
	}

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.Id = model.Id

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	return
}
