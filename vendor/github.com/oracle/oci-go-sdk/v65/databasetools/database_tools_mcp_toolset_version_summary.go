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

// DatabaseToolsMcpToolsetVersionSummary Summary of versions available for a specific MCP toolset type.
type DatabaseToolsMcpToolsetVersionSummary interface {

	// The default version for this toolset type.
	GetDefaultVersion() *int
}

type databasetoolsmcptoolsetversionsummary struct {
	JsonData       []byte
	DefaultVersion *int   `mandatory:"true" json:"defaultVersion"`
	Type           string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolsmcptoolsetversionsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolsmcptoolsetversionsummary databasetoolsmcptoolsetversionsummary
	s := struct {
		Model Unmarshalerdatabasetoolsmcptoolsetversionsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DefaultVersion = s.Model.DefaultVersion
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolsmcptoolsetversionsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CUSTOM_SQL_TOOL":
		mm := DatabaseToolsMcpToolsetVersionCustomSqlToolSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMIZABLE_REPORTING_TOOLS":
		mm := DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUILT_IN_SQL_TOOLS":
		mm := DatabaseToolsMcpToolsetVersionBuiltInSqlToolsSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENAI_SQL_ASSISTANT":
		mm := DatabaseToolsMcpToolsetVersionGenAiSqlAssistantSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabaseToolsMcpToolsetVersionSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetDefaultVersion returns DefaultVersion
func (m databasetoolsmcptoolsetversionsummary) GetDefaultVersion() *int {
	return m.DefaultVersion
}

func (m databasetoolsmcptoolsetversionsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolsmcptoolsetversionsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
