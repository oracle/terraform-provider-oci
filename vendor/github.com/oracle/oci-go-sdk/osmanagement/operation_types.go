// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesInstall   OperationTypesEnum = "INSTALL"
	OperationTypesUpdate    OperationTypesEnum = "UPDATE"
	OperationTypesRemove    OperationTypesEnum = "REMOVE"
	OperationTypesUpdateall OperationTypesEnum = "UPDATEALL"
)

var mappingOperationTypes = map[string]OperationTypesEnum{
	"INSTALL":   OperationTypesInstall,
	"UPDATE":    OperationTypesUpdate,
	"REMOVE":    OperationTypesRemove,
	"UPDATEALL": OperationTypesUpdateall,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypes {
		values = append(values, v)
	}
	return values
}
