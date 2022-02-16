// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	OperationTypesCreateBds                 OperationTypesEnum = "CREATE_BDS"
	OperationTypesUpdateBds                 OperationTypesEnum = "UPDATE_BDS"
	OperationTypesDeleteBds                 OperationTypesEnum = "DELETE_BDS"
	OperationTypesAddBlockStorage           OperationTypesEnum = "ADD_BLOCK_STORAGE"
	OperationTypesAddWorkerNodes            OperationTypesEnum = "ADD_WORKER_NODES"
	OperationTypesAddCloudSql               OperationTypesEnum = "ADD_CLOUD_SQL"
	OperationTypesRemoveCloudSql            OperationTypesEnum = "REMOVE_CLOUD_SQL"
	OperationTypesChangeCompartmentForBds   OperationTypesEnum = "CHANGE_COMPARTMENT_FOR_BDS"
	OperationTypesChangeShape               OperationTypesEnum = "CHANGE_SHAPE"
	OperationTypesUpdateInfra               OperationTypesEnum = "UPDATE_INFRA"
	OperationTypesRestartNode               OperationTypesEnum = "RESTART_NODE"
	OperationTypesAutoscaleConfig           OperationTypesEnum = "AUTOSCALE_CONFIG"
	OperationTypesAutoscaleRun              OperationTypesEnum = "AUTOSCALE_RUN"
	OperationTypesCreateApiKey              OperationTypesEnum = "CREATE_API_KEY"
	OperationTypesDeleteApiKey              OperationTypesEnum = "DELETE_API_KEY"
	OperationTypesTestObjectStoreConnection OperationTypesEnum = "TEST_OBJECT_STORE_CONNECTION"
	OperationTypesCreateMetastoreConfig     OperationTypesEnum = "CREATE_METASTORE_CONFIG"
	OperationTypesDeleteMetastoreConfig     OperationTypesEnum = "DELETE_METASTORE_CONFIG"
	OperationTypesUpdateMetastoreConfig     OperationTypesEnum = "UPDATE_METASTORE_CONFIG"
	OperationTypesActivateMetastoreConfig   OperationTypesEnum = "ACTIVATE_METASTORE_CONFIG"
	OperationTypesTestMetastoreConfig       OperationTypesEnum = "TEST_METASTORE_CONFIG"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_BDS":                   OperationTypesCreateBds,
	"UPDATE_BDS":                   OperationTypesUpdateBds,
	"DELETE_BDS":                   OperationTypesDeleteBds,
	"ADD_BLOCK_STORAGE":            OperationTypesAddBlockStorage,
	"ADD_WORKER_NODES":             OperationTypesAddWorkerNodes,
	"ADD_CLOUD_SQL":                OperationTypesAddCloudSql,
	"REMOVE_CLOUD_SQL":             OperationTypesRemoveCloudSql,
	"CHANGE_COMPARTMENT_FOR_BDS":   OperationTypesChangeCompartmentForBds,
	"CHANGE_SHAPE":                 OperationTypesChangeShape,
	"UPDATE_INFRA":                 OperationTypesUpdateInfra,
	"RESTART_NODE":                 OperationTypesRestartNode,
	"AUTOSCALE_CONFIG":             OperationTypesAutoscaleConfig,
	"AUTOSCALE_RUN":                OperationTypesAutoscaleRun,
	"CREATE_API_KEY":               OperationTypesCreateApiKey,
	"DELETE_API_KEY":               OperationTypesDeleteApiKey,
	"TEST_OBJECT_STORE_CONNECTION": OperationTypesTestObjectStoreConnection,
	"CREATE_METASTORE_CONFIG":      OperationTypesCreateMetastoreConfig,
	"DELETE_METASTORE_CONFIG":      OperationTypesDeleteMetastoreConfig,
	"UPDATE_METASTORE_CONFIG":      OperationTypesUpdateMetastoreConfig,
	"ACTIVATE_METASTORE_CONFIG":    OperationTypesActivateMetastoreConfig,
	"TEST_METASTORE_CONFIG":        OperationTypesTestMetastoreConfig,
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
		"ADD_WORKER_NODES",
		"ADD_CLOUD_SQL",
		"REMOVE_CLOUD_SQL",
		"CHANGE_COMPARTMENT_FOR_BDS",
		"CHANGE_SHAPE",
		"UPDATE_INFRA",
		"RESTART_NODE",
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
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	mappingOperationTypesEnumIgnoreCase := make(map[string]OperationTypesEnum)
	for k, v := range mappingOperationTypesEnum {
		mappingOperationTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperationTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
