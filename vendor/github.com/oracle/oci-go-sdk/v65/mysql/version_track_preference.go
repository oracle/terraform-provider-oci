// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// VersionTrackPreferenceEnum Enum with underlying type: string
type VersionTrackPreferenceEnum string

// Set of constants representing the allowable values for VersionTrackPreferenceEnum
const (
	VersionTrackPreferenceLongTermSupport VersionTrackPreferenceEnum = "LONG_TERM_SUPPORT"
	VersionTrackPreferenceInnovation      VersionTrackPreferenceEnum = "INNOVATION"
	VersionTrackPreferenceFollow          VersionTrackPreferenceEnum = "FOLLOW"
)

var mappingVersionTrackPreferenceEnum = map[string]VersionTrackPreferenceEnum{
	"LONG_TERM_SUPPORT": VersionTrackPreferenceLongTermSupport,
	"INNOVATION":        VersionTrackPreferenceInnovation,
	"FOLLOW":            VersionTrackPreferenceFollow,
}

var mappingVersionTrackPreferenceEnumLowerCase = map[string]VersionTrackPreferenceEnum{
	"long_term_support": VersionTrackPreferenceLongTermSupport,
	"innovation":        VersionTrackPreferenceInnovation,
	"follow":            VersionTrackPreferenceFollow,
}

// GetVersionTrackPreferenceEnumValues Enumerates the set of values for VersionTrackPreferenceEnum
func GetVersionTrackPreferenceEnumValues() []VersionTrackPreferenceEnum {
	values := make([]VersionTrackPreferenceEnum, 0)
	for _, v := range mappingVersionTrackPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetVersionTrackPreferenceEnumStringValues Enumerates the set of values in String for VersionTrackPreferenceEnum
func GetVersionTrackPreferenceEnumStringValues() []string {
	return []string{
		"LONG_TERM_SUPPORT",
		"INNOVATION",
		"FOLLOW",
	}
}

// GetMappingVersionTrackPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVersionTrackPreferenceEnum(val string) (VersionTrackPreferenceEnum, bool) {
	enum, ok := mappingVersionTrackPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
