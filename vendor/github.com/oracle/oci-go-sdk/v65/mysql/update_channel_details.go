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

// UpdateChannelDetails Details required to update a Channel
type UpdateChannelDetails struct {
	Source UpdateChannelSourceDetails `mandatory:"false" json:"source"`

	Target UpdateChannelTargetDetails `mandatory:"false" json:"target"`

	// The user-friendly name for the Channel. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the Channel should be enabled or disabled. Enabling a previously
	// disabled Channel will cause the Channel to be started. Conversely, disabling
	// a previously enabled Channel will stop the Channel. Both operations are
	// executed asynchronously.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// User provided description of the Channel.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateChannelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateChannelDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Source       updatechannelsourcedetails        `json:"source"`
		Target       updatechanneltargetdetails        `json:"target"`
		DisplayName  *string                           `json:"displayName"`
		IsEnabled    *bool                             `json:"isEnabled"`
		Description  *string                           `json:"description"`
		FreeformTags map[string]string                 `json:"freeformTags"`
		DefinedTags  map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(UpdateChannelSourceDetails)
	} else {
		m.Source = nil
	}

	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(UpdateChannelTargetDetails)
	} else {
		m.Target = nil
	}

	m.DisplayName = model.DisplayName

	m.IsEnabled = model.IsEnabled

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
