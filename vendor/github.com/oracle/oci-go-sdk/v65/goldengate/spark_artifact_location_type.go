// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SparkArtifactLocationTypeEnum Enum with underlying type: string
type SparkArtifactLocationTypeEnum string

// Set of constants representing the allowable values for SparkArtifactLocationTypeEnum
const (
	SparkArtifactLocationTypeOciObjectStorage SparkArtifactLocationTypeEnum = "OCI_OBJECT_STORAGE"
)

var mappingSparkArtifactLocationTypeEnum = map[string]SparkArtifactLocationTypeEnum{
	"OCI_OBJECT_STORAGE": SparkArtifactLocationTypeOciObjectStorage,
}

var mappingSparkArtifactLocationTypeEnumLowerCase = map[string]SparkArtifactLocationTypeEnum{
	"oci_object_storage": SparkArtifactLocationTypeOciObjectStorage,
}

// GetSparkArtifactLocationTypeEnumValues Enumerates the set of values for SparkArtifactLocationTypeEnum
func GetSparkArtifactLocationTypeEnumValues() []SparkArtifactLocationTypeEnum {
	values := make([]SparkArtifactLocationTypeEnum, 0)
	for _, v := range mappingSparkArtifactLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkArtifactLocationTypeEnumStringValues Enumerates the set of values in String for SparkArtifactLocationTypeEnum
func GetSparkArtifactLocationTypeEnumStringValues() []string {
	return []string{
		"OCI_OBJECT_STORAGE",
	}
}

// GetMappingSparkArtifactLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkArtifactLocationTypeEnum(val string) (SparkArtifactLocationTypeEnum, bool) {
	enum, ok := mappingSparkArtifactLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
