// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ToolConfig The configuration and type of Tool.
type ToolConfig interface {
}

type toolconfig struct {
	JsonData       []byte
	ToolConfigType string `json:"toolConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *toolconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertoolconfig toolconfig
	s := struct {
		Model Unmarshalertoolconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ToolConfigType = s.Model.ToolConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *toolconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ToolConfigType {
	case "SQL_TOOL_CONFIG":
		mm := SqlToolConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION_CALLING_TOOL_CONFIG":
		mm := FunctionCallingToolConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP_ENDPOINT_TOOL_CONFIG":
		mm := HttpEndpointToolConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AGENT_TOOL_CONFIG":
		mm := AgentToolConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RAG_TOOL_CONFIG":
		mm := RagToolConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ToolConfig: %s.", m.ToolConfigType)
		return *m, nil
	}
}

func (m toolconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m toolconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ToolConfigToolConfigTypeEnum Enum with underlying type: string
type ToolConfigToolConfigTypeEnum string

// Set of constants representing the allowable values for ToolConfigToolConfigTypeEnum
const (
	ToolConfigToolConfigTypeSqlToolConfig             ToolConfigToolConfigTypeEnum = "SQL_TOOL_CONFIG"
	ToolConfigToolConfigTypeRagToolConfig             ToolConfigToolConfigTypeEnum = "RAG_TOOL_CONFIG"
	ToolConfigToolConfigTypeFunctionCallingToolConfig ToolConfigToolConfigTypeEnum = "FUNCTION_CALLING_TOOL_CONFIG"
	ToolConfigToolConfigTypeHttpEndpointToolConfig    ToolConfigToolConfigTypeEnum = "HTTP_ENDPOINT_TOOL_CONFIG"
	ToolConfigToolConfigTypeAgentToolConfig           ToolConfigToolConfigTypeEnum = "AGENT_TOOL_CONFIG"
)

var mappingToolConfigToolConfigTypeEnum = map[string]ToolConfigToolConfigTypeEnum{
	"SQL_TOOL_CONFIG":              ToolConfigToolConfigTypeSqlToolConfig,
	"RAG_TOOL_CONFIG":              ToolConfigToolConfigTypeRagToolConfig,
	"FUNCTION_CALLING_TOOL_CONFIG": ToolConfigToolConfigTypeFunctionCallingToolConfig,
	"HTTP_ENDPOINT_TOOL_CONFIG":    ToolConfigToolConfigTypeHttpEndpointToolConfig,
	"AGENT_TOOL_CONFIG":            ToolConfigToolConfigTypeAgentToolConfig,
}

var mappingToolConfigToolConfigTypeEnumLowerCase = map[string]ToolConfigToolConfigTypeEnum{
	"sql_tool_config":              ToolConfigToolConfigTypeSqlToolConfig,
	"rag_tool_config":              ToolConfigToolConfigTypeRagToolConfig,
	"function_calling_tool_config": ToolConfigToolConfigTypeFunctionCallingToolConfig,
	"http_endpoint_tool_config":    ToolConfigToolConfigTypeHttpEndpointToolConfig,
	"agent_tool_config":            ToolConfigToolConfigTypeAgentToolConfig,
}

// GetToolConfigToolConfigTypeEnumValues Enumerates the set of values for ToolConfigToolConfigTypeEnum
func GetToolConfigToolConfigTypeEnumValues() []ToolConfigToolConfigTypeEnum {
	values := make([]ToolConfigToolConfigTypeEnum, 0)
	for _, v := range mappingToolConfigToolConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetToolConfigToolConfigTypeEnumStringValues Enumerates the set of values in String for ToolConfigToolConfigTypeEnum
func GetToolConfigToolConfigTypeEnumStringValues() []string {
	return []string{
		"SQL_TOOL_CONFIG",
		"RAG_TOOL_CONFIG",
		"FUNCTION_CALLING_TOOL_CONFIG",
		"HTTP_ENDPOINT_TOOL_CONFIG",
		"AGENT_TOOL_CONFIG",
	}
}

// GetMappingToolConfigToolConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingToolConfigToolConfigTypeEnum(val string) (ToolConfigToolConfigTypeEnum, bool) {
	enum, ok := mappingToolConfigToolConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
