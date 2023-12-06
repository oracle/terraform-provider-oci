// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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
	OperationTypesCreateBds                   OperationTypesEnum = "CREATE_BDS"
	OperationTypesUpdateBds                   OperationTypesEnum = "UPDATE_BDS"
	OperationTypesDeleteBds                   OperationTypesEnum = "DELETE_BDS"
	OperationTypesAddBlockStorage             OperationTypesEnum = "ADD_BLOCK_STORAGE"
	OperationTypesAddMasterNodes              OperationTypesEnum = "ADD_MASTER_NODES"
	OperationTypesAddUtilityNodes             OperationTypesEnum = "ADD_UTILITY_NODES"
	OperationTypesAddWorkerNodes              OperationTypesEnum = "ADD_WORKER_NODES"
	OperationTypesAddCloudSql                 OperationTypesEnum = "ADD_CLOUD_SQL"
	OperationTypesRemoveCloudSql              OperationTypesEnum = "REMOVE_CLOUD_SQL"
	OperationTypesChangeCompartmentForBds     OperationTypesEnum = "CHANGE_COMPARTMENT_FOR_BDS"
	OperationTypesChangeShape                 OperationTypesEnum = "CHANGE_SHAPE"
	OperationTypesUpdateInfra                 OperationTypesEnum = "UPDATE_INFRA"
	OperationTypesRestartNode                 OperationTypesEnum = "RESTART_NODE"
	OperationTypesRemoveNode                  OperationTypesEnum = "REMOVE_NODE"
	OperationTypesCreateAutoscaleConfig       OperationTypesEnum = "CREATE_AUTOSCALE_CONFIG"
	OperationTypesUpdateAutoscaleConfig       OperationTypesEnum = "UPDATE_AUTOSCALE_CONFIG"
	OperationTypesDeleteAutoscaleConfig       OperationTypesEnum = "DELETE_AUTOSCALE_CONFIG"
	OperationTypesAutoscaleConfig             OperationTypesEnum = "AUTOSCALE_CONFIG"
	OperationTypesAutoscaleRun                OperationTypesEnum = "AUTOSCALE_RUN"
	OperationTypesCreateApiKey                OperationTypesEnum = "CREATE_API_KEY"
	OperationTypesDeleteApiKey                OperationTypesEnum = "DELETE_API_KEY"
	OperationTypesTestObjectStoreConnection   OperationTypesEnum = "TEST_OBJECT_STORE_CONNECTION"
	OperationTypesCreateMetastoreConfig       OperationTypesEnum = "CREATE_METASTORE_CONFIG"
	OperationTypesDeleteMetastoreConfig       OperationTypesEnum = "DELETE_METASTORE_CONFIG"
	OperationTypesUpdateMetastoreConfig       OperationTypesEnum = "UPDATE_METASTORE_CONFIG"
	OperationTypesActivateMetastoreConfig     OperationTypesEnum = "ACTIVATE_METASTORE_CONFIG"
	OperationTypesTestMetastoreConfig         OperationTypesEnum = "TEST_METASTORE_CONFIG"
	OperationTypesPatchBds                    OperationTypesEnum = "PATCH_BDS"
	OperationTypesPatchOdh                    OperationTypesEnum = "PATCH_ODH"
	OperationTypesPatchOs                     OperationTypesEnum = "PATCH_OS"
	OperationTypesStopBds                     OperationTypesEnum = "STOP_BDS"
	OperationTypesStartBds                    OperationTypesEnum = "START_BDS"
	OperationTypesAddKafka                    OperationTypesEnum = "ADD_KAFKA"
	OperationTypesRemoveKafka                 OperationTypesEnum = "REMOVE_KAFKA"
	OperationTypesExecuteBootstrapScript      OperationTypesEnum = "EXECUTE_BOOTSTRAP_SCRIPT"
	OperationTypesOdhServiceCertificateUpdate OperationTypesEnum = "ODH_SERVICE_CERTIFICATE_UPDATE"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_BDS":                     OperationTypesCreateBds,
	"UPDATE_BDS":                     OperationTypesUpdateBds,
	"DELETE_BDS":                     OperationTypesDeleteBds,
	"ADD_BLOCK_STORAGE":              OperationTypesAddBlockStorage,
	"ADD_MASTER_NODES":               OperationTypesAddMasterNodes,
	"ADD_UTILITY_NODES":              OperationTypesAddUtilityNodes,
	"ADD_WORKER_NODES":               OperationTypesAddWorkerNodes,
	"ADD_CLOUD_SQL":                  OperationTypesAddCloudSql,
	"REMOVE_CLOUD_SQL":               OperationTypesRemoveCloudSql,
	"CHANGE_COMPARTMENT_FOR_BDS":     OperationTypesChangeCompartmentForBds,
	"CHANGE_SHAPE":                   OperationTypesChangeShape,
	"UPDATE_INFRA":                   OperationTypesUpdateInfra,
	"RESTART_NODE":                   OperationTypesRestartNode,
	"REMOVE_NODE":                    OperationTypesRemoveNode,
	"CREATE_AUTOSCALE_CONFIG":        OperationTypesCreateAutoscaleConfig,
	"UPDATE_AUTOSCALE_CONFIG":        OperationTypesUpdateAutoscaleConfig,
	"DELETE_AUTOSCALE_CONFIG":        OperationTypesDeleteAutoscaleConfig,
	"AUTOSCALE_CONFIG":               OperationTypesAutoscaleConfig,
	"AUTOSCALE_RUN":                  OperationTypesAutoscaleRun,
	"CREATE_API_KEY":                 OperationTypesCreateApiKey,
	"DELETE_API_KEY":                 OperationTypesDeleteApiKey,
	"TEST_OBJECT_STORE_CONNECTION":   OperationTypesTestObjectStoreConnection,
	"CREATE_METASTORE_CONFIG":        OperationTypesCreateMetastoreConfig,
	"DELETE_METASTORE_CONFIG":        OperationTypesDeleteMetastoreConfig,
	"UPDATE_METASTORE_CONFIG":        OperationTypesUpdateMetastoreConfig,
	"ACTIVATE_METASTORE_CONFIG":      OperationTypesActivateMetastoreConfig,
	"TEST_METASTORE_CONFIG":          OperationTypesTestMetastoreConfig,
	"PATCH_BDS":                      OperationTypesPatchBds,
	"PATCH_ODH":                      OperationTypesPatchOdh,
	"PATCH_OS":                       OperationTypesPatchOs,
	"STOP_BDS":                       OperationTypesStopBds,
	"START_BDS":                      OperationTypesStartBds,
	"ADD_KAFKA":                      OperationTypesAddKafka,
	"REMOVE_KAFKA":                   OperationTypesRemoveKafka,
	"EXECUTE_BOOTSTRAP_SCRIPT":       OperationTypesExecuteBootstrapScript,
	"ODH_SERVICE_CERTIFICATE_UPDATE": OperationTypesOdhServiceCertificateUpdate,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"create_bds":                     OperationTypesCreateBds,
	"update_bds":                     OperationTypesUpdateBds,
	"delete_bds":                     OperationTypesDeleteBds,
	"add_block_storage":              OperationTypesAddBlockStorage,
	"add_master_nodes":               OperationTypesAddMasterNodes,
	"add_utility_nodes":              OperationTypesAddUtilityNodes,
	"add_worker_nodes":               OperationTypesAddWorkerNodes,
	"add_cloud_sql":                  OperationTypesAddCloudSql,
	"remove_cloud_sql":               OperationTypesRemoveCloudSql,
	"change_compartment_for_bds":     OperationTypesChangeCompartmentForBds,
	"change_shape":                   OperationTypesChangeShape,
	"update_infra":                   OperationTypesUpdateInfra,
	"restart_node":                   OperationTypesRestartNode,
	"remove_node":                    OperationTypesRemoveNode,
	"create_autoscale_config":        OperationTypesCreateAutoscaleConfig,
	"update_autoscale_config":        OperationTypesUpdateAutoscaleConfig,
	"delete_autoscale_config":        OperationTypesDeleteAutoscaleConfig,
	"autoscale_config":               OperationTypesAutoscaleConfig,
	"autoscale_run":                  OperationTypesAutoscaleRun,
	"create_api_key":                 OperationTypesCreateApiKey,
	"delete_api_key":                 OperationTypesDeleteApiKey,
	"test_object_store_connection":   OperationTypesTestObjectStoreConnection,
	"create_metastore_config":        OperationTypesCreateMetastoreConfig,
	"delete_metastore_config":        OperationTypesDeleteMetastoreConfig,
	"update_metastore_config":        OperationTypesUpdateMetastoreConfig,
	"activate_metastore_config":      OperationTypesActivateMetastoreConfig,
	"test_metastore_config":          OperationTypesTestMetastoreConfig,
	"patch_bds":                      OperationTypesPatchBds,
	"patch_odh":                      OperationTypesPatchOdh,
	"patch_os":                       OperationTypesPatchOs,
	"stop_bds":                       OperationTypesStopBds,
	"start_bds":                      OperationTypesStartBds,
	"add_kafka":                      OperationTypesAddKafka,
	"remove_kafka":                   OperationTypesRemoveKafka,
	"execute_bootstrap_script":       OperationTypesExecuteBootstrapScript,
	"odh_service_certificate_update": OperationTypesOdhServiceCertificateUpdate,
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
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
