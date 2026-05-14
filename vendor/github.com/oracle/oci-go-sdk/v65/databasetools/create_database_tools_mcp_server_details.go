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

// CreateDatabaseToolsMcpServerDetails Details for the new Database Tools MCP server.
type CreateDatabaseToolsMcpServerDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
	GetCompartmentId() *string

	// A meaningful, human-readable label displayed to end users. Not required to be unique and can be changed after creation. Do not include confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	GetDatabaseToolsConnectionId() *string

	// Custom Roles associated with the MCP Server.
	GetCustomRoles() []DatabaseToolsMcpServerCustomRole

	// A human-readable description of the Database Tools MCP server.
	GetDescription() *string

	// Access token expiry in seconds
	GetAccessTokenExpiryInSeconds() *int

	// Refresh token expiry in seconds
	GetRefreshTokenExpiryInSeconds() *int

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Specifies the identity used when accessing OCI resources at runtime. AUTHENTICATED_PRINCIPAL to use the caller’s identity (On-Behalf-Of token), or RESOURCE_PRINCIPAL to use the MCP Server’s resource principal (RPST).
	GetRuntimeIdentity() DatabaseToolsMcpServerRuntimeIdentityEnum

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type createdatabasetoolsmcpserverdetails struct {
	JsonData                    []byte
	CustomRoles                 []DatabaseToolsMcpServerCustomRole        `mandatory:"false" json:"customRoles"`
	Description                 *string                                   `mandatory:"false" json:"description"`
	AccessTokenExpiryInSeconds  *int                                      `mandatory:"false" json:"accessTokenExpiryInSeconds"`
	RefreshTokenExpiryInSeconds *int                                      `mandatory:"false" json:"refreshTokenExpiryInSeconds"`
	DefinedTags                 map[string]map[string]interface{}         `mandatory:"false" json:"definedTags"`
	FreeformTags                map[string]string                         `mandatory:"false" json:"freeformTags"`
	RuntimeIdentity             DatabaseToolsMcpServerRuntimeIdentityEnum `mandatory:"false" json:"runtimeIdentity,omitempty"`
	Locks                       []ResourceLock                            `mandatory:"false" json:"locks"`
	CompartmentId               *string                                   `mandatory:"true" json:"compartmentId"`
	DisplayName                 *string                                   `mandatory:"true" json:"displayName"`
	DatabaseToolsConnectionId   *string                                   `mandatory:"true" json:"databaseToolsConnectionId"`
	Type                        string                                    `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsmcpserverdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsmcpserverdetails createdatabasetoolsmcpserverdetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsmcpserverdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.DatabaseToolsConnectionId = s.Model.DatabaseToolsConnectionId
	m.CustomRoles = s.Model.CustomRoles
	m.Description = s.Model.Description
	m.AccessTokenExpiryInSeconds = s.Model.AccessTokenExpiryInSeconds
	m.RefreshTokenExpiryInSeconds = s.Model.RefreshTokenExpiryInSeconds
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.RuntimeIdentity = s.Model.RuntimeIdentity
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsmcpserverdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := CreateDatabaseToolsMcpServerDefaultDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDatabaseToolsMcpServerDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetCustomRoles returns CustomRoles
func (m createdatabasetoolsmcpserverdetails) GetCustomRoles() []DatabaseToolsMcpServerCustomRole {
	return m.CustomRoles
}

// GetDescription returns Description
func (m createdatabasetoolsmcpserverdetails) GetDescription() *string {
	return m.Description
}

// GetAccessTokenExpiryInSeconds returns AccessTokenExpiryInSeconds
func (m createdatabasetoolsmcpserverdetails) GetAccessTokenExpiryInSeconds() *int {
	return m.AccessTokenExpiryInSeconds
}

// GetRefreshTokenExpiryInSeconds returns RefreshTokenExpiryInSeconds
func (m createdatabasetoolsmcpserverdetails) GetRefreshTokenExpiryInSeconds() *int {
	return m.RefreshTokenExpiryInSeconds
}

// GetDefinedTags returns DefinedTags
func (m createdatabasetoolsmcpserverdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m createdatabasetoolsmcpserverdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetRuntimeIdentity returns RuntimeIdentity
func (m createdatabasetoolsmcpserverdetails) GetRuntimeIdentity() DatabaseToolsMcpServerRuntimeIdentityEnum {
	return m.RuntimeIdentity
}

// GetLocks returns Locks
func (m createdatabasetoolsmcpserverdetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetCompartmentId returns CompartmentId
func (m createdatabasetoolsmcpserverdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolsmcpserverdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m createdatabasetoolsmcpserverdetails) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

func (m createdatabasetoolsmcpserverdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolsmcpserverdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsMcpServerRuntimeIdentityEnum(string(m.RuntimeIdentity)); !ok && m.RuntimeIdentity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuntimeIdentity: %s. Supported values are: %s.", m.RuntimeIdentity, strings.Join(GetDatabaseToolsMcpServerRuntimeIdentityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
