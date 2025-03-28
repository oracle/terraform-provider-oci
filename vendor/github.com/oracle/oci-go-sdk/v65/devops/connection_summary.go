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

	GetLastConnectionValidationResult() *ConnectionValidationResult

	// A detailed message describing the current state. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// The current state of the connection.
	GetLifecycleState() ConnectionLifecycleStateEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type connectionsummary struct {
	JsonData                       []byte
	DisplayName                    *string                           `mandatory:"false" json:"displayName"`
	Description                    *string                           `mandatory:"false" json:"description"`
	TimeCreated                    *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated                    *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LastConnectionValidationResult *ConnectionValidationResult       `mandatory:"false" json:"lastConnectionValidationResult"`
	LifecycleDetails               *string                           `mandatory:"false" json:"lifecycleDetails"`
	LifecycleState                 ConnectionLifecycleStateEnum      `mandatory:"false" json:"lifecycleState,omitempty"`
	FreeformTags                   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags                     map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                             *string                           `mandatory:"true" json:"id"`
	CompartmentId                  *string                           `mandatory:"true" json:"compartmentId"`
	ProjectId                      *string                           `mandatory:"true" json:"projectId"`
	ConnectionType                 string                            `json:"connectionType"`
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
	m.LastConnectionValidationResult = s.Model.LastConnectionValidationResult
	m.LifecycleDetails = s.Model.LifecycleDetails
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
	case "BITBUCKET_CLOUD_APP_PASSWORD":
		mm := BitbucketCloudAppPasswordConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VBS_ACCESS_TOKEN":
		mm := VbsAccessTokenConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB_ACCESS_TOKEN":
		mm := GithubAccessTokenConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_ACCESS_TOKEN":
		mm := GitlabAccessTokenConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_SERVER_ACCESS_TOKEN":
		mm := BitbucketServerTokenConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_SERVER_ACCESS_TOKEN":
		mm := GitlabServerAccessTokenConnectionSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ConnectionSummary: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m connectionsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m connectionsummary) GetDescription() *string {
	return m.Description
}

// GetTimeCreated returns TimeCreated
func (m connectionsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m connectionsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLastConnectionValidationResult returns LastConnectionValidationResult
func (m connectionsummary) GetLastConnectionValidationResult() *ConnectionValidationResult {
	return m.LastConnectionValidationResult
}

// GetLifecycleDetails returns LifecycleDetails
func (m connectionsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetLifecycleState returns LifecycleState
func (m connectionsummary) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m connectionsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m connectionsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m connectionsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m connectionsummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m connectionsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetProjectId returns ProjectId
func (m connectionsummary) GetProjectId() *string {
	return m.ProjectId
}

func (m connectionsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connectionsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
