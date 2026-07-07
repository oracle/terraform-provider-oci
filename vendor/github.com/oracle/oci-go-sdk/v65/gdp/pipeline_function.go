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

// PipelineFunctionEnum Enum with underlying type: string
type PipelineFunctionEnum string

// Set of constants representing the allowable values for PipelineFunctionEnum
const (
	PipelineFunctionP2P     PipelineFunctionEnum = "P2P"
	PipelineFunctionMp      PipelineFunctionEnum = "MP"
	PipelineFunctionService PipelineFunctionEnum = "SERVICE"
)

var mappingPipelineFunctionEnum = map[string]PipelineFunctionEnum{
	"P2P":     PipelineFunctionP2P,
	"MP":      PipelineFunctionMp,
	"SERVICE": PipelineFunctionService,
}

var mappingPipelineFunctionEnumLowerCase = map[string]PipelineFunctionEnum{
	"p2p":     PipelineFunctionP2P,
	"mp":      PipelineFunctionMp,
	"service": PipelineFunctionService,
}

// GetPipelineFunctionEnumValues Enumerates the set of values for PipelineFunctionEnum
func GetPipelineFunctionEnumValues() []PipelineFunctionEnum {
	values := make([]PipelineFunctionEnum, 0)
	for _, v := range mappingPipelineFunctionEnum {
		values = append(values, v)
	}
	return values
}

// GetPipelineFunctionEnumStringValues Enumerates the set of values in String for PipelineFunctionEnum
func GetPipelineFunctionEnumStringValues() []string {
	return []string{
		"P2P",
		"MP",
		"SERVICE",
	}
}

// GetMappingPipelineFunctionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPipelineFunctionEnum(val string) (PipelineFunctionEnum, bool) {
	enum, ok := mappingPipelineFunctionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
