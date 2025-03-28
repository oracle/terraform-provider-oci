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

// CreateConnectionDetails The details for creating a connection.
type CreateConnectionDetails interface {

	// The OCID of the DevOps project.
	GetProjectId() *string

	// Optional description about the connection.
	GetDescription() *string

	// Optional connection display name. Avoid entering confidential information.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createconnectiondetails struct {
	JsonData       []byte
	Description    *string                           `mandatory:"false" json:"description"`
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ProjectId      *string                           `mandatory:"true" json:"projectId"`
	ConnectionType string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *createconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateconnectiondetails createconnectiondetails
	s := struct {
		Model Unmarshalercreateconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ProjectId = s.Model.ProjectId
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "VBS_ACCESS_TOKEN":
		mm := CreateVbsAccessTokenConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_SERVER_ACCESS_TOKEN":
		mm := CreateGitlabServerAccessTokenConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_SERVER_ACCESS_TOKEN":
		mm := CreateBitbucketServerAccessTokenConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB_ACCESS_TOKEN":
		mm := CreateGithubAccessTokenConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_CLOUD_APP_PASSWORD":
		mm := CreateBitbucketCloudAppPasswordConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_ACCESS_TOKEN":
		mm := CreateGitlabAccessTokenConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateConnectionDetails: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createconnectiondetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m createconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m createconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetProjectId returns ProjectId
func (m createconnectiondetails) GetProjectId() *string {
	return m.ProjectId
}

func (m createconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
