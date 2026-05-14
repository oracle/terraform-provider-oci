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

// DatabaseToolsMcpToolsetBuiltInSqlTools Allows the creation, configuration and management of an McpToolset of type Built-in SQL tools.
type DatabaseToolsMcpToolsetBuiltInSqlTools struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools MCP Toolsets.
	Id *string `mandatory:"true" json:"id"`

	// The MCP toolset version
	Version *int `mandatory:"true" json:"version"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to move the Database Tools MCP Toolset to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools McpServer.
	DatabaseToolsMcpServerId *string `mandatory:"true" json:"databaseToolsMcpServerId"`

	// The time the Database Tools MCP server was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Database Tools MCP server was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A human readable description of the Database Tools MCP toolset.
	Description *string `mandatory:"false" json:"description"`

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

	// List of Database Tools Built-in SQL tools
	Tools []DatabaseToolsMcpToolsetToolDetails `mandatory:"false" json:"tools"`

	// The current state of the Database Tools MCP Toolset.
	LifecycleState DatabaseToolsMcpToolsetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The default execution type for the toolset. The default value is SYNCHRONOUS.
	// To use ASYNCHRONOUS execution, the MCP Server must have the storage property configured.
	DefaultExecutionType DatabaseToolsMcpToolsetDefaultExecutionTypeEnum `mandatory:"false" json:"defaultExecutionType,omitempty"`
}

// GetId returns Id
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetId() *string {
	return m.Id
}

// GetVersion returns Version
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetVersion() *int {
	return m.Version
}

// GetCompartmentId returns CompartmentId
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetDescription() *string {
	return m.Description
}

// GetDatabaseToolsMcpServerId returns DatabaseToolsMcpServerId
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetDatabaseToolsMcpServerId() *string {
	return m.DatabaseToolsMcpServerId
}

// GetLifecycleState returns LifecycleState
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetLifecycleState() DatabaseToolsMcpToolsetLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetDefinedTags returns DefinedTags
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) GetLocks() []ResourceLock {
	return m.Locks
}

func (m DatabaseToolsMcpToolsetBuiltInSqlTools) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseToolsMcpToolsetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsMcpToolsetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsMcpToolsetDefaultExecutionTypeEnum(string(m.DefaultExecutionType)); !ok && m.DefaultExecutionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultExecutionType: %s. Supported values are: %s.", m.DefaultExecutionType, strings.Join(GetDatabaseToolsMcpToolsetDefaultExecutionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsMcpToolsetBuiltInSqlTools) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsMcpToolsetBuiltInSqlTools DatabaseToolsMcpToolsetBuiltInSqlTools
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsMcpToolsetBuiltInSqlTools
	}{
		"BUILT_IN_SQL_TOOLS",
		(MarshalTypeDatabaseToolsMcpToolsetBuiltInSqlTools)(m),
	}

	return json.Marshal(&s)
}
