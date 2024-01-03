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

// ReleaseTypeEnum Enum with underlying type: string
type ReleaseTypeEnum string

// Set of constants representing the allowable values for ReleaseTypeEnum
const (
	ReleaseTypeMajor  ReleaseTypeEnum = "MAJOR"
	ReleaseTypeBundle ReleaseTypeEnum = "BUNDLE"
	ReleaseTypeMinor  ReleaseTypeEnum = "MINOR"
)

var mappingReleaseTypeEnum = map[string]ReleaseTypeEnum{
	"MAJOR":  ReleaseTypeMajor,
	"BUNDLE": ReleaseTypeBundle,
	"MINOR":  ReleaseTypeMinor,
}

var mappingReleaseTypeEnumLowerCase = map[string]ReleaseTypeEnum{
	"major":  ReleaseTypeMajor,
	"bundle": ReleaseTypeBundle,
	"minor":  ReleaseTypeMinor,
}

// GetReleaseTypeEnumValues Enumerates the set of values for ReleaseTypeEnum
func GetReleaseTypeEnumValues() []ReleaseTypeEnum {
	values := make([]ReleaseTypeEnum, 0)
	for _, v := range mappingReleaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReleaseTypeEnumStringValues Enumerates the set of values in String for ReleaseTypeEnum
func GetReleaseTypeEnumStringValues() []string {
	return []string{
		"MAJOR",
		"BUNDLE",
		"MINOR",
	}
}

// GetMappingReleaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReleaseTypeEnum(val string) (ReleaseTypeEnum, bool) {
	enum, ok := mappingReleaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
