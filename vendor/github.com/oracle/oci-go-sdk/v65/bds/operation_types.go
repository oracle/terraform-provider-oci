// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateBds                            OperationTypesEnum = "CREATE_BDS"
	OperationTypesUpdateBds                            OperationTypesEnum = "UPDATE_BDS"
	OperationTypesDeleteBds                            OperationTypesEnum = "DELETE_BDS"
	OperationTypesAddBlockStorage                      OperationTypesEnum = "ADD_BLOCK_STORAGE"
	OperationTypesAddMasterNodes                       OperationTypesEnum = "ADD_MASTER_NODES"
	OperationTypesAddUtilityNodes                      OperationTypesEnum = "ADD_UTILITY_NODES"
	OperationTypesAddWorkerNodes                       OperationTypesEnum = "ADD_WORKER_NODES"
	OperationTypesAddCloudSql                          OperationTypesEnum = "ADD_CLOUD_SQL"
	OperationTypesRemoveCloudSql                       OperationTypesEnum = "REMOVE_CLOUD_SQL"
	OperationTypesChangeCompartmentForBds              OperationTypesEnum = "CHANGE_COMPARTMENT_FOR_BDS"
	OperationTypesChangeShape                          OperationTypesEnum = "CHANGE_SHAPE"
	OperationTypesUpdateInfra                          OperationTypesEnum = "UPDATE_INFRA"
	OperationTypesRestartNode                          OperationTypesEnum = "RESTART_NODE"
	OperationTypesRemoveNode                           OperationTypesEnum = "REMOVE_NODE"
	OperationTypesCreateAutoscaleConfig                OperationTypesEnum = "CREATE_AUTOSCALE_CONFIG"
	OperationTypesUpdateAutoscaleConfig                OperationTypesEnum = "UPDATE_AUTOSCALE_CONFIG"
	OperationTypesDeleteAutoscaleConfig                OperationTypesEnum = "DELETE_AUTOSCALE_CONFIG"
	OperationTypesAutoscaleConfig                      OperationTypesEnum = "AUTOSCALE_CONFIG"
	OperationTypesAutoscaleRun                         OperationTypesEnum = "AUTOSCALE_RUN"
	OperationTypesCreateApiKey                         OperationTypesEnum = "CREATE_API_KEY"
	OperationTypesDeleteApiKey                         OperationTypesEnum = "DELETE_API_KEY"
	OperationTypesTestObjectStoreConnection            OperationTypesEnum = "TEST_OBJECT_STORE_CONNECTION"
	OperationTypesCreateMetastoreConfig                OperationTypesEnum = "CREATE_METASTORE_CONFIG"
	OperationTypesDeleteMetastoreConfig                OperationTypesEnum = "DELETE_METASTORE_CONFIG"
	OperationTypesUpdateMetastoreConfig                OperationTypesEnum = "UPDATE_METASTORE_CONFIG"
	OperationTypesActivateMetastoreConfig              OperationTypesEnum = "ACTIVATE_METASTORE_CONFIG"
	OperationTypesTestMetastoreConfig                  OperationTypesEnum = "TEST_METASTORE_CONFIG"
	OperationTypesPatchBds                             OperationTypesEnum = "PATCH_BDS"
	OperationTypesPatchOdh                             OperationTypesEnum = "PATCH_ODH"
	OperationTypesPatchOs                              OperationTypesEnum = "PATCH_OS"
	OperationTypesStopBds                              OperationTypesEnum = "STOP_BDS"
	OperationTypesStartBds                             OperationTypesEnum = "START_BDS"
	OperationTypesAddKafka                             OperationTypesEnum = "ADD_KAFKA"
	OperationTypesRemoveKafka                          OperationTypesEnum = "REMOVE_KAFKA"
	OperationTypesExecuteBootstrapScript               OperationTypesEnum = "EXECUTE_BOOTSTRAP_SCRIPT"
	OperationTypesOdhServiceCertificateUpdate          OperationTypesEnum = "ODH_SERVICE_CERTIFICATE_UPDATE"
	OperationTypesCreateBackupConfig                   OperationTypesEnum = "CREATE_BACKUP_CONFIG"
	OperationTypesUpdateBackupConfig                   OperationTypesEnum = "UPDATE_BACKUP_CONFIG"
	OperationTypesDeleteBackupConfig                   OperationTypesEnum = "DELETE_BACKUP_CONFIG"
	OperationTypesCreateNodeReplaceConfig              OperationTypesEnum = "CREATE_NODE_REPLACE_CONFIG"
	OperationTypesUpdateNodeReplaceConfig              OperationTypesEnum = "UPDATE_NODE_REPLACE_CONFIG"
	OperationTypesDeleteNodeReplaceConfig              OperationTypesEnum = "DELETE_NODE_REPLACE_CONFIG"
	OperationTypesBackupNodes                          OperationTypesEnum = "BACKUP_NODES"
	OperationTypesDeleteNodeBackup                     OperationTypesEnum = "DELETE_NODE_BACKUP"
	OperationTypesReplaceNode                          OperationTypesEnum = "REPLACE_NODE"
	OperationTypesCreateResourcePrincipalConfiguration OperationTypesEnum = "CREATE_RESOURCE_PRINCIPAL_CONFIGURATION"
	OperationTypesDeleteResourcePrincipalConfiguration OperationTypesEnum = "DELETE_RESOURCE_PRINCIPAL_CONFIGURATION"
	OperationTypesUpdateResourcePrincipalConfiguration OperationTypesEnum = "UPDATE_RESOURCE_PRINCIPAL_CONFIGURATION"
	OperationTypesRefreshResourcePrincipal             OperationTypesEnum = "REFRESH_RESOURCE_PRINCIPAL"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_BDS":                              OperationTypesCreateBds,
	"UPDATE_BDS":                              OperationTypesUpdateBds,
	"DELETE_BDS":                              OperationTypesDeleteBds,
	"ADD_BLOCK_STORAGE":                       OperationTypesAddBlockStorage,
	"ADD_MASTER_NODES":                        OperationTypesAddMasterNodes,
	"ADD_UTILITY_NODES":                       OperationTypesAddUtilityNodes,
	"ADD_WORKER_NODES":                        OperationTypesAddWorkerNodes,
	"ADD_CLOUD_SQL":                           OperationTypesAddCloudSql,
	"REMOVE_CLOUD_SQL":                        OperationTypesRemoveCloudSql,
	"CHANGE_COMPARTMENT_FOR_BDS":              OperationTypesChangeCompartmentForBds,
	"CHANGE_SHAPE":                            OperationTypesChangeShape,
	"UPDATE_INFRA":                            OperationTypesUpdateInfra,
	"RESTART_NODE":                            OperationTypesRestartNode,
	"REMOVE_NODE":                             OperationTypesRemoveNode,
	"CREATE_AUTOSCALE_CONFIG":                 OperationTypesCreateAutoscaleConfig,
	"UPDATE_AUTOSCALE_CONFIG":                 OperationTypesUpdateAutoscaleConfig,
	"DELETE_AUTOSCALE_CONFIG":                 OperationTypesDeleteAutoscaleConfig,
	"AUTOSCALE_CONFIG":                        OperationTypesAutoscaleConfig,
	"AUTOSCALE_RUN":                           OperationTypesAutoscaleRun,
	"CREATE_API_KEY":                          OperationTypesCreateApiKey,
	"DELETE_API_KEY":                          OperationTypesDeleteApiKey,
	"TEST_OBJECT_STORE_CONNECTION":            OperationTypesTestObjectStoreConnection,
	"CREATE_METASTORE_CONFIG":                 OperationTypesCreateMetastoreConfig,
	"DELETE_METASTORE_CONFIG":                 OperationTypesDeleteMetastoreConfig,
	"UPDATE_METASTORE_CONFIG":                 OperationTypesUpdateMetastoreConfig,
	"ACTIVATE_METASTORE_CONFIG":               OperationTypesActivateMetastoreConfig,
	"TEST_METASTORE_CONFIG":                   OperationTypesTestMetastoreConfig,
	"PATCH_BDS":                               OperationTypesPatchBds,
	"PATCH_ODH":                               OperationTypesPatchOdh,
	"PATCH_OS":                                OperationTypesPatchOs,
	"STOP_BDS":                                OperationTypesStopBds,
	"START_BDS":                               OperationTypesStartBds,
	"ADD_KAFKA":                               OperationTypesAddKafka,
	"REMOVE_KAFKA":                            OperationTypesRemoveKafka,
	"EXECUTE_BOOTSTRAP_SCRIPT":                OperationTypesExecuteBootstrapScript,
	"ODH_SERVICE_CERTIFICATE_UPDATE":          OperationTypesOdhServiceCertificateUpdate,
	"CREATE_BACKUP_CONFIG":                    OperationTypesCreateBackupConfig,
	"UPDATE_BACKUP_CONFIG":                    OperationTypesUpdateBackupConfig,
	"DELETE_BACKUP_CONFIG":                    OperationTypesDeleteBackupConfig,
	"CREATE_NODE_REPLACE_CONFIG":              OperationTypesCreateNodeReplaceConfig,
	"UPDATE_NODE_REPLACE_CONFIG":              OperationTypesUpdateNodeReplaceConfig,
	"DELETE_NODE_REPLACE_CONFIG":              OperationTypesDeleteNodeReplaceConfig,
	"BACKUP_NODES":                            OperationTypesBackupNodes,
	"DELETE_NODE_BACKUP":                      OperationTypesDeleteNodeBackup,
	"REPLACE_NODE":                            OperationTypesReplaceNode,
	"CREATE_RESOURCE_PRINCIPAL_CONFIGURATION": OperationTypesCreateResourcePrincipalConfiguration,
	"DELETE_RESOURCE_PRINCIPAL_CONFIGURATION": OperationTypesDeleteResourcePrincipalConfiguration,
	"UPDATE_RESOURCE_PRINCIPAL_CONFIGURATION": OperationTypesUpdateResourcePrincipalConfiguration,
	"REFRESH_RESOURCE_PRINCIPAL":              OperationTypesRefreshResourcePrincipal,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"create_bds":                              OperationTypesCreateBds,
	"update_bds":                              OperationTypesUpdateBds,
	"delete_bds":                              OperationTypesDeleteBds,
	"add_block_storage":                       OperationTypesAddBlockStorage,
	"add_master_nodes":                        OperationTypesAddMasterNodes,
	"add_utility_nodes":                       OperationTypesAddUtilityNodes,
	"add_worker_nodes":                        OperationTypesAddWorkerNodes,
	"add_cloud_sql":                           OperationTypesAddCloudSql,
	"remove_cloud_sql":                        OperationTypesRemoveCloudSql,
	"change_compartment_for_bds":              OperationTypesChangeCompartmentForBds,
	"change_shape":                            OperationTypesChangeShape,
	"update_infra":                            OperationTypesUpdateInfra,
	"restart_node":                            OperationTypesRestartNode,
	"remove_node":                             OperationTypesRemoveNode,
	"create_autoscale_config":                 OperationTypesCreateAutoscaleConfig,
	"update_autoscale_config":                 OperationTypesUpdateAutoscaleConfig,
	"delete_autoscale_config":                 OperationTypesDeleteAutoscaleConfig,
	"autoscale_config":                        OperationTypesAutoscaleConfig,
	"autoscale_run":                           OperationTypesAutoscaleRun,
	"create_api_key":                          OperationTypesCreateApiKey,
	"delete_api_key":                          OperationTypesDeleteApiKey,
	"test_object_store_connection":            OperationTypesTestObjectStoreConnection,
	"create_metastore_config":                 OperationTypesCreateMetastoreConfig,
	"delete_metastore_config":                 OperationTypesDeleteMetastoreConfig,
	"update_metastore_config":                 OperationTypesUpdateMetastoreConfig,
	"activate_metastore_config":               OperationTypesActivateMetastoreConfig,
	"test_metastore_config":                   OperationTypesTestMetastoreConfig,
	"patch_bds":                               OperationTypesPatchBds,
	"patch_odh":                               OperationTypesPatchOdh,
	"patch_os":                                OperationTypesPatchOs,
	"stop_bds":                                OperationTypesStopBds,
	"start_bds":                               OperationTypesStartBds,
	"add_kafka":                               OperationTypesAddKafka,
	"remove_kafka":                            OperationTypesRemoveKafka,
	"execute_bootstrap_script":                OperationTypesExecuteBootstrapScript,
	"odh_service_certificate_update":          OperationTypesOdhServiceCertificateUpdate,
	"create_backup_config":                    OperationTypesCreateBackupConfig,
	"update_backup_config":                    OperationTypesUpdateBackupConfig,
	"delete_backup_config":                    OperationTypesDeleteBackupConfig,
	"create_node_replace_config":              OperationTypesCreateNodeReplaceConfig,
	"update_node_replace_config":              OperationTypesUpdateNodeReplaceConfig,
	"delete_node_replace_config":              OperationTypesDeleteNodeReplaceConfig,
	"backup_nodes":                            OperationTypesBackupNodes,
	"delete_node_backup":                      OperationTypesDeleteNodeBackup,
	"replace_node":                            OperationTypesReplaceNode,
	"create_resource_principal_configuration": OperationTypesCreateResourcePrincipalConfiguration,
	"delete_resource_principal_configuration": OperationTypesDeleteResourcePrincipalConfiguration,
	"update_resource_principal_configuration": OperationTypesUpdateResourcePrincipalConfiguration,
	"refresh_resource_principal":              OperationTypesRefreshResourcePrincipal,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypesEnumStringValues Enumerates the set of values in String for OperationTypesEnum
func GetOperationTypesEnumStringValues() []string {
	return []string{
		"CREATE_BDS",
		"UPDATE_BDS",
		"DELETE_BDS",
		"ADD_BLOCK_STORAGE",
		"ADD_MASTER_NODES",
		"ADD_UTILITY_NODES",
		"ADD_WORKER_NODES",
		"ADD_CLOUD_SQL",
		"REMOVE_CLOUD_SQL",
		"CHANGE_COMPARTMENT_FOR_BDS",
		"CHANGE_SHAPE",
		"UPDATE_INFRA",
		"RESTART_NODE",
		"REMOVE_NODE",
		"CREATE_AUTOSCALE_CONFIG",
		"UPDATE_AUTOSCALE_CONFIG",
		"DELETE_AUTOSCALE_CONFIG",
		"AUTOSCALE_CONFIG",
		"AUTOSCALE_RUN",
		"CREATE_API_KEY",
		"DELETE_API_KEY",
		"TEST_OBJECT_STORE_CONNECTION",
		"CREATE_METASTORE_CONFIG",
		"DELETE_METASTORE_CONFIG",
		"UPDATE_METASTORE_CONFIG",
		"ACTIVATE_METASTORE_CONFIG",
		"TEST_METASTORE_CONFIG",
		"PATCH_BDS",
		"PATCH_ODH",
		"PATCH_OS",
		"STOP_BDS",
		"START_BDS",
		"ADD_KAFKA",
		"REMOVE_KAFKA",
		"EXECUTE_BOOTSTRAP_SCRIPT",
		"ODH_SERVICE_CERTIFICATE_UPDATE",
		"CREATE_BACKUP_CONFIG",
		"UPDATE_BACKUP_CONFIG",
		"DELETE_BACKUP_CONFIG",
		"CREATE_NODE_REPLACE_CONFIG",
		"UPDATE_NODE_REPLACE_CONFIG",
		"DELETE_NODE_REPLACE_CONFIG",
		"BACKUP_NODES",
		"DELETE_NODE_BACKUP",
		"REPLACE_NODE",
		"CREATE_RESOURCE_PRINCIPAL_CONFIGURATION",
		"DELETE_RESOURCE_PRINCIPAL_CONFIGURATION",
		"UPDATE_RESOURCE_PRINCIPAL_CONFIGURATION",
		"REFRESH_RESOURCE_PRINCIPAL",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
