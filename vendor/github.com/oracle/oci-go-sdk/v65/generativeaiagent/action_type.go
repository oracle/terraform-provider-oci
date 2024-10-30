// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// **Generative AI Agents API**
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

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
