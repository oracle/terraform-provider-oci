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

// CreateDatabaseToolsMcpServerDefaultDetails Details for the new Database Tools MCP server.
type CreateDatabaseToolsMcpServerDefaultDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A meaningful, human-readable label displayed to end users. Not required to be unique and can be changed after creation. Do not include confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	DatabaseToolsConnectionId *string `mandatory:"true" json:"databaseToolsConnectionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated identity domain.
	DomainId *string `mandatory:"true" json:"domainId"`

	Storage DatabaseToolsMcpServerStorage `mandatory:"true" json:"storage"`

	// Custom Roles associated with the MCP Server.
	CustomRoles []DatabaseToolsMcpServerCustomRole `mandatory:"false" json:"customRoles"`

	// A human-readable description of the Database Tools MCP server.
	Description *string `mandatory:"false" json:"description"`

	// Access token expiry in seconds
	AccessTokenExpiryInSeconds *int `mandatory:"false" json:"accessTokenExpiryInSeconds"`

	// Refresh token expiry in seconds
	RefreshTokenExpiryInSeconds *int `mandatory:"false" json:"refreshTokenExpiryInSeconds"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Specifies the identity used when accessing OCI resources at runtime. AUTHENTICATED_PRINCIPAL to use the caller’s identity (On-Behalf-Of token), or RESOURCE_PRINCIPAL to use the MCP Server’s resource principal (RPST).
	RuntimeIdentity DatabaseToolsMcpServerRuntimeIdentityEnum `mandatory:"false" json:"runtimeIdentity,omitempty"`
}

// GetCompartmentId returns CompartmentId
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetCustomRoles returns CustomRoles
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetCustomRoles() []DatabaseToolsMcpServerCustomRole {
	return m.CustomRoles
}

// GetDisplayName returns DisplayName
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetDescription() *string {
	return m.Description
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetAccessTokenExpiryInSeconds returns AccessTokenExpiryInSeconds
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetAccessTokenExpiryInSeconds() *int {
	return m.AccessTokenExpiryInSeconds
}

// GetRefreshTokenExpiryInSeconds returns RefreshTokenExpiryInSeconds
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetRefreshTokenExpiryInSeconds() *int {
	return m.RefreshTokenExpiryInSeconds
}

// GetDefinedTags returns DefinedTags
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetRuntimeIdentity returns RuntimeIdentity
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetRuntimeIdentity() DatabaseToolsMcpServerRuntimeIdentityEnum {
	return m.RuntimeIdentity
}

// GetLocks returns Locks
func (m CreateDatabaseToolsMcpServerDefaultDetails) GetLocks() []ResourceLock {
	return m.Locks
}

func (m CreateDatabaseToolsMcpServerDefaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsMcpServerDefaultDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsMcpServerRuntimeIdentityEnum(string(m.RuntimeIdentity)); !ok && m.RuntimeIdentity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeIdentity: %s. Supported values are: %s.", m.RuntimeIdentity, strings.Join(GetDatabaseToolsMcpServerRuntimeIdentityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseToolsMcpServerDefaultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseToolsMcpServerDefaultDetails CreateDatabaseToolsMcpServerDefaultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateDatabaseToolsMcpServerDefaultDetails
	}{
		"DEFAULT",
		(MarshalTypeCreateDatabaseToolsMcpServerDefaultDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDatabaseToolsMcpServerDefaultDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CustomRoles                 []DatabaseToolsMcpServerCustomRole        `json:"customRoles"`
		Description                 *string                                   `json:"description"`
		AccessTokenExpiryInSeconds  *int                                      `json:"accessTokenExpiryInSeconds"`
		RefreshTokenExpiryInSeconds *int                                      `json:"refreshTokenExpiryInSeconds"`
		DefinedTags                 map[string]map[string]interface{}         `json:"definedTags"`
		FreeformTags                map[string]string                         `json:"freeformTags"`
		RuntimeIdentity             DatabaseToolsMcpServerRuntimeIdentityEnum `json:"runtimeIdentity"`
		Locks                       []ResourceLock                            `json:"locks"`
		CompartmentId               *string                                   `json:"compartmentId"`
		DisplayName                 *string                                   `json:"displayName"`
		DatabaseToolsConnectionId   *string                                   `json:"databaseToolsConnectionId"`
		DomainId                    *string                                   `json:"domainId"`
		Storage                     databasetoolsmcpserverstorage             `json:"storage"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CustomRoles = make([]DatabaseToolsMcpServerCustomRole, len(model.CustomRoles))
	copy(m.CustomRoles, model.CustomRoles)
	m.Description = model.Description

	m.AccessTokenExpiryInSeconds = model.AccessTokenExpiryInSeconds

	m.RefreshTokenExpiryInSeconds = model.RefreshTokenExpiryInSeconds

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.RuntimeIdentity = model.RuntimeIdentity

	m.Locks = make([]ResourceLock, len(model.Locks))
	copy(m.Locks, model.Locks)
	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.DatabaseToolsConnectionId = model.DatabaseToolsConnectionId

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
