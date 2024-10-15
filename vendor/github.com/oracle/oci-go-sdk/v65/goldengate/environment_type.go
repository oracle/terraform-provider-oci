// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// EnvironmentTypeEnum Enum with underlying type: string
type EnvironmentTypeEnum string

// Set of constants representing the allowable values for EnvironmentTypeEnum
const (
	EnvironmentTypeProduction           EnvironmentTypeEnum = "PRODUCTION"
	EnvironmentTypeDevelopmentOrTesting EnvironmentTypeEnum = "DEVELOPMENT_OR_TESTING"
)

var mappingEnvironmentTypeEnum = map[string]EnvironmentTypeEnum{
	"PRODUCTION":             EnvironmentTypeProduction,
	"DEVELOPMENT_OR_TESTING": EnvironmentTypeDevelopmentOrTesting,
}

var mappingEnvironmentTypeEnumLowerCase = map[string]EnvironmentTypeEnum{
	"production":             EnvironmentTypeProduction,
	"development_or_testing": EnvironmentTypeDevelopmentOrTesting,
}

// GetEnvironmentTypeEnumValues Enumerates the set of values for EnvironmentTypeEnum
func GetEnvironmentTypeEnumValues() []EnvironmentTypeEnum {
	values := make([]EnvironmentTypeEnum, 0)
	for _, v := range mappingEnvironmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEnvironmentTypeEnumStringValues Enumerates the set of values in String for EnvironmentTypeEnum
func GetEnvironmentTypeEnumStringValues() []string {
	return []string{
		"PRODUCTION",
		"DEVELOPMENT_OR_TESTING",
	}
}

// GetMappingEnvironmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnvironmentTypeEnum(val string) (EnvironmentTypeEnum, bool) {
	enum, ok := mappingEnvironmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
