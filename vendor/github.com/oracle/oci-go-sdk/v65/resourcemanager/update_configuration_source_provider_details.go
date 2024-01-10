// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateConfigurationSourceProviderDetails Update details for a configuration source provider.
type UpdateConfigurationSourceProviderDetails interface {

	// Human-readable name of the configuration source provider. Avoid entering confidential information.
	GetDisplayName() *string

	// Description of the configuration source provider. Avoid entering confidential information.
	GetDescription() *string

	GetPrivateServerConfigDetails() *PrivateServerConfigDetails

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updateconfigurationsourceproviderdetails struct {
	JsonData                   []byte
	DisplayName                *string                           `mandatory:"false" json:"displayName"`
	Description                *string                           `mandatory:"false" json:"description"`
	PrivateServerConfigDetails *PrivateServerConfigDetails       `mandatory:"false" json:"privateServerConfigDetails"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ConfigSourceProviderType   string                            `json:"configSourceProviderType"`
}

// UnmarshalJSON unmarshals json
func (m *updateconfigurationsourceproviderdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconfigurationsourceproviderdetails updateconfigurationsourceproviderdetails
	s := struct {
		Model Unmarshalerupdateconfigurationsourceproviderdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.PrivateServerConfigDetails = s.Model.PrivateServerConfigDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConfigSourceProviderType = s.Model.ConfigSourceProviderType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconfigurationsourceproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceProviderType {
	case "BITBUCKET_CLOUD_USERNAME_APPPASSWORD":
		mm := UpdateBitbucketCloudUsernameAppPasswordConfigurationSourceProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_SERVER_ACCESS_TOKEN":
		mm := UpdateBitbucketServerAccessTokenConfigurationSourceProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_ACCESS_TOKEN":
		mm := UpdateGitlabAccessTokenConfigurationSourceProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB_ACCESS_TOKEN":
		mm := UpdateGithubAccessTokenConfigurationSourceProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateConfigurationSourceProviderDetails: %s.", m.ConfigSourceProviderType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updateconfigurationsourceproviderdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m updateconfigurationsourceproviderdetails) GetDescription() *string {
	return m.Description
}

// GetPrivateServerConfigDetails returns PrivateServerConfigDetails
func (m updateconfigurationsourceproviderdetails) GetPrivateServerConfigDetails() *PrivateServerConfigDetails {
	return m.PrivateServerConfigDetails
}

// GetFreeformTags returns FreeformTags
func (m updateconfigurationsourceproviderdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updateconfigurationsourceproviderdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updateconfigurationsourceproviderdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateconfigurationsourceproviderdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
