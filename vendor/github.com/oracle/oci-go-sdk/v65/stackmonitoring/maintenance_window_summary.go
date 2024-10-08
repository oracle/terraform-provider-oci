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

// MaintenanceWindowSummary General information of a Maintenance Window
type MaintenanceWindowSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of maintenance window.
	Id *string `mandatory:"true" json:"id"`

	// Maintenance Window name.
	Name *string `mandatory:"true" json:"name"`

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Number of resources of the Maintenance window.
	NumberOfResources *int `mandatory:"false" json:"numberOfResources"`

	// Lifecycle state of the monitored resource.
	LifecycleState MaintenanceWindowLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Lifecycle Details of the Maintenance Window.
	LifecycleDetails MaintenanceWindowLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The name of the most recent operation of the Maintenance window.
	OperationType MaintenanceWindowOperationTypeEnum `mandatory:"false" json:"operationType,omitempty"`

	// Status of the most recent operation of the Maintenance Window.
	OperationStatus MaintenanceWindowOperationStatusEnum `mandatory:"false" json:"operationStatus,omitempty"`

	Schedule MaintenanceWindowSchedule `mandatory:"false" json:"schedule"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MaintenanceWindowSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceWindowSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMaintenanceWindowLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaintenanceWindowLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceWindowLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetMaintenanceWindowLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceWindowOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetMaintenanceWindowOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceWindowOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetMaintenanceWindowOperationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MaintenanceWindowSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId     *string                               `json:"compartmentId"`
		NumberOfResources *int                                  `json:"numberOfResources"`
		LifecycleState    MaintenanceWindowLifecycleStateEnum   `json:"lifecycleState"`
		LifecycleDetails  MaintenanceWindowLifecycleDetailsEnum `json:"lifecycleDetails"`
		OperationType     MaintenanceWindowOperationTypeEnum    `json:"operationType"`
		OperationStatus   MaintenanceWindowOperationStatusEnum  `json:"operationStatus"`
		Schedule          maintenancewindowschedule             `json:"schedule"`
		FreeformTags      map[string]string                     `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags        map[string]map[string]interface{}     `json:"systemTags"`
		Id                *string                               `json:"id"`
		Name              *string                               `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.NumberOfResources = model.NumberOfResources

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.OperationType = model.OperationType

	m.OperationStatus = model.OperationStatus

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

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	return
}
