// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.cloud.oracle.com/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"strings"
)

// ModelCapabilityEnum Enum with underlying type: string
type ModelCapabilityEnum string

// Set of constants representing the allowable values for ModelCapabilityEnum
const (
	ModelCapabilityTextGeneration    ModelCapabilityEnum = "TEXT_GENERATION"
	ModelCapabilityTextSummarization ModelCapabilityEnum = "TEXT_SUMMARIZATION"
	ModelCapabilityTextEmbeddings    ModelCapabilityEnum = "TEXT_EMBEDDINGS"
	ModelCapabilityFineTune          ModelCapabilityEnum = "FINE_TUNE"
	ModelCapabilityChat              ModelCapabilityEnum = "CHAT"
)

var mappingModelCapabilityEnum = map[string]ModelCapabilityEnum{
	"TEXT_GENERATION":    ModelCapabilityTextGeneration,
	"TEXT_SUMMARIZATION": ModelCapabilityTextSummarization,
	"TEXT_EMBEDDINGS":    ModelCapabilityTextEmbeddings,
	"FINE_TUNE":          ModelCapabilityFineTune,
	"CHAT":               ModelCapabilityChat,
}

var mappingModelCapabilityEnumLowerCase = map[string]ModelCapabilityEnum{
	"text_generation":    ModelCapabilityTextGeneration,
	"text_summarization": ModelCapabilityTextSummarization,
	"text_embeddings":    ModelCapabilityTextEmbeddings,
	"fine_tune":          ModelCapabilityFineTune,
	"chat":               ModelCapabilityChat,
}

// GetModelCapabilityEnumValues Enumerates the set of values for ModelCapabilityEnum
func GetModelCapabilityEnumValues() []ModelCapabilityEnum {
	values := make([]ModelCapabilityEnum, 0)
	for _, v := range mappingModelCapabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetModelCapabilityEnumStringValues Enumerates the set of values in String for ModelCapabilityEnum
func GetModelCapabilityEnumStringValues() []string {
	return []string{
		"TEXT_GENERATION",
		"TEXT_SUMMARIZATION",
		"TEXT_EMBEDDINGS",
		"FINE_TUNE",
		"CHAT",
	}
}

// GetMappingModelCapabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelCapabilityEnum(val string) (ModelCapabilityEnum, bool) {
	enum, ok := mappingModelCapabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
