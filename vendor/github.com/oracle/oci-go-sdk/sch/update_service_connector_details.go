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

// UpdateServiceConnectorDetails The configuration details for updating a service connector.
type UpdateServiceConnectorDetails struct {

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the resource. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	Source SourceDetails `mandatory:"false" json:"source"`

	// The list of the tasks.
	Tasks []TaskDetails `mandatory:"false" json:"tasks"`

	Target TargetDetails `mandatory:"false" json:"target"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateServiceConnectorDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateServiceConnectorDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName  *string                           `json:"displayName"`
		Description  *string                           `json:"description"`
		Source       sourcedetails                     `json:"source"`
		Tasks        []taskdetails                     `json:"tasks"`
		Target       targetdetails                     `json:"target"`
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

	return
}
