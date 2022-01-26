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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ConfigurationSourceProvider The properties that define a configuration source provider.
// For more information, see
// Managing Configuration Source Providers (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/managingconfigurationsourceproviders.htm).
type ConfigurationSourceProvider interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration source provider.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the configuration source provider is located.
	GetCompartmentId() *string

	// Human-readable display name for the configuration source provider.
	GetDisplayName() *string

	// Description of the configuration source provider.
	GetDescription() *string

	// The date and time when the configuration source provider was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The current lifecycle state of the configuration source provider.
	// For more information about configuration source provider lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__CSPStates).
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

type configurationsourceprovider struct {
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
func (m *configurationsourceprovider) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigurationsourceprovider configurationsourceprovider
	s := struct {
		Model Unmarshalerconfigurationsourceprovider
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
func (m *configurationsourceprovider) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceProviderType {
	case "GITHUB_ACCESS_TOKEN":
		mm := GithubAccessTokenConfigurationSourceProvider{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_ACCESS_TOKEN":
		mm := GitlabAccessTokenConfigurationSourceProvider{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m configurationsourceprovider) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m configurationsourceprovider) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m configurationsourceprovider) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m configurationsourceprovider) GetDescription() *string {
	return m.Description
}

//GetTimeCreated returns TimeCreated
func (m configurationsourceprovider) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetLifecycleState returns LifecycleState
func (m configurationsourceprovider) GetLifecycleState() ConfigurationSourceProviderLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m configurationsourceprovider) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m configurationsourceprovider) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m configurationsourceprovider) String() string {
	return common.PointerString(m)
}

// ConfigurationSourceProviderLifecycleStateEnum Enum with underlying type: string
type ConfigurationSourceProviderLifecycleStateEnum string

// Set of constants representing the allowable values for ConfigurationSourceProviderLifecycleStateEnum
const (
	ConfigurationSourceProviderLifecycleStateActive ConfigurationSourceProviderLifecycleStateEnum = "ACTIVE"
)

var mappingConfigurationSourceProviderLifecycleState = map[string]ConfigurationSourceProviderLifecycleStateEnum{
	"ACTIVE": ConfigurationSourceProviderLifecycleStateActive,
}

// GetConfigurationSourceProviderLifecycleStateEnumValues Enumerates the set of values for ConfigurationSourceProviderLifecycleStateEnum
func GetConfigurationSourceProviderLifecycleStateEnumValues() []ConfigurationSourceProviderLifecycleStateEnum {
	values := make([]ConfigurationSourceProviderLifecycleStateEnum, 0)
	for _, v := range mappingConfigurationSourceProviderLifecycleState {
		values = append(values, v)
	}
	return values
}

// ConfigurationSourceProviderConfigSourceProviderTypeEnum Enum with underlying type: string
type ConfigurationSourceProviderConfigSourceProviderTypeEnum string

// Set of constants representing the allowable values for ConfigurationSourceProviderConfigSourceProviderTypeEnum
const (
	ConfigurationSourceProviderConfigSourceProviderTypeGitlabAccessToken ConfigurationSourceProviderConfigSourceProviderTypeEnum = "GITLAB_ACCESS_TOKEN"
	ConfigurationSourceProviderConfigSourceProviderTypeGithubAccessToken ConfigurationSourceProviderConfigSourceProviderTypeEnum = "GITHUB_ACCESS_TOKEN"
)

var mappingConfigurationSourceProviderConfigSourceProviderType = map[string]ConfigurationSourceProviderConfigSourceProviderTypeEnum{
	"GITLAB_ACCESS_TOKEN": ConfigurationSourceProviderConfigSourceProviderTypeGitlabAccessToken,
	"GITHUB_ACCESS_TOKEN": ConfigurationSourceProviderConfigSourceProviderTypeGithubAccessToken,
}

// GetConfigurationSourceProviderConfigSourceProviderTypeEnumValues Enumerates the set of values for ConfigurationSourceProviderConfigSourceProviderTypeEnum
func GetConfigurationSourceProviderConfigSourceProviderTypeEnumValues() []ConfigurationSourceProviderConfigSourceProviderTypeEnum {
	values := make([]ConfigurationSourceProviderConfigSourceProviderTypeEnum, 0)
	for _, v := range mappingConfigurationSourceProviderConfigSourceProviderType {
		values = append(values, v)
	}
	return values
}
