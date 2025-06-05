// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// PackageTypeEnum Enum with underlying type: string
type PackageTypeEnum string

// Set of constants representing the allowable values for PackageTypeEnum
const (
	PackageTypePackageFound      PackageTypeEnum = "PACKAGE_FOUND"
	PackageTypePackageMissing    PackageTypeEnum = "PACKAGE_MISSING"
	PackageTypePackageDependency PackageTypeEnum = "PACKAGE_DEPENDENCY"
)

var mappingPackageTypeEnum = map[string]PackageTypeEnum{
	"PACKAGE_FOUND":      PackageTypePackageFound,
	"PACKAGE_MISSING":    PackageTypePackageMissing,
	"PACKAGE_DEPENDENCY": PackageTypePackageDependency,
}

var mappingPackageTypeEnumLowerCase = map[string]PackageTypeEnum{
	"package_found":      PackageTypePackageFound,
	"package_missing":    PackageTypePackageMissing,
	"package_dependency": PackageTypePackageDependency,
}

// GetPackageTypeEnumValues Enumerates the set of values for PackageTypeEnum
func GetPackageTypeEnumValues() []PackageTypeEnum {
	values := make([]PackageTypeEnum, 0)
	for _, v := range mappingPackageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPackageTypeEnumStringValues Enumerates the set of values in String for PackageTypeEnum
func GetPackageTypeEnumStringValues() []string {
	return []string{
		"PACKAGE_FOUND",
		"PACKAGE_MISSING",
		"PACKAGE_DEPENDENCY",
	}
}

// GetMappingPackageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageTypeEnum(val string) (PackageTypeEnum, bool) {
	enum, ok := mappingPackageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
