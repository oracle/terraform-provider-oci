// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

// PackageUpdateTypesEnum Enum with underlying type: string
type PackageUpdateTypesEnum string

// Set of constants representing the allowable values for PackageUpdateTypesEnum
const (
	PackageUpdateTypesSecurity    PackageUpdateTypesEnum = "SECURITY"
	PackageUpdateTypesBugfix      PackageUpdateTypesEnum = "BUGFIX"
	PackageUpdateTypesEnhancement PackageUpdateTypesEnum = "ENHANCEMENT"
	PackageUpdateTypesAll         PackageUpdateTypesEnum = "ALL"
)

var mappingPackageUpdateTypes = map[string]PackageUpdateTypesEnum{
	"SECURITY":    PackageUpdateTypesSecurity,
	"BUGFIX":      PackageUpdateTypesBugfix,
	"ENHANCEMENT": PackageUpdateTypesEnhancement,
	"ALL":         PackageUpdateTypesAll,
}

// GetPackageUpdateTypesEnumValues Enumerates the set of values for PackageUpdateTypesEnum
func GetPackageUpdateTypesEnumValues() []PackageUpdateTypesEnum {
	values := make([]PackageUpdateTypesEnum, 0)
	for _, v := range mappingPackageUpdateTypes {
		values = append(values, v)
	}
	return values
}
