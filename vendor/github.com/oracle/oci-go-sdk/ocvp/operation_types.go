// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage the Oracle Cloud VMware Solution.
//

package ocvp

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateSddc     OperationTypesEnum = "CREATE_SDDC"
	OperationTypesDeleteSddc     OperationTypesEnum = "DELETE_SDDC"
	OperationTypesCreateEsxiHost OperationTypesEnum = "CREATE_ESXI_HOST"
	OperationTypesDeleteEsxiHost OperationTypesEnum = "DELETE_ESXI_HOST"
)

var mappingOperationTypes = map[string]OperationTypesEnum{
	"CREATE_SDDC":      OperationTypesCreateSddc,
	"DELETE_SDDC":      OperationTypesDeleteSddc,
	"CREATE_ESXI_HOST": OperationTypesCreateEsxiHost,
	"DELETE_ESXI_HOST": OperationTypesDeleteEsxiHost,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypes {
		values = append(values, v)
	}
	return values
}
