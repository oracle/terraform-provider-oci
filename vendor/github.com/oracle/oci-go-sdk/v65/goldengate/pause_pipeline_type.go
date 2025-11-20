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

// PausePipelineTypeEnum Enum with underlying type: string
type PausePipelineTypeEnum string

// Set of constants representing the allowable values for PausePipelineTypeEnum
const (
	PausePipelineTypeDefault PausePipelineTypeEnum = "DEFAULT"
)

var mappingPausePipelineTypeEnum = map[string]PausePipelineTypeEnum{
	"DEFAULT": PausePipelineTypeDefault,
}

var mappingPausePipelineTypeEnumLowerCase = map[string]PausePipelineTypeEnum{
	"default": PausePipelineTypeDefault,
}

// GetPausePipelineTypeEnumValues Enumerates the set of values for PausePipelineTypeEnum
func GetPausePipelineTypeEnumValues() []PausePipelineTypeEnum {
	values := make([]PausePipelineTypeEnum, 0)
	for _, v := range mappingPausePipelineTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPausePipelineTypeEnumStringValues Enumerates the set of values in String for PausePipelineTypeEnum
func GetPausePipelineTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingPausePipelineTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPausePipelineTypeEnum(val string) (PausePipelineTypeEnum, bool) {
	enum, ok := mappingPausePipelineTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
