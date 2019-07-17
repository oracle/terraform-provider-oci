// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

// WorkRequestOperationTypesEnum Enum with underlying type: string
type WorkRequestOperationTypesEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypesEnum
const (
	WorkRequestOperationTypesCreateWaasPolicy WorkRequestOperationTypesEnum = "CREATE_WAAS_POLICY"
	WorkRequestOperationTypesUpdateWaasPolicy WorkRequestOperationTypesEnum = "UPDATE_WAAS_POLICY"
	WorkRequestOperationTypesDeleteWaasPolicy WorkRequestOperationTypesEnum = "DELETE_WAAS_POLICY"
	WorkRequestOperationTypesPurgeWaasPolicy  WorkRequestOperationTypesEnum = "PURGE_WAAS_POLICY"
)

var mappingWorkRequestOperationTypes = map[string]WorkRequestOperationTypesEnum{
	"CREATE_WAAS_POLICY": WorkRequestOperationTypesCreateWaasPolicy,
	"UPDATE_WAAS_POLICY": WorkRequestOperationTypesUpdateWaasPolicy,
	"DELETE_WAAS_POLICY": WorkRequestOperationTypesDeleteWaasPolicy,
	"PURGE_WAAS_POLICY":  WorkRequestOperationTypesPurgeWaasPolicy,
}

// GetWorkRequestOperationTypesEnumValues Enumerates the set of values for WorkRequestOperationTypesEnum
func GetWorkRequestOperationTypesEnumValues() []WorkRequestOperationTypesEnum {
	values := make([]WorkRequestOperationTypesEnum, 0)
	for _, v := range mappingWorkRequestOperationTypes {
		values = append(values, v)
	}
	return values
}
