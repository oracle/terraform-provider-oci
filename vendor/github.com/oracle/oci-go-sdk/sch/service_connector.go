// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ServiceConnector The configuration details of the flow defined by the service connector.
// For more information about flows defined by service connectors, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
type ServiceConnector struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service connector.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the service connector.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time when the service connector was created.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time when the service connector was updated.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2020-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the service connector.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the resource. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail.
	// For example, the message might provide actionable
	// information for a resource in a `FAILED` state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	Source SourceDetails `mandatory:"false" json:"source"`

	// The list of tasks.
	Tasks []TaskDetails `mandatory:"false" json:"tasks"`

	Target TargetDetails `mandatory:"false" json:"target"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ServiceConnector) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ServiceConnector) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description     *string                           `json:"description"`
		LifecyleDetails *string                           `json:"lifecyleDetails"`
		Source          sourcedetails                     `json:"source"`
		Tasks           []taskdetails                     `json:"tasks"`
		Target          targetdetails                     `json:"target"`
		FreeformTags    map[string]string                 `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{} `json:"definedTags"`
		SystemTags      map[string]map[string]interface{} `json:"systemTags"`
		Id              *string                           `json:"id"`
		DisplayName     *string                           `json:"displayName"`
		CompartmentId   *string                           `json:"compartmentId"`
		TimeCreated     *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated     *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState  LifecycleStateEnum                `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecyleDetails = model.LifecyleDetails

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(SourceDetails)
	} else {
		m.Source = nil
	}

	m.Tasks = make([]TaskDetails, len(model.Tasks))
	for i, n := range model.Tasks {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Tasks[i] = nn.(TaskDetails)
		} else {
			m.Tasks[i] = nil
		}
	}

	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(TargetDetails)
	} else {
		m.Target = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	return
}
