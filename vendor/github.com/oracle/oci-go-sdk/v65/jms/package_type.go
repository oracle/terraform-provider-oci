// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// PackageTypeEnum Enum with underlying type: string
type PackageTypeEnum string

// Set of constants representing the allowable values for PackageTypeEnum
const (
	PackageTypeRpm PackageTypeEnum = "RPM"
	PackageTypeZip PackageTypeEnum = "ZIP"
)

var mappingPackageTypeEnum = map[string]PackageTypeEnum{
	"RPM": PackageTypeRpm,
	"ZIP": PackageTypeZip,
}

var mappingPackageTypeEnumLowerCase = map[string]PackageTypeEnum{
	"rpm": PackageTypeRpm,
	"zip": PackageTypeZip,
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
		"RPM",
		"ZIP",
	}
}

// GetMappingPackageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageTypeEnum(val string) (PackageTypeEnum, bool) {
	enum, ok := mappingPackageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
