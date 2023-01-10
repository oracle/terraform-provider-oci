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

// ChannelSummary Summary of a Channel.
type ChannelSummary struct {

	// The OCID of the Channel.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether the Channel has been enabled by the user.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	Source ChannelSource `mandatory:"true" json:"source"`

	Target ChannelTarget `mandatory:"true" json:"target"`

	// The state of the Channel.
	LifecycleState ChannelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the Channel. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the Channel was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Channel was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A message describing the state of the Channel.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ChannelSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChannelSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChannelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetChannelLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ChannelSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		IsEnabled        *bool                             `json:"isEnabled"`
		Source           channelsource                     `json:"source"`
		Target           channeltarget                     `json:"target"`
		LifecycleState   ChannelLifecycleStateEnum         `json:"lifecycleState"`
		DisplayName      *string                           `json:"displayName"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.IsEnabled = model.IsEnabled

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(ChannelSource)
	} else {
		m.Source = nil
	}

	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(ChannelTarget)
	} else {
		m.Target = nil
	}

	m.LifecycleState = model.LifecycleState

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}
