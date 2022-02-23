// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Control Plane API
//
// Use the Application Performance Monitoring Control Plane API to perform operations such as creating, updating,
// deleting and listing APM domains and monitoring the progress of these operations using the work request APIs. For more information, see Application Performance Monitoring (https://docs.cloud.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmcontrolplane

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateApmDomain  OperationTypesEnum = "CREATE_APM_DOMAIN"
	OperationTypesUpdateApmDomain  OperationTypesEnum = "UPDATE_APM_DOMAIN"
	OperationTypesDeleteApmDomain  OperationTypesEnum = "DELETE_APM_DOMAIN"
	OperationTypesGenerateDataKeys OperationTypesEnum = "GENERATE_DATA_KEYS"
	OperationTypesRemoveDataKeys   OperationTypesEnum = "REMOVE_DATA_KEYS"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_APM_DOMAIN":  OperationTypesCreateApmDomain,
	"UPDATE_APM_DOMAIN":  OperationTypesUpdateApmDomain,
	"DELETE_APM_DOMAIN":  OperationTypesDeleteApmDomain,
	"GENERATE_DATA_KEYS": OperationTypesGenerateDataKeys,
	"REMOVE_DATA_KEYS":   OperationTypesRemoveDataKeys,
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
		"CREATE_APM_DOMAIN",
		"UPDATE_APM_DOMAIN",
		"DELETE_APM_DOMAIN",
		"GENERATE_DATA_KEYS",
		"REMOVE_DATA_KEYS",
	}
}
