// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

import (
	"strings"
)

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeCreating   ActionTypeEnum = "CREATING"
	ActionTypeCreated    ActionTypeEnum = "CREATED"
	ActionTypeUpdated    ActionTypeEnum = "UPDATED"
	ActionTypeDeleted    ActionTypeEnum = "DELETED"
	ActionTypeInProgress ActionTypeEnum = "IN_PROGRESS"
	ActionTypeFailed     ActionTypeEnum = "FAILED"
)

var mappingActionTypeEnum = map[string]ActionTypeEnum{
	"CREATING":    ActionTypeCreating,
	"CREATED":     ActionTypeCreated,
	"UPDATED":     ActionTypeUpdated,
	"DELETED":     ActionTypeDeleted,
	"IN_PROGRESS": ActionTypeInProgress,
	"FAILED":      ActionTypeFailed,
}

var mappingActionTypeEnumLowerCase = map[string]ActionTypeEnum{
	"creating":    ActionTypeCreating,
	"created":     ActionTypeCreated,
	"updated":     ActionTypeUpdated,
	"deleted":     ActionTypeDeleted,
	"in_progress": ActionTypeInProgress,
	"failed":      ActionTypeFailed,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionTypeEnumStringValues Enumerates the set of values in String for ActionTypeEnum
func GetActionTypeEnumStringValues() []string {
	return []string{
		"CREATING",
		"CREATED",
		"UPDATED",
		"DELETED",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypeEnum(val string) (ActionTypeEnum, bool) {
	enum, ok := mappingActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
