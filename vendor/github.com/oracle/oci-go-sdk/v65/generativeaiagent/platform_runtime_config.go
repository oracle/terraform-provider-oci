// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PlatformRuntimeConfig Configuration for agent platform component.
type PlatformRuntimeConfig struct {

	// The type of Platform runtime config.
	PlatformRuntimeConfigType PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum `mandatory:"false" json:"platformRuntimeConfigType,omitempty"`

	// The version of the Core. The latest version will be displayed as default.
	Version *string `mandatory:"false" json:"version"`
}

func (m PlatformRuntimeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PlatformRuntimeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPlatformRuntimeConfigPlatformRuntimeConfigTypeEnum(string(m.PlatformRuntimeConfigType)); !ok && m.PlatformRuntimeConfigType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformRuntimeConfigType: %s. Supported values are: %s.", m.PlatformRuntimeConfigType, strings.Join(GetPlatformRuntimeConfigPlatformRuntimeConfigTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum Enum with underlying type: string
type PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum string

// Set of constants representing the allowable values for PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum
const (
	PlatformRuntimeConfigPlatformRuntimeConfigTypeAgentPlatform   PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum = "AGENT_PLATFORM"
	PlatformRuntimeConfigPlatformRuntimeConfigTypeFusionReasoning PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum = "FUSION_REASONING"
)

var mappingPlatformRuntimeConfigPlatformRuntimeConfigTypeEnum = map[string]PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum{
	"AGENT_PLATFORM":   PlatformRuntimeConfigPlatformRuntimeConfigTypeAgentPlatform,
	"FUSION_REASONING": PlatformRuntimeConfigPlatformRuntimeConfigTypeFusionReasoning,
}

var mappingPlatformRuntimeConfigPlatformRuntimeConfigTypeEnumLowerCase = map[string]PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum{
	"agent_platform":   PlatformRuntimeConfigPlatformRuntimeConfigTypeAgentPlatform,
	"fusion_reasoning": PlatformRuntimeConfigPlatformRuntimeConfigTypeFusionReasoning,
}

// GetPlatformRuntimeConfigPlatformRuntimeConfigTypeEnumValues Enumerates the set of values for PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum
func GetPlatformRuntimeConfigPlatformRuntimeConfigTypeEnumValues() []PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum {
	values := make([]PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum, 0)
	for _, v := range mappingPlatformRuntimeConfigPlatformRuntimeConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPlatformRuntimeConfigPlatformRuntimeConfigTypeEnumStringValues Enumerates the set of values in String for PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum
func GetPlatformRuntimeConfigPlatformRuntimeConfigTypeEnumStringValues() []string {
	return []string{
		"AGENT_PLATFORM",
		"FUSION_REASONING",
	}
}

// GetMappingPlatformRuntimeConfigPlatformRuntimeConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPlatformRuntimeConfigPlatformRuntimeConfigTypeEnum(val string) (PlatformRuntimeConfigPlatformRuntimeConfigTypeEnum, bool) {
	enum, ok := mappingPlatformRuntimeConfigPlatformRuntimeConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
