// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesInstall                    OperationTypesEnum = "INSTALL"
	OperationTypesUpdate                     OperationTypesEnum = "UPDATE"
	OperationTypesRemove                     OperationTypesEnum = "REMOVE"
	OperationTypesUpdateall                  OperationTypesEnum = "UPDATEALL"
	OperationTypesEnablemodulestream         OperationTypesEnum = "ENABLEMODULESTREAM"
	OperationTypesDisablemodulestream        OperationTypesEnum = "DISABLEMODULESTREAM"
	OperationTypesSwitchmodulestream         OperationTypesEnum = "SWITCHMODULESTREAM"
	OperationTypesInstallmodulestreamprofile OperationTypesEnum = "INSTALLMODULESTREAMPROFILE"
	OperationTypesRemovemodulestreamprofile  OperationTypesEnum = "REMOVEMODULESTREAMPROFILE"
	OperationTypesCompound                   OperationTypesEnum = "COMPOUND"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"INSTALL":                    OperationTypesInstall,
	"UPDATE":                     OperationTypesUpdate,
	"REMOVE":                     OperationTypesRemove,
	"UPDATEALL":                  OperationTypesUpdateall,
	"ENABLEMODULESTREAM":         OperationTypesEnablemodulestream,
	"DISABLEMODULESTREAM":        OperationTypesDisablemodulestream,
	"SWITCHMODULESTREAM":         OperationTypesSwitchmodulestream,
	"INSTALLMODULESTREAMPROFILE": OperationTypesInstallmodulestreamprofile,
	"REMOVEMODULESTREAMPROFILE":  OperationTypesRemovemodulestreamprofile,
	"COMPOUND":                   OperationTypesCompound,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"install":                    OperationTypesInstall,
	"update":                     OperationTypesUpdate,
	"remove":                     OperationTypesRemove,
	"updateall":                  OperationTypesUpdateall,
	"enablemodulestream":         OperationTypesEnablemodulestream,
	"disablemodulestream":        OperationTypesDisablemodulestream,
	"switchmodulestream":         OperationTypesSwitchmodulestream,
	"installmodulestreamprofile": OperationTypesInstallmodulestreamprofile,
	"removemodulestreamprofile":  OperationTypesRemovemodulestreamprofile,
	"compound":                   OperationTypesCompound,
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
		"INSTALL",
		"UPDATE",
		"REMOVE",
		"UPDATEALL",
		"ENABLEMODULESTREAM",
		"DISABLEMODULESTREAM",
		"SWITCHMODULESTREAM",
		"INSTALLMODULESTREAMPROFILE",
		"REMOVEMODULESTREAMPROFILE",
		"COMPOUND",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
