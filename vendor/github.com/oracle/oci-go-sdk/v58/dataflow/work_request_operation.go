// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// WorkRequestOperationEnum Enum with underlying type: string
type WorkRequestOperationEnum string

// Set of constants representing the allowable values for WorkRequestOperationEnum
const (
	WorkRequestOperationCreatePrivateEndpoint WorkRequestOperationEnum = "CREATE_PRIVATE_ENDPOINT"
	WorkRequestOperationUpdatePrivateEndpoint WorkRequestOperationEnum = "UPDATE_PRIVATE_ENDPOINT"
	WorkRequestOperationDeletePrivateEndpoint WorkRequestOperationEnum = "DELETE_PRIVATE_ENDPOINT"
	WorkRequestOperationMovePrivateEndpoint   WorkRequestOperationEnum = "MOVE_PRIVATE_ENDPOINT"
)

var mappingWorkRequestOperationEnum = map[string]WorkRequestOperationEnum{
	"CREATE_PRIVATE_ENDPOINT": WorkRequestOperationCreatePrivateEndpoint,
	"UPDATE_PRIVATE_ENDPOINT": WorkRequestOperationUpdatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT": WorkRequestOperationDeletePrivateEndpoint,
	"MOVE_PRIVATE_ENDPOINT":   WorkRequestOperationMovePrivateEndpoint,
}

// GetWorkRequestOperationEnumValues Enumerates the set of values for WorkRequestOperationEnum
func GetWorkRequestOperationEnumValues() []WorkRequestOperationEnum {
	values := make([]WorkRequestOperationEnum, 0)
	for _, v := range mappingWorkRequestOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationEnumStringValues Enumerates the set of values in String for WorkRequestOperationEnum
func GetWorkRequestOperationEnumStringValues() []string {
	return []string{
		"CREATE_PRIVATE_ENDPOINT",
		"UPDATE_PRIVATE_ENDPOINT",
		"DELETE_PRIVATE_ENDPOINT",
		"MOVE_PRIVATE_ENDPOINT",
	}
}

// GetMappingWorkRequestOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationEnum(val string) (WorkRequestOperationEnum, bool) {
	mappingWorkRequestOperationEnumIgnoreCase := make(map[string]WorkRequestOperationEnum)
	for k, v := range mappingWorkRequestOperationEnum {
		mappingWorkRequestOperationEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestOperationEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
