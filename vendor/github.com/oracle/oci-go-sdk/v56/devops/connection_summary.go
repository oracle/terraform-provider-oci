// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ConnectionSummary Summary information for a connection.
type ConnectionSummary interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of the compartment containing the connection.
	GetCompartmentId() *string

	// The OCID of the DevOps project.
	GetProjectId() *string

	// Connection display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	// Optional description about the connection.
	GetDescription() *string

	// The time the connection was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeCreated() *common.SDKTime

	// The time the connection was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeUpdated() *common.SDKTime

	// The current state of the connection.
	GetLifecycleState() ConnectionLifecycleStateEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type connectionsummary struct {
	JsonData       []byte
	Id             *string                           `mandatory:"true" json:"id"`
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	ProjectId      *string                           `mandatory:"true" json:"projectId"`
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	Description    *string                           `mandatory:"false" json:"description"`
	TimeCreated    *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated    *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState ConnectionLifecycleStateEnum      `mandatory:"false" json:"lifecycleState,omitempty"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags     map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	ConnectionType string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *connectionsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectionsummary connectionsummary
	s := struct {
		Model Unmarshalerconnectionsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.ProjectId = s.Model.ProjectId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectionsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "GITHUB_ACCESS_TOKEN":
		mm := GithubAccessTokenConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_ACCESS_TOKEN":
		mm := GitlabAccessTokenConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m connectionsummary) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m connectionsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetProjectId returns ProjectId
func (m connectionsummary) GetProjectId() *string {
	return m.ProjectId
}

//GetDisplayName returns DisplayName
func (m connectionsummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m connectionsummary) GetDescription() *string {
	return m.Description
}

//GetTimeCreated returns TimeCreated
func (m connectionsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m connectionsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m connectionsummary) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m connectionsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m connectionsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m connectionsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m connectionsummary) String() string {
	return common.PointerString(m)
}
