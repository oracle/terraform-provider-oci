// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"strings"
)

// ProbeModeEnum Enum with underlying type: string
type ProbeModeEnum string

// Set of constants representing the allowable values for ProbeModeEnum
const (
	ProbeModeSack ProbeModeEnum = "SACK"
	ProbeModeSyn  ProbeModeEnum = "SYN"
)

var mappingProbeModeEnum = map[string]ProbeModeEnum{
	"SACK": ProbeModeSack,
	"SYN":  ProbeModeSyn,
}

// GetProbeModeEnumValues Enumerates the set of values for ProbeModeEnum
func GetProbeModeEnumValues() []ProbeModeEnum {
	values := make([]ProbeModeEnum, 0)
	for _, v := range mappingProbeModeEnum {
		values = append(values, v)
	}
	return values
}

// GetProbeModeEnumStringValues Enumerates the set of values in String for ProbeModeEnum
func GetProbeModeEnumStringValues() []string {
	return []string{
		"SACK",
		"SYN",
	}
}

// GetMappingProbeModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProbeModeEnum(val string) (ProbeModeEnum, bool) {
	mappingProbeModeEnumIgnoreCase := make(map[string]ProbeModeEnum)
	for k, v := range mappingProbeModeEnum {
		mappingProbeModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingProbeModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
