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

// CreateUnifiedAgentConfigurationDetails Unified Agent configuration creation object.
type CreateUnifiedAgentConfigurationDetails struct {

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	ServiceConfiguration UnifiedAgentServiceConfigurationDetails `mandatory:"true" json:"serviceConfiguration"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of a user-friendly name. It has to be unique within enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	GroupAssociation *GroupAssociationDetails `mandatory:"false" json:"groupAssociation"`
}

func (m CreateUnifiedAgentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateUnifiedAgentConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                                 `json:"displayName"`
		DefinedTags          map[string]map[string]interface{}       `json:"definedTags"`
		FreeformTags         map[string]string                       `json:"freeformTags"`
		Description          *string                                 `json:"description"`
		GroupAssociation     *GroupAssociationDetails                `json:"groupAssociation"`
		IsEnabled            *bool                                   `json:"isEnabled"`
		ServiceConfiguration unifiedagentserviceconfigurationdetails `json:"serviceConfiguration"`
		CompartmentId        *string                                 `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.Description = model.Description

	m.GroupAssociation = model.GroupAssociation

	m.IsEnabled = model.IsEnabled

	nn, e = model.ServiceConfiguration.UnmarshalPolymorphicJSON(model.ServiceConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ServiceConfiguration = nn.(UnifiedAgentServiceConfigurationDetails)
	} else {
		m.ServiceConfiguration = nil
	}

	m.CompartmentId = model.CompartmentId

	return
}
