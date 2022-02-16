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

// CreateConfigurationSourceProviderDetails The details for creating a configuration source provider.
type CreateConfigurationSourceProviderDetails interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where
	// you want to create the configuration source provider.
	GetCompartmentId() *string

	// Human-readable name of the configuration source provider. Avoid entering confidential information.
	GetDisplayName() *string

	// Description of the configuration source provider. Avoid entering confidential information.
	GetDescription() *string

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createconfigurationsourceproviderdetails struct {
	JsonData                 []byte
	CompartmentId            *string                           `mandatory:"false" json:"compartmentId"`
	DisplayName              *string                           `mandatory:"false" json:"displayName"`
	Description              *string                           `mandatory:"false" json:"description"`
	FreeformTags             map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags              map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ConfigSourceProviderType string                            `json:"configSourceProviderType"`
}

// UnmarshalJSON unmarshals json
func (m *createconfigurationsourceproviderdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateconfigurationsourceproviderdetails createconfigurationsourceproviderdetails
	s := struct {
		Model Unmarshalercreateconfigurationsourceproviderdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConfigSourceProviderType = s.Model.ConfigSourceProviderType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconfigurationsourceproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceProviderType {
	case "GITLAB_ACCESS_TOKEN":
		mm := CreateGitlabAccessTokenConfigurationSourceProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB_ACCESS_TOKEN":
		mm := CreateGithubAccessTokenConfigurationSourceProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCompartmentId returns CompartmentId
func (m createconfigurationsourceproviderdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m createconfigurationsourceproviderdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m createconfigurationsourceproviderdetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m createconfigurationsourceproviderdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createconfigurationsourceproviderdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createconfigurationsourceproviderdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createconfigurationsourceproviderdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
