// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service API
//
// **Generative AI Service**
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable LLMs that cover a wide range of use cases for text generation. Use the playground to try out the models out-of-the-box or create and host your own fine-tuned custom models based on your own data on dedicated AI clusters.
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
)

var mappingModelCapabilityEnum = map[string]ModelCapabilityEnum{
	"TEXT_GENERATION":    ModelCapabilityTextGeneration,
	"TEXT_SUMMARIZATION": ModelCapabilityTextSummarization,
	"TEXT_EMBEDDINGS":    ModelCapabilityTextEmbeddings,
	"FINE_TUNE":          ModelCapabilityFineTune,
}

var mappingModelCapabilityEnumLowerCase = map[string]ModelCapabilityEnum{
	"text_generation":    ModelCapabilityTextGeneration,
	"text_summarization": ModelCapabilityTextSummarization,
	"text_embeddings":    ModelCapabilityTextEmbeddings,
	"fine_tune":          ModelCapabilityFineTune,
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
	}
}

// GetMappingModelCapabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelCapabilityEnum(val string) (ModelCapabilityEnum, bool) {
	enum, ok := mappingModelCapabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
