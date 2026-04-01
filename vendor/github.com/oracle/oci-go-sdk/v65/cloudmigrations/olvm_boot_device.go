// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"strings"
)

// OlvmBootDeviceEnum Enum with underlying type: string
type OlvmBootDeviceEnum string

// Set of constants representing the allowable values for OlvmBootDeviceEnum
const (
	OlvmBootDeviceCdrom   OlvmBootDeviceEnum = "CDROM"
	OlvmBootDeviceHd      OlvmBootDeviceEnum = "HD"
	OlvmBootDeviceNetwork OlvmBootDeviceEnum = "NETWORK"
)

var mappingOlvmBootDeviceEnum = map[string]OlvmBootDeviceEnum{
	"CDROM":   OlvmBootDeviceCdrom,
	"HD":      OlvmBootDeviceHd,
	"NETWORK": OlvmBootDeviceNetwork,
}

var mappingOlvmBootDeviceEnumLowerCase = map[string]OlvmBootDeviceEnum{
	"cdrom":   OlvmBootDeviceCdrom,
	"hd":      OlvmBootDeviceHd,
	"network": OlvmBootDeviceNetwork,
}

// GetOlvmBootDeviceEnumValues Enumerates the set of values for OlvmBootDeviceEnum
func GetOlvmBootDeviceEnumValues() []OlvmBootDeviceEnum {
	values := make([]OlvmBootDeviceEnum, 0)
	for _, v := range mappingOlvmBootDeviceEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmBootDeviceEnumStringValues Enumerates the set of values in String for OlvmBootDeviceEnum
func GetOlvmBootDeviceEnumStringValues() []string {
	return []string{
		"CDROM",
		"HD",
		"NETWORK",
	}
}

// GetMappingOlvmBootDeviceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmBootDeviceEnum(val string) (OlvmBootDeviceEnum, bool) {
	enum, ok := mappingOlvmBootDeviceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
