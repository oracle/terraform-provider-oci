// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// AiModelAuthTypeEnum Enum with underlying type: string
type AiModelAuthTypeEnum string

// Set of constants representing the allowable values for AiModelAuthTypeEnum
const (
	AiModelAuthTypeOciGenAi AiModelAuthTypeEnum = "OCI_GEN_AI"
	AiModelAuthTypeApiKey   AiModelAuthTypeEnum = "API_KEY"
)

var mappingAiModelAuthTypeEnum = map[string]AiModelAuthTypeEnum{
	"OCI_GEN_AI": AiModelAuthTypeOciGenAi,
	"API_KEY":    AiModelAuthTypeApiKey,
}

var mappingAiModelAuthTypeEnumLowerCase = map[string]AiModelAuthTypeEnum{
	"oci_gen_ai": AiModelAuthTypeOciGenAi,
	"api_key":    AiModelAuthTypeApiKey,
}

// GetAiModelAuthTypeEnumValues Enumerates the set of values for AiModelAuthTypeEnum
func GetAiModelAuthTypeEnumValues() []AiModelAuthTypeEnum {
	values := make([]AiModelAuthTypeEnum, 0)
	for _, v := range mappingAiModelAuthTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAiModelAuthTypeEnumStringValues Enumerates the set of values in String for AiModelAuthTypeEnum
func GetAiModelAuthTypeEnumStringValues() []string {
	return []string{
		"OCI_GEN_AI",
		"API_KEY",
	}
}

// GetMappingAiModelAuthTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAiModelAuthTypeEnum(val string) (AiModelAuthTypeEnum, bool) {
	enum, ok := mappingAiModelAuthTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
