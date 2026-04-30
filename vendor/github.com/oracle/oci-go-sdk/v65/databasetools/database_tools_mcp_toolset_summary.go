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

// DatabaseToolsMcpToolsetSummary Summary of the Database Tools MCP Toolset.
type DatabaseToolsMcpToolsetSummary interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools MCP Toolset.
	GetId() *string

	// The toolset version
	GetVersion() *int

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
	GetCompartmentId() *string

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools McpServer.
	GetDatabaseToolsMcpServerId() *string

	// The current state of the Database Tools MCP Toolset.
	GetLifecycleState() DatabaseToolsMcpToolsetLifecycleStateEnum

	// The time the Database Tools MCP server was created. An RFC3339 formatted datetime string.
	GetTimeCreated() *common.SDKTime

	// The time the Database Tools MCP server was updated. An RFC3339 formatted datetime string.
	GetTimeUpdated() *common.SDKTime

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

type databasetoolsmcptoolsetsummary struct {
	JsonData                 []byte
	LifecycleDetails         *string                                   `mandatory:"false" json:"lifecycleDetails"`
	DefinedTags              map[string]map[string]interface{}         `mandatory:"false" json:"definedTags"`
	FreeformTags             map[string]string                         `mandatory:"false" json:"freeformTags"`
	SystemTags               map[string]map[string]interface{}         `mandatory:"false" json:"systemTags"`
	Locks                    []ResourceLock                            `mandatory:"false" json:"locks"`
	Id                       *string                                   `mandatory:"true" json:"id"`
	Version                  *int                                      `mandatory:"true" json:"version"`
	CompartmentId            *string                                   `mandatory:"true" json:"compartmentId"`
	DisplayName              *string                                   `mandatory:"true" json:"displayName"`
	DatabaseToolsMcpServerId *string                                   `mandatory:"true" json:"databaseToolsMcpServerId"`
	LifecycleState           DatabaseToolsMcpToolsetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	TimeCreated              *common.SDKTime                           `mandatory:"true" json:"timeCreated"`
	TimeUpdated              *common.SDKTime                           `mandatory:"true" json:"timeUpdated"`
	Type                     string                                    `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsmcptoolsetsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsmcptoolsetsummary databasetoolsmcptoolsetsummary
	s := struct {
		Model Unmarshalerdatabasetoolsmcptoolsetsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.Version = s.Model.Version
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.DatabaseToolsMcpServerId = s.Model.DatabaseToolsMcpServerId
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.SystemTags = s.Model.SystemTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsmcptoolsetsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "BUILT_IN_SQL_TOOLS":
		mm := DatabaseToolsMcpToolsetBuiltInSqlToolsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMIZABLE_REPORTING_TOOLS":
		mm := DatabaseToolsMcpToolsetCustomizableReportingToolsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_SQL_TOOL":
		mm := DatabaseToolsMcpToolsetCustomSqlToolSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENAI_SQL_ASSISTANT":
		mm := DatabaseToolsMcpToolsetGenAiSqlAssistantSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsMcpToolsetSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetLifecycleDetails returns LifecycleDetails
func (m databasetoolsmcptoolsetsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDefinedTags returns DefinedTags
func (m databasetoolsmcptoolsetsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m databasetoolsmcptoolsetsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m databasetoolsmcptoolsetsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m databasetoolsmcptoolsetsummary) GetLocks() []ResourceLock {
	return m.Locks
}

// GetId returns Id
func (m databasetoolsmcptoolsetsummary) GetId() *string {
	return m.Id
}

// GetVersion returns Version
func (m databasetoolsmcptoolsetsummary) GetVersion() *int {
	return m.Version
}

// GetCompartmentId returns CompartmentId
func (m databasetoolsmcptoolsetsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m databasetoolsmcptoolsetsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsMcpServerId returns DatabaseToolsMcpServerId
func (m databasetoolsmcptoolsetsummary) GetDatabaseToolsMcpServerId() *string {
	return m.DatabaseToolsMcpServerId
}

// GetLifecycleState returns LifecycleState
func (m databasetoolsmcptoolsetsummary) GetLifecycleState() DatabaseToolsMcpToolsetLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m databasetoolsmcptoolsetsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m databasetoolsmcptoolsetsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m databasetoolsmcptoolsetsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsmcptoolsetsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolsMcpToolsetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsMcpToolsetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
