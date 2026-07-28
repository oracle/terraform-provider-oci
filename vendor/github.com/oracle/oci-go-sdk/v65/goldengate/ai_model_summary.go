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

// AiModelSummary Summary details of an AI model supported by a provider.
type AiModelSummary struct {

	// The identifier of the AI model offered by a provider.
	Key *string `mandatory:"true" json:"key"`

	// AI Provider type used by the AI Model Connection.
	ProviderType AiModelSummaryProviderTypeEnum `mandatory:"true" json:"providerType"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"true" json:"description"`
}

func (m AiModelSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AiModelSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAiModelSummaryProviderTypeEnum(string(m.ProviderType)); !ok && m.ProviderType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProviderType: %s. Supported values are: %s.", m.ProviderType, strings.Join(GetAiModelSummaryProviderTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AiModelSummaryProviderTypeEnum Enum with underlying type: string
type AiModelSummaryProviderTypeEnum string

// Set of constants representing the allowable values for AiModelSummaryProviderTypeEnum
const (
	AiModelSummaryProviderTypeOciGenerativeAi AiModelSummaryProviderTypeEnum = "OCI_GENERATIVE_AI"
	AiModelSummaryProviderTypeGemini          AiModelSummaryProviderTypeEnum = "GEMINI"
	AiModelSummaryProviderTypeOpenAi          AiModelSummaryProviderTypeEnum = "OPEN_AI"
	AiModelSummaryProviderTypeVoyageAi        AiModelSummaryProviderTypeEnum = "VOYAGE_AI"
)

var mappingAiModelSummaryProviderTypeEnum = map[string]AiModelSummaryProviderTypeEnum{
	"OCI_GENERATIVE_AI": AiModelSummaryProviderTypeOciGenerativeAi,
	"GEMINI":            AiModelSummaryProviderTypeGemini,
	"OPEN_AI":           AiModelSummaryProviderTypeOpenAi,
	"VOYAGE_AI":         AiModelSummaryProviderTypeVoyageAi,
}

var mappingAiModelSummaryProviderTypeEnumLowerCase = map[string]AiModelSummaryProviderTypeEnum{
	"oci_generative_ai": AiModelSummaryProviderTypeOciGenerativeAi,
	"gemini":            AiModelSummaryProviderTypeGemini,
	"open_ai":           AiModelSummaryProviderTypeOpenAi,
	"voyage_ai":         AiModelSummaryProviderTypeVoyageAi,
}

// GetAiModelSummaryProviderTypeEnumValues Enumerates the set of values for AiModelSummaryProviderTypeEnum
func GetAiModelSummaryProviderTypeEnumValues() []AiModelSummaryProviderTypeEnum {
	values := make([]AiModelSummaryProviderTypeEnum, 0)
	for _, v := range mappingAiModelSummaryProviderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAiModelSummaryProviderTypeEnumStringValues Enumerates the set of values in String for AiModelSummaryProviderTypeEnum
func GetAiModelSummaryProviderTypeEnumStringValues() []string {
	return []string{
		"OCI_GENERATIVE_AI",
		"GEMINI",
		"OPEN_AI",
		"VOYAGE_AI",
	}
}

// GetMappingAiModelSummaryProviderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAiModelSummaryProviderTypeEnum(val string) (AiModelSummaryProviderTypeEnum, bool) {
	enum, ok := mappingAiModelSummaryProviderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
