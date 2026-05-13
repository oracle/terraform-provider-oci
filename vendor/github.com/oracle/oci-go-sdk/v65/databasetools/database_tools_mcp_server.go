// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsMcpServer Allows the creation, configuration and management of an MCP server.
type DatabaseToolsMcpServer interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools MCP server.
	GetId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
	GetCompartmentId() *string

	// A meaningful, human-readable label displayed to end users. Not required to be unique and can be changed after creation. Do not include confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	GetDatabaseToolsConnectionId() *string

	GetRelatedResource() *DatabaseToolsMcpServerRelatedResource

	// The current state of the Database Tools MCP server.
	GetLifecycleState() DatabaseToolsMcpServerLifecycleStateEnum

	// The time the Database Tools MCP server was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the Database Tools MCP server was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

	// Specifies the identity used when accessing OCI resources at runtime. AUTHENTICATED_PRINCIPAL to use the caller’s identity (On-Behalf-Of token), or RESOURCE_PRINCIPAL to use the MCP Server’s resource principal (RPST).
	GetRuntimeIdentity() DatabaseToolsMcpServerRuntimeIdentityEnum

	// A human-readable description of the Database Tools MCP server.
	GetDescription() *string

	// Invoke endpoints for the MCP server.
	GetEndpoints() []DatabaseToolsMcpServerEndpoint

	// Built-in roles associated with the MCP Server.
	GetBuiltInRoles() []DatabaseToolsMcpServerBuiltInRole

	// Custom roles associated with the MCP Server.
	GetCustomRoles() []DatabaseToolsMcpServerCustomRole

	// Access token expiry in seconds
	GetAccessTokenExpiryInSeconds() *int

	// Refresh token expiry in seconds
	GetRefreshTokenExpiryInSeconds() *int

	// A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
	GetLifecycleDetails() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type databasetoolsmcpserver struct {
	JsonData                    []byte
	Description                 *string                                   `mandatory:"false" json:"description"`
	Endpoints                   []DatabaseToolsMcpServerEndpoint          `mandatory:"false" json:"endpoints"`
	BuiltInRoles                []DatabaseToolsMcpServerBuiltInRole       `mandatory:"false" json:"builtInRoles"`
	CustomRoles                 []DatabaseToolsMcpServerCustomRole        `mandatory:"false" json:"customRoles"`
	AccessTokenExpiryInSeconds  *int                                      `mandatory:"false" json:"accessTokenExpiryInSeconds"`
	RefreshTokenExpiryInSeconds *int                                      `mandatory:"false" json:"refreshTokenExpiryInSeconds"`
	LifecycleDetails            *string                                   `mandatory:"false" json:"lifecycleDetails"`
	DefinedTags                 map[string]map[string]interface{}         `mandatory:"false" json:"definedTags"`
	FreeformTags                map[string]string                         `mandatory:"false" json:"freeformTags"`
	SystemTags                  map[string]map[string]interface{}         `mandatory:"false" json:"systemTags"`
	Locks                       []ResourceLock                            `mandatory:"false" json:"locks"`
	Id                          *string                                   `mandatory:"true" json:"id"`
	CompartmentId               *string                                   `mandatory:"true" json:"compartmentId"`
	DisplayName                 *string                                   `mandatory:"true" json:"displayName"`
	DatabaseToolsConnectionId   *string                                   `mandatory:"true" json:"databaseToolsConnectionId"`
	RelatedResource             *DatabaseToolsMcpServerRelatedResource    `mandatory:"true" json:"relatedResource"`
	LifecycleState              DatabaseToolsMcpServerLifecycleStateEnum  `mandatory:"true" json:"lifecycleState"`
	TimeCreated                 *common.SDKTime                           `mandatory:"true" json:"timeCreated"`
	TimeUpdated                 *common.SDKTime                           `mandatory:"true" json:"timeUpdated"`
	RuntimeIdentity             DatabaseToolsMcpServerRuntimeIdentityEnum `mandatory:"true" json:"runtimeIdentity"`
	Type                        string                                    `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsmcpserver) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsmcpserver databasetoolsmcpserver
	s := struct {
		Model Unmarshalerdatabasetoolsmcpserver
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.DatabaseToolsConnectionId = s.Model.DatabaseToolsConnectionId
	m.RelatedResource = s.Model.RelatedResource
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.RuntimeIdentity = s.Model.RuntimeIdentity
	m.Description = s.Model.Description
	m.Endpoints = s.Model.Endpoints
	m.BuiltInRoles = s.Model.BuiltInRoles
	m.CustomRoles = s.Model.CustomRoles
	m.AccessTokenExpiryInSeconds = s.Model.AccessTokenExpiryInSeconds
	m.RefreshTokenExpiryInSeconds = s.Model.RefreshTokenExpiryInSeconds
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.SystemTags = s.Model.SystemTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsmcpserver) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DatabaseToolsMcpServerDefault{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsMcpServer: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m databasetoolsmcpserver) GetDescription() *string {
	return m.Description
}

// GetEndpoints returns Endpoints
func (m databasetoolsmcpserver) GetEndpoints() []DatabaseToolsMcpServerEndpoint {
	return m.Endpoints
}

// GetBuiltInRoles returns BuiltInRoles
func (m databasetoolsmcpserver) GetBuiltInRoles() []DatabaseToolsMcpServerBuiltInRole {
	return m.BuiltInRoles
}

// GetCustomRoles returns CustomRoles
func (m databasetoolsmcpserver) GetCustomRoles() []DatabaseToolsMcpServerCustomRole {
	return m.CustomRoles
}

// GetAccessTokenExpiryInSeconds returns AccessTokenExpiryInSeconds
func (m databasetoolsmcpserver) GetAccessTokenExpiryInSeconds() *int {
	return m.AccessTokenExpiryInSeconds
}

// GetRefreshTokenExpiryInSeconds returns RefreshTokenExpiryInSeconds
func (m databasetoolsmcpserver) GetRefreshTokenExpiryInSeconds() *int {
	return m.RefreshTokenExpiryInSeconds
}

// GetLifecycleDetails returns LifecycleDetails
func (m databasetoolsmcpserver) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDefinedTags returns DefinedTags
func (m databasetoolsmcpserver) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m databasetoolsmcpserver) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m databasetoolsmcpserver) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m databasetoolsmcpserver) GetLocks() []ResourceLock {
	return m.Locks
}

// GetId returns Id
func (m databasetoolsmcpserver) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m databasetoolsmcpserver) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m databasetoolsmcpserver) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m databasetoolsmcpserver) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetRelatedResource returns RelatedResource
func (m databasetoolsmcpserver) GetRelatedResource() *DatabaseToolsMcpServerRelatedResource {
	return m.RelatedResource
}

// GetLifecycleState returns LifecycleState
func (m databasetoolsmcpserver) GetLifecycleState() DatabaseToolsMcpServerLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m databasetoolsmcpserver) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolsmcpserver) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetRuntimeIdentity returns RuntimeIdentity
func (m databasetoolsmcpserver) GetRuntimeIdentity() DatabaseToolsMcpServerRuntimeIdentityEnum {
	return m.RuntimeIdentity
}

func (m databasetoolsmcpserver) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsmcpserver) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolsMcpServerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsMcpServerLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsMcpServerRuntimeIdentityEnum(string(m.RuntimeIdentity)); !ok && m.RuntimeIdentity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeIdentity: %s. Supported values are: %s.", m.RuntimeIdentity, strings.Join(GetDatabaseToolsMcpServerRuntimeIdentityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
