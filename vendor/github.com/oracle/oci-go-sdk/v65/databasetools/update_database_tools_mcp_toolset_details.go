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

// UpdateDatabaseToolsMcpToolsetDetails Database Tools MCP Toolset information to be updated.
type UpdateDatabaseToolsMcpToolsetDetails interface {

	// The MCP toolset version
	GetVersion() *int

	// A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
	GetDisplayName() *string

	// A human readable description of the Database Tools MCP toolset.
	GetDescription() *string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string
}

type updatedatabasetoolsmcptoolsetdetails struct {
	JsonData     []byte
	Version      *int                              `mandatory:"false" json:"version"`
	DisplayName  *string                           `mandatory:"false" json:"displayName"`
	Description  *string                           `mandatory:"false" json:"description"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	Type         string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatedatabasetoolsmcptoolsetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedatabasetoolsmcptoolsetdetails updatedatabasetoolsmcptoolsetdetails
	s := struct {
		Model Unmarshalerupdatedatabasetoolsmcptoolsetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Version = s.Model.Version
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.DefinedTags = s.Model.DefinedTags
	m.FreeformTags = s.Model.FreeformTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedatabasetoolsmcptoolsetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CUSTOMIZABLE_REPORTING_TOOLS":
		mm := UpdateDatabaseToolsMcpToolsetCustomizableReportingToolsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUILT_IN_SQL_TOOLS":
		mm := UpdateDatabaseToolsMcpToolsetBuiltInSqlToolsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_SQL_TOOL":
		mm := UpdateDatabaseToolsMcpToolsetCustomSqlToolDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENAI_SQL_ASSISTANT":
		mm := UpdateDatabaseToolsMcpToolsetGenAiSqlAssistantDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateDatabaseToolsMcpToolsetDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetVersion returns Version
func (m updatedatabasetoolsmcptoolsetdetails) GetVersion() *int {
	return m.Version
}

// GetDisplayName returns DisplayName
func (m updatedatabasetoolsmcptoolsetdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m updatedatabasetoolsmcptoolsetdetails) GetDescription() *string {
	return m.Description
}

// GetDefinedTags returns DefinedTags
func (m updatedatabasetoolsmcptoolsetdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m updatedatabasetoolsmcptoolsetdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

func (m updatedatabasetoolsmcptoolsetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedatabasetoolsmcptoolsetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
