// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"strings"
)

// ClusterTypeEnum Enum with underlying type: string
type ClusterTypeEnum string

// Set of constants representing the allowable values for ClusterTypeEnum
const (
	ClusterTypeHosting ClusterTypeEnum = "HOSTING"
)

var mappingClusterTypeEnum = map[string]ClusterTypeEnum{
	"HOSTING": ClusterTypeHosting,
}

var mappingClusterTypeEnumLowerCase = map[string]ClusterTypeEnum{
	"hosting": ClusterTypeHosting,
}

// GetClusterTypeEnumValues Enumerates the set of values for ClusterTypeEnum
func GetClusterTypeEnumValues() []ClusterTypeEnum {
	values := make([]ClusterTypeEnum, 0)
	for _, v := range mappingClusterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterTypeEnumStringValues Enumerates the set of values in String for ClusterTypeEnum
func GetClusterTypeEnumStringValues() []string {
	return []string{
		"HOSTING",
	}
}

// GetMappingClusterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterTypeEnum(val string) (ClusterTypeEnum, bool) {
	enum, ok := mappingClusterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
