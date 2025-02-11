// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// StreamCursorTypesEnum Enum with underlying type: string
type StreamCursorTypesEnum string

// Set of constants representing the allowable values for StreamCursorTypesEnum
const (
	StreamCursorTypesDefault     StreamCursorTypesEnum = "DEFAULT"
	StreamCursorTypesTrimHorizon StreamCursorTypesEnum = "TRIM_HORIZON"
	StreamCursorTypesLatest      StreamCursorTypesEnum = "LATEST"
	StreamCursorTypesAtTime      StreamCursorTypesEnum = "AT_TIME"
)

var mappingStreamCursorTypesEnum = map[string]StreamCursorTypesEnum{
	"DEFAULT":      StreamCursorTypesDefault,
	"TRIM_HORIZON": StreamCursorTypesTrimHorizon,
	"LATEST":       StreamCursorTypesLatest,
	"AT_TIME":      StreamCursorTypesAtTime,
}

var mappingStreamCursorTypesEnumLowerCase = map[string]StreamCursorTypesEnum{
	"default":      StreamCursorTypesDefault,
	"trim_horizon": StreamCursorTypesTrimHorizon,
	"latest":       StreamCursorTypesLatest,
	"at_time":      StreamCursorTypesAtTime,
}

// GetStreamCursorTypesEnumValues Enumerates the set of values for StreamCursorTypesEnum
func GetStreamCursorTypesEnumValues() []StreamCursorTypesEnum {
	values := make([]StreamCursorTypesEnum, 0)
	for _, v := range mappingStreamCursorTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamCursorTypesEnumStringValues Enumerates the set of values in String for StreamCursorTypesEnum
func GetStreamCursorTypesEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"TRIM_HORIZON",
		"LATEST",
		"AT_TIME",
	}
}

// GetMappingStreamCursorTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamCursorTypesEnum(val string) (StreamCursorTypesEnum, bool) {
	enum, ok := mappingStreamCursorTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
