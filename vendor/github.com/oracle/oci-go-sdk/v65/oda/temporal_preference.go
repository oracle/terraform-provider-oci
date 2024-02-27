// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// TemporalPreferenceEnum Enum with underlying type: string
type TemporalPreferenceEnum string

// Set of constants representing the allowable values for TemporalPreferenceEnum
const (
	TemporalPreferencePast    TemporalPreferenceEnum = "PAST"
	TemporalPreferenceFuture  TemporalPreferenceEnum = "FUTURE"
	TemporalPreferenceNearest TemporalPreferenceEnum = "NEAREST"
)

var mappingTemporalPreferenceEnum = map[string]TemporalPreferenceEnum{
	"PAST":    TemporalPreferencePast,
	"FUTURE":  TemporalPreferenceFuture,
	"NEAREST": TemporalPreferenceNearest,
}

var mappingTemporalPreferenceEnumLowerCase = map[string]TemporalPreferenceEnum{
	"past":    TemporalPreferencePast,
	"future":  TemporalPreferenceFuture,
	"nearest": TemporalPreferenceNearest,
}

// GetTemporalPreferenceEnumValues Enumerates the set of values for TemporalPreferenceEnum
func GetTemporalPreferenceEnumValues() []TemporalPreferenceEnum {
	values := make([]TemporalPreferenceEnum, 0)
	for _, v := range mappingTemporalPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetTemporalPreferenceEnumStringValues Enumerates the set of values in String for TemporalPreferenceEnum
func GetTemporalPreferenceEnumStringValues() []string {
	return []string{
		"PAST",
		"FUTURE",
		"NEAREST",
	}
}

// GetMappingTemporalPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTemporalPreferenceEnum(val string) (TemporalPreferenceEnum, bool) {
	enum, ok := mappingTemporalPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
