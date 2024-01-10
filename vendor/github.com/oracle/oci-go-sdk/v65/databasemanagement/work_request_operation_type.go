// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateDbManagementPrivateEndpoint WorkRequestOperationTypeEnum = "CREATE_DB_MANAGEMENT_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeDeleteDbManagementPrivateEndpoint WorkRequestOperationTypeEnum = "DELETE_DB_MANAGEMENT_PRIVATE_ENDPOINT"
	WorkRequestOperationTypeCreateDbSystemDiscovery           WorkRequestOperationTypeEnum = "CREATE_DB_SYSTEM_DISCOVERY"
	WorkRequestOperationTypeCreateDbSystem                    WorkRequestOperationTypeEnum = "CREATE_DB_SYSTEM"
	WorkRequestOperationTypeUpdateDbSystem                    WorkRequestOperationTypeEnum = "UPDATE_DB_SYSTEM"
	WorkRequestOperationTypeDbSystemEnableDbmgmt              WorkRequestOperationTypeEnum = "DB_SYSTEM_ENABLE_DBMGMT"
	WorkRequestOperationTypeDbSystemDisableDbmgmt             WorkRequestOperationTypeEnum = "DB_SYSTEM_DISABLE_DBMGMT"
	WorkRequestOperationTypeDeleteDbSystem                    WorkRequestOperationTypeEnum = "DELETE_DB_SYSTEM"
	WorkRequestOperationTypeUpdateExternalDbSystemConnector   WorkRequestOperationTypeEnum = "UPDATE_EXTERNAL_DB_SYSTEM_CONNECTOR"
	WorkRequestOperationTypeChangeExternalDbSystemCompartment WorkRequestOperationTypeEnum = "CHANGE_EXTERNAL_DB_SYSTEM_COMPARTMENT"
	WorkRequestOperationTypeDisableExadataInfrasturcture      WorkRequestOperationTypeEnum = "DISABLE_EXADATA_INFRASTURCTURE"
	WorkRequestOperationTypeEnableExadataInfrastructure       WorkRequestOperationTypeEnum = "ENABLE_EXADATA_INFRASTRUCTURE"
	WorkRequestOperationTypeDeleteExadataInfrastructure       WorkRequestOperationTypeEnum = "DELETE_EXADATA_INFRASTRUCTURE"
	WorkRequestOperationTypeChangeExadataCompartment          WorkRequestOperationTypeEnum = "CHANGE_EXADATA_COMPARTMENT"
)

var mappingWorkRequestOperationTypeEnum = map[string]WorkRequestOperationTypeEnum{
	"CREATE_DB_MANAGEMENT_PRIVATE_ENDPOINT": WorkRequestOperationTypeCreateDbManagementPrivateEndpoint,
	"DELETE_DB_MANAGEMENT_PRIVATE_ENDPOINT": WorkRequestOperationTypeDeleteDbManagementPrivateEndpoint,
	"CREATE_DB_SYSTEM_DISCOVERY":            WorkRequestOperationTypeCreateDbSystemDiscovery,
	"CREATE_DB_SYSTEM":                      WorkRequestOperationTypeCreateDbSystem,
	"UPDATE_DB_SYSTEM":                      WorkRequestOperationTypeUpdateDbSystem,
	"DB_SYSTEM_ENABLE_DBMGMT":               WorkRequestOperationTypeDbSystemEnableDbmgmt,
	"DB_SYSTEM_DISABLE_DBMGMT":              WorkRequestOperationTypeDbSystemDisableDbmgmt,
	"DELETE_DB_SYSTEM":                      WorkRequestOperationTypeDeleteDbSystem,
	"UPDATE_EXTERNAL_DB_SYSTEM_CONNECTOR":   WorkRequestOperationTypeUpdateExternalDbSystemConnector,
	"CHANGE_EXTERNAL_DB_SYSTEM_COMPARTMENT": WorkRequestOperationTypeChangeExternalDbSystemCompartment,
	"DISABLE_EXADATA_INFRASTURCTURE":        WorkRequestOperationTypeDisableExadataInfrasturcture,
	"ENABLE_EXADATA_INFRASTRUCTURE":         WorkRequestOperationTypeEnableExadataInfrastructure,
	"DELETE_EXADATA_INFRASTRUCTURE":         WorkRequestOperationTypeDeleteExadataInfrastructure,
	"CHANGE_EXADATA_COMPARTMENT":            WorkRequestOperationTypeChangeExadataCompartment,
}

var mappingWorkRequestOperationTypeEnumLowerCase = map[string]WorkRequestOperationTypeEnum{
	"create_db_management_private_endpoint": WorkRequestOperationTypeCreateDbManagementPrivateEndpoint,
	"delete_db_management_private_endpoint": WorkRequestOperationTypeDeleteDbManagementPrivateEndpoint,
	"create_db_system_discovery":            WorkRequestOperationTypeCreateDbSystemDiscovery,
	"create_db_system":                      WorkRequestOperationTypeCreateDbSystem,
	"update_db_system":                      WorkRequestOperationTypeUpdateDbSystem,
	"db_system_enable_dbmgmt":               WorkRequestOperationTypeDbSystemEnableDbmgmt,
	"db_system_disable_dbmgmt":              WorkRequestOperationTypeDbSystemDisableDbmgmt,
	"delete_db_system":                      WorkRequestOperationTypeDeleteDbSystem,
	"update_external_db_system_connector":   WorkRequestOperationTypeUpdateExternalDbSystemConnector,
	"change_external_db_system_compartment": WorkRequestOperationTypeChangeExternalDbSystemCompartment,
	"disable_exadata_infrasturcture":        WorkRequestOperationTypeDisableExadataInfrasturcture,
	"enable_exadata_infrastructure":         WorkRequestOperationTypeEnableExadataInfrastructure,
	"delete_exadata_infrastructure":         WorkRequestOperationTypeDeleteExadataInfrastructure,
	"change_exadata_compartment":            WorkRequestOperationTypeChangeExadataCompartment,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DB_MANAGEMENT_PRIVATE_ENDPOINT",
		"DELETE_DB_MANAGEMENT_PRIVATE_ENDPOINT",
		"CREATE_DB_SYSTEM_DISCOVERY",
		"CREATE_DB_SYSTEM",
		"UPDATE_DB_SYSTEM",
		"DB_SYSTEM_ENABLE_DBMGMT",
		"DB_SYSTEM_DISABLE_DBMGMT",
		"DELETE_DB_SYSTEM",
		"UPDATE_EXTERNAL_DB_SYSTEM_CONNECTOR",
		"CHANGE_EXTERNAL_DB_SYSTEM_COMPARTMENT",
		"DISABLE_EXADATA_INFRASTURCTURE",
		"ENABLE_EXADATA_INFRASTRUCTURE",
		"DELETE_EXADATA_INFRASTRUCTURE",
		"CHANGE_EXADATA_COMPARTMENT",
	}
}

// GetMappingWorkRequestOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationTypeEnum(val string) (WorkRequestOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
