// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateGitlabServerAccessTokenConnectionDetails The details for creating a connection of the type `GITLAB_SERVER_ACCESS_TOKEN`.
// This type corresponds to a connection in GitLab self hosted server that is authenticated with a personal access token.
type CreateGitlabServerAccessTokenConnectionDetails struct {

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of personal access token saved in secret store.
	AccessToken *string `mandatory:"true" json:"accessToken"`

	// The baseUrl of the hosted GitLabServer.
	BaseUrl *string `mandatory:"true" json:"baseUrl"`

	// Optional description about the connection.
	Description *string `mandatory:"false" json:"description"`

	// Optional connection display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	TlsVerifyConfig TlsVerifyConfig `mandatory:"false" json:"tlsVerifyConfig"`
}

// GetDescription returns Description
func (m CreateGitlabServerAccessTokenConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m CreateGitlabServerAccessTokenConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetProjectId returns ProjectId
func (m CreateGitlabServerAccessTokenConnectionDetails) GetProjectId() *string {
	return m.ProjectId
}

// GetFreeformTags returns FreeformTags
func (m CreateGitlabServerAccessTokenConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateGitlabServerAccessTokenConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateGitlabServerAccessTokenConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGitlabServerAccessTokenConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateGitlabServerAccessTokenConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateGitlabServerAccessTokenConnectionDetails CreateGitlabServerAccessTokenConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeCreateGitlabServerAccessTokenConnectionDetails
	}{
		"GITLAB_SERVER_ACCESS_TOKEN",
		(MarshalTypeCreateGitlabServerAccessTokenConnectionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateGitlabServerAccessTokenConnectionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description     *string                           `json:"description"`
		DisplayName     *string                           `json:"displayName"`
		FreeformTags    map[string]string                 `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{} `json:"definedTags"`
		TlsVerifyConfig tlsverifyconfig                   `json:"tlsVerifyConfig"`
		ProjectId       *string                           `json:"projectId"`
		AccessToken     *string                           `json:"accessToken"`
		BaseUrl         *string                           `json:"baseUrl"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.TlsVerifyConfig.UnmarshalPolymorphicJSON(model.TlsVerifyConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TlsVerifyConfig = nn.(TlsVerifyConfig)
	} else {
		m.TlsVerifyConfig = nil
	}

	m.ProjectId = model.ProjectId

	m.AccessToken = model.AccessToken

	m.BaseUrl = model.BaseUrl

	return
}
