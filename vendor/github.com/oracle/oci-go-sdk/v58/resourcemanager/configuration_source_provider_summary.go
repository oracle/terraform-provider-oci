// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__CSPStates).
	// Allowable values:
	// - ACTIVE
	GetLifecycleState() ConfigurationSourceProviderLifecycleStateEnum

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
	Id                       *string                                       `mandatory:"false" json:"id"`
	CompartmentId            *string                                       `mandatory:"false" json:"compartmentId"`
	DisplayName              *string                                       `mandatory:"false" json:"displayName"`
	Description              *string                                       `mandatory:"false" json:"description"`
	TimeCreated              *common.SDKTime                               `mandatory:"false" json:"timeCreated"`
	LifecycleState           ConfigurationSourceProviderLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	FreeformTags             map[string]string                             `mandatory:"false" json:"freeformTags"`
	DefinedTags              map[string]map[string]interface{}             `mandatory:"false" json:"definedTags"`
	ConfigSourceProviderType string                                        `json:"configSourceProviderType"`
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
	case "GITHUB_ACCESS_TOKEN":
		mm := GithubAccessTokenConfigurationSourceProviderSummary{}
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
func (m configurationsourceprovidersummary) GetLifecycleState() ConfigurationSourceProviderLifecycleStateEnum {
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configurationsourceprovidersummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigurationSourceProviderLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigurationSourceProviderLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
