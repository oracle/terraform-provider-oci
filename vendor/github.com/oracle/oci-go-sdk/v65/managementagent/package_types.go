// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// PackageTypesEnum Enum with underlying type: string
type PackageTypesEnum string

// Set of constants representing the allowable values for PackageTypesEnum
const (
	PackageTypesRpm PackageTypesEnum = "RPM"
	PackageTypesZip PackageTypesEnum = "ZIP"
)

var mappingPackageTypesEnum = map[string]PackageTypesEnum{
	"RPM": PackageTypesRpm,
	"ZIP": PackageTypesZip,
}

var mappingPackageTypesEnumLowerCase = map[string]PackageTypesEnum{
	"rpm": PackageTypesRpm,
	"zip": PackageTypesZip,
}

// GetPackageTypesEnumValues Enumerates the set of values for PackageTypesEnum
func GetPackageTypesEnumValues() []PackageTypesEnum {
	values := make([]PackageTypesEnum, 0)
	for _, v := range mappingPackageTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetPackageTypesEnumStringValues Enumerates the set of values in String for PackageTypesEnum
func GetPackageTypesEnumStringValues() []string {
	return []string{
		"RPM",
		"ZIP",
	}
}

// GetMappingPackageTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPackageTypesEnum(val string) (PackageTypesEnum, bool) {
	enum, ok := mappingPackageTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
