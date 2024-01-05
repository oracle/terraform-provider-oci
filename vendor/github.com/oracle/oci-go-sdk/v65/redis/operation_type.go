// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Redis Service API
//
// Use the Redis Service API to create and manage Redis clusters. A Redis cluster is a memory-based storage solution. For more information, see OCI Caching Service with Redis (https://docs.cloud.oracle.com/iaas/Content/redis/home.htm).
//

package redis

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateRedisCluster   OperationTypeEnum = "CREATE_REDIS_CLUSTER"
	OperationTypeUpdateRedisCluster   OperationTypeEnum = "UPDATE_REDIS_CLUSTER"
	OperationTypeDeleteRedisCluster   OperationTypeEnum = "DELETE_REDIS_CLUSTER"
	OperationTypeMoveRedisCluster     OperationTypeEnum = "MOVE_REDIS_CLUSTER"
	OperationTypeFailoverRedisCluster OperationTypeEnum = "FAILOVER_REDIS_CLUSTER"
	OperationTypeCreateRedisConfigSet OperationTypeEnum = "CREATE_REDIS_CONFIG_SET"
	OperationTypeUpdateRedisConfigSet OperationTypeEnum = "UPDATE_REDIS_CONFIG_SET"
	OperationTypeDeleteRedisConfigSet OperationTypeEnum = "DELETE_REDIS_CONFIG_SET"
	OperationTypeMoveRedisConfigSet   OperationTypeEnum = "MOVE_REDIS_CONFIG_SET"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_REDIS_CLUSTER":    OperationTypeCreateRedisCluster,
	"UPDATE_REDIS_CLUSTER":    OperationTypeUpdateRedisCluster,
	"DELETE_REDIS_CLUSTER":    OperationTypeDeleteRedisCluster,
	"MOVE_REDIS_CLUSTER":      OperationTypeMoveRedisCluster,
	"FAILOVER_REDIS_CLUSTER":  OperationTypeFailoverRedisCluster,
	"CREATE_REDIS_CONFIG_SET": OperationTypeCreateRedisConfigSet,
	"UPDATE_REDIS_CONFIG_SET": OperationTypeUpdateRedisConfigSet,
	"DELETE_REDIS_CONFIG_SET": OperationTypeDeleteRedisConfigSet,
	"MOVE_REDIS_CONFIG_SET":   OperationTypeMoveRedisConfigSet,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_redis_cluster":    OperationTypeCreateRedisCluster,
	"update_redis_cluster":    OperationTypeUpdateRedisCluster,
	"delete_redis_cluster":    OperationTypeDeleteRedisCluster,
	"move_redis_cluster":      OperationTypeMoveRedisCluster,
	"failover_redis_cluster":  OperationTypeFailoverRedisCluster,
	"create_redis_config_set": OperationTypeCreateRedisConfigSet,
	"update_redis_config_set": OperationTypeUpdateRedisConfigSet,
	"delete_redis_config_set": OperationTypeDeleteRedisConfigSet,
	"move_redis_config_set":   OperationTypeMoveRedisConfigSet,
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
		"CREATE_REDIS_CLUSTER",
		"UPDATE_REDIS_CLUSTER",
		"DELETE_REDIS_CLUSTER",
		"MOVE_REDIS_CLUSTER",
		"FAILOVER_REDIS_CLUSTER",
		"CREATE_REDIS_CONFIG_SET",
		"UPDATE_REDIS_CONFIG_SET",
		"DELETE_REDIS_CONFIG_SET",
		"MOVE_REDIS_CONFIG_SET",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
