// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ToolRuntimeConfig Configuration for the tool.
type ToolRuntimeConfig struct {

	// The type of the tool.
	ToolRuntimeConfigType ToolRuntimeConfigToolRuntimeConfigTypeEnum `mandatory:"true" json:"toolRuntimeConfigType"`

	// The version of the components.
	Version *string `mandatory:"false" json:"version"`
}

func (m ToolRuntimeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ToolRuntimeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingToolRuntimeConfigToolRuntimeConfigTypeEnum(string(m.ToolRuntimeConfigType)); !ok && m.ToolRuntimeConfigType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ToolRuntimeConfigType: %s. Supported values are: %s.", m.ToolRuntimeConfigType, strings.Join(GetToolRuntimeConfigToolRuntimeConfigTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ToolRuntimeConfigToolRuntimeConfigTypeEnum Enum with underlying type: string
type ToolRuntimeConfigToolRuntimeConfigTypeEnum string

// Set of constants representing the allowable values for ToolRuntimeConfigToolRuntimeConfigTypeEnum
const (
	ToolRuntimeConfigToolRuntimeConfigTypeRag      ToolRuntimeConfigToolRuntimeConfigTypeEnum = "RAG"
	ToolRuntimeConfigToolRuntimeConfigTypeSqlSmall ToolRuntimeConfigToolRuntimeConfigTypeEnum = "SQL_SMALL"
	ToolRuntimeConfigToolRuntimeConfigTypeSqlLarge ToolRuntimeConfigToolRuntimeConfigTypeEnum = "SQL_LARGE"
)

var mappingToolRuntimeConfigToolRuntimeConfigTypeEnum = map[string]ToolRuntimeConfigToolRuntimeConfigTypeEnum{
	"RAG":       ToolRuntimeConfigToolRuntimeConfigTypeRag,
	"SQL_SMALL": ToolRuntimeConfigToolRuntimeConfigTypeSqlSmall,
	"SQL_LARGE": ToolRuntimeConfigToolRuntimeConfigTypeSqlLarge,
}

var mappingToolRuntimeConfigToolRuntimeConfigTypeEnumLowerCase = map[string]ToolRuntimeConfigToolRuntimeConfigTypeEnum{
	"rag":       ToolRuntimeConfigToolRuntimeConfigTypeRag,
	"sql_small": ToolRuntimeConfigToolRuntimeConfigTypeSqlSmall,
	"sql_large": ToolRuntimeConfigToolRuntimeConfigTypeSqlLarge,
}

// GetToolRuntimeConfigToolRuntimeConfigTypeEnumValues Enumerates the set of values for ToolRuntimeConfigToolRuntimeConfigTypeEnum
func GetToolRuntimeConfigToolRuntimeConfigTypeEnumValues() []ToolRuntimeConfigToolRuntimeConfigTypeEnum {
	values := make([]ToolRuntimeConfigToolRuntimeConfigTypeEnum, 0)
	for _, v := range mappingToolRuntimeConfigToolRuntimeConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetToolRuntimeConfigToolRuntimeConfigTypeEnumStringValues Enumerates the set of values in String for ToolRuntimeConfigToolRuntimeConfigTypeEnum
func GetToolRuntimeConfigToolRuntimeConfigTypeEnumStringValues() []string {
	return []string{
		"RAG",
		"SQL_SMALL",
		"SQL_LARGE",
	}
}

// GetMappingToolRuntimeConfigToolRuntimeConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingToolRuntimeConfigToolRuntimeConfigTypeEnum(val string) (ToolRuntimeConfigToolRuntimeConfigTypeEnum, bool) {
	enum, ok := mappingToolRuntimeConfigToolRuntimeConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
