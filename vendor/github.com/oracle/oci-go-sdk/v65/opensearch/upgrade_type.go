// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// UpgradeTypeEnum Enum with underlying type: string
type UpgradeTypeEnum string

// Set of constants representing the allowable values for UpgradeTypeEnum
const (
	UpgradeTypeMinor UpgradeTypeEnum = "MINOR"
	UpgradeTypeMajor UpgradeTypeEnum = "MAJOR"
)

var mappingUpgradeTypeEnum = map[string]UpgradeTypeEnum{
	"MINOR": UpgradeTypeMinor,
	"MAJOR": UpgradeTypeMajor,
}

var mappingUpgradeTypeEnumLowerCase = map[string]UpgradeTypeEnum{
	"minor": UpgradeTypeMinor,
	"major": UpgradeTypeMajor,
}

// GetUpgradeTypeEnumValues Enumerates the set of values for UpgradeTypeEnum
func GetUpgradeTypeEnumValues() []UpgradeTypeEnum {
	values := make([]UpgradeTypeEnum, 0)
	for _, v := range mappingUpgradeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpgradeTypeEnumStringValues Enumerates the set of values in String for UpgradeTypeEnum
func GetUpgradeTypeEnumStringValues() []string {
	return []string{
		"MINOR",
		"MAJOR",
	}
}

// GetMappingUpgradeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpgradeTypeEnum(val string) (UpgradeTypeEnum, bool) {
	enum, ok := mappingUpgradeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
