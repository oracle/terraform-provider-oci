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
	"strings"
)

// ModelModalityEnum Enum with underlying type: string
type ModelModalityEnum string

// Set of constants representing the allowable values for ModelModalityEnum
const (
	ModelModalityAudio ModelModalityEnum = "AUDIO"
	ModelModalityVideo ModelModalityEnum = "VIDEO"
	ModelModalityText  ModelModalityEnum = "TEXT"
	ModelModalityImage ModelModalityEnum = "IMAGE"
)

var mappingModelModalityEnum = map[string]ModelModalityEnum{
	"AUDIO": ModelModalityAudio,
	"VIDEO": ModelModalityVideo,
	"TEXT":  ModelModalityText,
	"IMAGE": ModelModalityImage,
}

var mappingModelModalityEnumLowerCase = map[string]ModelModalityEnum{
	"audio": ModelModalityAudio,
	"video": ModelModalityVideo,
	"text":  ModelModalityText,
	"image": ModelModalityImage,
}

// GetModelModalityEnumValues Enumerates the set of values for ModelModalityEnum
func GetModelModalityEnumValues() []ModelModalityEnum {
	values := make([]ModelModalityEnum, 0)
	for _, v := range mappingModelModalityEnum {
		values = append(values, v)
	}
	return values
}

// GetModelModalityEnumStringValues Enumerates the set of values in String for ModelModalityEnum
func GetModelModalityEnumStringValues() []string {
	return []string{
		"AUDIO",
		"VIDEO",
		"TEXT",
		"IMAGE",
	}
}

// GetMappingModelModalityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelModalityEnum(val string) (ModelModalityEnum, bool) {
	enum, ok := mappingModelModalityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
