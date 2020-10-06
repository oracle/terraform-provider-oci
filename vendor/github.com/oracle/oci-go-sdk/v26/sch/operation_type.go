// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
//

package sch

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateServiceConnector     OperationTypeEnum = "CREATE_SERVICE_CONNECTOR"
	OperationTypeUpdateServiceConnector     OperationTypeEnum = "UPDATE_SERVICE_CONNECTOR"
	OperationTypeDeleteServiceConnector     OperationTypeEnum = "DELETE_SERVICE_CONNECTOR"
	OperationTypeActivateServiceConnector   OperationTypeEnum = "ACTIVATE_SERVICE_CONNECTOR"
	OperationTypeDeactivateServiceConnector OperationTypeEnum = "DEACTIVATE_SERVICE_CONNECTOR"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"CREATE_SERVICE_CONNECTOR":     OperationTypeCreateServiceConnector,
	"UPDATE_SERVICE_CONNECTOR":     OperationTypeUpdateServiceConnector,
	"DELETE_SERVICE_CONNECTOR":     OperationTypeDeleteServiceConnector,
	"ACTIVATE_SERVICE_CONNECTOR":   OperationTypeActivateServiceConnector,
	"DEACTIVATE_SERVICE_CONNECTOR": OperationTypeDeactivateServiceConnector,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
