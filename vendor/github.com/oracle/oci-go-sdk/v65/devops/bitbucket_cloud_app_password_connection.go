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

// BitbucketCloudAppPasswordConnection The properties that define a connection of the type `BITBUCKET_CLOUD_APP_PASSWORD`.
// This type corresponds to a connection in Bitbucket Cloud that is authenticated with a App Password along with username.
type BitbucketCloudAppPasswordConnection struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Public Bitbucket Cloud Username in plain text
	Username *string `mandatory:"true" json:"username"`

	// OCID of personal Bitbucket Cloud AppPassword saved in secret store
	AppPassword *string `mandatory:"true" json:"appPassword"`

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
func (m BitbucketCloudAppPasswordConnection) GetId() *string {
	return m.Id
}

// GetDescription returns Description
func (m BitbucketCloudAppPasswordConnection) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m BitbucketCloudAppPasswordConnection) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m BitbucketCloudAppPasswordConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetProjectId returns ProjectId
func (m BitbucketCloudAppPasswordConnection) GetProjectId() *string {
	return m.ProjectId
}

// GetTimeCreated returns TimeCreated
func (m BitbucketCloudAppPasswordConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m BitbucketCloudAppPasswordConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLastConnectionValidationResult returns LastConnectionValidationResult
func (m BitbucketCloudAppPasswordConnection) GetLastConnectionValidationResult() *ConnectionValidationResult {
	return m.LastConnectionValidationResult
}

// GetLifecycleDetails returns LifecycleDetails
func (m BitbucketCloudAppPasswordConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetLifecycleState returns LifecycleState
func (m BitbucketCloudAppPasswordConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m BitbucketCloudAppPasswordConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m BitbucketCloudAppPasswordConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m BitbucketCloudAppPasswordConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m BitbucketCloudAppPasswordConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BitbucketCloudAppPasswordConnection) ValidateEnumValue() (bool, error) {
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
func (m BitbucketCloudAppPasswordConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBitbucketCloudAppPasswordConnection BitbucketCloudAppPasswordConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeBitbucketCloudAppPasswordConnection
	}{
		"BITBUCKET_CLOUD_APP_PASSWORD",
		(MarshalTypeBitbucketCloudAppPasswordConnection)(m),
	}

	return json.Marshal(&s)
}
