// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cluster Placement Groups API
//
// API for managing cluster placement groups.
//

package clusterplacementgroups

import (
	"strings"
)

// ClusterPlacementGroupTypeEnum Enum with underlying type: string
type ClusterPlacementGroupTypeEnum string

// Set of constants representing the allowable values for ClusterPlacementGroupTypeEnum
const (
	ClusterPlacementGroupTypeStandard ClusterPlacementGroupTypeEnum = "STANDARD"
)

var mappingClusterPlacementGroupTypeEnum = map[string]ClusterPlacementGroupTypeEnum{
	"STANDARD": ClusterPlacementGroupTypeStandard,
}

var mappingClusterPlacementGroupTypeEnumLowerCase = map[string]ClusterPlacementGroupTypeEnum{
	"standard": ClusterPlacementGroupTypeStandard,
}

// GetClusterPlacementGroupTypeEnumValues Enumerates the set of values for ClusterPlacementGroupTypeEnum
func GetClusterPlacementGroupTypeEnumValues() []ClusterPlacementGroupTypeEnum {
	values := make([]ClusterPlacementGroupTypeEnum, 0)
	for _, v := range mappingClusterPlacementGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterPlacementGroupTypeEnumStringValues Enumerates the set of values in String for ClusterPlacementGroupTypeEnum
func GetClusterPlacementGroupTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
	}
}

// GetMappingClusterPlacementGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterPlacementGroupTypeEnum(val string) (ClusterPlacementGroupTypeEnum, bool) {
	enum, ok := mappingClusterPlacementGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
