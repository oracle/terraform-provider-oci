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

// Channel A Channel connecting a DB System to an external entity.
type Channel struct {

	// The OCID of the Channel.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Channel. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Whether the Channel has been enabled by the user.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	Source ChannelSource `mandatory:"true" json:"source"`

	Target ChannelTarget `mandatory:"true" json:"target"`

	// The state of the Channel.
	LifecycleState ChannelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Channel was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Channel was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// User provided description of the Channel.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the state of the Channel.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Channel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Channel) ValidateEnumValue() (bool, error) {
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
func (m *Channel) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		DisplayName      *string                           `json:"displayName"`
		IsEnabled        *bool                             `json:"isEnabled"`
		Source           channelsource                     `json:"source"`
		Target           channeltarget                     `json:"target"`
		LifecycleState   ChannelLifecycleStateEnum         `json:"lifecycleState"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

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

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// ChannelLifecycleStateEnum Enum with underlying type: string
type ChannelLifecycleStateEnum string

// Set of constants representing the allowable values for ChannelLifecycleStateEnum
const (
	ChannelLifecycleStateCreating       ChannelLifecycleStateEnum = "CREATING"
	ChannelLifecycleStateActive         ChannelLifecycleStateEnum = "ACTIVE"
	ChannelLifecycleStateNeedsAttention ChannelLifecycleStateEnum = "NEEDS_ATTENTION"
	ChannelLifecycleStateInactive       ChannelLifecycleStateEnum = "INACTIVE"
	ChannelLifecycleStateUpdating       ChannelLifecycleStateEnum = "UPDATING"
	ChannelLifecycleStateDeleting       ChannelLifecycleStateEnum = "DELETING"
	ChannelLifecycleStateDeleted        ChannelLifecycleStateEnum = "DELETED"
	ChannelLifecycleStateFailed         ChannelLifecycleStateEnum = "FAILED"
)

var mappingChannelLifecycleStateEnum = map[string]ChannelLifecycleStateEnum{
	"CREATING":        ChannelLifecycleStateCreating,
	"ACTIVE":          ChannelLifecycleStateActive,
	"NEEDS_ATTENTION": ChannelLifecycleStateNeedsAttention,
	"INACTIVE":        ChannelLifecycleStateInactive,
	"UPDATING":        ChannelLifecycleStateUpdating,
	"DELETING":        ChannelLifecycleStateDeleting,
	"DELETED":         ChannelLifecycleStateDeleted,
	"FAILED":          ChannelLifecycleStateFailed,
}

var mappingChannelLifecycleStateEnumLowerCase = map[string]ChannelLifecycleStateEnum{
	"creating":        ChannelLifecycleStateCreating,
	"active":          ChannelLifecycleStateActive,
	"needs_attention": ChannelLifecycleStateNeedsAttention,
	"inactive":        ChannelLifecycleStateInactive,
	"updating":        ChannelLifecycleStateUpdating,
	"deleting":        ChannelLifecycleStateDeleting,
	"deleted":         ChannelLifecycleStateDeleted,
	"failed":          ChannelLifecycleStateFailed,
}

// GetChannelLifecycleStateEnumValues Enumerates the set of values for ChannelLifecycleStateEnum
func GetChannelLifecycleStateEnumValues() []ChannelLifecycleStateEnum {
	values := make([]ChannelLifecycleStateEnum, 0)
	for _, v := range mappingChannelLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetChannelLifecycleStateEnumStringValues Enumerates the set of values in String for ChannelLifecycleStateEnum
func GetChannelLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingChannelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChannelLifecycleStateEnum(val string) (ChannelLifecycleStateEnum, bool) {
	enum, ok := mappingChannelLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
