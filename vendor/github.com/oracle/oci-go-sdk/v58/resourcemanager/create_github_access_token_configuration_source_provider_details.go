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

// CreateGithubAccessTokenConfigurationSourceProviderDetails The details for creating a configuration source provider of the type `GITHUB_ACCESS_TOKEN`.
// This type corresponds to a configuration source provider in GitHub that is authenticated with a personal access token.
type CreateGithubAccessTokenConfigurationSourceProviderDetails struct {

	// The GitHub service endpoint.
	// Example: `https://github.com/`
	ApiEndpoint *string `mandatory:"true" json:"apiEndpoint"`

	// The personal access token to be configured on the GitHub repository. Avoid entering confidential information.
	AccessToken *string `mandatory:"true" json:"accessToken"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where
	// you want to create the configuration source provider.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable name of the configuration source provider. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the configuration source provider. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetCompartmentId returns CompartmentId
func (m CreateGithubAccessTokenConfigurationSourceProviderDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m CreateGithubAccessTokenConfigurationSourceProviderDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m CreateGithubAccessTokenConfigurationSourceProviderDetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m CreateGithubAccessTokenConfigurationSourceProviderDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateGithubAccessTokenConfigurationSourceProviderDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateGithubAccessTokenConfigurationSourceProviderDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGithubAccessTokenConfigurationSourceProviderDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateGithubAccessTokenConfigurationSourceProviderDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateGithubAccessTokenConfigurationSourceProviderDetails CreateGithubAccessTokenConfigurationSourceProviderDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceProviderType"`
		MarshalTypeCreateGithubAccessTokenConfigurationSourceProviderDetails
	}{
		"GITHUB_ACCESS_TOKEN",
		(MarshalTypeCreateGithubAccessTokenConfigurationSourceProviderDetails)(m),
	}

	return json.Marshal(&s)
}
