// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
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
	OperationTypeMigrateCluster       OperationTypeEnum = "MIGRATE_CLUSTER"
	OperationTypeClusterRollback      OperationTypeEnum = "CLUSTER_ROLLBACK"
	OperationTypeAttachOciCacheUsers  OperationTypeEnum = "ATTACH_OCI_CACHE_USERS"
	OperationTypeDetachOciCacheUsers  OperationTypeEnum = "DETACH_OCI_CACHE_USERS"
	OperationTypeCreateOciCacheUser   OperationTypeEnum = "CREATE_OCI_CACHE_USER"
	OperationTypeUpdateOciCacheUser   OperationTypeEnum = "UPDATE_OCI_CACHE_USER"
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
	"MIGRATE_CLUSTER":         OperationTypeMigrateCluster,
	"CLUSTER_ROLLBACK":        OperationTypeClusterRollback,
	"ATTACH_OCI_CACHE_USERS":  OperationTypeAttachOciCacheUsers,
	"DETACH_OCI_CACHE_USERS":  OperationTypeDetachOciCacheUsers,
	"CREATE_OCI_CACHE_USER":   OperationTypeCreateOciCacheUser,
	"UPDATE_OCI_CACHE_USER":   OperationTypeUpdateOciCacheUser,
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
	"migrate_cluster":         OperationTypeMigrateCluster,
	"cluster_rollback":        OperationTypeClusterRollback,
	"attach_oci_cache_users":  OperationTypeAttachOciCacheUsers,
	"detach_oci_cache_users":  OperationTypeDetachOciCacheUsers,
	"create_oci_cache_user":   OperationTypeCreateOciCacheUser,
	"update_oci_cache_user":   OperationTypeUpdateOciCacheUser,
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
		"MIGRATE_CLUSTER",
		"CLUSTER_ROLLBACK",
		"ATTACH_OCI_CACHE_USERS",
		"DETACH_OCI_CACHE_USERS",
		"CREATE_OCI_CACHE_USER",
		"UPDATE_OCI_CACHE_USER",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
