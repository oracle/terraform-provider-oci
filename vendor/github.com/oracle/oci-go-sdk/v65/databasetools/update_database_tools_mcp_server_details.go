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

// UpdateDatabaseToolsMcpServerDetails Database Tools MCP server information to be updated.
type UpdateDatabaseToolsMcpServerDetails interface {

	// A meaningful, human-readable label displayed to end users. Not required to be unique and can be changed after creation. Do not include confidential information.
	GetDisplayName() *string

	// A human-readable description of the Database Tools MCP server.
	GetDescription() *string

	// Custom roles associated with the MCP Server.
	GetCustomRoles() []DatabaseToolsMcpServerCustomRole

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
}

type updatedatabasetoolsmcpserverdetails struct {
	JsonData                    []byte
	DisplayName                 *string                            `mandatory:"false" json:"displayName"`
	Description                 *string                            `mandatory:"false" json:"description"`
	CustomRoles                 []DatabaseToolsMcpServerCustomRole `mandatory:"false" json:"customRoles"`
	AccessTokenExpiryInSeconds  *int                               `mandatory:"false" json:"accessTokenExpiryInSeconds"`
	RefreshTokenExpiryInSeconds *int                               `mandatory:"false" json:"refreshTokenExpiryInSeconds"`
	DefinedTags                 map[string]map[string]interface{}  `mandatory:"false" json:"definedTags"`
	FreeformTags                map[string]string                  `mandatory:"false" json:"freeformTags"`
	Type                        string                             `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatedatabasetoolsmcpserverdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedatabasetoolsmcpserverdetails updatedatabasetoolsmcpserverdetails
	s := struct {
		Model Unmarshalerupdatedatabasetoolsmcpserverdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.CustomRoles = s.Model.CustomRoles
	m.AccessTokenExpiryInSeconds = s.Model.AccessTokenExpiryInSeconds
	m.RefreshTokenExpiryInSeconds = s.Model.RefreshTokenExpiryInSeconds
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedatabasetoolsmcpserverdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := UpdateDatabaseToolsMcpServerDetailsDefault{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateDatabaseToolsMcpServerDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updatedatabasetoolsmcpserverdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m updatedatabasetoolsmcpserverdetails) GetDescription() *string {
	return m.Description
}

// GetCustomRoles returns CustomRoles
func (m updatedatabasetoolsmcpserverdetails) GetCustomRoles() []DatabaseToolsMcpServerCustomRole {
	return m.CustomRoles
}

// GetAccessTokenExpiryInSeconds returns AccessTokenExpiryInSeconds
func (m updatedatabasetoolsmcpserverdetails) GetAccessTokenExpiryInSeconds() *int {
	return m.AccessTokenExpiryInSeconds
}

// GetRefreshTokenExpiryInSeconds returns RefreshTokenExpiryInSeconds
func (m updatedatabasetoolsmcpserverdetails) GetRefreshTokenExpiryInSeconds() *int {
	return m.RefreshTokenExpiryInSeconds
}

// GetDefinedTags returns DefinedTags
func (m updatedatabasetoolsmcpserverdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m updatedatabasetoolsmcpserverdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m updatedatabasetoolsmcpserverdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedatabasetoolsmcpserverdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
