// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// SystemModelEnumEnum Enum with underlying type: string
type SystemModelEnumEnum string

// Set of constants representing the allowable values for SystemModelEnumEnum
const (
	SystemModelEnumX11Ha768 SystemModelEnumEnum = "X11_HA_768"
	SystemModelEnumX10Ha512 SystemModelEnumEnum = "X10_HA_512"
	SystemModelEnumX8Ha384  SystemModelEnumEnum = "X8_HA_384"
)

var mappingSystemModelEnumEnum = map[string]SystemModelEnumEnum{
	"X11_HA_768": SystemModelEnumX11Ha768,
	"X10_HA_512": SystemModelEnumX10Ha512,
	"X8_HA_384":  SystemModelEnumX8Ha384,
}

var mappingSystemModelEnumEnumLowerCase = map[string]SystemModelEnumEnum{
	"x11_ha_768": SystemModelEnumX11Ha768,
	"x10_ha_512": SystemModelEnumX10Ha512,
	"x8_ha_384":  SystemModelEnumX8Ha384,
}

// GetSystemModelEnumEnumValues Enumerates the set of values for SystemModelEnumEnum
func GetSystemModelEnumEnumValues() []SystemModelEnumEnum {
	values := make([]SystemModelEnumEnum, 0)
	for _, v := range mappingSystemModelEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetSystemModelEnumEnumStringValues Enumerates the set of values in String for SystemModelEnumEnum
func GetSystemModelEnumEnumStringValues() []string {
	return []string{
		"X11_HA_768",
		"X10_HA_512",
		"X8_HA_384",
	}
}

// GetMappingSystemModelEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSystemModelEnumEnum(val string) (SystemModelEnumEnum, bool) {
	enum, ok := mappingSystemModelEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
