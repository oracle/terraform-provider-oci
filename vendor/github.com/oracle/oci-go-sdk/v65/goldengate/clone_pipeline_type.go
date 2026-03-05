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

// ClonePipelineTypeEnum Enum with underlying type: string
type ClonePipelineTypeEnum string

// Set of constants representing the allowable values for ClonePipelineTypeEnum
const (
	ClonePipelineTypeDefault ClonePipelineTypeEnum = "DEFAULT"
)

var mappingClonePipelineTypeEnum = map[string]ClonePipelineTypeEnum{
	"DEFAULT": ClonePipelineTypeDefault,
}

var mappingClonePipelineTypeEnumLowerCase = map[string]ClonePipelineTypeEnum{
	"default": ClonePipelineTypeDefault,
}

// GetClonePipelineTypeEnumValues Enumerates the set of values for ClonePipelineTypeEnum
func GetClonePipelineTypeEnumValues() []ClonePipelineTypeEnum {
	values := make([]ClonePipelineTypeEnum, 0)
	for _, v := range mappingClonePipelineTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetClonePipelineTypeEnumStringValues Enumerates the set of values in String for ClonePipelineTypeEnum
func GetClonePipelineTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingClonePipelineTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClonePipelineTypeEnum(val string) (ClonePipelineTypeEnum, bool) {
	enum, ok := mappingClonePipelineTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
