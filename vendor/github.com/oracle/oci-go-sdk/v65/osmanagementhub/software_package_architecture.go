// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// SoftwarePackageArchitectureEnum Enum with underlying type: string
type SoftwarePackageArchitectureEnum string

// Set of constants representing the allowable values for SoftwarePackageArchitectureEnum
const (
	SoftwarePackageArchitectureI386    SoftwarePackageArchitectureEnum = "I386"
	SoftwarePackageArchitectureI686    SoftwarePackageArchitectureEnum = "I686"
	SoftwarePackageArchitectureAarch64 SoftwarePackageArchitectureEnum = "AARCH64"
	SoftwarePackageArchitectureX8664   SoftwarePackageArchitectureEnum = "X86_64"
	SoftwarePackageArchitectureSrc     SoftwarePackageArchitectureEnum = "SRC"
	SoftwarePackageArchitectureNoarch  SoftwarePackageArchitectureEnum = "NOARCH"
	SoftwarePackageArchitectureOther   SoftwarePackageArchitectureEnum = "OTHER"
)

var mappingSoftwarePackageArchitectureEnum = map[string]SoftwarePackageArchitectureEnum{
	"I386":    SoftwarePackageArchitectureI386,
	"I686":    SoftwarePackageArchitectureI686,
	"AARCH64": SoftwarePackageArchitectureAarch64,
	"X86_64":  SoftwarePackageArchitectureX8664,
	"SRC":     SoftwarePackageArchitectureSrc,
	"NOARCH":  SoftwarePackageArchitectureNoarch,
	"OTHER":   SoftwarePackageArchitectureOther,
}

var mappingSoftwarePackageArchitectureEnumLowerCase = map[string]SoftwarePackageArchitectureEnum{
	"i386":    SoftwarePackageArchitectureI386,
	"i686":    SoftwarePackageArchitectureI686,
	"aarch64": SoftwarePackageArchitectureAarch64,
	"x86_64":  SoftwarePackageArchitectureX8664,
	"src":     SoftwarePackageArchitectureSrc,
	"noarch":  SoftwarePackageArchitectureNoarch,
	"other":   SoftwarePackageArchitectureOther,
}

// GetSoftwarePackageArchitectureEnumValues Enumerates the set of values for SoftwarePackageArchitectureEnum
func GetSoftwarePackageArchitectureEnumValues() []SoftwarePackageArchitectureEnum {
	values := make([]SoftwarePackageArchitectureEnum, 0)
	for _, v := range mappingSoftwarePackageArchitectureEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwarePackageArchitectureEnumStringValues Enumerates the set of values in String for SoftwarePackageArchitectureEnum
func GetSoftwarePackageArchitectureEnumStringValues() []string {
	return []string{
		"I386",
		"I686",
		"AARCH64",
		"X86_64",
		"SRC",
		"NOARCH",
		"OTHER",
	}
}

// GetMappingSoftwarePackageArchitectureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwarePackageArchitectureEnum(val string) (SoftwarePackageArchitectureEnum, bool) {
	enum, ok := mappingSoftwarePackageArchitectureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
