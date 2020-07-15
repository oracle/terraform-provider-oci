// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ConfigurationSourceProviderSummary Summary information for a configuration source provider.
type ConfigurationSourceProviderSummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration source provider.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the configuration source provider is located.
	GetCompartmentId() *string

	// Human-readable display name for the configuration source provider.
	GetDisplayName() *string

	// General description of the configuration source provider.
	GetDescription() *string

	// The date and time when the configuration source provider was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// Current state of the specified configuration source provider.
	// For more information about configuration source provider lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#CSPStates).
	// Allowable values:
	// - ACTIVE
	GetLifecycleState() ConfigurationSourceProviderSummary

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type configurationsourceprovidersummary struct {
	JsonData                 []byte
	Id                       *string                            `mandatory:"false" json:"id"`
	CompartmentId            *string                            `mandatory:"false" json:"compartmentId"`
	DisplayName              *string                            `mandatory:"false" json:"displayName"`
	Description              *string                            `mandatory:"false" json:"description"`
	TimeCreated              *common.SDKTime                    `mandatory:"false" json:"timeCreated"`
	LifecycleState           ConfigurationSourceProviderSummary `mandatory:"false" json:"lifecycleState,omitempty"`
	FreeformTags             map[string]string                  `mandatory:"false" json:"freeformTags"`
	DefinedTags              map[string]map[string]interface{}  `mandatory:"false" json:"definedTags"`
	ConfigSourceProviderType string                             `json:"configSourceProviderType"`
}

// UnmarshalJSON unmarshals json
func (m *configurationsourceprovidersummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigurationsourceprovidersummary configurationsourceprovidersummary
	s := struct {
		Model Unmarshalerconfigurationsourceprovidersummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConfigSourceProviderType = s.Model.ConfigSourceProviderType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configurationsourceprovidersummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceProviderType {
	case "GITLAB_ACCESS_TOKEN":
		mm := GitlabAccessTokenConfigurationSourceProviderSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m configurationsourceprovidersummary) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m configurationsourceprovidersummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m configurationsourceprovidersummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m configurationsourceprovidersummary) GetDescription() *string {
	return m.Description
}

//GetTimeCreated returns TimeCreated
func (m configurationsourceprovidersummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetLifecycleState returns LifecycleState
func (m configurationsourceprovidersummary) GetLifecycleState() ConfigurationSourceProviderSummary {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m configurationsourceprovidersummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m configurationsourceprovidersummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m configurationsourceprovidersummary) String() string {
	return common.PointerString(m)
}
