// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// MonitoredResourceTask The request details for importing resources from Telemetry.
type MonitoredResourceTask struct {

	// Task identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// Name of the task.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	TaskDetails MonitoredResourceTaskDetails `mandatory:"true" json:"taskDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy.
	TenantId *string `mandatory:"false" json:"tenantId"`

	// Identifiers OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for work requests submitted for this task.
	WorkRequestIds []string `mandatory:"false" json:"workRequestIds"`

	// The date and time when the stack monitoring resource task was created, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the stack monitoring resource task was last updated, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the stack monitoring resource task.
	LifecycleState MonitoredResourceTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m MonitoredResourceTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredResourceTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMonitoredResourceTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMonitoredResourceTaskLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *MonitoredResourceTask) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TenantId       *string                                 `json:"tenantId"`
		WorkRequestIds []string                                `json:"workRequestIds"`
		TimeCreated    *common.SDKTime                         `json:"timeCreated"`
		TimeUpdated    *common.SDKTime                         `json:"timeUpdated"`
		LifecycleState MonitoredResourceTaskLifecycleStateEnum `json:"lifecycleState"`
		FreeformTags   map[string]string                       `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{}       `json:"definedTags"`
		SystemTags     map[string]map[string]interface{}       `json:"systemTags"`
		Id             *string                                 `json:"id"`
		Name           *string                                 `json:"name"`
		CompartmentId  *string                                 `json:"compartmentId"`
		TaskDetails    monitoredresourcetaskdetails            `json:"taskDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TenantId = model.TenantId

	m.WorkRequestIds = make([]string, len(model.WorkRequestIds))
	copy(m.WorkRequestIds, model.WorkRequestIds)
	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	nn, e = model.TaskDetails.UnmarshalPolymorphicJSON(model.TaskDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TaskDetails = nn.(MonitoredResourceTaskDetails)
	} else {
		m.TaskDetails = nil
	}

	return
}
