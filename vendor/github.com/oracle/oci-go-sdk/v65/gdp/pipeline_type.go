// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Guarded Data Pipelines API
//
// Use Guarded Data Pipelines to facilitate data transfer between different security domains. The service provides physical, network, and logistical isolation between security domains, malware and vulnerability scanning, auditing, and logging, with deep content inspection capabilities.
//

package gdp

import (
	"strings"
)

// PipelineTypeEnum Enum with underlying type: string
type PipelineTypeEnum string

// Set of constants representing the allowable values for PipelineTypeEnum
const (
	PipelineTypeSender   PipelineTypeEnum = "SENDER"
	PipelineTypeReceiver PipelineTypeEnum = "RECEIVER"
)

var mappingPipelineTypeEnum = map[string]PipelineTypeEnum{
	"SENDER":   PipelineTypeSender,
	"RECEIVER": PipelineTypeReceiver,
}

var mappingPipelineTypeEnumLowerCase = map[string]PipelineTypeEnum{
	"sender":   PipelineTypeSender,
	"receiver": PipelineTypeReceiver,
}

// GetPipelineTypeEnumValues Enumerates the set of values for PipelineTypeEnum
func GetPipelineTypeEnumValues() []PipelineTypeEnum {
	values := make([]PipelineTypeEnum, 0)
	for _, v := range mappingPipelineTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineTypeEnumStringValues Enumerates the set of values in String for PipelineTypeEnum
func GetPipelineTypeEnumStringValues() []string {
	return []string{
		"SENDER",
		"RECEIVER",
	}
}

// GetMappingPipelineTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineTypeEnum(val string) (PipelineTypeEnum, bool) {
	enum, ok := mappingPipelineTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
