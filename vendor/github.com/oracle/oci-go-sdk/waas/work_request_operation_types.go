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
	WorkRequestOperationTypesCreateWaasPolicy           WorkRequestOperationTypesEnum = "CREATE_WAAS_POLICY"
	WorkRequestOperationTypesUpdateWaasPolicy           WorkRequestOperationTypesEnum = "UPDATE_WAAS_POLICY"
	WorkRequestOperationTypesDeleteWaasPolicy           WorkRequestOperationTypesEnum = "DELETE_WAAS_POLICY"
	WorkRequestOperationTypesPurgeWaasPolicyCache       WorkRequestOperationTypesEnum = "PURGE_WAAS_POLICY_CACHE"
	WorkRequestOperationTypesCreateCustomProtectionRule WorkRequestOperationTypesEnum = "CREATE_CUSTOM_PROTECTION_RULE"
	WorkRequestOperationTypesUpdateCustomProtectionRule WorkRequestOperationTypesEnum = "UPDATE_CUSTOM_PROTECTION_RULE"
	WorkRequestOperationTypesDeleteCustomProtectionRule WorkRequestOperationTypesEnum = "DELETE_CUSTOM_PROTECTION_RULE"
)

var mappingWorkRequestOperationTypes = map[string]WorkRequestOperationTypesEnum{
	"CREATE_WAAS_POLICY":            WorkRequestOperationTypesCreateWaasPolicy,
	"UPDATE_WAAS_POLICY":            WorkRequestOperationTypesUpdateWaasPolicy,
	"DELETE_WAAS_POLICY":            WorkRequestOperationTypesDeleteWaasPolicy,
	"PURGE_WAAS_POLICY_CACHE":       WorkRequestOperationTypesPurgeWaasPolicyCache,
	"CREATE_CUSTOM_PROTECTION_RULE": WorkRequestOperationTypesCreateCustomProtectionRule,
	"UPDATE_CUSTOM_PROTECTION_RULE": WorkRequestOperationTypesUpdateCustomProtectionRule,
	"DELETE_CUSTOM_PROTECTION_RULE": WorkRequestOperationTypesDeleteCustomProtectionRule,
}

// GetWorkRequestOperationTypesEnumValues Enumerates the set of values for WorkRequestOperationTypesEnum
func GetWorkRequestOperationTypesEnumValues() []WorkRequestOperationTypesEnum {
	values := make([]WorkRequestOperationTypesEnum, 0)
	for _, v := range mappingWorkRequestOperationTypes {
		values = append(values, v)
	}
	return values
}
