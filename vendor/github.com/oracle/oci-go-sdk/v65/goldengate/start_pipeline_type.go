// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// StartPipelineTypeEnum Enum with underlying type: string
type StartPipelineTypeEnum string

// Set of constants representing the allowable values for StartPipelineTypeEnum
const (
	StartPipelineTypeDefault StartPipelineTypeEnum = "DEFAULT"
)

var mappingStartPipelineTypeEnum = map[string]StartPipelineTypeEnum{
	"DEFAULT": StartPipelineTypeDefault,
}

var mappingStartPipelineTypeEnumLowerCase = map[string]StartPipelineTypeEnum{
	"default": StartPipelineTypeDefault,
}

// GetStartPipelineTypeEnumValues Enumerates the set of values for StartPipelineTypeEnum
func GetStartPipelineTypeEnumValues() []StartPipelineTypeEnum {
	values := make([]StartPipelineTypeEnum, 0)
	for _, v := range mappingStartPipelineTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStartPipelineTypeEnumStringValues Enumerates the set of values in String for StartPipelineTypeEnum
func GetStartPipelineTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingStartPipelineTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStartPipelineTypeEnum(val string) (StartPipelineTypeEnum, bool) {
	enum, ok := mappingStartPipelineTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
