// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v39/common"
)

// GithubAccessTokenConfigurationSourceProvider The properties that define a configuration source provider of the type `GITHUB_ACCESS_TOKEN`.
// This type corresponds to a configuration source provider in GitHub that is authenticated with a personal access token.
type GithubAccessTokenConfigurationSourceProvider struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration source provider.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the configuration source provider is located.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable display name for the configuration source provider.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the configuration source provider.
	Description *string `mandatory:"false" json:"description"`

	// The date and time when the configuration source provider was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The GitHub service endpoint.
	// Example: `https://github.com/`
	ApiEndpoint *string `mandatory:"false" json:"apiEndpoint"`

	// The current lifecycle state of the configuration source provider.
	// For more information about configuration source provider lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#CSPStates).
	LifecycleState ConfigurationSourceProviderLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m GithubAccessTokenConfigurationSourceProvider) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m GithubAccessTokenConfigurationSourceProvider) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m GithubAccessTokenConfigurationSourceProvider) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m GithubAccessTokenConfigurationSourceProvider) GetDescription() *string {
	return m.Description
}

//GetTimeCreated returns TimeCreated
func (m GithubAccessTokenConfigurationSourceProvider) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetLifecycleState returns LifecycleState
func (m GithubAccessTokenConfigurationSourceProvider) GetLifecycleState() ConfigurationSourceProviderLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m GithubAccessTokenConfigurationSourceProvider) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m GithubAccessTokenConfigurationSourceProvider) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m GithubAccessTokenConfigurationSourceProvider) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GithubAccessTokenConfigurationSourceProvider) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGithubAccessTokenConfigurationSourceProvider GithubAccessTokenConfigurationSourceProvider
	s := struct {
		DiscriminatorParam string `json:"configSourceProviderType"`
		MarshalTypeGithubAccessTokenConfigurationSourceProvider
	}{
		"GITHUB_ACCESS_TOKEN",
		(MarshalTypeGithubAccessTokenConfigurationSourceProvider)(m),
	}

	return json.Marshal(&s)
}
