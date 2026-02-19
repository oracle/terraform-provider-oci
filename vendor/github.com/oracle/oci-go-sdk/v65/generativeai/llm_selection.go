// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LlmSelection LLM selection configuration.
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
	case "GEN_AI_MODEL":
		mm := GenAiModelLlmSelection{}
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
	LlmSelectionLlmSelectionTypeGenAiModel LlmSelectionLlmSelectionTypeEnum = "GEN_AI_MODEL"
)

var mappingLlmSelectionLlmSelectionTypeEnum = map[string]LlmSelectionLlmSelectionTypeEnum{
	"GEN_AI_MODEL": LlmSelectionLlmSelectionTypeGenAiModel,
}

var mappingLlmSelectionLlmSelectionTypeEnumLowerCase = map[string]LlmSelectionLlmSelectionTypeEnum{
	"gen_ai_model": LlmSelectionLlmSelectionTypeGenAiModel,
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
		"GEN_AI_MODEL",
	}
}

// GetMappingLlmSelectionLlmSelectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLlmSelectionLlmSelectionTypeEnum(val string) (LlmSelectionLlmSelectionTypeEnum, bool) {
	enum, ok := mappingLlmSelectionLlmSelectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
