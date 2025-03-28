// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateUnifiedAgentConfigurationDetails Unified Agent configuration creation object.
type CreateUnifiedAgentConfigurationDetails struct {

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	ServiceConfiguration UnifiedAgentServiceConfigurationDetails `mandatory:"true" json:"serviceConfiguration"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Description for this resource.
	Description *string `mandatory:"true" json:"description"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	GroupAssociation *GroupAssociationDetails `mandatory:"false" json:"groupAssociation"`
}

func (m CreateUnifiedAgentConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateUnifiedAgentConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateUnifiedAgentConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags          map[string]map[string]interface{}       `json:"definedTags"`
		FreeformTags         map[string]string                       `json:"freeformTags"`
		GroupAssociation     *GroupAssociationDetails                `json:"groupAssociation"`
		DisplayName          *string                                 `json:"displayName"`
		IsEnabled            *bool                                   `json:"isEnabled"`
		ServiceConfiguration unifiedagentserviceconfigurationdetails `json:"serviceConfiguration"`
		CompartmentId        *string                                 `json:"compartmentId"`
		Description          *string                                 `json:"description"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.GroupAssociation = model.GroupAssociation

	m.DisplayName = model.DisplayName

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

	m.Description = model.Description

	return
}
