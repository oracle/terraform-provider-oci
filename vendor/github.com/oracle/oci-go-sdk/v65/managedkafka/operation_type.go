// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Streaming with Apache Kafka (OSAK) API
//
// Use Oracle Streaming with Apache Kafka Control Plane API to create/update/delete managed Kafka clusters.
//

package managedkafka

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateKafkaCluster OperationTypeEnum = "CREATE_KAFKA_CLUSTER"
	OperationTypeUpdateKafkaCluster OperationTypeEnum = "UPDATE_KAFKA_CLUSTER"
	OperationTypeDeleteKafkaCluster OperationTypeEnum = "DELETE_KAFKA_CLUSTER"
	OperationTypeMoveKafkaCluster   OperationTypeEnum = "MOVE_KAFKA_CLUSTER"
	OperationTypeEnableSuperuser    OperationTypeEnum = "ENABLE_SUPERUSER"
	OperationTypeDisableSuperuser   OperationTypeEnum = "DISABLE_SUPERUSER"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_KAFKA_CLUSTER": OperationTypeCreateKafkaCluster,
	"UPDATE_KAFKA_CLUSTER": OperationTypeUpdateKafkaCluster,
	"DELETE_KAFKA_CLUSTER": OperationTypeDeleteKafkaCluster,
	"MOVE_KAFKA_CLUSTER":   OperationTypeMoveKafkaCluster,
	"ENABLE_SUPERUSER":     OperationTypeEnableSuperuser,
	"DISABLE_SUPERUSER":    OperationTypeDisableSuperuser,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_kafka_cluster": OperationTypeCreateKafkaCluster,
	"update_kafka_cluster": OperationTypeUpdateKafkaCluster,
	"delete_kafka_cluster": OperationTypeDeleteKafkaCluster,
	"move_kafka_cluster":   OperationTypeMoveKafkaCluster,
	"enable_superuser":     OperationTypeEnableSuperuser,
	"disable_superuser":    OperationTypeDisableSuperuser,
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
		"CREATE_KAFKA_CLUSTER",
		"UPDATE_KAFKA_CLUSTER",
		"DELETE_KAFKA_CLUSTER",
		"MOVE_KAFKA_CLUSTER",
		"ENABLE_SUPERUSER",
		"DISABLE_SUPERUSER",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
