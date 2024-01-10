// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.cloud.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateQueue OperationTypeEnum = "CREATE_QUEUE"
	OperationTypeUpdateQueue OperationTypeEnum = "UPDATE_QUEUE"
	OperationTypeDeleteQueue OperationTypeEnum = "DELETE_QUEUE"
	OperationTypeMoveQueue   OperationTypeEnum = "MOVE_QUEUE"
	OperationTypePurgeQueue  OperationTypeEnum = "PURGE_QUEUE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_QUEUE": OperationTypeCreateQueue,
	"UPDATE_QUEUE": OperationTypeUpdateQueue,
	"DELETE_QUEUE": OperationTypeDeleteQueue,
	"MOVE_QUEUE":   OperationTypeMoveQueue,
	"PURGE_QUEUE":  OperationTypePurgeQueue,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_queue": OperationTypeCreateQueue,
	"update_queue": OperationTypeUpdateQueue,
	"delete_queue": OperationTypeDeleteQueue,
	"move_queue":   OperationTypeMoveQueue,
	"purge_queue":  OperationTypePurgeQueue,
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
		"CREATE_QUEUE",
		"UPDATE_QUEUE",
		"DELETE_QUEUE",
		"MOVE_QUEUE",
		"PURGE_QUEUE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
