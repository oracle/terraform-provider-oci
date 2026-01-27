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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LlmSelection LLM selection configuration - either DEFAULT or CUSTOM.
type LlmSelection interface {
}

type llmselection struct {
	JsonData         []byte
	LlmSelectionType string `json:"llmSelectionType"`
}

// UnmarshalJSON unmarshals json
func (m *llmselection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerllmselection llmselection
	s := struct {
		Model Unmarshalerllmselection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.LlmSelectionType = s.Model.LlmSelectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *llmselection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.LlmSelectionType {
	case "DEFAULT":
		mm := DefaultLlmSelection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_GEN_AI_ENDPOINT":
		mm := CustomGenAiEndpointLlmSelection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_GEN_AI_MODEL":
		mm := CustomGenAiModelLlmSelection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for LlmSelection: %s.", m.LlmSelectionType)
		return *m, nil
	}
}

func (m llmselection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m llmselection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LlmSelectionLlmSelectionTypeEnum Enum with underlying type: string
type LlmSelectionLlmSelectionTypeEnum string

// Set of constants representing the allowable values for LlmSelectionLlmSelectionTypeEnum
const (
	LlmSelectionLlmSelectionTypeDefault             LlmSelectionLlmSelectionTypeEnum = "DEFAULT"
	LlmSelectionLlmSelectionTypeCustomGenAiModel    LlmSelectionLlmSelectionTypeEnum = "CUSTOM_GEN_AI_MODEL"
	LlmSelectionLlmSelectionTypeCustomGenAiEndpoint LlmSelectionLlmSelectionTypeEnum = "CUSTOM_GEN_AI_ENDPOINT"
)

var mappingLlmSelectionLlmSelectionTypeEnum = map[string]LlmSelectionLlmSelectionTypeEnum{
	"DEFAULT":                LlmSelectionLlmSelectionTypeDefault,
	"CUSTOM_GEN_AI_MODEL":    LlmSelectionLlmSelectionTypeCustomGenAiModel,
	"CUSTOM_GEN_AI_ENDPOINT": LlmSelectionLlmSelectionTypeCustomGenAiEndpoint,
}

var mappingLlmSelectionLlmSelectionTypeEnumLowerCase = map[string]LlmSelectionLlmSelectionTypeEnum{
	"default":                LlmSelectionLlmSelectionTypeDefault,
	"custom_gen_ai_model":    LlmSelectionLlmSelectionTypeCustomGenAiModel,
	"custom_gen_ai_endpoint": LlmSelectionLlmSelectionTypeCustomGenAiEndpoint,
}

// GetLlmSelectionLlmSelectionTypeEnumValues Enumerates the set of values for LlmSelectionLlmSelectionTypeEnum
func GetLlmSelectionLlmSelectionTypeEnumValues() []LlmSelectionLlmSelectionTypeEnum {
	values := make([]LlmSelectionLlmSelectionTypeEnum, 0)
	for _, v := range mappingLlmSelectionLlmSelectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLlmSelectionLlmSelectionTypeEnumStringValues Enumerates the set of values in String for LlmSelectionLlmSelectionTypeEnum
func GetLlmSelectionLlmSelectionTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM_GEN_AI_MODEL",
		"CUSTOM_GEN_AI_ENDPOINT",
	}
}

// GetMappingLlmSelectionLlmSelectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLlmSelectionLlmSelectionTypeEnum(val string) (LlmSelectionLlmSelectionTypeEnum, bool) {
	enum, ok := mappingLlmSelectionLlmSelectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
