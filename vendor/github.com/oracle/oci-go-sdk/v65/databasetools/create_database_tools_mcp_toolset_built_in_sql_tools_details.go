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

// CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails Details for the new Database Tools MCP Toolset of type Built-in SQL tools.
type CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails struct {

	// The MCP toolset version
	Version *int `mandatory:"true" json:"version"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the Database Tools MCP Server
	DatabaseToolsMcpServerId *string `mandatory:"true" json:"databaseToolsMcpServerId"`

	// A human readable description of the Database Tools MCP toolset.
	Description *string `mandatory:"false" json:"description"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// List of Database Tools BUILT_IN_SQL_TOOLS toolset tool configurations
	Tools []CreateDatabaseToolsMcpToolsetToolDetails `mandatory:"false" json:"tools"`

	// The default execution type for the toolset. The default value is SYNCHRONOUS.
	// To use ASYNCHRONOUS execution, the MCP Server must have the storage property configured.
	DefaultExecutionType DatabaseToolsMcpToolsetDefaultExecutionTypeEnum `mandatory:"false" json:"defaultExecutionType,omitempty"`
}

// GetVersion returns Version
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) GetVersion() *int {
	return m.Version
}

// GetCompartmentId returns CompartmentId
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsMcpServerId returns DatabaseToolsMcpServerId
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) GetDatabaseToolsMcpServerId() *string {
	return m.DatabaseToolsMcpServerId
}

// GetDescription returns Description
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) GetDescription() *string {
	return m.Description
}

// GetDefinedTags returns DefinedTags
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) GetLocks() []ResourceLock {
	return m.Locks
}

func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsMcpToolsetDefaultExecutionTypeEnum(string(m.DefaultExecutionType)); !ok && m.DefaultExecutionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultExecutionType: %s. Supported values are: %s.", m.DefaultExecutionType, strings.Join(GetDatabaseToolsMcpToolsetDefaultExecutionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails
	}{
		"BUILT_IN_SQL_TOOLS",
		(MarshalTypeCreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails)(m),
	}

	return json.Marshal(&s)
}
