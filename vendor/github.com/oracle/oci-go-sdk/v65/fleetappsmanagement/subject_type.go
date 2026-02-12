// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// SubjectTypeEnum Enum with underlying type: string
type SubjectTypeEnum string

// Set of constants representing the allowable values for SubjectTypeEnum
const (
	SubjectTypePatchGroup         SubjectTypeEnum = "PATCH_GROUP"
	SubjectTypeRecommendedPatches SubjectTypeEnum = "RECOMMENDED_PATCHES"
	SubjectTypeOneOff             SubjectTypeEnum = "ONE_OFF"
)

var mappingSubjectTypeEnum = map[string]SubjectTypeEnum{
	"PATCH_GROUP":         SubjectTypePatchGroup,
	"RECOMMENDED_PATCHES": SubjectTypeRecommendedPatches,
	"ONE_OFF":             SubjectTypeOneOff,
}

var mappingSubjectTypeEnumLowerCase = map[string]SubjectTypeEnum{
	"patch_group":         SubjectTypePatchGroup,
	"recommended_patches": SubjectTypeRecommendedPatches,
	"one_off":             SubjectTypeOneOff,
}

// GetSubjectTypeEnumValues Enumerates the set of values for SubjectTypeEnum
func GetSubjectTypeEnumValues() []SubjectTypeEnum {
	values := make([]SubjectTypeEnum, 0)
	for _, v := range mappingSubjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubjectTypeEnumStringValues Enumerates the set of values in String for SubjectTypeEnum
func GetSubjectTypeEnumStringValues() []string {
	return []string{
		"PATCH_GROUP",
		"RECOMMENDED_PATCHES",
		"ONE_OFF",
	}
}

// GetMappingSubjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubjectTypeEnum(val string) (SubjectTypeEnum, bool) {
	enum, ok := mappingSubjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
