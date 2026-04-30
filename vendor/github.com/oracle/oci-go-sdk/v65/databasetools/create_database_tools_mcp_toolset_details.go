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

// CreateDatabaseToolsMcpToolsetDetails Details for the new Database Tools MCP server.
type CreateDatabaseToolsMcpToolsetDetails interface {

	// The MCP toolset version
	GetVersion() *int

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
	GetCompartmentId() *string

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the Database Tools MCP Server
	GetDatabaseToolsMcpServerId() *string

	// A human readable description of the Database Tools MCP toolset.
	GetDescription() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type createdatabasetoolsmcptoolsetdetails struct {
	JsonData                 []byte
	Description              *string                           `mandatory:"false" json:"description"`
	DefinedTags              map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags             map[string]string                 `mandatory:"false" json:"freeformTags"`
	Locks                    []ResourceLock                    `mandatory:"false" json:"locks"`
	Version                  *int                              `mandatory:"true" json:"version"`
	CompartmentId            *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName              *string                           `mandatory:"true" json:"displayName"`
	DatabaseToolsMcpServerId *string                           `mandatory:"true" json:"databaseToolsMcpServerId"`
	Type                     string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasetoolsmcptoolsetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasetoolsmcptoolsetdetails createdatabasetoolsmcptoolsetdetails
	s := struct {
		Model Unmarshalercreatedatabasetoolsmcptoolsetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Version = s.Model.Version
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.DatabaseToolsMcpServerId = s.Model.DatabaseToolsMcpServerId
	m.Description = s.Model.Description
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.Locks = s.Model.Locks
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasetoolsmcptoolsetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "BUILT_IN_SQL_TOOLS":
		mm := CreateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_SQL_TOOL":
		mm := CreateDatabaseToolsMcpToolsetCustomSqlToolDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENAI_SQL_ASSISTANT":
		mm := CreateDatabaseToolsMcpToolsetGenAiSqlAssistantDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMIZABLE_REPORTING_TOOLS":
		mm := CreateDatabaseToolsMcpToolsetCustomizableReportingToolsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDatabaseToolsMcpToolsetDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createdatabasetoolsmcptoolsetdetails) GetDescription() *string {
	return m.Description
}

// GetDefinedTags returns DefinedTags
func (m createdatabasetoolsmcptoolsetdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m createdatabasetoolsmcptoolsetdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m createdatabasetoolsmcptoolsetdetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetVersion returns Version
func (m createdatabasetoolsmcptoolsetdetails) GetVersion() *int {
	return m.Version
}

// GetCompartmentId returns CompartmentId
func (m createdatabasetoolsmcptoolsetdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m createdatabasetoolsmcptoolsetdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsMcpServerId returns DatabaseToolsMcpServerId
func (m createdatabasetoolsmcptoolsetdetails) GetDatabaseToolsMcpServerId() *string {
	return m.DatabaseToolsMcpServerId
}

func (m createdatabasetoolsmcptoolsetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasetoolsmcptoolsetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
