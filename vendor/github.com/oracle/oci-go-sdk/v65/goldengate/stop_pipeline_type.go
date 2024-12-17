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

// StopPipelineTypeEnum Enum with underlying type: string
type StopPipelineTypeEnum string

// Set of constants representing the allowable values for StopPipelineTypeEnum
const (
	StopPipelineTypeDefault StopPipelineTypeEnum = "DEFAULT"
)

var mappingStopPipelineTypeEnum = map[string]StopPipelineTypeEnum{
	"DEFAULT": StopPipelineTypeDefault,
}

var mappingStopPipelineTypeEnumLowerCase = map[string]StopPipelineTypeEnum{
	"default": StopPipelineTypeDefault,
}

// GetStopPipelineTypeEnumValues Enumerates the set of values for StopPipelineTypeEnum
func GetStopPipelineTypeEnumValues() []StopPipelineTypeEnum {
	values := make([]StopPipelineTypeEnum, 0)
	for _, v := range mappingStopPipelineTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStopPipelineTypeEnumStringValues Enumerates the set of values in String for StopPipelineTypeEnum
func GetStopPipelineTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingStopPipelineTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStopPipelineTypeEnum(val string) (StopPipelineTypeEnum, bool) {
	enum, ok := mappingStopPipelineTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
