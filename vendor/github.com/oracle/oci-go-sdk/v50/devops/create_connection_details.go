// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateConnectionDetails The details for creating a connection.
type CreateConnectionDetails interface {

	// Project Identifier
	GetProjectId() *string

	// Optional description about the Connection
	GetDescription() *string

	// Optional Connection display name
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createconnectiondetails struct {
	JsonData       []byte
	ProjectId      *string                           `mandatory:"true" json:"projectId"`
	Description    *string                           `mandatory:"false" json:"description"`
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
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
	case "GITHUB_ACCESS_TOKEN":
		mm := CreateGithubAccessTokenConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_ACCESS_TOKEN":
		mm := CreateGitlabAccessTokenConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetProjectId returns ProjectId
func (m createconnectiondetails) GetProjectId() *string {
	return m.ProjectId
}

//GetDescription returns Description
func (m createconnectiondetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m createconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m createconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createconnectiondetails) String() string {
	return common.PointerString(m)
}
