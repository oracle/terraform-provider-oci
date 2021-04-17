// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeEnableDatabaseInsight         OperationTypeEnum = "ENABLE_DATABASE_INSIGHT"
	OperationTypeDisableDatabaseInsight        OperationTypeEnum = "DISABLE_DATABASE_INSIGHT"
	OperationTypeUpdateDatabaseInsight         OperationTypeEnum = "UPDATE_DATABASE_INSIGHT"
	OperationTypeCreateDatabaseInsight         OperationTypeEnum = "CREATE_DATABASE_INSIGHT"
	OperationTypeMoveDatabaseInsight           OperationTypeEnum = "MOVE_DATABASE_INSIGHT"
	OperationTypeDeleteDatabaseInsight         OperationTypeEnum = "DELETE_DATABASE_INSIGHT"
	OperationTypeCreateEnterpriseManagerBridge OperationTypeEnum = "CREATE_ENTERPRISE_MANAGER_BRIDGE"
	OperationTypeUdpateEnterpriseManagerBridge OperationTypeEnum = "UDPATE_ENTERPRISE_MANAGER_BRIDGE"
	OperationTypeMoveEnterpriseManagerBridge   OperationTypeEnum = "MOVE_ENTERPRISE_MANAGER_BRIDGE"
	OperationTypeDeleteEnterpriseManagerBridge OperationTypeEnum = "DELETE_ENTERPRISE_MANAGER_BRIDGE"
	OperationTypeEnableHostInsight             OperationTypeEnum = "ENABLE_HOST_INSIGHT"
	OperationTypeDisableHostInsight            OperationTypeEnum = "DISABLE_HOST_INSIGHT"
	OperationTypeUpdateHostInsight             OperationTypeEnum = "UPDATE_HOST_INSIGHT"
	OperationTypeCreateHostInsight             OperationTypeEnum = "CREATE_HOST_INSIGHT"
	OperationTypeMoveHostInsight               OperationTypeEnum = "MOVE_HOST_INSIGHT"
	OperationTypeDeleteHostInsight             OperationTypeEnum = "DELETE_HOST_INSIGHT"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"ENABLE_DATABASE_INSIGHT":          OperationTypeEnableDatabaseInsight,
	"DISABLE_DATABASE_INSIGHT":         OperationTypeDisableDatabaseInsight,
	"UPDATE_DATABASE_INSIGHT":          OperationTypeUpdateDatabaseInsight,
	"CREATE_DATABASE_INSIGHT":          OperationTypeCreateDatabaseInsight,
	"MOVE_DATABASE_INSIGHT":            OperationTypeMoveDatabaseInsight,
	"DELETE_DATABASE_INSIGHT":          OperationTypeDeleteDatabaseInsight,
	"CREATE_ENTERPRISE_MANAGER_BRIDGE": OperationTypeCreateEnterpriseManagerBridge,
	"UDPATE_ENTERPRISE_MANAGER_BRIDGE": OperationTypeUdpateEnterpriseManagerBridge,
	"MOVE_ENTERPRISE_MANAGER_BRIDGE":   OperationTypeMoveEnterpriseManagerBridge,
	"DELETE_ENTERPRISE_MANAGER_BRIDGE": OperationTypeDeleteEnterpriseManagerBridge,
	"ENABLE_HOST_INSIGHT":              OperationTypeEnableHostInsight,
	"DISABLE_HOST_INSIGHT":             OperationTypeDisableHostInsight,
	"UPDATE_HOST_INSIGHT":              OperationTypeUpdateHostInsight,
	"CREATE_HOST_INSIGHT":              OperationTypeCreateHostInsight,
	"MOVE_HOST_INSIGHT":                OperationTypeMoveHostInsight,
	"DELETE_HOST_INSIGHT":              OperationTypeDeleteHostInsight,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
