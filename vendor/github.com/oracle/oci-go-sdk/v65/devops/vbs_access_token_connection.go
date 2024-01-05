// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VbsAccessTokenConnection The properties that define a connection of the type `VBS_ACCESS_TOKEN`.
// This type corresponds to a connection in Visual Builder Studio that is authenticated with a Personal Access Token.
type VbsAccessTokenConnection struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of personal access token saved in secret store.
	AccessToken *string `mandatory:"true" json:"accessToken"`

	// The Base URL of the hosted Visual Builder Studio server.
	BaseUrl *string `mandatory:"true" json:"baseUrl"`

	// Optional description about the connection.
	Description *string `mandatory:"false" json:"description"`

	// Connection display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the connection was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the connection was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	LastConnectionValidationResult *ConnectionValidationResult `mandatory:"false" json:"lastConnectionValidationResult"`

	// A detailed message describing the current state. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m VbsAccessTokenConnection) GetId() *string {
	return m.Id
}

// GetDescription returns Description
func (m VbsAccessTokenConnection) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m VbsAccessTokenConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m VbsAccessTokenConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetProjectId returns ProjectId
func (m VbsAccessTokenConnection) GetProjectId() *string {
	return m.ProjectId
}

// GetTimeCreated returns TimeCreated
func (m VbsAccessTokenConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m VbsAccessTokenConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLastConnectionValidationResult returns LastConnectionValidationResult
func (m VbsAccessTokenConnection) GetLastConnectionValidationResult() *ConnectionValidationResult {
	return m.LastConnectionValidationResult
}

// GetLifecycleDetails returns LifecycleDetails
func (m VbsAccessTokenConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetLifecycleState returns LifecycleState
func (m VbsAccessTokenConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m VbsAccessTokenConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m VbsAccessTokenConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m VbsAccessTokenConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m VbsAccessTokenConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VbsAccessTokenConnection) ValidateEnumValue() (bool, error) {
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
func (m VbsAccessTokenConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVbsAccessTokenConnection VbsAccessTokenConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeVbsAccessTokenConnection
	}{
		"VBS_ACCESS_TOKEN",
		(MarshalTypeVbsAccessTokenConnection)(m),
	}

	return json.Marshal(&s)
}
