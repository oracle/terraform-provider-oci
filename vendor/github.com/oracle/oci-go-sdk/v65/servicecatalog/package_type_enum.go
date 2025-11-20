// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Use the Service Catalog API to manage solutions in Oracle Cloud Infrastructure Service Catalog.
// For more information, see Overview of Service Catalog (https://docs.oracle.com/iaas/Content/service-catalog/overview_of_service_catalog.htm).
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
	PackageTypeEnumImage PackageTypeEnumEnum = "IMAGE"
)

var mappingPackageTypeEnumEnum = map[string]PackageTypeEnumEnum{
	"STACK": PackageTypeEnumStack,
	"IMAGE": PackageTypeEnumImage,
}

var mappingPackageTypeEnumEnumLowerCase = map[string]PackageTypeEnumEnum{
	"stack": PackageTypeEnumStack,
	"image": PackageTypeEnumImage,
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
		"IMAGE",
	}
}

// GetMappingPackageTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageTypeEnumEnum(val string) (PackageTypeEnumEnum, bool) {
	enum, ok := mappingPackageTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
