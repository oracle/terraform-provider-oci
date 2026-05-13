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

// UpdateDatabaseToolsMcpServerDetailsDefault Database Tools MCP server information to be updated for the default type.
type UpdateDatabaseToolsMcpServerDetailsDefault struct {

	// A meaningful, human-readable label displayed to end users. Not required to be unique and can be changed after creation. Do not include confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A human-readable description of the Database Tools MCP server.
	Description *string `mandatory:"false" json:"description"`

	// Custom roles associated with the MCP Server.
	CustomRoles []DatabaseToolsMcpServerCustomRole `mandatory:"false" json:"customRoles"`

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

	Storage DatabaseToolsMcpServerStorage `mandatory:"false" json:"storage"`
}

// GetDisplayName returns DisplayName
func (m UpdateDatabaseToolsMcpServerDetailsDefault) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateDatabaseToolsMcpServerDetailsDefault) GetDescription() *string {
	return m.Description
}

// GetCustomRoles returns CustomRoles
func (m UpdateDatabaseToolsMcpServerDetailsDefault) GetCustomRoles() []DatabaseToolsMcpServerCustomRole {
	return m.CustomRoles
}

// GetAccessTokenExpiryInSeconds returns AccessTokenExpiryInSeconds
func (m UpdateDatabaseToolsMcpServerDetailsDefault) GetAccessTokenExpiryInSeconds() *int {
	return m.AccessTokenExpiryInSeconds
}

// GetRefreshTokenExpiryInSeconds returns RefreshTokenExpiryInSeconds
func (m UpdateDatabaseToolsMcpServerDetailsDefault) GetRefreshTokenExpiryInSeconds() *int {
	return m.RefreshTokenExpiryInSeconds
}

// GetDefinedTags returns DefinedTags
func (m UpdateDatabaseToolsMcpServerDetailsDefault) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m UpdateDatabaseToolsMcpServerDetailsDefault) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m UpdateDatabaseToolsMcpServerDetailsDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDatabaseToolsMcpServerDetailsDefault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateDatabaseToolsMcpServerDetailsDefault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateDatabaseToolsMcpServerDetailsDefault UpdateDatabaseToolsMcpServerDetailsDefault
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateDatabaseToolsMcpServerDetailsDefault
	}{
		"DEFAULT",
		(MarshalTypeUpdateDatabaseToolsMcpServerDetailsDefault)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDatabaseToolsMcpServerDetailsDefault) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                 *string                            `json:"displayName"`
		Description                 *string                            `json:"description"`
		CustomRoles                 []DatabaseToolsMcpServerCustomRole `json:"customRoles"`
		AccessTokenExpiryInSeconds  *int                               `json:"accessTokenExpiryInSeconds"`
		RefreshTokenExpiryInSeconds *int                               `json:"refreshTokenExpiryInSeconds"`
		DefinedTags                 map[string]map[string]interface{}  `json:"definedTags"`
		FreeformTags                map[string]string                  `json:"freeformTags"`
		Storage                     databasetoolsmcpserverstorage      `json:"storage"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.CustomRoles = make([]DatabaseToolsMcpServerCustomRole, len(model.CustomRoles))
	copy(m.CustomRoles, model.CustomRoles)
	m.AccessTokenExpiryInSeconds = model.AccessTokenExpiryInSeconds

	m.RefreshTokenExpiryInSeconds = model.RefreshTokenExpiryInSeconds

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

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
