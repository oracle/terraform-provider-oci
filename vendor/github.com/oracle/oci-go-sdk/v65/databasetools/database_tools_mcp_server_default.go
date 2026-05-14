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

// DatabaseToolsMcpServerDefault Allows the creation, configuration and management of an MCP server.
type DatabaseToolsMcpServerDefault struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools MCP server.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A meaningful, human-readable label displayed to end users. Not required to be unique and can be changed after creation. Do not include confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	DatabaseToolsConnectionId *string `mandatory:"true" json:"databaseToolsConnectionId"`

	RelatedResource *DatabaseToolsMcpServerRelatedResource `mandatory:"true" json:"relatedResource"`

	// The time the Database Tools MCP server was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Database Tools MCP server was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated identity domain.
	DomainId *string `mandatory:"true" json:"domainId"`

	Storage DatabaseToolsMcpServerStorage `mandatory:"true" json:"storage"`

	// A human-readable description of the Database Tools MCP server.
	Description *string `mandatory:"false" json:"description"`

	// Invoke endpoints for the MCP server.
	Endpoints []DatabaseToolsMcpServerEndpoint `mandatory:"false" json:"endpoints"`

	// Built-in roles associated with the MCP Server.
	BuiltInRoles []DatabaseToolsMcpServerBuiltInRole `mandatory:"false" json:"builtInRoles"`

	// Custom roles associated with the MCP Server.
	CustomRoles []DatabaseToolsMcpServerCustomRole `mandatory:"false" json:"customRoles"`

	// Access token expiry in seconds
	AccessTokenExpiryInSeconds *int `mandatory:"false" json:"accessTokenExpiryInSeconds"`

	// Refresh token expiry in seconds
	RefreshTokenExpiryInSeconds *int `mandatory:"false" json:"refreshTokenExpiryInSeconds"`

	// A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated domain application (Oracle Cloud Service).
	DomainAppId *string `mandatory:"false" json:"domainAppId"`

	// The current state of the Database Tools MCP server.
	LifecycleState DatabaseToolsMcpServerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Specifies the identity used when accessing OCI resources at runtime. AUTHENTICATED_PRINCIPAL to use the caller’s identity (On-Behalf-Of token), or RESOURCE_PRINCIPAL to use the MCP Server’s resource principal (RPST).
	RuntimeIdentity DatabaseToolsMcpServerRuntimeIdentityEnum `mandatory:"true" json:"runtimeIdentity"`
}

// GetId returns Id
func (m DatabaseToolsMcpServerDefault) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m DatabaseToolsMcpServerDefault) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsMcpServerDefault) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m DatabaseToolsMcpServerDefault) GetDescription() *string {
	return m.Description
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m DatabaseToolsMcpServerDefault) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetEndpoints returns Endpoints
func (m DatabaseToolsMcpServerDefault) GetEndpoints() []DatabaseToolsMcpServerEndpoint {
	return m.Endpoints
}

// GetBuiltInRoles returns BuiltInRoles
func (m DatabaseToolsMcpServerDefault) GetBuiltInRoles() []DatabaseToolsMcpServerBuiltInRole {
	return m.BuiltInRoles
}

// GetCustomRoles returns CustomRoles
func (m DatabaseToolsMcpServerDefault) GetCustomRoles() []DatabaseToolsMcpServerCustomRole {
	return m.CustomRoles
}

// GetAccessTokenExpiryInSeconds returns AccessTokenExpiryInSeconds
func (m DatabaseToolsMcpServerDefault) GetAccessTokenExpiryInSeconds() *int {
	return m.AccessTokenExpiryInSeconds
}

// GetRefreshTokenExpiryInSeconds returns RefreshTokenExpiryInSeconds
func (m DatabaseToolsMcpServerDefault) GetRefreshTokenExpiryInSeconds() *int {
	return m.RefreshTokenExpiryInSeconds
}

// GetRelatedResource returns RelatedResource
func (m DatabaseToolsMcpServerDefault) GetRelatedResource() *DatabaseToolsMcpServerRelatedResource {
	return m.RelatedResource
}

// GetLifecycleState returns LifecycleState
func (m DatabaseToolsMcpServerDefault) GetLifecycleState() DatabaseToolsMcpServerLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m DatabaseToolsMcpServerDefault) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsMcpServerDefault) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsMcpServerDefault) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetDefinedTags returns DefinedTags
func (m DatabaseToolsMcpServerDefault) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m DatabaseToolsMcpServerDefault) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m DatabaseToolsMcpServerDefault) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetRuntimeIdentity returns RuntimeIdentity
func (m DatabaseToolsMcpServerDefault) GetRuntimeIdentity() DatabaseToolsMcpServerRuntimeIdentityEnum {
	return m.RuntimeIdentity
}

// GetLocks returns Locks
func (m DatabaseToolsMcpServerDefault) GetLocks() []ResourceLock {
	return m.Locks
}

func (m DatabaseToolsMcpServerDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsMcpServerDefault) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m DatabaseToolsMcpServerDefault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsMcpServerDefault DatabaseToolsMcpServerDefault
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsMcpServerDefault
	}{
		"DEFAULT",
		(MarshalTypeDatabaseToolsMcpServerDefault)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsMcpServerDefault) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                 *string                                   `json:"description"`
		Endpoints                   []DatabaseToolsMcpServerEndpoint          `json:"endpoints"`
		BuiltInRoles                []DatabaseToolsMcpServerBuiltInRole       `json:"builtInRoles"`
		CustomRoles                 []DatabaseToolsMcpServerCustomRole        `json:"customRoles"`
		AccessTokenExpiryInSeconds  *int                                      `json:"accessTokenExpiryInSeconds"`
		RefreshTokenExpiryInSeconds *int                                      `json:"refreshTokenExpiryInSeconds"`
		LifecycleDetails            *string                                   `json:"lifecycleDetails"`
		DefinedTags                 map[string]map[string]interface{}         `json:"definedTags"`
		FreeformTags                map[string]string                         `json:"freeformTags"`
		SystemTags                  map[string]map[string]interface{}         `json:"systemTags"`
		Locks                       []ResourceLock                            `json:"locks"`
		DomainAppId                 *string                                   `json:"domainAppId"`
		Id                          *string                                   `json:"id"`
		CompartmentId               *string                                   `json:"compartmentId"`
		DisplayName                 *string                                   `json:"displayName"`
		DatabaseToolsConnectionId   *string                                   `json:"databaseToolsConnectionId"`
		RelatedResource             *DatabaseToolsMcpServerRelatedResource    `json:"relatedResource"`
		LifecycleState              DatabaseToolsMcpServerLifecycleStateEnum  `json:"lifecycleState"`
		TimeCreated                 *common.SDKTime                           `json:"timeCreated"`
		TimeUpdated                 *common.SDKTime                           `json:"timeUpdated"`
		RuntimeIdentity             DatabaseToolsMcpServerRuntimeIdentityEnum `json:"runtimeIdentity"`
		DomainId                    *string                                   `json:"domainId"`
		Storage                     databasetoolsmcpserverstorage             `json:"storage"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Endpoints = make([]DatabaseToolsMcpServerEndpoint, len(model.Endpoints))
	copy(m.Endpoints, model.Endpoints)
	m.BuiltInRoles = make([]DatabaseToolsMcpServerBuiltInRole, len(model.BuiltInRoles))
	copy(m.BuiltInRoles, model.BuiltInRoles)
	m.CustomRoles = make([]DatabaseToolsMcpServerCustomRole, len(model.CustomRoles))
	copy(m.CustomRoles, model.CustomRoles)
	m.AccessTokenExpiryInSeconds = model.AccessTokenExpiryInSeconds

	m.RefreshTokenExpiryInSeconds = model.RefreshTokenExpiryInSeconds

	m.LifecycleDetails = model.LifecycleDetails

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.SystemTags = model.SystemTags

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.DomainAppId = model.DomainAppId

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.DatabaseToolsConnectionId = model.DatabaseToolsConnectionId

	m.RelatedResource = model.RelatedResource

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.RuntimeIdentity = model.RuntimeIdentity

	m.DomainId = model.DomainId

	nn, e = model.Storage.UnmarshalPolymorphicJSON(model.Storage.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Storage = nn.(DatabaseToolsMcpServerStorage)
	} else {
		m.Storage = nil
	}

	return
}
