// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ArtifactSourceTypeEnum Enum with underlying type: string
type ArtifactSourceTypeEnum string

// Set of constants representing the allowable values for ArtifactSourceTypeEnum
const (
	ArtifactSourceTypeOracleObjectStorage ArtifactSourceTypeEnum = "ORACLE_OBJECT_STORAGE"
)

var mappingArtifactSourceTypeEnum = map[string]ArtifactSourceTypeEnum{
	"ORACLE_OBJECT_STORAGE": ArtifactSourceTypeOracleObjectStorage,
}

var mappingArtifactSourceTypeEnumLowerCase = map[string]ArtifactSourceTypeEnum{
	"oracle_object_storage": ArtifactSourceTypeOracleObjectStorage,
}

// GetArtifactSourceTypeEnumValues Enumerates the set of values for ArtifactSourceTypeEnum
func GetArtifactSourceTypeEnumValues() []ArtifactSourceTypeEnum {
	values := make([]ArtifactSourceTypeEnum, 0)
	for _, v := range mappingArtifactSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactSourceTypeEnumStringValues Enumerates the set of values in String for ArtifactSourceTypeEnum
func GetArtifactSourceTypeEnumStringValues() []string {
	return []string{
		"ORACLE_OBJECT_STORAGE",
	}
}

// GetMappingArtifactSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactSourceTypeEnum(val string) (ArtifactSourceTypeEnum, bool) {
	enum, ok := mappingArtifactSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
