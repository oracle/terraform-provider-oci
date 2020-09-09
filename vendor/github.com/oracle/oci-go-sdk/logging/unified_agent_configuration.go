// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UnifiedAgentConfiguration Top Unified Agent configuration object.
type UnifiedAgentConfiguration struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of a user-friendly name. It has to be unique within enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of an pipeline.
	LifecycleState LogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// State of unified agent service configuration.
	ConfigurationState UnifiedAgentServiceConfigurationStatesEnum `mandatory:"true" json:"configurationState"`

	ServiceConfiguration UnifiedAgentServiceConfigurationDetails `mandatory:"true" json:"serviceConfiguration"`

	GroupAssociation *GroupAssociationDetails `mandatory:"true" json:"groupAssociation"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`
}

func (m UnifiedAgentConfiguration) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UnifiedAgentConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description          *string                                    `json:"description"`
		DefinedTags          map[string]map[string]interface{}          `json:"definedTags"`
		FreeformTags         map[string]string                          `json:"freeformTags"`
		TimeCreated          *common.SDKTime                            `json:"timeCreated"`
		TimeLastModified     *common.SDKTime                            `json:"timeLastModified"`
		Id                   *string                                    `json:"id"`
		CompartmentId        *string                                    `json:"compartmentId"`
		DisplayName          *string                                    `json:"displayName"`
		LifecycleState       LogLifecycleStateEnum                      `json:"lifecycleState"`
		IsEnabled            *bool                                      `json:"isEnabled"`
		ConfigurationState   UnifiedAgentServiceConfigurationStatesEnum `json:"configurationState"`
		ServiceConfiguration unifiedagentserviceconfigurationdetails    `json:"serviceConfiguration"`
		GroupAssociation     *GroupAssociationDetails                   `json:"groupAssociation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.TimeCreated = model.TimeCreated

	m.TimeLastModified = model.TimeLastModified

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.LifecycleState = model.LifecycleState

	m.IsEnabled = model.IsEnabled

	m.ConfigurationState = model.ConfigurationState

	nn, e = model.ServiceConfiguration.UnmarshalPolymorphicJSON(model.ServiceConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ServiceConfiguration = nn.(UnifiedAgentServiceConfigurationDetails)
	} else {
		m.ServiceConfiguration = nil
	}

	m.GroupAssociation = model.GroupAssociation

	return
}
