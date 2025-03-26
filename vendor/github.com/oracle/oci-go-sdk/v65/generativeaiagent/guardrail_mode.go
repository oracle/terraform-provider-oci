// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"strings"
)

// GuardrailModeEnum Enum with underlying type: string
type GuardrailModeEnum string

// Set of constants representing the allowable values for GuardrailModeEnum
const (
	GuardrailModeDisable GuardrailModeEnum = "DISABLE"
	GuardrailModeBlock   GuardrailModeEnum = "BLOCK"
	GuardrailModeInform  GuardrailModeEnum = "INFORM"
)

var mappingGuardrailModeEnum = map[string]GuardrailModeEnum{
	"DISABLE": GuardrailModeDisable,
	"BLOCK":   GuardrailModeBlock,
	"INFORM":  GuardrailModeInform,
}

var mappingGuardrailModeEnumLowerCase = map[string]GuardrailModeEnum{
	"disable": GuardrailModeDisable,
	"block":   GuardrailModeBlock,
	"inform":  GuardrailModeInform,
}

// GetGuardrailModeEnumValues Enumerates the set of values for GuardrailModeEnum
func GetGuardrailModeEnumValues() []GuardrailModeEnum {
	values := make([]GuardrailModeEnum, 0)
	for _, v := range mappingGuardrailModeEnum {
		values = append(values, v)
	}
	return values
}

// GetGuardrailModeEnumStringValues Enumerates the set of values in String for GuardrailModeEnum
func GetGuardrailModeEnumStringValues() []string {
	return []string{
		"DISABLE",
		"BLOCK",
		"INFORM",
	}
}

// GetMappingGuardrailModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGuardrailModeEnum(val string) (GuardrailModeEnum, bool) {
	enum, ok := mappingGuardrailModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
