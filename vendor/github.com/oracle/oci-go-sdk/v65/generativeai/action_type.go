// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Management API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service management API to create and manage DedicatedAiCluster, Endpoint, Model, and WorkRequest in the Generative AI service. For example, create a custom model by fine-tuning an out-of-the-box model using your own data, on a fine-tuning dedicated AI cluster. Then, create a hosting dedicated AI cluster with an endpoint to host your custom model.
// To access your custom model endpoints, or to try the out-of-the-box models to generate text, summarize, and create text embeddings see the Generative AI Inference API (https://docs.cloud.oracle.com/iaas/api/#/en/generative-ai-inference/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeai

import (
	"strings"
)

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeCreated    ActionTypeEnum = "CREATED"
	ActionTypeUpdated    ActionTypeEnum = "UPDATED"
	ActionTypeDeleted    ActionTypeEnum = "DELETED"
	ActionTypeInProgress ActionTypeEnum = "IN_PROGRESS"
	ActionTypeRelated    ActionTypeEnum = "RELATED"
	ActionTypeFailed     ActionTypeEnum = "FAILED"
)

var mappingActionTypeEnum = map[string]ActionTypeEnum{
	"CREATED":     ActionTypeCreated,
	"UPDATED":     ActionTypeUpdated,
	"DELETED":     ActionTypeDeleted,
	"IN_PROGRESS": ActionTypeInProgress,
	"RELATED":     ActionTypeRelated,
	"FAILED":      ActionTypeFailed,
}

var mappingActionTypeEnumLowerCase = map[string]ActionTypeEnum{
	"created":     ActionTypeCreated,
	"updated":     ActionTypeUpdated,
	"deleted":     ActionTypeDeleted,
	"in_progress": ActionTypeInProgress,
	"related":     ActionTypeRelated,
	"failed":      ActionTypeFailed,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionTypeEnumStringValues Enumerates the set of values in String for ActionTypeEnum
func GetActionTypeEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"IN_PROGRESS",
		"RELATED",
		"FAILED",
	}
}

// GetMappingActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypeEnum(val string) (ActionTypeEnum, bool) {
	enum, ok := mappingActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
