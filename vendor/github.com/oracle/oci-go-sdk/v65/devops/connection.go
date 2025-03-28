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

// Connection The properties that define a connection to external repositories.
type Connection interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of the compartment containing the connection.
	GetCompartmentId() *string

	// The OCID of the DevOps project.
	GetProjectId() *string

	// Optional description about the connection.
	GetDescription() *string

	// Connection display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

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

type connection struct {
	JsonData                       []byte
	Description                    *string                           `mandatory:"false" json:"description"`
	DisplayName                    *string                           `mandatory:"false" json:"displayName"`
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
func (m *connection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnection connection
	s := struct {
		Model Unmarshalerconnection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.ProjectId = s.Model.ProjectId
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
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
func (m *connection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "BITBUCKET_SERVER_ACCESS_TOKEN":
		mm := BitbucketServerAccessTokenConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_ACCESS_TOKEN":
		mm := GitlabAccessTokenConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB_ACCESS_TOKEN":
		mm := GithubAccessTokenConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BITBUCKET_CLOUD_APP_PASSWORD":
		mm := BitbucketCloudAppPasswordConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB_SERVER_ACCESS_TOKEN":
		mm := GitlabServerAccessTokenConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VBS_ACCESS_TOKEN":
		mm := VbsAccessTokenConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Connection: %s.", m.ConnectionType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m connection) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m connection) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m connection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m connection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLastConnectionValidationResult returns LastConnectionValidationResult
func (m connection) GetLastConnectionValidationResult() *ConnectionValidationResult {
	return m.LastConnectionValidationResult
}

// GetLifecycleDetails returns LifecycleDetails
func (m connection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetLifecycleState returns LifecycleState
func (m connection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m connection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m connection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m connection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m connection) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m connection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetProjectId returns ProjectId
func (m connection) GetProjectId() *string {
	return m.ProjectId
}

func (m connection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionLifecycleStateEnum Enum with underlying type: string
type ConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for ConnectionLifecycleStateEnum
const (
	ConnectionLifecycleStateActive   ConnectionLifecycleStateEnum = "ACTIVE"
	ConnectionLifecycleStateDeleting ConnectionLifecycleStateEnum = "DELETING"
)

var mappingConnectionLifecycleStateEnum = map[string]ConnectionLifecycleStateEnum{
	"ACTIVE":   ConnectionLifecycleStateActive,
	"DELETING": ConnectionLifecycleStateDeleting,
}

var mappingConnectionLifecycleStateEnumLowerCase = map[string]ConnectionLifecycleStateEnum{
	"active":   ConnectionLifecycleStateActive,
	"deleting": ConnectionLifecycleStateDeleting,
}

// GetConnectionLifecycleStateEnumValues Enumerates the set of values for ConnectionLifecycleStateEnum
func GetConnectionLifecycleStateEnumValues() []ConnectionLifecycleStateEnum {
	values := make([]ConnectionLifecycleStateEnum, 0)
	for _, v := range mappingConnectionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionLifecycleStateEnumStringValues Enumerates the set of values in String for ConnectionLifecycleStateEnum
func GetConnectionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETING",
	}
}

// GetMappingConnectionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionLifecycleStateEnum(val string) (ConnectionLifecycleStateEnum, bool) {
	enum, ok := mappingConnectionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConnectionConnectionTypeEnum Enum with underlying type: string
type ConnectionConnectionTypeEnum string

// Set of constants representing the allowable values for ConnectionConnectionTypeEnum
const (
	ConnectionConnectionTypeGithubAccessToken          ConnectionConnectionTypeEnum = "GITHUB_ACCESS_TOKEN"
	ConnectionConnectionTypeGitlabAccessToken          ConnectionConnectionTypeEnum = "GITLAB_ACCESS_TOKEN"
	ConnectionConnectionTypeGitlabServerAccessToken    ConnectionConnectionTypeEnum = "GITLAB_SERVER_ACCESS_TOKEN"
	ConnectionConnectionTypeBitbucketServerAccessToken ConnectionConnectionTypeEnum = "BITBUCKET_SERVER_ACCESS_TOKEN"
	ConnectionConnectionTypeBitbucketCloudAppPassword  ConnectionConnectionTypeEnum = "BITBUCKET_CLOUD_APP_PASSWORD"
	ConnectionConnectionTypeVbsAccessToken             ConnectionConnectionTypeEnum = "VBS_ACCESS_TOKEN"
)

var mappingConnectionConnectionTypeEnum = map[string]ConnectionConnectionTypeEnum{
	"GITHUB_ACCESS_TOKEN":           ConnectionConnectionTypeGithubAccessToken,
	"GITLAB_ACCESS_TOKEN":           ConnectionConnectionTypeGitlabAccessToken,
	"GITLAB_SERVER_ACCESS_TOKEN":    ConnectionConnectionTypeGitlabServerAccessToken,
	"BITBUCKET_SERVER_ACCESS_TOKEN": ConnectionConnectionTypeBitbucketServerAccessToken,
	"BITBUCKET_CLOUD_APP_PASSWORD":  ConnectionConnectionTypeBitbucketCloudAppPassword,
	"VBS_ACCESS_TOKEN":              ConnectionConnectionTypeVbsAccessToken,
}

var mappingConnectionConnectionTypeEnumLowerCase = map[string]ConnectionConnectionTypeEnum{
	"github_access_token":           ConnectionConnectionTypeGithubAccessToken,
	"gitlab_access_token":           ConnectionConnectionTypeGitlabAccessToken,
	"gitlab_server_access_token":    ConnectionConnectionTypeGitlabServerAccessToken,
	"bitbucket_server_access_token": ConnectionConnectionTypeBitbucketServerAccessToken,
	"bitbucket_cloud_app_password":  ConnectionConnectionTypeBitbucketCloudAppPassword,
	"vbs_access_token":              ConnectionConnectionTypeVbsAccessToken,
}

// GetConnectionConnectionTypeEnumValues Enumerates the set of values for ConnectionConnectionTypeEnum
func GetConnectionConnectionTypeEnumValues() []ConnectionConnectionTypeEnum {
	values := make([]ConnectionConnectionTypeEnum, 0)
	for _, v := range mappingConnectionConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionConnectionTypeEnumStringValues Enumerates the set of values in String for ConnectionConnectionTypeEnum
func GetConnectionConnectionTypeEnumStringValues() []string {
	return []string{
		"GITHUB_ACCESS_TOKEN",
		"GITLAB_ACCESS_TOKEN",
		"GITLAB_SERVER_ACCESS_TOKEN",
		"BITBUCKET_SERVER_ACCESS_TOKEN",
		"BITBUCKET_CLOUD_APP_PASSWORD",
		"VBS_ACCESS_TOKEN",
	}
}

// GetMappingConnectionConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionConnectionTypeEnum(val string) (ConnectionConnectionTypeEnum, bool) {
	enum, ok := mappingConnectionConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
