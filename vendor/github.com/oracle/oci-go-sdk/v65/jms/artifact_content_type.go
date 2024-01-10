// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// ArtifactContentTypeEnum Enum with underlying type: string
type ArtifactContentTypeEnum string

// Set of constants representing the allowable values for ArtifactContentTypeEnum
const (
	ArtifactContentTypeJdk       ArtifactContentTypeEnum = "JDK"
	ArtifactContentTypeJre       ArtifactContentTypeEnum = "JRE"
	ArtifactContentTypeServerJre ArtifactContentTypeEnum = "SERVER_JRE"
)

var mappingArtifactContentTypeEnum = map[string]ArtifactContentTypeEnum{
	"JDK":        ArtifactContentTypeJdk,
	"JRE":        ArtifactContentTypeJre,
	"SERVER_JRE": ArtifactContentTypeServerJre,
}

var mappingArtifactContentTypeEnumLowerCase = map[string]ArtifactContentTypeEnum{
	"jdk":        ArtifactContentTypeJdk,
	"jre":        ArtifactContentTypeJre,
	"server_jre": ArtifactContentTypeServerJre,
}

// GetArtifactContentTypeEnumValues Enumerates the set of values for ArtifactContentTypeEnum
func GetArtifactContentTypeEnumValues() []ArtifactContentTypeEnum {
	values := make([]ArtifactContentTypeEnum, 0)
	for _, v := range mappingArtifactContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactContentTypeEnumStringValues Enumerates the set of values in String for ArtifactContentTypeEnum
func GetArtifactContentTypeEnumStringValues() []string {
	return []string{
		"JDK",
		"JRE",
		"SERVER_JRE",
	}
}

// GetMappingArtifactContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactContentTypeEnum(val string) (ArtifactContentTypeEnum, bool) {
	enum, ok := mappingArtifactContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
