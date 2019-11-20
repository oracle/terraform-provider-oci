// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
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
