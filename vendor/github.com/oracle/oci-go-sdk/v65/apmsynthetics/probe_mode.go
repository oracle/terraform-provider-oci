// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
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

var mappingProbeModeEnumLowerCase = map[string]ProbeModeEnum{
	"sack": ProbeModeSack,
	"syn":  ProbeModeSyn,
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
	enum, ok := mappingProbeModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
