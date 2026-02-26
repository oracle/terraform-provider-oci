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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnvironmentVariable The environment variables for the Hosted Application
type EnvironmentVariable struct {

	// Name of the environment variable.
	Name *string `mandatory:"true" json:"name"`

	// Type of the environment variable (PLAINTEXT or HASHED, no default value).
	Type EnvironmentVariableTypeEnum `mandatory:"true" json:"type"`

	// Value of the environment variable.
	Value *interface{} `mandatory:"true" json:"value"`
}

func (m EnvironmentVariable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnvironmentVariable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnvironmentVariableTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetEnvironmentVariableTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnvironmentVariableTypeEnum Enum with underlying type: string
type EnvironmentVariableTypeEnum string

// Set of constants representing the allowable values for EnvironmentVariableTypeEnum
const (
	EnvironmentVariableTypePlaintext EnvironmentVariableTypeEnum = "PLAINTEXT"
	EnvironmentVariableTypeVault     EnvironmentVariableTypeEnum = "VAULT"
)

var mappingEnvironmentVariableTypeEnum = map[string]EnvironmentVariableTypeEnum{
	"PLAINTEXT": EnvironmentVariableTypePlaintext,
	"VAULT":     EnvironmentVariableTypeVault,
}

var mappingEnvironmentVariableTypeEnumLowerCase = map[string]EnvironmentVariableTypeEnum{
	"plaintext": EnvironmentVariableTypePlaintext,
	"vault":     EnvironmentVariableTypeVault,
}

// GetEnvironmentVariableTypeEnumValues Enumerates the set of values for EnvironmentVariableTypeEnum
func GetEnvironmentVariableTypeEnumValues() []EnvironmentVariableTypeEnum {
	values := make([]EnvironmentVariableTypeEnum, 0)
	for _, v := range mappingEnvironmentVariableTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEnvironmentVariableTypeEnumStringValues Enumerates the set of values in String for EnvironmentVariableTypeEnum
func GetEnvironmentVariableTypeEnumStringValues() []string {
	return []string{
		"PLAINTEXT",
		"VAULT",
	}
}

// GetMappingEnvironmentVariableTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnvironmentVariableTypeEnum(val string) (EnvironmentVariableTypeEnum, bool) {
	enum, ok := mappingEnvironmentVariableTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
