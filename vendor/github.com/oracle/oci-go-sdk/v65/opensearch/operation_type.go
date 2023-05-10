// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateOpensearchCluster                           OperationTypeEnum = "CREATE_OPENSEARCH_CLUSTER"
	OperationTypeUpdateOpensearchCluster                           OperationTypeEnum = "UPDATE_OPENSEARCH_CLUSTER"
	OperationTypeDeleteOpensearchCluster                           OperationTypeEnum = "DELETE_OPENSEARCH_CLUSTER"
	OperationTypeMoveOpensearchCluster                             OperationTypeEnum = "MOVE_OPENSEARCH_CLUSTER"
	OperationTypeRestoreOpensearchCluster                          OperationTypeEnum = "RESTORE_OPENSEARCH_CLUSTER"
	OperationTypeBackupOpensearchCluster                           OperationTypeEnum = "BACKUP_OPENSEARCH_CLUSTER"
	OperationTypeUpdateOpensearchClusterBackup                     OperationTypeEnum = "UPDATE_OPENSEARCH_CLUSTER_BACKUP"
	OperationTypeMoveOpensearchClusterBackup                       OperationTypeEnum = "MOVE_OPENSEARCH_CLUSTER_BACKUP"
	OperationTypeDeleteOpensearchClusterBackup                     OperationTypeEnum = "DELETE_OPENSEARCH_CLUSTER_BACKUP"
	OperationTypeUpdateOpensearchClusterSecurityConfig             OperationTypeEnum = "UPDATE_OPENSEARCH_CLUSTER_SECURITY_CONFIG"
	OperationTypeUpdateOpensearchCrossClusterConfig                OperationTypeEnum = "UPDATE_OPENSEARCH_CROSS_CLUSTER_CONFIG"
	OperationTypeUpdateOpensearchClusterReverseConnectionEndpoints OperationTypeEnum = "UPDATE_OPENSEARCH_CLUSTER_REVERSE_CONNECTION_ENDPOINTS"
	OperationTypeConfigureRemoteCluster                            OperationTypeEnum = "CONFIGURE_REMOTE_CLUSTER"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_OPENSEARCH_CLUSTER":                              OperationTypeCreateOpensearchCluster,
	"UPDATE_OPENSEARCH_CLUSTER":                              OperationTypeUpdateOpensearchCluster,
	"DELETE_OPENSEARCH_CLUSTER":                              OperationTypeDeleteOpensearchCluster,
	"MOVE_OPENSEARCH_CLUSTER":                                OperationTypeMoveOpensearchCluster,
	"RESTORE_OPENSEARCH_CLUSTER":                             OperationTypeRestoreOpensearchCluster,
	"BACKUP_OPENSEARCH_CLUSTER":                              OperationTypeBackupOpensearchCluster,
	"UPDATE_OPENSEARCH_CLUSTER_BACKUP":                       OperationTypeUpdateOpensearchClusterBackup,
	"MOVE_OPENSEARCH_CLUSTER_BACKUP":                         OperationTypeMoveOpensearchClusterBackup,
	"DELETE_OPENSEARCH_CLUSTER_BACKUP":                       OperationTypeDeleteOpensearchClusterBackup,
	"UPDATE_OPENSEARCH_CLUSTER_SECURITY_CONFIG":              OperationTypeUpdateOpensearchClusterSecurityConfig,
	"UPDATE_OPENSEARCH_CROSS_CLUSTER_CONFIG":                 OperationTypeUpdateOpensearchCrossClusterConfig,
	"UPDATE_OPENSEARCH_CLUSTER_REVERSE_CONNECTION_ENDPOINTS": OperationTypeUpdateOpensearchClusterReverseConnectionEndpoints,
	"CONFIGURE_REMOTE_CLUSTER":                               OperationTypeConfigureRemoteCluster,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_opensearch_cluster":                              OperationTypeCreateOpensearchCluster,
	"update_opensearch_cluster":                              OperationTypeUpdateOpensearchCluster,
	"delete_opensearch_cluster":                              OperationTypeDeleteOpensearchCluster,
	"move_opensearch_cluster":                                OperationTypeMoveOpensearchCluster,
	"restore_opensearch_cluster":                             OperationTypeRestoreOpensearchCluster,
	"backup_opensearch_cluster":                              OperationTypeBackupOpensearchCluster,
	"update_opensearch_cluster_backup":                       OperationTypeUpdateOpensearchClusterBackup,
	"move_opensearch_cluster_backup":                         OperationTypeMoveOpensearchClusterBackup,
	"delete_opensearch_cluster_backup":                       OperationTypeDeleteOpensearchClusterBackup,
	"update_opensearch_cluster_security_config":              OperationTypeUpdateOpensearchClusterSecurityConfig,
	"update_opensearch_cross_cluster_config":                 OperationTypeUpdateOpensearchCrossClusterConfig,
	"update_opensearch_cluster_reverse_connection_endpoints": OperationTypeUpdateOpensearchClusterReverseConnectionEndpoints,
	"configure_remote_cluster":                               OperationTypeConfigureRemoteCluster,
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
		"CREATE_OPENSEARCH_CLUSTER",
		"UPDATE_OPENSEARCH_CLUSTER",
		"DELETE_OPENSEARCH_CLUSTER",
		"MOVE_OPENSEARCH_CLUSTER",
		"RESTORE_OPENSEARCH_CLUSTER",
		"BACKUP_OPENSEARCH_CLUSTER",
		"UPDATE_OPENSEARCH_CLUSTER_BACKUP",
		"MOVE_OPENSEARCH_CLUSTER_BACKUP",
		"DELETE_OPENSEARCH_CLUSTER_BACKUP",
		"UPDATE_OPENSEARCH_CLUSTER_SECURITY_CONFIG",
		"UPDATE_OPENSEARCH_CROSS_CLUSTER_CONFIG",
		"UPDATE_OPENSEARCH_CLUSTER_REVERSE_CONNECTION_ENDPOINTS",
		"CONFIGURE_REMOTE_CLUSTER",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
