// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"strings"
)

// ImportedModelCapabilityEnum Enum with underlying type: string
type ImportedModelCapabilityEnum string

// Set of constants representing the allowable values for ImportedModelCapabilityEnum
const (
	ImportedModelCapabilityTextToText      ImportedModelCapabilityEnum = "TEXT_TO_TEXT"
	ImportedModelCapabilityImageTextToText ImportedModelCapabilityEnum = "IMAGE_TEXT_TO_TEXT"
	ImportedModelCapabilityEmbedding       ImportedModelCapabilityEnum = "EMBEDDING"
	ImportedModelCapabilityRerank          ImportedModelCapabilityEnum = "RERANK"
)

var mappingImportedModelCapabilityEnum = map[string]ImportedModelCapabilityEnum{
	"TEXT_TO_TEXT":       ImportedModelCapabilityTextToText,
	"IMAGE_TEXT_TO_TEXT": ImportedModelCapabilityImageTextToText,
	"EMBEDDING":          ImportedModelCapabilityEmbedding,
	"RERANK":             ImportedModelCapabilityRerank,
}

var mappingImportedModelCapabilityEnumLowerCase = map[string]ImportedModelCapabilityEnum{
	"text_to_text":       ImportedModelCapabilityTextToText,
	"image_text_to_text": ImportedModelCapabilityImageTextToText,
	"embedding":          ImportedModelCapabilityEmbedding,
	"rerank":             ImportedModelCapabilityRerank,
}

// GetImportedModelCapabilityEnumValues Enumerates the set of values for ImportedModelCapabilityEnum
func GetImportedModelCapabilityEnumValues() []ImportedModelCapabilityEnum {
	values := make([]ImportedModelCapabilityEnum, 0)
	for _, v := range mappingImportedModelCapabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetImportedModelCapabilityEnumStringValues Enumerates the set of values in String for ImportedModelCapabilityEnum
func GetImportedModelCapabilityEnumStringValues() []string {
	return []string{
		"TEXT_TO_TEXT",
		"IMAGE_TEXT_TO_TEXT",
		"EMBEDDING",
		"RERANK",
	}
}

// GetMappingImportedModelCapabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportedModelCapabilityEnum(val string) (ImportedModelCapabilityEnum, bool) {
	enum, ok := mappingImportedModelCapabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
