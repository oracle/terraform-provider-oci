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

// ModelCapabilityEnum Enum with underlying type: string
type ModelCapabilityEnum string

// Set of constants representing the allowable values for ModelCapabilityEnum
const (
	ModelCapabilityTextGeneration    ModelCapabilityEnum = "TEXT_GENERATION"
	ModelCapabilityTextSummarization ModelCapabilityEnum = "TEXT_SUMMARIZATION"
	ModelCapabilityTextEmbeddings    ModelCapabilityEnum = "TEXT_EMBEDDINGS"
	ModelCapabilityFineTune          ModelCapabilityEnum = "FINE_TUNE"
	ModelCapabilityChat              ModelCapabilityEnum = "CHAT"
	ModelCapabilityTextRerank        ModelCapabilityEnum = "TEXT_RERANK"
	ModelCapabilityTextToImage       ModelCapabilityEnum = "TEXT_TO_IMAGE"
	ModelCapabilityImageTextToImage  ModelCapabilityEnum = "IMAGE_TEXT_TO_IMAGE"
	ModelCapabilityImageTextToText   ModelCapabilityEnum = "IMAGE_TEXT_TO_TEXT"
	ModelCapabilityImageTextToVideo  ModelCapabilityEnum = "IMAGE_TEXT_TO_VIDEO"
	ModelCapabilityImageToImage      ModelCapabilityEnum = "IMAGE_TO_IMAGE"
	ModelCapabilityRealtime          ModelCapabilityEnum = "REALTIME"
	ModelCapabilityAudioToAudio      ModelCapabilityEnum = "AUDIO_TO_AUDIO"
	ModelCapabilityAudioToText       ModelCapabilityEnum = "AUDIO_TO_TEXT"
	ModelCapabilityTextToAudio       ModelCapabilityEnum = "TEXT_TO_AUDIO"
	ModelCapabilityTextToVideo       ModelCapabilityEnum = "TEXT_TO_VIDEO"
)

var mappingModelCapabilityEnum = map[string]ModelCapabilityEnum{
	"TEXT_GENERATION":     ModelCapabilityTextGeneration,
	"TEXT_SUMMARIZATION":  ModelCapabilityTextSummarization,
	"TEXT_EMBEDDINGS":     ModelCapabilityTextEmbeddings,
	"FINE_TUNE":           ModelCapabilityFineTune,
	"CHAT":                ModelCapabilityChat,
	"TEXT_RERANK":         ModelCapabilityTextRerank,
	"TEXT_TO_IMAGE":       ModelCapabilityTextToImage,
	"IMAGE_TEXT_TO_IMAGE": ModelCapabilityImageTextToImage,
	"IMAGE_TEXT_TO_TEXT":  ModelCapabilityImageTextToText,
	"IMAGE_TEXT_TO_VIDEO": ModelCapabilityImageTextToVideo,
	"IMAGE_TO_IMAGE":      ModelCapabilityImageToImage,
	"REALTIME":            ModelCapabilityRealtime,
	"AUDIO_TO_AUDIO":      ModelCapabilityAudioToAudio,
	"AUDIO_TO_TEXT":       ModelCapabilityAudioToText,
	"TEXT_TO_AUDIO":       ModelCapabilityTextToAudio,
	"TEXT_TO_VIDEO":       ModelCapabilityTextToVideo,
}

var mappingModelCapabilityEnumLowerCase = map[string]ModelCapabilityEnum{
	"text_generation":     ModelCapabilityTextGeneration,
	"text_summarization":  ModelCapabilityTextSummarization,
	"text_embeddings":     ModelCapabilityTextEmbeddings,
	"fine_tune":           ModelCapabilityFineTune,
	"chat":                ModelCapabilityChat,
	"text_rerank":         ModelCapabilityTextRerank,
	"text_to_image":       ModelCapabilityTextToImage,
	"image_text_to_image": ModelCapabilityImageTextToImage,
	"image_text_to_text":  ModelCapabilityImageTextToText,
	"image_text_to_video": ModelCapabilityImageTextToVideo,
	"image_to_image":      ModelCapabilityImageToImage,
	"realtime":            ModelCapabilityRealtime,
	"audio_to_audio":      ModelCapabilityAudioToAudio,
	"audio_to_text":       ModelCapabilityAudioToText,
	"text_to_audio":       ModelCapabilityTextToAudio,
	"text_to_video":       ModelCapabilityTextToVideo,
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
		"TEXT_RERANK",
		"TEXT_TO_IMAGE",
		"IMAGE_TEXT_TO_IMAGE",
		"IMAGE_TEXT_TO_TEXT",
		"IMAGE_TEXT_TO_VIDEO",
		"IMAGE_TO_IMAGE",
		"REALTIME",
		"AUDIO_TO_AUDIO",
		"AUDIO_TO_TEXT",
		"TEXT_TO_AUDIO",
		"TEXT_TO_VIDEO",
	}
}

// GetMappingModelCapabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelCapabilityEnum(val string) (ModelCapabilityEnum, bool) {
	enum, ok := mappingModelCapabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
