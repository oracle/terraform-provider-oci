// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AiProviderSummary Summary details of an AI provider and its supported models.
type AiProviderSummary struct {

	// AI Provider type used by the AI Model Connection.
	ProviderType AiProviderSummaryProviderTypeEnum `mandatory:"true" json:"providerType"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"true" json:"description"`

	// Authentication types supported by the AI provider.
	AuthType []AiModelAuthTypeEnum `mandatory:"true" json:"authType"`

	// Default base URL for the AI provider.
	DefaultBaseUrl *string `mandatory:"true" json:"defaultBaseUrl"`

	// List of AI models supported by this provider, when available. This
	// field is null when the provider's models can be retrieved only after
	// supplying additional context. For example, OCI_GENERATIVE_AI model
	// availability may vary by region.
	Models []AiModelSummary `mandatory:"false" json:"models"`
}

func (m AiProviderSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AiProviderSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAiProviderSummaryProviderTypeEnum(string(m.ProviderType)); !ok && m.ProviderType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProviderType: %s. Supported values are: %s.", m.ProviderType, strings.Join(GetAiProviderSummaryProviderTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AiProviderSummaryProviderTypeEnum Enum with underlying type: string
type AiProviderSummaryProviderTypeEnum string

// Set of constants representing the allowable values for AiProviderSummaryProviderTypeEnum
const (
	AiProviderSummaryProviderTypeOciGenerativeAi AiProviderSummaryProviderTypeEnum = "OCI_GENERATIVE_AI"
	AiProviderSummaryProviderTypeGemini          AiProviderSummaryProviderTypeEnum = "GEMINI"
	AiProviderSummaryProviderTypeOpenAi          AiProviderSummaryProviderTypeEnum = "OPEN_AI"
	AiProviderSummaryProviderTypeVoyageAi        AiProviderSummaryProviderTypeEnum = "VOYAGE_AI"
)

var mappingAiProviderSummaryProviderTypeEnum = map[string]AiProviderSummaryProviderTypeEnum{
	"OCI_GENERATIVE_AI": AiProviderSummaryProviderTypeOciGenerativeAi,
	"GEMINI":            AiProviderSummaryProviderTypeGemini,
	"OPEN_AI":           AiProviderSummaryProviderTypeOpenAi,
	"VOYAGE_AI":         AiProviderSummaryProviderTypeVoyageAi,
}

var mappingAiProviderSummaryProviderTypeEnumLowerCase = map[string]AiProviderSummaryProviderTypeEnum{
	"oci_generative_ai": AiProviderSummaryProviderTypeOciGenerativeAi,
	"gemini":            AiProviderSummaryProviderTypeGemini,
	"open_ai":           AiProviderSummaryProviderTypeOpenAi,
	"voyage_ai":         AiProviderSummaryProviderTypeVoyageAi,
}

// GetAiProviderSummaryProviderTypeEnumValues Enumerates the set of values for AiProviderSummaryProviderTypeEnum
func GetAiProviderSummaryProviderTypeEnumValues() []AiProviderSummaryProviderTypeEnum {
	values := make([]AiProviderSummaryProviderTypeEnum, 0)
	for _, v := range mappingAiProviderSummaryProviderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAiProviderSummaryProviderTypeEnumStringValues Enumerates the set of values in String for AiProviderSummaryProviderTypeEnum
func GetAiProviderSummaryProviderTypeEnumStringValues() []string {
	return []string{
		"OCI_GENERATIVE_AI",
		"GEMINI",
		"OPEN_AI",
		"VOYAGE_AI",
	}
}

// GetMappingAiProviderSummaryProviderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAiProviderSummaryProviderTypeEnum(val string) (AiProviderSummaryProviderTypeEnum, bool) {
	enum, ok := mappingAiProviderSummaryProviderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
