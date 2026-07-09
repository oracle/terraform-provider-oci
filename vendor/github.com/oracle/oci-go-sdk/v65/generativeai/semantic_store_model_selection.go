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

// SemanticStoreModelSelection Model selection configuration for SemanticStore enrichment.
type SemanticStoreModelSelection interface {
}

type semanticstoremodelselection struct {
	JsonData           []byte
	ModelSelectionType string `json:"modelSelectionType"`
}

// UnmarshalJSON unmarshals json
func (m *semanticstoremodelselection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersemanticstoremodelselection semanticstoremodelselection
	s := struct {
		Model Unmarshalersemanticstoremodelselection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelSelectionType = s.Model.ModelSelectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *semanticstoremodelselection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelSelectionType {
	case "DEFAULT":
		mm := DefaultSemanticStoreModelSelection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM":
		mm := CustomSemanticStoreModelSelection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SemanticStoreModelSelection: %s.", m.ModelSelectionType)
		return *m, nil
	}
}

func (m semanticstoremodelselection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m semanticstoremodelselection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SemanticStoreModelSelectionModelSelectionTypeEnum Enum with underlying type: string
type SemanticStoreModelSelectionModelSelectionTypeEnum string

// Set of constants representing the allowable values for SemanticStoreModelSelectionModelSelectionTypeEnum
const (
	SemanticStoreModelSelectionModelSelectionTypeDefault SemanticStoreModelSelectionModelSelectionTypeEnum = "DEFAULT"
	SemanticStoreModelSelectionModelSelectionTypeCustom  SemanticStoreModelSelectionModelSelectionTypeEnum = "CUSTOM"
)

var mappingSemanticStoreModelSelectionModelSelectionTypeEnum = map[string]SemanticStoreModelSelectionModelSelectionTypeEnum{
	"DEFAULT": SemanticStoreModelSelectionModelSelectionTypeDefault,
	"CUSTOM":  SemanticStoreModelSelectionModelSelectionTypeCustom,
}

var mappingSemanticStoreModelSelectionModelSelectionTypeEnumLowerCase = map[string]SemanticStoreModelSelectionModelSelectionTypeEnum{
	"default": SemanticStoreModelSelectionModelSelectionTypeDefault,
	"custom":  SemanticStoreModelSelectionModelSelectionTypeCustom,
}

// GetSemanticStoreModelSelectionModelSelectionTypeEnumValues Enumerates the set of values for SemanticStoreModelSelectionModelSelectionTypeEnum
func GetSemanticStoreModelSelectionModelSelectionTypeEnumValues() []SemanticStoreModelSelectionModelSelectionTypeEnum {
	values := make([]SemanticStoreModelSelectionModelSelectionTypeEnum, 0)
	for _, v := range mappingSemanticStoreModelSelectionModelSelectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSemanticStoreModelSelectionModelSelectionTypeEnumStringValues Enumerates the set of values in String for SemanticStoreModelSelectionModelSelectionTypeEnum
func GetSemanticStoreModelSelectionModelSelectionTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM",
	}
}

// GetMappingSemanticStoreModelSelectionModelSelectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSemanticStoreModelSelectionModelSelectionTypeEnum(val string) (SemanticStoreModelSelectionModelSelectionTypeEnum, bool) {
	enum, ok := mappingSemanticStoreModelSelectionModelSelectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
