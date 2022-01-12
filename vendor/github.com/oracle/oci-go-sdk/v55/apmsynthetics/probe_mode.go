// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

// ProbeModeEnum Enum with underlying type: string
type ProbeModeEnum string

// Set of constants representing the allowable values for ProbeModeEnum
const (
	ProbeModeSack ProbeModeEnum = "SACK"
	ProbeModeSyn  ProbeModeEnum = "SYN"
)

var mappingProbeMode = map[string]ProbeModeEnum{
	"SACK": ProbeModeSack,
	"SYN":  ProbeModeSyn,
}

// GetProbeModeEnumValues Enumerates the set of values for ProbeModeEnum
func GetProbeModeEnumValues() []ProbeModeEnum {
	values := make([]ProbeModeEnum, 0)
	for _, v := range mappingProbeMode {
		values = append(values, v)
	}
	return values
}
