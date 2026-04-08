// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CpuArchTypeEnum Enum with underlying type: string
type CpuArchTypeEnum string

// Set of constants representing the allowable values for CpuArchTypeEnum
const (
	CpuArchTypeX8664   CpuArchTypeEnum = "X86_64"
	CpuArchTypeAarch64 CpuArchTypeEnum = "AARCH64"
)

var mappingCpuArchTypeEnum = map[string]CpuArchTypeEnum{
	"X86_64":  CpuArchTypeX8664,
	"AARCH64": CpuArchTypeAarch64,
}

var mappingCpuArchTypeEnumLowerCase = map[string]CpuArchTypeEnum{
	"x86_64":  CpuArchTypeX8664,
	"aarch64": CpuArchTypeAarch64,
}

// GetCpuArchTypeEnumValues Enumerates the set of values for CpuArchTypeEnum
func GetCpuArchTypeEnumValues() []CpuArchTypeEnum {
	values := make([]CpuArchTypeEnum, 0)
	for _, v := range mappingCpuArchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCpuArchTypeEnumStringValues Enumerates the set of values in String for CpuArchTypeEnum
func GetCpuArchTypeEnumStringValues() []string {
	return []string{
		"X86_64",
		"AARCH64",
	}
}

// GetMappingCpuArchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCpuArchTypeEnum(val string) (CpuArchTypeEnum, bool) {
	enum, ok := mappingCpuArchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
