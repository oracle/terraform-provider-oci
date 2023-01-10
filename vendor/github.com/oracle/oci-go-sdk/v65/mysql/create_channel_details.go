// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateChannelDetails Details required to create a Channel.
type CreateChannelDetails struct {
	Source CreateChannelSourceDetails `mandatory:"true" json:"source"`

	Target CreateChannelTargetDetails `mandatory:"true" json:"target"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The user-friendly name for the Channel. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the Channel should be enabled upon creation. If set to true, the Channel
	// will be asynchronously started as a result of the create Channel operation.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// User provided information about the Channel.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateChannelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateChannelDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId *string                           `json:"compartmentId"`
		DisplayName   *string                           `json:"displayName"`
		IsEnabled     *bool                             `json:"isEnabled"`
		Description   *string                           `json:"description"`
		FreeformTags  map[string]string                 `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		Source        createchannelsourcedetails        `json:"source"`
		Target        createchanneltargetdetails        `json:"target"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.IsEnabled = model.IsEnabled

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(CreateChannelSourceDetails)
	} else {
		m.Source = nil
	}

	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(CreateChannelTargetDetails)
	} else {
		m.Target = nil
	}

	return
}
