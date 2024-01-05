// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"strings"
)

// PackageTypeEnumEnum Enum with underlying type: string
type PackageTypeEnumEnum string

// Set of constants representing the allowable values for PackageTypeEnumEnum
const (
	PackageTypeEnumStack PackageTypeEnumEnum = "STACK"
)

var mappingPackageTypeEnumEnum = map[string]PackageTypeEnumEnum{
	"STACK": PackageTypeEnumStack,
}

var mappingPackageTypeEnumEnumLowerCase = map[string]PackageTypeEnumEnum{
	"stack": PackageTypeEnumStack,
}

// GetPackageTypeEnumEnumValues Enumerates the set of values for PackageTypeEnumEnum
func GetPackageTypeEnumEnumValues() []PackageTypeEnumEnum {
	values := make([]PackageTypeEnumEnum, 0)
	for _, v := range mappingPackageTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetPackageTypeEnumEnumStringValues Enumerates the set of values in String for PackageTypeEnumEnum
func GetPackageTypeEnumEnumStringValues() []string {
	return []string{
		"STACK",
	}
}

// GetMappingPackageTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageTypeEnumEnum(val string) (PackageTypeEnumEnum, bool) {
	enum, ok := mappingPackageTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
