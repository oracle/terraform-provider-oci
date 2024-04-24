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

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateClusterPlacementGroup OperationTypeEnum = "CREATE_CLUSTER_PLACEMENT_GROUP"
	OperationTypeUpdateClusterPlacementGroup OperationTypeEnum = "UPDATE_CLUSTER_PLACEMENT_GROUP"
	OperationTypeDeleteClusterPlacementGroup OperationTypeEnum = "DELETE_CLUSTER_PLACEMENT_GROUP"
	OperationTypeMoveClusterPlacementGroup   OperationTypeEnum = "MOVE_CLUSTER_PLACEMENT_GROUP"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_CLUSTER_PLACEMENT_GROUP": OperationTypeCreateClusterPlacementGroup,
	"UPDATE_CLUSTER_PLACEMENT_GROUP": OperationTypeUpdateClusterPlacementGroup,
	"DELETE_CLUSTER_PLACEMENT_GROUP": OperationTypeDeleteClusterPlacementGroup,
	"MOVE_CLUSTER_PLACEMENT_GROUP":   OperationTypeMoveClusterPlacementGroup,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_cluster_placement_group": OperationTypeCreateClusterPlacementGroup,
	"update_cluster_placement_group": OperationTypeUpdateClusterPlacementGroup,
	"delete_cluster_placement_group": OperationTypeDeleteClusterPlacementGroup,
	"move_cluster_placement_group":   OperationTypeMoveClusterPlacementGroup,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_CLUSTER_PLACEMENT_GROUP",
		"UPDATE_CLUSTER_PLACEMENT_GROUP",
		"DELETE_CLUSTER_PLACEMENT_GROUP",
		"MOVE_CLUSTER_PLACEMENT_GROUP",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
