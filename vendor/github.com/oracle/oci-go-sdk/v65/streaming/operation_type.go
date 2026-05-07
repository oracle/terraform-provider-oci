// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.oracle.com/iaas/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateStream         OperationTypeEnum = "CREATE_STREAM"
	OperationTypeUpdateStream         OperationTypeEnum = "UPDATE_STREAM"
	OperationTypeDeleteStream         OperationTypeEnum = "DELETE_STREAM"
	OperationTypeCreateStreamPool     OperationTypeEnum = "CREATE_STREAM_POOL"
	OperationTypeUpdateStreamPool     OperationTypeEnum = "UPDATE_STREAM_POOL"
	OperationTypeDeleteStreamPool     OperationTypeEnum = "DELETE_STREAM_POOL"
	OperationTypeCreateConnectHarness OperationTypeEnum = "CREATE_CONNECT_HARNESS"
	OperationTypeUpdateConnectHarness OperationTypeEnum = "UPDATE_CONNECT_HARNESS"
	OperationTypeDeleteConnectHarness OperationTypeEnum = "DELETE_CONNECT_HARNESS"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_STREAM":          OperationTypeCreateStream,
	"UPDATE_STREAM":          OperationTypeUpdateStream,
	"DELETE_STREAM":          OperationTypeDeleteStream,
	"CREATE_STREAM_POOL":     OperationTypeCreateStreamPool,
	"UPDATE_STREAM_POOL":     OperationTypeUpdateStreamPool,
	"DELETE_STREAM_POOL":     OperationTypeDeleteStreamPool,
	"CREATE_CONNECT_HARNESS": OperationTypeCreateConnectHarness,
	"UPDATE_CONNECT_HARNESS": OperationTypeUpdateConnectHarness,
	"DELETE_CONNECT_HARNESS": OperationTypeDeleteConnectHarness,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_stream":          OperationTypeCreateStream,
	"update_stream":          OperationTypeUpdateStream,
	"delete_stream":          OperationTypeDeleteStream,
	"create_stream_pool":     OperationTypeCreateStreamPool,
	"update_stream_pool":     OperationTypeUpdateStreamPool,
	"delete_stream_pool":     OperationTypeDeleteStreamPool,
	"create_connect_harness": OperationTypeCreateConnectHarness,
	"update_connect_harness": OperationTypeUpdateConnectHarness,
	"delete_connect_harness": OperationTypeDeleteConnectHarness,
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
		"CREATE_STREAM",
		"UPDATE_STREAM",
		"DELETE_STREAM",
		"CREATE_STREAM_POOL",
		"UPDATE_STREAM_POOL",
		"DELETE_STREAM_POOL",
		"CREATE_CONNECT_HARNESS",
		"UPDATE_CONNECT_HARNESS",
		"DELETE_CONNECT_HARNESS",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
