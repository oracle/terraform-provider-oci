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

// BitbucketServerTokenConnectionSummary Summary information for a connection of the type `BITBUCKET_SERVER_ACCESS_TOKEN`.
// This type corresponds to a connection in Bitbucket that is authenticated with a personal access token.
type BitbucketServerTokenConnectionSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of personal access token saved in secret store.
	AccessToken *string `mandatory:"true" json:"accessToken"`

	// The Base URL of the hosted BitbucketServer.
	BaseUrl *string `mandatory:"true" json:"baseUrl"`

	// Connection display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the connection.
	Description *string `mandatory:"false" json:"description"`

	// The time the connection was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the connection was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	LastConnectionValidationResult *ConnectionValidationResult `mandatory:"false" json:"lastConnectionValidationResult"`

	// A detailed message describing the current state. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	TlsVerifyConfig TlsVerifyConfig `mandatory:"false" json:"tlsVerifyConfig"`

	// The current state of the connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m BitbucketServerTokenConnectionSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m BitbucketServerTokenConnectionSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m BitbucketServerTokenConnectionSummary) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m BitbucketServerTokenConnectionSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetProjectId returns ProjectId
func (m BitbucketServerTokenConnectionSummary) GetProjectId() *string {
	return m.ProjectId
}

// GetTimeCreated returns TimeCreated
func (m BitbucketServerTokenConnectionSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m BitbucketServerTokenConnectionSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLastConnectionValidationResult returns LastConnectionValidationResult
func (m BitbucketServerTokenConnectionSummary) GetLastConnectionValidationResult() *ConnectionValidationResult {
	return m.LastConnectionValidationResult
}

// GetLifecycleDetails returns LifecycleDetails
func (m BitbucketServerTokenConnectionSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetLifecycleState returns LifecycleState
func (m BitbucketServerTokenConnectionSummary) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m BitbucketServerTokenConnectionSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m BitbucketServerTokenConnectionSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m BitbucketServerTokenConnectionSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m BitbucketServerTokenConnectionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BitbucketServerTokenConnectionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BitbucketServerTokenConnectionSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBitbucketServerTokenConnectionSummary BitbucketServerTokenConnectionSummary
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeBitbucketServerTokenConnectionSummary
	}{
		"BITBUCKET_SERVER_ACCESS_TOKEN",
		(MarshalTypeBitbucketServerTokenConnectionSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *BitbucketServerTokenConnectionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                    *string                           `json:"displayName"`
		Description                    *string                           `json:"description"`
		TimeCreated                    *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                    *common.SDKTime                   `json:"timeUpdated"`
		LastConnectionValidationResult *ConnectionValidationResult       `json:"lastConnectionValidationResult"`
		LifecycleDetails               *string                           `json:"lifecycleDetails"`
		LifecycleState                 ConnectionLifecycleStateEnum      `json:"lifecycleState"`
		FreeformTags                   map[string]string                 `json:"freeformTags"`
		DefinedTags                    map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                     map[string]map[string]interface{} `json:"systemTags"`
		TlsVerifyConfig                tlsverifyconfig                   `json:"tlsVerifyConfig"`
		Id                             *string                           `json:"id"`
		CompartmentId                  *string                           `json:"compartmentId"`
		ProjectId                      *string                           `json:"projectId"`
		AccessToken                    *string                           `json:"accessToken"`
		BaseUrl                        *string                           `json:"baseUrl"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LastConnectionValidationResult = model.LastConnectionValidationResult

	m.LifecycleDetails = model.LifecycleDetails

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	nn, e = model.TlsVerifyConfig.UnmarshalPolymorphicJSON(model.TlsVerifyConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TlsVerifyConfig = nn.(TlsVerifyConfig)
	} else {
		m.TlsVerifyConfig = nil
	}

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	m.AccessToken = model.AccessToken

	m.BaseUrl = model.BaseUrl

	return
}
